// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermTrafficManagerProfileInvalidProfileStatusRule checks the pattern is valid
type AzurermTrafficManagerProfileInvalidProfileStatusRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermTrafficManagerProfileInvalidProfileStatusRule returns new rule with default attributes
func NewAzurermTrafficManagerProfileInvalidProfileStatusRule() *AzurermTrafficManagerProfileInvalidProfileStatusRule {
	return &AzurermTrafficManagerProfileInvalidProfileStatusRule{
		resourceType:  "azurerm_traffic_manager_profile",
		attributeName: "profile_status",
		enum: []string{
			"Enabled",
			"Disabled",
		},
	}
}

// Name returns the rule name
func (r *AzurermTrafficManagerProfileInvalidProfileStatusRule) Name() string {
	return "azurerm_traffic_manager_profile_invalid_profile_status"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermTrafficManagerProfileInvalidProfileStatusRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermTrafficManagerProfileInvalidProfileStatusRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermTrafficManagerProfileInvalidProfileStatusRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermTrafficManagerProfileInvalidProfileStatusRule) Check(runner tflint.Runner) error {
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
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		err = runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as profile_status`, truncateLongMessage(val)),
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
