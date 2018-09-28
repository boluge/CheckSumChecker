// Copyright © 2018 Stéphane Deluce
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

	"github.com/fatih/color"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	emoji "gopkg.in/kyokomi/emoji.v1"
)

var source string
var key string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "CheckSumChecker",
	Short: "Application to check file checksum",
	Long: `This application allows to calculate file checksum,
to check them and this for different formats like :
[md5, sha1, sha256, sha512, pgp]`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) { },
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
	cobra.OnInitialize(initConfig)

	// Flag for the file
	rootCmd.PersistentFlags().StringVarP(&source, "file", "f", "", "Path of the hash file")
	// Flag for the key
	rootCmd.PersistentFlags().StringVarP(&key, "key", "k", "", "Expected key")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("path", "p", , "Help message for toggle")
}

// Compare Hash
func isSame(hash string, key string) bool {
	if hash == key {
		return true
	} else {
		return false
	}
}

func showResult(hash string, key string) {
	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	hashSuccess := color.New(color.Bold, color.FgGreen).SprintFunc()
	hashFail := color.New(color.Bold, color.FgRed).SprintFunc()
	space := "    "
	if key != "" {
		if isSame(hash, key) {
			emoji.Printf("%s:beer: %s\n", space, hashSuccess("The hash is Valid"))
			fmt.Printf("%sHash\t => %s \n", space, cyan(hash))
			fmt.Printf("%sExpected\t => %s \n", space, cyan(key))
		} else {
			emoji.Printf("%s:warning: %s\n", space, hashFail("The hash is not Valid"))
			fmt.Printf("%sHash\t => %s \n", space, yellow(hash))
			fmt.Printf("%sExpected\t => %s \n", space, yellow(key))
		}
	} else {
		fmt.Printf("%sHash => %s \n", space, green(hash))
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if source != "" {
		// Use config file from the flag.
		viper.SetConfigFile(source)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".CheckSumChecker" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".CheckSumChecker")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
