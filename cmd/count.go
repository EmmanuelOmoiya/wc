/*
MIT License

Copyright © 2024 Emmanuel Omoiya <emmanuelomoiya6@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

package cmd

import (
	"fmt"

	"github.com/EmmanuelOmoiya/wc/counters"
	"github.com/spf13/cobra"
)

var wcCmd = &cobra.Command{
	Use:   "count",
	Short: "count - print characters, line, word, and byte counts for each file",
	Long:  `Print  characters, line, word, and  byte counts for each FILE, and a total line if more than one FILE is specified.  A word is a non-zero-length sequence of characters delimited by white space.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			word, _ := cmd.Flags().GetBool("word")
			line, _ := cmd.Flags().GetBool("line")
			char, _ := cmd.Flags().GetBool("character")
			byte, _ := cmd.Flags().GetBool("byte")

			if byte {
				counters.CountByte(args[0])
			} else if char {
				counters.CountChars((args[0]))
			} else if word {
				counters.CountWord(args[0])
			} else if line {
				counters.CountLine(args[0])
			} else {
				counters.CountWord(args[0])
				fmt.Println("")
				counters.CountByte(args[0])
				fmt.Println("")
				counters.CountLine(args[0])
			}
		} else {
			fmt.Println("Invalid file, kindly specify path to file e.g wc count test.txt")
		}
	},
}

func init() {
	rootCmd.AddCommand(wcCmd)
	wcCmd.Flags().BoolP("word", "w", false, "Display the number of words in that file")
	wcCmd.Flags().BoolP("line", "l", false, "Display the number of lines in that file")
	wcCmd.Flags().BoolP("character", "m", false, "Display the number of characters in that file")
	wcCmd.Flags().BoolP("byte", "c", false, "Display the number of bytes in that file")
}
