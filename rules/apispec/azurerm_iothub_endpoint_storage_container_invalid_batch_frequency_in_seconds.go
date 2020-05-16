// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermIothubEndpointStorageContainerInvalidBatchFrequencyInSecondsRule checks the pattern is valid
type AzurermIothubEndpointStorageContainerInvalidBatchFrequencyInSecondsRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAzurermIothubEndpointStorageContainerInvalidBatchFrequencyInSecondsRule returns new rule with default attributes
func NewAzurermIothubEndpointStorageContainerInvalidBatchFrequencyInSecondsRule() *AzurermIothubEndpointStorageContainerInvalidBatchFrequencyInSecondsRule {
	return &AzurermIothubEndpointStorageContainerInvalidBatchFrequencyInSecondsRule{
		resourceType:  "azurerm_iothub_endpoint_storage_container",
		attributeName: "batch_frequency_in_seconds",
		max:           720,
		min:           60,
	}
}

// Name returns the rule name
func (r *AzurermIothubEndpointStorageContainerInvalidBatchFrequencyInSecondsRule) Name() string {
	return "azurerm_iothub_endpoint_storage_container_invalid_batch_frequency_in_seconds"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermIothubEndpointStorageContainerInvalidBatchFrequencyInSecondsRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermIothubEndpointStorageContainerInvalidBatchFrequencyInSecondsRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermIothubEndpointStorageContainerInvalidBatchFrequencyInSecondsRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermIothubEndpointStorageContainerInvalidBatchFrequencyInSecondsRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val int
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if val > r.max {
				runner.EmitIssue(
					r,
					"batch_frequency_in_seconds must be 720 or less",
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			if val < r.min {
				runner.EmitIssue(
					r,
					"batch_frequency_in_seconds must be 60 or higher",
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
