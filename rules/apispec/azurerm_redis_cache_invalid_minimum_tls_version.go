// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermRedisCacheInvalidMinimumTLSVersionRule checks the pattern is valid
type AzurermRedisCacheInvalidMinimumTLSVersionRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermRedisCacheInvalidMinimumTLSVersionRule returns new rule with default attributes
func NewAzurermRedisCacheInvalidMinimumTLSVersionRule() *AzurermRedisCacheInvalidMinimumTLSVersionRule {
	return &AzurermRedisCacheInvalidMinimumTLSVersionRule{
		resourceType:  "azurerm_redis_cache",
		attributeName: "minimum_tls_version",
		enum: []string{
			"1.0",
			"1.1",
			"1.2",
		},
	}
}

// Name returns the rule name
func (r *AzurermRedisCacheInvalidMinimumTLSVersionRule) Name() string {
	return "azurerm_redis_cache_invalid_minimum_tls_version"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermRedisCacheInvalidMinimumTLSVersionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermRedisCacheInvalidMinimumTLSVersionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermRedisCacheInvalidMinimumTLSVersionRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermRedisCacheInvalidMinimumTLSVersionRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as minimum_tls_version`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
