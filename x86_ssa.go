// +build ignore

// 3 reg intrinsics
	case ssa.OpAMD64IntrinMm256AddEpi32, ssa.OpAMD64IntrinMm256AddEpi64, ssa.OpAMD64IntrinMm256AndSi256, ssa.OpAMD64IntrinMm256CmpeqEpi8, ssa.OpAMD64IntrinMm256OrSi256, ssa.OpAMD64IntrinMm256ShuffleEpi8, ssa.OpAMD64IntrinMm256XorSi256:
		r := v.Reg()
		opregregreg(s, v.Op.Asm(), r, v.Args[0].Reg(), v.Args[1].Reg())
		// 1 reg intrinsics
	case ssa.OpAMD64IntrinMmAesimcSi128, ssa.OpAMD64IntrinMmBroadcastbEpi8, ssa.OpAMD64IntrinMmBroadcastsdPd, ssa.OpAMD64IntrinMmBroadcastssPs, ssa.OpAMD64IntrinMmCeilPd, ssa.OpAMD64IntrinMmCeilPs, ssa.OpAMD64IntrinMmCvtepi16Epi32, ssa.OpAMD64IntrinMmCvtepi16Epi64, ssa.OpAMD64IntrinMmCvtepi32Epi64, ssa.OpAMD64IntrinMmCvtepi8Epi16, ssa.OpAMD64IntrinMmCvtepi8Epi32, ssa.OpAMD64IntrinMmCvtepi8Epi64, ssa.OpAMD64IntrinMmCvtepu16Epi32, ssa.OpAMD64IntrinMmCvtepu16Epi64, ssa.OpAMD64IntrinMmCvtepu32Epi64, ssa.OpAMD64IntrinMmCvtepu8Epi16, ssa.OpAMD64IntrinMmCvtepu8Epi32, ssa.OpAMD64IntrinMmCvtepu8Epi64, ssa.OpAMD64IntrinMmCvtpdPs, ssa.OpAMD64IntrinMmCvtpsPd, ssa.OpAMD64IntrinMmFloorPd, ssa.OpAMD64IntrinMmFloorPs, ssa.OpAMD64IntrinMmMinposEpu16, ssa.OpAMD64IntrinMmMoveEpi64, ssa.OpAMD64IntrinMmMovedupPd, ssa.OpAMD64IntrinMm256MovedupPd, ssa.OpAMD64IntrinMmMovehdupPs, ssa.OpAMD64IntrinMm256MovehdupPs, ssa.OpAMD64IntrinMmMoveldupPs, ssa.OpAMD64IntrinMm256MoveldupPs, ssa.OpAMD64IntrinMmRcpPs, ssa.OpAMD64IntrinMmRcpSs, ssa.OpAMD64IntrinMmRsqrtPs, ssa.OpAMD64IntrinMmRsqrtSs, ssa.OpAMD64IntrinMmSqrtPd, ssa.OpAMD64IntrinMmSqrtPs, ssa.OpAMD64IntrinMmSqrtSs:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
