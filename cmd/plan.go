// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
)

// planCmd represents the plan command
var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "View available plans.",
	// Long: ``,
}

var listPlanCmd = &cobra.Command{
	Use:	"list",
	Short:	"Print out available plans.",
	RunE:	func(cmd *cobra.Command, args []string) error {
		err := ListPlans()
		return err
	},
}

func init() {
	planCmd.AddCommand(listPlanCmd)
	RootCmd.AddCommand(planCmd)
}
