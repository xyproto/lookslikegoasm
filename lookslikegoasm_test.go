package lookslikegoasm

import "testing"

// TestConsider tests the Consider function to check if it correctly detects Go/Plan9 assembly.
func TestConsider(t *testing.T) {
	// Test case: Go/Plan9 assembly code
	goPlan9Source := `
	TEXT hello(SB), $0-0
	MOVQ AX, BX
	ADDQ $1, AX
	CALL somefunction
	`

	if !Consider(goPlan9Source) {
		t.Errorf("Expected Go/Plan9 Assembly to be detected, but it was not.")
	}

	// Test case: Intel assembly code
	intelSource := `
	mov eax, ebx
	add eax, 1
	jmp label
	call function
	`

	if Consider(intelSource) {
		t.Errorf("Expected Intel Assembly to be detected, but Go/Plan9 Assembly was incorrectly identified.")
	}

	// Test case: AT&T assembly code
	atntSource := `
	movl %eax, %ebx
	addl $1, %eax
	jmpl label
	call function
	`

	if Consider(atntSource) {
		t.Errorf("Expected AT&T Assembly to be detected, but Go/Plan9 Assembly was incorrectly identified.")
	}

	// Test case: Mixed code (should lean towards Go/Plan9)
	mixedSource := `
	TEXT myFunc(SB), $0-0
	mov eax, ebx
	ADDQ $1, AX
	CALL anotherfunc
	`

	if !Consider(mixedSource) {
		t.Errorf("Expected Go/Plan9 Assembly to be detected in mixed code, but it was not.")
	}

	// Test case: Edge case - Empty source code
	emptySource := ""

	if Consider(emptySource) {
		t.Errorf("Expected no detection for empty source, but Go/Plan9 Assembly was incorrectly identified.")
	}
}
