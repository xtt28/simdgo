//go:build !noavx

#include "textflag.h"

TEXT ·Vec4x64Add(SB), NOSPLIT, $0-24
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), CX
    MOVQ d+16(FP), DX
    VMOVAPD (AX), Y0
    VMOVAPD (CX), Y1
    VADDPD Y1, Y0, Y0
    VMOVAPD Y0, (DX)
    VZEROUPPER
    RET

TEXT ·Vec4x64Sub(SB), NOSPLIT, $0-24
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), CX
    MOVQ d+16(FP), DX
    VMOVAPD (AX), Y0
    VMOVAPD (CX), Y1
    VSUBPD Y1, Y0, Y0
    VMOVAPD Y0, (DX)
    VZEROUPPER
    RET

TEXT ·Vec4x64Mul(SB), NOSPLIT, $0-24
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), CX
    MOVQ d+16(FP), DX
    VMOVAPD (AX), Y0
    VMOVAPD (CX), Y1
    VMULPD Y1, Y0, Y0
    VMOVAPD Y0, (DX)
    VZEROUPPER
    RET

TEXT ·Vec4x64Div(SB), NOSPLIT, $0-24
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), CX
    MOVQ d+16(FP), DX
    VMOVAPD (AX), Y0
    VMOVAPD (CX), Y1
    VDIVPD Y1, Y0, Y0
    VMOVAPD Y0, (DX)
    VZEROUPPER
    RET
