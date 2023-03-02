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

// RefillFull test the data structures performance by sequentially adding n items to the data structures and then removing all added items
// repeating the test 100 times using the same data structure instance. But before running the test, fills the data structures
// with n items.
// RefillFull rests the data structures ability to fill again once it has been filled and emptied back to a certain level.
func (t *Tests) RefillFull(b *testing.B, initInstance func(), add func(v interface{}), remove func() (interface{}, bool), empty func() bool) {
	initInstance()
	for i := 0; i < fillCount; i++ {
		add(GetTestValue(i))
	}

	for i, test := range tests {
		// Doesn't run the first (0 items) and last (1mi) items tests
		// as 0 items makes no sense for this test and 1mi is too slow.
		if i == 0 || i > 6 {
			continue
		}

		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				for k := 0; k < refillCount; k++ {
					for i := 0; i < test.count; i++ {
						add(GetTestValue(i))
					}
					for i := 0; i < test.count; i++ {
						tmp, tmp2 = remove()
					}
				}
			}
		})
	}

	for !empty() {
		tmp, tmp2 = remove()
	}
}

// RefillFullTestObject test the data structures performance by sequentially adding n items to the data structures and then removing all added items
// repeating the test 100 times using the same data structure instance. But before running the test, fills the data structures
// with n items.
// RefillFullTestObject rests the data structures ability to fill again once it has been filled and emptied back to a certain level.
// RefillFullTestObject is a copy of RefillFull that operates on *TestValue object which allows data structures that suport
// generics to not need to perform any type cast in the benchmark tests.
func (t *Tests) RefillFullTestObject(b *testing.B, initInstance func(), add func(v *TestValue), remove func() (*TestValue, bool), empty func() bool) {
	initInstance()
	for i := 0; i < fillCount; i++ {
		add(GetTestValue(i))
	}

	for i, test := range tests {
		// Doesn't run the first (0 items) and last (1mi) items tests
		// as 0 items makes no sense for this test and 1mi is too slow.
		if i == 0 || i > 6 {
			continue
		}

		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				for k := 0; k < refillCount; k++ {
					for i := 0; i < test.count; i++ {
						add(GetTestValue(i))
					}
					for i := 0; i < test.count; i++ {
						tmp, tmp2 = remove()
					}
				}
			}
		})
	}

	for !empty() {
		tmp, tmp2 = remove()
	}
}
