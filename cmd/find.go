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
	bundles   *[]string
	addresses *[]string
	tags      *[]string
	approvees *[]string
)

func toTrytes(input []string) ([]giota.Trytes, error) {
	var output []giota.Trytes
	for _, t := range input {
		trytes, err := giota.ToTrytes(t)
		if err != nil {
			return nil, err
		}
		output = append(output, trytes)
	}
	return output, nil
}
func toAddress(input []string) ([]giota.Address, error) {
	var output []giota.Address
	for _, t := range input {
		trytes, err := giota.ToAddress(t)
		if err != nil {
			return nil, err
		}
		output = append(output, trytes)
	}
	return output, nil
}

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Search for transactions",
	Long:  `Find the transactions which match the specified input and return. All input values are lists, for which a list of return values (transaction hashes), in the same order, is returned for all individual elements. The input fields can either be bundles, addresses, tags or approvees. Using multiple of these input fields returns the intersection of the values.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Set the default http client to use
		httpClient := &http.Client{
			Timeout: viper.GetDuration("timeout"),
		}

		log.Println("NODE", viper.GetString("node"))
		log.Println("DUR", viper.GetDuration("timeout"))
		api := giota.NewAPI(viper.GetString("node"), httpClient)

		bundleTrytes, err := toTrytes(*bundles)
		if err != nil {
			log.Fatal(err)
		}

		addrTrytes, err := toAddress(*addresses)
		if err != nil {
			log.Fatal(err)
		}
		tagTrytes, err := toTrytes(*tags)
		if err != nil {
			log.Fatal(err)
		}
		childrenTrytes, err := toTrytes(*approvees)
		if err != nil {
			log.Fatal(err)
		}
		nodeInfo, err := api.FindTransactions(&giota.FindTransactionsRequest{
			Bundles:   bundleTrytes,
			Addresses: addrTrytes,
			Tags:      tagTrytes,
			Approvees: childrenTrytes,
		})

		if err != nil {
			pp.Print(tagTrytes)
			log.Println(len(tagTrytes))
			log.Fatal(err)
		}
		pp.Print(*nodeInfo)
	},
}

func init() {
	RootCmd.AddCommand(findCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	bundles = findCmd.Flags().StringSliceP("bundle", "b", nil, "bundle to search for")
	addresses = findCmd.Flags().StringSliceP("address", "a", nil, "address to search for")
	tags = findCmd.Flags().StringSliceP("tag", "t", nil, "tag to search for")
	approvees = findCmd.Flags().StringSliceP("child", "c", nil, "return transactions that are a parent of the supplied child (aka approvees)")
}
