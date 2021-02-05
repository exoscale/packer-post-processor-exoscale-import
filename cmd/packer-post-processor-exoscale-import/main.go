package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/packer-plugin-sdk/plugin"

	exoscaleimport "github.com/exoscale/packer-post-processor-exoscale-import"
)

func main() {
	ps := plugin.NewSet()
	ps.RegisterPostProcessor(plugin.DEFAULT_NAME, new(exoscaleimport.PostProcessor))

	if err := ps.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
