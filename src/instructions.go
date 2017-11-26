package main

type sub func([]int, *Registers)

func makeInstructions() {
	mappedFunctions := map[int]sub{
		POP: popInstr,
	}

}

func popInstr(bin_code []int, registers *Registers) {

}
