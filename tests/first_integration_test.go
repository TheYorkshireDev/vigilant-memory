//go:build integration

package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestRepoNamingModule(t *testing.T) {
	product := "aaa"
	environment := "bbb"

	expectedName := "aaa-bbb"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "fixtures",
		Vars: map[string]interface{}{
			"product":     product,
			"environment": environment,
		},
	})

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// Apply the infrastructure
	terraform.InitAndApply(t, terraformOptions)
	output := terraform.Output(t, terraformOptions, "repository_name")
	assert.Equal(t, expectedName, output)

	// Run perpetual diff
	planResult := terraform.Plan(t, terratestOptions)

	// Make sure the plan shows zero changes
	assert.Contains(t, planResult, "No changes.")
}
