// Code generated by "stringer -type=FeatureID,Vendor"; DO NOT EDIT.

package cpuid

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ADX-1]
	_ = x[AESNI-2]
	_ = x[AMD3DNOW-3]
	_ = x[AMD3DNOWEXT-4]
	_ = x[AMXBF16-5]
	_ = x[AMXINT8-6]
	_ = x[AMXTILE-7]
	_ = x[AVX-8]
	_ = x[AVX2-9]
	_ = x[AVX512BF16-10]
	_ = x[AVX512BITALG-11]
	_ = x[AVX512BW-12]
	_ = x[AVX512CD-13]
	_ = x[AVX512DQ-14]
	_ = x[AVX512ER-15]
	_ = x[AVX512F-16]
	_ = x[AVX512FP16-17]
	_ = x[AVX512IFMA-18]
	_ = x[AVX512PF-19]
	_ = x[AVX512VBMI-20]
	_ = x[AVX512VBMI2-21]
	_ = x[AVX512VL-22]
	_ = x[AVX512VNNI-23]
	_ = x[AVX512VP2INTERSECT-24]
	_ = x[AVX512VPOPCNTDQ-25]
	_ = x[AVXSLOW-26]
	_ = x[AVXVNNI-27]
	_ = x[BMI1-28]
	_ = x[BMI2-29]
	_ = x[CETIBT-30]
	_ = x[CETSS-31]
	_ = x[CLDEMOTE-32]
	_ = x[CLMUL-33]
	_ = x[CLZERO-34]
	_ = x[CMOV-35]
	_ = x[CMPSB_SCADBS_SHORT-36]
	_ = x[CMPXCHG8-37]
	_ = x[CPBOOST-38]
	_ = x[CX16-39]
	_ = x[ENQCMD-40]
	_ = x[ERMS-41]
	_ = x[F16C-42]
	_ = x[FLUSH_L1D-43]
	_ = x[FMA3-44]
	_ = x[FMA4-45]
	_ = x[FSRM-46]
	_ = x[FXSR-47]
	_ = x[FXSROPT-48]
	_ = x[GFNI-49]
	_ = x[HLE-50]
	_ = x[HRESET-51]
	_ = x[HTT-52]
	_ = x[HWA-53]
	_ = x[HYBRID_CPU-54]
	_ = x[HYPERVISOR-55]
	_ = x[IA32_ARCH_CAP-56]
	_ = x[IA32_CORE_CAP-57]
	_ = x[IBPB-58]
	_ = x[IBS-59]
	_ = x[IBSBRNTRGT-60]
	_ = x[IBSFETCHSAM-61]
	_ = x[IBSFFV-62]
	_ = x[IBSOPCNT-63]
	_ = x[IBSOPCNTEXT-64]
	_ = x[IBSOPSAM-65]
	_ = x[IBSRDWROPCNT-66]
	_ = x[IBSRIPINVALIDCHK-67]
	_ = x[IBS_PREVENTHOST-68]
	_ = x[INT_WBINVD-69]
	_ = x[INVLPGB-70]
	_ = x[LAHF-71]
	_ = x[LAM-72]
	_ = x[LBRVIRT-73]
	_ = x[LZCNT-74]
	_ = x[MCAOVERFLOW-75]
	_ = x[MCDT_NO-76]
	_ = x[MCOMMIT-77]
	_ = x[MD_CLEAR-78]
	_ = x[MMX-79]
	_ = x[MMXEXT-80]
	_ = x[MOVBE-81]
	_ = x[MOVDIR64B-82]
	_ = x[MOVDIRI-83]
	_ = x[MOVSB_ZL-84]
	_ = x[MPX-85]
	_ = x[MSRIRC-86]
	_ = x[MSR_PAGEFLUSH-87]
	_ = x[NRIPS-88]
	_ = x[NX-89]
	_ = x[OSXSAVE-90]
	_ = x[PCONFIG-91]
	_ = x[POPCNT-92]
	_ = x[RDPRU-93]
	_ = x[RDRAND-94]
	_ = x[RDSEED-95]
	_ = x[RDTSCP-96]
	_ = x[RTM-97]
	_ = x[RTM_ALWAYS_ABORT-98]
	_ = x[SERIALIZE-99]
	_ = x[SEV-100]
	_ = x[SEV_64BIT-101]
	_ = x[SEV_ALTERNATIVE-102]
	_ = x[SEV_DEBUGSWAP-103]
	_ = x[SEV_ES-104]
	_ = x[SEV_RESTRICTED-105]
	_ = x[SEV_SNP-106]
	_ = x[SGX-107]
	_ = x[SGXLC-108]
	_ = x[SHA-109]
	_ = x[SME-110]
	_ = x[SME_COHERENT-111]
	_ = x[SPEC_CTRL_SSBD-112]
	_ = x[SRBDS_CTRL-113]
	_ = x[SSE-114]
	_ = x[SSE2-115]
	_ = x[SSE3-116]
	_ = x[SSE4-117]
	_ = x[SSE42-118]
	_ = x[SSE4A-119]
	_ = x[SSSE3-120]
	_ = x[STIBP-121]
	_ = x[STOSB_SHORT-122]
	_ = x[SUCCOR-123]
	_ = x[SVM-124]
	_ = x[SVMDA-125]
	_ = x[SVMFBASID-126]
	_ = x[SVML-127]
	_ = x[SVMNP-128]
	_ = x[SVMPF-129]
	_ = x[SVMPFT-130]
	_ = x[SYSCALL-131]
	_ = x[SYSEE-132]
	_ = x[TBM-133]
	_ = x[TME-134]
	_ = x[TOPEXT-135]
	_ = x[TSCRATEMSR-136]
	_ = x[TSXLDTRK-137]
	_ = x[VAES-138]
	_ = x[VMCBCLEAN-139]
	_ = x[VMPL-140]
	_ = x[VMSA_REGPROT-141]
	_ = x[VMX-142]
	_ = x[VPCLMULQDQ-143]
	_ = x[VTE-144]
	_ = x[WAITPKG-145]
	_ = x[WBNOINVD-146]
	_ = x[X87-147]
	_ = x[XGETBV1-148]
	_ = x[XOP-149]
	_ = x[XSAVE-150]
	_ = x[XSAVEC-151]
	_ = x[XSAVEOPT-152]
	_ = x[XSAVES-153]
	_ = x[AESARM-154]
	_ = x[ARMCPUID-155]
	_ = x[ASIMD-156]
	_ = x[ASIMDDP-157]
	_ = x[ASIMDHP-158]
	_ = x[ASIMDRDM-159]
	_ = x[ATOMICS-160]
	_ = x[CRC32-161]
	_ = x[DCPOP-162]
	_ = x[EVTSTRM-163]
	_ = x[FCMA-164]
	_ = x[FP-165]
	_ = x[FPHP-166]
	_ = x[GPA-167]
	_ = x[JSCVT-168]
	_ = x[LRCPC-169]
	_ = x[PMULL-170]
	_ = x[SHA1-171]
	_ = x[SHA2-172]
	_ = x[SHA3-173]
	_ = x[SHA512-174]
	_ = x[SM3-175]
	_ = x[SM4-176]
	_ = x[SVE-177]
	_ = x[lastID-178]
	_ = x[firstID-0]
}

