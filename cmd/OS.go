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

// OSCmd represents the OS command
var OSCmd = &cobra.Command{
	Use:   "OS",
	Short: "View available operating systems",
	// Long: ``,
	// Run: func(cmd *cobra.Command, args []string) {
	// TODO: Work your own magic here
	//fmt.Println("OS called")
	// },
}

// OSListCmd represents the OS list command
var OSListCmd = &cobra.Command{
	Use:   "list",
	Short: "Print out available operating systems",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := ListOS()
		return err
	},
}

func init() {
	OSCmd.AddCommand(OSListCmd)
	RootCmd.AddCommand(OSCmd)
}
