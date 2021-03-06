package exoscale

import (
	"context"
	"errors"
	"fmt"

	"github.com/exoscale/egoscale"
	"github.com/hashicorp/packer/builder/file"
	"github.com/hashicorp/packer/builder/qemu"
	"github.com/hashicorp/packer/common"
	"github.com/hashicorp/packer/helper/multistep"
	"github.com/hashicorp/packer/packer"
	"github.com/hashicorp/packer/post-processor/artifice"
	"github.com/hashicorp/packer/version"
)

func init() {
	egoscale.UserAgent = "Exoscale-Packer-Post-Processor/" + version.FormattedVersion() + " " + egoscale.UserAgent
}

type PostProcessor struct {
	config *Config
	runner multistep.Runner
	exo    *egoscale.Client
}

func (p *PostProcessor) Configure(raws ...interface{}) error {
	config, err := NewConfig(raws...)
	if err != nil {
		return err
	}
	p.config = config

	packer.LogSecretFilter.Set(p.config.APIKey, p.config.APISecret)

	return nil
}

func (p *PostProcessor) PostProcess(ctx context.Context, ui packer.Ui, a packer.Artifact) (packer.Artifact, bool, bool, error) {
	switch a.BuilderId() {
	case qemu.BuilderId, file.BuilderId, artifice.BuilderId:
		break
	default:
		err := fmt.Errorf("unsupported artifact type %q: this post-processor only imports "+
			"artifacts from QEMU/file builders and Artifice post-processor", a.BuilderId())
		return nil, false, false, err
	}

	p.exo = egoscale.NewClient(p.config.APIEndpoint, p.config.APIKey, p.config.APISecret)

	state := new(multistep.BasicStateBag)
	state.Put("config", p.config)
	state.Put("exo", p.exo)
	state.Put("ui", ui)
	state.Put("artifact", a)

	steps := []multistep.Step{
		new(stepUploadImage),
		new(stepRegisterTemplate),
		new(stepDeleteImage),
	}

	p.runner = common.NewRunnerWithPauseFn(steps, p.config.PackerConfig, ui, state)
	p.runner.Run(ctx, state)

	if rawErr, ok := state.GetOk("error"); ok {
		return nil, false, false, rawErr.(error)
	}

	if _, ok := state.GetOk(multistep.StateCancelled); ok {
		return nil, false, false, errors.New("post-processing cancelled")
	}

	if _, ok := state.GetOk(multistep.StateHalted); ok {
		return nil, false, false, errors.New("post-processing halted")
	}

	v, ok := state.GetOk("template")
	if !ok {
		return nil, false, false, errors.New("unable to find template in state")
	}

	return &Artifact{
		template: v.(egoscale.Template),
		exo:      p.exo,
	}, false, false, nil
}
