// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/miyazi777/git-desc/git"
	config "github.com/miyazi777/git-desc/git/config"
	"github.com/miyazi777/git-desc/shell"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"path/filepath"
)

var gitConfig config.Config
var description config.Description
var page config.Page

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gitdesc",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	SilenceErrors: true,
	SilenceUsage:  true,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	setup()
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gitdesc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.AddConfigPath(filepath.Join(home, ".config", "git-desc"))
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		// setting defaults
		viper.Set("editor", "vi")
		viper.Set("list.branch.width", 0.35)
		viper.Set("list.description.width", 0.55)
	}
}

func setup() {
	command := &shell.CommandImpl{}

	git := &git.GitImpl{
		Command: command,
	}

	gitConfig = &config.ConfigImpl{
		Git: git,
	}
	description = &config.DescriptionImpl{
		Git: git,
	}
	page = &config.PageImpl{
		Git: git,
	}
}
