//go:build unit

package test

import (
	"path/filepath"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestJoinHypenValidity(t *testing.T) {
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

func TestJoinHypenValidity1(t *testing.T) {
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

func TestJoinHypenValidity2(t *testing.T) {
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

func TestJoinHypenValidity3(t *testing.T) {
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

func TestJoinHypenValidity4(t *testing.T) {
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

func TestJoinHypenValidity5(t *testing.T) {
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

func TestJoinHypenValidity6(t *testing.T) {
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

func TestJoinHypenValidityFail(t *testing.T) {
	t.Parallel()

	_fixturesDir := "./fixtures"
	exampleDir := filepath.Join(_fixturesDir, "join-hypen")

	terratestOptions := &terraform.Options{
		TerraformDir: exampleDir,
		NoColor:      true,
	}

	t.Logf("Running in %s", exampleDir)
	output := terraform.InitAndPlan(t, terratestOptions)

	assert.Contains(t, output, "Some changes. Your infrastructure matches the configuration.")
}

func TestJoinHypenValidity7(t *testing.T) {
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

func TestJoinHypenValidity8(t *testing.T) {
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

func TestJoinHypenValidity9(t *testing.T) {
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
