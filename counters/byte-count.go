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

package counters

import (
	"bufio"
	"fmt"
	"os"
	"sync/atomic"
)

func CountByte(fileName string) (int64, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return -1, fmt.Errorf("File not found")
	}
	defer file.Close()

	var byteCount int64

	scanner := bufio.NewScanner(bufio.NewReader(file))

	scanner.Split(bufio.ScanBytes)

	for scanner.Scan() {
		atomic.AddInt64(&byteCount, 1)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v %v", byteCount, fileName)

	return byteCount, nil
}
