package rules

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"regexp"
)

// ResourceCountRule checks the number of resources declared in each terraform file
type ResourceCountRule struct{}

// NewResourceCountRule returns a new rule
func NewResourceCountRule() *ResourceCountRule {
	return &ResourceCountRule{}
}

// Name returns the rule name
func (r *ResourceCountRule) Name() string {
	return "resource_count"
}

// Enabled returns whether the rule is enabled by default
func (r *ResourceCountRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *ResourceCountRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *ResourceCountRule) Link() string {
	return ""
}

// Check checks the number of resources declared in each terraform file
func (r *ResourceCountRule) Check(runner tflint.Runner) error {

	resourceLimit := 12 // Max limit of resource declarations per file

	pattern := regexp.MustCompile("\\sresource \"")

	files, _ := runner.Files()

	for name, file := range files {

		matches := pattern.FindAllSubmatchIndex(file.Bytes, -1)
		if len(matches) > resourceLimit {

			errorMsg := fmt.Sprintf("Found %d resources. The limit is %d. Consider splitting resources into multiple files", len(matches), resourceLimit)
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
