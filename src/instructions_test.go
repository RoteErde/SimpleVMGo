package main

import (
	"testing"
)

func Test_popInstr(t *testing.T) {
	t.Run("makeInstructions::Testing POP AX instructions",
		func(t *testing.T) {
			var test_code_block = [...]int{
				PSH, 11,
				POP, AX,
				POP, BX,
				HLT,
			}
			var vmTestRegister Registers
			var mappedFunctions = makeInstructions()
			mappedFunctions[test_code_block[vmTestRegister.IP]](test_code_block[:], &vmTestRegister)
			mappedFunctions[test_code_block[vmTestRegister.IP]](test_code_block[:], &vmTestRegister)
			dumpRegister(vmTestRegister)
			t.Log("Check if AX has value of 11")
			if vmTestRegister.AX == 11 {
				testLogSuccess(t)
			} else {
				t.Errorf("AX wrong value, expect 11, got %d", vmTestRegister.AX)
			}
		})
}

func Test_loadInstr(t *testing.T) {
	type args struct {
		binCode   []int
		registers *Registers
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loadInstr(tt.args.binCode, tt.args.registers)
		})
	}
}

func Test_pushInstr(t *testing.T) {

	t.Run("makeInstructions::Testing PSH instructions",
		func(t *testing.T) {
			var test_code_block = [...]int{
				PSH, 11,
				PSH, 12,
				POP, AX,
				POP, BX,
				HLT,
			}
			var vmTestRegister Registers
			var mappedFunctions = makeInstructions()
			var testStack []int
			var initialStackCount = len(vmTestRegister.stack)

			mappedFunctions[test_code_block[vmTestRegister.IP]](test_code_block[:], &vmTestRegister)
			t.Log("Testing IP")
			if vmTestRegister.IP == 2 {
				testLogSuccess(t)
			} else {
				t.Errorf("Failed, IP incorrect value, Expected vmTestRegister.IP = 2, got %d",
					vmTestRegister.IP)
			}
			t.Log("Testing Stack")
			var endStackCount = len(vmTestRegister.stack)
			if endStackCount-initialStackCount == 1 {
				testLogSuccess(t)
			} else {
				t.Errorf("Failed, stack incorrect length, Expected stack == 2, got %d", len(testStack))
			}
			//t.Errorf("PSH not found")
		})

}

func Test_hltInstr(t *testing.T) {
	type args struct {
		bin_code  []int
		registers *Registers
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hltInstr(tt.args.bin_code, tt.args.registers)
		})
	}
}
