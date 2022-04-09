//go:build unit

package test

import (
	"path/filepath"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestStaticSiteValidity(t *testing.T) {
	t.Parallel()

	_fixturesDir := "./fixtures"
	exampleDir := filepath.Join(_fixturesDir, "join-hypen")

	terratestOptions := &terraform.Options{
		TerraformDir: exampleDir,
		NoColor:      true,
	}

	t.Logf("Running in %s", exampleDir)
	output := terraform.InitAndPlan(t, terratestOptions)

	assert.Contains(t, output, "No changes. Your infrastructure matches the configuration.")
}
