package main

import (
	"fmt"
	"strconv"
)


/*
MAIN.GO
Alvin NG
Stack machine for fun
 */

const (
	PSH = 101
	POP = 102
	ADD = 103
	SET = 104
	MOV = 105
	HLT	= 99
	AX=0
	BX=1
)
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


var register_ax int;
var register_bx int;
var stack[] int;
var register[] int ;

var end_status= false
func main() {
	var code_ptr =0
	for code_ptr < len(code_block) &&  end_status == false{
		var instr = fetchInstruction(code_block[:], code_ptr)
		switch instr {
		case MOV:
			code_ptr++
			var register = fetchInstruction(code_block[:], code_ptr)
			code_ptr++
			var value = fetchInstruction(code_block[:], code_ptr)
			switch(register){
			case AX:
				register_ax = value;
				break
			case BX:
				register_bx = value
				break
			}
		case PSH:
			code_ptr++
			var token = fetchInstruction(code_block[:], code_ptr)
			switch(token){
			case AX:
				stack= append(stack, register_ax)
				break
			case BX:
				stack = append(stack, register_bx)
				break
			default:
				//check if value is numeric
				stack = append(stack, token)
				break
			}
			break
		case POP:
			debugLog(" register len:"+ strconv.Itoa(len(register)))
			code_ptr++
			var register = fetchInstruction(code_block[:], code_ptr)
			var popped_val=0
			popped_val = stack[len(stack)-1]
			stack=stack[:len(stack)-1]

			switch(register){
			case AX:
				register_ax = popped_val
				break
			case BX:
				register_bx = popped_val
				break
			}
			debugLog("Pop "+strconv.Itoa(popped_val))

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
}

func fetchInstruction(code[] int, ptr int) int{
	return code[ptr]
}

/**
debugging logger
 */
func debugLog(text string ){
	fmt.Println(text)
}


func dumpRegister(){

}
