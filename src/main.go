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
		instructionMap[code_block[vmRegister.SP]](code_block[:], &vmRegister)
	}
	/*
		for code_ptr < len(code_block) && end_status == false {
			var instr = fetchInstruction(code_block[:], code_ptr)
			switch instr {
			case LOAD:
				code_ptr++
				var register = fetchInstruction(code_block[:], code_ptr)
				code_ptr++
				var value = fetchInstruction(code_block[:], code_ptr)
				switch register {
				case AX:
					vmRegister.AX = value
					break
				case BX:
					vmRegister.BX = value
					break
				}
			case PSH:
				code_ptr++
				var token = fetchInstruction(code_block[:], code_ptr)
				switch token {
				case AX:
					stack = append(stack, vmRegister.AX)
					break
				case BX:
					stack = append(stack, vmRegister.BX)
					break
				default:
					//check if value is numeric
					stack = append(stack, token)
					break
				}
				break
			case POP:
				debugLog(" register len:" + strconv.Itoa(len(register)))
				code_ptr++
				var register = fetchInstruction(code_block[:], code_ptr)
				var poppedVal = 0
				poppedVal = stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				switch register {
				case AX:
					vmRegister.AX = poppedVal
					break
				case BX:
					vmRegister.BX = poppedVal
					break
				}
				debugLog("Pop " + strconv.Itoa(poppedVal))

				code_ptr++
				break
			case HLT:
				debugLog("end")
				end_status = true
				break
			default:
				code_ptr++
			}
			fmt.Print(" ")
		}
		dumpRegister(vmRegister)
	*/
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
