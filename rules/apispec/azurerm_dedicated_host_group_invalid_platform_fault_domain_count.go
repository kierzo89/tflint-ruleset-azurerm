// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermDedicatedHostGroupInvalidPlatformFaultDomainCountRule checks the pattern is valid
type AzurermDedicatedHostGroupInvalidPlatformFaultDomainCountRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	min           int
}

// NewAzurermDedicatedHostGroupInvalidPlatformFaultDomainCountRule returns new rule with default attributes
func NewAzurermDedicatedHostGroupInvalidPlatformFaultDomainCountRule() *AzurermDedicatedHostGroupInvalidPlatformFaultDomainCountRule {
	return &AzurermDedicatedHostGroupInvalidPlatformFaultDomainCountRule{
		resourceType:  "azurerm_dedicated_host_group",
		attributeName: "platform_fault_domain_count",
		min:           1,
	}
}

// Name returns the rule name
func (r *AzurermDedicatedHostGroupInvalidPlatformFaultDomainCountRule) Name() string {
	return "azurerm_dedicated_host_group_invalid_platform_fault_domain_count"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermDedicatedHostGroupInvalidPlatformFaultDomainCountRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermDedicatedHostGroupInvalidPlatformFaultDomainCountRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermDedicatedHostGroupInvalidPlatformFaultDomainCountRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermDedicatedHostGroupInvalidPlatformFaultDomainCountRule) Check(runner tflint.Runner) error {
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
			if val < r.min {
				runner.EmitIssue(
					r,
					"platform_fault_domain_count must be 1 or higher",
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
