// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.
//
// This file is auto-generated - do not modify

// +build amd64,!gccgo,!appengine

#include "textflag.h"

TEXT ·testAsm(SB),NOSPLIT,$0
	MOVQ src+0(FP), SI
	MOVQ len+8(FP), BX
	MOVB value+16(FP), AX
	CMPQ BX, $16
	JB loop
	PINSRB $0, AX, X0
	PXOR X1, X1
	PSHUFB X1, X0
bigloop:
	MOVOU -16(SI)(BX*1), X1
	PXOR X0, X1
	// PTEST X1, X1
	BYTE $0x66; BYTE $0x0f; BYTE $0x38; BYTE $0x17; BYTE $0xc9
	JNZ ret_false
	SUBQ $16, BX
	JZ ret_true
	CMPQ BX, $16
	JAE bigloop
loop:
	MOVB -1(SI)(BX*1), R15
	CMPB AX, R15
	JNE ret_false
	DECQ BX
	JNZ loop
ret_true:
	MOVB $1, ret+24(FP)
	RET
ret_false:
	MOVB $0, ret+24(FP)
	RET