// 2 ops intrinsics
	case ssa.OpAMD64IntrinMmAddEpi16, ssa.OpAMD64IntrinMmAddEpi64, ssa.OpAMD64IntrinMmAddEpi8, ssa.OpAMD64IntrinMmAddPd, ssa.OpAMD64IntrinMmAddPs, ssa.OpAMD64IntrinMmAddSd, ssa.OpAMD64IntrinMmAddSs, ssa.OpAMD64IntrinMmAddsEpi16, ssa.OpAMD64IntrinMmAddsEpi8, ssa.OpAMD64IntrinMmAddsEpu16, ssa.OpAMD64IntrinMmAddsEpu8, ssa.OpAMD64IntrinMmAesdecSi128, ssa.OpAMD64IntrinMmAesdeclastSi128, ssa.OpAMD64IntrinMmAesencSi128, ssa.OpAMD64IntrinMmAesenclastSi128, ssa.OpAMD64IntrinMmAndPd, ssa.OpAMD64IntrinMmAndPs, ssa.OpAMD64IntrinMmAndSi128, ssa.OpAMD64IntrinMmAndnotPd, ssa.OpAMD64IntrinMmAndnotPs, ssa.OpAMD64IntrinMmAndnotSi128, ssa.OpAMD64IntrinMmAvgEpu16, ssa.OpAMD64IntrinMmAvgEpu8, ssa.OpAMD64IntrinMmCeilSd, ssa.OpAMD64IntrinMmCeilSs, ssa.OpAMD64IntrinMmCmpeqEpi16, ssa.OpAMD64IntrinMmCmpeqEpi8, ssa.OpAMD64IntrinMmCmpeqPd, ssa.OpAMD64IntrinMmCmpeqPs, ssa.OpAMD64IntrinMmCmpeqSd, ssa.OpAMD64IntrinMmCmpeqSs, ssa.OpAMD64IntrinMmCmpgePd, ssa.OpAMD64IntrinMmCmpgePs, ssa.OpAMD64IntrinMmCmpgeSd, ssa.OpAMD64IntrinMmCmpgeSs, ssa.OpAMD64IntrinMmCmpgtEpi16, ssa.OpAMD64IntrinMmCmpgtEpi8, ssa.OpAMD64IntrinMmCmpgtPd, ssa.OpAMD64IntrinMmCmpgtPs, ssa.OpAMD64IntrinMmCmpgtSd, ssa.OpAMD64IntrinMmCmpgtSs, ssa.OpAMD64IntrinMmCmplePd, ssa.OpAMD64IntrinMmCmplePs, ssa.OpAMD64IntrinMmCmpleSd, ssa.OpAMD64IntrinMmCmpleSs, ssa.OpAMD64IntrinMmCmpltEpi16, ssa.OpAMD64IntrinMmCmpltEpi8, ssa.OpAMD64IntrinMmCmpltPd, ssa.OpAMD64IntrinMmCmpltPs, ssa.OpAMD64IntrinMmCmpltSd, ssa.OpAMD64IntrinMmCmpltSs, ssa.OpAMD64IntrinMmCmpneqPd, ssa.OpAMD64IntrinMmCmpneqPs, ssa.OpAMD64IntrinMmCmpneqSd, ssa.OpAMD64IntrinMmCmpneqSs, ssa.OpAMD64IntrinMmCmpngePd, ssa.OpAMD64IntrinMmCmpngePs, ssa.OpAMD64IntrinMmCmpngeSd, ssa.OpAMD64IntrinMmCmpngeSs, ssa.OpAMD64IntrinMmCmpngtPd, ssa.OpAMD64IntrinMmCmpngtPs, ssa.OpAMD64IntrinMmCmpngtSd, ssa.OpAMD64IntrinMmCmpngtSs, ssa.OpAMD64IntrinMmCmpnlePd, ssa.OpAMD64IntrinMmCmpnlePs, ssa.OpAMD64IntrinMmCmpnleSd, ssa.OpAMD64IntrinMmCmpnleSs, ssa.OpAMD64IntrinMmCmpnltPd, ssa.OpAMD64IntrinMmCmpnltPs, ssa.OpAMD64IntrinMmCmpnltSd, ssa.OpAMD64IntrinMmCmpnltSs, ssa.OpAMD64IntrinMmCmpordPd, ssa.OpAMD64IntrinMmCmpordPs, ssa.OpAMD64IntrinMmCmpordSd, ssa.OpAMD64IntrinMmCmpordSs, ssa.OpAMD64IntrinMmCmpunordPd, ssa.OpAMD64IntrinMmCmpunordPs, ssa.OpAMD64IntrinMmCmpunordSd, ssa.OpAMD64IntrinMmCmpunordSs, ssa.OpAMD64IntrinMmCvtsdSs, ssa.OpAMD64IntrinMmCvtssSd, ssa.OpAMD64IntrinMmDivPd, ssa.OpAMD64IntrinMmDivPs, ssa.OpAMD64IntrinMmDivSd, ssa.OpAMD64IntrinMmDivSs, ssa.OpAMD64IntrinMmFloorSd, ssa.OpAMD64IntrinMmFloorSs, ssa.OpAMD64IntrinMmHaddEpi16, ssa.OpAMD64IntrinMmHaddEpi32, ssa.OpAMD64IntrinMmHaddPd, ssa.OpAMD64IntrinMmHaddPs, ssa.OpAMD64IntrinMmHaddsEpi16, ssa.OpAMD64IntrinMmHsubEpi16, ssa.OpAMD64IntrinMmHsubEpi32, ssa.OpAMD64IntrinMmHsubPd, ssa.OpAMD64IntrinMmHsubPs, ssa.OpAMD64IntrinMmHsubsEpi16, ssa.OpAMD64IntrinMmMaxEpi16, ssa.OpAMD64IntrinMmMaxEpu8, ssa.OpAMD64IntrinMmMaxPd, ssa.OpAMD64IntrinMmMaxPs, ssa.OpAMD64IntrinMmMaxSd, ssa.OpAMD64IntrinMmMaxSs, ssa.OpAMD64IntrinMmMinEpi16, ssa.OpAMD64IntrinMmMinEpu8, ssa.OpAMD64IntrinMmMinPd, ssa.OpAMD64IntrinMmMinPs, ssa.OpAMD64IntrinMmMinSd, ssa.OpAMD64IntrinMmMinSs, ssa.OpAMD64IntrinMmMoveSd, ssa.OpAMD64IntrinMmMoveSs, ssa.OpAMD64IntrinMmMovehlPs, ssa.OpAMD64IntrinMmMovelhPs, ssa.OpAMD64IntrinMmMulEpi32, ssa.OpAMD64IntrinMmMulPd, ssa.OpAMD64IntrinMmMulPs, ssa.OpAMD64IntrinMmMulSd, ssa.OpAMD64IntrinMmMulSs, ssa.OpAMD64IntrinMmMulhiEpi16, ssa.OpAMD64IntrinMmMulhiEpu16, ssa.OpAMD64IntrinMmMulloEpi16, ssa.OpAMD64IntrinMmMulloEpi32, ssa.OpAMD64IntrinMmOrPd, ssa.OpAMD64IntrinMmOrPs, ssa.OpAMD64IntrinMmOrSi128, ssa.OpAMD64IntrinMmPacksEpi16, ssa.OpAMD64IntrinMmPackusEpi16, ssa.OpAMD64IntrinMmSadEpu8, ssa.OpAMD64IntrinMmShuffleEpi8, ssa.OpAMD64IntrinMmSllEpi16, ssa.OpAMD64IntrinMmSllEpi64, ssa.OpAMD64IntrinMmSqrtSd, ssa.OpAMD64IntrinMmSraEpi16, ssa.OpAMD64IntrinMmSrlEpi16, ssa.OpAMD64IntrinMmSrlEpi64, ssa.OpAMD64IntrinMmSubEpi16, ssa.OpAMD64IntrinMmSubEpi64, ssa.OpAMD64IntrinMmSubEpi8, ssa.OpAMD64IntrinMmSubPd, ssa.OpAMD64IntrinMmSubPs, ssa.OpAMD64IntrinMmSubSd, ssa.OpAMD64IntrinMmSubSs, ssa.OpAMD64IntrinMmSubsEpi16, ssa.OpAMD64IntrinMmSubsEpi8, ssa.OpAMD64IntrinMmSubsEpu16, ssa.OpAMD64IntrinMmSubsEpu8, ssa.OpAMD64IntrinMmUnpackhiEpi64, ssa.OpAMD64IntrinMmUnpackhiEpi8, ssa.OpAMD64IntrinMmUnpackhiPd, ssa.OpAMD64IntrinMmUnpackhiPs, ssa.OpAMD64IntrinMmUnpackloEpi64, ssa.OpAMD64IntrinMmUnpackloEpi8, ssa.OpAMD64IntrinMmUnpackloPd, ssa.OpAMD64IntrinMmUnpackloPs, ssa.OpAMD64IntrinMmXorPd, ssa.OpAMD64IntrinMmXorPs, ssa.OpAMD64IntrinMmXorSi128:
		r := v.Reg()
		if r != v.Args[0].Reg() {
			v.Fatalf("input[0] and output not in same register %s", v.LongString())
		}
		opregreg(s, v.Op.Asm(), r, v.Args[1].Reg())
		