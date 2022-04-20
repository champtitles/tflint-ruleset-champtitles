package rules

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"regexp"
)

// MultilineCommentRule checks for comments which span more than a predefined number of lines
type MultilineCommentRule struct {
	tflint.DefaultRule
}

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
func (r *MultilineCommentRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *MultilineCommentRule) Link() string {
	return ""
}

// Check checks for comments which span more than a predefined number of lines
func (r *MultilineCommentRule) Check(runner tflint.Runner) error {

	lineLimit := 2 // Max number of lines to allow for comment lines

	errorMsg := fmt.Sprintf("avoid the use of comments which span more than %d lines", lineLimit)

	pattern := fmt.Sprintf("(?m:^\\s*[//|#].*\n){%d,}", lineLimit+1)

	files, _ := runner.GetFiles()

	for name, file := range files {

		matched, _ := regexp.MatchString(pattern, string(file.Bytes))
		if matched == true {
			return runner.EmitIssue(r, errorMsg, hcl.Range{
				Filename: name,
				// TODO: Set the start and end to the offending section of the file
				Start: hcl.Pos{},
				End:   hcl.Pos{},
			})
		}
	}
	return nil
}
