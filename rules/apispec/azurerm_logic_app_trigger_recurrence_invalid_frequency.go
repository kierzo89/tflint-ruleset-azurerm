// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermLogicAppTriggerRecurrenceInvalidFrequencyRule checks the pattern is valid
type AzurermLogicAppTriggerRecurrenceInvalidFrequencyRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermLogicAppTriggerRecurrenceInvalidFrequencyRule returns new rule with default attributes
func NewAzurermLogicAppTriggerRecurrenceInvalidFrequencyRule() *AzurermLogicAppTriggerRecurrenceInvalidFrequencyRule {
	return &AzurermLogicAppTriggerRecurrenceInvalidFrequencyRule{
		resourceType:  "azurerm_logic_app_trigger_recurrence",
		attributeName: "frequency",
		enum: []string{
			"NotSpecified",
			"Second",
			"Minute",
			"Hour",
			"Day",
			"Week",
			"Month",
			"Year",
		},
	}
}

// Name returns the rule name
func (r *AzurermLogicAppTriggerRecurrenceInvalidFrequencyRule) Name() string {
	return "azurerm_logic_app_trigger_recurrence_invalid_frequency"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermLogicAppTriggerRecurrenceInvalidFrequencyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermLogicAppTriggerRecurrenceInvalidFrequencyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermLogicAppTriggerRecurrenceInvalidFrequencyRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermLogicAppTriggerRecurrenceInvalidFrequencyRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as frequency`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
