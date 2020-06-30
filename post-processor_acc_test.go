package exoscale

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	pkrbuilderqemu "github.com/hashicorp/packer/builder/qemu"
	"github.com/hashicorp/packer/packer"
	"github.com/stretchr/testify/require"
)

var (
	testAccImageBucket         = "eat-template-images"
	testAccTemplateName        = "test-packer-builder-exoscale"
	testAccTemplateZone        = "ch-dk-2"
	testAccTemplateDescription = "Built with Packer"
	testAccTemplateUsername    = "packer"
	testAccImageFile           = "./testdata/test-packer-post-processor-exoscale-import.qcow2"
)

type testMockArtifact struct {
	files []string
}

func (a *testMockArtifact) BuilderId() string          { return pkrbuilderqemu.BuilderId }
func (a *testMockArtifact) Files() []string            { return a.files }
func (a *testMockArtifact) Id() string                 { return "" }
func (a *testMockArtifact) String() string             { return "" }
func (a *testMockArtifact) State(_ string) interface{} { return nil }
func (a *testMockArtifact) Destroy() error             { return nil }

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("EXOSCALE_API_KEY"); v == "" {
		t.Fatal("EXOSCALE_API_KEY must be set for acceptance tests")
	}

	if v := os.Getenv("EXOSCALE_API_SECRET"); v == "" {
		t.Fatal("EXOSCALE_API_SECRET must be set for acceptance tests")
	}
}

func testAccCheckPostProcessorArtifact(t *testing.T, artifact packer.Artifact) error {
	var a = artifact.(*Artifact)

	if a.template.ID == nil {
		return fmt.Errorf("artifact template ID is not set")
	}

	if a.template.ZoneName != testAccTemplateZone {
		return fmt.Errorf("expected template zone %q, got %q",
			testAccTemplateZone,
			a.template.ZoneName)
	}

	if a.template.Name != testAccTemplateName {
		return fmt.Errorf("expected template name %q, got %q",
			testAccTemplateName,
			a.template.Name)
	}

	if a.template.DisplayText != testAccTemplateDescription {
		return fmt.Errorf("expected template description %q, got %q",
			testAccTemplateDescription,
			a.template.DisplayText)
	}

	if username, ok := a.template.Details["username"]; !ok {
		return errors.New("artifact username not set")
	} else if username != testAccTemplateUsername {
		return fmt.Errorf("expected template username %q, got %q",
			testAccTemplateUsername,
			username)
	}

	return nil
}

func TestAccPostProcessor(t *testing.T) {
	var postProcessor PostProcessor

	testAccPreCheck(t)

	err := postProcessor.Configure([]interface{}{map[string]interface{}{
		"api_key":              os.Getenv("EXOSCALE_API_KEY"),
		"api_secret":           os.Getenv("EXOSCALE_API_SECRET"),
		"image_bucket":         testAccImageBucket,
		"template_zone":        testAccTemplateZone,
		"template_name":        testAccTemplateName,
		"template_description": testAccTemplateDescription,
		"template_username":    testAccTemplateUsername,
	}}...)
	require.NoError(t, err)

	artifact, _, _, err := postProcessor.PostProcess(
		context.Background(),
		&packer.NoopUi{},
		&testMockArtifact{files: []string{testAccImageFile}})
	require.NoError(t, err)

	err = testAccCheckPostProcessorArtifact(t, artifact)
	require.NoError(t, err)

	err = artifact.Destroy()
	require.NoError(t, err)
}
