// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermCosmosdbGremlinGraphInvalidResourceGroupNameRule checks the pattern is valid
type AzurermCosmosdbGremlinGraphInvalidResourceGroupNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermCosmosdbGremlinGraphInvalidResourceGroupNameRule returns new rule with default attributes
func NewAzurermCosmosdbGremlinGraphInvalidResourceGroupNameRule() *AzurermCosmosdbGremlinGraphInvalidResourceGroupNameRule {
	return &AzurermCosmosdbGremlinGraphInvalidResourceGroupNameRule{
		resourceType:  "azurerm_cosmosdb_gremlin_graph",
		attributeName: "resource_group_name",
		pattern:       regexp.MustCompile(`^[-\w\._\(\)]+$`),
	}
}

// Name returns the rule name
func (r *AzurermCosmosdbGremlinGraphInvalidResourceGroupNameRule) Name() string {
	return "azurerm_cosmosdb_gremlin_graph_invalid_resource_group_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermCosmosdbGremlinGraphInvalidResourceGroupNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermCosmosdbGremlinGraphInvalidResourceGroupNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermCosmosdbGremlinGraphInvalidResourceGroupNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermCosmosdbGremlinGraphInvalidResourceGroupNameRule) Check(runner tflint.Runner) error {
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
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[-\w\._\(\)]+$`),
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
