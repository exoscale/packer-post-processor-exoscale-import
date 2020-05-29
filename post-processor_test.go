package exoscaleimport

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testSOSEndpoint             = "https://sos." + testTemplateZone + ".exo.io"
	testAPIEndpoint             = "https://api.exoscale.com/v1"
	testAPIKey                  = "EXOabcdef0123456789abcdef01"
	testAPISecret               = "ABCDEFGHIJKLMNOPRQSTUVWXYZ0123456789abcdefg"
	testImageBucket             = "my_bucket"
	testTemplateZone            = "ch-dk-2"
	testTemplateName            = "test"
	testTemplateDescription     = "test description"
	testTemplateUsername        = "test"
	testTemplateDisablePassword = true
	testTemplateDisableSSHKey   = true
	testSkipClean               = true
)

func testConfig() map[string]interface{} {
	return map[string]interface{}{
		"api_key":              testAPIKey,
		"api_secret":           testAPISecret,
		"image_bucket":         testImageBucket,
		"template_name":        testTemplateName,
		"template_description": testTemplateDescription,
	}
}

func testConfigOptionalOptions() map[string]interface{} {
	return map[string]interface{}{
		"api_key":                   testAPIKey,
		"api_secret":                testAPISecret,
		"image_bucket":              testImageBucket,
		"template_name":             testTemplateName,
		"template_description":      testTemplateDescription,
		"sos_endpoint":              testSOSEndpoint,
		"api_endpoint":              testAPIEndpoint,
		"template_username":         testTemplateUsername,
		"template_disable_password": testTemplateDisablePassword,
		"template_disable_sshkey":   testTemplateDisableSSHKey,
		"skip_clean":                testSkipClean,
	}
}

func TestConfigureWithRequiredOptions(t *testing.T) {
	p := new(PostProcessor)
	config := testConfig()

	err := p.Configure(config)

	require.NoError(t, err)
}

func TestConfigureWithoutRequiredOptions(t *testing.T) {
	p := new(PostProcessor)

	err := p.Configure(map[string]interface{}{})

	require.Error(t, err)
}

func TestConfigureWithoutAPIKeyOption(t *testing.T) {
	p := new(PostProcessor)
	config := testConfig()

	delete(config, "api_key")

	err := p.Configure(config)

	require.Error(t, err)
}

func TestConfigureWithoutAPISecretOption(t *testing.T) {
	p := new(PostProcessor)
	config := testConfig()

	delete(config, "api_secret")

	err := p.Configure(config)

	require.Error(t, err)
}
func TestConfigureWithoutImageBucketOption(t *testing.T) {
	p := new(PostProcessor)
	config := testConfig()

	delete(config, "image_bucket")

	err := p.Configure(config)

	require.Error(t, err)
}
func TestConfigureWithoutTemplateNameOption(t *testing.T) {
	p := new(PostProcessor)
	config := testConfig()

	delete(config, "template_name")

	err := p.Configure(config)

	require.Error(t, err)
}
func TestConfigureWithoutTemplateDescriptionOption(t *testing.T) {
	p := new(PostProcessor)
	config := testConfig()

	delete(config, "template_description")

	err := p.Configure(config)

	require.Error(t, err)
}

func TestConfigureWithbadOption(t *testing.T) {
	p := new(PostProcessor)
	config := testConfig()

	config["test"] = "test"

	err := p.Configure(config)

	require.Error(t, err)
}

func TestConfigureWithOptionalOptions(t *testing.T) {
	p := new(PostProcessor)
	config := testConfigOptionalOptions()

	err := p.Configure(config)

	require.NoError(t, err)
}
