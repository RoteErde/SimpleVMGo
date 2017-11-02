package main
import "fmt"

const (
	PSH = 1
	POP = 2
	ADD = 3
	SET = 4
	MOV = 5
	HLT	= 99
)

var code_block = [...]int{PSH, 10 ,POP, HLT}

func main() {
	var code_ptr =0
	//var inst_code[100] int
	fmt.Println("Start")
	for code_ptr < len(code_block){
		fmt.Print(fetchInstruction(code_block[:], code_ptr))
		code_ptr++
	}
}

func fetchInstruction(code[] int, ptr int) int{
	return code[ptr]
}
