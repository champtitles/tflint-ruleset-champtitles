package rules

import (
	"github.com/hashicorp/hcl/v2"
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ModuleCallUsingHashRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "version module reference",
			Content: `
module "foo" {
 source = "git::git@github.com:champtitles/my-repo.git?ref=1.0.0"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewModuleCallUsingHashRule(),
					Message: "git module source should use a hash reference",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 11},
						End:      hcl.Pos{Line: 3, Column: 66},
					},
				},
			},
		},
		{
			Name: "name module reference",
			Content: `
module "foo" {
 source = "git::git@github.com:champtitles/my-repo.git?ref=main"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewModuleCallUsingHashRule(),
					Message: "git module source should use a hash reference",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 11},
						End:      hcl.Pos{Line: 3, Column: 65},
					},
				},
			},
		},
		{
			Name: "no module reference",
			Content: `
module "foo" {
 source = "git::git@github.com:champtitles/my-repo.git"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewModuleCallUsingHashRule(),
					Message: "git module source should use a hash reference",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 11},
						End:      hcl.Pos{Line: 3, Column: 56},
					},
				},
			},
		},
		{
			Name: "hash module reference",
			Content: `
module "foo" {
 source = "git::git@github.com:champtitles/my-repo.git?ref=c64220a05fa5a34d68f6d836eb4dfcf6e9753dc7"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "local module reference",
			Content: `
module "foo" {
  source = "../test"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewModuleCallUsingHashRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
