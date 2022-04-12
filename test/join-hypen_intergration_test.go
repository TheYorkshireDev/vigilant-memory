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

	_examplesDir := "../examples"
	exampleDir := filepath.Join(_examplesDir, "join-hypen")

	input_one := "one"
	input_two := "two"
	expected_output := "joined_name = \"one-two\""

	terratestOptions := &terraform.Options{
		TerraformDir: exampleDir,
		NoColor:      true,

		Vars: map[string]interface{}{
			"example_input_one": input_one,
			"example_input_two": input_two,
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terratestOptions)

	// Apply the infrastructure
	applyResult := terraform.InitAndApply(t, terratestOptions)
	assert.Contains(t, applyResult, expected_output)

	// Run perpetual diff
	planResult := terraform.Plan(t, terratestOptions)

	// Make sure the plan shows zero changes
	assert.Contains(t, planResult, "No changes.")
}

func TestJoinHypenDeployment1(t *testing.T) {
	t.Parallel()

	_examplesDir := "../examples"
	exampleDir := filepath.Join(_examplesDir, "join-hypen")

	input_one := "one"
	input_two := "two"
	expected_output := "joined_name = \"one-two\""

	terratestOptions := &terraform.Options{
		TerraformDir: exampleDir,
		NoColor:      true,

		Vars: map[string]interface{}{
			"example_input_one": input_one,
			"example_input_two": input_two,
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terratestOptions)

	// Apply the infrastructure
	applyResult := terraform.InitAndApply(t, terratestOptions)
	assert.Contains(t, applyResult, expected_output)

	// Run perpetual diff
	planResult := terraform.Plan(t, terratestOptions)

	// Make sure the plan shows zero changes
	assert.Contains(t, planResult, "No changes.")
}

func TestJoinHypenDeployment2(t *testing.T) {
	t.Parallel()

	_examplesDir := "../examples"
	exampleDir := filepath.Join(_examplesDir, "join-hypen")

	input_one := "one"
	input_two := "two"
	expected_output := "joined_name = \"one-two\""

	terratestOptions := &terraform.Options{
		TerraformDir: exampleDir,
		NoColor:      true,

		Vars: map[string]interface{}{
			"example_input_one": input_one,
			"example_input_two": input_two,
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terratestOptions)

	// Apply the infrastructure
	applyResult := terraform.InitAndApply(t, terratestOptions)
	assert.Contains(t, applyResult, expected_output)

	// Run perpetual diff
	planResult := terraform.Plan(t, terratestOptions)

	// Make sure the plan shows zero changes
	assert.Contains(t, planResult, "No changes.")
}

func TestJoinHypenDeployment3(t *testing.T) {
	t.Parallel()

	_examplesDir := "../examples"
	exampleDir := filepath.Join(_examplesDir, "join-hypen")

	input_one := "one"
	input_two := "two"
	expected_output := "joined_name = \"one-two\""

	terratestOptions := &terraform.Options{
		TerraformDir: exampleDir,
		NoColor:      true,

		Vars: map[string]interface{}{
			"example_input_one": input_one,
			"example_input_two": input_two,
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terratestOptions)

	// Apply the infrastructure
	applyResult := terraform.InitAndApply(t, terratestOptions)
	assert.Contains(t, applyResult, expected_output)

	// Run perpetual diff
	planResult := terraform.Plan(t, terratestOptions)

	// Make sure the plan shows zero changes
	assert.Contains(t, planResult, "No changes.")
}

func TestJoinHypenDeployment4(t *testing.T) {
	t.Parallel()

	_examplesDir := "../examples"
	exampleDir := filepath.Join(_examplesDir, "join-hypen")

	input_one := "one"
	input_two := "two"
	expected_output := "joined_name = \"one-two\""

	terratestOptions := &terraform.Options{
		TerraformDir: exampleDir,
		NoColor:      true,

		Vars: map[string]interface{}{
			"example_input_one": input_one,
			"example_input_two": input_two,
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terratestOptions)

	// Apply the infrastructure
	applyResult := terraform.InitAndApply(t, terratestOptions)
	assert.Contains(t, applyResult, expected_output)

	// Run perpetual diff
	planResult := terraform.Plan(t, terratestOptions)

	// Make sure the plan shows zero changes
	assert.Contains(t, planResult, "No changes.")
}

func TestJoinHypenDeployment5(t *testing.T) {
	t.Parallel()

	_examplesDir := "../examples"
	exampleDir := filepath.Join(_examplesDir, "join-hypen")

	input_one := "one"
	input_two := "two"
	expected_output := "joined_name = \"one-two\""

	terratestOptions := &terraform.Options{
		TerraformDir: exampleDir,
		NoColor:      true,

		Vars: map[string]interface{}{
			"example_input_one": input_one,
			"example_input_two": input_two,
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terratestOptions)

	// Apply the infrastructure
	applyResult := terraform.InitAndApply(t, terratestOptions)
	assert.Contains(t, applyResult, expected_output)

	// Run perpetual diff
	planResult := terraform.Plan(t, terratestOptions)

	// Make sure the plan shows zero changes
	assert.Contains(t, planResult, "No changes.")
}

func TestJoinHypenDeployment6(t *testing.T) {
	t.Parallel()

	_examplesDir := "../examples"
	exampleDir := filepath.Join(_examplesDir, "join-hypen")

	input_one := "one"
	input_two := "two"
	expected_output := "joined_name = \"one-two\""

	terratestOptions := &terraform.Options{
		TerraformDir: exampleDir,
		NoColor:      true,

		Vars: map[string]interface{}{
			"example_input_one": input_one,
			"example_input_two": input_two,
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terratestOptions)

	// Apply the infrastructure
	applyResult := terraform.InitAndApply(t, terratestOptions)
	assert.Contains(t, applyResult, expected_output)

	// Run perpetual diff
	planResult := terraform.Plan(t, terratestOptions)

	// Make sure the plan shows zero changes
	assert.Contains(t, planResult, "No changes.")
}

func TestJoinHypenDeployment7(t *testing.T) {
	t.Parallel()

	_examplesDir := "../examples"
	exampleDir := filepath.Join(_examplesDir, "join-hypen")

	input_one := "one"
	input_two := "two"
	expected_output := "joined_name = \"one-two\""

	terratestOptions := &terraform.Options{
		TerraformDir: exampleDir,
		NoColor:      true,

		Vars: map[string]interface{}{
			"example_input_one": input_one,
			"example_input_two": input_two,
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terratestOptions)

	// Apply the infrastructure
	applyResult := terraform.InitAndApply(t, terratestOptions)
	assert.Contains(t, applyResult, expected_output)

	// Run perpetual diff
	planResult := terraform.Plan(t, terratestOptions)

	// Make sure the plan shows zero changes
	assert.Contains(t, planResult, "No changes.")
}

func TestJoinHypenDeploymentFail(t *testing.T) {
	t.Parallel()

	_examplesDir := "../examples"
	exampleDir := filepath.Join(_examplesDir, "join-hypen")

	input_one := "one"
	input_two := "two"
	expected_output := "joined_name = \"one-two\""

	terratestOptions := &terraform.Options{
		TerraformDir: exampleDir,
		NoColor:      true,

		Vars: map[string]interface{}{
			"example_input_one": input_one,
			"example_input_two": input_two,
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terratestOptions)

	// Apply the infrastructure
	applyResult := terraform.InitAndApply(t, terratestOptions)
	assert.Contains(t, applyResult, expected_output)

	// Run perpetual diff
	planResult := terraform.Plan(t, terratestOptions)

	// Make sure the plan shows zero changes
	assert.Contains(t, planResult, "No changes.")
}

func TestJoinHypenDeployment8(t *testing.T) {
	t.Parallel()

	_examplesDir := "../examples"
	exampleDir := filepath.Join(_examplesDir, "join-hypen")

	input_one := "one"
	input_two := "two"
	expected_output := "joined_name = \"one-two\""

	terratestOptions := &terraform.Options{
		TerraformDir: exampleDir,
		NoColor:      true,

		Vars: map[string]interface{}{
			"example_input_one": input_one,
			"example_input_two": input_two,
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terratestOptions)

	// Apply the infrastructure
	applyResult := terraform.InitAndApply(t, terratestOptions)
	assert.Contains(t, applyResult, expected_output)

	// Run perpetual diff
	planResult := terraform.Plan(t, terratestOptions)

	// Make sure the plan shows zero changes
	assert.Contains(t, planResult, "No changes.")
}

func TestJoinHypenDeployment9(t *testing.T) {
	t.Parallel()

	_examplesDir := "../examples"
	exampleDir := filepath.Join(_examplesDir, "join-hypen")

	input_one := "one"
	input_two := "two"
	expected_output := "joined_name = \"one-two\""

	terratestOptions := &terraform.Options{
		TerraformDir: exampleDir,
		NoColor:      true,

		Vars: map[string]interface{}{
			"example_input_one": input_one,
			"example_input_two": input_two,
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terratestOptions)

	// Apply the infrastructure
	applyResult := terraform.InitAndApply(t, terratestOptions)
	assert.Contains(t, applyResult, expected_output)

	// Run perpetual diff
	planResult := terraform.Plan(t, terratestOptions)

	// Make sure the plan shows zero changes
	assert.Contains(t, planResult, "No changes.")
}
