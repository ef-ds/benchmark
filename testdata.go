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

package benchmark

// TestValue is used as the value added in each push call to the queues.
// A struct is being used as structs should be more representative of real
// world uses of a queue. A second f2 field was added as the users structs
// are likely to contain more than one field.
type TestValue struct {
	count int
	f2    int
}

// testData contains the number of items to add to the queues in each test.
type testData struct {
	count int
}

var (
	tests = []testData{
		{count: 0},
		{count: 1},
		{count: 10},
		{count: 100},
		{count: 1000},    // 1k
		{count: 10000},   //10k
		{count: 100000},  // 100k
		{count: 1000000}, // 1mi
	}

	// Used to store temp values, avoiding any compiler optimizations.
	tmp  interface{}
	tmp2 bool

	fillCount   = 10000
	refillCount = 100
)

// Helper methods-----------------------------------------------------------------------------------

// GetTestValue returns an initialized instance of *TestValue.
func GetTestValue(i int) *TestValue {
	return &TestValue{
		count: i,
		f2:    1, // Initializes f2 to some random value (1).
	}
}
