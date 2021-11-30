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
			Name: "indented 3 comments",
			Content: `
module "foo" {
	# test 1
 	# test 2
 	# test 3
 	source = "git::git@github.com:champtitles/my-repo.git"
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewMultilineCommentRule(),
					Message: "avoid the use of comments which span more than 2 lines",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{},
						End:      hcl.Pos{},
					},
				},
			},
		},
		{
			Name: "3 hash comments",
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
						Start:    hcl.Pos{},
						End:      hcl.Pos{},
					},
				},
			},
		},
		{
			Name: "3 slash comments",
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
						Start:    hcl.Pos{},
						End:      hcl.Pos{},
					},
				},
			},
		},
		{
			Name: "2 hash comments",
			Content: `
# test 1
# test 2
module "foo" {
 source = "git::git@github.com:champtitles/my-repo.git"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "2 slash comments",
			Content: `
// test 1
// test 2
module "foo" {
 source = "git::git@github.com:champtitles/my-repo.git"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "1 hash comment",
			Content: `
# test 1
module "foo" {
 source = "git::git@github.com:champtitles/my-repo.git"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "1 slash comments",
			Content: `
// test 1
module "foo" {
 source = "git::git@github.com:champtitles/my-repo.git"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "no comments",
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
