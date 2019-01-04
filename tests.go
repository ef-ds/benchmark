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

// Package benchmark contains benchmark tests targetted to test the performance
// and efficiency of data structures.
package benchmark

import (
	"strconv"
	"testing"
)

// Benchmark contains benchmark tests targeted to test the performance and efficiency of data structures.
type Benchmark struct {
}

// Fill test the data structures performance by sequentially adding n items to the data structure and then removing all added items.
// Fill tests the data structures ability for quickly expand and shrink.
func (t *Benchmark) Fill(b *testing.B, initInstance func(), add func(v interface{}), remove func() (interface{}, bool), empty func() bool) {
	for _, test := range tests {
		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				initInstance()
				for i := 0; i < test.count; i++ {
					add(GetTestValue(i))
				}
				for !empty() {
					tmp, tmp2 = remove()
				}
			}
		})
	}
}

// Refill test the data structures performance by sequentially adding n items to the data structure and then removing all added items
// repeating the test 100 times using the same data structure instance.
// Refill tests the data structures ability to fill again once it has been filled and emptied.
func (t *Benchmark) Refill(b *testing.B, initInstance func(), add func(v interface{}), remove func() (interface{}, bool), empty func() bool) {
	for i, test := range tests {
		// Doesn't run the first (0 items) and last (1mi) items tests
		// as 0 items makes no sense for this test and 1mi is too slow.
		if i == 0 || i > 6 {
			continue
		}

		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			initInstance()
			for n := 0; n < b.N; n++ {
				for n := 0; n < refillCount; n++ {
					for i := 0; i < test.count; i++ {
						add(GetTestValue(i))
					}
					for !empty() {
						tmp, tmp2 = remove()
					}
				}
			}
		})
	}
}

// RefillFull test the data structures performance by sequentially adding n items to the data structures and then removing all added items
// repeating the test 100 times using the same data structure instance. But before running the test, fills the data structures
// with n items.
// RefillFull rests the data structures ability to fill again once it has been filled and emptied back to a certain level.
func (t *Benchmark) RefillFull(b *testing.B, initInstance func(), add func(v interface{}), remove func() (interface{}, bool), empty func() bool) {
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

// SlowDecrease tests the data structures performance by sequentially adding 2 items and then removing 1.
// SlowDecrease tests the data structures ability to slowly expand while removing some elements from the data structure.
func (t *Benchmark) SlowDecrease(b *testing.B, initInstance func(), add func(v interface{}), remove func() (interface{}, bool), empty func() bool) {
	initInstance()
	for _, test := range tests {
		items := test.count / 2
		for i := 0; i <= items; i++ {
			add(GetTestValue(i))
		}
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
					if !empty() {
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

// SlowIncrease tests the data structures performance by filling the data structures with n items, and then
// sequentially removing 2 items and adding 1.
// SlowIncrease tests the data structures ability to slowly shrink while adding some elements to the data structure.
func (t *Benchmark) SlowIncrease(b *testing.B, initInstance func(), add func(v interface{}), remove func() (interface{}, bool), empty func() bool) {
	for i, test := range tests {
		// Doesn't run the first (0 items) test as 0 items makes no sense for this test.
		if i == 0 {
			continue
		}

		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				initInstance()
				for i := 0; i < test.count; i++ {
					add(GetTestValue(i))
					add(GetTestValue(i))
					tmp, tmp2 = remove()
				}
				for !empty() {
					tmp, tmp2 = remove()
				}
			}
		})
	}
}

// Stable tests the data structures performance by adding 1 item and removing it.
// Stable  tests the data structures ability to handle constant add/remove over n iterations.
func (t *Benchmark) Stable(b *testing.B, initInstance func(), add func(v interface{}), remove func() (interface{}, bool), empty func() bool) {
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

// Microservice tests the data structures performance by simulating the data structure being used by microservice
// and serverless systems when running in production environments.
func (t *Benchmark) Microservice(b *testing.B, initInstance func(), add func(v interface{}), remove func() (interface{}, bool), empty func() bool) {
	for _, test := range tests {
		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				initInstance()

				// Simulate stable traffic
				for i := 0; i < test.count; i++ {
					add(GetTestValue(i))
					remove()
				}

				// Simulate slowly increasing traffic
				for i := 0; i < test.count; i++ {
					add(GetTestValue(i))
					add(GetTestValue(i))
					remove()
				}

				// Simulate slowly decreasing traffic, bringing traffic back to normal
				for i := 0; i < test.count; i++ {
					remove()
					if !empty() {
						remove()
					}
					add(GetTestValue(i))
				}

				// Simulate quick traffic spike (DDOS attack, etc)
				for i := 0; i < test.count; i++ {
					add(GetTestValue(i))
				}

				// Simulate stable traffic while at high traffic
				for i := 0; i < test.count; i++ {
					add(GetTestValue(i))
					remove()
				}

				// Simulate going back to normal (DDOS attack fended off)
				for i := 0; i < test.count; i++ {
					remove()
				}

				// Simulate stable traffic (now that is back to normal)
				for i := 0; i < test.count; i++ {
					add(GetTestValue(i))
					remove()
				}
			}
		})
	}
}
