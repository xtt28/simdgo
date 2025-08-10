//go:build !noavx

#include "textflag.h"

TEXT ·Vec8x32Add(SB), NOSPLIT, $0-24
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), CX
    MOVQ d+16(FP), DX
    VMOVAPS (AX), Y0
    VMOVAPS (CX), Y1
    VADDPS Y1, Y0, Y0
    VMOVAPS Y0, (DX)
    VZEROUPPER
    RET

TEXT ·Vec8x32Sub(SB), NOSPLIT, $0-24
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), CX
    MOVQ d+16(FP), DX
    VMOVAPS (AX), Y0
    VMOVAPS (CX), Y1
    VSUBPS Y1, Y0, Y0
    VMOVAPS Y0, (DX)
    VZEROUPPER
    RET

TEXT ·Vec8x32Mul(SB), NOSPLIT, $0-24
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), CX
    MOVQ d+16(FP), DX
    VMOVAPS (AX), Y0
    VMOVAPS (CX), Y1
    VMULPS Y1, Y0, Y0
    VMOVAPS Y0, (DX)
    VZEROUPPER
    RET

TEXT ·Vec8x32Div(SB), NOSPLIT, $0-24
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), CX
    MOVQ d+16(FP), DX
    VMOVAPS (AX), Y0
    VMOVAPS (CX), Y1
    VDIVPS Y1, Y0, Y0
    VMOVAPS Y0, (DX)
    VZEROUPPER
    RET
