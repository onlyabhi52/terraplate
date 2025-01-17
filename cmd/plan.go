/*
Copyright © 2021 Verifa <info@verifa.io>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/verifa/terraplate/builder"
	"github.com/verifa/terraplate/parser"
	"github.com/verifa/terraplate/runner"
)

var (
	runBuild bool
	runInit  bool
)

// planCmd represents the plan command
var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Runs terraform plan on all subdirectories",
	Long:  `Runs terraform plan on all subdirectories.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := parser.Parse(&config.ParserConfig)
		if err != nil {
			return fmt.Errorf("parsing terraplate: %w", err)
		}
		if runBuild {
			if err := builder.Build(config); err != nil {
				return fmt.Errorf("building terrplate: %w", err)
			}
		}
		runOpts := []func(r *runner.TerraRun){
			runner.RunPlan(),
		}
		if runInit {
			runOpts = append(runOpts, runner.RunInit())
		}
		return runner.Run(config, runOpts...)
	},
}

func init() {
	rootCmd.AddCommand(planCmd)

	planCmd.Flags().BoolVar(&runBuild, "build", false, "Run build process also")
	planCmd.Flags().BoolVar(&runInit, "init", false, "Run terraform init also")
}