const _FeatureID_name = "firstIDADXAESNIAMD3DNOWAMD3DNOWEXTAMXBF16AMXINT8AMXTILEAVXAVX2AVX512BF16AVX512BITALGAVX512BWAVX512CDAVX512DQAVX512ERAVX512FAVX512FP16AVX512IFMAAVX512PFAVX512VBMIAVX512VBMI2AVX512VLAVX512VNNIAVX512VP2INTERSECTAVX512VPOPCNTDQAVXSLOWAVXVNNIBMI1BMI2CETIBTCETSSCLDEMOTECLMULCLZEROCMOVCMPSB_SCADBS_SHORTCMPXCHG8CPBOOSTCX16ENQCMDERMSF16CFLUSH_L1DFMA3FMA4FSRMFXSRFXSROPTGFNIHLEHRESETHTTHWAHYBRID_CPUHYPERVISORIA32_ARCH_CAPIA32_CORE_CAPIBPBIBSIBSBRNTRGTIBSFETCHSAMIBSFFVIBSOPCNTIBSOPCNTEXTIBSOPSAMIBSRDWROPCNTIBSRIPINVALIDCHKIBS_PREVENTHOSTINT_WBINVDINVLPGBLAHFLAMLBRVIRTLZCNTMCAOVERFLOWMCDT_NOMCOMMITMD_CLEARMMXMMXEXTMOVBEMOVDIR64BMOVDIRIMOVSB_ZLMPXMSRIRCMSR_PAGEFLUSHNRIPSNXOSXSAVEPCONFIGPOPCNTRDPRURDRANDRDSEEDRDTSCPRTMRTM_ALWAYS_ABORTSERIALIZESEVSEV_64BITSEV_ALTERNATIVESEV_DEBUGSWAPSEV_ESSEV_RESTRICTEDSEV_SNPSGXSGXLCSHASMESME_COHERENTSPEC_CTRL_SSBDSRBDS_CTRLSSESSE2SSE3SSE4SSE42SSE4ASSSE3STIBPSTOSB_SHORTSUCCORSVMSVMDASVMFBASIDSVMLSVMNPSVMPFSVMPFTSYSCALLSYSEETBMTMETOPEXTTSCRATEMSRTSXLDTRKVAESVMCBCLEANVMPLVMSA_REGPROTVMXVPCLMULQDQVTEWAITPKGWBNOINVDX87XGETBV1XOPXSAVEXSAVECXSAVEOPTXSAVESAESARMARMCPUIDASIMDASIMDDPASIMDHPASIMDRDMATOMICSCRC32DCPOPEVTSTRMFCMAFPFPHPGPAJSCVTLRCPCPMULLSHA1SHA2SHA3SHA512SM3SM4SVElastID"

