// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermWindowsVirtualMachineScaleSetInvalidPriorityRule checks the pattern is valid
type AzurermWindowsVirtualMachineScaleSetInvalidPriorityRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermWindowsVirtualMachineScaleSetInvalidPriorityRule returns new rule with default attributes
func NewAzurermWindowsVirtualMachineScaleSetInvalidPriorityRule() *AzurermWindowsVirtualMachineScaleSetInvalidPriorityRule {
	return &AzurermWindowsVirtualMachineScaleSetInvalidPriorityRule{
		resourceType:  "azurerm_windows_virtual_machine_scale_set",
		attributeName: "priority",
		enum: []string{
			"Regular",
			"Low",
			"Spot",
		},
	}
}

// Name returns the rule name
func (r *AzurermWindowsVirtualMachineScaleSetInvalidPriorityRule) Name() string {
	return "azurerm_windows_virtual_machine_scale_set_invalid_priority"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermWindowsVirtualMachineScaleSetInvalidPriorityRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermWindowsVirtualMachineScaleSetInvalidPriorityRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermWindowsVirtualMachineScaleSetInvalidPriorityRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermWindowsVirtualMachineScaleSetInvalidPriorityRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as priority`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
