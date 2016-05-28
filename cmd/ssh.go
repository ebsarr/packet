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

// sshCmd represents the ssh command
var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Manage your ssh keys for secure login in the packet platform.",
	// Long: ``,
}

var listSSHKeysCmd = &cobra.Command{
	Use:	"listall",
	Short:	"View all configured SSH keys",
	RunE:	func(cmd *cobra.Command, args []string) error {
		err := ListSSHKeys()
		return err
	},
}

var listSSHKeyCmd = &cobra.Command{
	Use:	"list",
	Short:	"View configured SSH key associated with the given ID.",
	RunE:	func(cmd *cobra.Command, args []string) error {
		sshKeyID := cmd.Flag("key-id").Value.String()
		err := ListSSHKey(sshKeyID)
		return err
	},
}

var createSSHKeyCmd = &cobra.Command{
	Use:	"create",
	Short:	"Configure a new SSH key",
	RunE:	func(cmd *cobra.Command, args []string) error {
		key := cmd.Flag("ssh-key").Value.String()
		label := cmd.Flag("label").Value.String()
		err := CreateSSHKey(label, key)
		return err
	},
}

var deleteSSHKeyCmd = &cobra.Command{
	Use:	"delete",
	Short:	"Delete SSH key associated with the given ID.",
	RunE:	func(cmd *cobra.Command, args []string) error {
		sshKeyID := cmd.Flag("key-id").Value.String()
		err := DeleteSSHKey(sshKeyID)
		return err
	},
}

var updateSSHKeyCmd = &cobra.Command{
	Use:	"update",
	Short:	"Update a SSH key: change the key or its label",
	RunE:	func(cmd *cobra.Command, args []string) error {
		sshKeyID := cmd.Flag("key-id").Value.String()
		key := cmd.Flag("ssh-key").Value.String()
		label := cmd.Flag("label").Value.String()
		err := UpdateSSHKey(sshKeyID, label, key)
		return err
	},
}

func init() {
	// Subcommands
	sshCmd.AddCommand(listSSHKeysCmd, listSSHKeyCmd, createSSHKeyCmd, deleteSSHKeyCmd, updateSSHKeyCmd)
	RootCmd.AddCommand(sshCmd)
	
	// Flags for command: packet ssh list
	listSSHKeyCmd.Flags().String("key-id", "", "SSH key ID")
	
	//Flags for command: packet ssh create
	listSSHKeyCmd.Flags().String("label", "", "Label to assign to the key")
	listSSHKeyCmd.Flags().String("ssh-key", "", "SSH key: public key to deploy to servers")
	
	// Flags for command: packet ssh delete
	listSSHKeyCmd.Flags().String("key-id", "", "SSH key ID")
	
	// Flags for command: packet ssh update
	listSSHKeyCmd.Flags().String("key-id", "", "SSH key ID")
	listSSHKeyCmd.Flags().String("label", "", "Label to assign to the key")
	listSSHKeyCmd.Flags().String("ssh-key", "", "SSH key: public key to deploy to servers")
}
