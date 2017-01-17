// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package bytetest

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"

	"github.com/tmthrgd/go-memset"
)

func TestCorrect(t *testing.T) {
	if err := quick.Check(Test, &quick.Config{
		Values: func(args []reflect.Value, rand *rand.Rand) {
			off := rand.Intn(32)

			data := make([]byte, off+1+rand.Intn(128*1024))

			value := byte(rand.Intn(0x100))
			memset.Memset(data, value)

			args[0] = reflect.ValueOf(data[off:])
			args[1] = reflect.ValueOf(value)
		},

		MaxCountScale: 100,
	}); err != nil {
		t.Error(err)
	}

	if err := quick.Check(func(data []byte, value byte) bool {
		return !Test(data, value)
	}, &quick.Config{
		Values: func(args []reflect.Value, rand *rand.Rand) {
			off := rand.Intn(32)

			data := make([]byte, off+1+rand.Intn(128*1024))

			value := byte(rand.Intn(0x100))
			memset.Memset(data, value)

			data[len(data)-1] ^= 0xff

			args[0] = reflect.ValueOf(data[off:])
			args[1] = reflect.ValueOf(value)
		},

		MaxCountScale: 100,
	}); err != nil {
		t.Error(err)
	}
}

var sizes = []struct {
	name string
	l    int
}{
	{"32", 32},
	{"128", 128},
	{"1K", 1 * 1024},
	{"16K", 16 * 1024},
	{"128K", 128 * 1024},
	{"1M", 1024 * 1024},
	{"16M", 16 * 1024 * 1024},
	{"128M", 128 * 1024 * 1024},
	{"512M", 512 * 1024 * 1024},
}

func BenchmarkTest(b *testing.B) {
	for _, size := range sizes {
		b.Run(size.name, func(b *testing.B) {
			data := make([]byte, size.l)

			value := byte(rand.Intn(0x100))
			memset.Memset(data, value)

			b.SetBytes(int64(size.l))
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				if !Test(data, value) {
					b.Fatal("Test failed")
				}
			}
		})
	}
}

func BenchmarkGoTest(b *testing.B) {
	for _, size := range sizes {
		b.Run(size.name, func(b *testing.B) {
			data := make([]byte, size.l)

			value := byte(rand.Intn(0x100))
			memset.Memset(data, value)

			b.SetBytes(int64(size.l))
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				for _, v := range data {
					if v != value {
						b.Fatal("test loop failed")
					}
				}
			}
		})
	}
}
