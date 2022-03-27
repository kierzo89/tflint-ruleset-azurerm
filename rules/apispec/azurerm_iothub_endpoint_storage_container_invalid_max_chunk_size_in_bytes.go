// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermIothubEndpointStorageContainerInvalidMaxChunkSizeInBytesRule checks the pattern is valid
type AzurermIothubEndpointStorageContainerInvalidMaxChunkSizeInBytesRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAzurermIothubEndpointStorageContainerInvalidMaxChunkSizeInBytesRule returns new rule with default attributes
func NewAzurermIothubEndpointStorageContainerInvalidMaxChunkSizeInBytesRule() *AzurermIothubEndpointStorageContainerInvalidMaxChunkSizeInBytesRule {
	return &AzurermIothubEndpointStorageContainerInvalidMaxChunkSizeInBytesRule{
		resourceType:  "azurerm_iothub_endpoint_storage_container",
		attributeName: "max_chunk_size_in_bytes",
		max:           524288000,
		min:           10485760,
	}
}

// Name returns the rule name
func (r *AzurermIothubEndpointStorageContainerInvalidMaxChunkSizeInBytesRule) Name() string {
	return "azurerm_iothub_endpoint_storage_container_invalid_max_chunk_size_in_bytes"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermIothubEndpointStorageContainerInvalidMaxChunkSizeInBytesRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermIothubEndpointStorageContainerInvalidMaxChunkSizeInBytesRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermIothubEndpointStorageContainerInvalidMaxChunkSizeInBytesRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermIothubEndpointStorageContainerInvalidMaxChunkSizeInBytesRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}
		var val int
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		err = runner.EnsureNoError(err, func() error {
			if val > r.max {
				runner.EmitIssue(
					r,
					"max_chunk_size_in_bytes must be 524288000 or less",
					attribute.Expr.Range(),
				)
			}
			if val < r.min {
				runner.EmitIssue(
					r,
					"max_chunk_size_in_bytes must be 10485760 or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
