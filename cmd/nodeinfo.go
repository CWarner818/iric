// Copyright © 2018 Chris Warner
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"log"
	"net/http"

	"github.com/cwarner818/giota"
	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// nodeinfoCmd represents the nodeinfo command
var nodeinfoCmd = &cobra.Command{
	Use:   "nodeinfo",
	Short: "Returns information about the node",
	Long:  "Returns information about the node",
	Run: func(cmd *cobra.Command, args []string) {
		// Set the default http client to use
		httpClient := &http.Client{
			Timeout: viper.GetDuration("timeout"),
		}
		api := giota.NewAPI(viper.GetString("node"), httpClient)

		nodeInfo, err := api.GetNodeInfo()

		if err != nil {
			log.Fatal(err)
		}
		pp.Print(*nodeInfo)
	},
}

func init() {
	RootCmd.AddCommand(nodeinfoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nodeinfoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nodeinfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
