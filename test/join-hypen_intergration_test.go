//go:build integration

package test

import (
	"path/filepath"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestJoinHypenDeployment(t *testing.T) {
	t.Parallel()

	_examplesDir := "../../examples"
	exampleDir := filepath.Join(_fixturesDir, "join-hypen")

	input_one := "one"
	input_two := "two"
	expected_output := "one-two"

	terratestOptions := &terraform.Options{
		TerraformDir: exampleDir,
		NoColor:      true,

		Vars: map[string]interface{}{
			"example_input_one": input_one,
			"example_input_two": input_two,
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer test_structure.RunTestStage(t, "teardown", func() {
		terraform.Destroy(t, terratestOptions)
	})

	// Apply the infrastructure
	test_structure.RunTestStage(t, "apply", func() {
		// Save the terraform oprions for future reference
		terraform.InitAndApply(t, terratestOptions)
	})

	// Apply the infrastructure
	test_structure.RunTestStage(t, "apply", func() {
		terraform.InitAndApply(t, terratestOptions)
	})

	// Run perpetual diff
	test_structure.RunTestStage(t, "perpetual_diff", func() {
		planResult := terraform.Plan(t, terratestOptions)

		// Make sure the plan shows zero changes
		assert.Contains(t, planResult, "No changes.")
	})
}
