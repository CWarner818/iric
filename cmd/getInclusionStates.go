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

var (
	txns *[]string
	tips *[]string
)

// getInclusionStatesCmd represents the getInclusionStates command
var getInclusionStatesCmd = &cobra.Command{
	Use:   "getInclusionStates",
	Short: "See if the specified transaction is confirmed by the specified tip",
	Long: `Get the inclusion states of a set of transactions. This is for determining if a transaction was accepted and confirmed by the network or not. You can search for multiple tips (and thus, milestones) to get past inclusion states of transactions.

This API call simply returns a list of boolean values in the same order as the transaction list you submitted, thus you get a true/false whether a transaction is confirmed or not.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Set the default http client to use
		httpClient := &http.Client{
			Timeout: viper.GetDuration("timeout"),
		}
		api := giota.NewAPI(viper.GetString("node"), httpClient)

		txnTrytes, err := toTrytes(*txns)
		if err != nil {
			log.Fatal(err)
		}
		tipTrytes, err := toTrytes(*tips)
		if err != nil {
			log.Fatal(err)
		}
		states, err := api.GetInclusionStates(txnTrytes, tipTrytes)
		if err != nil {
			log.Fatal(err)
		}
		pp.Print(states.States)
	},
}

func init() {
	RootCmd.AddCommand(getInclusionStatesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getInclusionStatesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getInclusionStatesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	txns = getInclusionStatesCmd.Flags().StringSlice("txn", nil, "txn you want to get the inclusion state for")
	tips = getInclusionStatesCmd.Flags().StringSlice("tip", nil, "tip (including milestones) you want to search for the inclusion state")
}
