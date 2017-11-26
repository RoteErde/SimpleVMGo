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

}

func pushInstr(bin_code []int, registers *Registers) {
	registers.SP++
	var token = fetchInstruction(bin_code[:], registers.SP)
	switch token {
	case AX:
		stack = append(stack, registers.AX)
		break
	case BX:
		stack = append(stack, registers.BX)
		break
	default:
		//check if value is numeric
		stack = append(stack, token)
		break
	}

}

func hltInstr(bin_code []int, registers *Registers) {

}
