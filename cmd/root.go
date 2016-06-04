package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const version = "0.1"

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:          "packet",
	Short:        "A unified tool to manage your packet services",
	SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {
		showVersion, _ := cmd.Flags().GetBool("version")
		if showVersion {
			fmt.Printf("%s\n", version)
		}
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		// fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global flag for setting the API key.
	RootCmd.PersistentFlags().StringP("key", "k", "", "Specify the api key")
	RootCmd.PersistentFlags().BoolP("version", "v", false, "Show version and exit")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".packet") // name of config file (without extension)
	viper.AddConfigPath("$HOME")   // adding home directory as first search path
	viper.AutomaticEnv()           // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
