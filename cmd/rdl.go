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
	"github.com/burpOverflow/goRecz/rdl"
	"github.com/spf13/cobra"
)

// var (
// 	file       string
// 	outputfile string
// 	showlines  bool
// )

// rdlCmd represents the rdl command
var rdlCmd = &cobra.Command{
	Use:   "rdl",
	Short: "rdl used to remove duplicate lines from a file",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: rdlHandler,
}

func init() {
	rootCmd.AddCommand(rdlCmd)

	rdlCmd.Flags().StringP("file", "f", "", "Source file")
	rdlCmd.Flags().StringP("output", "o", "", "Output file")
	rdlCmd.Flags().BoolP("showlines", "s", false, "Show output lines")

	rdlCmd.MarkFlagRequired("file")
}

func rdlHandler(cmd *cobra.Command, args []string) {
	file, _ := cmd.Flags().GetString("file")
	output, _ := cmd.Flags().GetString("output")
	showlines, _ := cmd.Flags().GetBool("showlines")

	rdl.Handle(&file, &output, &showlines)
}
