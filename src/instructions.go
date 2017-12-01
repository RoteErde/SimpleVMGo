package main

type sub func([]int, *Registers)

/**
Alvin Ng : skeleton code to map instructions to functions

aim is to avoid lengthy case statements
*/
func makeInstructions() map[int]sub {
	mappedFunctions := map[int]sub{
		PSH: pushInstr,
	}
	return mappedFunctions
}

func popInstr(bin_code []int, registers *Registers) {

}

func movInstr(bin_code []int, registers *Registers) {
	registers.IP++
	var opcode1 = fetchInstruction(bin_code[:], registers.IP)
	registers.IP++
	var opcode2 = fetchInstruction(bin_code[:], registers.IP)
	switch opcode1 {
	case AX:

		break
	case BX:
		break
	default:
		// not an oprand
		switch opcode2 {
		case AX:
			break
		case BX:
			break
		default:
			//likely opcode to be an error
			break

		}
		break
	}
}

func pushInstr(bin_code []int, registers *Registers) {
	//increment code pointer register
	registers.IP++
	var token = fetchInstruction(bin_code[:], registers.IP)
	// on successful parsing, add value to the
	// stack and increment stack pointer register
	switch token {
	case AX:
		stack = append(stack, registers.AX)
		registers.SP++
		break
	case BX:
		stack = append(stack, registers.BX)
		registers.SP++
		break
	default:
		//TODO:check if value is numeric
		stack = append(stack, token)
		registers.SP++

		// TODO handle parse failure i.e. non numerical value
		break
	}

}

func hltInstr(bin_code []int, registers *Registers) {

}
