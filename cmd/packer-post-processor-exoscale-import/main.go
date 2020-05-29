package main

import (
	"log"

	exoscaleimport "github.com/exoscale/packer-post-processor-exoscale-import"
	"github.com/hashicorp/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		log.Fatal(err)
	}

	if err := server.RegisterPostProcessor(new(exoscaleimport.PostProcessor)); err != nil {
		log.Fatal(err)
	}

	server.Serve()
}
