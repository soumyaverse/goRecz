/*
Copyright © 2021 burpOverflow

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
	"github.com/burpOverflow/goRecz/diff"
	"github.com/spf13/cobra"
)

// diffCmd represents the diff command
var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Different Lines bet two files",

	Run: diffHandler,
}

func init() {
	rootCmd.AddCommand(diffCmd)

	diffCmd.Flags().StringP("firstfile", "f", "", "first file name")
	diffCmd.Flags().StringP("secondfile", "s", "", "second file name")
	diffCmd.Flags().StringP("outputfile", "o", "", "output file name")

	diffCmd.MarkFlagRequired("firstfile")
	diffCmd.MarkFlagRequired("secondfile")
}

func diffHandler(cmd *cobra.Command, args []string) {
	firstfile, _ := cmd.Flags().GetString("firstfile")
	secondfile, _ := cmd.Flags().GetString("secondfile")
	outputfile, _ := cmd.Flags().GetString("outputfile")

	diff.Handle(&firstfile, &secondfile, &outputfile, false)
}
