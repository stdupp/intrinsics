// func clflushopt(p uintptr) 
TEXT ·clflushopt(SB),7,$0
	MOVQ p+0(FP),R8

	// TODO: Code missing
	// Could be:
	// CLFLUSHOPT R8

	RET

