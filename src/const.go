package main

const (
	PSH    = 101
	POP    = 102
	ADD    = 103
	SET    = 104
	LOAD   = 105
	HLT    = 99
	AX     = 0
	BX     = 1
	POP_AX = 120
	POP_BX = 121

	HALTMEM = 65535 // magic number to signify instruction has ended
)
