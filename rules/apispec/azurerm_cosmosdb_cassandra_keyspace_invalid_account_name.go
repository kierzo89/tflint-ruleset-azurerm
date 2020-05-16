// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermCosmosdbCassandraKeyspaceInvalidAccountNameRule checks the pattern is valid
type AzurermCosmosdbCassandraKeyspaceInvalidAccountNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermCosmosdbCassandraKeyspaceInvalidAccountNameRule returns new rule with default attributes
func NewAzurermCosmosdbCassandraKeyspaceInvalidAccountNameRule() *AzurermCosmosdbCassandraKeyspaceInvalidAccountNameRule {
	return &AzurermCosmosdbCassandraKeyspaceInvalidAccountNameRule{
		resourceType:  "azurerm_cosmosdb_cassandra_keyspace",
		attributeName: "account_name",
		pattern:       regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*`),
	}
}

// Name returns the rule name
func (r *AzurermCosmosdbCassandraKeyspaceInvalidAccountNameRule) Name() string {
	return "azurerm_cosmosdb_cassandra_keyspace_invalid_account_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermCosmosdbCassandraKeyspaceInvalidAccountNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermCosmosdbCassandraKeyspaceInvalidAccountNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermCosmosdbCassandraKeyspaceInvalidAccountNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermCosmosdbCassandraKeyspaceInvalidAccountNameRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-z0-9]+(-[a-z0-9]+)*`),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
