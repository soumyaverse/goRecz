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
	"github.com/burpOverflow/goRecz/ft"
	"github.com/spf13/cobra"
)

// ftCmd represents the ft command
var ftCmd = &cobra.Command{
	Use:   "ft",
	Short: "Find title from domain list",

	Run: ftHandler,
}

func init() {
	rootCmd.AddCommand(ftCmd)

	ftCmd.Flags().StringP("domainlist", "d", "", "domain list")
	ftCmd.MarkFlagRequired("domainlist")
}

func ftHandler(cmd *cobra.Command, args []string) {
	domainlist, _ := cmd.Flags().GetString("domainlist")
	ft.Handler(&domainlist, &wg)
	wg.Wait()
}
