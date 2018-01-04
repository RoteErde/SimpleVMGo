package main

type sub func([]int, *Registers)

/**
Alvin Ng : skeleton code to map instructions to functions

aim is to avoid lengthy case statements
*/
func makeInstructions() map[int]sub {
	mappedFunctions := map[int]sub{
		PSH:    pushInstr,
		POP:    popInstr,
		HLT:    hltInstr,
		POP_AX: popAXInstr,
		POP_BX: popBXInstr,
	}
	return mappedFunctions
}

func popInstr(binCode []int, registers *Registers) {
	log("Popping")
	registers.IP++
	var opcode = fetchInstruction(binCode[:], registers.IP)
	registers.IP++
	switch opcode {
	case AX:
		if registers.SP > 0 {
			//TODO: handle stack function
			registers.AX = registers.stack[len(registers.stack)-1]
			registers.stack = append(registers.stack[:0], registers.stack[1:]...) //varadic function
			//delete(registers.stack, registers.stack[len(registers.stack)-1])
			registers.SP--

		} else {
			//TODO: handle stack empty exception
		}

		break
	case BX:
		if registers.SP > 0 {
			registers.BX = registers.stack[len(registers.stack)-1]
		} else {
			//TODO: handle stck empty exception
		}

		break
	}
	log("Done Popping")
}

func popAXInstr(binCode []int, registers *Registers) {
	log("Popping")
	registers.IP++
	//var opcode = fetchInstruction(binCode[:], registers.IP)
	registers.IP++
	if registers.SP > 0 {
		//TODO: handle stack function
		registers.AX = registers.stack[len(registers.stack)-1]
		registers.stack = append(registers.stack[:0], registers.stack[1:]...) //varadic function
		//delete(registers.stack, registers.stack[len(registers.stack)-1])
		registers.SP--

	} else {
		//TODO: handle stack empty exception
	}

	log("Done Popping")
}

func popBXInstr(binCode []int, registers *Registers) {
	log("Popping")
	registers.IP++
	//var opcode = fetchInstruction(binCode[:], registers.IP)
	registers.IP++
	if registers.SP > 0 {
		//TODO: handle stack function
		registers.BX = registers.stack[len(registers.stack)-1]
		registers.stack = append(registers.stack[:0], registers.stack[1:]...) //varadic function
		//delete(registers.stack, registers.stack[len(registers.stack)-1])
		registers.SP--

	} else {
		//TODO: handle stack empty exception
	}

	log("Done Popping")
}

func loadInstr(binCode []int, registers *Registers) {
	log("Loading")
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

func _pushInstr(binCode []int, registers *Registers) {
	log("Pushing")
	//increment code pointer register
	registers.IP++
	log("Done Pushing")
}

func pushInstr(binCode []int, registers *Registers) {
	log("Pushing")
	//increment code pointer register
	registers.IP++
	var token = fetchInstruction(binCode[:], registers.IP)
	registers.IP++
	// on successful parsing, add value to the
	// stack and increment stack pointer register

	registers.stack = append(registers.stack, token)
	registers.SP++

	log("Done Pushing")

}

func hltInstr(bin_code []int, registers *Registers) {
	registers.IP = HALTMEM // magic code to signify end of instructions
}
