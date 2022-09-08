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

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
	output := terraform.Output(t, terraformOptions, "repository_name")
	assert.Equal(t, expectedName, output)
}
