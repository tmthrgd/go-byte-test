# go-byte-test

[![GoDoc](https://godoc.org/github.com/tmthrgd/go-byte-test?status.svg)](https://godoc.org/github.com/tmthrgd/go-byte-test)
[![Build Status](https://travis-ci.org/tmthrgd/go-byte-test.svg?branch=master)](https://travis-ci.org/tmthrgd/go-byte-test)

An efficient byte test implementation for Golang.

It is SSE accelerated equivalent of the following function:
```
// Test returns true iff each byte in data is equal to value.
func Test(data []byte, value byte) bool {
	for _, v := range data {
		if v != value {
			return false
		}
	}

	return true
}
```

## Download

```
go get github.com/tmthrgd/go-byte-test
```

## Benchmark

```
BenchmarkTest/32-8  	200000000	         6.35 ns/op	5040.66 MB/s
BenchmarkTest/128-8 	100000000	        12.4 ns/op	10295.25 MB/s
BenchmarkTest/1K-8  	20000000	        68.0 ns/op	15066.47 MB/s
BenchmarkTest/16K-8 	 2000000	       826 ns/op	19815.88 MB/s
BenchmarkTest/128K-8         	  200000	      6577 ns/op	19926.39 MB/s
BenchmarkTest/1M-8           	   30000	     58091 ns/op	18050.46 MB/s
BenchmarkTest/16M-8          	    1000	   1418267 ns/op	11829.37 MB/s
BenchmarkTest/128M-8         	     100	  10938682 ns/op	12270.01 MB/s
BenchmarkTest/512M-8         	      30	  43501966 ns/op	12341.30 MB/s
```

```
BenchmarkGoTest/32-8         	50000000	        35.8 ns/op	 893.80 MB/s
BenchmarkGoTest/128-8        	10000000	       120 ns/op	1058.03 MB/s
BenchmarkGoTest/1K-8         	 2000000	       869 ns/op	1177.11 MB/s
BenchmarkGoTest/16K-8        	  100000	     13760 ns/op	1190.62 MB/s
BenchmarkGoTest/128K-8       	   10000	    109813 ns/op	1193.59 MB/s
BenchmarkGoTest/1M-8         	    2000	    878439 ns/op	1193.68 MB/s
BenchmarkGoTest/16M-8        	     100	  14339512 ns/op	1170.00 MB/s
BenchmarkGoTest/128M-8       	      10	 114336485 ns/op	1173.88 MB/s
BenchmarkGoTest/512M-8       	       3	 457974138 ns/op	1172.27 MB/s
```

go -> go-byte-test:
```
benchmark                old ns/op     new ns/op     delta
BenchmarkTest/32-8       35.8          6.35          -82.26%
BenchmarkTest/128-8      120           12.4          -89.67%
BenchmarkTest/1K-8       869           68.0          -92.17%
BenchmarkTest/16K-8      13760         826           -94.00%
BenchmarkTest/128K-8     109813        6577          -94.01%
BenchmarkTest/1M-8       878439        58091         -93.39%
BenchmarkTest/16M-8      14339512      1418267       -90.11%
BenchmarkTest/128M-8     114336485     10938682      -90.43%
BenchmarkTest/512M-8     457974138     43501966      -90.50%

benchmark                old MB/s     new MB/s     speedup
BenchmarkTest/32-8       893.80       5040.66      5.64x
BenchmarkTest/128-8      1058.03      10295.25     9.73x
BenchmarkTest/1K-8       1177.11      15066.47     12.80x
BenchmarkTest/16K-8      1190.62      19815.88     16.64x
BenchmarkTest/128K-8     1193.59      19926.39     16.69x
BenchmarkTest/1M-8       1193.68      18050.46     15.12x
BenchmarkTest/16M-8      1170.00      11829.37     10.11x
BenchmarkTest/128M-8     1173.88      12270.01     10.45x
BenchmarkTest/512M-8     1172.27      12341.30     10.53x
```

## License

Unless otherwise noted, the go-byte-test source files are distributed under the Modified BSD License
found in the LICENSE file.
