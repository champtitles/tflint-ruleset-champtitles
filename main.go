package main

import (
	"github.com/champtitles/tflint-ruleset-champtitles/rules"
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "template",
			Version: "0.1.0",
			Rules: []tflint.Rule{
				rules.NewAwsInstanceExampleTypeRule(),
				rules.NewAwsS3BucketExampleLifecycleRuleRule(),
				rules.NewLocalFileExampleProvisionerRule(),
				rules.NewTerraformBackendTypeRule(),
			},
		},
	})
}
