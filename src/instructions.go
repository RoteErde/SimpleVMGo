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

func popInstr(binCode []int, registers *Registers) {
	registers.IP++
	var opcode = fetchInstruction(binCode[:], registers.IP)

	switch opcode {
	case AX:
		if registers.SP > 0 {
			//TODO: handle stack function
		}

		break
	case BX:
		break
	}
}

func loadInstr(binCode []int, registers *Registers) {
	registers.IP++
	var opcode1 = fetchInstruction(binCode[:], registers.IP)
	registers.IP++
	var opcode2 = fetchInstruction(binCode[:], registers.IP)
	switch opcode1 {
	case AX:
		registers.AX = opcode2
		break
	case BX:
		registers.BX = opcode2
		break
	default:
		break
	}
}

func pushInstr(binCode []int, registers *Registers) {
	//increment code pointer register
	registers.IP++
	var token = fetchInstruction(binCode[:], registers.IP)
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
