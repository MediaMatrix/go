// Code generated from gen/MIPS.rules; DO NOT EDIT.
// generated with: cd gen; go run *.go

package ssa

import "cmd/compile/internal/types"

func rewriteValueMIPS(v *Value) bool {
	switch v.Op {
	case OpAdd16:
		v.Op = OpMIPSADD
		return true
	case OpAdd32:
		v.Op = OpMIPSADD
		return true
	case OpAdd32F:
		v.Op = OpMIPSADDF
		return true
	case OpAdd32withcarry:
		return rewriteValueMIPS_OpAdd32withcarry(v)
	case OpAdd64F:
		v.Op = OpMIPSADDD
		return true
	case OpAdd8:
		v.Op = OpMIPSADD
		return true
	case OpAddPtr:
		v.Op = OpMIPSADD
		return true
	case OpAddr:
		v.Op = OpMIPSMOVWaddr
		return true
	case OpAnd16:
		v.Op = OpMIPSAND
		return true
	case OpAnd32:
		v.Op = OpMIPSAND
		return true
	case OpAnd8:
		v.Op = OpMIPSAND
		return true
	case OpAndB:
		v.Op = OpMIPSAND
		return true
	case OpAtomicAdd32:
		v.Op = OpMIPSLoweredAtomicAdd
		return true
	case OpAtomicAnd8:
		return rewriteValueMIPS_OpAtomicAnd8(v)
	case OpAtomicCompareAndSwap32:
		v.Op = OpMIPSLoweredAtomicCas
		return true
	case OpAtomicExchange32:
		v.Op = OpMIPSLoweredAtomicExchange
		return true
	case OpAtomicLoad32:
		v.Op = OpMIPSLoweredAtomicLoad32
		return true
	case OpAtomicLoad8:
		v.Op = OpMIPSLoweredAtomicLoad8
		return true
	case OpAtomicLoadPtr:
		v.Op = OpMIPSLoweredAtomicLoad32
		return true
	case OpAtomicOr8:
		return rewriteValueMIPS_OpAtomicOr8(v)
	case OpAtomicStore32:
		v.Op = OpMIPSLoweredAtomicStore32
		return true
	case OpAtomicStore8:
		v.Op = OpMIPSLoweredAtomicStore8
		return true
	case OpAtomicStorePtrNoWB:
		v.Op = OpMIPSLoweredAtomicStore32
		return true
	case OpAvg32u:
		return rewriteValueMIPS_OpAvg32u(v)
	case OpBitLen32:
		return rewriteValueMIPS_OpBitLen32(v)
	case OpClosureCall:
		v.Op = OpMIPSCALLclosure
		return true
	case OpCom16:
		return rewriteValueMIPS_OpCom16(v)
	case OpCom32:
		return rewriteValueMIPS_OpCom32(v)
	case OpCom8:
		return rewriteValueMIPS_OpCom8(v)
	case OpConst16:
		v.Op = OpMIPSMOVWconst
		return true
	case OpConst32:
		v.Op = OpMIPSMOVWconst
		return true
	case OpConst32F:
		v.Op = OpMIPSMOVFconst
		return true
	case OpConst64F:
		v.Op = OpMIPSMOVDconst
		return true
	case OpConst8:
		v.Op = OpMIPSMOVWconst
		return true
	case OpConstBool:
		v.Op = OpMIPSMOVWconst
		return true
	case OpConstNil:
		return rewriteValueMIPS_OpConstNil(v)
	case OpCtz32:
		return rewriteValueMIPS_OpCtz32(v)
	case OpCtz32NonZero:
		v.Op = OpCtz32
		return true
	case OpCvt32Fto32:
		v.Op = OpMIPSTRUNCFW
		return true
	case OpCvt32Fto64F:
		v.Op = OpMIPSMOVFD
		return true
	case OpCvt32to32F:
		v.Op = OpMIPSMOVWF
		return true
	case OpCvt32to64F:
		v.Op = OpMIPSMOVWD
		return true
	case OpCvt64Fto32:
		v.Op = OpMIPSTRUNCDW
		return true
	case OpCvt64Fto32F:
		v.Op = OpMIPSMOVDF
		return true
	case OpCvtBoolToUint8:
		v.Op = OpCopy
		return true
	case OpDiv16:
		return rewriteValueMIPS_OpDiv16(v)
	case OpDiv16u:
		return rewriteValueMIPS_OpDiv16u(v)
	case OpDiv32:
		return rewriteValueMIPS_OpDiv32(v)
	case OpDiv32F:
		v.Op = OpMIPSDIVF
		return true
	case OpDiv32u:
		return rewriteValueMIPS_OpDiv32u(v)
	case OpDiv64F:
		v.Op = OpMIPSDIVD
		return true
	case OpDiv8:
		return rewriteValueMIPS_OpDiv8(v)
	case OpDiv8u:
		return rewriteValueMIPS_OpDiv8u(v)
	case OpEq16:
		return rewriteValueMIPS_OpEq16(v)
	case OpEq32:
		return rewriteValueMIPS_OpEq32(v)
	case OpEq32F:
		return rewriteValueMIPS_OpEq32F(v)
	case OpEq64F:
		return rewriteValueMIPS_OpEq64F(v)
	case OpEq8:
		return rewriteValueMIPS_OpEq8(v)
	case OpEqB:
		return rewriteValueMIPS_OpEqB(v)
	case OpEqPtr:
		return rewriteValueMIPS_OpEqPtr(v)
	case OpGeq32F:
		return rewriteValueMIPS_OpGeq32F(v)
	case OpGeq64F:
		return rewriteValueMIPS_OpGeq64F(v)
	case OpGetCallerPC:
		v.Op = OpMIPSLoweredGetCallerPC
		return true
	case OpGetCallerSP:
		v.Op = OpMIPSLoweredGetCallerSP
		return true
	case OpGetClosurePtr:
		v.Op = OpMIPSLoweredGetClosurePtr
		return true
	case OpGreater32F:
		return rewriteValueMIPS_OpGreater32F(v)
	case OpGreater64F:
		return rewriteValueMIPS_OpGreater64F(v)
	case OpHmul32:
		return rewriteValueMIPS_OpHmul32(v)
	case OpHmul32u:
		return rewriteValueMIPS_OpHmul32u(v)
	case OpInterCall:
		v.Op = OpMIPSCALLinter
		return true
	case OpIsInBounds:
		return rewriteValueMIPS_OpIsInBounds(v)
	case OpIsNonNil:
		return rewriteValueMIPS_OpIsNonNil(v)
	case OpIsSliceInBounds:
		return rewriteValueMIPS_OpIsSliceInBounds(v)
	case OpLeq16:
		return rewriteValueMIPS_OpLeq16(v)
	case OpLeq16U:
		return rewriteValueMIPS_OpLeq16U(v)
	case OpLeq32:
		return rewriteValueMIPS_OpLeq32(v)
	case OpLeq32F:
		return rewriteValueMIPS_OpLeq32F(v)
	case OpLeq32U:
		return rewriteValueMIPS_OpLeq32U(v)
	case OpLeq64F:
		return rewriteValueMIPS_OpLeq64F(v)
	case OpLeq8:
		return rewriteValueMIPS_OpLeq8(v)
	case OpLeq8U:
		return rewriteValueMIPS_OpLeq8U(v)
	case OpLess16:
		return rewriteValueMIPS_OpLess16(v)
	case OpLess16U:
		return rewriteValueMIPS_OpLess16U(v)
	case OpLess32:
		return rewriteValueMIPS_OpLess32(v)
	case OpLess32F:
		return rewriteValueMIPS_OpLess32F(v)
	case OpLess32U:
		return rewriteValueMIPS_OpLess32U(v)
	case OpLess64F:
		return rewriteValueMIPS_OpLess64F(v)
	case OpLess8:
		return rewriteValueMIPS_OpLess8(v)
	case OpLess8U:
		return rewriteValueMIPS_OpLess8U(v)
	case OpLoad:
		return rewriteValueMIPS_OpLoad(v)
	case OpLocalAddr:
		return rewriteValueMIPS_OpLocalAddr(v)
	case OpLsh16x16:
		return rewriteValueMIPS_OpLsh16x16(v)
	case OpLsh16x32:
		return rewriteValueMIPS_OpLsh16x32(v)
	case OpLsh16x64:
		return rewriteValueMIPS_OpLsh16x64(v)
	case OpLsh16x8:
		return rewriteValueMIPS_OpLsh16x8(v)
	case OpLsh32x16:
		return rewriteValueMIPS_OpLsh32x16(v)
	case OpLsh32x32:
		return rewriteValueMIPS_OpLsh32x32(v)
	case OpLsh32x64:
		return rewriteValueMIPS_OpLsh32x64(v)
	case OpLsh32x8:
		return rewriteValueMIPS_OpLsh32x8(v)
	case OpLsh8x16:
		return rewriteValueMIPS_OpLsh8x16(v)
	case OpLsh8x32:
		return rewriteValueMIPS_OpLsh8x32(v)
	case OpLsh8x64:
		return rewriteValueMIPS_OpLsh8x64(v)
	case OpLsh8x8:
		return rewriteValueMIPS_OpLsh8x8(v)
	case OpMIPSADD:
		return rewriteValueMIPS_OpMIPSADD(v)
	case OpMIPSADDconst:
		return rewriteValueMIPS_OpMIPSADDconst(v)
	case OpMIPSAND:
		return rewriteValueMIPS_OpMIPSAND(v)
	case OpMIPSANDconst:
		return rewriteValueMIPS_OpMIPSANDconst(v)
	case OpMIPSCMOVZ:
		return rewriteValueMIPS_OpMIPSCMOVZ(v)
	case OpMIPSCMOVZzero:
		return rewriteValueMIPS_OpMIPSCMOVZzero(v)
	case OpMIPSLoweredAtomicAdd:
		return rewriteValueMIPS_OpMIPSLoweredAtomicAdd(v)
	case OpMIPSLoweredAtomicStore32:
		return rewriteValueMIPS_OpMIPSLoweredAtomicStore32(v)
	case OpMIPSMOVBUload:
		return rewriteValueMIPS_OpMIPSMOVBUload(v)
	case OpMIPSMOVBUreg:
		return rewriteValueMIPS_OpMIPSMOVBUreg(v)
	case OpMIPSMOVBload:
		return rewriteValueMIPS_OpMIPSMOVBload(v)
	case OpMIPSMOVBreg:
		return rewriteValueMIPS_OpMIPSMOVBreg(v)
	case OpMIPSMOVBstore:
		return rewriteValueMIPS_OpMIPSMOVBstore(v)
	case OpMIPSMOVBstorezero:
		return rewriteValueMIPS_OpMIPSMOVBstorezero(v)
	case OpMIPSMOVDload:
		return rewriteValueMIPS_OpMIPSMOVDload(v)
	case OpMIPSMOVDstore:
		return rewriteValueMIPS_OpMIPSMOVDstore(v)
	case OpMIPSMOVFload:
		return rewriteValueMIPS_OpMIPSMOVFload(v)
	case OpMIPSMOVFstore:
		return rewriteValueMIPS_OpMIPSMOVFstore(v)
	case OpMIPSMOVHUload:
		return rewriteValueMIPS_OpMIPSMOVHUload(v)
	case OpMIPSMOVHUreg:
		return rewriteValueMIPS_OpMIPSMOVHUreg(v)
	case OpMIPSMOVHload:
		return rewriteValueMIPS_OpMIPSMOVHload(v)
	case OpMIPSMOVHreg:
		return rewriteValueMIPS_OpMIPSMOVHreg(v)
	case OpMIPSMOVHstore:
		return rewriteValueMIPS_OpMIPSMOVHstore(v)
	case OpMIPSMOVHstorezero:
		return rewriteValueMIPS_OpMIPSMOVHstorezero(v)
	case OpMIPSMOVWload:
		return rewriteValueMIPS_OpMIPSMOVWload(v)
	case OpMIPSMOVWreg:
		return rewriteValueMIPS_OpMIPSMOVWreg(v)
	case OpMIPSMOVWstore:
		return rewriteValueMIPS_OpMIPSMOVWstore(v)
	case OpMIPSMOVWstorezero:
		return rewriteValueMIPS_OpMIPSMOVWstorezero(v)
	case OpMIPSMUL:
		return rewriteValueMIPS_OpMIPSMUL(v)
	case OpMIPSNEG:
		return rewriteValueMIPS_OpMIPSNEG(v)
	case OpMIPSNOR:
		return rewriteValueMIPS_OpMIPSNOR(v)
	case OpMIPSNORconst:
		return rewriteValueMIPS_OpMIPSNORconst(v)
	case OpMIPSOR:
		return rewriteValueMIPS_OpMIPSOR(v)
	case OpMIPSORconst:
		return rewriteValueMIPS_OpMIPSORconst(v)
	case OpMIPSSGT:
		return rewriteValueMIPS_OpMIPSSGT(v)
	case OpMIPSSGTU:
		return rewriteValueMIPS_OpMIPSSGTU(v)
	case OpMIPSSGTUconst:
		return rewriteValueMIPS_OpMIPSSGTUconst(v)
	case OpMIPSSGTUzero:
		return rewriteValueMIPS_OpMIPSSGTUzero(v)
	case OpMIPSSGTconst:
		return rewriteValueMIPS_OpMIPSSGTconst(v)
	case OpMIPSSGTzero:
		return rewriteValueMIPS_OpMIPSSGTzero(v)
	case OpMIPSSLL:
		return rewriteValueMIPS_OpMIPSSLL(v)
	case OpMIPSSLLconst:
		return rewriteValueMIPS_OpMIPSSLLconst(v)
	case OpMIPSSRA:
		return rewriteValueMIPS_OpMIPSSRA(v)
	case OpMIPSSRAconst:
		return rewriteValueMIPS_OpMIPSSRAconst(v)
	case OpMIPSSRL:
		return rewriteValueMIPS_OpMIPSSRL(v)
	case OpMIPSSRLconst:
		return rewriteValueMIPS_OpMIPSSRLconst(v)
	case OpMIPSSUB:
		return rewriteValueMIPS_OpMIPSSUB(v)
	case OpMIPSSUBconst:
		return rewriteValueMIPS_OpMIPSSUBconst(v)
	case OpMIPSXOR:
		return rewriteValueMIPS_OpMIPSXOR(v)
	case OpMIPSXORconst:
		return rewriteValueMIPS_OpMIPSXORconst(v)
	case OpMod16:
		return rewriteValueMIPS_OpMod16(v)
	case OpMod16u:
		return rewriteValueMIPS_OpMod16u(v)
	case OpMod32:
		return rewriteValueMIPS_OpMod32(v)
	case OpMod32u:
		return rewriteValueMIPS_OpMod32u(v)
	case OpMod8:
		return rewriteValueMIPS_OpMod8(v)
	case OpMod8u:
		return rewriteValueMIPS_OpMod8u(v)
	case OpMove:
		return rewriteValueMIPS_OpMove(v)
	case OpMul16:
		v.Op = OpMIPSMUL
		return true
	case OpMul32:
		v.Op = OpMIPSMUL
		return true
	case OpMul32F:
		v.Op = OpMIPSMULF
		return true
	case OpMul32uhilo:
		v.Op = OpMIPSMULTU
		return true
	case OpMul64F:
		v.Op = OpMIPSMULD
		return true
	case OpMul8:
		v.Op = OpMIPSMUL
		return true
	case OpNeg16:
		v.Op = OpMIPSNEG
		return true
	case OpNeg32:
		v.Op = OpMIPSNEG
		return true
	case OpNeg32F:
		v.Op = OpMIPSNEGF
		return true
	case OpNeg64F:
		v.Op = OpMIPSNEGD
		return true
	case OpNeg8:
		v.Op = OpMIPSNEG
		return true
	case OpNeq16:
		return rewriteValueMIPS_OpNeq16(v)
	case OpNeq32:
		return rewriteValueMIPS_OpNeq32(v)
	case OpNeq32F:
		return rewriteValueMIPS_OpNeq32F(v)
	case OpNeq64F:
		return rewriteValueMIPS_OpNeq64F(v)
	case OpNeq8:
		return rewriteValueMIPS_OpNeq8(v)
	case OpNeqB:
		v.Op = OpMIPSXOR
		return true
	case OpNeqPtr:
		return rewriteValueMIPS_OpNeqPtr(v)
	case OpNilCheck:
		v.Op = OpMIPSLoweredNilCheck
		return true
	case OpNot:
		return rewriteValueMIPS_OpNot(v)
	case OpOffPtr:
		return rewriteValueMIPS_OpOffPtr(v)
	case OpOr16:
		v.Op = OpMIPSOR
		return true
	case OpOr32:
		v.Op = OpMIPSOR
		return true
	case OpOr8:
		v.Op = OpMIPSOR
		return true
	case OpOrB:
		v.Op = OpMIPSOR
		return true
	case OpPanicBounds:
		return rewriteValueMIPS_OpPanicBounds(v)
	case OpPanicExtend:
		return rewriteValueMIPS_OpPanicExtend(v)
	case OpRotateLeft16:
		return rewriteValueMIPS_OpRotateLeft16(v)
	case OpRotateLeft32:
		return rewriteValueMIPS_OpRotateLeft32(v)
	case OpRotateLeft64:
		return rewriteValueMIPS_OpRotateLeft64(v)
	case OpRotateLeft8:
		return rewriteValueMIPS_OpRotateLeft8(v)
	case OpRound32F:
		v.Op = OpCopy
		return true
	case OpRound64F:
		v.Op = OpCopy
		return true
	case OpRsh16Ux16:
		return rewriteValueMIPS_OpRsh16Ux16(v)
	case OpRsh16Ux32:
		return rewriteValueMIPS_OpRsh16Ux32(v)
	case OpRsh16Ux64:
		return rewriteValueMIPS_OpRsh16Ux64(v)
	case OpRsh16Ux8:
		return rewriteValueMIPS_OpRsh16Ux8(v)
	case OpRsh16x16:
		return rewriteValueMIPS_OpRsh16x16(v)
	case OpRsh16x32:
		return rewriteValueMIPS_OpRsh16x32(v)
	case OpRsh16x64:
		return rewriteValueMIPS_OpRsh16x64(v)
	case OpRsh16x8:
		return rewriteValueMIPS_OpRsh16x8(v)
	case OpRsh32Ux16:
		return rewriteValueMIPS_OpRsh32Ux16(v)
	case OpRsh32Ux32:
		return rewriteValueMIPS_OpRsh32Ux32(v)
	case OpRsh32Ux64:
		return rewriteValueMIPS_OpRsh32Ux64(v)
	case OpRsh32Ux8:
		return rewriteValueMIPS_OpRsh32Ux8(v)
	case OpRsh32x16:
		return rewriteValueMIPS_OpRsh32x16(v)
	case OpRsh32x32:
		return rewriteValueMIPS_OpRsh32x32(v)
	case OpRsh32x64:
		return rewriteValueMIPS_OpRsh32x64(v)
	case OpRsh32x8:
		return rewriteValueMIPS_OpRsh32x8(v)
	case OpRsh8Ux16:
		return rewriteValueMIPS_OpRsh8Ux16(v)
	case OpRsh8Ux32:
		return rewriteValueMIPS_OpRsh8Ux32(v)
	case OpRsh8Ux64:
		return rewriteValueMIPS_OpRsh8Ux64(v)
	case OpRsh8Ux8:
		return rewriteValueMIPS_OpRsh8Ux8(v)
	case OpRsh8x16:
		return rewriteValueMIPS_OpRsh8x16(v)
	case OpRsh8x32:
		return rewriteValueMIPS_OpRsh8x32(v)
	case OpRsh8x64:
		return rewriteValueMIPS_OpRsh8x64(v)
	case OpRsh8x8:
		return rewriteValueMIPS_OpRsh8x8(v)
	case OpSelect0:
		return rewriteValueMIPS_OpSelect0(v)
	case OpSelect1:
		return rewriteValueMIPS_OpSelect1(v)
	case OpSignExt16to32:
		v.Op = OpMIPSMOVHreg
		return true
	case OpSignExt8to16:
		v.Op = OpMIPSMOVBreg
		return true
	case OpSignExt8to32:
		v.Op = OpMIPSMOVBreg
		return true
	case OpSignmask:
		return rewriteValueMIPS_OpSignmask(v)
	case OpSlicemask:
		return rewriteValueMIPS_OpSlicemask(v)
	case OpSqrt:
		v.Op = OpMIPSSQRTD
		return true
	case OpStaticCall:
		v.Op = OpMIPSCALLstatic
		return true
	case OpStore:
		return rewriteValueMIPS_OpStore(v)
	case OpSub16:
		v.Op = OpMIPSSUB
		return true
	case OpSub32:
		v.Op = OpMIPSSUB
		return true
	case OpSub32F:
		v.Op = OpMIPSSUBF
		return true
	case OpSub32withcarry:
		return rewriteValueMIPS_OpSub32withcarry(v)
	case OpSub64F:
		v.Op = OpMIPSSUBD
		return true
	case OpSub8:
		v.Op = OpMIPSSUB
		return true
	case OpSubPtr:
		v.Op = OpMIPSSUB
		return true
	case OpTrunc16to8:
		v.Op = OpCopy
		return true
	case OpTrunc32to16:
		v.Op = OpCopy
		return true
	case OpTrunc32to8:
		v.Op = OpCopy
		return true
	case OpWB:
		v.Op = OpMIPSLoweredWB
		return true
	case OpXor16:
		v.Op = OpMIPSXOR
		return true
	case OpXor32:
		v.Op = OpMIPSXOR
		return true
	case OpXor8:
		v.Op = OpMIPSXOR
		return true
	case OpZero:
		return rewriteValueMIPS_OpZero(v)
	case OpZeroExt16to32:
		v.Op = OpMIPSMOVHUreg
		return true
	case OpZeroExt8to16:
		v.Op = OpMIPSMOVBUreg
		return true
	case OpZeroExt8to32:
		v.Op = OpMIPSMOVBUreg
		return true
	case OpZeromask:
		return rewriteValueMIPS_OpZeromask(v)
	}
	return false
}
func rewriteValueMIPS_OpAdd32withcarry(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Add32withcarry <t> x y c)
	// result: (ADD c (ADD <t> x y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		c := v_2
		v.reset(OpMIPSADD)
		v0 := b.NewValue0(v.Pos, OpMIPSADD, t)
		v0.AddArg2(x, y)
		v.AddArg2(c, v0)
		return true
	}
}
func rewriteValueMIPS_OpAtomicAnd8(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (AtomicAnd8 ptr val mem)
	// cond: !config.BigEndian
	// result: (LoweredAtomicAnd (AND <typ.UInt32Ptr> (MOVWconst [^3]) ptr) (OR <typ.UInt32> (SLL <typ.UInt32> (ZeroExt8to32 val) (SLLconst <typ.UInt32> [3] (ANDconst <typ.UInt32> [3] ptr))) (NORconst [0] <typ.UInt32> (SLL <typ.UInt32> (MOVWconst [0xff]) (SLLconst <typ.UInt32> [3] (ANDconst <typ.UInt32> [3] ptr))))) mem)
	for {
		ptr := v_0
		val := v_1
		mem := v_2
		if !(!config.BigEndian) {
			break
		}
		v.reset(OpMIPSLoweredAtomicAnd)
		v0 := b.NewValue0(v.Pos, OpMIPSAND, typ.UInt32Ptr)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = ^3
		v0.AddArg2(v1, ptr)
		v2 := b.NewValue0(v.Pos, OpMIPSOR, typ.UInt32)
		v3 := b.NewValue0(v.Pos, OpMIPSSLL, typ.UInt32)
		v4 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v4.AddArg(val)
		v5 := b.NewValue0(v.Pos, OpMIPSSLLconst, typ.UInt32)
		v5.AuxInt = 3
		v6 := b.NewValue0(v.Pos, OpMIPSANDconst, typ.UInt32)
		v6.AuxInt = 3
		v6.AddArg(ptr)
		v5.AddArg(v6)
		v3.AddArg2(v4, v5)
		v7 := b.NewValue0(v.Pos, OpMIPSNORconst, typ.UInt32)
		v7.AuxInt = 0
		v8 := b.NewValue0(v.Pos, OpMIPSSLL, typ.UInt32)
		v9 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v9.AuxInt = 0xff
		v10 := b.NewValue0(v.Pos, OpMIPSSLLconst, typ.UInt32)
		v10.AuxInt = 3
		v11 := b.NewValue0(v.Pos, OpMIPSANDconst, typ.UInt32)
		v11.AuxInt = 3
		v11.AddArg(ptr)
		v10.AddArg(v11)
		v8.AddArg2(v9, v10)
		v7.AddArg(v8)
		v2.AddArg2(v3, v7)
		v.AddArg3(v0, v2, mem)
		return true
	}
	// match: (AtomicAnd8 ptr val mem)
	// cond: config.BigEndian
	// result: (LoweredAtomicAnd (AND <typ.UInt32Ptr> (MOVWconst [^3]) ptr) (OR <typ.UInt32> (SLL <typ.UInt32> (ZeroExt8to32 val) (SLLconst <typ.UInt32> [3] (ANDconst <typ.UInt32> [3] (XORconst <typ.UInt32> [3] ptr)))) (NORconst [0] <typ.UInt32> (SLL <typ.UInt32> (MOVWconst [0xff]) (SLLconst <typ.UInt32> [3] (ANDconst <typ.UInt32> [3] (XORconst <typ.UInt32> [3] ptr)))))) mem)
	for {
		ptr := v_0
		val := v_1
		mem := v_2
		if !(config.BigEndian) {
			break
		}
		v.reset(OpMIPSLoweredAtomicAnd)
		v0 := b.NewValue0(v.Pos, OpMIPSAND, typ.UInt32Ptr)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = ^3
		v0.AddArg2(v1, ptr)
		v2 := b.NewValue0(v.Pos, OpMIPSOR, typ.UInt32)
		v3 := b.NewValue0(v.Pos, OpMIPSSLL, typ.UInt32)
		v4 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v4.AddArg(val)
		v5 := b.NewValue0(v.Pos, OpMIPSSLLconst, typ.UInt32)
		v5.AuxInt = 3
		v6 := b.NewValue0(v.Pos, OpMIPSANDconst, typ.UInt32)
		v6.AuxInt = 3
		v7 := b.NewValue0(v.Pos, OpMIPSXORconst, typ.UInt32)
		v7.AuxInt = 3
		v7.AddArg(ptr)
		v6.AddArg(v7)
		v5.AddArg(v6)
		v3.AddArg2(v4, v5)
		v8 := b.NewValue0(v.Pos, OpMIPSNORconst, typ.UInt32)
		v8.AuxInt = 0
		v9 := b.NewValue0(v.Pos, OpMIPSSLL, typ.UInt32)
		v10 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v10.AuxInt = 0xff
		v11 := b.NewValue0(v.Pos, OpMIPSSLLconst, typ.UInt32)
		v11.AuxInt = 3
		v12 := b.NewValue0(v.Pos, OpMIPSANDconst, typ.UInt32)
		v12.AuxInt = 3
		v13 := b.NewValue0(v.Pos, OpMIPSXORconst, typ.UInt32)
		v13.AuxInt = 3
		v13.AddArg(ptr)
		v12.AddArg(v13)
		v11.AddArg(v12)
		v9.AddArg2(v10, v11)
		v8.AddArg(v9)
		v2.AddArg2(v3, v8)
		v.AddArg3(v0, v2, mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpAtomicOr8(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (AtomicOr8 ptr val mem)
	// cond: !config.BigEndian
	// result: (LoweredAtomicOr (AND <typ.UInt32Ptr> (MOVWconst [^3]) ptr) (SLL <typ.UInt32> (ZeroExt8to32 val) (SLLconst <typ.UInt32> [3] (ANDconst <typ.UInt32> [3] ptr))) mem)
	for {
		ptr := v_0
		val := v_1
		mem := v_2
		if !(!config.BigEndian) {
			break
		}
		v.reset(OpMIPSLoweredAtomicOr)
		v0 := b.NewValue0(v.Pos, OpMIPSAND, typ.UInt32Ptr)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = ^3
		v0.AddArg2(v1, ptr)
		v2 := b.NewValue0(v.Pos, OpMIPSSLL, typ.UInt32)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v3.AddArg(val)
		v4 := b.NewValue0(v.Pos, OpMIPSSLLconst, typ.UInt32)
		v4.AuxInt = 3
		v5 := b.NewValue0(v.Pos, OpMIPSANDconst, typ.UInt32)
		v5.AuxInt = 3
		v5.AddArg(ptr)
		v4.AddArg(v5)
		v2.AddArg2(v3, v4)
		v.AddArg3(v0, v2, mem)
		return true
	}
	// match: (AtomicOr8 ptr val mem)
	// cond: config.BigEndian
	// result: (LoweredAtomicOr (AND <typ.UInt32Ptr> (MOVWconst [^3]) ptr) (SLL <typ.UInt32> (ZeroExt8to32 val) (SLLconst <typ.UInt32> [3] (ANDconst <typ.UInt32> [3] (XORconst <typ.UInt32> [3] ptr)))) mem)
	for {
		ptr := v_0
		val := v_1
		mem := v_2
		if !(config.BigEndian) {
			break
		}
		v.reset(OpMIPSLoweredAtomicOr)
		v0 := b.NewValue0(v.Pos, OpMIPSAND, typ.UInt32Ptr)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = ^3
		v0.AddArg2(v1, ptr)
		v2 := b.NewValue0(v.Pos, OpMIPSSLL, typ.UInt32)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v3.AddArg(val)
		v4 := b.NewValue0(v.Pos, OpMIPSSLLconst, typ.UInt32)
		v4.AuxInt = 3
		v5 := b.NewValue0(v.Pos, OpMIPSANDconst, typ.UInt32)
		v5.AuxInt = 3
		v6 := b.NewValue0(v.Pos, OpMIPSXORconst, typ.UInt32)
		v6.AuxInt = 3
		v6.AddArg(ptr)
		v5.AddArg(v6)
		v4.AddArg(v5)
		v2.AddArg2(v3, v4)
		v.AddArg3(v0, v2, mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpAvg32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Avg32u <t> x y)
	// result: (ADD (SRLconst <t> (SUB <t> x y) [1]) y)
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPSADD)
		v0 := b.NewValue0(v.Pos, OpMIPSSRLconst, t)
		v0.AuxInt = 1
		v1 := b.NewValue0(v.Pos, OpMIPSSUB, t)
		v1.AddArg2(x, y)
		v0.AddArg(v1)
		v.AddArg2(v0, y)
		return true
	}
}
func rewriteValueMIPS_OpBitLen32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (BitLen32 <t> x)
	// result: (SUB (MOVWconst [32]) (CLZ <t> x))
	for {
		t := v.Type
		x := v_0
		v.reset(OpMIPSSUB)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v0.AuxInt = 32
		v1 := b.NewValue0(v.Pos, OpMIPSCLZ, t)
		v1.AddArg(x)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueMIPS_OpCom16(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Com16 x)
	// result: (NORconst [0] x)
	for {
		x := v_0
		v.reset(OpMIPSNORconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpCom32(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Com32 x)
	// result: (NORconst [0] x)
	for {
		x := v_0
		v.reset(OpMIPSNORconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpCom8(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Com8 x)
	// result: (NORconst [0] x)
	for {
		x := v_0
		v.reset(OpMIPSNORconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpConstNil(v *Value) bool {
	// match: (ConstNil)
	// result: (MOVWconst [0])
	for {
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
}
func rewriteValueMIPS_OpCtz32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Ctz32 <t> x)
	// result: (SUB (MOVWconst [32]) (CLZ <t> (SUBconst <t> [1] (AND <t> x (NEG <t> x)))))
	for {
		t := v.Type
		x := v_0
		v.reset(OpMIPSSUB)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v0.AuxInt = 32
		v1 := b.NewValue0(v.Pos, OpMIPSCLZ, t)
		v2 := b.NewValue0(v.Pos, OpMIPSSUBconst, t)
		v2.AuxInt = 1
		v3 := b.NewValue0(v.Pos, OpMIPSAND, t)
		v4 := b.NewValue0(v.Pos, OpMIPSNEG, t)
		v4.AddArg(x)
		v3.AddArg2(x, v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueMIPS_OpDiv16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16 x y)
	// result: (Select1 (DIV (SignExt16to32 x) (SignExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPSDIV, types.NewTuple(typ.Int32, typ.Int32))
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpDiv16u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16u x y)
	// result: (Select1 (DIVU (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPSDIVU, types.NewTuple(typ.UInt32, typ.UInt32))
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpDiv32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div32 x y)
	// result: (Select1 (DIV x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPSDIV, types.NewTuple(typ.Int32, typ.Int32))
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpDiv32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div32u x y)
	// result: (Select1 (DIVU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPSDIVU, types.NewTuple(typ.UInt32, typ.UInt32))
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpDiv8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8 x y)
	// result: (Select1 (DIV (SignExt8to32 x) (SignExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPSDIV, types.NewTuple(typ.Int32, typ.Int32))
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpDiv8u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8u x y)
	// result: (Select1 (DIVU (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPSDIVU, types.NewTuple(typ.UInt32, typ.UInt32))
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpEq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq16 x y)
	// result: (SGTUconst [1] (XOR (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSSGTUconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSXOR, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpEq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq32 x y)
	// result: (SGTUconst [1] (XOR x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSSGTUconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSXOR, typ.UInt32)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpEq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Eq32F x y)
	// result: (FPFlagTrue (CMPEQF x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPEQF, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpEq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Eq64F x y)
	// result: (FPFlagTrue (CMPEQD x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPEQD, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpEq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq8 x y)
	// result: (SGTUconst [1] (XOR (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSSGTUconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSXOR, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpEqB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (EqB x y)
	// result: (XORconst [1] (XOR <typ.Bool> x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSXOR, typ.Bool)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpEqPtr(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (EqPtr x y)
	// result: (SGTUconst [1] (XOR x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSSGTUconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSXOR, typ.UInt32)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpGeq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Geq32F x y)
	// result: (FPFlagTrue (CMPGEF x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPGEF, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpGeq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Geq64F x y)
	// result: (FPFlagTrue (CMPGED x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPGED, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpGreater32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Greater32F x y)
	// result: (FPFlagTrue (CMPGTF x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPGTF, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpGreater64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Greater64F x y)
	// result: (FPFlagTrue (CMPGTD x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPGTD, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpHmul32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Hmul32 x y)
	// result: (Select0 (MULT x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPSMULT, types.NewTuple(typ.Int32, typ.Int32))
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpHmul32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Hmul32u x y)
	// result: (Select0 (MULTU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPSMULTU, types.NewTuple(typ.UInt32, typ.UInt32))
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpIsInBounds(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (IsInBounds idx len)
	// result: (SGTU len idx)
	for {
		idx := v_0
		len := v_1
		v.reset(OpMIPSSGTU)
		v.AddArg2(len, idx)
		return true
	}
}
func rewriteValueMIPS_OpIsNonNil(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (IsNonNil ptr)
	// result: (SGTU ptr (MOVWconst [0]))
	for {
		ptr := v_0
		v.reset(OpMIPSSGTU)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg2(ptr, v0)
		return true
	}
}
func rewriteValueMIPS_OpIsSliceInBounds(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (IsSliceInBounds idx len)
	// result: (XORconst [1] (SGTU idx len))
	for {
		idx := v_0
		len := v_1
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGTU, typ.Bool)
		v0.AddArg2(idx, len)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpLeq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq16 x y)
	// result: (XORconst [1] (SGT (SignExt16to32 x) (SignExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGT, typ.Bool)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpLeq16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq16U x y)
	// result: (XORconst [1] (SGTU (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGTU, typ.Bool)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpLeq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq32 x y)
	// result: (XORconst [1] (SGT x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGT, typ.Bool)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpLeq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Leq32F x y)
	// result: (FPFlagTrue (CMPGEF y x))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPGEF, types.TypeFlags)
		v0.AddArg2(y, x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpLeq32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq32U x y)
	// result: (XORconst [1] (SGTU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGTU, typ.Bool)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpLeq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Leq64F x y)
	// result: (FPFlagTrue (CMPGED y x))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPGED, types.TypeFlags)
		v0.AddArg2(y, x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpLeq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq8 x y)
	// result: (XORconst [1] (SGT (SignExt8to32 x) (SignExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGT, typ.Bool)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpLeq8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq8U x y)
	// result: (XORconst [1] (SGTU (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGTU, typ.Bool)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpLess16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less16 x y)
	// result: (SGT (SignExt16to32 y) (SignExt16to32 x))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSSGT)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(x)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueMIPS_OpLess16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less16U x y)
	// result: (SGTU (ZeroExt16to32 y) (ZeroExt16to32 x))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSSGTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueMIPS_OpLess32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Less32 x y)
	// result: (SGT y x)
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSSGT)
		v.AddArg2(y, x)
		return true
	}
}
func rewriteValueMIPS_OpLess32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Less32F x y)
	// result: (FPFlagTrue (CMPGTF y x))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPGTF, types.TypeFlags)
		v0.AddArg2(y, x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpLess32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Less32U x y)
	// result: (SGTU y x)
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSSGTU)
		v.AddArg2(y, x)
		return true
	}
}
func rewriteValueMIPS_OpLess64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Less64F x y)
	// result: (FPFlagTrue (CMPGTD y x))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPGTD, types.TypeFlags)
		v0.AddArg2(y, x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpLess8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less8 x y)
	// result: (SGT (SignExt8to32 y) (SignExt8to32 x))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSSGT)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(x)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueMIPS_OpLess8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less8U x y)
	// result: (SGTU (ZeroExt8to32 y) (ZeroExt8to32 x))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSSGTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueMIPS_OpLoad(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Load <t> ptr mem)
	// cond: t.IsBoolean()
	// result: (MOVBUload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(t.IsBoolean()) {
			break
		}
		v.reset(OpMIPSMOVBUload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is8BitInt(t) && isSigned(t))
	// result: (MOVBload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is8BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpMIPSMOVBload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is8BitInt(t) && !isSigned(t))
	// result: (MOVBUload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is8BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpMIPSMOVBUload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is16BitInt(t) && isSigned(t))
	// result: (MOVHload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is16BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpMIPSMOVHload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is16BitInt(t) && !isSigned(t))
	// result: (MOVHUload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is16BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpMIPSMOVHUload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is32BitInt(t) || isPtr(t))
	// result: (MOVWload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is32BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpMIPSMOVWload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is32BitFloat(t)
	// result: (MOVFload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is32BitFloat(t)) {
			break
		}
		v.reset(OpMIPSMOVFload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is64BitFloat(t)
	// result: (MOVDload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is64BitFloat(t)) {
			break
		}
		v.reset(OpMIPSMOVDload)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpLocalAddr(v *Value) bool {
	v_0 := v.Args[0]
	// match: (LocalAddr {sym} base _)
	// result: (MOVWaddr {sym} base)
	for {
		sym := v.Aux
		base := v_0
		v.reset(OpMIPSMOVWaddr)
		v.Aux = sym
		v.AddArg(base)
		return true
	}
}
func rewriteValueMIPS_OpLsh16x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x16 <t> x y)
	// result: (CMOVZ (SLL <t> x (ZeroExt16to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt16to32 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSLL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg3(v0, v2, v3)
		return true
	}
}
func rewriteValueMIPS_OpLsh16x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x32 <t> x y)
	// result: (CMOVZ (SLL <t> x y) (MOVWconst [0]) (SGTUconst [32] y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSLL, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = 0
		v2 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v2.AuxInt = 32
		v2.AddArg(y)
		v.AddArg3(v0, v1, v2)
		return true
	}
}
func rewriteValueMIPS_OpLsh16x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Lsh16x64 x (Const64 [c]))
	// cond: uint32(c) < 16
	// result: (SLLconst x [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 16) {
			break
		}
		v.reset(OpMIPSSLLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh16x64 _ (Const64 [c]))
	// cond: uint32(c) >= 16
	// result: (MOVWconst [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) >= 16) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS_OpLsh16x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x8 <t> x y)
	// result: (CMOVZ (SLL <t> x (ZeroExt8to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt8to32 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSLL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg3(v0, v2, v3)
		return true
	}
}
func rewriteValueMIPS_OpLsh32x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x16 <t> x y)
	// result: (CMOVZ (SLL <t> x (ZeroExt16to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt16to32 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSLL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg3(v0, v2, v3)
		return true
	}
}
func rewriteValueMIPS_OpLsh32x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x32 <t> x y)
	// result: (CMOVZ (SLL <t> x y) (MOVWconst [0]) (SGTUconst [32] y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSLL, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = 0
		v2 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v2.AuxInt = 32
		v2.AddArg(y)
		v.AddArg3(v0, v1, v2)
		return true
	}
}
func rewriteValueMIPS_OpLsh32x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Lsh32x64 x (Const64 [c]))
	// cond: uint32(c) < 32
	// result: (SLLconst x [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 32) {
			break
		}
		v.reset(OpMIPSSLLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh32x64 _ (Const64 [c]))
	// cond: uint32(c) >= 32
	// result: (MOVWconst [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) >= 32) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS_OpLsh32x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x8 <t> x y)
	// result: (CMOVZ (SLL <t> x (ZeroExt8to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt8to32 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSLL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg3(v0, v2, v3)
		return true
	}
}
func rewriteValueMIPS_OpLsh8x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x16 <t> x y)
	// result: (CMOVZ (SLL <t> x (ZeroExt16to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt16to32 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSLL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg3(v0, v2, v3)
		return true
	}
}
func rewriteValueMIPS_OpLsh8x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x32 <t> x y)
	// result: (CMOVZ (SLL <t> x y) (MOVWconst [0]) (SGTUconst [32] y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSLL, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = 0
		v2 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v2.AuxInt = 32
		v2.AddArg(y)
		v.AddArg3(v0, v1, v2)
		return true
	}
}
func rewriteValueMIPS_OpLsh8x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Lsh8x64 x (Const64 [c]))
	// cond: uint32(c) < 8
	// result: (SLLconst x [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 8) {
			break
		}
		v.reset(OpMIPSSLLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh8x64 _ (Const64 [c]))
	// cond: uint32(c) >= 8
	// result: (MOVWconst [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) >= 8) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS_OpLsh8x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x8 <t> x y)
	// result: (CMOVZ (SLL <t> x (ZeroExt8to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt8to32 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSLL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg3(v0, v2, v3)
		return true
	}
}
func rewriteValueMIPS_OpMIPSADD(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ADD x (MOVWconst [c]))
	// result: (ADDconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMIPSMOVWconst {
				continue
			}
			c := v_1.AuxInt
			v.reset(OpMIPSADDconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ADD x (NEG y))
	// result: (SUB x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMIPSNEG {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpMIPSSUB)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	return false
}
func rewriteValueMIPS_OpMIPSADDconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ADDconst [off1] (MOVWaddr [off2] {sym} ptr))
	// result: (MOVWaddr [off1+off2] {sym} ptr)
	for {
		off1 := v.AuxInt
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym := v_0.Aux
		ptr := v_0.Args[0]
		v.reset(OpMIPSMOVWaddr)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		return true
	}
	// match: (ADDconst [0] x)
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v_0
		v.copyOf(x)
		return true
	}
	// match: (ADDconst [c] (MOVWconst [d]))
	// result: (MOVWconst [int64(int32(c+d))])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(c + d))
		return true
	}
	// match: (ADDconst [c] (ADDconst [d] x))
	// result: (ADDconst [int64(int32(c+d))] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPSADDconst)
		v.AuxInt = int64(int32(c + d))
		v.AddArg(x)
		return true
	}
	// match: (ADDconst [c] (SUBconst [d] x))
	// result: (ADDconst [int64(int32(c-d))] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSSUBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPSADDconst)
		v.AuxInt = int64(int32(c - d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSAND(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (AND x (MOVWconst [c]))
	// result: (ANDconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMIPSMOVWconst {
				continue
			}
			c := v_1.AuxInt
			v.reset(OpMIPSANDconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (AND x x)
	// result: x
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (AND (SGTUconst [1] x) (SGTUconst [1] y))
	// result: (SGTUconst [1] (OR <x.Type> x y))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpMIPSSGTUconst || v_0.AuxInt != 1 {
				continue
			}
			x := v_0.Args[0]
			if v_1.Op != OpMIPSSGTUconst || v_1.AuxInt != 1 {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpMIPSSGTUconst)
			v.AuxInt = 1
			v0 := b.NewValue0(v.Pos, OpMIPSOR, x.Type)
			v0.AddArg2(x, y)
			v.AddArg(v0)
			return true
		}
		break
	}
	return false
}
func rewriteValueMIPS_OpMIPSANDconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ANDconst [0] _)
	// result: (MOVWconst [0])
	for {
		if v.AuxInt != 0 {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (ANDconst [-1] x)
	// result: x
	for {
		if v.AuxInt != -1 {
			break
		}
		x := v_0
		v.copyOf(x)
		return true
	}
	// match: (ANDconst [c] (MOVWconst [d]))
	// result: (MOVWconst [c&d])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = c & d
		return true
	}
	// match: (ANDconst [c] (ANDconst [d] x))
	// result: (ANDconst [c&d] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSANDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPSANDconst)
		v.AuxInt = c & d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSCMOVZ(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (CMOVZ _ f (MOVWconst [0]))
	// result: f
	for {
		f := v_1
		if v_2.Op != OpMIPSMOVWconst || v_2.AuxInt != 0 {
			break
		}
		v.copyOf(f)
		return true
	}
	// match: (CMOVZ a _ (MOVWconst [c]))
	// cond: c!=0
	// result: a
	for {
		a := v_0
		if v_2.Op != OpMIPSMOVWconst {
			break
		}
		c := v_2.AuxInt
		if !(c != 0) {
			break
		}
		v.copyOf(a)
		return true
	}
	// match: (CMOVZ a (MOVWconst [0]) c)
	// result: (CMOVZzero a c)
	for {
		a := v_0
		if v_1.Op != OpMIPSMOVWconst || v_1.AuxInt != 0 {
			break
		}
		c := v_2
		v.reset(OpMIPSCMOVZzero)
		v.AddArg2(a, c)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSCMOVZzero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (CMOVZzero _ (MOVWconst [0]))
	// result: (MOVWconst [0])
	for {
		if v_1.Op != OpMIPSMOVWconst || v_1.AuxInt != 0 {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (CMOVZzero a (MOVWconst [c]))
	// cond: c!=0
	// result: a
	for {
		a := v_0
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(c != 0) {
			break
		}
		v.copyOf(a)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSLoweredAtomicAdd(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (LoweredAtomicAdd ptr (MOVWconst [c]) mem)
	// cond: is16Bit(c)
	// result: (LoweredAtomicAddconst [c] ptr mem)
	for {
		ptr := v_0
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		mem := v_2
		if !(is16Bit(c)) {
			break
		}
		v.reset(OpMIPSLoweredAtomicAddconst)
		v.AuxInt = c
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSLoweredAtomicStore32(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (LoweredAtomicStore32 ptr (MOVWconst [0]) mem)
	// result: (LoweredAtomicStorezero ptr mem)
	for {
		ptr := v_0
		if v_1.Op != OpMIPSMOVWconst || v_1.AuxInt != 0 {
			break
		}
		mem := v_2
		v.reset(OpMIPSLoweredAtomicStorezero)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVBUload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBUload [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVBUload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v_0
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v_1
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVBUload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBUload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVBUload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVBUload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBUload [off] {sym} ptr (MOVBstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: (MOVBUreg x)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPSMOVBstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpMIPSMOVBUreg)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVBUreg(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (MOVBUreg x:(MOVBUload _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpMIPSMOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBUreg x:(MOVBUreg _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpMIPSMOVBUreg {
			break
		}
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBUreg <t> x:(MOVBload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVBUload <t> [off] {sym} ptr mem)
	for {
		t := v.Type
		x := v_0
		if x.Op != OpMIPSMOVBload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[1]
		ptr := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpMIPSMOVBUload, t)
		v.copyOf(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBUreg (ANDconst [c] x))
	// result: (ANDconst [c&0xff] x)
	for {
		if v_0.Op != OpMIPSANDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPSANDconst)
		v.AuxInt = c & 0xff
		v.AddArg(x)
		return true
	}
	// match: (MOVBUreg (MOVWconst [c]))
	// result: (MOVWconst [int64(uint8(c))])
	for {
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(uint8(c))
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVBload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBload [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVBload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v_0
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v_1
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVBload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVBload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVBload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBload [off] {sym} ptr (MOVBstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: (MOVBreg x)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPSMOVBstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpMIPSMOVBreg)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVBreg(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (MOVBreg x:(MOVBload _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpMIPSMOVBload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg x:(MOVBreg _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpMIPSMOVBreg {
			break
		}
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg <t> x:(MOVBUload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVBload <t> [off] {sym} ptr mem)
	for {
		t := v.Type
		x := v_0
		if x.Op != OpMIPSMOVBUload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[1]
		ptr := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpMIPSMOVBload, t)
		v.copyOf(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBreg (ANDconst [c] x))
	// cond: c & 0x80 == 0
	// result: (ANDconst [c&0x7f] x)
	for {
		if v_0.Op != OpMIPSANDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c&0x80 == 0) {
			break
		}
		v.reset(OpMIPSANDconst)
		v.AuxInt = c & 0x7f
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg (MOVWconst [c]))
	// result: (MOVWconst [int64(int8(c))])
	for {
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int8(c))
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVBstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBstore [off1] {sym} x:(ADDconst [off2] ptr) val mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVBstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v_0
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		val := v_1
		mem := v_2
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVBstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVBstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVWconst [0]) mem)
	// result: (MOVBstorezero [off] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPSMOVWconst || v_1.AuxInt != 0 {
			break
		}
		mem := v_2
		v.reset(OpMIPSMOVBstorezero)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVBreg x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPSMOVBreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg3(ptr, x, mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVBUreg x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPSMOVBUreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg3(ptr, x, mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVHreg x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPSMOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg3(ptr, x, mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVHUreg x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPSMOVHUreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg3(ptr, x, mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVWreg x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPSMOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg3(ptr, x, mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVBstorezero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBstorezero [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVBstorezero [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v_0
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v_1
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVBstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBstorezero [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVBstorezero [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVBstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVDload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVDload [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVDload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v_0
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v_1
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVDload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVDload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVDload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVDload [off] {sym} ptr (MOVDstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: x
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPSMOVDstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVDstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVDstore [off1] {sym} x:(ADDconst [off2] ptr) val mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVDstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v_0
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		val := v_1
		mem := v_2
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVDstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVDstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVFload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVFload [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVFload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v_0
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v_1
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVFload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVFload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVFload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVFload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVFload [off] {sym} ptr (MOVFstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: x
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPSMOVFstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVFstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVFstore [off1] {sym} x:(ADDconst [off2] ptr) val mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVFstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v_0
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		val := v_1
		mem := v_2
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVFstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVFstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVFstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVFstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVHUload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHUload [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVHUload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v_0
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v_1
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVHUload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHUload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVHUload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVHUload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHUload [off] {sym} ptr (MOVHstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: (MOVHUreg x)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPSMOVHstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpMIPSMOVHUreg)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVHUreg(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (MOVHUreg x:(MOVBUload _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpMIPSMOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg x:(MOVHUload _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpMIPSMOVHUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg x:(MOVBUreg _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpMIPSMOVBUreg {
			break
		}
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg x:(MOVHUreg _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpMIPSMOVHUreg {
			break
		}
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg <t> x:(MOVHload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVHUload <t> [off] {sym} ptr mem)
	for {
		t := v.Type
		x := v_0
		if x.Op != OpMIPSMOVHload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[1]
		ptr := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpMIPSMOVHUload, t)
		v.copyOf(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHUreg (ANDconst [c] x))
	// result: (ANDconst [c&0xffff] x)
	for {
		if v_0.Op != OpMIPSANDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPSANDconst)
		v.AuxInt = c & 0xffff
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg (MOVWconst [c]))
	// result: (MOVWconst [int64(uint16(c))])
	for {
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(uint16(c))
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVHload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHload [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVHload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v_0
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v_1
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVHload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVHload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVHload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHload [off] {sym} ptr (MOVHstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: (MOVHreg x)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPSMOVHstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpMIPSMOVHreg)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVHreg(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (MOVHreg x:(MOVBload _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpMIPSMOVBload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBUload _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpMIPSMOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVHload _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpMIPSMOVHload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBreg _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpMIPSMOVBreg {
			break
		}
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBUreg _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpMIPSMOVBUreg {
			break
		}
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVHreg _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpMIPSMOVHreg {
			break
		}
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg <t> x:(MOVHUload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVHload <t> [off] {sym} ptr mem)
	for {
		t := v.Type
		x := v_0
		if x.Op != OpMIPSMOVHUload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[1]
		ptr := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpMIPSMOVHload, t)
		v.copyOf(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHreg (ANDconst [c] x))
	// cond: c & 0x8000 == 0
	// result: (ANDconst [c&0x7fff] x)
	for {
		if v_0.Op != OpMIPSANDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c&0x8000 == 0) {
			break
		}
		v.reset(OpMIPSANDconst)
		v.AuxInt = c & 0x7fff
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg (MOVWconst [c]))
	// result: (MOVWconst [int64(int16(c))])
	for {
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int16(c))
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVHstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHstore [off1] {sym} x:(ADDconst [off2] ptr) val mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVHstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v_0
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		val := v_1
		mem := v_2
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVHstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVHstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVHstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVHstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVWconst [0]) mem)
	// result: (MOVHstorezero [off] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPSMOVWconst || v_1.AuxInt != 0 {
			break
		}
		mem := v_2
		v.reset(OpMIPSMOVHstorezero)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVHreg x) mem)
	// result: (MOVHstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPSMOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpMIPSMOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg3(ptr, x, mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVHUreg x) mem)
	// result: (MOVHstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPSMOVHUreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpMIPSMOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg3(ptr, x, mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVWreg x) mem)
	// result: (MOVHstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPSMOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpMIPSMOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg3(ptr, x, mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVHstorezero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHstorezero [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVHstorezero [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v_0
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v_1
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVHstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHstorezero [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVHstorezero [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVHstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVWload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWload [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVWload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v_0
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v_1
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVWload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVWload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVWload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVWload [off] {sym} ptr (MOVWstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: x
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPSMOVWstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVWreg(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MOVWreg x)
	// cond: x.Uses == 1
	// result: (MOVWnop x)
	for {
		x := v_0
		if !(x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVWnop)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg (MOVWconst [c]))
	// result: (MOVWconst [c])
	for {
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = c
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVWstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWstore [off1] {sym} x:(ADDconst [off2] ptr) val mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVWstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v_0
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		val := v_1
		mem := v_2
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVWstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVWstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVWstore [off] {sym} ptr (MOVWconst [0]) mem)
	// result: (MOVWstorezero [off] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPSMOVWconst || v_1.AuxInt != 0 {
			break
		}
		mem := v_2
		v.reset(OpMIPSMOVWstorezero)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVWstore [off] {sym} ptr (MOVWreg x) mem)
	// result: (MOVWstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPSMOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpMIPSMOVWstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg3(ptr, x, mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVWstorezero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWstorezero [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVWstorezero [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v_0
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v_1
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVWstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVWstorezero [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVWstorezero [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVWstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMUL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MUL (MOVWconst [0]) _ )
	// result: (MOVWconst [0])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpMIPSMOVWconst || v_0.AuxInt != 0 {
				continue
			}
			v.reset(OpMIPSMOVWconst)
			v.AuxInt = 0
			return true
		}
		break
	}
	// match: (MUL (MOVWconst [1]) x )
	// result: x
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpMIPSMOVWconst || v_0.AuxInt != 1 {
				continue
			}
			x := v_1
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (MUL (MOVWconst [-1]) x )
	// result: (NEG x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpMIPSMOVWconst || v_0.AuxInt != -1 {
				continue
			}
			x := v_1
			v.reset(OpMIPSNEG)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (MUL (MOVWconst [c]) x )
	// cond: isPowerOfTwo(int64(uint32(c)))
	// result: (SLLconst [log2(int64(uint32(c)))] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpMIPSMOVWconst {
				continue
			}
			c := v_0.AuxInt
			x := v_1
			if !(isPowerOfTwo(int64(uint32(c)))) {
				continue
			}
			v.reset(OpMIPSSLLconst)
			v.AuxInt = log2(int64(uint32(c)))
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (MUL (MOVWconst [c]) (MOVWconst [d]))
	// result: (MOVWconst [int64(int32(c)*int32(d))])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpMIPSMOVWconst {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpMIPSMOVWconst {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpMIPSMOVWconst)
			v.AuxInt = int64(int32(c) * int32(d))
			return true
		}
		break
	}
	return false
}
func rewriteValueMIPS_OpMIPSNEG(v *Value) bool {
	v_0 := v.Args[0]
	// match: (NEG (MOVWconst [c]))
	// result: (MOVWconst [int64(int32(-c))])
	for {
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(-c))
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSNOR(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (NOR x (MOVWconst [c]))
	// result: (NORconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMIPSMOVWconst {
				continue
			}
			c := v_1.AuxInt
			v.reset(OpMIPSNORconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValueMIPS_OpMIPSNORconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (NORconst [c] (MOVWconst [d]))
	// result: (MOVWconst [^(c|d)])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = ^(c | d)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSOR(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (OR x (MOVWconst [c]))
	// result: (ORconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMIPSMOVWconst {
				continue
			}
			c := v_1.AuxInt
			v.reset(OpMIPSORconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (OR x x)
	// result: x
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (OR (SGTUzero x) (SGTUzero y))
	// result: (SGTUzero (OR <x.Type> x y))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpMIPSSGTUzero {
				continue
			}
			x := v_0.Args[0]
			if v_1.Op != OpMIPSSGTUzero {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpMIPSSGTUzero)
			v0 := b.NewValue0(v.Pos, OpMIPSOR, x.Type)
			v0.AddArg2(x, y)
			v.AddArg(v0)
			return true
		}
		break
	}
	return false
}
func rewriteValueMIPS_OpMIPSORconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ORconst [0] x)
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v_0
		v.copyOf(x)
		return true
	}
	// match: (ORconst [-1] _)
	// result: (MOVWconst [-1])
	for {
		if v.AuxInt != -1 {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = -1
		return true
	}
	// match: (ORconst [c] (MOVWconst [d]))
	// result: (MOVWconst [c|d])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = c | d
		return true
	}
	// match: (ORconst [c] (ORconst [d] x))
	// result: (ORconst [c|d] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSORconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPSORconst)
		v.AuxInt = c | d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSGT(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SGT (MOVWconst [c]) x)
	// result: (SGTconst [c] x)
	for {
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpMIPSSGTconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (SGT x (MOVWconst [0]))
	// result: (SGTzero x)
	for {
		x := v_0
		if v_1.Op != OpMIPSMOVWconst || v_1.AuxInt != 0 {
			break
		}
		v.reset(OpMIPSSGTzero)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSGTU(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SGTU (MOVWconst [c]) x)
	// result: (SGTUconst [c] x)
	for {
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpMIPSSGTUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (SGTU x (MOVWconst [0]))
	// result: (SGTUzero x)
	for {
		x := v_0
		if v_1.Op != OpMIPSMOVWconst || v_1.AuxInt != 0 {
			break
		}
		v.reset(OpMIPSSGTUzero)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSGTUconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SGTUconst [c] (MOVWconst [d]))
	// cond: uint32(c)>uint32(d)
	// result: (MOVWconst [1])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		if !(uint32(c) > uint32(d)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTUconst [c] (MOVWconst [d]))
	// cond: uint32(c)<=uint32(d)
	// result: (MOVWconst [0])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		if !(uint32(c) <= uint32(d)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (SGTUconst [c] (MOVBUreg _))
	// cond: 0xff < uint32(c)
	// result: (MOVWconst [1])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSMOVBUreg || !(0xff < uint32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTUconst [c] (MOVHUreg _))
	// cond: 0xffff < uint32(c)
	// result: (MOVWconst [1])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSMOVHUreg || !(0xffff < uint32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTUconst [c] (ANDconst [m] _))
	// cond: uint32(m) < uint32(c)
	// result: (MOVWconst [1])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSANDconst {
			break
		}
		m := v_0.AuxInt
		if !(uint32(m) < uint32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTUconst [c] (SRLconst _ [d]))
	// cond: uint32(d) <= 31 && 0xffffffff>>uint32(d) < uint32(c)
	// result: (MOVWconst [1])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSSRLconst {
			break
		}
		d := v_0.AuxInt
		if !(uint32(d) <= 31 && 0xffffffff>>uint32(d) < uint32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSGTUzero(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SGTUzero (MOVWconst [d]))
	// cond: uint32(d) != 0
	// result: (MOVWconst [1])
	for {
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		if !(uint32(d) != 0) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTUzero (MOVWconst [d]))
	// cond: uint32(d) == 0
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		if !(uint32(d) == 0) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSGTconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SGTconst [c] (MOVWconst [d]))
	// cond: int32(c) > int32(d)
	// result: (MOVWconst [1])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		if !(int32(c) > int32(d)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTconst [c] (MOVWconst [d]))
	// cond: int32(c) <= int32(d)
	// result: (MOVWconst [0])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		if !(int32(c) <= int32(d)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (SGTconst [c] (MOVBreg _))
	// cond: 0x7f < int32(c)
	// result: (MOVWconst [1])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSMOVBreg || !(0x7f < int32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTconst [c] (MOVBreg _))
	// cond: int32(c) <= -0x80
	// result: (MOVWconst [0])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSMOVBreg || !(int32(c) <= -0x80) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (SGTconst [c] (MOVBUreg _))
	// cond: 0xff < int32(c)
	// result: (MOVWconst [1])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSMOVBUreg || !(0xff < int32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTconst [c] (MOVBUreg _))
	// cond: int32(c) < 0
	// result: (MOVWconst [0])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSMOVBUreg || !(int32(c) < 0) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (SGTconst [c] (MOVHreg _))
	// cond: 0x7fff < int32(c)
	// result: (MOVWconst [1])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSMOVHreg || !(0x7fff < int32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTconst [c] (MOVHreg _))
	// cond: int32(c) <= -0x8000
	// result: (MOVWconst [0])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSMOVHreg || !(int32(c) <= -0x8000) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (SGTconst [c] (MOVHUreg _))
	// cond: 0xffff < int32(c)
	// result: (MOVWconst [1])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSMOVHUreg || !(0xffff < int32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTconst [c] (MOVHUreg _))
	// cond: int32(c) < 0
	// result: (MOVWconst [0])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSMOVHUreg || !(int32(c) < 0) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (SGTconst [c] (ANDconst [m] _))
	// cond: 0 <= int32(m) && int32(m) < int32(c)
	// result: (MOVWconst [1])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSANDconst {
			break
		}
		m := v_0.AuxInt
		if !(0 <= int32(m) && int32(m) < int32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTconst [c] (SRLconst _ [d]))
	// cond: 0 <= int32(c) && uint32(d) <= 31 && 0xffffffff>>uint32(d) < uint32(c)
	// result: (MOVWconst [1])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSSRLconst {
			break
		}
		d := v_0.AuxInt
		if !(0 <= int32(c) && uint32(d) <= 31 && 0xffffffff>>uint32(d) < uint32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSGTzero(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SGTzero (MOVWconst [d]))
	// cond: int32(d) > 0
	// result: (MOVWconst [1])
	for {
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		if !(int32(d) > 0) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTzero (MOVWconst [d]))
	// cond: int32(d) <= 0
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		if !(int32(d) <= 0) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SLL _ (MOVWconst [c]))
	// cond: uint32(c)>=32
	// result: (MOVWconst [0])
	for {
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) >= 32) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (SLL x (MOVWconst [c]))
	// result: (SLLconst x [c])
	for {
		x := v_0
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPSSLLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSLLconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SLLconst [c] (MOVWconst [d]))
	// result: (MOVWconst [int64(int32(uint32(d)<<uint32(c)))])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(uint32(d) << uint32(c)))
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SRA x (MOVWconst [c]))
	// cond: uint32(c)>=32
	// result: (SRAconst x [31])
	for {
		x := v_0
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) >= 32) {
			break
		}
		v.reset(OpMIPSSRAconst)
		v.AuxInt = 31
		v.AddArg(x)
		return true
	}
	// match: (SRA x (MOVWconst [c]))
	// result: (SRAconst x [c])
	for {
		x := v_0
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPSSRAconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSRAconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SRAconst [c] (MOVWconst [d]))
	// result: (MOVWconst [int64(int32(d)>>uint32(c))])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(d) >> uint32(c))
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SRL _ (MOVWconst [c]))
	// cond: uint32(c)>=32
	// result: (MOVWconst [0])
	for {
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) >= 32) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (SRL x (MOVWconst [c]))
	// result: (SRLconst x [c])
	for {
		x := v_0
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPSSRLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSRLconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SRLconst [c] (MOVWconst [d]))
	// result: (MOVWconst [int64(uint32(d)>>uint32(c))])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(uint32(d) >> uint32(c))
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSUB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SUB x (MOVWconst [c]))
	// result: (SUBconst [c] x)
	for {
		x := v_0
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPSSUBconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (SUB x x)
	// result: (MOVWconst [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (SUB (MOVWconst [0]) x)
	// result: (NEG x)
	for {
		if v_0.Op != OpMIPSMOVWconst || v_0.AuxInt != 0 {
			break
		}
		x := v_1
		v.reset(OpMIPSNEG)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSUBconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SUBconst [0] x)
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v_0
		v.copyOf(x)
		return true
	}
	// match: (SUBconst [c] (MOVWconst [d]))
	// result: (MOVWconst [int64(int32(d-c))])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(d - c))
		return true
	}
	// match: (SUBconst [c] (SUBconst [d] x))
	// result: (ADDconst [int64(int32(-c-d))] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSSUBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPSADDconst)
		v.AuxInt = int64(int32(-c - d))
		v.AddArg(x)
		return true
	}
	// match: (SUBconst [c] (ADDconst [d] x))
	// result: (ADDconst [int64(int32(-c+d))] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPSADDconst)
		v.AuxInt = int64(int32(-c + d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSXOR(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (XOR x (MOVWconst [c]))
	// result: (XORconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMIPSMOVWconst {
				continue
			}
			c := v_1.AuxInt
			v.reset(OpMIPSXORconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (XOR x x)
	// result: (MOVWconst [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSXORconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (XORconst [0] x)
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v_0
		v.copyOf(x)
		return true
	}
	// match: (XORconst [-1] x)
	// result: (NORconst [0] x)
	for {
		if v.AuxInt != -1 {
			break
		}
		x := v_0
		v.reset(OpMIPSNORconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}
	// match: (XORconst [c] (MOVWconst [d]))
	// result: (MOVWconst [c^d])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = c ^ d
		return true
	}
	// match: (XORconst [c] (XORconst [d] x))
	// result: (XORconst [c^d] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPSXORconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPSXORconst)
		v.AuxInt = c ^ d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMod16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16 x y)
	// result: (Select0 (DIV (SignExt16to32 x) (SignExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPSDIV, types.NewTuple(typ.Int32, typ.Int32))
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpMod16u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16u x y)
	// result: (Select0 (DIVU (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPSDIVU, types.NewTuple(typ.UInt32, typ.UInt32))
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpMod32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod32 x y)
	// result: (Select0 (DIV x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPSDIV, types.NewTuple(typ.Int32, typ.Int32))
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpMod32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod32u x y)
	// result: (Select0 (DIVU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPSDIVU, types.NewTuple(typ.UInt32, typ.UInt32))
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpMod8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8 x y)
	// result: (Select0 (DIV (SignExt8to32 x) (SignExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPSDIV, types.NewTuple(typ.Int32, typ.Int32))
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpMod8u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8u x y)
	// result: (Select0 (DIVU (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPSDIVU, types.NewTuple(typ.UInt32, typ.UInt32))
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpMove(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (Move [0] _ _ mem)
	// result: mem
	for {
		if v.AuxInt != 0 {
			break
		}
		mem := v_2
		v.copyOf(mem)
		return true
	}
	// match: (Move [1] dst src mem)
	// result: (MOVBstore dst (MOVBUload src mem) mem)
	for {
		if v.AuxInt != 1 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpMIPSMOVBstore)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVBUload, typ.UInt8)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [2] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore dst (MOVHUload src mem) mem)
	for {
		if v.AuxInt != 2 {
			break
		}
		t := v.Aux
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.(*types.Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpMIPSMOVHstore)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVHUload, typ.UInt16)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [2] dst src mem)
	// result: (MOVBstore [1] dst (MOVBUload [1] src mem) (MOVBstore dst (MOVBUload src mem) mem))
	for {
		if v.AuxInt != 2 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSMOVBUload, typ.UInt8)
		v0.AuxInt = 1
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVBstore, types.TypeMem)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVBUload, typ.UInt8)
		v2.AddArg2(src, mem)
		v1.AddArg3(dst, v2, mem)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [4] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%4 == 0
	// result: (MOVWstore dst (MOVWload src mem) mem)
	for {
		if v.AuxInt != 4 {
			break
		}
		t := v.Aux
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.(*types.Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpMIPSMOVWstore)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWload, typ.UInt32)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [4] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore [2] dst (MOVHUload [2] src mem) (MOVHstore dst (MOVHUload src mem) mem))
	for {
		if v.AuxInt != 4 {
			break
		}
		t := v.Aux
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.(*types.Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpMIPSMOVHstore)
		v.AuxInt = 2
		v0 := b.NewValue0(v.Pos, OpMIPSMOVHUload, typ.UInt16)
		v0.AuxInt = 2
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVHstore, types.TypeMem)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVHUload, typ.UInt16)
		v2.AddArg2(src, mem)
		v1.AddArg3(dst, v2, mem)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [4] dst src mem)
	// result: (MOVBstore [3] dst (MOVBUload [3] src mem) (MOVBstore [2] dst (MOVBUload [2] src mem) (MOVBstore [1] dst (MOVBUload [1] src mem) (MOVBstore dst (MOVBUload src mem) mem))))
	for {
		if v.AuxInt != 4 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = 3
		v0 := b.NewValue0(v.Pos, OpMIPSMOVBUload, typ.UInt8)
		v0.AuxInt = 3
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVBstore, types.TypeMem)
		v1.AuxInt = 2
		v2 := b.NewValue0(v.Pos, OpMIPSMOVBUload, typ.UInt8)
		v2.AuxInt = 2
		v2.AddArg2(src, mem)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVBstore, types.TypeMem)
		v3.AuxInt = 1
		v4 := b.NewValue0(v.Pos, OpMIPSMOVBUload, typ.UInt8)
		v4.AuxInt = 1
		v4.AddArg2(src, mem)
		v5 := b.NewValue0(v.Pos, OpMIPSMOVBstore, types.TypeMem)
		v6 := b.NewValue0(v.Pos, OpMIPSMOVBUload, typ.UInt8)
		v6.AddArg2(src, mem)
		v5.AddArg3(dst, v6, mem)
		v3.AddArg3(dst, v4, v5)
		v1.AddArg3(dst, v2, v3)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [3] dst src mem)
	// result: (MOVBstore [2] dst (MOVBUload [2] src mem) (MOVBstore [1] dst (MOVBUload [1] src mem) (MOVBstore dst (MOVBUload src mem) mem)))
	for {
		if v.AuxInt != 3 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = 2
		v0 := b.NewValue0(v.Pos, OpMIPSMOVBUload, typ.UInt8)
		v0.AuxInt = 2
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVBstore, types.TypeMem)
		v1.AuxInt = 1
		v2 := b.NewValue0(v.Pos, OpMIPSMOVBUload, typ.UInt8)
		v2.AuxInt = 1
		v2.AddArg2(src, mem)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVBstore, types.TypeMem)
		v4 := b.NewValue0(v.Pos, OpMIPSMOVBUload, typ.UInt8)
		v4.AddArg2(src, mem)
		v3.AddArg3(dst, v4, mem)
		v1.AddArg3(dst, v2, v3)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [8] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%4 == 0
	// result: (MOVWstore [4] dst (MOVWload [4] src mem) (MOVWstore dst (MOVWload src mem) mem))
	for {
		if v.AuxInt != 8 {
			break
		}
		t := v.Aux
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.(*types.Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpMIPSMOVWstore)
		v.AuxInt = 4
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWload, typ.UInt32)
		v0.AuxInt = 4
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWstore, types.TypeMem)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWload, typ.UInt32)
		v2.AddArg2(src, mem)
		v1.AddArg3(dst, v2, mem)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [8] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore [6] dst (MOVHload [6] src mem) (MOVHstore [4] dst (MOVHload [4] src mem) (MOVHstore [2] dst (MOVHload [2] src mem) (MOVHstore dst (MOVHload src mem) mem))))
	for {
		if v.AuxInt != 8 {
			break
		}
		t := v.Aux
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.(*types.Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpMIPSMOVHstore)
		v.AuxInt = 6
		v0 := b.NewValue0(v.Pos, OpMIPSMOVHload, typ.Int16)
		v0.AuxInt = 6
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVHstore, types.TypeMem)
		v1.AuxInt = 4
		v2 := b.NewValue0(v.Pos, OpMIPSMOVHload, typ.Int16)
		v2.AuxInt = 4
		v2.AddArg2(src, mem)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVHstore, types.TypeMem)
		v3.AuxInt = 2
		v4 := b.NewValue0(v.Pos, OpMIPSMOVHload, typ.Int16)
		v4.AuxInt = 2
		v4.AddArg2(src, mem)
		v5 := b.NewValue0(v.Pos, OpMIPSMOVHstore, types.TypeMem)
		v6 := b.NewValue0(v.Pos, OpMIPSMOVHload, typ.Int16)
		v6.AddArg2(src, mem)
		v5.AddArg3(dst, v6, mem)
		v3.AddArg3(dst, v4, v5)
		v1.AddArg3(dst, v2, v3)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [6] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore [4] dst (MOVHload [4] src mem) (MOVHstore [2] dst (MOVHload [2] src mem) (MOVHstore dst (MOVHload src mem) mem)))
	for {
		if v.AuxInt != 6 {
			break
		}
		t := v.Aux
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.(*types.Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpMIPSMOVHstore)
		v.AuxInt = 4
		v0 := b.NewValue0(v.Pos, OpMIPSMOVHload, typ.Int16)
		v0.AuxInt = 4
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVHstore, types.TypeMem)
		v1.AuxInt = 2
		v2 := b.NewValue0(v.Pos, OpMIPSMOVHload, typ.Int16)
		v2.AuxInt = 2
		v2.AddArg2(src, mem)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVHstore, types.TypeMem)
		v4 := b.NewValue0(v.Pos, OpMIPSMOVHload, typ.Int16)
		v4.AddArg2(src, mem)
		v3.AddArg3(dst, v4, mem)
		v1.AddArg3(dst, v2, v3)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [12] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%4 == 0
	// result: (MOVWstore [8] dst (MOVWload [8] src mem) (MOVWstore [4] dst (MOVWload [4] src mem) (MOVWstore dst (MOVWload src mem) mem)))
	for {
		if v.AuxInt != 12 {
			break
		}
		t := v.Aux
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.(*types.Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpMIPSMOVWstore)
		v.AuxInt = 8
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWload, typ.UInt32)
		v0.AuxInt = 8
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWstore, types.TypeMem)
		v1.AuxInt = 4
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWload, typ.UInt32)
		v2.AuxInt = 4
		v2.AddArg2(src, mem)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWstore, types.TypeMem)
		v4 := b.NewValue0(v.Pos, OpMIPSMOVWload, typ.UInt32)
		v4.AddArg2(src, mem)
		v3.AddArg3(dst, v4, mem)
		v1.AddArg3(dst, v2, v3)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [16] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%4 == 0
	// result: (MOVWstore [12] dst (MOVWload [12] src mem) (MOVWstore [8] dst (MOVWload [8] src mem) (MOVWstore [4] dst (MOVWload [4] src mem) (MOVWstore dst (MOVWload src mem) mem))))
	for {
		if v.AuxInt != 16 {
			break
		}
		t := v.Aux
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.(*types.Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpMIPSMOVWstore)
		v.AuxInt = 12
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWload, typ.UInt32)
		v0.AuxInt = 12
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWstore, types.TypeMem)
		v1.AuxInt = 8
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWload, typ.UInt32)
		v2.AuxInt = 8
		v2.AddArg2(src, mem)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWstore, types.TypeMem)
		v3.AuxInt = 4
		v4 := b.NewValue0(v.Pos, OpMIPSMOVWload, typ.UInt32)
		v4.AuxInt = 4
		v4.AddArg2(src, mem)
		v5 := b.NewValue0(v.Pos, OpMIPSMOVWstore, types.TypeMem)
		v6 := b.NewValue0(v.Pos, OpMIPSMOVWload, typ.UInt32)
		v6.AddArg2(src, mem)
		v5.AddArg3(dst, v6, mem)
		v3.AddArg3(dst, v4, v5)
		v1.AddArg3(dst, v2, v3)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [s] {t} dst src mem)
	// cond: (s > 16 || t.(*types.Type).Alignment()%4 != 0)
	// result: (LoweredMove [t.(*types.Type).Alignment()] dst src (ADDconst <src.Type> src [s-moveSize(t.(*types.Type).Alignment(), config)]) mem)
	for {
		s := v.AuxInt
		t := v.Aux
		dst := v_0
		src := v_1
		mem := v_2
		if !(s > 16 || t.(*types.Type).Alignment()%4 != 0) {
			break
		}
		v.reset(OpMIPSLoweredMove)
		v.AuxInt = t.(*types.Type).Alignment()
		v0 := b.NewValue0(v.Pos, OpMIPSADDconst, src.Type)
		v0.AuxInt = s - moveSize(t.(*types.Type).Alignment(), config)
		v0.AddArg(src)
		v.AddArg4(dst, src, v0, mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpNeq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq16 x y)
	// result: (SGTU (XOR (ZeroExt16to32 x) (ZeroExt16to32 y)) (MOVWconst [0]))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSSGTU)
		v0 := b.NewValue0(v.Pos, OpMIPSXOR, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v3.AuxInt = 0
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValueMIPS_OpNeq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq32 x y)
	// result: (SGTU (XOR x y) (MOVWconst [0]))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSSGTU)
		v0 := b.NewValue0(v.Pos, OpMIPSXOR, typ.UInt32)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = 0
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueMIPS_OpNeq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Neq32F x y)
	// result: (FPFlagFalse (CMPEQF x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSFPFlagFalse)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPEQF, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpNeq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Neq64F x y)
	// result: (FPFlagFalse (CMPEQD x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSFPFlagFalse)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPEQD, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpNeq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq8 x y)
	// result: (SGTU (XOR (ZeroExt8to32 x) (ZeroExt8to32 y)) (MOVWconst [0]))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSSGTU)
		v0 := b.NewValue0(v.Pos, OpMIPSXOR, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v3.AuxInt = 0
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValueMIPS_OpNeqPtr(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (NeqPtr x y)
	// result: (SGTU (XOR x y) (MOVWconst [0]))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSSGTU)
		v0 := b.NewValue0(v.Pos, OpMIPSXOR, typ.UInt32)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = 0
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueMIPS_OpNot(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Not x)
	// result: (XORconst [1] x)
	for {
		x := v_0
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpOffPtr(v *Value) bool {
	v_0 := v.Args[0]
	// match: (OffPtr [off] ptr:(SP))
	// result: (MOVWaddr [off] ptr)
	for {
		off := v.AuxInt
		ptr := v_0
		if ptr.Op != OpSP {
			break
		}
		v.reset(OpMIPSMOVWaddr)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}
	// match: (OffPtr [off] ptr)
	// result: (ADDconst [off] ptr)
	for {
		off := v.AuxInt
		ptr := v_0
		v.reset(OpMIPSADDconst)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}
}
func rewriteValueMIPS_OpPanicBounds(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (PanicBounds [kind] x y mem)
	// cond: boundsABI(kind) == 0
	// result: (LoweredPanicBoundsA [kind] x y mem)
	for {
		kind := v.AuxInt
		x := v_0
		y := v_1
		mem := v_2
		if !(boundsABI(kind) == 0) {
			break
		}
		v.reset(OpMIPSLoweredPanicBoundsA)
		v.AuxInt = kind
		v.AddArg3(x, y, mem)
		return true
	}
	// match: (PanicBounds [kind] x y mem)
	// cond: boundsABI(kind) == 1
	// result: (LoweredPanicBoundsB [kind] x y mem)
	for {
		kind := v.AuxInt
		x := v_0
		y := v_1
		mem := v_2
		if !(boundsABI(kind) == 1) {
			break
		}
		v.reset(OpMIPSLoweredPanicBoundsB)
		v.AuxInt = kind
		v.AddArg3(x, y, mem)
		return true
	}
	// match: (PanicBounds [kind] x y mem)
	// cond: boundsABI(kind) == 2
	// result: (LoweredPanicBoundsC [kind] x y mem)
	for {
		kind := v.AuxInt
		x := v_0
		y := v_1
		mem := v_2
		if !(boundsABI(kind) == 2) {
			break
		}
		v.reset(OpMIPSLoweredPanicBoundsC)
		v.AuxInt = kind
		v.AddArg3(x, y, mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpPanicExtend(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (PanicExtend [kind] hi lo y mem)
	// cond: boundsABI(kind) == 0
	// result: (LoweredPanicExtendA [kind] hi lo y mem)
	for {
		kind := v.AuxInt
		hi := v_0
		lo := v_1
		y := v_2
		mem := v_3
		if !(boundsABI(kind) == 0) {
			break
		}
		v.reset(OpMIPSLoweredPanicExtendA)
		v.AuxInt = kind
		v.AddArg4(hi, lo, y, mem)
		return true
	}
	// match: (PanicExtend [kind] hi lo y mem)
	// cond: boundsABI(kind) == 1
	// result: (LoweredPanicExtendB [kind] hi lo y mem)
	for {
		kind := v.AuxInt
		hi := v_0
		lo := v_1
		y := v_2
		mem := v_3
		if !(boundsABI(kind) == 1) {
			break
		}
		v.reset(OpMIPSLoweredPanicExtendB)
		v.AuxInt = kind
		v.AddArg4(hi, lo, y, mem)
		return true
	}
	// match: (PanicExtend [kind] hi lo y mem)
	// cond: boundsABI(kind) == 2
	// result: (LoweredPanicExtendC [kind] hi lo y mem)
	for {
		kind := v.AuxInt
		hi := v_0
		lo := v_1
		y := v_2
		mem := v_3
		if !(boundsABI(kind) == 2) {
			break
		}
		v.reset(OpMIPSLoweredPanicExtendC)
		v.AuxInt = kind
		v.AddArg4(hi, lo, y, mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpRotateLeft16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft16 <t> x (MOVWconst [c]))
	// result: (Or16 (Lsh16x32 <t> x (MOVWconst [c&15])) (Rsh16Ux32 <t> x (MOVWconst [-c&15])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOr16)
		v0 := b.NewValue0(v.Pos, OpLsh16x32, t)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = c & 15
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpRsh16Ux32, t)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v3.AuxInt = -c & 15
		v2.AddArg2(x, v3)
		v.AddArg2(v0, v2)
		return true
	}
	return false
}
func rewriteValueMIPS_OpRotateLeft32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft32 <t> x (MOVWconst [c]))
	// result: (Or32 (Lsh32x32 <t> x (MOVWconst [c&31])) (Rsh32Ux32 <t> x (MOVWconst [-c&31])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOr32)
		v0 := b.NewValue0(v.Pos, OpLsh32x32, t)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = c & 31
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpRsh32Ux32, t)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v3.AuxInt = -c & 31
		v2.AddArg2(x, v3)
		v.AddArg2(v0, v2)
		return true
	}
	return false
}
func rewriteValueMIPS_OpRotateLeft64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft64 <t> x (MOVWconst [c]))
	// result: (Or64 (Lsh64x32 <t> x (MOVWconst [c&63])) (Rsh64Ux32 <t> x (MOVWconst [-c&63])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOr64)
		v0 := b.NewValue0(v.Pos, OpLsh64x32, t)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = c & 63
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpRsh64Ux32, t)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v3.AuxInt = -c & 63
		v2.AddArg2(x, v3)
		v.AddArg2(v0, v2)
		return true
	}
	return false
}
func rewriteValueMIPS_OpRotateLeft8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft8 <t> x (MOVWconst [c]))
	// result: (Or8 (Lsh8x32 <t> x (MOVWconst [c&7])) (Rsh8Ux32 <t> x (MOVWconst [-c&7])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOr8)
		v0 := b.NewValue0(v.Pos, OpLsh8x32, t)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = c & 7
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpRsh8Ux32, t)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v3.AuxInt = -c & 7
		v2.AddArg2(x, v3)
		v.AddArg2(v0, v2)
		return true
	}
	return false
}
func rewriteValueMIPS_OpRsh16Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux16 <t> x y)
	// result: (CMOVZ (SRL <t> (ZeroExt16to32 x) (ZeroExt16to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt16to32 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v3.AuxInt = 0
		v4 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v4.AuxInt = 32
		v5 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg3(v0, v3, v4)
		return true
	}
}
func rewriteValueMIPS_OpRsh16Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux32 <t> x y)
	// result: (CMOVZ (SRL <t> (ZeroExt16to32 x) y) (MOVWconst [0]) (SGTUconst [32] y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg2(v1, y)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v3.AuxInt = 32
		v3.AddArg(y)
		v.AddArg3(v0, v2, v3)
		return true
	}
}
func rewriteValueMIPS_OpRsh16Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux64 x (Const64 [c]))
	// cond: uint32(c) < 16
	// result: (SRLconst (SLLconst <typ.UInt32> x [16]) [c+16])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 16) {
			break
		}
		v.reset(OpMIPSSRLconst)
		v.AuxInt = c + 16
		v0 := b.NewValue0(v.Pos, OpMIPSSLLconst, typ.UInt32)
		v0.AuxInt = 16
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16Ux64 _ (Const64 [c]))
	// cond: uint32(c) >= 16
	// result: (MOVWconst [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) >= 16) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS_OpRsh16Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux8 <t> x y)
	// result: (CMOVZ (SRL <t> (ZeroExt16to32 x) (ZeroExt8to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt8to32 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v3.AuxInt = 0
		v4 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v4.AuxInt = 32
		v5 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg3(v0, v3, v4)
		return true
	}
}
func rewriteValueMIPS_OpRsh16x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x16 x y)
	// result: (SRA (SignExt16to32 x) ( CMOVZ <typ.UInt32> (ZeroExt16to32 y) (MOVWconst [-1]) (SGTUconst [32] (ZeroExt16to32 y))))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpMIPSCMOVZ, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v3.AuxInt = -1
		v4 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v4.AuxInt = 32
		v5 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v5.AddArg(y)
		v4.AddArg(v5)
		v1.AddArg3(v2, v3, v4)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueMIPS_OpRsh16x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x32 x y)
	// result: (SRA (SignExt16to32 x) ( CMOVZ <typ.UInt32> y (MOVWconst [-1]) (SGTUconst [32] y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpMIPSCMOVZ, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = -1
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v3.AuxInt = 32
		v3.AddArg(y)
		v1.AddArg3(y, v2, v3)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueMIPS_OpRsh16x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x64 x (Const64 [c]))
	// cond: uint32(c) < 16
	// result: (SRAconst (SLLconst <typ.UInt32> x [16]) [c+16])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 16) {
			break
		}
		v.reset(OpMIPSSRAconst)
		v.AuxInt = c + 16
		v0 := b.NewValue0(v.Pos, OpMIPSSLLconst, typ.UInt32)
		v0.AuxInt = 16
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16x64 x (Const64 [c]))
	// cond: uint32(c) >= 16
	// result: (SRAconst (SLLconst <typ.UInt32> x [16]) [31])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) >= 16) {
			break
		}
		v.reset(OpMIPSSRAconst)
		v.AuxInt = 31
		v0 := b.NewValue0(v.Pos, OpMIPSSLLconst, typ.UInt32)
		v0.AuxInt = 16
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueMIPS_OpRsh16x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x8 x y)
	// result: (SRA (SignExt16to32 x) ( CMOVZ <typ.UInt32> (ZeroExt8to32 y) (MOVWconst [-1]) (SGTUconst [32] (ZeroExt8to32 y))))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpMIPSCMOVZ, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v3.AuxInt = -1
		v4 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v4.AuxInt = 32
		v5 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v5.AddArg(y)
		v4.AddArg(v5)
		v1.AddArg3(v2, v3, v4)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueMIPS_OpRsh32Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux16 <t> x y)
	// result: (CMOVZ (SRL <t> x (ZeroExt16to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt16to32 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg3(v0, v2, v3)
		return true
	}
}
func rewriteValueMIPS_OpRsh32Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux32 <t> x y)
	// result: (CMOVZ (SRL <t> x y) (MOVWconst [0]) (SGTUconst [32] y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSRL, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = 0
		v2 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v2.AuxInt = 32
		v2.AddArg(y)
		v.AddArg3(v0, v1, v2)
		return true
	}
}
func rewriteValueMIPS_OpRsh32Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Rsh32Ux64 x (Const64 [c]))
	// cond: uint32(c) < 32
	// result: (SRLconst x [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 32) {
			break
		}
		v.reset(OpMIPSSRLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh32Ux64 _ (Const64 [c]))
	// cond: uint32(c) >= 32
	// result: (MOVWconst [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) >= 32) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS_OpRsh32Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux8 <t> x y)
	// result: (CMOVZ (SRL <t> x (ZeroExt8to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt8to32 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg3(v0, v2, v3)
		return true
	}
}
func rewriteValueMIPS_OpRsh32x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x16 x y)
	// result: (SRA x ( CMOVZ <typ.UInt32> (ZeroExt16to32 y) (MOVWconst [-1]) (SGTUconst [32] (ZeroExt16to32 y))))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSSRA)
		v0 := b.NewValue0(v.Pos, OpMIPSCMOVZ, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = -1
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v0.AddArg3(v1, v2, v3)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueMIPS_OpRsh32x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x32 x y)
	// result: (SRA x ( CMOVZ <typ.UInt32> y (MOVWconst [-1]) (SGTUconst [32] y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSSRA)
		v0 := b.NewValue0(v.Pos, OpMIPSCMOVZ, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = -1
		v2 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v2.AuxInt = 32
		v2.AddArg(y)
		v0.AddArg3(y, v1, v2)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueMIPS_OpRsh32x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Rsh32x64 x (Const64 [c]))
	// cond: uint32(c) < 32
	// result: (SRAconst x [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 32) {
			break
		}
		v.reset(OpMIPSSRAconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh32x64 x (Const64 [c]))
	// cond: uint32(c) >= 32
	// result: (SRAconst x [31])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) >= 32) {
			break
		}
		v.reset(OpMIPSSRAconst)
		v.AuxInt = 31
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpRsh32x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x8 x y)
	// result: (SRA x ( CMOVZ <typ.UInt32> (ZeroExt8to32 y) (MOVWconst [-1]) (SGTUconst [32] (ZeroExt8to32 y))))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSSRA)
		v0 := b.NewValue0(v.Pos, OpMIPSCMOVZ, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = -1
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v0.AddArg3(v1, v2, v3)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueMIPS_OpRsh8Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux16 <t> x y)
	// result: (CMOVZ (SRL <t> (ZeroExt8to32 x) (ZeroExt16to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt16to32 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v3.AuxInt = 0
		v4 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v4.AuxInt = 32
		v5 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg3(v0, v3, v4)
		return true
	}
}
func rewriteValueMIPS_OpRsh8Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux32 <t> x y)
	// result: (CMOVZ (SRL <t> (ZeroExt8to32 x) y) (MOVWconst [0]) (SGTUconst [32] y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg2(v1, y)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v3.AuxInt = 32
		v3.AddArg(y)
		v.AddArg3(v0, v2, v3)
		return true
	}
}
func rewriteValueMIPS_OpRsh8Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux64 x (Const64 [c]))
	// cond: uint32(c) < 8
	// result: (SRLconst (SLLconst <typ.UInt32> x [24]) [c+24])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 8) {
			break
		}
		v.reset(OpMIPSSRLconst)
		v.AuxInt = c + 24
		v0 := b.NewValue0(v.Pos, OpMIPSSLLconst, typ.UInt32)
		v0.AuxInt = 24
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8Ux64 _ (Const64 [c]))
	// cond: uint32(c) >= 8
	// result: (MOVWconst [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) >= 8) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS_OpRsh8Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux8 <t> x y)
	// result: (CMOVZ (SRL <t> (ZeroExt8to32 x) (ZeroExt8to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt8to32 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v3.AuxInt = 0
		v4 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v4.AuxInt = 32
		v5 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg3(v0, v3, v4)
		return true
	}
}
func rewriteValueMIPS_OpRsh8x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x16 x y)
	// result: (SRA (SignExt16to32 x) ( CMOVZ <typ.UInt32> (ZeroExt16to32 y) (MOVWconst [-1]) (SGTUconst [32] (ZeroExt16to32 y))))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpMIPSCMOVZ, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v3.AuxInt = -1
		v4 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v4.AuxInt = 32
		v5 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v5.AddArg(y)
		v4.AddArg(v5)
		v1.AddArg3(v2, v3, v4)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueMIPS_OpRsh8x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x32 x y)
	// result: (SRA (SignExt16to32 x) ( CMOVZ <typ.UInt32> y (MOVWconst [-1]) (SGTUconst [32] y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpMIPSCMOVZ, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = -1
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v3.AuxInt = 32
		v3.AddArg(y)
		v1.AddArg3(y, v2, v3)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueMIPS_OpRsh8x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x64 x (Const64 [c]))
	// cond: uint32(c) < 8
	// result: (SRAconst (SLLconst <typ.UInt32> x [24]) [c+24])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 8) {
			break
		}
		v.reset(OpMIPSSRAconst)
		v.AuxInt = c + 24
		v0 := b.NewValue0(v.Pos, OpMIPSSLLconst, typ.UInt32)
		v0.AuxInt = 24
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8x64 x (Const64 [c]))
	// cond: uint32(c) >= 8
	// result: (SRAconst (SLLconst <typ.UInt32> x [24]) [31])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) >= 8) {
			break
		}
		v.reset(OpMIPSSRAconst)
		v.AuxInt = 31
		v0 := b.NewValue0(v.Pos, OpMIPSSLLconst, typ.UInt32)
		v0.AuxInt = 24
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueMIPS_OpRsh8x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x8 x y)
	// result: (SRA (SignExt16to32 x) ( CMOVZ <typ.UInt32> (ZeroExt8to32 y) (MOVWconst [-1]) (SGTUconst [32] (ZeroExt8to32 y))))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPSSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpMIPSCMOVZ, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v3.AuxInt = -1
		v4 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v4.AuxInt = 32
		v5 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v5.AddArg(y)
		v4.AddArg(v5)
		v1.AddArg3(v2, v3, v4)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueMIPS_OpSelect0(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Select0 (Add32carry <t> x y))
	// result: (ADD <t.FieldType(0)> x y)
	for {
		if v_0.Op != OpAdd32carry {
			break
		}
		t := v_0.Type
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpMIPSADD)
		v.Type = t.FieldType(0)
		v.AddArg2(x, y)
		return true
	}
	// match: (Select0 (Sub32carry <t> x y))
	// result: (SUB <t.FieldType(0)> x y)
	for {
		if v_0.Op != OpSub32carry {
			break
		}
		t := v_0.Type
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpMIPSSUB)
		v.Type = t.FieldType(0)
		v.AddArg2(x, y)
		return true
	}
	// match: (Select0 (MULTU (MOVWconst [0]) _ ))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpMIPSMOVWconst || v_0_0.AuxInt != 0 {
				continue
			}
			v.reset(OpMIPSMOVWconst)
			v.AuxInt = 0
			return true
		}
		break
	}
	// match: (Select0 (MULTU (MOVWconst [1]) _ ))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpMIPSMOVWconst || v_0_0.AuxInt != 1 {
				continue
			}
			v.reset(OpMIPSMOVWconst)
			v.AuxInt = 0
			return true
		}
		break
	}
	// match: (Select0 (MULTU (MOVWconst [-1]) x ))
	// result: (CMOVZ (ADDconst <x.Type> [-1] x) (MOVWconst [0]) x)
	for {
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpMIPSMOVWconst || v_0_0.AuxInt != -1 {
				continue
			}
			x := v_0_1
			v.reset(OpMIPSCMOVZ)
			v0 := b.NewValue0(v.Pos, OpMIPSADDconst, x.Type)
			v0.AuxInt = -1
			v0.AddArg(x)
			v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
			v1.AuxInt = 0
			v.AddArg3(v0, v1, x)
			return true
		}
		break
	}
	// match: (Select0 (MULTU (MOVWconst [c]) x ))
	// cond: isPowerOfTwo(int64(uint32(c)))
	// result: (SRLconst [32-log2(int64(uint32(c)))] x)
	for {
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpMIPSMOVWconst {
				continue
			}
			c := v_0_0.AuxInt
			x := v_0_1
			if !(isPowerOfTwo(int64(uint32(c)))) {
				continue
			}
			v.reset(OpMIPSSRLconst)
			v.AuxInt = 32 - log2(int64(uint32(c)))
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (Select0 (MULTU (MOVWconst [c]) (MOVWconst [d])))
	// result: (MOVWconst [(c*d)>>32])
	for {
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpMIPSMOVWconst {
				continue
			}
			c := v_0_0.AuxInt
			if v_0_1.Op != OpMIPSMOVWconst {
				continue
			}
			d := v_0_1.AuxInt
			v.reset(OpMIPSMOVWconst)
			v.AuxInt = (c * d) >> 32
			return true
		}
		break
	}
	// match: (Select0 (DIV (MOVWconst [c]) (MOVWconst [d])))
	// result: (MOVWconst [int64(int32(c)%int32(d))])
	for {
		if v_0.Op != OpMIPSDIV {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(c) % int32(d))
		return true
	}
	// match: (Select0 (DIVU (MOVWconst [c]) (MOVWconst [d])))
	// result: (MOVWconst [int64(int32(uint32(c)%uint32(d)))])
	for {
		if v_0.Op != OpMIPSDIVU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(uint32(c) % uint32(d)))
		return true
	}
	return false
}
func rewriteValueMIPS_OpSelect1(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Select1 (Add32carry <t> x y))
	// result: (SGTU <typ.Bool> x (ADD <t.FieldType(0)> x y))
	for {
		if v_0.Op != OpAdd32carry {
			break
		}
		t := v_0.Type
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpMIPSSGTU)
		v.Type = typ.Bool
		v0 := b.NewValue0(v.Pos, OpMIPSADD, t.FieldType(0))
		v0.AddArg2(x, y)
		v.AddArg2(x, v0)
		return true
	}
	// match: (Select1 (Sub32carry <t> x y))
	// result: (SGTU <typ.Bool> (SUB <t.FieldType(0)> x y) x)
	for {
		if v_0.Op != OpSub32carry {
			break
		}
		t := v_0.Type
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpMIPSSGTU)
		v.Type = typ.Bool
		v0 := b.NewValue0(v.Pos, OpMIPSSUB, t.FieldType(0))
		v0.AddArg2(x, y)
		v.AddArg2(v0, x)
		return true
	}
	// match: (Select1 (MULTU (MOVWconst [0]) _ ))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpMIPSMOVWconst || v_0_0.AuxInt != 0 {
				continue
			}
			v.reset(OpMIPSMOVWconst)
			v.AuxInt = 0
			return true
		}
		break
	}
	// match: (Select1 (MULTU (MOVWconst [1]) x ))
	// result: x
	for {
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpMIPSMOVWconst || v_0_0.AuxInt != 1 {
				continue
			}
			x := v_0_1
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (Select1 (MULTU (MOVWconst [-1]) x ))
	// result: (NEG <x.Type> x)
	for {
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpMIPSMOVWconst || v_0_0.AuxInt != -1 {
				continue
			}
			x := v_0_1
			v.reset(OpMIPSNEG)
			v.Type = x.Type
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (Select1 (MULTU (MOVWconst [c]) x ))
	// cond: isPowerOfTwo(int64(uint32(c)))
	// result: (SLLconst [log2(int64(uint32(c)))] x)
	for {
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpMIPSMOVWconst {
				continue
			}
			c := v_0_0.AuxInt
			x := v_0_1
			if !(isPowerOfTwo(int64(uint32(c)))) {
				continue
			}
			v.reset(OpMIPSSLLconst)
			v.AuxInt = log2(int64(uint32(c)))
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (Select1 (MULTU (MOVWconst [c]) (MOVWconst [d])))
	// result: (MOVWconst [int64(int32(uint32(c)*uint32(d)))])
	for {
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpMIPSMOVWconst {
				continue
			}
			c := v_0_0.AuxInt
			if v_0_1.Op != OpMIPSMOVWconst {
				continue
			}
			d := v_0_1.AuxInt
			v.reset(OpMIPSMOVWconst)
			v.AuxInt = int64(int32(uint32(c) * uint32(d)))
			return true
		}
		break
	}
	// match: (Select1 (DIV (MOVWconst [c]) (MOVWconst [d])))
	// result: (MOVWconst [int64(int32(c)/int32(d))])
	for {
		if v_0.Op != OpMIPSDIV {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(c) / int32(d))
		return true
	}
	// match: (Select1 (DIVU (MOVWconst [c]) (MOVWconst [d])))
	// result: (MOVWconst [int64(int32(uint32(c)/uint32(d)))])
	for {
		if v_0.Op != OpMIPSDIVU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(uint32(c) / uint32(d)))
		return true
	}
	return false
}
func rewriteValueMIPS_OpSignmask(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Signmask x)
	// result: (SRAconst x [31])
	for {
		x := v_0
		v.reset(OpMIPSSRAconst)
		v.AuxInt = 31
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpSlicemask(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (Slicemask <t> x)
	// result: (SRAconst (NEG <t> x) [31])
	for {
		t := v.Type
		x := v_0
		v.reset(OpMIPSSRAconst)
		v.AuxInt = 31
		v0 := b.NewValue0(v.Pos, OpMIPSNEG, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpStore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 1
	// result: (MOVBstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 1) {
			break
		}
		v.reset(OpMIPSMOVBstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 2
	// result: (MOVHstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 2) {
			break
		}
		v.reset(OpMIPSMOVHstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 4 && !is32BitFloat(val.Type)
	// result: (MOVWstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 4 && !is32BitFloat(val.Type)) {
			break
		}
		v.reset(OpMIPSMOVWstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 4 && is32BitFloat(val.Type)
	// result: (MOVFstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 4 && is32BitFloat(val.Type)) {
			break
		}
		v.reset(OpMIPSMOVFstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 8 && is64BitFloat(val.Type)
	// result: (MOVDstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 8 && is64BitFloat(val.Type)) {
			break
		}
		v.reset(OpMIPSMOVDstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpSub32withcarry(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Sub32withcarry <t> x y c)
	// result: (SUB (SUB <t> x y) c)
	for {
		t := v.Type
		x := v_0
		y := v_1
		c := v_2
		v.reset(OpMIPSSUB)
		v0 := b.NewValue0(v.Pos, OpMIPSSUB, t)
		v0.AddArg2(x, y)
		v.AddArg2(v0, c)
		return true
	}
}
func rewriteValueMIPS_OpZero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (Zero [0] _ mem)
	// result: mem
	for {
		if v.AuxInt != 0 {
			break
		}
		mem := v_1
		v.copyOf(mem)
		return true
	}
	// match: (Zero [1] ptr mem)
	// result: (MOVBstore ptr (MOVWconst [0]) mem)
	for {
		if v.AuxInt != 1 {
			break
		}
		ptr := v_0
		mem := v_1
		v.reset(OpMIPSMOVBstore)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg3(ptr, v0, mem)
		return true
	}
	// match: (Zero [2] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore ptr (MOVWconst [0]) mem)
	for {
		if v.AuxInt != 2 {
			break
		}
		t := v.Aux
		ptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpMIPSMOVHstore)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg3(ptr, v0, mem)
		return true
	}
	// match: (Zero [2] ptr mem)
	// result: (MOVBstore [1] ptr (MOVWconst [0]) (MOVBstore [0] ptr (MOVWconst [0]) mem))
	for {
		if v.AuxInt != 2 {
			break
		}
		ptr := v_0
		mem := v_1
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v1 := b.NewValue0(v.Pos, OpMIPSMOVBstore, types.TypeMem)
		v1.AuxInt = 0
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v1.AddArg3(ptr, v2, mem)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [4] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%4 == 0
	// result: (MOVWstore ptr (MOVWconst [0]) mem)
	for {
		if v.AuxInt != 4 {
			break
		}
		t := v.Aux
		ptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpMIPSMOVWstore)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg3(ptr, v0, mem)
		return true
	}
	// match: (Zero [4] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore [2] ptr (MOVWconst [0]) (MOVHstore [0] ptr (MOVWconst [0]) mem))
	for {
		if v.AuxInt != 4 {
			break
		}
		t := v.Aux
		ptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpMIPSMOVHstore)
		v.AuxInt = 2
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v1 := b.NewValue0(v.Pos, OpMIPSMOVHstore, types.TypeMem)
		v1.AuxInt = 0
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v1.AddArg3(ptr, v2, mem)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [4] ptr mem)
	// result: (MOVBstore [3] ptr (MOVWconst [0]) (MOVBstore [2] ptr (MOVWconst [0]) (MOVBstore [1] ptr (MOVWconst [0]) (MOVBstore [0] ptr (MOVWconst [0]) mem))))
	for {
		if v.AuxInt != 4 {
			break
		}
		ptr := v_0
		mem := v_1
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = 3
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v1 := b.NewValue0(v.Pos, OpMIPSMOVBstore, types.TypeMem)
		v1.AuxInt = 2
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v3 := b.NewValue0(v.Pos, OpMIPSMOVBstore, types.TypeMem)
		v3.AuxInt = 1
		v4 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v4.AuxInt = 0
		v5 := b.NewValue0(v.Pos, OpMIPSMOVBstore, types.TypeMem)
		v5.AuxInt = 0
		v6 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v6.AuxInt = 0
		v5.AddArg3(ptr, v6, mem)
		v3.AddArg3(ptr, v4, v5)
		v1.AddArg3(ptr, v2, v3)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [3] ptr mem)
	// result: (MOVBstore [2] ptr (MOVWconst [0]) (MOVBstore [1] ptr (MOVWconst [0]) (MOVBstore [0] ptr (MOVWconst [0]) mem)))
	for {
		if v.AuxInt != 3 {
			break
		}
		ptr := v_0
		mem := v_1
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = 2
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v1 := b.NewValue0(v.Pos, OpMIPSMOVBstore, types.TypeMem)
		v1.AuxInt = 1
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v3 := b.NewValue0(v.Pos, OpMIPSMOVBstore, types.TypeMem)
		v3.AuxInt = 0
		v4 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v4.AuxInt = 0
		v3.AddArg3(ptr, v4, mem)
		v1.AddArg3(ptr, v2, v3)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [6] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore [4] ptr (MOVWconst [0]) (MOVHstore [2] ptr (MOVWconst [0]) (MOVHstore [0] ptr (MOVWconst [0]) mem)))
	for {
		if v.AuxInt != 6 {
			break
		}
		t := v.Aux
		ptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpMIPSMOVHstore)
		v.AuxInt = 4
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v1 := b.NewValue0(v.Pos, OpMIPSMOVHstore, types.TypeMem)
		v1.AuxInt = 2
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v3 := b.NewValue0(v.Pos, OpMIPSMOVHstore, types.TypeMem)
		v3.AuxInt = 0
		v4 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v4.AuxInt = 0
		v3.AddArg3(ptr, v4, mem)
		v1.AddArg3(ptr, v2, v3)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [8] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%4 == 0
	// result: (MOVWstore [4] ptr (MOVWconst [0]) (MOVWstore [0] ptr (MOVWconst [0]) mem))
	for {
		if v.AuxInt != 8 {
			break
		}
		t := v.Aux
		ptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpMIPSMOVWstore)
		v.AuxInt = 4
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWstore, types.TypeMem)
		v1.AuxInt = 0
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v1.AddArg3(ptr, v2, mem)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [12] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%4 == 0
	// result: (MOVWstore [8] ptr (MOVWconst [0]) (MOVWstore [4] ptr (MOVWconst [0]) (MOVWstore [0] ptr (MOVWconst [0]) mem)))
	for {
		if v.AuxInt != 12 {
			break
		}
		t := v.Aux
		ptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpMIPSMOVWstore)
		v.AuxInt = 8
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWstore, types.TypeMem)
		v1.AuxInt = 4
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWstore, types.TypeMem)
		v3.AuxInt = 0
		v4 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v4.AuxInt = 0
		v3.AddArg3(ptr, v4, mem)
		v1.AddArg3(ptr, v2, v3)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [16] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%4 == 0
	// result: (MOVWstore [12] ptr (MOVWconst [0]) (MOVWstore [8] ptr (MOVWconst [0]) (MOVWstore [4] ptr (MOVWconst [0]) (MOVWstore [0] ptr (MOVWconst [0]) mem))))
	for {
		if v.AuxInt != 16 {
			break
		}
		t := v.Aux
		ptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpMIPSMOVWstore)
		v.AuxInt = 12
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWstore, types.TypeMem)
		v1.AuxInt = 8
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWstore, types.TypeMem)
		v3.AuxInt = 4
		v4 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v4.AuxInt = 0
		v5 := b.NewValue0(v.Pos, OpMIPSMOVWstore, types.TypeMem)
		v5.AuxInt = 0
		v6 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v6.AuxInt = 0
		v5.AddArg3(ptr, v6, mem)
		v3.AddArg3(ptr, v4, v5)
		v1.AddArg3(ptr, v2, v3)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [s] {t} ptr mem)
	// cond: (s > 16 || t.(*types.Type).Alignment()%4 != 0)
	// result: (LoweredZero [t.(*types.Type).Alignment()] ptr (ADDconst <ptr.Type> ptr [s-moveSize(t.(*types.Type).Alignment(), config)]) mem)
	for {
		s := v.AuxInt
		t := v.Aux
		ptr := v_0
		mem := v_1
		if !(s > 16 || t.(*types.Type).Alignment()%4 != 0) {
			break
		}
		v.reset(OpMIPSLoweredZero)
		v.AuxInt = t.(*types.Type).Alignment()
		v0 := b.NewValue0(v.Pos, OpMIPSADDconst, ptr.Type)
		v0.AuxInt = s - moveSize(t.(*types.Type).Alignment(), config)
		v0.AddArg(ptr)
		v.AddArg3(ptr, v0, mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpZeromask(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Zeromask x)
	// result: (NEG (SGTU x (MOVWconst [0])))
	for {
		x := v_0
		v.reset(OpMIPSNEG)
		v0 := b.NewValue0(v.Pos, OpMIPSSGTU, typ.Bool)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = 0
		v0.AddArg2(x, v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteBlockMIPS(b *Block) bool {
	switch b.Kind {
	case BlockMIPSEQ:
		// match: (EQ (FPFlagTrue cmp) yes no)
		// result: (FPF cmp yes no)
		for b.Controls[0].Op == OpMIPSFPFlagTrue {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockMIPSFPF)
			b.AddControl(cmp)
			return true
		}
		// match: (EQ (FPFlagFalse cmp) yes no)
		// result: (FPT cmp yes no)
		for b.Controls[0].Op == OpMIPSFPFlagFalse {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockMIPSFPT)
			b.AddControl(cmp)
			return true
		}
		// match: (EQ (XORconst [1] cmp:(SGT _ _)) yes no)
		// result: (NE cmp yes no)
		for b.Controls[0].Op == OpMIPSXORconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpMIPSSGT {
				break
			}
			_ = cmp.Args[1]
			b.Reset(BlockMIPSNE)
			b.AddControl(cmp)
			return true
		}
		// match: (EQ (XORconst [1] cmp:(SGTU _ _)) yes no)
		// result: (NE cmp yes no)
		for b.Controls[0].Op == OpMIPSXORconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpMIPSSGTU {
				break
			}
			_ = cmp.Args[1]
			b.Reset(BlockMIPSNE)
			b.AddControl(cmp)
			return true
		}
		// match: (EQ (XORconst [1] cmp:(SGTconst _)) yes no)
		// result: (NE cmp yes no)
		for b.Controls[0].Op == OpMIPSXORconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpMIPSSGTconst {
				break
			}
			b.Reset(BlockMIPSNE)
			b.AddControl(cmp)
			return true
		}
		// match: (EQ (XORconst [1] cmp:(SGTUconst _)) yes no)
		// result: (NE cmp yes no)
		for b.Controls[0].Op == OpMIPSXORconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpMIPSSGTUconst {
				break
			}
			b.Reset(BlockMIPSNE)
			b.AddControl(cmp)
			return true
		}
		// match: (EQ (XORconst [1] cmp:(SGTzero _)) yes no)
		// result: (NE cmp yes no)
		for b.Controls[0].Op == OpMIPSXORconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpMIPSSGTzero {
				break
			}
			b.Reset(BlockMIPSNE)
			b.AddControl(cmp)
			return true
		}
		// match: (EQ (XORconst [1] cmp:(SGTUzero _)) yes no)
		// result: (NE cmp yes no)
		for b.Controls[0].Op == OpMIPSXORconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpMIPSSGTUzero {
				break
			}
			b.Reset(BlockMIPSNE)
			b.AddControl(cmp)
			return true
		}
		// match: (EQ (SGTUconst [1] x) yes no)
		// result: (NE x yes no)
		for b.Controls[0].Op == OpMIPSSGTUconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 1 {
				break
			}
			x := v_0.Args[0]
			b.Reset(BlockMIPSNE)
			b.AddControl(x)
			return true
		}
		// match: (EQ (SGTUzero x) yes no)
		// result: (EQ x yes no)
		for b.Controls[0].Op == OpMIPSSGTUzero {
			v_0 := b.Controls[0]
			x := v_0.Args[0]
			b.Reset(BlockMIPSEQ)
			b.AddControl(x)
			return true
		}
		// match: (EQ (SGTconst [0] x) yes no)
		// result: (GEZ x yes no)
		for b.Controls[0].Op == OpMIPSSGTconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			x := v_0.Args[0]
			b.Reset(BlockMIPSGEZ)
			b.AddControl(x)
			return true
		}
		// match: (EQ (SGTzero x) yes no)
		// result: (LEZ x yes no)
		for b.Controls[0].Op == OpMIPSSGTzero {
			v_0 := b.Controls[0]
			x := v_0.Args[0]
			b.Reset(BlockMIPSLEZ)
			b.AddControl(x)
			return true
		}
		// match: (EQ (MOVWconst [0]) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpMIPSMOVWconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (EQ (MOVWconst [c]) yes no)
		// cond: c != 0
		// result: (First no yes)
		for b.Controls[0].Op == OpMIPSMOVWconst {
			v_0 := b.Controls[0]
			c := v_0.AuxInt
			if !(c != 0) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	case BlockMIPSGEZ:
		// match: (GEZ (MOVWconst [c]) yes no)
		// cond: int32(c) >= 0
		// result: (First yes no)
		for b.Controls[0].Op == OpMIPSMOVWconst {
			v_0 := b.Controls[0]
			c := v_0.AuxInt
			if !(int32(c) >= 0) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (GEZ (MOVWconst [c]) yes no)
		// cond: int32(c) < 0
		// result: (First no yes)
		for b.Controls[0].Op == OpMIPSMOVWconst {
			v_0 := b.Controls[0]
			c := v_0.AuxInt
			if !(int32(c) < 0) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	case BlockMIPSGTZ:
		// match: (GTZ (MOVWconst [c]) yes no)
		// cond: int32(c) > 0
		// result: (First yes no)
		for b.Controls[0].Op == OpMIPSMOVWconst {
			v_0 := b.Controls[0]
			c := v_0.AuxInt
			if !(int32(c) > 0) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (GTZ (MOVWconst [c]) yes no)
		// cond: int32(c) <= 0
		// result: (First no yes)
		for b.Controls[0].Op == OpMIPSMOVWconst {
			v_0 := b.Controls[0]
			c := v_0.AuxInt
			if !(int32(c) <= 0) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	case BlockIf:
		// match: (If cond yes no)
		// result: (NE cond yes no)
		for {
			cond := b.Controls[0]
			b.Reset(BlockMIPSNE)
			b.AddControl(cond)
			return true
		}
	case BlockMIPSLEZ:
		// match: (LEZ (MOVWconst [c]) yes no)
		// cond: int32(c) <= 0
		// result: (First yes no)
		for b.Controls[0].Op == OpMIPSMOVWconst {
			v_0 := b.Controls[0]
			c := v_0.AuxInt
			if !(int32(c) <= 0) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (LEZ (MOVWconst [c]) yes no)
		// cond: int32(c) > 0
		// result: (First no yes)
		for b.Controls[0].Op == OpMIPSMOVWconst {
			v_0 := b.Controls[0]
			c := v_0.AuxInt
			if !(int32(c) > 0) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	case BlockMIPSLTZ:
		// match: (LTZ (MOVWconst [c]) yes no)
		// cond: int32(c) < 0
		// result: (First yes no)
		for b.Controls[0].Op == OpMIPSMOVWconst {
			v_0 := b.Controls[0]
			c := v_0.AuxInt
			if !(int32(c) < 0) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (LTZ (MOVWconst [c]) yes no)
		// cond: int32(c) >= 0
		// result: (First no yes)
		for b.Controls[0].Op == OpMIPSMOVWconst {
			v_0 := b.Controls[0]
			c := v_0.AuxInt
			if !(int32(c) >= 0) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	case BlockMIPSNE:
		// match: (NE (FPFlagTrue cmp) yes no)
		// result: (FPT cmp yes no)
		for b.Controls[0].Op == OpMIPSFPFlagTrue {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockMIPSFPT)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (FPFlagFalse cmp) yes no)
		// result: (FPF cmp yes no)
		for b.Controls[0].Op == OpMIPSFPFlagFalse {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockMIPSFPF)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (XORconst [1] cmp:(SGT _ _)) yes no)
		// result: (EQ cmp yes no)
		for b.Controls[0].Op == OpMIPSXORconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpMIPSSGT {
				break
			}
			_ = cmp.Args[1]
			b.Reset(BlockMIPSEQ)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (XORconst [1] cmp:(SGTU _ _)) yes no)
		// result: (EQ cmp yes no)
		for b.Controls[0].Op == OpMIPSXORconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpMIPSSGTU {
				break
			}
			_ = cmp.Args[1]
			b.Reset(BlockMIPSEQ)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (XORconst [1] cmp:(SGTconst _)) yes no)
		// result: (EQ cmp yes no)
		for b.Controls[0].Op == OpMIPSXORconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpMIPSSGTconst {
				break
			}
			b.Reset(BlockMIPSEQ)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (XORconst [1] cmp:(SGTUconst _)) yes no)
		// result: (EQ cmp yes no)
		for b.Controls[0].Op == OpMIPSXORconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpMIPSSGTUconst {
				break
			}
			b.Reset(BlockMIPSEQ)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (XORconst [1] cmp:(SGTzero _)) yes no)
		// result: (EQ cmp yes no)
		for b.Controls[0].Op == OpMIPSXORconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpMIPSSGTzero {
				break
			}
			b.Reset(BlockMIPSEQ)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (XORconst [1] cmp:(SGTUzero _)) yes no)
		// result: (EQ cmp yes no)
		for b.Controls[0].Op == OpMIPSXORconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpMIPSSGTUzero {
				break
			}
			b.Reset(BlockMIPSEQ)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (SGTUconst [1] x) yes no)
		// result: (EQ x yes no)
		for b.Controls[0].Op == OpMIPSSGTUconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 1 {
				break
			}
			x := v_0.Args[0]
			b.Reset(BlockMIPSEQ)
			b.AddControl(x)
			return true
		}
		// match: (NE (SGTUzero x) yes no)
		// result: (NE x yes no)
		for b.Controls[0].Op == OpMIPSSGTUzero {
			v_0 := b.Controls[0]
			x := v_0.Args[0]
			b.Reset(BlockMIPSNE)
			b.AddControl(x)
			return true
		}
		// match: (NE (SGTconst [0] x) yes no)
		// result: (LTZ x yes no)
		for b.Controls[0].Op == OpMIPSSGTconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			x := v_0.Args[0]
			b.Reset(BlockMIPSLTZ)
			b.AddControl(x)
			return true
		}
		// match: (NE (SGTzero x) yes no)
		// result: (GTZ x yes no)
		for b.Controls[0].Op == OpMIPSSGTzero {
			v_0 := b.Controls[0]
			x := v_0.Args[0]
			b.Reset(BlockMIPSGTZ)
			b.AddControl(x)
			return true
		}
		// match: (NE (MOVWconst [0]) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpMIPSMOVWconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (NE (MOVWconst [c]) yes no)
		// cond: c != 0
		// result: (First yes no)
		for b.Controls[0].Op == OpMIPSMOVWconst {
			v_0 := b.Controls[0]
			c := v_0.AuxInt
			if !(c != 0) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
	}
	return false
}
