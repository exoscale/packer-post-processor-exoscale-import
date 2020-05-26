package main

import (
	exoscaleimport "github.com/exoscale/packer-post-processor-exoscale-import"
	"github.com/hashicorp/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}

	if err = server.RegisterPostProcessor(new(exoscaleimport.PostProcessor)); err != nil {
		panic(err)
	}

	server.Serve()
}
