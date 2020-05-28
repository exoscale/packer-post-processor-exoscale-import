# Packer Post-Processor plugin: Exoscale Import

[![Actions Status](https://github.com/exoscale/packer-post-processor-exoscale-import/workflows/CI/badge.svg)](https://github.com/exoscale/packer-post-processor-exoscale-import/actions?query=workflow%3ACI)

This is Exoscale import plugin to be used with HashiCorp [Packer](https://www.packer.io/). This plugin takes an image artifact from the QEMU, Artifice, or File builders and imports it to Exoscale.

This guide assumes you have already used Packer and have a basic understanding of how Packer works. Otherwise, first read this guide on how to [get started with Packer](https://www.packer.io/intro/getting-started/).

## Installation

### Using pre-built releases

You can find pre-built releases of the plugin here. Once you have downloaded the latest archive corresponding to your target OS, uncompress it to retrieve the `packer-post-processor-exoscale-import` plugin binary file.

### From Sources

If you prefer to build the plugin from sources, clone the GitHub repository locally and run the command `make build` from the root of the sources directory. Upon successful compilation, a `packer-post-processor-exoscale-import` plugin binary file can be found in the `/bin` directory.

## Configuration

to install the plugin in Packer, please follow this guideline to [installing a Packer plugin](https://www.packer.io/docs/extending/plugins/#installing-plugins)

There are some configuration options available for the post-processor.

Required:

- `api_key` (string) - The API key used to communicate with Exoscale
  services. This may also be set using the `EXOSCALE_API_KEY` environmental
  variable.

- `api_secret` (string) - The API secret used to communicate with Exoscale
  services. This may also be set using the `EXOSCALE_API_SECRET`
  environmental variable.

- `image_bucket` (string) - The name of the bucket in which to upload the
  template image to SOS. The bucket must exist when the post-processor is
  run.

- `template_name` (string) - The name to be used for registering the template.

- `template_description` (string) - The description for the registered template.

Optional:

- `api_endpoint` (string) - The API endpoint used to communicate with the
  Exoscale API. Defaults to `https://api.exoscale.com/compute`.

- `sos_endpoint` (string) - The endpoint used to communicate with SOS.
  Defaults to `https://sos-ch-gva-2.exo.io`.

- `template_zone` (string) - The Exoscale [zone](https://www.exoscale.com/datacenters/)
  in which to register the template. Defaults to `ch-gva-2`.

- `template_username` (string) - An optional username to be used to log into
  Compute instances using this template.

- `template_disable_password` (boolean) - Whether the registered template
  should disable Compute instance password reset. Defaults to `false`.

- `template_disable_sshkey` (boolean) - Whether the registered template
  should disable SSH key installation during Compute instance creation.
  Defaults to `false`.

- `skip_clean` (boolean) - Whether we should skip removing the image file
  uploaded to SOS after the import process has completed. "true" means that
  we should leave it in the bucket, "false" means deleting it.
  Defaults to `false`.

## Usage

### Basic Example

```json
"post-processors": [
        {
            "type": "exoscale-import",
            "api_key": "{{user `exoscale_api_key`}}",
            "api_secret": "{{user `exoscale_api_secret`}}",
            "api_endpoint": "https://api.exoscale.com/v1",
            "image_bucket": "joh-cloud",
            "template_name": "hello",
            "template_description": "My hello application",
            "template_username": "ubuntu"
        }
    ]
```

Please follow this post to discover [How to creating custom templates using Packer](https://www.exoscale.com/syslog/creating-custom-templates-using-packer/)