//go:build unit

package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestRepoNamingModuleUnit(t *testing.T) {
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

	// Plan the infrastructure
	planResult := terraform.InitAndPlan(t, terraformOptions)

	// Make sure the plan shows expected name
	assert.Contains(t, planResult, expectedName)
}
