// Copyright (c) 2018 ef-ds
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Package benchmark contains benchmark tests targeted to test the performance
// and efficiency of data structures.
package benchmark

import (
	"strconv"
	"testing"
)

// Stable tests the data structures performance by adding 1 item and removing it.
// Stable tests the data structures ability to handle constant add/remove over n iterations.
func (t *Tests) Stable(b *testing.B, initInstance func(), add func(v interface{}), remove func() (interface{}, bool), empty func() bool) {
	initInstance()
	for i := 0; i < fillCount; i++ {
		add(GetTestValue(i))
	}

	for i, test := range tests {
		// Doesn't run the first (0 items) test as 0 items makes no sense for this test.
		if i == 0 {
			continue
		}

		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				for i := 0; i < test.count; i++ {
					add(GetTestValue(i))
					tmp, tmp2 = remove()
				}

			}
		})
	}

	for !empty() {
		tmp, tmp2 = remove()
	}
}

// StableTestObject tests the data structures performance by adding 1 item and removing it.
// StableTestObject tests the data structures ability to handle constant add/remove over n iterations.
// StableTestObject is a copy of Stable that operates on *TestValue object which allows data structures that suport
// generics to not need to perform any type cast in the benchmark tests.
func (t *Tests) StableTestObject(b *testing.B, initInstance func(), add func(v *TestValue), remove func() (*TestValue, bool), empty func() bool) {
	initInstance()
	for i := 0; i < fillCount; i++ {
		add(GetTestValue(i))
	}

	for i, test := range tests {
		// Doesn't run the first (0 items) test as 0 items makes no sense for this test.
		if i == 0 {
			continue
		}

		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				for i := 0; i < test.count; i++ {
					add(GetTestValue(i))
					tmp, tmp2 = remove()
				}

			}
		})
	}

	for !empty() {
		tmp, tmp2 = remove()
	}
}
