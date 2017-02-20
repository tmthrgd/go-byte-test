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
BenchmarkTest/32-8         	200000000	         6.60 ns/op	4845.48 MB/s
BenchmarkTest/128-8        	200000000	         8.39 ns/op	15256.81 MB/s
BenchmarkTest/1K-8         	50000000	        29.9 ns/op	34288.78 MB/s
BenchmarkTest/16K-8        	 5000000	       361 ns/op	45345.02 MB/s
BenchmarkTest/128K-8       	  500000	      3483 ns/op	37631.76 MB/s
BenchmarkTest/1M-8         	   30000	     41177 ns/op	25464.56 MB/s
BenchmarkTest/16M-8        	    2000	   1066258 ns/op	15734.66 MB/s
BenchmarkTest/128M-8       	     200	   8561493 ns/op	15676.91 MB/s
BenchmarkTest/512M-8       	      50	  34319085 ns/op	15643.51 MB/s
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
BenchmarkTest/32-8       35.8          6.60          -81.56%
BenchmarkTest/128-8      120           8.39          -93.01%
BenchmarkTest/1K-8       869           29.9          -96.56%
BenchmarkTest/16K-8      13760         361           -97.38%
BenchmarkTest/128K-8     109813        3483          -96.83%
BenchmarkTest/1M-8       878439        41177         -95.31%
BenchmarkTest/16M-8      14339512      1066258       -92.56%
BenchmarkTest/128M-8     114336485     8561493       -92.51%
BenchmarkTest/512M-8     457974138     34319085      -92.51%

benchmark                old MB/s     new MB/s     speedup
BenchmarkTest/32-8       893.80       4845.48      5.42x
BenchmarkTest/128-8      1058.03      15256.81     14.42x
BenchmarkTest/1K-8       1177.11      34288.78     29.13x
BenchmarkTest/16K-8      1190.62      45320.13     38.06x
BenchmarkTest/128K-8     1193.59      37631.76     31.53x
BenchmarkTest/1M-8       1193.68      25464.56     21.33x
BenchmarkTest/16M-8      1170.00      15734.66     13.45x
BenchmarkTest/128M-8     1173.88      15676.91     13.35x
BenchmarkTest/512M-8     1172.27      15643.51     13.34x
```

## License

Unless otherwise noted, the go-byte-test source files are distributed under the Modified BSD License
found in the LICENSE file.
