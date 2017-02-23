// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

// +build ignore

package main

import "github.com/tmthrgd/asm"

const header = `// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.
//
// This file is auto-generated - do not modify

// +build amd64,!gccgo,!appengine
`

func testAsm(a *asm.Asm) {
	a.NewFunction("testAsm")
	a.NoSplit()

	dst := a.Argument("src", 8)
	length := a.Argument("len", 8)
	value := a.Argument("value", 8)
	ret := a.Argument("ret", 8)

	a.Start()

	massiveloop := a.NewLabel("massiveloop")
	hugeloop := a.NewLabel("hugeloop")
	bigloop := a.NewLabel("bigloop")
	loop := a.NewLabel("loop")
	retLabel := a.NewLabel("ret")
	retTrue := retLabel.Suffix("true")
	retFalse := retLabel.Suffix("false")

	si, ax, cx := asm.SI, asm.AX, asm.BX

	a.Movq(si, dst)
	a.Movq(cx, length)
	a.Movb(ax, value)

	a.Cmpq(asm.Constant(16), cx)
	a.Jb(loop)

	a.Pinsrb(asm.X0, ax, asm.Constant(0))
	a.Pxor(asm.X1, asm.X1)
	a.Pshufb(asm.X0, asm.X1)

	a.Cmpq(asm.Constant(64), cx)
	a.Jb(bigloop)

	a.Cmpq(asm.Constant(128), cx)
	a.Jb(hugeloop)

	a.Label(massiveloop)

	a.Movou(asm.X1, asm.Address(si, cx, asm.SX1, -16))
	a.Movou(asm.X2, asm.Address(si, cx, asm.SX1, -32))
	a.Movou(asm.X3, asm.Address(si, cx, asm.SX1, -48))
	a.Movou(asm.X4, asm.Address(si, cx, asm.SX1, -64))
	a.Movou(asm.X5, asm.Address(si, cx, asm.SX1, -80))
	a.Movou(asm.X6, asm.Address(si, cx, asm.SX1, -96))
	a.Movou(asm.X7, asm.Address(si, cx, asm.SX1, -112))
	a.Movou(asm.X8, asm.Address(si, cx, asm.SX1, -128))

	a.Pxor(asm.X1, asm.X0)
	a.Pxor(asm.X2, asm.X0)
	a.Pxor(asm.X3, asm.X0)
	a.Pxor(asm.X4, asm.X0)
	a.Pxor(asm.X5, asm.X0)
	a.Pxor(asm.X6, asm.X0)
	a.Pxor(asm.X7, asm.X0)
	a.Pxor(asm.X8, asm.X0)

	a.Por(asm.X1, asm.X2)
	a.Por(asm.X1, asm.X3)
	a.Por(asm.X1, asm.X4)
	a.Por(asm.X1, asm.X5)
	a.Por(asm.X1, asm.X6)
	a.Por(asm.X1, asm.X7)
	a.Por(asm.X1, asm.X8)

	a.Ptest(asm.X1, asm.X1)
	a.Jnz(retFalse)

	a.Subq(cx, asm.Constant(128))
	a.Jz(retTrue)

	a.Cmpq(asm.Constant(128), cx)
	a.Jae(massiveloop)

	a.Cmpq(asm.Constant(16), cx)
	a.Jb(loop)

	a.Cmpq(asm.Constant(64), cx)
	a.Jb(bigloop)

	a.Label(hugeloop)

	a.Movou(asm.X1, asm.Address(si, cx, asm.SX1, -16))
	a.Movou(asm.X2, asm.Address(si, cx, asm.SX1, -32))
	a.Movou(asm.X3, asm.Address(si, cx, asm.SX1, -48))
	a.Movou(asm.X4, asm.Address(si, cx, asm.SX1, -64))

	a.Pxor(asm.X1, asm.X0)
	a.Pxor(asm.X2, asm.X0)
	a.Pxor(asm.X3, asm.X0)
	a.Pxor(asm.X4, asm.X0)

	a.Por(asm.X1, asm.X2)
	a.Por(asm.X1, asm.X3)
	a.Por(asm.X1, asm.X4)

	a.Ptest(asm.X1, asm.X1)
	a.Jnz(retFalse)

	a.Subq(cx, asm.Constant(64))
	a.Jz(retTrue)

	a.Cmpq(asm.Constant(64), cx)
	a.Jae(hugeloop)

	a.Cmpq(asm.Constant(16), cx)
	a.Jb(loop)

	a.Label(bigloop)

	a.Movou(asm.X1, asm.Address(si, cx, asm.SX1, -16))

	a.Pxor(asm.X1, asm.X0)
	a.Ptest(asm.X1, asm.X1)
	a.Jnz(retFalse)

	a.Subq(cx, asm.Constant(16))
	a.Jz(retTrue)

	a.Cmpq(asm.Constant(16), cx)
	a.Jae(bigloop)

	a.Label(loop)

	a.Movb(asm.R15, asm.Address(si, cx, asm.SX1, -1))

	a.Cmpb(asm.R15, ax)
	a.Jne(retFalse)

	a.Decq(cx)
	a.Jnz(loop)

	a.Label(retTrue)
	a.Movb(ret, asm.Constant(0x01))
	a.Ret()

	a.Label(retFalse)
	a.Movb(ret, asm.Constant(0x00))
	a.Ret()
}

func main() {
	if err := asm.Do("test_amd64.s", header, testAsm); err != nil {
		panic(err)
	}
}
