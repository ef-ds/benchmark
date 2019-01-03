# benchmark [![Build Status](https://travis-ci.com/ef-ds/benchmark.svg?branch=master)](https://travis-ci.com/ef-ds/benchmark)[![Go Report Card](https://goreportcard.com/badge/github.com/ef-ds/benchmark)](https://goreportcard.com/report/github.com/ef-ds/benchmark)  [![GoDoc](https://godoc.org/github.com/ef-ds/benchmark?status.svg)](https://godoc.org/github.com/ef-ds/benchmark)

Package benchmark contains benchmark tests targetted to test the performance and efficiency of data structures.

## Install
From a configured [Go environment](https://golang.org/doc/install#testing):
```sh
go get -u github.com/ef-ds/benchmark
```

If you are using dep:
```sh
dep ensure -add github.com/ef-ds/benchmark@1.0.1
```

We recommend to target only released versions for production use.


## How to Use

Below runs the benchmark [Fill](https://github.com/ef-ds/benchmark/blob/master/tests.go) tests using the standard [list package](https://github.com/golang/go/tree/master/src/container/list) as a LIFO stack.

```go
package main

import (
	"fmt"

	"github.com/ef-ds/benchmark"
)

func BenchmarkFillList(b *testing.B) {
	var l *list.List
	var tests benchmark.Benchmark
	tests.Fill(
		b,
		func() {
			l = list.New()
		},
		func(v interface{}) {
			l.PushBack(v)
		},
		func() (interface{}, bool) {
			return l.Remove(l.Back()), true
		},
		func() bool {
			return l.Front() == nil
		},
	)
}
```

## Tests
The benchmark tests are composed of test suites and ranges.


## Test Suites
The test suites were designed to test the data structures with different add/remove patterns under different scenarios such as low and high stress.

- [Fill](tests.go): test the data structures performance by sequentially adding n items to the data structure and then removing all added items. Tests the data structures ability for quickly expand and shrink.
- [Refill](tests.go): same test as Fill, but repeat the test 100 times using the same data structure instance. Tests the data structures ability to fill again once it has been filled and emptied.
- [RefillFull](tests.go): same test as Refill, but before running the test, fills the data structures with n items to fill at least three internal slices. Tests the data structures ability to fill again once it has been filled and emptied back to a certain level (10k items).
- [SlowIncrease](tests.go): test the data structures performance by sequentially adding 2 items and then removing 1. Tests the data structures ability to slowly expand while removing some elements from the data structure.
- [SlowDecrease](tests.go): test the data structures performance by filling the data structures with n items to fill at least three internal slices, and then sequentially removing 2 items and adding 1. Tests the data structures ability to slowly shrink while adding some elements to the data structure.
- [Stable](tests.go): Add 1 item to the data structure and remove it. Tests the data structures ability to handle constant push/pop over n iterations.


### The Microservice Test
It is very common on production [Microservices](https://en.wikipedia.org/wiki/Microservices) and [serverless](https://en.wikipedia.org/wiki/Serverless_computing) systems to use more resources, be it memory or CPU, as the traffic it is serving increases. Keeping this fact in mind, this is a composite test designed to test the data structures in a production like microservice scenario. The test idea is that every time the Microservice using the data structure receives a request, it would add an item to the data structure. As soon as the request is served, the Microservice removes an item from the data structure.

The test start by running a stable test to simulate stable traffic (i.e. the system is able to handle the traffic without stress).

Next the test simulates the system facing some stress in the form of a slowly increasing traffic where the system is forced to use more resources (more items in the data structure) to serve the extra traffic.

Next the test simulates the system handling decreasing traffic where more items are removed from the data structures, moving back to a low traffic level.

Next the test simulates the system handling quick spikes in traffic (i.e. DDOS attack) where n items are added to the data structure but none is removed.

Next the test simulates the system handling the traffic while under stress (high constant traffic with high number of items in the data structure).

Next the test simulates the system handling the traffic going back to normal quickly (i.e. DDOS attack fended off).

Finally, the test simulates the system handling the regular, stable, traffic again.

The Microservice test can be found [here.](tests.go)


## Test Ranges

The test ranges are designed to test the data structures with different loads. The tests will add and remove below number of items to the data structures according to each test suites pattern.

- 0 items
- 1 items
- 10 items
- 100 items
- 1000 items //1k
- 10000 items // 10k
- 100000 items // 100k
- 1000000 items // 1mi

The 0 items test runs only for the [Fill](tests.go) and [Microservice](tests.go) tests and is designed to test the data structures initialization time only.


## Tests Type
In order to try to simulate real world usage scenarios as much as possible, all tests create and add/remove below testValue struct to the data structures, as structs being pushed into the data structures should be the most common scenario.

```
// testValue is used as the value added in each push call to the data structures.
// A struct is being used as structs should be more representative of real
// world uses of a data structure. A second f2 field was added as the users structs
// are likely to contain more than one field.
type testValue struct {
	count int
	f2    int
}
```

## Supported Go Versions
See [supported_go_versions.md](https://github.com/ef-ds/docs/blob/master/supported_go_versions.md).

## License
MIT, see [LICENSE](LICENSE).

"Use, abuse, have fun and contribute back!"
