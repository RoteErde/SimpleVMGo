package main

import (
	"fmt"
)

/*
MAIN.GO
Alvin NG
Toy VM
*/

/*
var code_block = [...]int{MOV, AX, 10,
						PSH, AX,
						PSH, AX,
						MOV,BX ,11,
						PSH, BX ,
						POP,AX,
						POP,AX,
						POP,BX,
						HLT}
*/

var code_block = [...]int{PSH, 10,
	PSH, 11,
	PSH, 12,
	POP, AX,
	POP, BX,
	HLT}

var stack []int
var register []int
var vmRegister Registers
var end_status = false

func main() {
	var instructionMap = makeInstructions()
	for vmRegister.IP < len(code_block) || vmRegister.IP == 65535 {

		instructionMap[code_block[vmRegister.IP]](code_block[:], &vmRegister)
	}

}
func fetchInstruction(code []int, ptr int) int {
	return code[ptr]
}

/**
debugging logger
*/
func debugLog(text string) {
	fmt.Println(text)
}

func dumpRegister(registers Registers) {
	fmt.Println("Registers:")
	fmt.Println("AX:", vmRegister.AX)
	fmt.Println("BX", vmRegister.BX)

}
