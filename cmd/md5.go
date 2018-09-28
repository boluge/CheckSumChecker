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
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// md5Cmd represents the md5 command
var md5Cmd = &cobra.Command{
	Use:   "md5",
	Short: "Application to check file md5 checksum",
	Long: `The Application allows to calculate file checksum, to check them and this for different formats like :
	md5, sha1, sha256, pgp.`,
	Run: func(cmd *cobra.Command, args []string) {
		if source == "" {
			fmt.Println("You must add a file path")
		} else {
			hash := getMd5Hash()
			showResult(hash, key)
		}
	},
}

// getHash return the hash of a file
func getMd5Hash() string {
	f, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func init() {
	rootCmd.AddCommand(md5Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// md5Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// md5Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
