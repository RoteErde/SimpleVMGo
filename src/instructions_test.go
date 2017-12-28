package main

import (
	"testing"
)

func Test_makeInstructions(t *testing.T) {

	t.Run("makeInstructions::Testing PSH instructions",
		func(t *testing.T) {
			var test_code_block = [...]int{PSH, 10,
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
				t.Log("✓ Success")
			} else {
				t.Errorf("Failed, IP incorrect value, Expected vmTestRegister.IP = 2, got %d",
					vmTestRegister.IP)
			}
			t.Log("Testing Stack")
			var endStackCount = len(vmTestRegister.stack)
			if endStackCount-initialStackCount == 1 {
				t.Log("✓ Success")
			} else {
				t.Errorf("Failed, stack incorrect length, Expected stack == 2, got %d", len(testStack))
			}
			//t.Errorf("PSH not found")
		})
}

/*func Test_makeInstructions(t *testing.T) {
	tests := []struct {
		name string
		want map[int]sub
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeInstructions(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeInstructions() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
