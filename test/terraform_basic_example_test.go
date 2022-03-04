package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformBasicExample(t *testing.T) {

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/terraform-basic-example",

		Logger: logger.Discard
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.Init(t, terraformOptions)
}
