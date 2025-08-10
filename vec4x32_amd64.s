//go:build !nosse

#include "textflag.h"

TEXT ·Vec4x32Add(SB), NOSPLIT, $0-24
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), CX
    MOVQ d+16(FP), DX
    MOVAPS (AX), X0
    MOVAPS (CX), X1
    ADDPS X1, X0
    MOVAPS X0, (DX)
    RET

TEXT ·Vec4x32Sub(SB), NOSPLIT, $0-24
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), CX
    MOVQ d+16(FP), DX
    MOVAPS (AX), X0
    MOVAPS (CX), X1
    SUBPS X1, X0
    MOVAPS X0, (DX)
    RET

TEXT ·Vec4x32Mul(SB), NOSPLIT, $0-24
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), CX
    MOVQ d+16(FP), DX
    MOVAPS (AX), X0
    MOVAPS (CX), X1
    MULPS X1, X0
    MOVAPS X0, (DX)
    RET

TEXT ·Vec4x32Div(SB), NOSPLIT, $0-24
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), CX
    MOVQ d+16(FP), DX
    MOVAPS (AX), X0
    MOVAPS (CX), X1
    DIVPS X1, X0
    MOVAPS X0, (DX)
    RET
