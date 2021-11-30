package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/terraform/configs"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"regexp"
)

// MultilineCommentRule checks for comments which span more than a set number of lines
type MultilineCommentRule struct{}

// NewMultilineCommentRule returns a new rule
func NewMultilineCommentRule() *MultilineCommentRule {
	return &MultilineCommentRule{}
}

// Name returns the rule name
func (r *MultilineCommentRule) Name() string {
	return "multiline_comment"
}

// Enabled returns whether the rule is enabled by default
func (r *MultilineCommentRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *MultilineCommentRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *MultilineCommentRule) Link() string {
	return ""
}

// Check checks for comments which span more than a set number of lines
func (r *MultilineCommentRule) Check(runner tflint.Runner) error {
	return runner.WalkModuleCalls(func(call *configs.ModuleCall) error {

		lineLimit := 2 // Max number of lines to allow for comment lines

		errorMsg := fmt.Sprintf("avoid the use of comments which span more than %d lines", lineLimit)

		// Regex pattern to match contiguous lines of comments
		pattern := fmt.Sprintf("(?m:^[//|#].*\n){%d,}", lineLimit+1)

		files, _ := runner.Files()

		for _, file := range files {
			matched, err := regexp.MatchString(pattern, string(file.Bytes))
			if err != nil {
				panic(err)
			}
			if matched == true {
				return runner.EmitIssue(r, errorMsg, call.SourceAddrRange)
			}
		}

		return nil
	})
}
