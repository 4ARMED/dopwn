// Copyright © 2018 Marc Wickenden <marc@4armed.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"os"

	"github.com/4armed/dopwn/pkg/exploit"
	"github.com/kubicorn/kubicorn/pkg/logger"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dopwn",
	Short: "Take over a DigitalOcean Kubernetes cluster via sensitive metadata",
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(exploit.Command())

	rootCmd.PersistentFlags().IntVarP(&logger.Level, "verbose", "v", 3, "set log level, use 0 to silence, 4 for debugging")
	rootCmd.PersistentFlags().BoolVarP(&logger.Color, "color", "C", true, "toggle colorized logs")
}
