package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"regexp"
	"strings"
)

// ModuleCallUsingHashRule checks whether git-based module calls use a hash reference instead of a tag or branch name
type ModuleCallUsingHashRule struct {
	tflint.DefaultRule
}

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
func (r *ModuleCallUsingHashRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *ModuleCallUsingHashRule) Link() string {
	return ""
}

// Check checks whether git-based module calls use a hash reference instead of a tag or branch name
func (r *ModuleCallUsingHashRule) Check(runner tflint.Runner) error {

	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "module",
				LabelNames: []string{"name"},
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{{Name: "source"}},
				},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, module := range content.Blocks {
		attribute, exists := module.Body.Attributes["source"]
		if !exists {
			continue
		}

		var sourceValue string
		err := runner.EvaluateExpr(attribute.Expr, &sourceValue, nil)
		if err != nil {
			return err
		}

		if !strings.Contains(sourceValue, "git::") {
			return nil
		}

		matched, _ := regexp.MatchString(`\w{40}$`, sourceValue)
		if matched == false {
			return runner.EmitIssue(r, "git module source should use a hash reference", attribute.Expr.Range())
		}
	}

	return nil
}
