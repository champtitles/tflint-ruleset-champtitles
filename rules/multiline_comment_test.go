package rules

import (
	"github.com/hashicorp/hcl/v2"
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_MultilineCommentRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "no module reference",
			Content: `
# test 1
# test 2
# test 3
module "foo" {
 source = "git::git@github.com:champtitles/my-repo.git"
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewMultilineCommentRule(),
					Message: "avoid the use of comments which span more than 2 lines",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 6, Column: 11},
						End:      hcl.Pos{Line: 6, Column: 56},
					},
				},
			},
		},
		{
			Name: "no module reference",
			Content: `
// test 1
// test 2
// test 3
module "foo" {
 source = "git::git@github.com:champtitles/my-repo.git"
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewMultilineCommentRule(),
					Message: "avoid the use of comments which span more than 2 lines",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 6, Column: 11},
						End:      hcl.Pos{Line: 6, Column: 56},
					},
				},
			},
		},
		{
			Name: "hash module reference",
			Content: `
# test 1
# test 2
module "foo" {
 source = "git::git@github.com:champtitles/my-repo.git"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "local module reference",
			Content: `
// test 1
// test 2
module "foo" {
 source = "git::git@github.com:champtitles/my-repo.git"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "hash module reference",
			Content: `
# test 1
module "foo" {
 source = "git::git@github.com:champtitles/my-repo.git"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "local module reference",
			Content: `
// test 1
module "foo" {
 source = "git::git@github.com:champtitles/my-repo.git"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "hash module reference",
			Content: `
module "foo" {
 source = "git::git@github.com:champtitles/my-repo.git"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "local module reference",
			Content: `
module "foo" {
 source = "git::git@github.com:champtitles/my-repo.git"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewMultilineCommentRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
