// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermStorageAccountInvalidAccountKindRule checks the pattern is valid
type AzurermStorageAccountInvalidAccountKindRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermStorageAccountInvalidAccountKindRule returns new rule with default attributes
func NewAzurermStorageAccountInvalidAccountKindRule() *AzurermStorageAccountInvalidAccountKindRule {
	return &AzurermStorageAccountInvalidAccountKindRule{
		resourceType:  "azurerm_storage_account",
		attributeName: "account_kind",
		enum: []string{
			"Storage",
			"StorageV2",
			"BlobStorage",
			"FileStorage",
			"BlockBlobStorage",
		},
	}
}

// Name returns the rule name
func (r *AzurermStorageAccountInvalidAccountKindRule) Name() string {
	return "azurerm_storage_account_invalid_account_kind"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermStorageAccountInvalidAccountKindRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermStorageAccountInvalidAccountKindRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermStorageAccountInvalidAccountKindRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermStorageAccountInvalidAccountKindRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as account_kind`, truncateLongMessage(val)),
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
