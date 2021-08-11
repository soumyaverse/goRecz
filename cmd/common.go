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

// commonCmd represents the common command
var commonCmd = &cobra.Command{
	Use:   "common",
	Short: "Common Lines between two files",
	Run:   commonHandler,
}

func init() {
	rootCmd.AddCommand(commonCmd)

	commonCmd.Flags().StringP("firstfile", "f", "", "first file")
	commonCmd.Flags().StringP("secondfile", "s", "", "second file")
	commonCmd.Flags().StringP("outputfile", "o", "", "output file")

}

func commonHandler(cmd *cobra.Command, args []string) {
	firstfile, _ := cmd.Flags().GetString("firstfile")
	secondfile, _ := cmd.Flags().GetString("secondfile")
	outputfile, _ := cmd.Flags().GetString("outputfile")

	diff.Handle(&firstfile, &secondfile, &outputfile, true)
}