var _FeatureID_index = [...]uint16{0, 7, 10, 15, 23, 34, 41, 48, 55, 58, 62, 72, 84, 92, 100, 108, 116, 123, 133, 143, 151, 161, 172, 180, 190, 208, 223, 230, 237, 241, 245, 251, 256, 264, 269, 275, 279, 297, 305, 312, 316, 322, 326, 330, 339, 343, 347, 351, 355, 362, 366, 369, 375, 378, 381, 391, 401, 414, 427, 431, 434, 444, 455, 461, 469, 480, 488, 500, 516, 531, 541, 548, 552, 555, 562, 567, 578, 585, 592, 600, 603, 609, 614, 623, 630, 638, 641, 647, 660, 665, 667, 674, 681, 687, 692, 698, 704, 710, 713, 729, 738, 741, 750, 765, 778, 784, 798, 805, 808, 813, 816, 819, 831, 845, 855, 858, 862, 866, 870, 875, 880, 885, 890, 901, 907, 910, 915, 924, 928, 933, 938, 944, 951, 956, 959, 962, 968, 978, 986, 990, 999, 1003, 1015, 1018, 1028, 1031, 1038, 1046, 1049, 1056, 1059, 1064, 1070, 1078, 1084, 1090, 1098, 1103, 1110, 1117, 1125, 1132, 1137, 1142, 1149, 1153, 1155, 1159, 1162, 1167, 1172, 1177, 1181, 1185, 1189, 1195, 1198, 1201, 1204, 1210}

func (i FeatureID) String() string {
	if i < 0 || i >= FeatureID(len(_FeatureID_index)-1) {
		return "FeatureID(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FeatureID_name[_FeatureID_index[i]:_FeatureID_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[VendorUnknown-0]
	_ = x[Intel-1]
	_ = x[AMD-2]
	_ = x[VIA-3]
	_ = x[Transmeta-4]
	_ = x[NSC-5]
	_ = x[KVM-6]
	_ = x[MSVM-7]
	_ = x[VMware-8]
	_ = x[XenHVM-9]
	_ = x[Bhyve-10]
	_ = x[Hygon-11]
	_ = x[SiS-12]
	_ = x[RDC-13]
	_ = x[Ampere-14]
	_ = x[ARM-15]
	_ = x[Broadcom-16]
	_ = x[Cavium-17]
	_ = x[DEC-18]
	_ = x[Fujitsu-19]
	_ = x[Infineon-20]
	_ = x[Motorola-21]
	_ = x[NVIDIA-22]
	_ = x[AMCC-23]
	_ = x[Qualcomm-24]
	_ = x[Marvell-25]
	_ = x[lastVendor-26]
}

const _Vendor_name = "VendorUnknownIntelAMDVIATransmetaNSCKVMMSVMVMwareXenHVMBhyveHygonSiSRDCAmpereARMBroadcomCaviumDECFujitsuInfineonMotorolaNVIDIAAMCCQualcommMarvelllastVendor"

var _Vendor_index = [...]uint8{0, 13, 18, 21, 24, 33, 36, 39, 43, 49, 55, 60, 65, 68, 71, 77, 80, 88, 94, 97, 104, 112, 120, 126, 130, 138, 145, 155}

func (i Vendor) String() string {
	if i < 0 || i >= Vendor(len(_Vendor_index)-1) {
		return "Vendor(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Vendor_name[_Vendor_index[i]:_Vendor_index[i+1]]
}
