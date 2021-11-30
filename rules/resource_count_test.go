package rules

import (
	"github.com/hashicorp/hcl/v2"
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ResourceCountRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "13 resources",
			Content: `
resource "null_resource" "example1" {}
resource "null_resource" "example1" {}
resource "null_resource" "example1" {}
resource "null_resource" "example1" {}
resource "null_resource" "example1" {}
resource "null_resource" "example1" {}
resource "null_resource" "example1" {}
resource "null_resource" "example1" {}
resource "null_resource" "example1" {}
resource "null_resource" "example1" {}
resource "null_resource" "example1" {}
resource "null_resource" "example1" {}
resource "null_resource" "example1" {}
`,
			Expected: helper.Issues{
				{
					Rule:    NewResourceCountRule(),
					Message: "Found 13 resources. The limit is 12. Consider splitting resources into multiple files",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{},
						End:      hcl.Pos{},
					},
				},
			},
		},
		{
			Name: "12 resources",
			Content: `
resource "null_resource" "example1" {}
resource "null_resource" "example1" {}
resource "null_resource" "example1" {}
resource "null_resource" "example1" {}
resource "null_resource" "example1" {}
resource "null_resource" "example1" {}
resource "null_resource" "example1" {}
resource "null_resource" "example1" {}
resource "null_resource" "example1" {}
resource "null_resource" "example1" {}
resource "null_resource" "example1" {}
resource "null_resource" "example1" {}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "2 resources",
			Content: `
resource "null_resource" "example1" {}
resource "null_resource" "example1" {}
`,
			Expected: helper.Issues{},
		},
	}

	rule := NewResourceCountRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
