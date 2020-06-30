//go:generate mapstructure-to-hcl2 -type Config

package exoscale

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer/common"
	pkrconfig "github.com/hashicorp/packer/helper/config"
	"github.com/hashicorp/packer/packer"
)

var defaultAPIEndpoint = "https://api.exoscale.com/v1"

type Config struct {
	common.PackerConfig `mapstructure:",squash"`
	SkipClean           bool `mapstructure:"skip_clean"`

	SOSEndpoint             string `mapstructure:"sos_endpoint"`
	APIEndpoint             string `mapstructure:"api_endpoint"`
	APIKey                  string `mapstructure:"api_key"`
	APISecret               string `mapstructure:"api_secret"`
	ImageBucket             string `mapstructure:"image_bucket"`
	TemplateZone            string `mapstructure:"template_zone"`
	TemplateName            string `mapstructure:"template_name"`
	TemplateDescription     string `mapstructure:"template_description"`
	TemplateUsername        string `mapstructure:"template_username"`
	TemplateDisablePassword bool   `mapstructure:"template_disable_password"`
	TemplateDisableSSHKey   bool   `mapstructure:"template_disable_sshkey"`
}

func NewConfig(raws ...interface{}) (*Config, error) {
	var config = Config{
		APIEndpoint: defaultAPIEndpoint,
	}

	if err := pkrconfig.Decode(&config, nil, raws...); err != nil {
		return nil, err
	}

	requiredArgs := map[string]*string{
		"api_key":       &config.APIKey,
		"api_secret":    &config.APISecret,
		"api_endpoint":  &config.APIEndpoint,
		"image_bucket":  &config.ImageBucket,
		"template_zone": &config.TemplateZone,
		"template_name": &config.TemplateName,
	}

	errs := new(packer.MultiError)
	for k, v := range requiredArgs {
		if *v == "" {
			errs = packer.MultiErrorAppend(
				errs, fmt.Errorf("%s must be set", k))
		}
	}

	if config.SOSEndpoint == "" {
		config.SOSEndpoint = "https://sos-" + config.TemplateZone + ".exo.io"
	}

	if len(errs.Errors) > 0 {
		return nil, errs
	}

	return &config, nil
}

// ConfigSpec returns HCL object spec
func (p *PostProcessor) ConfigSpec() hcldec.ObjectSpec {
	return p.config.FlatMapstructure().HCL2Spec()
}
