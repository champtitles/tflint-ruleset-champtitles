package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/terraform/configs"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"regexp"
	"strings"
)

// ModuleCallUsingHashRule checks whether git-based module calls use a hash reference instead of a tag or branch name
type ModuleCallUsingHashRule struct{}

// NewModuleCallUsingHashRule returns a new rule
func NewModuleCallUsingHashRule() *ModuleCallUsingHashRule {
	return &ModuleCallUsingHashRule{}
}

// Name returns the rule name
func (r *ModuleCallUsingHashRule) Name() string {
	return "module_call_using_hash"
}

// Enabled returns whether the rule is enabled by default
func (r *ModuleCallUsingHashRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *ModuleCallUsingHashRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *ModuleCallUsingHashRule) Link() string {
	return ""
}

// Check checks whether git-based module calls use a hash reference instead of a tag or branch name
func (r *ModuleCallUsingHashRule) Check(runner tflint.Runner) error {
	return runner.WalkModuleCalls(func(call *configs.ModuleCall) error {

		if !strings.Contains(call.SourceAddr, "git::") {
			return nil
		}

		matched, _ := regexp.MatchString(`\w{40}$`, call.SourceAddr)
		if matched == false {
			return runner.EmitIssue(r, "git module source should use a hash reference", call.SourceAddrRange)
		}

		return nil
	})
}
