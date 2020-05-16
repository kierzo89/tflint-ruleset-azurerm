// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermPublicIPInvalidSkuRule checks the pattern is valid
type AzurermPublicIPInvalidSkuRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermPublicIPInvalidSkuRule returns new rule with default attributes
func NewAzurermPublicIPInvalidSkuRule() *AzurermPublicIPInvalidSkuRule {
	return &AzurermPublicIPInvalidSkuRule{
		resourceType:  "azurerm_public_ip",
		attributeName: "sku",
		enum: []string{
			"Basic",
			"Standard",
		},
	}
}

// Name returns the rule name
func (r *AzurermPublicIPInvalidSkuRule) Name() string {
	return "azurerm_public_ip_invalid_sku"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermPublicIPInvalidSkuRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermPublicIPInvalidSkuRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermPublicIPInvalidSkuRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermPublicIPInvalidSkuRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as sku`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
