/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode [base64 string to decode]",
	Short: "decode to string from base64 encoded string",
	Long:  `decode to string from base64 encoded string`,
	Args:  ValidateArgument,
	Run:   decode,
}

func decode(cmd *cobra.Command, args []string) {
	decoded, err := base64.StdEncoding.DecodeString(args[0])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", decoded)
}

func init() {
	rootCmd.AddCommand(decodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
