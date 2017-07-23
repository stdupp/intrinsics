// +build ignore

	{name: "IntrinMmAddEpi16", argLength:2, reg:fp21, asm:"PADDW", resultInArg0: true},
	{name: "IntrinMm256AddEpi32", argLength:2, reg:avx21, asm:"VPADDD"},
	{name: "IntrinMmAddEpi64", argLength:2, reg:fp21, asm:"PADDQ", resultInArg0: true},
	{name: "IntrinMm256AddEpi64", argLength:2, reg:avx21, asm:"VPADDQ"},
	{name: "IntrinMmAddEpi8", argLength:2, reg:fp21, asm:"PADDB", resultInArg0: true},
	{name: "IntrinMmAddPd", argLength:2, reg:fp21, asm:"ADDPD", resultInArg0: true},
	{name: "IntrinMmAddPs", argLength:2, reg:fp21, asm:"ADDPS", resultInArg0: true},
	{name: "IntrinMmAddSd", argLength:2, reg:fp21, asm:"ADDSD", resultInArg0: true},
	{name: "IntrinMmAddSs", argLength:2, reg:fp21, asm:"ADDSS", resultInArg0: true},
	{name: "IntrinMmAddsEpi16", argLength:2, reg:fp21, asm:"PADDSW", resultInArg0: true},
	{name: "IntrinMmAddsEpi8", argLength:2, reg:fp21, asm:"PADDSB", resultInArg0: true},
	{name: "IntrinMmAddsEpu16", argLength:2, reg:fp21, asm:"PADDUSW", resultInArg0: true},
	{name: "IntrinMmAddsEpu8", argLength:2, reg:fp21, asm:"PADDUSB", resultInArg0: true},
	{name: "IntrinMmAesdecSi128", argLength:2, reg:fp21, asm:"AESDEC", resultInArg0: true},
	{name: "IntrinMmAesdeclastSi128", argLength:2, reg:fp21, asm:"AESDECLAST", resultInArg0: true},
	{name: "IntrinMmAesencSi128", argLength:2, reg:fp21, asm:"AESENC", resultInArg0: true},
	{name: "IntrinMmAesenclastSi128", argLength:2, reg:fp21, asm:"AESENCLAST", resultInArg0: true},
	{name: "IntrinMmAesimcSi128", argLength:1, reg:fp11, asm:"AESIMC", resultInArg0: true},
	{name: "IntrinMmAndPd", argLength:2, reg:fp21, asm:"ANDPD", resultInArg0: true},
	{name: "IntrinMmAndPs", argLength:2, reg:fp21, asm:"ANDPS", resultInArg0: true},
	{name: "IntrinMmAndSi128", argLength:2, reg:fp21, asm:"PAND", resultInArg0: true},
	{name: "IntrinMm256AndSi256", argLength:2, reg:avx21, asm:"VPAND"},
	{name: "IntrinMmAndnotPd", argLength:2, reg:fp21, asm:"ANDNPD", resultInArg0: true},
	{name: "IntrinMmAndnotPs", argLength:2, reg:fp21, asm:"ANDNPS", resultInArg0: true},
	{name: "IntrinMmAndnotSi128", argLength:2, reg:fp21, asm:"PANDN", resultInArg0: true},
	{name: "IntrinMmAvgEpu16", argLength:2, reg:fp21, asm:"PAVGW", resultInArg0: true},
	{name: "IntrinMmAvgEpu8", argLength:2, reg:fp21, asm:"PAVGB", resultInArg0: true},
	{name: "IntrinMmBroadcastbEpi8", argLength:1, reg:fp11, asm:"VPBROADCASTB", resultInArg0: true},
	{name: "IntrinMmBroadcastsdPd", argLength:1, reg:fp11, asm:"MOVDDUP", resultInArg0: true},
	{name: "IntrinMmBroadcastssPs", argLength:1, reg:fp11, asm:"VBROADCASTSS", resultInArg0: true},
	{name: "IntrinMmCeilPd", argLength:1, reg:fp11, asm:"ROUNDPD", resultInArg0: true},
	{name: "IntrinMmCeilPs", argLength:1, reg:fp11, asm:"ROUNDPS", resultInArg0: true},
	{name: "IntrinMmCeilSd", argLength:2, reg:fp21, asm:"ROUNDSD", resultInArg0: true},
	{name: "IntrinMmCeilSs", argLength:2, reg:fp21, asm:"ROUNDSS", resultInArg0: true},
	{name: "IntrinMmCmpeqEpi16", argLength:2, reg:fp21, asm:"PCMPEQW", resultInArg0: true},
	{name: "IntrinMmCmpeqEpi8", argLength:2, reg:fp21, asm:"PCMPEQB", resultInArg0: true},
	{name: "IntrinMm256CmpeqEpi8", argLength:2, reg:avx21, asm:"VPCMPEQB"},
	{name: "IntrinMmCmpeqPd", argLength:2, reg:fp21, asm:"CMPPD", resultInArg0: true},
	{name: "IntrinMmCmpeqPs", argLength:2, reg:fp21, asm:"CMPPS", resultInArg0: true},
	{name: "IntrinMmCmpeqSd", argLength:2, reg:fp21, asm:"CMPSD", resultInArg0: true},
	{name: "IntrinMmCmpeqSs", argLength:2, reg:fp21, asm:"CMPSS", resultInArg0: true},
	{name: "IntrinMmCmpgePd", argLength:2, reg:fp21, asm:"CMPPD", resultInArg0: true},
	{name: "IntrinMmCmpgePs", argLength:2, reg:fp21, asm:"CMPPS", resultInArg0: true},
	{name: "IntrinMmCmpgeSd", argLength:2, reg:fp21, asm:"CMPSD", resultInArg0: true},
	{name: "IntrinMmCmpgeSs", argLength:2, reg:fp21, asm:"CMPSS", resultInArg0: true},
	{name: "IntrinMmCmpgtEpi16", argLength:2, reg:fp21, asm:"PCMPGTW", resultInArg0: true},
	{name: "IntrinMmCmpgtEpi8", argLength:2, reg:fp21, asm:"PCMPGTB", resultInArg0: true},
	{name: "IntrinMmCmpgtPd", argLength:2, reg:fp21, asm:"CMPPD", resultInArg0: true},
	{name: "IntrinMmCmpgtPs", argLength:2, reg:fp21, asm:"CMPPS", resultInArg0: true},
	{name: "IntrinMmCmpgtSd", argLength:2, reg:fp21, asm:"CMPSD", resultInArg0: true},
	{name: "IntrinMmCmpgtSs", argLength:2, reg:fp21, asm:"CMPSS", resultInArg0: true},
	{name: "IntrinMmCmplePd", argLength:2, reg:fp21, asm:"CMPPD", resultInArg0: true},
	{name: "IntrinMmCmplePs", argLength:2, reg:fp21, asm:"CMPPS", resultInArg0: true},
	{name: "IntrinMmCmpleSd", argLength:2, reg:fp21, asm:"CMPSD", resultInArg0: true},
	{name: "IntrinMmCmpleSs", argLength:2, reg:fp21, asm:"CMPSS", resultInArg0: true},
	{name: "IntrinMmCmpltEpi16", argLength:2, reg:fp21, asm:"PCMPGTW", resultInArg0: true},
	{name: "IntrinMmCmpltEpi8", argLength:2, reg:fp21, asm:"PCMPGTB", resultInArg0: true},
	{name: "IntrinMmCmpltPd", argLength:2, reg:fp21, asm:"CMPPD", resultInArg0: true},
	{name: "IntrinMmCmpltPs", argLength:2, reg:fp21, asm:"CMPPS", resultInArg0: true},
	{name: "IntrinMmCmpltSd", argLength:2, reg:fp21, asm:"CMPSD", resultInArg0: true},
	{name: "IntrinMmCmpltSs", argLength:2, reg:fp21, asm:"CMPSS", resultInArg0: true},
	{name: "IntrinMmCmpneqPd", argLength:2, reg:fp21, asm:"CMPPD", resultInArg0: true},
	{name: "IntrinMmCmpneqPs", argLength:2, reg:fp21, asm:"CMPPS", resultInArg0: true},
	{name: "IntrinMmCmpneqSd", argLength:2, reg:fp21, asm:"CMPSD", resultInArg0: true},
	{name: "IntrinMmCmpneqSs", argLength:2, reg:fp21, asm:"CMPSS", resultInArg0: true},
	{name: "IntrinMmCmpngePd", argLength:2, reg:fp21, asm:"CMPPD", resultInArg0: true},
	{name: "IntrinMmCmpngePs", argLength:2, reg:fp21, asm:"CMPPS", resultInArg0: true},
	{name: "IntrinMmCmpngeSd", argLength:2, reg:fp21, asm:"CMPSD", resultInArg0: true},
	{name: "IntrinMmCmpngeSs", argLength:2, reg:fp21, asm:"CMPSS", resultInArg0: true},
	{name: "IntrinMmCmpngtPd", argLength:2, reg:fp21, asm:"CMPPD", resultInArg0: true},
	{name: "IntrinMmCmpngtPs", argLength:2, reg:fp21, asm:"CMPPS", resultInArg0: true},
	{name: "IntrinMmCmpngtSd", argLength:2, reg:fp21, asm:"CMPSD", resultInArg0: true},
	{name: "IntrinMmCmpngtSs", argLength:2, reg:fp21, asm:"CMPSS", resultInArg0: true},
	{name: "IntrinMmCmpnlePd", argLength:2, reg:fp21, asm:"CMPPD", resultInArg0: true},
	{name: "IntrinMmCmpnlePs", argLength:2, reg:fp21, asm:"CMPPS", resultInArg0: true},
	{name: "IntrinMmCmpnleSd", argLength:2, reg:fp21, asm:"CMPSD", resultInArg0: true},
	{name: "IntrinMmCmpnleSs", argLength:2, reg:fp21, asm:"CMPSS", resultInArg0: true},
	{name: "IntrinMmCmpnltPd", argLength:2, reg:fp21, asm:"CMPPD", resultInArg0: true},
	{name: "IntrinMmCmpnltPs", argLength:2, reg:fp21, asm:"CMPPS", resultInArg0: true},
	{name: "IntrinMmCmpnltSd", argLength:2, reg:fp21, asm:"CMPSD", resultInArg0: true},
	{name: "IntrinMmCmpnltSs", argLength:2, reg:fp21, asm:"CMPSS", resultInArg0: true},
	{name: "IntrinMmCmpordPd", argLength:2, reg:fp21, asm:"CMPPD", resultInArg0: true},
	{name: "IntrinMmCmpordPs", argLength:2, reg:fp21, asm:"CMPPS", resultInArg0: true},
	{name: "IntrinMmCmpordSd", argLength:2, reg:fp21, asm:"CMPSD", resultInArg0: true},
	{name: "IntrinMmCmpordSs", argLength:2, reg:fp21, asm:"CMPSS", resultInArg0: true},
	{name: "IntrinMmCmpunordPd", argLength:2, reg:fp21, asm:"CMPPD", resultInArg0: true},
	{name: "IntrinMmCmpunordPs", argLength:2, reg:fp21, asm:"CMPPS", resultInArg0: true},
	{name: "IntrinMmCmpunordSd", argLength:2, reg:fp21, asm:"CMPSD", resultInArg0: true},
	{name: "IntrinMmCmpunordSs", argLength:2, reg:fp21, asm:"CMPSS", resultInArg0: true},
	{name: "IntrinMmCvtepi16Epi32", argLength:1, reg:fp11, asm:"PMOVSXWD", resultInArg0: true},
	{name: "IntrinMmCvtepi16Epi64", argLength:1, reg:fp11, asm:"PMOVSXWQ", resultInArg0: true},
	{name: "IntrinMmCvtepi32Epi64", argLength:1, reg:fp11, asm:"PMOVSXDQ", resultInArg0: true},
	{name: "IntrinMmCvtepi8Epi16", argLength:1, reg:fp11, asm:"PMOVSXBW", resultInArg0: true},
	{name: "IntrinMmCvtepi8Epi32", argLength:1, reg:fp11, asm:"PMOVSXBD", resultInArg0: true},
	{name: "IntrinMmCvtepi8Epi64", argLength:1, reg:fp11, asm:"PMOVSXBQ", resultInArg0: true},
	{name: "IntrinMmCvtepu16Epi32", argLength:1, reg:fp11, asm:"PMOVZXWD", resultInArg0: true},
	{name: "IntrinMmCvtepu16Epi64", argLength:1, reg:fp11, asm:"PMOVZXWQ", resultInArg0: true},
	{name: "IntrinMmCvtepu32Epi64", argLength:1, reg:fp11, asm:"PMOVZXDQ", resultInArg0: true},
	{name: "IntrinMmCvtepu8Epi16", argLength:1, reg:fp11, asm:"PMOVZXBW", resultInArg0: true},
	{name: "IntrinMmCvtepu8Epi32", argLength:1, reg:fp11, asm:"PMOVZXBD", resultInArg0: true},
	{name: "IntrinMmCvtepu8Epi64", argLength:1, reg:fp11, asm:"PMOVZXBQ", resultInArg0: true},
	{name: "IntrinMmCvtpdPs", argLength:1, reg:fp11, asm:"CVTPD2PS", resultInArg0: true},
	{name: "IntrinMmCvtpsPd", argLength:1, reg:fp11, asm:"CVTPS2PD", resultInArg0: true},
	{name: "IntrinMmCvtsdSs", argLength:2, reg:fp21, asm:"CVTSD2SS", resultInArg0: true},
	{name: "IntrinMmCvtssSd", argLength:2, reg:fp21, asm:"CVTSS2SD", resultInArg0: true},
	{name: "IntrinMmDivPd", argLength:2, reg:fp21, asm:"DIVPD", resultInArg0: true},
	{name: "IntrinMmDivPs", argLength:2, reg:fp21, asm:"DIVPS", resultInArg0: true},
	{name: "IntrinMmDivSd", argLength:2, reg:fp21, asm:"DIVSD", resultInArg0: true},
	{name: "IntrinMmDivSs", argLength:2, reg:fp21, asm:"DIVSS", resultInArg0: true},
	{name: "IntrinMmFloorPd", argLength:1, reg:fp11, asm:"ROUNDPD", resultInArg0: true},
	{name: "IntrinMmFloorPs", argLength:1, reg:fp11, asm:"ROUNDPS", resultInArg0: true},
	{name: "IntrinMmFloorSd", argLength:2, reg:fp21, asm:"ROUNDSD", resultInArg0: true},
	{name: "IntrinMmFloorSs", argLength:2, reg:fp21, asm:"ROUNDSS", resultInArg0: true},
	{name: "IntrinMmHaddEpi16", argLength:2, reg:fp21, asm:"PHADDW", resultInArg0: true},
	{name: "IntrinMmHaddEpi32", argLength:2, reg:fp21, asm:"PHADDD", resultInArg0: true},
	{name: "IntrinMmHaddPd", argLength:2, reg:fp21, asm:"HADDPD", resultInArg0: true},
	{name: "IntrinMmHaddPs", argLength:2, reg:fp21, asm:"HADDPS", resultInArg0: true},
	{name: "IntrinMmHaddsEpi16", argLength:2, reg:fp21, asm:"PHADDSW", resultInArg0: true},
	{name: "IntrinMmHsubEpi16", argLength:2, reg:fp21, asm:"PHSUBW", resultInArg0: true},
	{name: "IntrinMmHsubEpi32", argLength:2, reg:fp21, asm:"PHSUBD", resultInArg0: true},
	{name: "IntrinMmHsubPd", argLength:2, reg:fp21, asm:"HSUBPD", resultInArg0: true},
	{name: "IntrinMmHsubPs", argLength:2, reg:fp21, asm:"HSUBPS", resultInArg0: true},
	{name: "IntrinMmHsubsEpi16", argLength:2, reg:fp21, asm:"PHSUBSW", resultInArg0: true},
	{name: "IntrinMmMaxEpi16", argLength:2, reg:fp21, asm:"PMAXSW", resultInArg0: true},
	{name: "IntrinMmMaxEpu8", argLength:2, reg:fp21, asm:"PMAXUB", resultInArg0: true},
	{name: "IntrinMmMaxPd", argLength:2, reg:fp21, asm:"MAXPD", resultInArg0: true},
	{name: "IntrinMmMaxPs", argLength:2, reg:fp21, asm:"MAXPS", resultInArg0: true},
	{name: "IntrinMmMaxSd", argLength:2, reg:fp21, asm:"MAXSD", resultInArg0: true},
	{name: "IntrinMmMaxSs", argLength:2, reg:fp21, asm:"MAXSS", resultInArg0: true},
	{name: "IntrinMmMinEpi16", argLength:2, reg:fp21, asm:"PMINSW", resultInArg0: true},
	{name: "IntrinMmMinEpu8", argLength:2, reg:fp21, asm:"PMINUB", resultInArg0: true},
	{name: "IntrinMmMinPd", argLength:2, reg:fp21, asm:"MINPD", resultInArg0: true},
	{name: "IntrinMmMinPs", argLength:2, reg:fp21, asm:"MINPS", resultInArg0: true},
	{name: "IntrinMmMinSd", argLength:2, reg:fp21, asm:"MINSD", resultInArg0: true},
	{name: "IntrinMmMinSs", argLength:2, reg:fp21, asm:"MINSS", resultInArg0: true},
	{name: "IntrinMmMinposEpu16", argLength:1, reg:fp11, asm:"PHMINPOSUW", resultInArg0: true},
	{name: "IntrinMmMoveEpi64", argLength:1, reg:fp11, asm:"MOVQ", resultInArg0: true},
	{name: "IntrinMmMoveSd", argLength:2, reg:fp21, asm:"MOVSD", resultInArg0: true},
	{name: "IntrinMmMoveSs", argLength:2, reg:fp21, asm:"MOVSS", resultInArg0: true},
	{name: "IntrinMmMovedupPd", argLength:1, reg:fp11, asm:"MOVDDUP", resultInArg0: true},
	{name: "IntrinMm256MovedupPd", argLength:1, reg:avx11, asm:"VMOVDDUP", resultInArg0: true},
	{name: "IntrinMmMovehdupPs", argLength:1, reg:fp11, asm:"MOVSHDUP", resultInArg0: true},
	{name: "IntrinMm256MovehdupPs", argLength:1, reg:avx11, asm:"VMOVSHDUP", resultInArg0: true},
	{name: "IntrinMmMovehlPs", argLength:2, reg:fp21, asm:"MOVHLPS", resultInArg0: true},
	{name: "IntrinMmMoveldupPs", argLength:1, reg:fp11, asm:"MOVSLDUP", resultInArg0: true},
	{name: "IntrinMm256MoveldupPs", argLength:1, reg:avx11, asm:"VMOVSLDUP", resultInArg0: true},
	{name: "IntrinMmMovelhPs", argLength:2, reg:fp21, asm:"MOVLHPS", resultInArg0: true},
	{name: "IntrinMmMulEpi32", argLength:2, reg:fp21, asm:"PMULDQ", resultInArg0: true},
	{name: "IntrinMmMulPd", argLength:2, reg:fp21, asm:"MULPD", resultInArg0: true},
	{name: "IntrinMmMulPs", argLength:2, reg:fp21, asm:"MULPS", resultInArg0: true},
	{name: "IntrinMmMulSd", argLength:2, reg:fp21, asm:"MULSD", resultInArg0: true},
	{name: "IntrinMmMulSs", argLength:2, reg:fp21, asm:"MULSS", resultInArg0: true},
	{name: "IntrinMmMulhiEpi16", argLength:2, reg:fp21, asm:"PMULHW", resultInArg0: true},
	{name: "IntrinMmMulhiEpu16", argLength:2, reg:fp21, asm:"PMULHUW", resultInArg0: true},
	{name: "IntrinMmMulloEpi16", argLength:2, reg:fp21, asm:"PMULLW", resultInArg0: true},
	{name: "IntrinMmMulloEpi32", argLength:2, reg:fp21, asm:"PMULLD", resultInArg0: true},
	{name: "IntrinMmOrPd", argLength:2, reg:fp21, asm:"ORPD", resultInArg0: true},
	{name: "IntrinMmOrPs", argLength:2, reg:fp21, asm:"ORPS", resultInArg0: true},
	{name: "IntrinMmOrSi128", argLength:2, reg:fp21, asm:"POR", resultInArg0: true},
	{name: "IntrinMm256OrSi256", argLength:2, reg:avx21, asm:"VPOR"},
	{name: "IntrinMmPacksEpi16", argLength:2, reg:fp21, asm:"PACKSSWB", resultInArg0: true},
	{name: "IntrinMmPackusEpi16", argLength:2, reg:fp21, asm:"PACKUSWB", resultInArg0: true},
	{name: "IntrinMmRcpPs", argLength:1, reg:fp11, asm:"RCPPS", resultInArg0: true},
	{name: "IntrinMmRcpSs", argLength:1, reg:fp11, asm:"RCPSS", resultInArg0: true},
	{name: "IntrinMmRsqrtPs", argLength:1, reg:fp11, asm:"RSQRTPS", resultInArg0: true},
	{name: "IntrinMmRsqrtSs", argLength:1, reg:fp11, asm:"RSQRTSS", resultInArg0: true},
	{name: "IntrinMmSadEpu8", argLength:2, reg:fp21, asm:"PSADBW", resultInArg0: true},
	{name: "IntrinMmShuffleEpi8", argLength:2, reg:fp21, asm:"PSHUFB", resultInArg0: true},
	{name: "IntrinMm256ShuffleEpi8", argLength:2, reg:avx21, asm:"VPSHUFB"},
	{name: "IntrinMmSllEpi16", argLength:2, reg:fp21, asm:"PSLLW", resultInArg0: true},
	{name: "IntrinMmSllEpi64", argLength:2, reg:fp21, asm:"PSLLQ", resultInArg0: true},
	{name: "IntrinMmSqrtPd", argLength:1, reg:fp11, asm:"SQRTPD", resultInArg0: true},
	{name: "IntrinMmSqrtPs", argLength:1, reg:fp11, asm:"SQRTPS", resultInArg0: true},
	{name: "IntrinMmSqrtSd", argLength:2, reg:fp21, asm:"SQRTSD", resultInArg0: true},
	{name: "IntrinMmSqrtSs", argLength:1, reg:fp11, asm:"SQRTSS", resultInArg0: true},
	{name: "IntrinMmSraEpi16", argLength:2, reg:fp21, asm:"PSRAW", resultInArg0: true},
	{name: "IntrinMmSrlEpi16", argLength:2, reg:fp21, asm:"PSRLW", resultInArg0: true},
	{name: "IntrinMmSrlEpi64", argLength:2, reg:fp21, asm:"PSRLQ", resultInArg0: true},
	{name: "IntrinMmSubEpi16", argLength:2, reg:fp21, asm:"PSUBW", resultInArg0: true},
	{name: "IntrinMmSubEpi64", argLength:2, reg:fp21, asm:"PSUBQ", resultInArg0: true},
	{name: "IntrinMmSubEpi8", argLength:2, reg:fp21, asm:"PSUBB", resultInArg0: true},
	{name: "IntrinMmSubPd", argLength:2, reg:fp21, asm:"SUBPD", resultInArg0: true},
	{name: "IntrinMmSubPs", argLength:2, reg:fp21, asm:"SUBPS", resultInArg0: true},
	{name: "IntrinMmSubSd", argLength:2, reg:fp21, asm:"SUBSD", resultInArg0: true},
	{name: "IntrinMmSubSs", argLength:2, reg:fp21, asm:"SUBSS", resultInArg0: true},
	{name: "IntrinMmSubsEpi16", argLength:2, reg:fp21, asm:"PSUBSW", resultInArg0: true},
	{name: "IntrinMmSubsEpi8", argLength:2, reg:fp21, asm:"PSUBSB", resultInArg0: true},
	{name: "IntrinMmSubsEpu16", argLength:2, reg:fp21, asm:"PSUBUSW", resultInArg0: true},
	{name: "IntrinMmSubsEpu8", argLength:2, reg:fp21, asm:"PSUBUSB", resultInArg0: true},
	{name: "IntrinMmUnpackhiEpi64", argLength:2, reg:fp21, asm:"PUNPCKHQDQ", resultInArg0: true},
	{name: "IntrinMmUnpackhiEpi8", argLength:2, reg:fp21, asm:"PUNPCKHBW", resultInArg0: true},
	{name: "IntrinMmUnpackhiPd", argLength:2, reg:fp21, asm:"UNPCKHPD", resultInArg0: true},
	{name: "IntrinMmUnpackhiPs", argLength:2, reg:fp21, asm:"UNPCKHPS", resultInArg0: true},
	{name: "IntrinMmUnpackloEpi64", argLength:2, reg:fp21, asm:"PUNPCKLQDQ", resultInArg0: true},
	{name: "IntrinMmUnpackloEpi8", argLength:2, reg:fp21, asm:"PUNPCKLBW", resultInArg0: true},
	{name: "IntrinMmUnpackloPd", argLength:2, reg:fp21, asm:"UNPCKLPD", resultInArg0: true},
	{name: "IntrinMmUnpackloPs", argLength:2, reg:fp21, asm:"UNPCKLPS", resultInArg0: true},
	{name: "IntrinMmXorPd", argLength:2, reg:fp21, asm:"XORPD", resultInArg0: true},
	{name: "IntrinMmXorPs", argLength:2, reg:fp21, asm:"XORPS", resultInArg0: true},
	{name: "IntrinMmXorSi128", argLength:2, reg:fp21, asm:"PXOR", resultInArg0: true},
	{name: "IntrinMm256XorSi256", argLength:2, reg:avx21, asm:"VPXOR"},
