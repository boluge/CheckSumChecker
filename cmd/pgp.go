// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

	"github.com/spf13/cobra"
)

// pgpCmd represents the pgp command
var pgpCmd = &cobra.Command{
	Use:   "pgp",
	Short: "Application to check file pgp checksum",
	Long: `The Application allows to calculate file checksum, to check them and this for different formats like :
	md5, sha1, sha256, pgp.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pgp called")
	},
}

func init() {
	rootCmd.AddCommand(pgpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pgpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pgpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
