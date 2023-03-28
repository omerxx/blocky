/*
Copyright © 2023 Omer Hamerman <omer.hamerman@zesty.co>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	block "blocky/pkg"
	"blocky/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/txn2/txeh"
)

// blockCmd represents the block command
var blockCmd = &cobra.Command{
	Use:   "block",
	Short: "The block subcommand will be creating the blocks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		hosts, err := txeh.NewHostsDefault()
		if err != nil {
			panic(err)
		}
		viperSites := viper.GetStringSlice("sites")
		viperSites = utils.AddWww(viperSites)

		// NOTE: so far sites were not using the www prefix
		//we probably need to add it automatically
		block.BlockSites(hosts, viperSites)
	},
}

func init() {
	rootCmd.AddCommand(blockCmd)
}
