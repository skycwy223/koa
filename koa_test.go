/*
 * Copyright 2018 De-labtory
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package koa

import (
	"errors"
	"os"
	"testing"

	"bytes"
	"encoding/hex"

	"github.com/DE-labtory/koa/abi"
	"github.com/DE-labtory/koa/translate"
)

type testData struct {
	fileName string
	asm      *translate.Asm
	err      error
}

func defineAsm() []testData {
	return []testData{
		// TODO: implement test cases with function jumper :-)
		//{
		//	fileName: "test/hello.koa",
		//	asm: &translate.Asm{
		//		AsmCodes: []translate.AsmCode{
		//			{
		//				RawByte: []byte{byte(opcode.Push)},
		//				Value:   "Push",
		//			},
		//			{
		//				RawByte: []byte{0x22, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x21, 0x22},
		//				Value:   "2268656c6c6f2122",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Returning)},
		//				Value:   "Returning",
		//			},
		//		},
		//	},
		//	err: nil,
		//},
		{
			fileName: "test/jun.koa",
			asm:      nil,
			err:      errors.New("[junbeomlee] definition doesn't exist"),
		},
		//{
		//	fileName: "test/add1.koa",
		//	asm: &translate.Asm{
		//		AsmCodes: []translate.AsmCode{
		//			{
		//				RawByte: []byte{byte(opcode.Push)},
		//				Value:   "Push",
		//			},
		//			{
		//				RawByte: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
		//				Value:   "0000000000000001",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Push)},
		//				Value:   "Push",
		//			},
		//			{
		//				RawByte: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02},
		//				Value:   "0000000000000002",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Add)},
		//				Value:   "Add",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Push)},
		//				Value:   "Push",
		//			},
		//			{
		//				RawByte: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08},
		//				Value:   "0000000000000008",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Push)},
		//				Value:   "Push",
		//			},
		//			{
		//				RawByte: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		//				Value:   "0000000000000000",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Mstore)},
		//				Value:   "Mstore",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Push)},
		//				Value:   "Push",
		//			},
		//			{
		//				RawByte: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03},
		//				Value:   "0000000000000003",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Push)},
		//				Value:   "Push",
		//			},
		//			{
		//				RawByte: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04},
		//				Value:   "0000000000000004",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Add)},
		//				Value:   "Add",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Push)},
		//				Value:   "Push",
		//			},
		//			{
		//				RawByte: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08},
		//				Value:   "0000000000000008",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Push)},
		//				Value:   "Push",
		//			},
		//			{
		//				RawByte: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08},
		//				Value:   "0000000000000008",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Mstore)},
		//				Value:   "Mstore",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Push)},
		//				Value:   "Push",
		//			},
		//			{
		//				RawByte: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05},
		//				Value:   "0000000000000005",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Push)},
		//				Value:   "Push",
		//			},
		//			{
		//				RawByte: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x06},
		//				Value:   "0000000000000006",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Mul)},
		//				Value:   "Mul",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Push)},
		//				Value:   "Push",
		//			},
		//			{
		//				RawByte: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08},
		//				Value:   "0000000000000008",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Push)},
		//				Value:   "Push",
		//			},
		//			{
		//				RawByte: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10},
		//				Value:   "0000000000000010",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Mstore)},
		//				Value:   "Mstore",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Push)},
		//				Value:   "Push",
		//			},
		//			{
		//				RawByte: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08},
		//				Value:   "0000000000000008",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Push)},
		//				Value:   "Push",
		//			},
		//			{
		//				RawByte: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04},
		//				Value:   "0000000000000004",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Div)},
		//				Value:   "Div",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Push)},
		//				Value:   "Push",
		//			},
		//			{
		//				RawByte: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08},
		//				Value:   "0000000000000008",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Push)},
		//				Value:   "Push",
		//			},
		//			{
		//				RawByte: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x18},
		//				Value:   "0000000000000018",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Mstore)},
		//				Value:   "Mstore",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Push)},
		//				Value:   "Push",
		//			},
		//			{
		//				RawByte: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08},
		//				Value:   "0000000000000008",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Push)},
		//				Value:   "Push",
		//			},
		//			{
		//				RawByte: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		//				Value:   "0000000000000000",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Mload)},
		//				Value:   "Mload",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Push)},
		//				Value:   "Push",
		//			},
		//			{
		//				RawByte: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08},
		//				Value:   "0000000000000008",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Push)},
		//				Value:   "Push",
		//			},
		//			{
		//				RawByte: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08},
		//				Value:   "0000000000000008",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Mload)},
		//				Value:   "Mload",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Push)},
		//				Value:   "Push",
		//			},
		//			{
		//				RawByte: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08},
		//				Value:   "0000000000000008",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Push)},
		//				Value:   "Push",
		//			},
		//			{
		//				RawByte: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10},
		//				Value:   "0000000000000010",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Mload)},
		//				Value:   "Mload",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Sub)},
		//				Value:   "Sub",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Push)},
		//				Value:   "Push",
		//			},
		//			{
		//				RawByte: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08},
		//				Value:   "0000000000000008",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Push)},
		//				Value:   "Push",
		//			},
		//			{
		//				RawByte: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x18},
		//				Value:   "0000000000000018",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Mload)},
		//				Value:   "Mload",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Mul)},
		//				Value:   "Mul",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Add)},
		//				Value:   "Add",
		//			},
		//			{
		//				RawByte: []byte{byte(opcode.Returning)},
		//				Value:   "Returning",
		//			},
		//		},
		//	},
		//	err: nil,
		//},
	}
}

func readFile(fileName string) (string, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, os.FileMode(644))
	if err != nil {
		return "", err
	}

	fs, err := file.Stat()
	if err != nil {
		return "", err
	}

	data := make([]byte, fs.Size())
	if err != nil {
		return "", err
	}

	_, err = file.Read(data)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func TestCompile(t *testing.T) {
	tests := defineAsm()

	for i, test := range tests {
		str, err := readFile(test.fileName)
		if err != nil {
			continue
		}

		asm, _, err := Compile(str)

		// TODO : after implements all functions in compile
		if err != nil && err.Error() != test.err.Error() {
			t.Fatalf("[test %d] - TestCompile() returns wrong error.\nexpected=%v\ngot=%v", i, test.err, err)
		}

		if test.asm != nil && !asm.Equal(*test.asm) {
			t.Fatalf("[test %d] - TestCompile() returns wrong asm.\nexpected=%v\ngot=%v", i, test.asm, asm)
		}
	}
}

func TestExecute(t *testing.T) {
	/*
		contract {
			func addVariable() int {
				int a = 5
				int b = 10
				return a + b
			}

			func addNative() int {
				return 5 + 10
			}

			func addArgs(a int, b int) int {
				return a + b
			}
		}
	*/
	firstContractRawBytecode, err := hex.DecodeString("2100000000000000202421000000000000001e25312100000000652f6077141521000000000000001f30312100000000a82ed9f7141521000000000000003930312100000000c6be6f42141521000000000000003f30332100000000000000052100000000000000082100000000000000002321000000000000000a210000000000000008210000000000000008232100000000000000082100000000000000002221000000000000000821000000000000000822012721000000000000000521000000000000000a01272100000000000000002621000000000000000821000000000000001023210000000000000001262100000000000000082100000000000000182321000000000000000821000000000000001022210000000000000008210000000000000018220127")
	if err != nil {
		t.Error(err)
	}

	/*
		contract {
			func hello() string{
				return "hello!"
			}
		}
	*/
	secondContractRawBytecode, err := hex.DecodeString("2100000000000000002421000000000000000e2531210000000019ff1d21141521000000000000000f3033212268656c6c6f212227")
	if err != nil {
		t.Error(err)
	}

	firstArgs1, err := abi.Encode()
	firstArgs2, err := abi.Encode()
	firstArgs3, err := abi.Encode(5, 10)
	second1, err := abi.Encode()

	tests := []struct {
		RawBytecode []byte
		Func        []byte
		Args        []byte
		output      []byte
	}{
		{
			RawBytecode: firstContractRawBytecode,
			Func:        abi.Selector("addVariable()"),
			Args:        firstArgs1,
			output:      []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0f},
		},
		{
			RawBytecode: firstContractRawBytecode,
			Func:        abi.Selector("addNative()"),
			Args:        firstArgs2,
			output:      []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0f},
		},
		{
			RawBytecode: firstContractRawBytecode,
			Func:        abi.Selector("addArgs(int,int)"),
			Args:        firstArgs3,
			output:      []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0f},
		},
		{
			RawBytecode: secondContractRawBytecode,
			Func:        abi.Selector("hello()"),
			Args:        second1,
			output:      []byte{0x22, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x21, 0x22},
		},
	}

	for _, test := range tests {
		output, err := Execute(test.RawBytecode, test.Func, test.Args)
		if err != nil {
			t.Error(err)
		}

		if !bytes.Equal(test.output, output) {
			t.Errorf("Invalid output - expected=%x, got=%x ", test.output, output)
		}
	}
}
