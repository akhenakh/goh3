// Code generated for linux/arm64 by 'ccgo -Isrc/h3lib/include -I../src/h3lib/include ../src/h3lib/lib/algos.c ../src/h3lib/lib/baseCells.c ../src/h3lib/lib/bbox.c ../src/h3lib/lib/coordijk.c ../src/h3lib/lib/directedEdge.c ../src/h3lib/lib/faceijk.c ../src/h3lib/lib/h3Assert.c ../src/h3lib/lib/h3Index.c ../src/h3lib/lib/iterators.c ../src/h3lib/lib/latLng.c ../src/h3lib/lib/linkedGeo.c ../src/h3lib/lib/localij.c ../src/h3lib/lib/mathExtensions.c ../src/h3lib/lib/polyfill.c ../src/h3lib/lib/polygon.c ../src/h3lib/lib/vec2d.c ../src/h3lib/lib/vec3d.c ../src/h3lib/lib/vertex.c ../src/h3lib/lib/vertexGraph.c --package-name=ch3 --prefix-enumerator=_ --prefix-external=X --prefix-field=F --prefix-macro=D --prefix-static-internal=_ --prefix-static-none=_ --prefix-tagged-enum=_ --prefix-tagged-struct=T --prefix-tagged-union=T --prefix-typename=T --prefix-undefined=_ -ignore-unsupported-alignment', DO NOT EDIT.

//go:build linux && arm64

package ch3

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

const DBUFSIZ = 1024
const DDBL_DECIMAL_DIG = 17
const DDBL_DIG = 15
const DDBL_EPSILON = 2.22044604925031308085e-16
const DDBL_HAS_SUBNORM = 1
const DDBL_MANT_DIG = 53
const DDBL_MAX = 1.79769313486231570815e+308
const DDBL_MAX_10_EXP = 308
const DDBL_MAX_EXP = 1024
const DDBL_MIN = 2.22507385850720138309e-308
const DDBL_TRUE_MIN = 4.94065645841246544177e-324
const DDECIMAL_DIG = 17
const DEARTH_RADIUS_KM = 6371.007180918475
const DEPSILON = 0.0000000000000001
const DEPSILON_DEG = .000000001
const DEXIT_FAILURE = 1
const DEXIT_SUCCESS = 0
const DFILENAME_MAX = 4096
const DFLT_DECIMAL_DIG = 9
const DFLT_DIG = 6
const DFLT_EPSILON = 1.1920928955078125e-07
const DFLT_EVAL_METHOD = 0
const DFLT_HAS_SUBNORM = 1
const DFLT_MANT_DIG = 24
const DFLT_MAX = 3.40282346638528859812e+38
const DFLT_MAX_10_EXP = 38
const DFLT_MAX_EXP = 128
const DFLT_MIN = 1.17549435082228750797e-38
const DFLT_RADIX = 2
const DFLT_TRUE_MIN = 1.40129846432481707092e-45
const DFOPEN_MAX = 1000
const DFP_FAST_FMA = 1
const DFP_FAST_FMAF = 1
const DFP_ILOGB0 = "FP_ILOGBNAN"
const DFP_INFINITE = 1
const DFP_NAN = 0
const DFP_NORMAL = 4
const DFP_SUBNORMAL = 3
const DFP_ZERO = 2
const DH3_BC_OFFSET = 45
const DH3_CELL_MODE = 1
const DH3_DIRECTEDEDGE_MODE = 2
const DH3_EDGE_MODE = 3
const DH3_MAX_OFFSET = 63
const DH3_MODE_OFFSET = 59
const DH3_NULL = 0
const DH3_NUM_BITS = 64
const DH3_PER_DIGIT_OFFSET = 3
const DH3_RESERVED_OFFSET = 56
const DH3_RES_OFFSET = 52
const DH3_VERSION_MAJOR = 4
const DH3_VERSION_MINOR = 2
const DH3_VERSION_PATCH = 1
const DH3_VERTEX_MODE = 4
const DHUGE = 3.40282346638528859812e+38
const DHUGE_VALF = "INFINITY"
const DIJ = 1
const DINT16_MAX = 0x7fff
const DINT32_MAX = 0x7fffffff
const DINT64_MAX = 0x7fffffffffffffff
const DINT8_MAX = 0x7f
const DINTMAX_MAX = "INT64_MAX"
const DINTMAX_MIN = "INT64_MIN"
const DINTPTR_MAX = "INT64_MAX"
const DINTPTR_MIN = "INT64_MIN"
const DINT_FAST16_MAX = "INT32_MAX"
const DINT_FAST16_MIN = "INT32_MIN"
const DINT_FAST32_MAX = "INT32_MAX"
const DINT_FAST32_MIN = "INT32_MIN"
const DINT_FAST64_MAX = "INT64_MAX"
const DINT_FAST64_MIN = "INT64_MIN"
const DINT_FAST8_MAX = "INT8_MAX"
const DINT_FAST8_MIN = "INT8_MIN"
const DINT_LEAST16_MAX = "INT16_MAX"
const DINT_LEAST16_MIN = "INT16_MIN"
const DINT_LEAST32_MAX = "INT32_MAX"
const DINT_LEAST32_MIN = "INT32_MIN"
const DINT_LEAST64_MAX = "INT64_MAX"
const DINT_LEAST64_MIN = "INT64_MIN"
const DINT_LEAST8_MAX = "INT8_MAX"
const DINT_LEAST8_MIN = "INT8_MIN"
const DINVALID_BASE_CELL = 127
const DINV_RES0_U_GNOMONIC = 2.61803398874989588842
const DJK = 3
const DKI = 2
const DLDBL_DECIMAL_DIG = "DECIMAL_DIG"
const DLDBL_DIG = 15
const DLDBL_EPSILON = 2.22044604925031308085e-16
const DLDBL_HAS_SUBNORM = 1
const DLDBL_MANT_DIG = 53
const DLDBL_MAX = 1.79769313486231570815e+308
const DLDBL_MAX_10_EXP = 308
const DLDBL_MAX_EXP = 1024
const DLDBL_MIN = 2.22507385850720138309e-308
const DLDBL_TRUE_MIN = 4.94065645841246544177e-324
const DL_ctermid = 20
const DL_cuserid = 20
const DL_tmpnam = 20
const DMATH_ERREXCEPT = 2
const DMATH_ERRNO = 1
const DMAX_CELL_BNDRY_VERTS = 10
const DMAX_FACE_COORD = 2
const DMAX_H3_RES = 15
const DMAX_ONE_RING_SIZE = 7
const DM_180_PI = 57.29577951308232087679815481410517033240547
const DM_1_PI = 0.31830988618379067154
const DM_2PI = 6.28318530717958647692528676655900576839433
const DM_2_PI = 0.63661977236758134308
const DM_2_SQRTPI = 1.12837916709551257390
const DM_AP7_ROT_RADS = 0.333473172251832115336090755351601070065900389
const DM_COS_AP7_ROT = 0.9449111825230680680167902
const DM_E = 2.7182818284590452354
const DM_LN10 = 2.30258509299404568402
const DM_LN2 = 0.69314718055994530942
const DM_LOG10E = 0.43429448190325182765
const DM_LOG2E = 1.4426950408889634074
const DM_ONESEVENTH = 0.14285714285714285714285714285714285
const DM_ONETHIRD = 0.333333333333333333333333333333333333333
const DM_PI = 3.14159265358979323846
const DM_PI_180 = 0.0174532925199432957692369076848861271111
const DM_PI_4 = 0.78539816339744830962
const DM_RSIN60 = 1.1547005383792515290182975610039149112953
const DM_SIN60 = "M_SQRT3_2"
const DM_SIN_AP7_ROT = 0.3273268353539885718950318
const DM_SQRT1_2 = 0.70710678118654752440
const DM_SQRT2 = 1.41421356237309504880
const DM_SQRT3_2 = 0.8660254037844386467637231707529361834714
const DNUM_BASE_CELLS = 122
const DNUM_HEX_VERTS = 6
const DNUM_ICOSA_FACES = 20
const DNUM_PENTAGONS = 12
const DNUM_PENT_VERTS = 5
const DPOLYGON_TO_CELLS_BUFFER = 12
const DPTRDIFF_MAX = "INT64_MAX"
const DPTRDIFF_MIN = "INT64_MIN"
const DP_tmpdir = "/tmp"
const DRAND_MAX = 0x7fffffff
const DRES0_U_GNOMONIC = 0.38196601125010500003
const DSIG_ATOMIC_MAX = "INT32_MAX"
const DSIG_ATOMIC_MIN = "INT32_MIN"
const DSIZE_MAX = "UINT64_MAX"
const DTMP_MAX = 10000
const DUINT16_MAX = 0xffff
const DUINT32_MAX = "0xffffffffu"
const DUINT64_MAX = "0xffffffffffffffffu"
const DUINT8_MAX = 0xff
const DUINTMAX_MAX = "UINT64_MAX"
const DUINTPTR_MAX = "UINT64_MAX"
const DUINT_FAST16_MAX = "UINT32_MAX"
const DUINT_FAST32_MAX = "UINT32_MAX"
const DUINT_FAST64_MAX = "UINT64_MAX"
const DUINT_FAST8_MAX = "UINT8_MAX"
const DUINT_LEAST16_MAX = "UINT16_MAX"
const DUINT_LEAST32_MAX = "UINT32_MAX"
const DUINT_LEAST64_MAX = "UINT64_MAX"
const DUINT_LEAST8_MAX = "UINT8_MAX"
const DWINT_MAX = "UINT32_MAX"
const DWINT_MIN = 0
const DWNOHANG = 1
const DWUNTRACED = 2
const D_GNU_SOURCE = 1
const D_IOFBF = 0
const D_IOLBF = 1
const D_IONBF = 2
const D_LP64 = 1
const D_STDC_PREDEF_H = 1
const D__AARCH64EL__ = 1
const D__AARCH64_CMODEL_SMALL__ = 1
const D__ARM_64BIT_STATE = 1
const D__ARM_ALIGN_MAX_PWR = 28
const D__ARM_ALIGN_MAX_STACK_PWR = 16
const D__ARM_ARCH = 8
const D__ARM_ARCH_8A = 1
const D__ARM_ARCH_ISA_A64 = 1
const D__ARM_ARCH_PROFILE = 65
const D__ARM_FEATURE_CLZ = 1
const D__ARM_FEATURE_FMA = 1
const D__ARM_FEATURE_IDIV = 1
const D__ARM_FEATURE_NUMERIC_MAXMIN = 1
const D__ARM_FEATURE_UNALIGNED = 1
const D__ARM_FP = 14
const D__ARM_FP16_ARGS = 1
const D__ARM_FP16_FORMAT_IEEE = 1
const D__ARM_NEON = 1
const D__ARM_PCS_AAPCS64 = 1
const D__ARM_SIZEOF_MINIMAL_ENUM = 4
const D__ARM_SIZEOF_WCHAR_T = 4
const D__ATOMIC_ACQUIRE = 2
const D__ATOMIC_ACQ_REL = 4
const D__ATOMIC_CONSUME = 1
const D__ATOMIC_RELAXED = 0
const D__ATOMIC_RELEASE = 3
const D__ATOMIC_SEQ_CST = 5
const D__BIGGEST_ALIGNMENT__ = 16
const D__BIG_ENDIAN = 4321
const D__BYTE_ORDER = 1234
const D__BYTE_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const D__CCGO__ = 1
const D__CHAR_BIT__ = 8
const D__CHAR_UNSIGNED__ = 1
const D__DBL_DECIMAL_DIG__ = 17
const D__DBL_DIG__ = 15
const D__DBL_HAS_DENORM__ = 1
const D__DBL_HAS_INFINITY__ = 1
const D__DBL_HAS_QUIET_NAN__ = 1
const D__DBL_IS_IEC_60559__ = 2
const D__DBL_MANT_DIG__ = 53
const D__DBL_MAX_10_EXP__ = 308
const D__DBL_MAX_EXP__ = 1024
const D__DECIMAL_DIG__ = 36
const D__DEC_EVAL_METHOD__ = 2
const D__ELF__ = 1
const D__FINITE_MATH_ONLY__ = 0
const D__FLOAT_WORD_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const D__FLT128_DECIMAL_DIG__ = 36
const D__FLT128_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const D__FLT128_DIG__ = 33
const D__FLT128_EPSILON__ = 1.92592994438723585305597794258492732e-34
const D__FLT128_HAS_DENORM__ = 1
const D__FLT128_HAS_INFINITY__ = 1
const D__FLT128_HAS_QUIET_NAN__ = 1
const D__FLT128_IS_IEC_60559__ = 2
const D__FLT128_MANT_DIG__ = 113
const D__FLT128_MAX_10_EXP__ = 4932
const D__FLT128_MAX_EXP__ = 16384
const D__FLT128_MAX__ = "1.18973149535723176508575932662800702e+4932"
const D__FLT128_MIN__ = 3.36210314311209350626267781732175260e-4932
const D__FLT128_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const D__FLT16_DECIMAL_DIG__ = 5
const D__FLT16_DENORM_MIN__ = 5.96046447753906250000000000000000000e-8
const D__FLT16_DIG__ = 3
const D__FLT16_EPSILON__ = 9.76562500000000000000000000000000000e-4
const D__FLT16_HAS_DENORM__ = 1
const D__FLT16_HAS_INFINITY__ = 1
const D__FLT16_HAS_QUIET_NAN__ = 1
const D__FLT16_IS_IEC_60559__ = 2
const D__FLT16_MANT_DIG__ = 11
const D__FLT16_MAX_10_EXP__ = 4
const D__FLT16_MAX_EXP__ = 16
const D__FLT16_MAX__ = 6.55040000000000000000000000000000000e+4
const D__FLT16_MIN__ = 6.10351562500000000000000000000000000e-5
const D__FLT16_NORM_MAX__ = 6.55040000000000000000000000000000000e+4
const D__FLT32X_DECIMAL_DIG__ = 17
const D__FLT32X_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const D__FLT32X_DIG__ = 15
const D__FLT32X_EPSILON__ = 2.22044604925031308084726333618164062e-16
const D__FLT32X_HAS_DENORM__ = 1
const D__FLT32X_HAS_INFINITY__ = 1
const D__FLT32X_HAS_QUIET_NAN__ = 1
const D__FLT32X_IS_IEC_60559__ = 2
const D__FLT32X_MANT_DIG__ = 53
const D__FLT32X_MAX_10_EXP__ = 308
const D__FLT32X_MAX_EXP__ = 1024
const D__FLT32X_MAX__ = 1.79769313486231570814527423731704357e+308
const D__FLT32X_MIN__ = 2.22507385850720138309023271733240406e-308
const D__FLT32X_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const D__FLT32_DECIMAL_DIG__ = 9
const D__FLT32_DENORM_MIN__ = 1.40129846432481707092372958328991613e-45
const D__FLT32_DIG__ = 6
const D__FLT32_EPSILON__ = 1.19209289550781250000000000000000000e-7
const D__FLT32_HAS_DENORM__ = 1
const D__FLT32_HAS_INFINITY__ = 1
const D__FLT32_HAS_QUIET_NAN__ = 1
const D__FLT32_IS_IEC_60559__ = 2
const D__FLT32_MANT_DIG__ = 24
const D__FLT32_MAX_10_EXP__ = 38
const D__FLT32_MAX_EXP__ = 128
const D__FLT32_MAX__ = 3.40282346638528859811704183484516925e+38
const D__FLT32_MIN__ = 1.17549435082228750796873653722224568e-38
const D__FLT32_NORM_MAX__ = 3.40282346638528859811704183484516925e+38
const D__FLT64X_DECIMAL_DIG__ = 36
const D__FLT64X_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const D__FLT64X_DIG__ = 33
const D__FLT64X_EPSILON__ = 1.92592994438723585305597794258492732e-34
const D__FLT64X_HAS_DENORM__ = 1
const D__FLT64X_HAS_INFINITY__ = 1
const D__FLT64X_HAS_QUIET_NAN__ = 1
const D__FLT64X_IS_IEC_60559__ = 2
const D__FLT64X_MANT_DIG__ = 113
const D__FLT64X_MAX_10_EXP__ = 4932
const D__FLT64X_MAX_EXP__ = 16384
const D__FLT64X_MAX__ = "1.18973149535723176508575932662800702e+4932"
const D__FLT64X_MIN__ = 3.36210314311209350626267781732175260e-4932
const D__FLT64X_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const D__FLT64_DECIMAL_DIG__ = 17
const D__FLT64_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const D__FLT64_DIG__ = 15
const D__FLT64_EPSILON__ = 2.22044604925031308084726333618164062e-16
const D__FLT64_HAS_DENORM__ = 1
const D__FLT64_HAS_INFINITY__ = 1
const D__FLT64_HAS_QUIET_NAN__ = 1
const D__FLT64_IS_IEC_60559__ = 2
const D__FLT64_MANT_DIG__ = 53
const D__FLT64_MAX_10_EXP__ = 308
const D__FLT64_MAX_EXP__ = 1024
const D__FLT64_MAX__ = 1.79769313486231570814527423731704357e+308
const D__FLT64_MIN__ = 2.22507385850720138309023271733240406e-308
const D__FLT64_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const D__FLT_DECIMAL_DIG__ = 9
const D__FLT_DENORM_MIN__ = 1.40129846432481707092372958328991613e-45
const D__FLT_DIG__ = 6
const D__FLT_EPSILON__ = 1.19209289550781250000000000000000000e-7
const D__FLT_EVAL_METHOD_C99__ = 0
const D__FLT_EVAL_METHOD_TS_18661_3__ = 0
const D__FLT_EVAL_METHOD__ = 0
const D__FLT_HAS_DENORM__ = 1
const D__FLT_HAS_INFINITY__ = 1
const D__FLT_HAS_QUIET_NAN__ = 1
const D__FLT_IS_IEC_60559__ = 2
const D__FLT_MANT_DIG__ = 24
const D__FLT_MAX_10_EXP__ = 38
const D__FLT_MAX_EXP__ = 128
const D__FLT_MAX__ = 3.40282346638528859811704183484516925e+38
const D__FLT_MIN__ = 1.17549435082228750796873653722224568e-38
const D__FLT_NORM_MAX__ = 3.40282346638528859811704183484516925e+38
const D__FLT_RADIX__ = 2
const D__FP_FAST_FMA = 1
const D__FP_FAST_FMAF = 1
const D__FP_FAST_FMAF32 = 1
const D__FP_FAST_FMAF32x = 1
const D__FP_FAST_FMAF64 = 1
const D__FUNCTION__ = "__func__"
const D__GCC_ASM_FLAG_OUTPUTS__ = 1
const D__GCC_ATOMIC_BOOL_LOCK_FREE = 2
const D__GCC_ATOMIC_CHAR16_T_LOCK_FREE = 2
const D__GCC_ATOMIC_CHAR32_T_LOCK_FREE = 2
const D__GCC_ATOMIC_CHAR_LOCK_FREE = 2
const D__GCC_ATOMIC_INT_LOCK_FREE = 2
const D__GCC_ATOMIC_LLONG_LOCK_FREE = 2
const D__GCC_ATOMIC_LONG_LOCK_FREE = 2
const D__GCC_ATOMIC_POINTER_LOCK_FREE = 2
const D__GCC_ATOMIC_SHORT_LOCK_FREE = 2
const D__GCC_ATOMIC_TEST_AND_SET_TRUEVAL = 1
const D__GCC_ATOMIC_WCHAR_T_LOCK_FREE = 2
const D__GCC_CONSTRUCTIVE_SIZE = 64
const D__GCC_DESTRUCTIVE_SIZE = 256
const D__GCC_HAVE_DWARF2_CFI_ASM = 1
const D__GCC_HAVE_SYNC_COMPARE_AND_SWAP_1 = 1
const D__GCC_HAVE_SYNC_COMPARE_AND_SWAP_16 = 1
const D__GCC_HAVE_SYNC_COMPARE_AND_SWAP_2 = 1
const D__GCC_HAVE_SYNC_COMPARE_AND_SWAP_4 = 1
const D__GCC_HAVE_SYNC_COMPARE_AND_SWAP_8 = 1
const D__GCC_IEC_559 = 2
const D__GCC_IEC_559_COMPLEX = 2
const D__GNUC_EXECUTION_CHARSET_NAME = "UTF-8"
const D__GNUC_MINOR__ = 2
const D__GNUC_PATCHLEVEL__ = 0
const D__GNUC_STDC_INLINE__ = 1
const D__GNUC_WIDE_EXECUTION_CHARSET_NAME = "UTF-32LE"
const D__GNUC__ = 12
const D__GXX_ABI_VERSION = 1017
const D__HAVE_SPECULATION_SAFE_VALUE = 1
const D__INT16_MAX__ = 0x7fff
const D__INT32_MAX__ = 0x7fffffff
const D__INT32_TYPE__ = "int"
const D__INT64_MAX__ = 0x7fffffffffffffff
const D__INT8_MAX__ = 0x7f
const D__INTMAX_MAX__ = 0x7fffffffffffffff
const D__INTMAX_WIDTH__ = 64
const D__INTPTR_MAX__ = 0x7fffffffffffffff
const D__INTPTR_WIDTH__ = 64
const D__INT_FAST16_MAX__ = 0x7fffffffffffffff
const D__INT_FAST16_WIDTH__ = 64
const D__INT_FAST32_MAX__ = 0x7fffffffffffffff
const D__INT_FAST32_WIDTH__ = 64
const D__INT_FAST64_MAX__ = 0x7fffffffffffffff
const D__INT_FAST64_WIDTH__ = 64
const D__INT_FAST8_MAX__ = 0x7f
const D__INT_FAST8_WIDTH__ = 8
const D__INT_LEAST16_MAX__ = 0x7fff
const D__INT_LEAST16_WIDTH__ = 16
const D__INT_LEAST32_MAX__ = 0x7fffffff
const D__INT_LEAST32_TYPE__ = "int"
const D__INT_LEAST32_WIDTH__ = 32
const D__INT_LEAST64_MAX__ = 0x7fffffffffffffff
const D__INT_LEAST64_WIDTH__ = 64
const D__INT_LEAST8_MAX__ = 0x7f
const D__INT_LEAST8_WIDTH__ = 8
const D__INT_MAX__ = 0x7fffffff
const D__INT_WIDTH__ = 32
const D__LDBL_DECIMAL_DIG__ = 36
const D__LDBL_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const D__LDBL_DIG__ = 33
const D__LDBL_EPSILON__ = 1.92592994438723585305597794258492732e-34
const D__LDBL_HAS_DENORM__ = 1
const D__LDBL_HAS_INFINITY__ = 1
const D__LDBL_HAS_QUIET_NAN__ = 1
const D__LDBL_IS_IEC_60559__ = 2
const D__LDBL_MANT_DIG__ = 113
const D__LDBL_MAX_10_EXP__ = 4932
const D__LDBL_MAX_EXP__ = 16384
const D__LDBL_MAX__ = "1.18973149535723176508575932662800702e+4932"
const D__LDBL_MIN__ = 3.36210314311209350626267781732175260e-4932
const D__LDBL_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const D__LITTLE_ENDIAN = 1234
const D__LONG_LONG_MAX__ = 0x7fffffffffffffff
const D__LONG_LONG_WIDTH__ = 64
const D__LONG_MAX = 0x7fffffffffffffff
const D__LONG_MAX__ = 0x7fffffffffffffff
const D__LONG_WIDTH__ = 64
const D__LP64__ = 1
const D__NO_INLINE__ = 1
const D__ORDER_BIG_ENDIAN__ = 4321
const D__ORDER_LITTLE_ENDIAN__ = 1234
const D__ORDER_PDP_ENDIAN__ = 3412
const D__PIC__ = 2
const D__PIE__ = 2
const D__PRAGMA_REDEFINE_EXTNAME = 1
const D__PRETTY_FUNCTION__ = "__func__"
const D__PTRDIFF_MAX__ = 0x7fffffffffffffff
const D__PTRDIFF_WIDTH__ = 64
const D__SCHAR_MAX__ = 0x7f
const D__SCHAR_WIDTH__ = 8
const D__SHRT_MAX__ = 0x7fff
const D__SHRT_WIDTH__ = 16
const D__SIG_ATOMIC_MAX__ = 0x7fffffff
const D__SIG_ATOMIC_TYPE__ = "int"
const D__SIG_ATOMIC_WIDTH__ = 32
const D__SIZEOF_DOUBLE__ = 8
const D__SIZEOF_FLOAT__ = 4
const D__SIZEOF_INT128__ = 16
const D__SIZEOF_INT__ = 4
const D__SIZEOF_LONG_DOUBLE__ = 8
const D__SIZEOF_LONG_LONG__ = 8
const D__SIZEOF_LONG__ = 8
const D__SIZEOF_POINTER__ = 8
const D__SIZEOF_PTRDIFF_T__ = 8
const D__SIZEOF_SHORT__ = 2
const D__SIZEOF_SIZE_T__ = 8
const D__SIZEOF_WCHAR_T__ = 4
const D__SIZEOF_WINT_T__ = 4
const D__SIZE_MAX__ = 0xffffffffffffffff
const D__SIZE_WIDTH__ = 64
const D__STDC_HOSTED__ = 1
const D__STDC_IEC_559_COMPLEX__ = 1
const D__STDC_IEC_559__ = 1
const D__STDC_IEC_60559_BFP__ = 201404
const D__STDC_IEC_60559_COMPLEX__ = 201404
const D__STDC_ISO_10646__ = 201706
const D__STDC_UTF_16__ = 1
const D__STDC_UTF_32__ = 1
const D__STDC_VERSION__ = 201710
const D__STDC__ = 1
const D__UINT16_MAX__ = 0xffff
const D__UINT32_MAX__ = 0xffffffff
const D__UINT64_MAX__ = 0xffffffffffffffff
const D__UINT8_MAX__ = 0xff
const D__UINTMAX_MAX__ = 0xffffffffffffffff
const D__UINTPTR_MAX__ = 0xffffffffffffffff
const D__UINT_FAST16_MAX__ = 0xffffffffffffffff
const D__UINT_FAST32_MAX__ = 0xffffffffffffffff
const D__UINT_FAST64_MAX__ = 0xffffffffffffffff
const D__UINT_FAST8_MAX__ = 0xff
const D__UINT_LEAST16_MAX__ = 0xffff
const D__UINT_LEAST32_MAX__ = 0xffffffff
const D__UINT_LEAST64_MAX__ = 0xffffffffffffffff
const D__UINT_LEAST8_MAX__ = 0xff
const D__USE_TIME_BITS64 = 1
const D__VERSION__ = "12.2.0"
const D__WCHAR_MAX__ = 0xffffffff
const D__WCHAR_MIN__ = 0
const D__WCHAR_WIDTH__ = 32
const D__WINT_MAX__ = 0xffffffff
const D__WINT_MIN__ = 0
const D__WINT_WIDTH__ = 32
const D__aarch64__ = 1
const D__bool_true_false_are_defined = 1
const D__gnu_linux__ = 1
const D__inline = "inline"
const D__linux = 1
const D__linux__ = 1
const D__pic__ = 2
const D__pie__ = 2
const D__restrict = "restrict"
const D__restrict_arr = "restrict"
const D__unix = 1
const D__unix__ = 1
const Dalloca = "__builtin_alloca"
const Dbool = "_Bool"
const Dfalse = 0
const Dlinux = 1
const Dmath_errhandling = 2
const Dstatic_assert = "_Static_assert"
const Dtrue = 1
const Dunix = 1

type T__builtin_va_list = uintptr

type T__predefined_size_t = uint64

type T__predefined_wchar_t = uint32

type T__predefined_ptrdiff_t = int64

type Tuintptr_t = uint64

type Tintptr_t = int64

type Tint8_t = int8

type Tint16_t = int16

type Tint32_t = int32

type Tint64_t = int64

type Tintmax_t = int64

type Tuint8_t = uint8

type Tuint16_t = uint16

type Tuint32_t = uint32

type Tuint64_t = uint64

type Tuintmax_t = uint64

type Tint_fast8_t = int8

type Tint_fast64_t = int64

type Tint_least8_t = int8

type Tint_least16_t = int16

type Tint_least32_t = int32

type Tint_least64_t = int64

type Tuint_fast8_t = uint8

type Tuint_fast64_t = uint64

type Tuint_least8_t = uint8

type Tuint_least16_t = uint16

type Tuint_least32_t = uint32

type Tuint_least64_t = uint64

type Tint_fast16_t = int32

type Tint_fast32_t = int32

type Tuint_fast16_t = uint32

type Tuint_fast32_t = uint32

type Twchar_t = uint32

type Tsize_t = uint64

type Tdiv_t = struct {
	Fquot int32
	Frem  int32
}

type Tldiv_t = struct {
	Fquot int64
	Frem  int64
}

type Tlldiv_t = struct {
	Fquot int64
	Frem  int64
}

type TH3Index = uint64

type TH3Error = uint32

type TH3ErrorCodes = int32

const _E_SUCCESS = 0
const _E_FAILED = 1
const _E_DOMAIN = 2
const _E_LATLNG_DOMAIN = 3
const _E_RES_DOMAIN = 4
const _E_CELL_INVALID = 5
const _E_DIR_EDGE_INVALID = 6
const _E_UNDIR_EDGE_INVALID = 7
const _E_VERTEX_INVALID = 8
const _E_PENTAGON = 9
const _E_DUPLICATE_INPUT = 10
const _E_NOT_NEIGHBORS = 11
const _E_RES_MISMATCH = 12
const _E_MEMORY_ALLOC = 13
const _E_MEMORY_BOUNDS = 14
const _E_OPTION_INVALID = 15

type TLatLng = struct {
	Flat float64
	Flng float64
}

type TCellBoundary = struct {
	FnumVerts int32
	Fverts    [10]TLatLng
}

type TGeoLoop = struct {
	FnumVerts int32
	Fverts    uintptr
}

type TGeoPolygon = struct {
	Fgeoloop  TGeoLoop
	FnumHoles int32
	Fholes    uintptr
}

type TGeoMultiPolygon = struct {
	FnumPolygons int32
	Fpolygons    uintptr
}

type TContainmentMode = int32

const _CONTAINMENT_CENTER = 0
const _CONTAINMENT_FULL = 1
const _CONTAINMENT_OVERLAPPING = 2
const _CONTAINMENT_OVERLAPPING_BBOX = 3
const _CONTAINMENT_INVALID = 4

type TLinkedLatLng = struct {
	Fvertex TLatLng
	Fnext   uintptr
}

type TLinkedLatLng1 = struct {
	Fvertex TLatLng
	Fnext   uintptr
}

type TLinkedGeoLoop = struct {
	Ffirst uintptr
	Flast  uintptr
	Fnext  uintptr
}

type TLinkedGeoLoop1 = struct {
	Ffirst uintptr
	Flast  uintptr
	Fnext  uintptr
}

type TLinkedGeoPolygon = struct {
	Ffirst uintptr
	Flast  uintptr
	Fnext  uintptr
}

type TLinkedGeoPolygon1 = struct {
	Ffirst uintptr
	Flast  uintptr
	Fnext  uintptr
}

type TCoordIJ = struct {
	Fi int32
	Fj int32
}

type Tssize_t = int64

type Toff_t = int64

type Tva_list = uintptr

type T__isoc_va_list = uintptr

type Tfpos_t = struct {
	F__lldata [0]int64
	F__align  [0]float64
	F__opaque [16]uint8
}

type T_G_fpos64_t = Tfpos_t

type Tcookie_io_functions_t = struct {
	Fread   uintptr
	Fwrite  uintptr
	Fseek   uintptr
	Fclose1 uintptr
}

type T_IO_cookie_io_functions_t = Tcookie_io_functions_t

type TLongitudeNormalization = int32

const _NORMALIZE_NONE = 0
const _NORMALIZE_EAST = 1
const _NORMALIZE_WEST = 2

type TBBox = struct {
	Fnorth float64
	Fsouth float64
	Feast  float64
	Fwest  float64
}

type TVec2d = struct {
	Fx float64
	Fy float64
}

type TCoordIJK = struct {
	Fi int32
	Fj int32
	Fk int32
}

var _UNIT_VECS = [7]TCoordIJK{
	0: {},
	1: {
		Fk: int32(1),
	},
	2: {
		Fj: int32(1),
	},
	3: {
		Fj: int32(1),
		Fk: int32(1),
	},
	4: {
		Fi: int32(1),
	},
	5: {
		Fi: int32(1),
		Fk: int32(1),
	},
	6: {
		Fi: int32(1),
		Fj: int32(1),
	},
}

type TDirection = int32

const _CENTER_DIGIT = 0
const _K_AXES_DIGIT = 1
const _J_AXES_DIGIT = 2
const _JK_AXES_DIGIT = 3
const _I_AXES_DIGIT = 4
const _IK_AXES_DIGIT = 5
const _IJ_AXES_DIGIT = 6
const _INVALID_DIGIT = 7
const _NUM_DIGITS = 7
const _PENTAGON_SKIPPED_DIGIT = 1

type TVertexNode = struct {
	Ffrom TLatLng
	Fto   TLatLng
	Fnext uintptr
}

type TVertexNode1 = struct {
	Ffrom TLatLng
	Fto   TLatLng
	Fnext uintptr
}

type TVertexGraph = struct {
	Fbuckets    uintptr
	FnumBuckets int32
	Fsize       int32
	Fres        int32
}

type Tfloat_t = float32

type Tdouble_t = float64

type Tlocale_t = uintptr

type TFaceIJK = struct {
	Fface  int32
	Fcoord TCoordIJK
}

type TFaceOrientIJK = struct {
	Fface      int32
	Ftranslate TCoordIJK
	FccwRot60  int32
}

type TOverage = int32

const _NO_OVERAGE = 0
const _FACE_EDGE = 1
const _NEW_FACE = 2

type TBaseCellData = struct {
	FhomeFijk     TFaceIJK
	FisPentagon   int32
	FcwOffsetPent [2]int32
}

/*
 * Copyright 2017, 2020-2021 Uber Technologies, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/** @file vertexGraph.h
 * @brief   Data structure for storing a graph of vertices
 */

/*
 * Return codes from gridDiskUnsafe and related functions.
 */

// C documentation
//
//	/**
//	 * Directions used for traversing a hexagonal ring counterclockwise around
//	 * {1, 0, 0}
//	 *
//	 * <pre>
//	 *      _
//	 *    _/ \\_
//	 *   / \\5/ \ *   \\0/ \\4/
//	 *   / \\_/ \ *   \\1/ \\3/
//	 *     \\2/
//	 * </pre>
//	 */
var _DIRECTIONS = [6]TDirection{
	0: int32(_J_AXES_DIGIT),
	1: int32(_JK_AXES_DIGIT),
	2: int32(_K_AXES_DIGIT),
	3: int32(_IK_AXES_DIGIT),
	4: int32(_I_AXES_DIGIT),
	5: int32(_IJ_AXES_DIGIT),
}

// C documentation
//
//	/**
//	 * Direction used for traversing to the next outward hexagonal ring.
//	 */
var _NEXT_RING_DIRECTION = int32(_I_AXES_DIGIT)

// C documentation
//
//	/**
//	 * New digit when traversing along class II grids.
//	 *
//	 * Current digit -> direction -> new digit.
//	 */
var _NEW_DIGIT_II = [7][7]TDirection{
	0: {
		1: int32(_K_AXES_DIGIT),
		2: int32(_J_AXES_DIGIT),
		3: int32(_JK_AXES_DIGIT),
		4: int32(_I_AXES_DIGIT),
		5: int32(_IK_AXES_DIGIT),
		6: int32(_IJ_AXES_DIGIT),
	},
	1: {
		0: int32(_K_AXES_DIGIT),
		1: int32(_I_AXES_DIGIT),
		2: int32(_JK_AXES_DIGIT),
		3: int32(_IJ_AXES_DIGIT),
		4: int32(_IK_AXES_DIGIT),
		5: int32(_J_AXES_DIGIT),
	},
	2: {
		0: int32(_J_AXES_DIGIT),
		1: int32(_JK_AXES_DIGIT),
		2: int32(_K_AXES_DIGIT),
		3: int32(_I_AXES_DIGIT),
		4: int32(_IJ_AXES_DIGIT),
		6: int32(_IK_AXES_DIGIT),
	},
	3: {
		0: int32(_JK_AXES_DIGIT),
		1: int32(_IJ_AXES_DIGIT),
		2: int32(_I_AXES_DIGIT),
		3: int32(_IK_AXES_DIGIT),
		5: int32(_K_AXES_DIGIT),
		6: int32(_J_AXES_DIGIT),
	},
	4: {
		0: int32(_I_AXES_DIGIT),
		1: int32(_IK_AXES_DIGIT),
		2: int32(_IJ_AXES_DIGIT),
		4: int32(_J_AXES_DIGIT),
		5: int32(_JK_AXES_DIGIT),
		6: int32(_K_AXES_DIGIT),
	},
	5: {
		0: int32(_IK_AXES_DIGIT),
		1: int32(_J_AXES_DIGIT),
		3: int32(_K_AXES_DIGIT),
		4: int32(_JK_AXES_DIGIT),
		5: int32(_IJ_AXES_DIGIT),
		6: int32(_I_AXES_DIGIT),
	},
	6: {
		0: int32(_IJ_AXES_DIGIT),
		2: int32(_IK_AXES_DIGIT),
		3: int32(_J_AXES_DIGIT),
		4: int32(_K_AXES_DIGIT),
		5: int32(_I_AXES_DIGIT),
		6: int32(_JK_AXES_DIGIT),
	},
}

// C documentation
//
//	/**
//	 * New traversal direction when traversing along class II grids.
//	 *
//	 * Current digit -> direction -> new ap7 move (at coarser level).
//	 */
var _NEW_ADJUSTMENT_II = [7][7]TDirection{
	0: {},
	1: {
		1: int32(_K_AXES_DIGIT),
		3: int32(_K_AXES_DIGIT),
		5: int32(_IK_AXES_DIGIT),
	},
	2: {
		2: int32(_J_AXES_DIGIT),
		3: int32(_JK_AXES_DIGIT),
		6: int32(_J_AXES_DIGIT),
	},
	3: {
		1: int32(_K_AXES_DIGIT),
		2: int32(_JK_AXES_DIGIT),
		3: int32(_JK_AXES_DIGIT),
	},
	4: {
		4: int32(_I_AXES_DIGIT),
		5: int32(_I_AXES_DIGIT),
		6: int32(_IJ_AXES_DIGIT),
	},
	5: {
		1: int32(_IK_AXES_DIGIT),
		4: int32(_I_AXES_DIGIT),
		5: int32(_IK_AXES_DIGIT),
	},
	6: {
		2: int32(_J_AXES_DIGIT),
		4: int32(_IJ_AXES_DIGIT),
		6: int32(_IJ_AXES_DIGIT),
	},
}

// C documentation
//
//	/**
//	 * New traversal direction when traversing along class III grids.
//	 *
//	 * Current digit -> direction -> new ap7 move (at coarser level).
//	 */
var _NEW_DIGIT_III = [7][7]TDirection{
	0: {
		1: int32(_K_AXES_DIGIT),
		2: int32(_J_AXES_DIGIT),
		3: int32(_JK_AXES_DIGIT),
		4: int32(_I_AXES_DIGIT),
		5: int32(_IK_AXES_DIGIT),
		6: int32(_IJ_AXES_DIGIT),
	},
	1: {
		0: int32(_K_AXES_DIGIT),
		1: int32(_J_AXES_DIGIT),
		2: int32(_JK_AXES_DIGIT),
		3: int32(_I_AXES_DIGIT),
		4: int32(_IK_AXES_DIGIT),
		5: int32(_IJ_AXES_DIGIT),
	},
	2: {
		0: int32(_J_AXES_DIGIT),
		1: int32(_JK_AXES_DIGIT),
		2: int32(_I_AXES_DIGIT),
		3: int32(_IK_AXES_DIGIT),
		4: int32(_IJ_AXES_DIGIT),
		6: int32(_K_AXES_DIGIT),
	},
	3: {
		0: int32(_JK_AXES_DIGIT),
		1: int32(_I_AXES_DIGIT),
		2: int32(_IK_AXES_DIGIT),
		3: int32(_IJ_AXES_DIGIT),
		5: int32(_K_AXES_DIGIT),
		6: int32(_J_AXES_DIGIT),
	},
	4: {
		0: int32(_I_AXES_DIGIT),
		1: int32(_IK_AXES_DIGIT),
		2: int32(_IJ_AXES_DIGIT),
		4: int32(_K_AXES_DIGIT),
		5: int32(_J_AXES_DIGIT),
		6: int32(_JK_AXES_DIGIT),
	},
	5: {
		0: int32(_IK_AXES_DIGIT),
		1: int32(_IJ_AXES_DIGIT),
		3: int32(_K_AXES_DIGIT),
		4: int32(_J_AXES_DIGIT),
		5: int32(_JK_AXES_DIGIT),
		6: int32(_I_AXES_DIGIT),
	},
	6: {
		0: int32(_IJ_AXES_DIGIT),
		2: int32(_K_AXES_DIGIT),
		3: int32(_J_AXES_DIGIT),
		4: int32(_JK_AXES_DIGIT),
		5: int32(_I_AXES_DIGIT),
		6: int32(_IK_AXES_DIGIT),
	},
}

// C documentation
//
//	/**
//	 * New traversal direction when traversing along class III grids.
//	 *
//	 * Current digit -> direction -> new ap7 move (at coarser level).
//	 */
var _NEW_ADJUSTMENT_III = [7][7]TDirection{
	0: {},
	1: {
		1: int32(_K_AXES_DIGIT),
		3: int32(_JK_AXES_DIGIT),
		5: int32(_K_AXES_DIGIT),
	},
	2: {
		2: int32(_J_AXES_DIGIT),
		3: int32(_J_AXES_DIGIT),
		6: int32(_IJ_AXES_DIGIT),
	},
	3: {
		1: int32(_JK_AXES_DIGIT),
		2: int32(_J_AXES_DIGIT),
		3: int32(_JK_AXES_DIGIT),
	},
	4: {
		4: int32(_I_AXES_DIGIT),
		5: int32(_IK_AXES_DIGIT),
		6: int32(_I_AXES_DIGIT),
	},
	5: {
		1: int32(_K_AXES_DIGIT),
		4: int32(_IK_AXES_DIGIT),
		5: int32(_IK_AXES_DIGIT),
	},
	6: {
		2: int32(_IJ_AXES_DIGIT),
		4: int32(_I_AXES_DIGIT),
		6: int32(_IJ_AXES_DIGIT),
	},
}

// C documentation
//
//	/**
//	 * k value which will encompass all cells at resolution 15.
//	 * This is the largest possible k in the H3 grid system.
//	 */
var _K_ALL_CELLS_AT_RES_15 = int32(13780510)

// C documentation
//
//	/**
//	 * Maximum number of cells that result from the gridDisk algorithm with the
//	 * given k. Formula source and proof: https://oeis.org/A003215
//	 *
//	 * @param   k   k value, k >= 0.
//	 * @param out   size in indexes
//	 */
func XmaxGridDiskSize(tls *libc.TLS, k int32, out uintptr) (r TH3Error) {
	if k < 0 {
		return uint32(_E_DOMAIN)
	}
	if k >= _K_ALL_CELLS_AT_RES_15 {
		// If a k value of this value or above is provided, this function will
		// estimate more cells than exist in the H3 grid at the finest
		// resolution. This is a problem since the function does signed integer
		// arithmetic on `k`, which could overflow. To prevent that, instead
		// substitute the maximum number of cells in the grid, as it should not
		// be possible for the gridDisk functions to exceed that. Note this is
		// not resolution specific. So, when resolution < 15, this function may
		// still estimate a size larger than the number of cells in the grid.
		return XgetNumCells(tls, int32(DMAX_H3_RES), out)
	}
	*(*Tint64_t)(unsafe.Pointer(out)) = int64(3)*int64(k)*(int64(k)+int64(1)) + int64(1)
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * Produce cells within grid distance k of the origin cell.
//	 *
//	 * k-ring 0 is defined as the origin cell, k-ring 1 is defined as k-ring 0 and
//	 * all neighboring cells, and so on.
//	 *
//	 * Output is placed in the provided array in no particular order. Elements of
//	 * the output array may be left zero, as can happen when crossing a pentagon.
//	 *
//	 * @param  origin   origin cell
//	 * @param  k        k >= 0
//	 * @param  out      zero-filled array which must be of size maxGridDiskSize(k)
//	 */
func XgridDisk(tls *libc.TLS, origin TH3Index, k int32, out uintptr) (r TH3Error) {
	return XgridDiskDistances(tls, origin, k, out, libc.UintptrFromInt32(0))
}

// C documentation
//
//	/**
//	 * Produce cells and their distances from the given origin cell, up to
//	 * distance k.
//	 *
//	 * k-ring 0 is defined as the origin cell, k-ring 1 is defined as k-ring 0 and
//	 * all neighboring cells, and so on.
//	 *
//	 * Output is placed in the provided array in no particular order. Elements of
//	 * the output array may be left zero, as can happen when crossing a pentagon.
//	 *
//	 * @param  origin      origin cell
//	 * @param  k           k >= 0
//	 * @param  out         zero-filled array which must be of size
//	 * maxGridDiskSize(k)
//	 * @param  distances   NULL or a zero-filled array which must be of size
//	 *                     maxGridDiskSize(k)
//	 */
func XgridDiskDistances(tls *libc.TLS, origin TH3Index, k int32, out uintptr, distances uintptr) (r TH3Error) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var err, failed, result TH3Error
	var _ /* maxIdx at bp+0 */ Tint64_t
	_, _, _ = err, failed, result
	// Optimistically try the faster gridDiskUnsafe algorithm first
	failed = XgridDiskDistancesUnsafe(tls, origin, k, out, distances)
	if failed != 0 {
		err = XmaxGridDiskSize(tls, k, bp)
		if err != 0 {
			return err
		}
		// Fast algo failed, fall back to slower, correct algo
		// and also wipe out array because contents untrustworthy
		libc.Xmemset(tls, out, 0, libc.Uint64FromInt64(*(*Tint64_t)(unsafe.Pointer(bp)))*uint64(8))
		if distances == libc.UintptrFromInt32(0) {
			distances = libc.Xcalloc(tls, libc.Uint64FromInt64(*(*Tint64_t)(unsafe.Pointer(bp))), uint64(4))
			if !(distances != 0) {
				return uint32(_E_MEMORY_ALLOC)
			}
			result = X_gridDiskDistancesInternal(tls, origin, k, out, distances, *(*Tint64_t)(unsafe.Pointer(bp)), 0)
			libc.Xfree(tls, distances)
			return result
		} else {
			libc.Xmemset(tls, distances, 0, libc.Uint64FromInt64(*(*Tint64_t)(unsafe.Pointer(bp)))*uint64(4))
			return X_gridDiskDistancesInternal(tls, origin, k, out, distances, *(*Tint64_t)(unsafe.Pointer(bp)), 0)
		}
	} else {
		return uint32(_E_SUCCESS)
	}
	return r
}

// C documentation
//
//	/**
//	 * Internal algorithm for the safe but slow version of gridDiskDistances
//	 *
//	 * Adds the origin cell to the output set (treating it as a hash set)
//	 * and recurses to its neighbors, if needed.
//	 *
//	 * @param  origin      Origin cell
//	 * @param  k           Maximum distance to move from the origin
//	 * @param  out         Array treated as a hash set, elements being either
//	 *                     H3Index or 0.
//	 * @param  distances   Scratch area, with elements paralleling the out array.
//	 *                     Elements indicate ijk distance from the origin cell to
//	 *                     the output cell
//	 * @param  maxIdx      Size of out and scratch arrays (must be
//	 * maxGridDiskSize(k))
//	 * @param  curK        Current distance from the origin
//	 */
func X_gridDiskDistancesInternal(tls *libc.TLS, origin TH3Index, k int32, out uintptr, distances uintptr, maxIdx Tint64_t, curK int32) (r TH3Error) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i int32
	var neighborResult TH3Error
	var off Tint64_t
	var _ /* nextNeighbor at bp+8 */ TH3Index
	var _ /* rotations at bp+0 */ int32
	_, _, _ = i, neighborResult, off
	// Put origin in the output array. out is used as a hash set.
	off = libc.Int64FromUint64(origin % libc.Uint64FromInt64(maxIdx))
	for *(*TH3Index)(unsafe.Pointer(out + uintptr(off)*8)) != uint64(0) && *(*TH3Index)(unsafe.Pointer(out + uintptr(off)*8)) != origin {
		off = (off + int64(1)) % maxIdx
	}
	// We either got a free slot in the hash set or hit a duplicate
	// We might need to process the duplicate anyways because we got
	// here on a longer path before.
	if *(*TH3Index)(unsafe.Pointer(out + uintptr(off)*8)) == origin && *(*int32)(unsafe.Pointer(distances + uintptr(off)*4)) <= curK {
		return uint32(_E_SUCCESS)
	}
	*(*TH3Index)(unsafe.Pointer(out + uintptr(off)*8)) = origin
	*(*int32)(unsafe.Pointer(distances + uintptr(off)*4)) = curK
	// Base case: reached an index k away from the origin.
	if curK >= k {
		return uint32(_E_SUCCESS)
	}
	// Recurse to all neighbors in no particular order.
	i = 0
	for {
		if !(i < int32(6)) {
			break
		}
		*(*int32)(unsafe.Pointer(bp)) = 0
		neighborResult = Xh3NeighborRotations(tls, origin, _DIRECTIONS[i], bp, bp+8)
		if neighborResult != uint32(_E_PENTAGON) {
			// E_PENTAGON is an expected case when trying to traverse off of
			// pentagons.
			if neighborResult != uint32(_E_SUCCESS) {
				return neighborResult
			}
			neighborResult = X_gridDiskDistancesInternal(tls, *(*TH3Index)(unsafe.Pointer(bp + 8)), k, out, distances, maxIdx, curK+int32(1))
			if neighborResult != 0 {
				return neighborResult
			}
		}
		goto _1
	_1:
		;
		i++
	}
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * Safe but slow version of gridDiskDistances (also called by it when needed).
//	 *
//	 * Adds the origin cell to the output set (treating it as a hash set)
//	 * and recurses to its neighbors, if needed.
//	 *
//	 * @param  origin      Origin cell
//	 * @param  k           Maximum distance to move from the origin
//	 * @param  out         Array treated as a hash set, elements being either
//	 *                     H3Index or 0.
//	 * @param  distances   Scratch area, with elements paralleling the out array.
//	 *                     Elements indicate ijk distance from the origin cell to
//	 *                     the output cell
//	 */
func XgridDiskDistancesSafe(tls *libc.TLS, origin TH3Index, k int32, out uintptr, distances uintptr) (r TH3Error) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var err TH3Error
	var _ /* maxIdx at bp+0 */ Tint64_t
	_ = err
	err = XmaxGridDiskSize(tls, k, bp)
	if err != 0 {
		return err
	}
	return X_gridDiskDistancesInternal(tls, origin, k, out, distances, *(*Tint64_t)(unsafe.Pointer(bp)), 0)
}

// C documentation
//
//	/**
//	 * Returns the hexagon index neighboring the origin, in the direction dir.
//	 *
//	 * Implementation note: The only reachable case where this returns 0 is if the
//	 * origin is a pentagon and the translation is in the k direction. Thus,
//	 * 0 can only be returned if origin is a pentagon.
//	 *
//	 * @param origin Origin index
//	 * @param dir Direction to move in
//	 * @param rotations Number of ccw rotations to perform to reorient the
//	 *                  translation vector. Will be modified to the new number of
//	 *                  rotations to perform (such as when crossing a face edge.)
//	 * @param out H3Index of the specified neighbor if succesful
//	 * @return E_SUCCESS on success
//	 */
func Xh3NeighborRotations(tls *libc.TLS, origin TH3Index, dir TDirection, rotations uintptr, out uintptr) (r1 TH3Error) {
	var alreadyAdjustedKSubsequence, i, i1, i2, newBaseCell, newRotations, oldBaseCell, r, v2, v4 int32
	var current TH3Index
	var nextDir, oldDigit, oldLeadingDigit TDirection
	var v3, v5 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = alreadyAdjustedKSubsequence, current, i, i1, i2, newBaseCell, newRotations, nextDir, oldBaseCell, oldDigit, oldLeadingDigit, r, v2, v3, v4, v5
	current = origin
	if dir < int32(_CENTER_DIGIT) || dir >= int32(_INVALID_DIGIT) {
		return uint32(_E_FAILED)
	}
	// Ensure that rotations is modulo'd by 6 before any possible addition,
	// to protect against signed integer overflow.
	*(*int32)(unsafe.Pointer(rotations)) = *(*int32)(unsafe.Pointer(rotations)) % int32(6)
	i = 0
	for {
		if !(i < *(*int32)(unsafe.Pointer(rotations))) {
			break
		}
		dir = X_rotate60ccw(tls, dir)
		goto _1
	_1:
		;
		i++
	}
	newRotations = 0
	oldBaseCell = libc.Int32FromUint64(current & (libc.Uint64FromInt32(libc.Int32FromInt32(127)) << libc.Int32FromInt32(DH3_BC_OFFSET)) >> libc.Int32FromInt32(DH3_BC_OFFSET))
	if oldBaseCell < 0 {
		if v3 = libc.Bool(0 != 0); !v3 {
			libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+2, int32(368), uintptr(unsafe.Pointer(&___func__)))
		}
		_ = v3 || libc.Bool(libc.Int32FromInt32(0) != 0)
		v2 = libc.Int32FromInt32(1)
	} else {
		v2 = 0
	}
	if v2 != 0 || oldBaseCell >= int32(DNUM_BASE_CELLS) {
		// Base cells less than zero can not be represented in an index
		return uint32(_E_CELL_INVALID)
	}
	oldLeadingDigit = X_h3LeadingNonZeroDigit(tls, current)
	// Adjust the indexing digits and, if needed, the base cell.
	r = libc.Int32FromUint64(current&(libc.Uint64FromUint64(15)<<libc.Int32FromInt32(DH3_RES_OFFSET))>>libc.Int32FromInt32(DH3_RES_OFFSET)) - int32(1)
	for int32(Dtrue) != 0 {
		if r == -int32(1) {
			current = current & ^(libc.Uint64FromInt32(libc.Int32FromInt32(127))<<libc.Int32FromInt32(DH3_BC_OFFSET)) | libc.Uint64FromInt32(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&XbaseCellNeighbors)) + uintptr(oldBaseCell)*28 + uintptr(dir)*4)))<<libc.Int32FromInt32(DH3_BC_OFFSET)
			newRotations = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&XbaseCellNeighbor60CCWRots)) + uintptr(oldBaseCell)*28 + uintptr(dir)*4))
			if libc.Int32FromUint64(current&(libc.Uint64FromInt32(libc.Int32FromInt32(127))<<libc.Int32FromInt32(DH3_BC_OFFSET))>>libc.Int32FromInt32(DH3_BC_OFFSET)) == int32(DINVALID_BASE_CELL) {
				// Adjust for the deleted k vertex at the base cell level.
				// This edge actually borders a different neighbor.
				current = current & ^(libc.Uint64FromInt32(libc.Int32FromInt32(127))<<libc.Int32FromInt32(DH3_BC_OFFSET)) | libc.Uint64FromInt32(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&XbaseCellNeighbors)) + uintptr(oldBaseCell)*28 + uintptr(_IK_AXES_DIGIT)*4)))<<libc.Int32FromInt32(DH3_BC_OFFSET)
				newRotations = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&XbaseCellNeighbor60CCWRots)) + uintptr(oldBaseCell)*28 + uintptr(_IK_AXES_DIGIT)*4))
				// perform the adjustment for the k-subsequence we're skipping
				// over.
				current = X_h3Rotate60ccw(tls, current)
				*(*int32)(unsafe.Pointer(rotations)) = *(*int32)(unsafe.Pointer(rotations)) + int32(1)
			}
			break
		} else {
			oldDigit = libc.Int32FromUint64(current >> ((libc.Int32FromInt32(DMAX_H3_RES) - (r + libc.Int32FromInt32(1))) * libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET)) & libc.Uint64FromInt32(libc.Int32FromInt32(7)))
			if oldDigit == int32(_INVALID_DIGIT) {
				// Only possible on invalid input
				return uint32(_E_CELL_INVALID)
			} else {
				if XisResolutionClassIII(tls, r+int32(1)) != 0 {
					current = current & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<((libc.Int32FromInt32(DMAX_H3_RES)-(r+libc.Int32FromInt32(1)))*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))) | libc.Uint64FromInt32(*(*TDirection)(unsafe.Pointer(uintptr(unsafe.Pointer(&_NEW_DIGIT_II)) + uintptr(oldDigit)*28 + uintptr(dir)*4)))<<((libc.Int32FromInt32(DMAX_H3_RES)-(r+libc.Int32FromInt32(1)))*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))
					nextDir = *(*TDirection)(unsafe.Pointer(uintptr(unsafe.Pointer(&_NEW_ADJUSTMENT_II)) + uintptr(oldDigit)*28 + uintptr(dir)*4))
				} else {
					current = current & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<((libc.Int32FromInt32(DMAX_H3_RES)-(r+libc.Int32FromInt32(1)))*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))) | libc.Uint64FromInt32(*(*TDirection)(unsafe.Pointer(uintptr(unsafe.Pointer(&_NEW_DIGIT_III)) + uintptr(oldDigit)*28 + uintptr(dir)*4)))<<((libc.Int32FromInt32(DMAX_H3_RES)-(r+libc.Int32FromInt32(1)))*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))
					nextDir = *(*TDirection)(unsafe.Pointer(uintptr(unsafe.Pointer(&_NEW_ADJUSTMENT_III)) + uintptr(oldDigit)*28 + uintptr(dir)*4))
				}
			}
			if nextDir != int32(_CENTER_DIGIT) {
				dir = nextDir
				r--
			} else {
				// No more adjustment to perform
				break
			}
		}
	}
	newBaseCell = libc.Int32FromUint64(current & (libc.Uint64FromInt32(libc.Int32FromInt32(127)) << libc.Int32FromInt32(DH3_BC_OFFSET)) >> libc.Int32FromInt32(DH3_BC_OFFSET))
	if X_isBaseCellPentagon(tls, newBaseCell) != 0 {
		alreadyAdjustedKSubsequence = 0
		// force rotation out of missing k-axes sub-sequence
		if X_h3LeadingNonZeroDigit(tls, current) == int32(_K_AXES_DIGIT) {
			if oldBaseCell != newBaseCell {
				// in this case, we traversed into the deleted
				// k subsequence of a pentagon base cell.
				// We need to rotate out of that case depending
				// on how we got here.
				// check for a cw/ccw offset face; default is ccw
				if X_baseCellIsCwOffset(tls, newBaseCell, XbaseCellData[oldBaseCell].FhomeFijk.Fface) != 0 {
					v4 = int32(1)
				} else {
					if v5 = libc.Bool(0 != 0); !v5 {
						libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+2, int32(434), uintptr(unsafe.Pointer(&___func__)))
					}
					_ = v5 || libc.Bool(libc.Int32FromInt32(0) != 0)
					v4 = libc.Int32FromInt32(0)
				}
				if v4 != 0 {
					current = X_h3Rotate60cw(tls, current)
				} else {
					// See cwOffsetPent in testGridDisk.c for why this is
					// unreachable.
					current = X_h3Rotate60ccw(tls, current)
				}
				alreadyAdjustedKSubsequence = int32(1)
			} else {
				// In this case, we traversed into the deleted
				// k subsequence from within the same pentagon
				// base cell.
				if oldLeadingDigit == int32(_CENTER_DIGIT) {
					// Undefined: the k direction is deleted from here
					return uint32(_E_PENTAGON)
				} else {
					if oldLeadingDigit == int32(_JK_AXES_DIGIT) {
						// Rotate out of the deleted k subsequence
						// We also need an additional change to the direction we're
						// moving in
						current = X_h3Rotate60ccw(tls, current)
						*(*int32)(unsafe.Pointer(rotations)) = *(*int32)(unsafe.Pointer(rotations)) + int32(1)
					} else {
						if oldLeadingDigit == int32(_IK_AXES_DIGIT) {
							// Rotate out of the deleted k subsequence
							// We also need an additional change to the direction we're
							// moving in
							current = X_h3Rotate60cw(tls, current)
							*(*int32)(unsafe.Pointer(rotations)) = *(*int32)(unsafe.Pointer(rotations)) + int32(5)
						} else {
							// TODO: Should never occur, but is reachable by fuzzer
							return uint32(_E_FAILED)
						}
					}
				}
			}
		}
		i1 = 0
		for {
			if !(i1 < newRotations) {
				break
			}
			current = X_h3RotatePent60ccw(tls, current)
			goto _6
		_6:
			;
			i1++
		}
		// Account for differing orientation of the base cells (this edge
		// might not follow properties of some other edges.)
		if oldBaseCell != newBaseCell {
			if X_isBaseCellPolarPentagon(tls, newBaseCell) != 0 {
				// 'polar' base cells behave differently because they have all
				// i neighbors.
				if oldBaseCell != int32(118) && oldBaseCell != int32(8) && X_h3LeadingNonZeroDigit(tls, current) != int32(_JK_AXES_DIGIT) {
					*(*int32)(unsafe.Pointer(rotations)) = *(*int32)(unsafe.Pointer(rotations)) + int32(1)
				}
			} else {
				if X_h3LeadingNonZeroDigit(tls, current) == int32(_IK_AXES_DIGIT) && !(alreadyAdjustedKSubsequence != 0) {
					// account for distortion introduced to the 5 neighbor by the
					// deleted k subsequence.
					*(*int32)(unsafe.Pointer(rotations)) = *(*int32)(unsafe.Pointer(rotations)) + int32(1)
				}
			}
		}
	} else {
		i2 = 0
		for {
			if !(i2 < newRotations) {
				break
			}
			current = X_h3Rotate60ccw(tls, current)
			goto _7
		_7:
			;
			i2++
		}
	}
	*(*int32)(unsafe.Pointer(rotations)) = (*(*int32)(unsafe.Pointer(rotations)) + newRotations) % int32(6)
	*(*TH3Index)(unsafe.Pointer(out)) = current
	return uint32(_E_SUCCESS)
}

var ___func__ = [20]uint8{'h', '3', 'N', 'e', 'i', 'g', 'h', 'b', 'o', 'r', 'R', 'o', 't', 'a', 't', 'i', 'o', 'n', 's'}

// C documentation
//
//	/**
//	 * Get the direction from the origin to a given neighbor. This is effectively
//	 * the reverse operation for h3NeighborRotations. Returns INVALID_DIGIT if the
//	 * cells are not neighbors.
//	 *
//	 * TODO: This is currently a brute-force algorithm, but as it's O(6) that's
//	 * probably acceptable.
//	 */
func XdirectionForNeighbor(tls *libc.TLS, origin TH3Index, destination TH3Index) (r TDirection) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var direction TDirection
	var isPent uint8
	var neighborError TH3Error
	var v2 int32
	var _ /* neighbor at bp+0 */ TH3Index
	var _ /* rotations at bp+8 */ int32
	_, _, _, _ = direction, isPent, neighborError, v2
	isPent = libc.Uint8FromInt32(XisPentagon(tls, origin))
	// Checks each neighbor, in order, to determine which direction the
	// destination neighbor is located. Skips CENTER_DIGIT since that
	// would be the origin; skips deleted K direction for pentagons.
	if isPent != 0 {
		v2 = int32(_J_AXES_DIGIT)
	} else {
		v2 = int32(_K_AXES_DIGIT)
	}
	direction = v2
	for {
		if !(direction < int32(_NUM_DIGITS)) {
			break
		}
		*(*int32)(unsafe.Pointer(bp + 8)) = 0
		neighborError = Xh3NeighborRotations(tls, origin, direction, bp+8, bp)
		if !(neighborError != 0) && *(*TH3Index)(unsafe.Pointer(bp)) == destination {
			return direction
		}
		goto _1
	_1:
		;
		direction++
	}
	return int32(_INVALID_DIGIT)
}

// C documentation
//
//	/**
//	 * gridDiskUnsafe produces indexes within k distance of the origin index.
//	 * Output behavior is undefined when one of the indexes returned by this
//	 * function is a pentagon or is in the pentagon distortion area.
//	 *
//	 * k-ring 0 is defined as the origin index, k-ring 1 is defined as k-ring 0 and
//	 * all neighboring indexes, and so on.
//	 *
//	 * Output is placed in the provided array in order of increasing distance from
//	 * the origin.
//	 *
//	 * @param origin Origin location.
//	 * @param k k >= 0
//	 * @param out Array which must be of size maxGridDiskSize(k).
//	 * @return 0 if no pentagon or pentagonal distortion area was encountered.
//	 */
func XgridDiskUnsafe(tls *libc.TLS, origin TH3Index, k int32, out uintptr) (r TH3Error) {
	return XgridDiskDistancesUnsafe(tls, origin, k, out, libc.UintptrFromInt32(0))
}

// C documentation
//
//	/**
//	 * gridDiskDistancesUnsafe produces indexes within k distance of the origin
//	 * index. Output behavior is undefined when one of the indexes returned by this
//	 * function is a pentagon or is in the pentagon distortion area.
//	 *
//	 * k-ring 0 is defined as the origin index, k-ring 1 is defined as k-ring 0 and
//	 * all neighboring indexes, and so on.
//	 *
//	 * Output is placed in the provided array in order of increasing distance from
//	 * the origin. The distances in hexagons is placed in the distances array at
//	 * the same offset.
//	 *
//	 * @param origin Origin location.
//	 * @param k k >= 0
//	 * @param out Array which must be of size maxGridDiskSize(k).
//	 * @param distances Null or array which must be of size maxGridDiskSize(k).
//	 * @return 0 if no pentagon or pentagonal distortion area was encountered.
//	 */
func XgridDiskDistancesUnsafe(tls *libc.TLS, _origin TH3Index, k int32, out uintptr, distances uintptr) (r TH3Error) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*TH3Index)(unsafe.Pointer(bp)) = _origin
	var direction, i, idx, ring int32
	var neighborResult, neighborResult1 TH3Error
	var _ /* rotations at bp+8 */ int32
	_, _, _, _, _, _ = direction, i, idx, neighborResult, neighborResult1, ring
	// Return codes:
	// 1 Pentagon was encountered
	// 2 Pentagon distortion (deleted k subsequence) was encountered
	// Pentagon being encountered is not itself a problem; really the deleted
	// k-subsequence is the problem, but for compatibility reasons we fail on
	// the pentagon.
	if k < 0 {
		return uint32(_E_DOMAIN)
	}
	// k must be >= 0, so origin is always needed
	idx = 0
	*(*TH3Index)(unsafe.Pointer(out + uintptr(idx)*8)) = *(*TH3Index)(unsafe.Pointer(bp))
	if distances != 0 {
		*(*int32)(unsafe.Pointer(distances + uintptr(idx)*4)) = 0
	}
	idx++
	if XisPentagon(tls, *(*TH3Index)(unsafe.Pointer(bp))) != 0 {
		// Pentagon was encountered; bail out as user doesn't want this.
		return uint32(_E_PENTAGON)
	}
	// 0 < ring <= k, current ring
	ring = int32(1)
	// 0 <= direction < 6, current side of the ring
	direction = 0
	// 0 <= i < ring, current position on the side of the ring
	i = 0
	// Number of 60 degree ccw rotations to perform on the direction (based on
	// which faces have been crossed.)
	*(*int32)(unsafe.Pointer(bp + 8)) = 0
	for ring <= k {
		if direction == 0 && i == 0 {
			// Not putting in the output set as it will be done later, at
			// the end of this ring.
			neighborResult = Xh3NeighborRotations(tls, *(*TH3Index)(unsafe.Pointer(bp)), _NEXT_RING_DIRECTION, bp+8, bp)
			if neighborResult != 0 {
				// Should not be possible because `origin` would have to be a
				// pentagon
				// TODO: Reachable via fuzzer
				return neighborResult
			}
			if XisPentagon(tls, *(*TH3Index)(unsafe.Pointer(bp))) != 0 {
				// Pentagon was encountered; bail out as user doesn't want this.
				return uint32(_E_PENTAGON)
			}
		}
		neighborResult1 = Xh3NeighborRotations(tls, *(*TH3Index)(unsafe.Pointer(bp)), _DIRECTIONS[direction], bp+8, bp)
		if neighborResult1 != 0 {
			return neighborResult1
		}
		*(*TH3Index)(unsafe.Pointer(out + uintptr(idx)*8)) = *(*TH3Index)(unsafe.Pointer(bp))
		if distances != 0 {
			*(*int32)(unsafe.Pointer(distances + uintptr(idx)*4)) = ring
		}
		idx++
		i++
		// Check if end of this side of the k-ring
		if i == ring {
			i = 0
			direction++
			// Check if end of this ring.
			if direction == int32(6) {
				direction = 0
				ring++
			}
		}
		if XisPentagon(tls, *(*TH3Index)(unsafe.Pointer(bp))) != 0 {
			// Pentagon was encountered; bail out as user doesn't want this.
			return uint32(_E_PENTAGON)
		}
	}
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * gridDisksUnsafe takes an array of input hex IDs and a max k-ring and returns
//	 * an array of hexagon IDs sorted first by the original hex IDs and then by the
//	 * k-ring (0 to max), with no guaranteed sorting within each k-ring group.
//	 *
//	 * @param h3Set A pointer to an array of H3Indexes
//	 * @param length The total number of H3Indexes in h3Set
//	 * @param k The number of rings to generate
//	 * @param out A pointer to the output memory to dump the new set of H3Indexes to
//	 *            The memory block should be equal to maxGridDiskSize(k) * length
//	 * @return 0 if no pentagon is encountered. Cannot trust output otherwise
//	 */
func XgridDisksUnsafe(tls *libc.TLS, h3Set uintptr, length int32, k int32, out uintptr) (r TH3Error) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var err, failed TH3Error
	var i int32
	var segment uintptr
	var _ /* segmentSize at bp+0 */ Tint64_t
	_, _, _, _ = err, failed, i, segment
	err = XmaxGridDiskSize(tls, k, bp)
	if err != 0 {
		return err
	}
	i = 0
	for {
		if !(i < length) {
			break
		}
		// Determine the appropriate segment of the output array to operate on
		segment = out + uintptr(int64(i)**(*Tint64_t)(unsafe.Pointer(bp)))*8
		failed = XgridDiskUnsafe(tls, *(*TH3Index)(unsafe.Pointer(h3Set + uintptr(i)*8)), k, segment)
		if failed != 0 {
			return failed
		}
		goto _1
	_1:
		;
		i++
	}
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * Returns the "hollow" ring of hexagons at exactly grid distance k from
//	 * the origin hexagon. In particular, k=0 returns just the origin hexagon.
//	 *
//	 * A nonzero failure code may be returned in some cases, for example,
//	 * if a pentagon is encountered.
//	 * Failure cases may be fixed in future versions.
//	 *
//	 * @param origin Origin location.
//	 * @param k k >= 0
//	 * @param out Array which must be of size 6 * k (or 1 if k == 0)
//	 * @return 0 if successful; nonzero otherwise.
//	 */
func XgridRingUnsafe(tls *libc.TLS, _origin TH3Index, k int32, out uintptr) (r TH3Error) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*TH3Index)(unsafe.Pointer(bp)) = _origin
	var direction, idx, pos, ring int32
	var lastIndex TH3Index
	var neighborResult, neighborResult1 TH3Error
	var _ /* rotations at bp+8 */ int32
	_, _, _, _, _, _, _ = direction, idx, lastIndex, neighborResult, neighborResult1, pos, ring
	// Short-circuit on 'identity' ring
	if k == 0 {
		*(*TH3Index)(unsafe.Pointer(out)) = *(*TH3Index)(unsafe.Pointer(bp))
		return uint32(_E_SUCCESS)
	}
	idx = 0
	// Number of 60 degree ccw rotations to perform on the direction (based on
	// which faces have been crossed.)
	*(*int32)(unsafe.Pointer(bp + 8)) = 0
	// Scratch structure for checking for pentagons
	if XisPentagon(tls, *(*TH3Index)(unsafe.Pointer(bp))) != 0 {
		// Pentagon was encountered; bail out as user doesn't want this.
		return uint32(_E_PENTAGON)
	}
	ring = 0
	for {
		if !(ring < k) {
			break
		}
		neighborResult = Xh3NeighborRotations(tls, *(*TH3Index)(unsafe.Pointer(bp)), _NEXT_RING_DIRECTION, bp+8, bp)
		if neighborResult != 0 {
			// Should not be possible because `origin` would have to be a
			// pentagon
			// TODO: Reachable via fuzzer
			return neighborResult
		}
		if XisPentagon(tls, *(*TH3Index)(unsafe.Pointer(bp))) != 0 {
			return uint32(_E_PENTAGON)
		}
		goto _1
	_1:
		;
		ring++
	}
	lastIndex = *(*TH3Index)(unsafe.Pointer(bp))
	*(*TH3Index)(unsafe.Pointer(out + uintptr(idx)*8)) = *(*TH3Index)(unsafe.Pointer(bp))
	idx++
	direction = 0
	for {
		if !(direction < int32(6)) {
			break
		}
		pos = 0
		for {
			if !(pos < k) {
				break
			}
			neighborResult1 = Xh3NeighborRotations(tls, *(*TH3Index)(unsafe.Pointer(bp)), _DIRECTIONS[direction], bp+8, bp)
			if neighborResult1 != 0 {
				// Should not be possible because `origin` would have to be a
				// pentagon
				// TODO: Reachable via fuzzer
				return neighborResult1
			}
			// Skip the very last index, it was already added. We do
			// however need to traverse to it because of the pentagonal
			// distortion check, below.
			if pos != k-int32(1) || direction != int32(5) {
				*(*TH3Index)(unsafe.Pointer(out + uintptr(idx)*8)) = *(*TH3Index)(unsafe.Pointer(bp))
				idx++
				if XisPentagon(tls, *(*TH3Index)(unsafe.Pointer(bp))) != 0 {
					return uint32(_E_PENTAGON)
				}
			}
			goto _3
		_3:
			;
			pos++
		}
		goto _2
	_2:
		;
		direction++
	}
	// Check that this matches the expected lastIndex, if it doesn't,
	// it indicates pentagonal distortion occurred and we should report
	// failure.
	if lastIndex != *(*TH3Index)(unsafe.Pointer(bp)) {
		return uint32(_E_PENTAGON)
	} else {
		return uint32(_E_SUCCESS)
	}
	return r
}

// C documentation
//
//	/**
//	 * maxPolygonToCellsSize returns the number of cells to allocate space for
//	 * when performing a polygonToCells on the given GeoJSON-like data structure.
//	 *
//	 * The size is the maximum of either the number of points in the geoloop or the
//	 * number of cells in the bounding box of the geoloop.
//	 *
//	 * @param geoPolygon A GeoJSON-like data structure indicating the poly to fill
//	 * @param res Hexagon resolution (0-15)
//	 * @param out number of cells to allocate for
//	 * @return 0 (E_SUCCESS) on success.
//	 */
func XmaxPolygonToCellsSize(tls *libc.TLS, geoPolygon uintptr, res int32, flags Tuint32_t, out uintptr) (r TH3Error) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var estimateErr, flagErr TH3Error
	var i, totalVerts int32
	var _ /* bbox at bp+0 */ TBBox
	var _ /* geoloop at bp+32 */ TGeoLoop
	var _ /* numHexagons at bp+48 */ Tint64_t
	_, _, _, _ = estimateErr, flagErr, i, totalVerts
	flagErr = XvalidatePolygonFlags(tls, flags)
	if flagErr != 0 {
		return flagErr
	}
	*(*TGeoLoop)(unsafe.Pointer(bp + 32)) = (*TGeoPolygon)(unsafe.Pointer(geoPolygon)).Fgeoloop
	XbboxFromGeoLoop(tls, bp+32, bp)
	estimateErr = XbboxHexEstimate(tls, bp, res, bp+48)
	if estimateErr != 0 {
		return estimateErr
	}
	// This algorithm assumes that the number of vertices is usually less than
	// the number of hexagons, but when it's wrong, this will keep it from
	// failing
	totalVerts = (*(*TGeoLoop)(unsafe.Pointer(bp + 32))).FnumVerts
	i = 0
	for {
		if !(i < (*TGeoPolygon)(unsafe.Pointer(geoPolygon)).FnumHoles) {
			break
		}
		totalVerts += (*(*TGeoLoop)(unsafe.Pointer((*TGeoPolygon)(unsafe.Pointer(geoPolygon)).Fholes + uintptr(i)*16))).FnumVerts
		goto _1
	_1:
		;
		i++
	}
	if *(*Tint64_t)(unsafe.Pointer(bp + 48)) < int64(totalVerts) {
		*(*Tint64_t)(unsafe.Pointer(bp + 48)) = int64(totalVerts)
	}
	// When the polygon is very small, near an icosahedron edge and is an odd
	// resolution, the line tracing needs an extra buffer than the estimator
	// function provides (but beefing that up to cover causes most situations to
	// overallocate memory)
	*(*Tint64_t)(unsafe.Pointer(bp + 48)) += int64(DPOLYGON_TO_CELLS_BUFFER)
	*(*Tint64_t)(unsafe.Pointer(out)) = *(*Tint64_t)(unsafe.Pointer(bp + 48))
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * _getEdgeHexagons takes a given geoloop ring (either the main geoloop or
//	 * one of the holes) and traces it with hexagons and updates the search and
//	 * found memory blocks. This is used for determining the initial hexagon set
//	 * for the polygonToCells algorithm to execute on.
//	 *
//	 * @param geoloop The geoloop (or hole) to be traced
//	 * @param numHexagons The maximum number of hexagons possible for the geoloop
//	 *                    (also the bounds of the search and found arrays)
//	 * @param res The hexagon resolution (0-15)
//	 * @param numSearchHexes The number of hexagons found so far to be searched
//	 * @param search The block of memory containing the hexagons to search from
//	 * @param found The block of memory containing the hexagons found from the
//	 * search
//	 *
//	 * @return An error code if the hash function cannot insert a found hexagon
//	 *         into the found array.
//	 */
func X_getEdgeHexagons(tls *libc.TLS, geoloop uintptr, numHexagons Tint64_t, res int32, numSearchHexes uintptr, search uintptr, found uintptr) (r TH3Error) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var e, estimateErr TH3Error
	var i int32
	var invNumHexesEst float64
	var j, loc, loopCount Tint64_t
	var v2 TLatLng
	var _ /* destination at bp+16 */ TLatLng
	var _ /* interpolate at bp+40 */ TLatLng
	var _ /* numHexesEstimate at bp+32 */ Tint64_t
	var _ /* origin at bp+0 */ TLatLng
	var _ /* pointHex at bp+56 */ TH3Index
	_, _, _, _, _, _, _, _ = e, estimateErr, i, invNumHexesEst, j, loc, loopCount, v2
	i = 0
	for {
		if !(i < (*TGeoLoop)(unsafe.Pointer(geoloop)).FnumVerts) {
			break
		}
		*(*TLatLng)(unsafe.Pointer(bp)) = *(*TLatLng)(unsafe.Pointer((*TGeoLoop)(unsafe.Pointer(geoloop)).Fverts + uintptr(i)*16))
		if i == (*TGeoLoop)(unsafe.Pointer(geoloop)).FnumVerts-int32(1) {
			v2 = *(*TLatLng)(unsafe.Pointer((*TGeoLoop)(unsafe.Pointer(geoloop)).Fverts))
		} else {
			v2 = *(*TLatLng)(unsafe.Pointer((*TGeoLoop)(unsafe.Pointer(geoloop)).Fverts + uintptr(i+int32(1))*16))
		}
		*(*TLatLng)(unsafe.Pointer(bp + 16)) = v2
		estimateErr = XlineHexEstimate(tls, bp, bp+16, res, bp+32)
		if estimateErr != 0 {
			return estimateErr
		}
		j = 0
		for {
			if !(j < *(*Tint64_t)(unsafe.Pointer(bp + 32))) {
				break
			}
			invNumHexesEst = float64(1) / float64(*(*Tint64_t)(unsafe.Pointer(bp + 32)))
			(*(*TLatLng)(unsafe.Pointer(bp + 40))).Flat = (*(*TLatLng)(unsafe.Pointer(bp))).Flat*float64(*(*Tint64_t)(unsafe.Pointer(bp + 32))-j)*invNumHexesEst + (*(*TLatLng)(unsafe.Pointer(bp + 16))).Flat*float64(j)*invNumHexesEst
			(*(*TLatLng)(unsafe.Pointer(bp + 40))).Flng = (*(*TLatLng)(unsafe.Pointer(bp))).Flng*float64(*(*Tint64_t)(unsafe.Pointer(bp + 32))-j)*invNumHexesEst + (*(*TLatLng)(unsafe.Pointer(bp + 16))).Flng*float64(j)*invNumHexesEst
			e = XlatLngToCell(tls, bp+40, res, bp+56)
			if e != 0 {
				return e
			}
			// A simple hash to store the hexagon, or move to another place if
			// needed
			loc = libc.Int64FromUint64(*(*TH3Index)(unsafe.Pointer(bp + 56)) % libc.Uint64FromInt64(numHexagons))
			loopCount = 0
			for *(*TH3Index)(unsafe.Pointer(found + uintptr(loc)*8)) != uint64(0) {
				// If this conditional is reached, the `found` memory block is
				// too small for the given polygon. This should not happen.
				// TODO: Reachable via fuzzer
				if loopCount > numHexagons {
					return uint32(_E_FAILED)
				}
				if *(*TH3Index)(unsafe.Pointer(found + uintptr(loc)*8)) == *(*TH3Index)(unsafe.Pointer(bp + 56)) {
					break
				} // At least two points of the geoloop index to the
				// same cell
				loc = (loc + int64(1)) % numHexagons
				loopCount++
			}
			if *(*TH3Index)(unsafe.Pointer(found + uintptr(loc)*8)) == *(*TH3Index)(unsafe.Pointer(bp + 56)) {
				goto _3
			} // Skip this hex, already exists in the found hash
			// Otherwise, set it in the found hash for now
			*(*TH3Index)(unsafe.Pointer(found + uintptr(loc)*8)) = *(*TH3Index)(unsafe.Pointer(bp + 56))
			*(*TH3Index)(unsafe.Pointer(search + uintptr(*(*Tint64_t)(unsafe.Pointer(numSearchHexes)))*8)) = *(*TH3Index)(unsafe.Pointer(bp + 56))
			*(*Tint64_t)(unsafe.Pointer(numSearchHexes))++
			goto _3
		_3:
			;
			j++
		}
		goto _1
	_1:
		;
		i++
	}
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * polygonToCells takes a given GeoJSON-like data structure and preallocated,
//	 * zeroed memory, and fills it with the hexagons that are contained by
//	 * the GeoJSON-like data structure.
//	 *
//	 * This implementation traces the GeoJSON geoloop(s) in cartesian space with
//	 * hexagons, tests them and their neighbors to be contained by the geoloop(s),
//	 * and then any newly found hexagons are used to test again until no new
//	 * hexagons are found.
//	 *
//	 * @param geoPolygon The geoloop and holes defining the relevant area
//	 * @param res The Hexagon resolution (0-15)
//	 * @param out The slab of zeroed memory to write to. Assumed to be big enough.
//	 */
func XpolygonToCells(tls *libc.TLS, geoPolygon uintptr, res int32, flags Tuint32_t, out uintptr) (r TH3Error) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var bboxes, found, hole, search, temp uintptr
	var currentSearchNum, i1, i2, j1, loc, loopCount, numFoundHexes Tint64_t
	var edgeHexError, flagErr, numHexagonsError TH3Error
	var hex, searchHex TH3Index
	var i, j int32
	var _ /* geoloop at bp+16 */ TGeoLoop
	var _ /* hexCenter at bp+88 */ TLatLng
	var _ /* numHexagons at bp+0 */ Tint64_t
	var _ /* numSearchHexes at bp+8 */ Tint64_t
	var _ /* ring at bp+32 */ [7]TH3Index
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bboxes, currentSearchNum, edgeHexError, flagErr, found, hex, hole, i, i1, i2, j, j1, loc, loopCount, numFoundHexes, numHexagonsError, search, searchHex, temp
	flagErr = XvalidatePolygonFlags(tls, flags)
	if flagErr != 0 {
		return flagErr
	}
	// One of the goals of the polygonToCells algorithm is that two adjacent
	// polygons with zero overlap have zero overlapping hexagons. That the
	// hexagons are uniquely assigned. There are a few approaches to take here,
	// such as deciding based on which polygon has the greatest overlapping area
	// of the hexagon, or the most number of contained points on the hexagon
	// (using the center point as a tiebreaker).
	//
	// But if the polygons are convex, both of these more complex algorithms can
	// be reduced down to checking whether or not the center of the hexagon is
	// contained in the polygon, and so this is the approach that this
	// polygonToCells algorithm will follow, as it's simpler, faster, and the
	// error for concave polygons is still minimal (only affecting concave
	// shapes on the order of magnitude of the hexagon size or smaller, not
	// impacting larger concave shapes)
	//
	// This first part is identical to the maxPolygonToCellsSize above.
	// Get the bounding boxes for the polygon and any holes
	bboxes = libc.Xmalloc(tls, libc.Uint64FromInt32((*TGeoPolygon)(unsafe.Pointer(geoPolygon)).FnumHoles+libc.Int32FromInt32(1))*uint64(32))
	if !(bboxes != 0) {
		return uint32(_E_MEMORY_ALLOC)
	}
	XbboxesFromGeoPolygon(tls, geoPolygon, bboxes)
	numHexagonsError = XmaxPolygonToCellsSize(tls, geoPolygon, res, flags, bp)
	if numHexagonsError != 0 {
		libc.Xfree(tls, bboxes)
		return numHexagonsError
	}
	search = libc.Xcalloc(tls, libc.Uint64FromInt64(*(*Tint64_t)(unsafe.Pointer(bp))), uint64(8))
	if !(search != 0) {
		libc.Xfree(tls, bboxes)
		return uint32(_E_MEMORY_ALLOC)
	}
	found = libc.Xcalloc(tls, libc.Uint64FromInt64(*(*Tint64_t)(unsafe.Pointer(bp))), uint64(8))
	if !(found != 0) {
		libc.Xfree(tls, bboxes)
		libc.Xfree(tls, search)
		return uint32(_E_MEMORY_ALLOC)
	}
	// Some metadata for tracking the state of the search and found memory
	// blocks
	*(*Tint64_t)(unsafe.Pointer(bp + 8)) = 0
	numFoundHexes = 0
	// 1. Trace the hexagons along the polygon defining the outer geoloop and
	// add them to the search hash. The hexagon containing the geoloop point
	// may or may not be contained by the geoloop (as the hexagon's center
	// point may be outside of the boundary.)
	*(*TGeoLoop)(unsafe.Pointer(bp + 16)) = (*TGeoPolygon)(unsafe.Pointer(geoPolygon)).Fgeoloop
	edgeHexError = X_getEdgeHexagons(tls, bp+16, *(*Tint64_t)(unsafe.Pointer(bp)), res, bp+8, search, found)
	// If this branch is reached, we have exceeded the maximum number of
	// hexagons possible and need to clean up the allocated memory.
	// TODO: Reachable via fuzzer
	if edgeHexError != 0 {
		libc.Xfree(tls, search)
		libc.Xfree(tls, found)
		libc.Xfree(tls, bboxes)
		return edgeHexError
	}
	// 2. Iterate over all holes, trace the polygons defining the holes with
	// hexagons and add to only the search hash. We're going to temporarily use
	// the `found` hash to use for dedupe purposes and then re-zero it once
	// we're done here, otherwise we'd have to scan the whole set on each insert
	// to make sure there's no duplicates, which is very inefficient.
	i = 0
	for {
		if !(i < (*TGeoPolygon)(unsafe.Pointer(geoPolygon)).FnumHoles) {
			break
		}
		hole = (*TGeoPolygon)(unsafe.Pointer(geoPolygon)).Fholes + uintptr(i)*16
		edgeHexError = X_getEdgeHexagons(tls, hole, *(*Tint64_t)(unsafe.Pointer(bp)), res, bp+8, search, found)
		// If this branch is reached, we have exceeded the maximum number of
		// hexagons possible and need to clean up the allocated memory.
		// TODO: Reachable via fuzzer
		if edgeHexError != 0 {
			libc.Xfree(tls, search)
			libc.Xfree(tls, found)
			libc.Xfree(tls, bboxes)
			return edgeHexError
		}
		goto _1
	_1:
		;
		i++
	}
	// 3. Re-zero the found hash so it can be used in the main loop below
	i1 = 0
	for {
		if !(i1 < *(*Tint64_t)(unsafe.Pointer(bp))) {
			break
		}
		*(*TH3Index)(unsafe.Pointer(found + uintptr(i1)*8)) = uint64(DH3_NULL)
		goto _2
	_2:
		;
		i1++
	}
	// 4. Begin main loop. While the search hash is not empty do the following
	for *(*Tint64_t)(unsafe.Pointer(bp + 8)) > 0 {
		// Iterate through all hexagons in the current search hash, then loop
		// through all neighbors and test Point-in-Poly, if point-in-poly
		// succeeds, add to out and found hashes if not already there.
		currentSearchNum = 0
		i2 = 0
		for currentSearchNum < *(*Tint64_t)(unsafe.Pointer(bp + 8)) {
			*(*[7]TH3Index)(unsafe.Pointer(bp + 32)) = [7]TH3Index{}
			searchHex = *(*TH3Index)(unsafe.Pointer(search + uintptr(i2)*8))
			XgridDisk(tls, searchHex, int32(1), bp+32)
			j = 0
			for {
				if !(j < int32(DMAX_ONE_RING_SIZE)) {
					break
				}
				if (*(*[7]TH3Index)(unsafe.Pointer(bp + 32)))[j] == uint64(DH3_NULL) {
					goto _3 // Skip if this was a pentagon and only had 5
					// neighbors
				}
				hex = (*(*[7]TH3Index)(unsafe.Pointer(bp + 32)))[j]
				// A simple hash to store the hexagon, or move to another place
				// if needed. This MUST be done before the point-in-poly check
				// since that's far more expensive
				loc = libc.Int64FromUint64(hex % libc.Uint64FromInt64(*(*Tint64_t)(unsafe.Pointer(bp))))
				loopCount = 0
				for *(*TH3Index)(unsafe.Pointer(out + uintptr(loc)*8)) != uint64(0) {
					// If this branch is reached, we have exceeded the maximum
					// number of hexagons possible and need to clean up the
					// allocated memory.
					// TODO: Reachable via fuzzer
					if loopCount > *(*Tint64_t)(unsafe.Pointer(bp)) {
						libc.Xfree(tls, search)
						libc.Xfree(tls, found)
						libc.Xfree(tls, bboxes)
						return uint32(_E_FAILED)
					}
					if *(*TH3Index)(unsafe.Pointer(out + uintptr(loc)*8)) == hex {
						break
					} // Skip duplicates found
					loc = (loc + int64(1)) % *(*Tint64_t)(unsafe.Pointer(bp))
					loopCount++
				}
				if *(*TH3Index)(unsafe.Pointer(out + uintptr(loc)*8)) == hex {
					goto _3 // Skip this hex, already exists in the out hash
				}
				XcellToLatLng(tls, hex, bp+88)
				// If not, skip
				if !(XpointInsidePolygon(tls, geoPolygon, bboxes, bp+88) != 0) {
					goto _3
				}
				// Otherwise set it in the output array
				*(*TH3Index)(unsafe.Pointer(out + uintptr(loc)*8)) = hex
				// Set the hexagon in the found hash
				*(*TH3Index)(unsafe.Pointer(found + uintptr(numFoundHexes)*8)) = hex
				numFoundHexes++
				goto _3
			_3:
				;
				j++
			}
			currentSearchNum++
			i2++
		}
		// Swap the search and found pointers, copy the found hex count to the
		// search hex count, and zero everything related to the found memory.
		temp = search
		search = found
		found = temp
		j1 = 0
		for {
			if !(j1 < *(*Tint64_t)(unsafe.Pointer(bp + 8))) {
				break
			}
			*(*TH3Index)(unsafe.Pointer(found + uintptr(j1)*8)) = uint64(0)
			goto _4
		_4:
			;
			j1++
		}
		*(*Tint64_t)(unsafe.Pointer(bp + 8)) = numFoundHexes
		numFoundHexes = 0
		// Repeat until no new hexagons are found
	}
	// The out memory structure should be complete, end it here
	libc.Xfree(tls, bboxes)
	libc.Xfree(tls, search)
	libc.Xfree(tls, found)
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * Internal: Create a vertex graph from a set of hexagons. It is the
//	 * responsibility of the caller to call destroyVertexGraph on the populated
//	 * graph, otherwise the memory in the graph nodes will not be freed.
//	 * @private
//	 * @param h3Set    Set of hexagons
//	 * @param numHexes Number of hexagons in the set
//	 * @param graph    Output graph
//	 */
func Xh3SetToVertexGraph(tls *libc.TLS, h3Set uintptr, numHexes int32, graph uintptr) (r TH3Error) {
	bp := tls.Alloc(176)
	defer tls.Free(176)
	var boundaryErr TH3Error
	var edge, fromVtx, toVtx uintptr
	var i, j, minBuckets, numBuckets, res, v1 int32
	var _ /* vertices at bp+0 */ TCellBoundary
	_, _, _, _, _, _, _, _, _, _ = boundaryErr, edge, fromVtx, i, j, minBuckets, numBuckets, res, toVtx, v1
	if numHexes < int32(1) {
		// We still need to init the graph, or calls to destroyVertexGraph will
		// fail
		XinitVertexGraph(tls, graph, 0, 0)
		return uint32(_E_SUCCESS)
	}
	res = libc.Int32FromUint64(*(*TH3Index)(unsafe.Pointer(h3Set)) & (libc.Uint64FromUint64(15) << libc.Int32FromInt32(DH3_RES_OFFSET)) >> libc.Int32FromInt32(DH3_RES_OFFSET))
	minBuckets = int32(6)
	if numHexes > minBuckets {
		v1 = numHexes
	} else {
		v1 = minBuckets
	}
	// TODO: Better way to calculate/guess?
	numBuckets = v1
	XinitVertexGraph(tls, graph, numBuckets, res)
	// Iterate through every hexagon
	i = 0
	for {
		if !(i < numHexes) {
			break
		}
		boundaryErr = XcellToBoundary(tls, *(*TH3Index)(unsafe.Pointer(h3Set + uintptr(i)*8)), bp)
		if boundaryErr != 0 {
			// Destroy vertex graph as caller will not know to do so.
			XdestroyVertexGraph(tls, graph)
			return boundaryErr
		}
		// iterate through every edge
		j = 0
		for {
			if !(j < (*(*TCellBoundary)(unsafe.Pointer(bp))).FnumVerts) {
				break
			}
			fromVtx = bp + 8 + uintptr(j)*16
			toVtx = bp + 8 + uintptr((j+int32(1))%(*(*TCellBoundary)(unsafe.Pointer(bp))).FnumVerts)*16
			// If we've seen this edge already, it will be reversed
			edge = XfindNodeForEdge(tls, graph, toVtx, fromVtx)
			if edge != libc.UintptrFromInt32(0) {
				// If we've seen it, drop it. No edge is shared by more than 2
				// hexagons, so we'll never see it again.
				XremoveVertexNode(tls, graph, edge)
			} else {
				// Add a new node for this edge
				XaddVertexNode(tls, graph, fromVtx, toVtx)
			}
			goto _3
		_3:
			;
			j++
		}
		goto _2
	_2:
		;
		i++
	}
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * Internal: Create a LinkedGeoPolygon from a vertex graph. It is the
//	 * responsibility of the caller to call destroyLinkedMultiPolygon on the
//	 * populated linked geo structure, or the memory for that structure will not be
//	 * freed.
//	 * @private
//	 * @param graph Input graph
//	 * @param out   Output polygon
//	 */
func X_vertexGraphToLinkedGeo(tls *libc.TLS, graph uintptr, out uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var edge, loop, v1 uintptr
	var _ /* nextVtx at bp+0 */ TLatLng
	_, _, _ = edge, loop, v1
	*(*TLinkedGeoPolygon)(unsafe.Pointer(out)) = TLinkedGeoPolygon{}
	// Find the next unused entry point
	for {
		v1 = XfirstVertexNode(tls, graph)
		edge = v1
		if !(v1 != libc.UintptrFromInt32(0)) {
			break
		}
		loop = XaddNewLinkedLoop(tls, out)
		// Walk the graph to get the outline
		for cond := true; cond; cond = edge != 0 {
			XaddLinkedCoord(tls, loop, edge)
			*(*TLatLng)(unsafe.Pointer(bp)) = (*TVertexNode)(unsafe.Pointer(edge)).Fto
			// Remove frees the node, so we can't use edge after this
			XremoveVertexNode(tls, graph, edge)
			edge = XfindNodeForVertex(tls, graph, bp)
		}
	}
}

// C documentation
//
//	/**
//	 * Create a LinkedGeoPolygon describing the outline(s) of a set of  hexagons.
//	 * Polygon outlines will follow GeoJSON MultiPolygon order: Each polygon will
//	 * have one outer loop, which is first in the list, followed by any holes.
//	 *
//	 * It is the responsibility of the caller to call destroyLinkedMultiPolygon on
//	 * the populated linked geo structure, or the memory for that structure will not
//	 * be freed.
//	 *
//	 * It is expected that all hexagons in the set have the same resolution and
//	 * that the set contains no duplicates. Behavior is undefined if duplicates
//	 * or multiple resolutions are present, and the algorithm may produce
//	 * unexpected or invalid output.
//	 *
//	 * @param h3Set    Set of hexagons
//	 * @param numHexes Number of hexagons in set
//	 * @param out      Output polygon
//	 */
func XcellsToLinkedMultiPolygon(tls *libc.TLS, h3Set uintptr, numHexes int32, out uintptr) (r TH3Error) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var err, normalizeResult TH3Error
	var _ /* graph at bp+0 */ TVertexGraph
	_, _ = err, normalizeResult
	err = Xh3SetToVertexGraph(tls, h3Set, numHexes, bp)
	if err != 0 {
		return err
	}
	X_vertexGraphToLinkedGeo(tls, bp, out)
	XdestroyVertexGraph(tls, bp)
	normalizeResult = XnormalizeMultiPolygon(tls, out)
	if normalizeResult != 0 {
		XdestroyLinkedMultiPolygon(tls, out)
	}
	return normalizeResult
}

const DM_PI_2 = 1.5707963267948966

var _UNIT_VECS1 = [7]TCoordIJK{
	0: {},
	1: {
		Fk: int32(1),
	},
	2: {
		Fj: int32(1),
	},
	3: {
		Fj: int32(1),
		Fk: int32(1),
	},
	4: {
		Fi: int32(1),
	},
	5: {
		Fi: int32(1),
		Fk: int32(1),
	},
	6: {
		Fi: int32(1),
		Fj: int32(1),
	},
}

// C documentation
//
//	/** @struct BaseCellRotation
//	 *  @brief base cell at a given ijk and required rotations into its system
//	 */
type TBaseCellRotation = struct {
	FbaseCell int32
	FccwRot60 int32
}

// C documentation
//
//	/** @brief Resolution 0 base cell lookup table for each face.
//	 *
//	 * Given the face number and a resolution 0 ijk+ coordinate in that face's
//	 * face-centered ijk coordinate system, gives the base cell located at that
//	 * coordinate and the number of 60 ccw rotations to rotate into that base
//	 * cell's orientation.
//	 *
//	 * Valid lookup coordinates are from (0, 0, 0) to (2, 2, 2).
//	 *
//	 * This table can be accessed using the functions `_faceIjkToBaseCell` and
//	 * `_faceIjkToBaseCellCCWrot60`
//	 */
var _faceIjkBaseCells = [20][3][3][3]TBaseCellRotation{
	0: {
		0: {
			0: {
				0: {
					FbaseCell: int32(16),
				},
				1: {
					FbaseCell: int32(18),
				},
				2: {
					FbaseCell: int32(24),
				},
			},
			1: {
				0: {
					FbaseCell: int32(33),
				},
				1: {
					FbaseCell: int32(30),
				},
				2: {
					FbaseCell: int32(32),
					FccwRot60: int32(3),
				},
			},
			2: {
				0: {
					FbaseCell: int32(49),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(48),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(50),
					FccwRot60: int32(3),
				},
			},
		},
		1: {
			0: {
				0: {
					FbaseCell: int32(8),
				},
				1: {
					FbaseCell: int32(5),
					FccwRot60: int32(5),
				},
				2: {
					FbaseCell: int32(10),
					FccwRot60: int32(5),
				},
			},
			1: {
				0: {
					FbaseCell: int32(22),
				},
				1: {
					FbaseCell: int32(16),
				},
				2: {
					FbaseCell: int32(18),
				},
			},
			2: {
				0: {
					FbaseCell: int32(41),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(33),
				},
				2: {
					FbaseCell: int32(30),
				},
			},
		},
		2: {
			0: {
				0: {
					FbaseCell: int32(4),
				},
				1: {
					FccwRot60: int32(5),
				},
				2: {
					FbaseCell: int32(2),
					FccwRot60: int32(5),
				},
			},
			1: {
				0: {
					FbaseCell: int32(15),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(8),
				},
				2: {
					FbaseCell: int32(5),
					FccwRot60: int32(5),
				},
			},
			2: {
				0: {
					FbaseCell: int32(31),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(22),
				},
				2: {
					FbaseCell: int32(16),
				},
			},
		},
	},
	1: {
		0: {
			0: {
				0: {
					FbaseCell: int32(2),
				},
				1: {
					FbaseCell: int32(6),
				},
				2: {
					FbaseCell: int32(14),
				},
			},
			1: {
				0: {
					FbaseCell: int32(10),
				},
				1: {
					FbaseCell: int32(11),
				},
				2: {
					FbaseCell: int32(17),
					FccwRot60: int32(3),
				},
			},
			2: {
				0: {
					FbaseCell: int32(24),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(23),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(25),
					FccwRot60: int32(3),
				},
			},
		},
		1: {
			0: {
				0: {},
				1: {
					FbaseCell: int32(1),
					FccwRot60: int32(5),
				},
				2: {
					FbaseCell: int32(9),
					FccwRot60: int32(5),
				},
			},
			1: {
				0: {
					FbaseCell: int32(5),
				},
				1: {
					FbaseCell: int32(2),
				},
				2: {
					FbaseCell: int32(6),
				},
			},
			2: {
				0: {
					FbaseCell: int32(18),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(10),
				},
				2: {
					FbaseCell: int32(11),
				},
			},
		},
		2: {
			0: {
				0: {
					FbaseCell: int32(4),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(3),
					FccwRot60: int32(5),
				},
				2: {
					FbaseCell: int32(7),
					FccwRot60: int32(5),
				},
			},
			1: {
				0: {
					FbaseCell: int32(8),
					FccwRot60: int32(1),
				},
				1: {},
				2: {
					FbaseCell: int32(1),
					FccwRot60: int32(5),
				},
			},
			2: {
				0: {
					FbaseCell: int32(16),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(5),
				},
				2: {
					FbaseCell: int32(2),
				},
			},
		},
	},
	2: {
		0: {
			0: {
				0: {
					FbaseCell: int32(7),
				},
				1: {
					FbaseCell: int32(21),
				},
				2: {
					FbaseCell: int32(38),
				},
			},
			1: {
				0: {
					FbaseCell: int32(9),
				},
				1: {
					FbaseCell: int32(19),
				},
				2: {
					FbaseCell: int32(34),
					FccwRot60: int32(3),
				},
			},
			2: {
				0: {
					FbaseCell: int32(14),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(20),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(36),
					FccwRot60: int32(3),
				},
			},
		},
		1: {
			0: {
				0: {
					FbaseCell: int32(3),
				},
				1: {
					FbaseCell: int32(13),
					FccwRot60: int32(5),
				},
				2: {
					FbaseCell: int32(29),
					FccwRot60: int32(5),
				},
			},
			1: {
				0: {
					FbaseCell: int32(1),
				},
				1: {
					FbaseCell: int32(7),
				},
				2: {
					FbaseCell: int32(21),
				},
			},
			2: {
				0: {
					FbaseCell: int32(6),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(9),
				},
				2: {
					FbaseCell: int32(19),
				},
			},
		},
		2: {
			0: {
				0: {
					FbaseCell: int32(4),
					FccwRot60: int32(2),
				},
				1: {
					FbaseCell: int32(12),
					FccwRot60: int32(5),
				},
				2: {
					FbaseCell: int32(26),
					FccwRot60: int32(5),
				},
			},
			1: {
				0: {
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(3),
				},
				2: {
					FbaseCell: int32(13),
					FccwRot60: int32(5),
				},
			},
			2: {
				0: {
					FbaseCell: int32(2),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(1),
				},
				2: {
					FbaseCell: int32(7),
				},
			},
		},
	},
	3: {
		0: {
			0: {
				0: {
					FbaseCell: int32(26),
				},
				1: {
					FbaseCell: int32(42),
				},
				2: {
					FbaseCell: int32(58),
				},
			},
			1: {
				0: {
					FbaseCell: int32(29),
				},
				1: {
					FbaseCell: int32(43),
				},
				2: {
					FbaseCell: int32(62),
					FccwRot60: int32(3),
				},
			},
			2: {
				0: {
					FbaseCell: int32(38),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(47),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(64),
					FccwRot60: int32(3),
				},
			},
		},
		1: {
			0: {
				0: {
					FbaseCell: int32(12),
				},
				1: {
					FbaseCell: int32(28),
					FccwRot60: int32(5),
				},
				2: {
					FbaseCell: int32(44),
					FccwRot60: int32(5),
				},
			},
			1: {
				0: {
					FbaseCell: int32(13),
				},
				1: {
					FbaseCell: int32(26),
				},
				2: {
					FbaseCell: int32(42),
				},
			},
			2: {
				0: {
					FbaseCell: int32(21),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(29),
				},
				2: {
					FbaseCell: int32(43),
				},
			},
		},
		2: {
			0: {
				0: {
					FbaseCell: int32(4),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(15),
					FccwRot60: int32(5),
				},
				2: {
					FbaseCell: int32(31),
					FccwRot60: int32(5),
				},
			},
			1: {
				0: {
					FbaseCell: int32(3),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(12),
				},
				2: {
					FbaseCell: int32(28),
					FccwRot60: int32(5),
				},
			},
			2: {
				0: {
					FbaseCell: int32(7),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(13),
				},
				2: {
					FbaseCell: int32(26),
				},
			},
		},
	},
	4: {
		0: {
			0: {
				0: {
					FbaseCell: int32(31),
				},
				1: {
					FbaseCell: int32(41),
				},
				2: {
					FbaseCell: int32(49),
				},
			},
			1: {
				0: {
					FbaseCell: int32(44),
				},
				1: {
					FbaseCell: int32(53),
				},
				2: {
					FbaseCell: int32(61),
					FccwRot60: int32(3),
				},
			},
			2: {
				0: {
					FbaseCell: int32(58),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(65),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(75),
					FccwRot60: int32(3),
				},
			},
		},
		1: {
			0: {
				0: {
					FbaseCell: int32(15),
				},
				1: {
					FbaseCell: int32(22),
					FccwRot60: int32(5),
				},
				2: {
					FbaseCell: int32(33),
					FccwRot60: int32(5),
				},
			},
			1: {
				0: {
					FbaseCell: int32(28),
				},
				1: {
					FbaseCell: int32(31),
				},
				2: {
					FbaseCell: int32(41),
				},
			},
			2: {
				0: {
					FbaseCell: int32(42),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(44),
				},
				2: {
					FbaseCell: int32(53),
				},
			},
		},
		2: {
			0: {
				0: {
					FbaseCell: int32(4),
					FccwRot60: int32(4),
				},
				1: {
					FbaseCell: int32(8),
					FccwRot60: int32(5),
				},
				2: {
					FbaseCell: int32(16),
					FccwRot60: int32(5),
				},
			},
			1: {
				0: {
					FbaseCell: int32(12),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(15),
				},
				2: {
					FbaseCell: int32(22),
					FccwRot60: int32(5),
				},
			},
			2: {
				0: {
					FbaseCell: int32(26),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(28),
				},
				2: {
					FbaseCell: int32(31),
				},
			},
		},
	},
	5: {
		0: {
			0: {
				0: {
					FbaseCell: int32(50),
				},
				1: {
					FbaseCell: int32(48),
				},
				2: {
					FbaseCell: int32(49),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(32),
				},
				1: {
					FbaseCell: int32(30),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(33),
					FccwRot60: int32(3),
				},
			},
			2: {
				0: {
					FbaseCell: int32(24),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(18),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(16),
					FccwRot60: int32(3),
				},
			},
		},
		1: {
			0: {
				0: {
					FbaseCell: int32(70),
				},
				1: {
					FbaseCell: int32(67),
				},
				2: {
					FbaseCell: int32(66),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(52),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(50),
				},
				2: {
					FbaseCell: int32(48),
				},
			},
			2: {
				0: {
					FbaseCell: int32(37),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(32),
				},
				2: {
					FbaseCell: int32(30),
					FccwRot60: int32(3),
				},
			},
		},
		2: {
			0: {
				0: {
					FbaseCell: int32(83),
				},
				1: {
					FbaseCell: int32(87),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(85),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(74),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(70),
				},
				2: {
					FbaseCell: int32(67),
				},
			},
			2: {
				0: {
					FbaseCell: int32(57),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(52),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(50),
				},
			},
		},
	},
	6: {
		0: {
			0: {
				0: {
					FbaseCell: int32(25),
				},
				1: {
					FbaseCell: int32(23),
				},
				2: {
					FbaseCell: int32(24),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(17),
				},
				1: {
					FbaseCell: int32(11),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(10),
					FccwRot60: int32(3),
				},
			},
			2: {
				0: {
					FbaseCell: int32(14),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(6),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(2),
					FccwRot60: int32(3),
				},
			},
		},
		1: {
			0: {
				0: {
					FbaseCell: int32(45),
				},
				1: {
					FbaseCell: int32(39),
				},
				2: {
					FbaseCell: int32(37),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(35),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(25),
				},
				2: {
					FbaseCell: int32(23),
				},
			},
			2: {
				0: {
					FbaseCell: int32(27),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(17),
				},
				2: {
					FbaseCell: int32(11),
					FccwRot60: int32(3),
				},
			},
		},
		2: {
			0: {
				0: {
					FbaseCell: int32(63),
				},
				1: {
					FbaseCell: int32(59),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(57),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(56),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(45),
				},
				2: {
					FbaseCell: int32(39),
				},
			},
			2: {
				0: {
					FbaseCell: int32(46),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(35),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(25),
				},
			},
		},
	},
	7: {
		0: {
			0: {
				0: {
					FbaseCell: int32(36),
				},
				1: {
					FbaseCell: int32(20),
				},
				2: {
					FbaseCell: int32(14),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(34),
				},
				1: {
					FbaseCell: int32(19),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(9),
					FccwRot60: int32(3),
				},
			},
			2: {
				0: {
					FbaseCell: int32(38),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(21),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(7),
					FccwRot60: int32(3),
				},
			},
		},
		1: {
			0: {
				0: {
					FbaseCell: int32(55),
				},
				1: {
					FbaseCell: int32(40),
				},
				2: {
					FbaseCell: int32(27),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(54),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(36),
				},
				2: {
					FbaseCell: int32(20),
				},
			},
			2: {
				0: {
					FbaseCell: int32(51),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(34),
				},
				2: {
					FbaseCell: int32(19),
					FccwRot60: int32(3),
				},
			},
		},
		2: {
			0: {
				0: {
					FbaseCell: int32(72),
				},
				1: {
					FbaseCell: int32(60),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(46),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(73),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(55),
				},
				2: {
					FbaseCell: int32(40),
				},
			},
			2: {
				0: {
					FbaseCell: int32(71),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(54),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(36),
				},
			},
		},
	},
	8: {
		0: {
			0: {
				0: {
					FbaseCell: int32(64),
				},
				1: {
					FbaseCell: int32(47),
				},
				2: {
					FbaseCell: int32(38),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(62),
				},
				1: {
					FbaseCell: int32(43),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(29),
					FccwRot60: int32(3),
				},
			},
			2: {
				0: {
					FbaseCell: int32(58),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(42),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(26),
					FccwRot60: int32(3),
				},
			},
		},
		1: {
			0: {
				0: {
					FbaseCell: int32(84),
				},
				1: {
					FbaseCell: int32(69),
				},
				2: {
					FbaseCell: int32(51),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(82),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(64),
				},
				2: {
					FbaseCell: int32(47),
				},
			},
			2: {
				0: {
					FbaseCell: int32(76),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(62),
				},
				2: {
					FbaseCell: int32(43),
					FccwRot60: int32(3),
				},
			},
		},
		2: {
			0: {
				0: {
					FbaseCell: int32(97),
				},
				1: {
					FbaseCell: int32(89),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(71),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(98),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(84),
				},
				2: {
					FbaseCell: int32(69),
				},
			},
			2: {
				0: {
					FbaseCell: int32(96),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(82),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(64),
				},
			},
		},
	},
	9: {
		0: {
			0: {
				0: {
					FbaseCell: int32(75),
				},
				1: {
					FbaseCell: int32(65),
				},
				2: {
					FbaseCell: int32(58),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(61),
				},
				1: {
					FbaseCell: int32(53),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(44),
					FccwRot60: int32(3),
				},
			},
			2: {
				0: {
					FbaseCell: int32(49),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(41),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(31),
					FccwRot60: int32(3),
				},
			},
		},
		1: {
			0: {
				0: {
					FbaseCell: int32(94),
				},
				1: {
					FbaseCell: int32(86),
				},
				2: {
					FbaseCell: int32(76),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(81),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(75),
				},
				2: {
					FbaseCell: int32(65),
				},
			},
			2: {
				0: {
					FbaseCell: int32(66),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(61),
				},
				2: {
					FbaseCell: int32(53),
					FccwRot60: int32(3),
				},
			},
		},
		2: {
			0: {
				0: {
					FbaseCell: int32(107),
				},
				1: {
					FbaseCell: int32(104),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(96),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(101),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(94),
				},
				2: {
					FbaseCell: int32(86),
				},
			},
			2: {
				0: {
					FbaseCell: int32(85),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(81),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(75),
				},
			},
		},
	},
	10: {
		0: {
			0: {
				0: {
					FbaseCell: int32(57),
				},
				1: {
					FbaseCell: int32(59),
				},
				2: {
					FbaseCell: int32(63),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(74),
				},
				1: {
					FbaseCell: int32(78),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(79),
					FccwRot60: int32(3),
				},
			},
			2: {
				0: {
					FbaseCell: int32(83),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(92),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(95),
					FccwRot60: int32(3),
				},
			},
		},
		1: {
			0: {
				0: {
					FbaseCell: int32(37),
				},
				1: {
					FbaseCell: int32(39),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(45),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(52),
				},
				1: {
					FbaseCell: int32(57),
				},
				2: {
					FbaseCell: int32(59),
				},
			},
			2: {
				0: {
					FbaseCell: int32(70),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(74),
				},
				2: {
					FbaseCell: int32(78),
					FccwRot60: int32(3),
				},
			},
		},
		2: {
			0: {
				0: {
					FbaseCell: int32(24),
				},
				1: {
					FbaseCell: int32(23),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(25),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(32),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(37),
				},
				2: {
					FbaseCell: int32(39),
					FccwRot60: int32(3),
				},
			},
			2: {
				0: {
					FbaseCell: int32(50),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(52),
				},
				2: {
					FbaseCell: int32(57),
				},
			},
		},
	},
	11: {
		0: {
			0: {
				0: {
					FbaseCell: int32(46),
				},
				1: {
					FbaseCell: int32(60),
				},
				2: {
					FbaseCell: int32(72),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(56),
				},
				1: {
					FbaseCell: int32(68),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(80),
					FccwRot60: int32(3),
				},
			},
			2: {
				0: {
					FbaseCell: int32(63),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(77),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(90),
					FccwRot60: int32(3),
				},
			},
		},
		1: {
			0: {
				0: {
					FbaseCell: int32(27),
				},
				1: {
					FbaseCell: int32(40),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(55),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(35),
				},
				1: {
					FbaseCell: int32(46),
				},
				2: {
					FbaseCell: int32(60),
				},
			},
			2: {
				0: {
					FbaseCell: int32(45),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(56),
				},
				2: {
					FbaseCell: int32(68),
					FccwRot60: int32(3),
				},
			},
		},
		2: {
			0: {
				0: {
					FbaseCell: int32(14),
				},
				1: {
					FbaseCell: int32(20),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(36),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(17),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(27),
				},
				2: {
					FbaseCell: int32(40),
					FccwRot60: int32(3),
				},
			},
			2: {
				0: {
					FbaseCell: int32(25),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(35),
				},
				2: {
					FbaseCell: int32(46),
				},
			},
		},
	},
	12: {
		0: {
			0: {
				0: {
					FbaseCell: int32(71),
				},
				1: {
					FbaseCell: int32(89),
				},
				2: {
					FbaseCell: int32(97),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(73),
				},
				1: {
					FbaseCell: int32(91),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(103),
					FccwRot60: int32(3),
				},
			},
			2: {
				0: {
					FbaseCell: int32(72),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(88),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(105),
					FccwRot60: int32(3),
				},
			},
		},
		1: {
			0: {
				0: {
					FbaseCell: int32(51),
				},
				1: {
					FbaseCell: int32(69),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(84),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(54),
				},
				1: {
					FbaseCell: int32(71),
				},
				2: {
					FbaseCell: int32(89),
				},
			},
			2: {
				0: {
					FbaseCell: int32(55),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(73),
				},
				2: {
					FbaseCell: int32(91),
					FccwRot60: int32(3),
				},
			},
		},
		2: {
			0: {
				0: {
					FbaseCell: int32(38),
				},
				1: {
					FbaseCell: int32(47),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(64),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(34),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(51),
				},
				2: {
					FbaseCell: int32(69),
					FccwRot60: int32(3),
				},
			},
			2: {
				0: {
					FbaseCell: int32(36),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(54),
				},
				2: {
					FbaseCell: int32(71),
				},
			},
		},
	},
	13: {
		0: {
			0: {
				0: {
					FbaseCell: int32(96),
				},
				1: {
					FbaseCell: int32(104),
				},
				2: {
					FbaseCell: int32(107),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(98),
				},
				1: {
					FbaseCell: int32(110),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(115),
					FccwRot60: int32(3),
				},
			},
			2: {
				0: {
					FbaseCell: int32(97),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(111),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(119),
					FccwRot60: int32(3),
				},
			},
		},
		1: {
			0: {
				0: {
					FbaseCell: int32(76),
				},
				1: {
					FbaseCell: int32(86),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(94),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(82),
				},
				1: {
					FbaseCell: int32(96),
				},
				2: {
					FbaseCell: int32(104),
				},
			},
			2: {
				0: {
					FbaseCell: int32(84),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(98),
				},
				2: {
					FbaseCell: int32(110),
					FccwRot60: int32(3),
				},
			},
		},
		2: {
			0: {
				0: {
					FbaseCell: int32(58),
				},
				1: {
					FbaseCell: int32(65),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(75),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(62),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(76),
				},
				2: {
					FbaseCell: int32(86),
					FccwRot60: int32(3),
				},
			},
			2: {
				0: {
					FbaseCell: int32(64),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(82),
				},
				2: {
					FbaseCell: int32(96),
				},
			},
		},
	},
	14: {
		0: {
			0: {
				0: {
					FbaseCell: int32(85),
				},
				1: {
					FbaseCell: int32(87),
				},
				2: {
					FbaseCell: int32(83),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(101),
				},
				1: {
					FbaseCell: int32(102),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(100),
					FccwRot60: int32(3),
				},
			},
			2: {
				0: {
					FbaseCell: int32(107),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(112),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(114),
					FccwRot60: int32(3),
				},
			},
		},
		1: {
			0: {
				0: {
					FbaseCell: int32(66),
				},
				1: {
					FbaseCell: int32(67),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(70),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(81),
				},
				1: {
					FbaseCell: int32(85),
				},
				2: {
					FbaseCell: int32(87),
				},
			},
			2: {
				0: {
					FbaseCell: int32(94),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(101),
				},
				2: {
					FbaseCell: int32(102),
					FccwRot60: int32(3),
				},
			},
		},
		2: {
			0: {
				0: {
					FbaseCell: int32(49),
				},
				1: {
					FbaseCell: int32(48),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(50),
					FccwRot60: int32(3),
				},
			},
			1: {
				0: {
					FbaseCell: int32(61),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(66),
				},
				2: {
					FbaseCell: int32(67),
					FccwRot60: int32(3),
				},
			},
			2: {
				0: {
					FbaseCell: int32(75),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(81),
				},
				2: {
					FbaseCell: int32(85),
				},
			},
		},
	},
	15: {
		0: {
			0: {
				0: {
					FbaseCell: int32(95),
				},
				1: {
					FbaseCell: int32(92),
				},
				2: {
					FbaseCell: int32(83),
				},
			},
			1: {
				0: {
					FbaseCell: int32(79),
				},
				1: {
					FbaseCell: int32(78),
				},
				2: {
					FbaseCell: int32(74),
					FccwRot60: int32(3),
				},
			},
			2: {
				0: {
					FbaseCell: int32(63),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(59),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(57),
					FccwRot60: int32(3),
				},
			},
		},
		1: {
			0: {
				0: {
					FbaseCell: int32(109),
				},
				1: {
					FbaseCell: int32(108),
				},
				2: {
					FbaseCell: int32(100),
					FccwRot60: int32(5),
				},
			},
			1: {
				0: {
					FbaseCell: int32(93),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(95),
				},
				2: {
					FbaseCell: int32(92),
				},
			},
			2: {
				0: {
					FbaseCell: int32(77),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(79),
				},
				2: {
					FbaseCell: int32(78),
				},
			},
		},
		2: {
			0: {
				0: {
					FbaseCell: int32(117),
					FccwRot60: int32(4),
				},
				1: {
					FbaseCell: int32(118),
					FccwRot60: int32(5),
				},
				2: {
					FbaseCell: int32(114),
					FccwRot60: int32(5),
				},
			},
			1: {
				0: {
					FbaseCell: int32(106),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(109),
				},
				2: {
					FbaseCell: int32(108),
				},
			},
			2: {
				0: {
					FbaseCell: int32(90),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(93),
					FccwRot60: int32(1),
				},
				2: {
					FbaseCell: int32(95),
				},
			},
		},
	},
	16: {
		0: {
			0: {
				0: {
					FbaseCell: int32(90),
				},
				1: {
					FbaseCell: int32(77),
				},
				2: {
					FbaseCell: int32(63),
				},
			},
			1: {
				0: {
					FbaseCell: int32(80),
				},
				1: {
					FbaseCell: int32(68),
				},
				2: {
					FbaseCell: int32(56),
					FccwRot60: int32(3),
				},
			},
			2: {
				0: {
					FbaseCell: int32(72),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(60),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(46),
					FccwRot60: int32(3),
				},
			},
		},
		1: {
			0: {
				0: {
					FbaseCell: int32(106),
				},
				1: {
					FbaseCell: int32(93),
				},
				2: {
					FbaseCell: int32(79),
					FccwRot60: int32(5),
				},
			},
			1: {
				0: {
					FbaseCell: int32(99),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(90),
				},
				2: {
					FbaseCell: int32(77),
				},
			},
			2: {
				0: {
					FbaseCell: int32(88),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(80),
				},
				2: {
					FbaseCell: int32(68),
				},
			},
		},
		2: {
			0: {
				0: {
					FbaseCell: int32(117),
					FccwRot60: int32(3),
				},
				1: {
					FbaseCell: int32(109),
					FccwRot60: int32(5),
				},
				2: {
					FbaseCell: int32(95),
					FccwRot60: int32(5),
				},
			},
			1: {
				0: {
					FbaseCell: int32(113),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(106),
				},
				2: {
					FbaseCell: int32(93),
				},
			},
			2: {
				0: {
					FbaseCell: int32(105),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(99),
					FccwRot60: int32(1),
				},
				2: {
					FbaseCell: int32(90),
				},
			},
		},
	},
	17: {
		0: {
			0: {
				0: {
					FbaseCell: int32(105),
				},
				1: {
					FbaseCell: int32(88),
				},
				2: {
					FbaseCell: int32(72),
				},
			},
			1: {
				0: {
					FbaseCell: int32(103),
				},
				1: {
					FbaseCell: int32(91),
				},
				2: {
					FbaseCell: int32(73),
					FccwRot60: int32(3),
				},
			},
			2: {
				0: {
					FbaseCell: int32(97),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(89),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(71),
					FccwRot60: int32(3),
				},
			},
		},
		1: {
			0: {
				0: {
					FbaseCell: int32(113),
				},
				1: {
					FbaseCell: int32(99),
				},
				2: {
					FbaseCell: int32(80),
					FccwRot60: int32(5),
				},
			},
			1: {
				0: {
					FbaseCell: int32(116),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(105),
				},
				2: {
					FbaseCell: int32(88),
				},
			},
			2: {
				0: {
					FbaseCell: int32(111),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(103),
				},
				2: {
					FbaseCell: int32(91),
				},
			},
		},
		2: {
			0: {
				0: {
					FbaseCell: int32(117),
					FccwRot60: int32(2),
				},
				1: {
					FbaseCell: int32(106),
					FccwRot60: int32(5),
				},
				2: {
					FbaseCell: int32(90),
					FccwRot60: int32(5),
				},
			},
			1: {
				0: {
					FbaseCell: int32(121),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(113),
				},
				2: {
					FbaseCell: int32(99),
				},
			},
			2: {
				0: {
					FbaseCell: int32(119),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(116),
					FccwRot60: int32(1),
				},
				2: {
					FbaseCell: int32(105),
				},
			},
		},
	},
	18: {
		0: {
			0: {
				0: {
					FbaseCell: int32(119),
				},
				1: {
					FbaseCell: int32(111),
				},
				2: {
					FbaseCell: int32(97),
				},
			},
			1: {
				0: {
					FbaseCell: int32(115),
				},
				1: {
					FbaseCell: int32(110),
				},
				2: {
					FbaseCell: int32(98),
					FccwRot60: int32(3),
				},
			},
			2: {
				0: {
					FbaseCell: int32(107),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(104),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(96),
					FccwRot60: int32(3),
				},
			},
		},
		1: {
			0: {
				0: {
					FbaseCell: int32(121),
				},
				1: {
					FbaseCell: int32(116),
				},
				2: {
					FbaseCell: int32(103),
					FccwRot60: int32(5),
				},
			},
			1: {
				0: {
					FbaseCell: int32(120),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(119),
				},
				2: {
					FbaseCell: int32(111),
				},
			},
			2: {
				0: {
					FbaseCell: int32(112),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(115),
				},
				2: {
					FbaseCell: int32(110),
				},
			},
		},
		2: {
			0: {
				0: {
					FbaseCell: int32(117),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(113),
					FccwRot60: int32(5),
				},
				2: {
					FbaseCell: int32(105),
					FccwRot60: int32(5),
				},
			},
			1: {
				0: {
					FbaseCell: int32(118),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(121),
				},
				2: {
					FbaseCell: int32(116),
				},
			},
			2: {
				0: {
					FbaseCell: int32(114),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(120),
					FccwRot60: int32(1),
				},
				2: {
					FbaseCell: int32(119),
				},
			},
		},
	},
	19: {
		0: {
			0: {
				0: {
					FbaseCell: int32(114),
				},
				1: {
					FbaseCell: int32(112),
				},
				2: {
					FbaseCell: int32(107),
				},
			},
			1: {
				0: {
					FbaseCell: int32(100),
				},
				1: {
					FbaseCell: int32(102),
				},
				2: {
					FbaseCell: int32(101),
					FccwRot60: int32(3),
				},
			},
			2: {
				0: {
					FbaseCell: int32(83),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(87),
					FccwRot60: int32(3),
				},
				2: {
					FbaseCell: int32(85),
					FccwRot60: int32(3),
				},
			},
		},
		1: {
			0: {
				0: {
					FbaseCell: int32(118),
				},
				1: {
					FbaseCell: int32(120),
				},
				2: {
					FbaseCell: int32(115),
					FccwRot60: int32(5),
				},
			},
			1: {
				0: {
					FbaseCell: int32(108),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(114),
				},
				2: {
					FbaseCell: int32(112),
				},
			},
			2: {
				0: {
					FbaseCell: int32(92),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(100),
				},
				2: {
					FbaseCell: int32(102),
				},
			},
		},
		2: {
			0: {
				0: {
					FbaseCell: int32(117),
				},
				1: {
					FbaseCell: int32(121),
					FccwRot60: int32(5),
				},
				2: {
					FbaseCell: int32(119),
					FccwRot60: int32(5),
				},
			},
			1: {
				0: {
					FbaseCell: int32(109),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(118),
				},
				2: {
					FbaseCell: int32(120),
				},
			},
			2: {
				0: {
					FbaseCell: int32(95),
					FccwRot60: int32(1),
				},
				1: {
					FbaseCell: int32(108),
					FccwRot60: int32(1),
				},
				2: {
					FbaseCell: int32(114),
				},
			},
		},
	},
}

// C documentation
//
//	/** @brief Return whether or not the indicated base cell is a pentagon. */
func X_isBaseCellPentagon(tls *libc.TLS, baseCell int32) (r int32) {
	if baseCell < 0 || baseCell >= int32(DNUM_BASE_CELLS) {
		// Base cells less than zero can not be represented in an index
		return Dfalse
	}
	return XbaseCellData[baseCell].FisPentagon
}

// C documentation
//
//	/** @brief Return whether the indicated base cell is a pentagon where all
//	 * neighbors are oriented towards it. */
func X_isBaseCellPolarPentagon(tls *libc.TLS, baseCell int32) (r uint8) {
	return libc.BoolUint8(baseCell == int32(4) || baseCell == int32(117))
}

// C documentation
//
//	/** @brief Find base cell given FaceIJK.
//	 *
//	 * Given the face number and a resolution 0 ijk+ coordinate in that face's
//	 * face-centered ijk coordinate system, return the base cell located at that
//	 * coordinate.
//	 *
//	 * Valid ijk+ lookup coordinates are from (0, 0, 0) to (2, 2, 2).
//	 */
func X_faceIjkToBaseCell(tls *libc.TLS, h uintptr) (r int32) {
	return (*(*TBaseCellRotation)(unsafe.Pointer(uintptr(unsafe.Pointer(&_faceIjkBaseCells)) + uintptr((*TFaceIJK)(unsafe.Pointer(h)).Fface)*216 + uintptr((*TFaceIJK)(unsafe.Pointer(h)).Fcoord.Fi)*72 + uintptr((*TFaceIJK)(unsafe.Pointer(h)).Fcoord.Fj)*24 + uintptr((*TFaceIJK)(unsafe.Pointer(h)).Fcoord.Fk)*8))).FbaseCell
}

// C documentation
//
//	/** @brief Find base cell given FaceIJK.
//	 *
//	 * Given the face number and a resolution 0 ijk+ coordinate in that face's
//	 * face-centered ijk coordinate system, return the number of 60' ccw rotations
//	 * to rotate into the coordinate system of the base cell at that coordinates.
//	 *
//	 * Valid ijk+ lookup coordinates are from (0, 0, 0) to (2, 2, 2).
//	 */
func X_faceIjkToBaseCellCCWrot60(tls *libc.TLS, h uintptr) (r int32) {
	return (*(*TBaseCellRotation)(unsafe.Pointer(uintptr(unsafe.Pointer(&_faceIjkBaseCells)) + uintptr((*TFaceIJK)(unsafe.Pointer(h)).Fface)*216 + uintptr((*TFaceIJK)(unsafe.Pointer(h)).Fcoord.Fi)*72 + uintptr((*TFaceIJK)(unsafe.Pointer(h)).Fcoord.Fj)*24 + uintptr((*TFaceIJK)(unsafe.Pointer(h)).Fcoord.Fk)*8))).FccwRot60
}

// C documentation
//
//	/** @brief Find the FaceIJK given a base cell.
//	 */
func X_baseCellToFaceIjk(tls *libc.TLS, baseCell int32, h uintptr) {
	*(*TFaceIJK)(unsafe.Pointer(h)) = XbaseCellData[baseCell].FhomeFijk
}

// C documentation
//
//	/**
//	 * @brief Given a base cell and the face it appears on, return
//	 *        the number of 60' ccw rotations for that base cell's
//	 *        coordinate system.
//	 * @returns The number of rotations, or INVALID_ROTATIONS if the base
//	 *          cell is not found on the given face
//	 */
func X_baseCellToCCWrot60(tls *libc.TLS, baseCell int32, face int32) (r int32) {
	var i, j, k int32
	_, _, _ = i, j, k
	if face < 0 || face > int32(DNUM_ICOSA_FACES) {
		return -int32(1)
	}
	i = 0
	for {
		if !(i < int32(3)) {
			break
		}
		j = 0
		for {
			if !(j < int32(3)) {
				break
			}
			k = 0
			for {
				if !(k < int32(3)) {
					break
				}
				if (*(*TBaseCellRotation)(unsafe.Pointer(uintptr(unsafe.Pointer(&_faceIjkBaseCells)) + uintptr(face)*216 + uintptr(i)*72 + uintptr(j)*24 + uintptr(k)*8))).FbaseCell == baseCell {
					return (*(*TBaseCellRotation)(unsafe.Pointer(uintptr(unsafe.Pointer(&_faceIjkBaseCells)) + uintptr(face)*216 + uintptr(i)*72 + uintptr(j)*24 + uintptr(k)*8))).FccwRot60
				}
				goto _3
			_3:
				;
				k++
			}
			goto _2
		_2:
			;
			j++
		}
		goto _1
	_1:
		;
		i++
	}
	return -int32(1)
}

// C documentation
//
//	/** @brief Return whether or not the tested face is a cw offset face.
//	 */
func X_baseCellIsCwOffset(tls *libc.TLS, baseCell int32, testFace int32) (r uint8) {
	return libc.BoolUint8(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&XbaseCellData)) + uintptr(baseCell)*28 + 20)) == testFace || *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&XbaseCellData)) + uintptr(baseCell)*28 + 20 + 1*4)) == testFace)
}

// C documentation
//
//	/** @brief Return the neighboring base cell in the given direction.
//	 */
func X_getBaseCellNeighbor(tls *libc.TLS, baseCell int32, dir TDirection) (r int32) {
	return *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&XbaseCellNeighbors)) + uintptr(baseCell)*28 + uintptr(dir)*4))
}

// C documentation
//
//	/** @brief Return the direction from the origin base cell to the neighbor.
//	 * Returns INVALID_DIGIT if the base cells are not neighbors.
//	 */
func X_getBaseCellDirection(tls *libc.TLS, originBaseCell int32, neighboringBaseCell int32) (r TDirection) {
	var dir TDirection
	var testBaseCell int32
	_, _ = dir, testBaseCell
	dir = int32(_CENTER_DIGIT)
	for {
		if !(dir < int32(_NUM_DIGITS)) {
			break
		}
		testBaseCell = X_getBaseCellNeighbor(tls, originBaseCell, dir)
		if testBaseCell == neighboringBaseCell {
			return dir
		}
		goto _1
	_1:
		;
		dir++
	}
	return int32(_INVALID_DIGIT)
}

// C documentation
//
//	/**
//	 * res0CellCount returns the number of resolution 0 cells
//	 *
//	 * @return int count of resolution 0 cells
//	 */
func Xres0CellCount(tls *libc.TLS) (r int32) {
	return int32(DNUM_BASE_CELLS)
}

// C documentation
//
//	/**
//	 * getRes0Cells generates all base cells storing them into the provided
//	 * memory pointer. Buffer must be of size NUM_BASE_CELLS * sizeof(H3Index).
//	 *
//	 * @param out H3Index* the memory to store the resulting base cells in
//	 * @returns E_SUCCESS.
//	 */
func XgetRes0Cells(tls *libc.TLS, out uintptr) (r TH3Error) {
	var baseCell TH3Index
	var bc int32
	_, _ = baseCell, bc
	bc = 0
	for {
		if !(bc < int32(DNUM_BASE_CELLS)) {
			break
		}
		baseCell = libc.Uint64FromUint64(35184372088831)
		baseCell = baseCell & ^(libc.Uint64FromInt32(libc.Int32FromInt32(15))<<libc.Int32FromInt32(DH3_MODE_OFFSET)) | libc.Uint64FromInt32(libc.Int32FromInt32(DH3_CELL_MODE))<<libc.Int32FromInt32(DH3_MODE_OFFSET)
		baseCell = baseCell & ^(libc.Uint64FromInt32(libc.Int32FromInt32(127))<<libc.Int32FromInt32(DH3_BC_OFFSET)) | libc.Uint64FromInt32(bc)<<libc.Int32FromInt32(DH3_BC_OFFSET)
		*(*TH3Index)(unsafe.Pointer(out + uintptr(bc)*8)) = baseCell
		goto _1
	_1:
		;
		bc++
	}
	return uint32(_E_SUCCESS)
}

const DM_2PI1 = 6.283185307179586
const DM_PI1 = 3.141592653589793

var _UNIT_VECS2 = [7]TCoordIJK{
	0: {},
	1: {
		Fk: int32(1),
	},
	2: {
		Fj: int32(1),
	},
	3: {
		Fj: int32(1),
		Fk: int32(1),
	},
	4: {
		Fi: int32(1),
	},
	5: {
		Fi: int32(1),
		Fk: int32(1),
	},
	6: {
		Fi: int32(1),
		Fj: int32(1),
	},
}

/*
 * Copyright 2016-2021 Uber Technologies, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/** @file latLng.h
 * @brief   Geodetic (lat/lng) functions.
 */

// C documentation
//
//	/**
//	 * Width of the bounding box, in rads
//	 * @param  bbox Bounding box to inspect
//	 * @return      width, in rads
//	 */
func XbboxWidthRads(tls *libc.TLS, bbox uintptr) (r float64) {
	var v1 float64
	_ = v1
	if XbboxIsTransmeridian(tls, bbox) != 0 {
		v1 = (*TBBox)(unsafe.Pointer(bbox)).Feast - (*TBBox)(unsafe.Pointer(bbox)).Fwest + float64(6.283185307179586)
	} else {
		v1 = (*TBBox)(unsafe.Pointer(bbox)).Feast - (*TBBox)(unsafe.Pointer(bbox)).Fwest
	}
	return v1
}

// C documentation
//
//	/**
//	 * Height of the bounding box
//	 * @param  bbox Bounding box to inspect
//	 * @return      height, in rads
//	 */
func XbboxHeightRads(tls *libc.TLS, bbox uintptr) (r float64) {
	return (*TBBox)(unsafe.Pointer(bbox)).Fnorth - (*TBBox)(unsafe.Pointer(bbox)).Fsouth
}

// C documentation
//
//	/**
//	 * Whether the given bounding box crosses the antimeridian
//	 * @param  bbox Bounding box to inspect
//	 * @return      is transmeridian
//	 */
func XbboxIsTransmeridian(tls *libc.TLS, bbox uintptr) (r uint8) {
	return libc.BoolUint8((*TBBox)(unsafe.Pointer(bbox)).Feast < (*TBBox)(unsafe.Pointer(bbox)).Fwest)
}

// C documentation
//
//	/**
//	 * Get the center of a bounding box
//	 * @param bbox   Input bounding box
//	 * @param center Output center coordinate
//	 */
func XbboxCenter(tls *libc.TLS, bbox uintptr, center uintptr) {
	var east, v1 float64
	_, _ = east, v1
	(*TLatLng)(unsafe.Pointer(center)).Flat = ((*TBBox)(unsafe.Pointer(bbox)).Fnorth + (*TBBox)(unsafe.Pointer(bbox)).Fsouth) * float64(0.5)
	if XbboxIsTransmeridian(tls, bbox) != 0 {
		v1 = (*TBBox)(unsafe.Pointer(bbox)).Feast + float64(6.283185307179586)
	} else {
		v1 = (*TBBox)(unsafe.Pointer(bbox)).Feast
	}
	// If the bbox crosses the antimeridian, shift east 360 degrees
	east = v1
	(*TLatLng)(unsafe.Pointer(center)).Flng = XconstrainLng(tls, (east+(*TBBox)(unsafe.Pointer(bbox)).Fwest)*float64(0.5))
}

// C documentation
//
//	/**
//	 * Whether the bounding box contains a given point
//	 * @param  bbox  Bounding box
//	 * @param  point Point to test
//	 * @return       Whether the point is contained
//	 */
func XbboxContains(tls *libc.TLS, bbox uintptr, point uintptr) (r uint8) {
	var v1 int32
	var v2 bool
	_, _ = v1, v2
	if v2 = (*TLatLng)(unsafe.Pointer(point)).Flat >= (*TBBox)(unsafe.Pointer(bbox)).Fsouth && (*TLatLng)(unsafe.Pointer(point)).Flat <= (*TBBox)(unsafe.Pointer(bbox)).Fnorth; v2 {
		if XbboxIsTransmeridian(tls, bbox) != 0 {
			v1 = libc.BoolInt32((*TLatLng)(unsafe.Pointer(point)).Flng >= (*TBBox)(unsafe.Pointer(bbox)).Fwest || (*TLatLng)(unsafe.Pointer(point)).Flng <= (*TBBox)(unsafe.Pointer(bbox)).Feast)
		} else {
			v1 = libc.BoolInt32((*TLatLng)(unsafe.Pointer(point)).Flng >= (*TBBox)(unsafe.Pointer(bbox)).Fwest && (*TLatLng)(unsafe.Pointer(point)).Flng <= (*TBBox)(unsafe.Pointer(bbox)).Feast)
		}
	}
	return libc.BoolUint8(v2 && v1 != 0)
}

// C documentation
//
//	/**
//	 * Whether two bounding boxes overlap
//	 * @param  a First bounding box
//	 * @param  b Second bounding box
//	 * @return   Whether the bounding boxes overlap
//	 */
func XbboxOverlapsBBox(tls *libc.TLS, a uintptr, b uintptr) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* aNormalization at bp+0 */ TLongitudeNormalization
	var _ /* bNormalization at bp+4 */ TLongitudeNormalization
	// Check whether latitude coords overlap
	if (*TBBox)(unsafe.Pointer(a)).Fnorth < (*TBBox)(unsafe.Pointer(b)).Fsouth || (*TBBox)(unsafe.Pointer(a)).Fsouth > (*TBBox)(unsafe.Pointer(b)).Fnorth {
		return uint8(Dfalse)
	}
	XbboxNormalization(tls, a, b, bp, bp+4)
	if XnormalizeLng(tls, (*TBBox)(unsafe.Pointer(a)).Feast, *(*TLongitudeNormalization)(unsafe.Pointer(bp))) < XnormalizeLng(tls, (*TBBox)(unsafe.Pointer(b)).Fwest, *(*TLongitudeNormalization)(unsafe.Pointer(bp + 4))) || XnormalizeLng(tls, (*TBBox)(unsafe.Pointer(a)).Fwest, *(*TLongitudeNormalization)(unsafe.Pointer(bp))) > XnormalizeLng(tls, (*TBBox)(unsafe.Pointer(b)).Feast, *(*TLongitudeNormalization)(unsafe.Pointer(bp + 4))) {
		return uint8(Dfalse)
	}
	return uint8(Dtrue)
}

// C documentation
//
//	/**
//	 * Whether one bounding box contains another
//	 * @param  a First bounding box
//	 * @param  b Second bounding box
//	 * @return   Whether a contains b
//	 */
func XbboxContainsBBox(tls *libc.TLS, a uintptr, b uintptr) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* aNormalization at bp+0 */ TLongitudeNormalization
	var _ /* bNormalization at bp+4 */ TLongitudeNormalization
	// Check whether latitude coords are contained
	if (*TBBox)(unsafe.Pointer(a)).Fnorth < (*TBBox)(unsafe.Pointer(b)).Fnorth || (*TBBox)(unsafe.Pointer(a)).Fsouth > (*TBBox)(unsafe.Pointer(b)).Fsouth {
		return uint8(Dfalse)
	}
	XbboxNormalization(tls, a, b, bp, bp+4)
	return libc.BoolUint8(XnormalizeLng(tls, (*TBBox)(unsafe.Pointer(a)).Fwest, *(*TLongitudeNormalization)(unsafe.Pointer(bp))) <= XnormalizeLng(tls, (*TBBox)(unsafe.Pointer(b)).Fwest, *(*TLongitudeNormalization)(unsafe.Pointer(bp + 4))) && XnormalizeLng(tls, (*TBBox)(unsafe.Pointer(a)).Feast, *(*TLongitudeNormalization)(unsafe.Pointer(bp))) >= XnormalizeLng(tls, (*TBBox)(unsafe.Pointer(b)).Feast, *(*TLongitudeNormalization)(unsafe.Pointer(bp + 4))))
}

// C documentation
//
//	/**
//	 * Whether two bounding boxes are strictly equal
//	 * @param  b1 Bounding box 1
//	 * @param  b2 Bounding box 2
//	 * @return    Whether the boxes are equal
//	 */
func XbboxEquals(tls *libc.TLS, b1 uintptr, b2 uintptr) (r uint8) {
	return libc.BoolUint8((*TBBox)(unsafe.Pointer(b1)).Fnorth == (*TBBox)(unsafe.Pointer(b2)).Fnorth && (*TBBox)(unsafe.Pointer(b1)).Fsouth == (*TBBox)(unsafe.Pointer(b2)).Fsouth && (*TBBox)(unsafe.Pointer(b1)).Feast == (*TBBox)(unsafe.Pointer(b2)).Feast && (*TBBox)(unsafe.Pointer(b1)).Fwest == (*TBBox)(unsafe.Pointer(b2)).Fwest)
}

func XbboxToCellBoundary(tls *libc.TLS, bbox uintptr) (r TCellBoundary) {
	var bboxBoundary TCellBoundary
	_ = bboxBoundary
	// Convert bbox to cell boundary, CCW vertex order
	bboxBoundary = TCellBoundary{
		FnumVerts: int32(4),
		Fverts: [10]TLatLng{
			0: {
				Flat: (*TBBox)(unsafe.Pointer(bbox)).Fnorth,
				Flng: (*TBBox)(unsafe.Pointer(bbox)).Feast,
			},
			1: {
				Flat: (*TBBox)(unsafe.Pointer(bbox)).Fnorth,
				Flng: (*TBBox)(unsafe.Pointer(bbox)).Fwest,
			},
			2: {
				Flat: (*TBBox)(unsafe.Pointer(bbox)).Fsouth,
				Flng: (*TBBox)(unsafe.Pointer(bbox)).Fwest,
			},
			3: {
				Flat: (*TBBox)(unsafe.Pointer(bbox)).Fsouth,
				Flng: (*TBBox)(unsafe.Pointer(bbox)).Feast,
			},
		},
	}
	return bboxBoundary
}

// C documentation
//
//	/**
//	 * _hexRadiusKm returns the radius of a given hexagon in Km
//	 *
//	 * @param h3Index the index of the hexagon
//	 * @return the radius of the hexagon in Km
//	 */
func X_hexRadiusKm(tls *libc.TLS, h3Index TH3Index) (r float64) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var _ /* h3Boundary at bp+16 */ TCellBoundary
	var _ /* h3Center at bp+0 */ TLatLng
	XcellToLatLng(tls, h3Index, bp)
	XcellToBoundary(tls, h3Index, bp+16)
	return XgreatCircleDistanceKm(tls, bp, bp+16+8)
}

// C documentation
//
//	/**
//	 * bboxHexEstimate returns an estimated number of hexagons that fit
//	 *                 within the cartesian-projected bounding box
//	 *
//	 * @param bbox the bounding box to estimate the hexagon fill level
//	 * @param res the resolution of the H3 hexagons to fill the bounding box
//	 * @param out the estimated number of hexagons to fill the bounding box
//	 * @return E_SUCCESS (0) on success, or another value otherwise.
//	 */
func XbboxHexEstimate(tls *libc.TLS, bbox uintptr, res int32, out uintptr) (r TH3Error) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var a, d, estimateDouble, latDiff, length, lngDiff, pentagonAreaKm2, pentagonRadiusKm, ratio, width float64
	var estimate Tint64_t
	var pentagonsErr TH3Error
	var v1 uint64
	var _ /* __u at bp+0 */ struct {
		F__i [0]uint64
		F__f float64
	}
	var _ /* p1 at bp+104 */ TLatLng
	var _ /* p2 at bp+120 */ TLatLng
	var _ /* pentagons at bp+8 */ [12]TH3Index
	_, _, _, _, _, _, _, _, _, _, _, _, _ = a, d, estimate, estimateDouble, latDiff, length, lngDiff, pentagonAreaKm2, pentagonRadiusKm, pentagonsErr, ratio, width, v1
	// Get the area of the pentagon as the maximally-distorted area possible
	*(*[12]TH3Index)(unsafe.Pointer(bp + 8)) = [12]TH3Index{}
	pentagonsErr = XgetPentagons(tls, res, bp+8)
	if pentagonsErr != 0 {
		return pentagonsErr
	}
	pentagonRadiusKm = X_hexRadiusKm(tls, (*(*[12]TH3Index)(unsafe.Pointer(bp + 8)))[0])
	// Area of a regular hexagon is 3/2*sqrt(3) * r * r
	// The pentagon has the most distortion (smallest edges) and shares its
	// edges with hexagons, so the most-distorted hexagons have this area,
	// shrunk by 20% off chance that the bounding box perfectly bounds a
	// pentagon.
	pentagonAreaKm2 = float64(0.8) * (float64(2.59807621135) * pentagonRadiusKm * pentagonRadiusKm)
	(*(*TLatLng)(unsafe.Pointer(bp + 104))).Flat = (*TBBox)(unsafe.Pointer(bbox)).Fnorth
	(*(*TLatLng)(unsafe.Pointer(bp + 104))).Flng = (*TBBox)(unsafe.Pointer(bbox)).Feast
	(*(*TLatLng)(unsafe.Pointer(bp + 120))).Flat = (*TBBox)(unsafe.Pointer(bbox)).Fsouth
	(*(*TLatLng)(unsafe.Pointer(bp + 120))).Flng = (*TBBox)(unsafe.Pointer(bbox)).Fwest
	d = XgreatCircleDistanceKm(tls, bp+104, bp+120)
	lngDiff = libc.Xfabs(tls, (*(*TLatLng)(unsafe.Pointer(bp + 104))).Flng-(*(*TLatLng)(unsafe.Pointer(bp + 120))).Flng)
	latDiff = libc.Xfabs(tls, (*(*TLatLng)(unsafe.Pointer(bp + 104))).Flat-(*(*TLatLng)(unsafe.Pointer(bp + 120))).Flat)
	if lngDiff == libc.Float64FromInt32(0) || latDiff == libc.Float64FromInt32(0) {
		return uint32(_E_FAILED)
	}
	length = libc.X__builtin_fmax(tls, lngDiff, latDiff)
	width = libc.X__builtin_fmin(tls, lngDiff, latDiff)
	ratio = length / width
	// Derived constant based on: https://math.stackexchange.com/a/1921940
	// Clamped to 3 as higher values tend to rapidly drag the estimate to zero.
	a = d * d / libc.X__builtin_fmin(tls, float64(3), ratio)
	// Divide the two to get an estimate of the number of hexagons needed
	estimateDouble = libc.Xceil(tls, a/pentagonAreaKm2)
	*(*float64)(unsafe.Pointer(bp)) = estimateDouble
	v1 = *(*uint64)(unsafe.Pointer(bp))
	goto _2
_2:
	if !(libc.BoolInt32(v1&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) < libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0) {
		return uint32(_E_FAILED)
	}
	estimate = int64(estimateDouble)
	if estimate == 0 {
		estimate = int64(1)
	}
	*(*Tint64_t)(unsafe.Pointer(out)) = estimate
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * lineHexEstimate returns an estimated number of hexagons that trace
//	 *                 the cartesian-projected line
//	 *
//	 * @param origin the origin coordinates
//	 * @param destination the destination coordinates
//	 * @param res the resolution of the H3 hexagons to trace the line
//	 * @param out Out parameter for the estimated number of hexagons required to
//	 * trace the line
//	 * @return E_SUCCESS (0) on success or another value otherwise.
//	 */
func XlineHexEstimate(tls *libc.TLS, origin uintptr, destination uintptr, res int32, out uintptr) (r TH3Error) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var dist, distCeil, pentagonRadiusKm float64
	var estimate Tint64_t
	var pentagonsErr TH3Error
	var v1 uint64
	var _ /* __u at bp+0 */ struct {
		F__i [0]uint64
		F__f float64
	}
	var _ /* pentagons at bp+8 */ [12]TH3Index
	_, _, _, _, _, _ = dist, distCeil, estimate, pentagonRadiusKm, pentagonsErr, v1
	// Get the area of the pentagon as the maximally-distorted area possible
	*(*[12]TH3Index)(unsafe.Pointer(bp + 8)) = [12]TH3Index{}
	pentagonsErr = XgetPentagons(tls, res, bp+8)
	if pentagonsErr != 0 {
		return pentagonsErr
	}
	pentagonRadiusKm = X_hexRadiusKm(tls, (*(*[12]TH3Index)(unsafe.Pointer(bp + 8)))[0])
	dist = XgreatCircleDistanceKm(tls, origin, destination)
	distCeil = libc.Xceil(tls, dist/(libc.Float64FromInt32(2)*pentagonRadiusKm))
	*(*float64)(unsafe.Pointer(bp)) = distCeil
	v1 = *(*uint64)(unsafe.Pointer(bp))
	goto _2
_2:
	if !(libc.BoolInt32(v1&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) < libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0) {
		return uint32(_E_FAILED)
	}
	estimate = int64(distCeil)
	if estimate == 0 {
		estimate = int64(1)
	}
	*(*Tint64_t)(unsafe.Pointer(out)) = estimate
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * Scale a given bounding box by some factor. Scales both width and height
//	 * by the factor, rather than scaling area, which will scale at scale^2.
//	 * Note that this function is meant to handle bounding boxes and scales,
//	 * within a reasonable domain, and does not guarantee reasonable results for
//	 * extreme values.
//	 * @param bbox Bounding box to scale, in-place
//	 * @param scale Scale factor
//	 */
func XscaleBBox(tls *libc.TLS, bbox uintptr, scale float64) {
	var height, heightBuffer, width, widthBuffer float64
	_, _, _, _ = height, heightBuffer, width, widthBuffer
	width = XbboxWidthRads(tls, bbox)
	height = XbboxHeightRads(tls, bbox)
	widthBuffer = (width*scale - width) * float64(0.5)
	heightBuffer = (height*scale - height) * float64(0.5)
	// Scale north and south, clamping to latitude domain
	*(*float64)(unsafe.Pointer(bbox)) += heightBuffer
	if (*TBBox)(unsafe.Pointer(bbox)).Fnorth > float64(1.5707963267948966) {
		(*TBBox)(unsafe.Pointer(bbox)).Fnorth = float64(1.5707963267948966)
	}
	*(*float64)(unsafe.Pointer(bbox + 8)) -= heightBuffer
	if (*TBBox)(unsafe.Pointer(bbox)).Fsouth < -libc.Float64FromFloat64(1.5707963267948966) {
		(*TBBox)(unsafe.Pointer(bbox)).Fsouth = -libc.Float64FromFloat64(1.5707963267948966)
	}
	// Scale east and west, clamping to longitude domain
	*(*float64)(unsafe.Pointer(bbox + 16)) += widthBuffer
	if (*TBBox)(unsafe.Pointer(bbox)).Feast > float64(3.141592653589793) {
		*(*float64)(unsafe.Pointer(bbox + 16)) -= float64(6.283185307179586)
	}
	if (*TBBox)(unsafe.Pointer(bbox)).Feast < -libc.Float64FromFloat64(3.141592653589793) {
		*(*float64)(unsafe.Pointer(bbox + 16)) += float64(6.283185307179586)
	}
	*(*float64)(unsafe.Pointer(bbox + 24)) -= widthBuffer
	if (*TBBox)(unsafe.Pointer(bbox)).Fwest > float64(3.141592653589793) {
		*(*float64)(unsafe.Pointer(bbox + 24)) -= float64(6.283185307179586)
	}
	if (*TBBox)(unsafe.Pointer(bbox)).Fwest < -libc.Float64FromFloat64(3.141592653589793) {
		*(*float64)(unsafe.Pointer(bbox + 24)) += float64(6.283185307179586)
	}
}

// C documentation
//
//	/**
//	 * Determine the longitude normalization scheme for two bounding boxes, either
//	 * or both of which might cross the antimeridian. The goal is to transform
//	 * latitudes in one or both boxes so that they are in the same frame of
//	 * reference and can be operated on with standard Cartesian functions.
//	 * @param a              First bounding box
//	 * @param b              Second bounding box
//	 * @param aNormalization Output: Normalization for longitudes in the first box
//	 * @param bNormalization Output: Normalization for longitudes in the second box
//	 */
func XbboxNormalization(tls *libc.TLS, a uintptr, b uintptr, aNormalization uintptr, bNormalization uintptr) {
	var aIsTransmeridian, aToBTrendsEast, bIsTransmeridian uint8
	var v1, v2, v3, v4, v5, v6 int32
	_, _, _, _, _, _, _, _, _ = aIsTransmeridian, aToBTrendsEast, bIsTransmeridian, v1, v2, v3, v4, v5, v6
	aIsTransmeridian = XbboxIsTransmeridian(tls, a)
	bIsTransmeridian = XbboxIsTransmeridian(tls, b)
	aToBTrendsEast = libc.BoolUint8((*TBBox)(unsafe.Pointer(a)).Fwest-(*TBBox)(unsafe.Pointer(b)).Feast < (*TBBox)(unsafe.Pointer(b)).Fwest-(*TBBox)(unsafe.Pointer(a)).Feast)
	// If neither is transmeridian, no normalization.
	// If both are transmeridian, normalize east by convention.
	// If one is transmeridian and one is not, normalize toward the other.
	if !(aIsTransmeridian != 0) {
		v1 = int32(_NORMALIZE_NONE)
	} else {
		if bIsTransmeridian != 0 {
			v2 = int32(_NORMALIZE_EAST)
		} else {
			if aToBTrendsEast != 0 {
				v3 = int32(_NORMALIZE_EAST)
			} else {
				v3 = int32(_NORMALIZE_WEST)
			}
			v2 = v3
		}
		v1 = v2
	}
	*(*TLongitudeNormalization)(unsafe.Pointer(aNormalization)) = v1
	if !(bIsTransmeridian != 0) {
		v4 = int32(_NORMALIZE_NONE)
	} else {
		if aIsTransmeridian != 0 {
			v5 = int32(_NORMALIZE_EAST)
		} else {
			if aToBTrendsEast != 0 {
				v6 = int32(_NORMALIZE_WEST)
			} else {
				v6 = int32(_NORMALIZE_EAST)
			}
			v5 = v6
		}
		v4 = v5
	}
	*(*TLongitudeNormalization)(unsafe.Pointer(bNormalization)) = v4
}

const DINT32_MAX1 = 2147483647
const DM_2PI2 = 6.28318530717958647692528676655900576839433
const DM_ONESEVENTH1 = 0.14285714285714285
const DM_PI2 = 3.14159265358979323846
const DM_RSIN601 = 1.1547005383792515
const DM_SQRT3_21 = 0.8660254037844386

var _UNIT_VECS3 = [7]TCoordIJK{
	0: {},
	1: {
		Fk: int32(1),
	},
	2: {
		Fj: int32(1),
	},
	3: {
		Fj: int32(1),
		Fk: int32(1),
	},
	4: {
		Fi: int32(1),
	},
	5: {
		Fi: int32(1),
		Fk: int32(1),
	},
	6: {
		Fi: int32(1),
		Fj: int32(1),
	},
}

// C documentation
//
//	/**
//	 * Sets an IJK coordinate to the specified component values.
//	 *
//	 * @param ijk The IJK coordinate to set.
//	 * @param i The desired i component value.
//	 * @param j The desired j component value.
//	 * @param k The desired k component value.
//	 */
func X_setIJK(tls *libc.TLS, ijk uintptr, i int32, j int32, k int32) {
	(*TCoordIJK)(unsafe.Pointer(ijk)).Fi = i
	(*TCoordIJK)(unsafe.Pointer(ijk)).Fj = j
	(*TCoordIJK)(unsafe.Pointer(ijk)).Fk = k
}

// C documentation
//
//	/**
//	 * Determine the containing hex in ijk+ coordinates for a 2D cartesian
//	 * coordinate vector (from DGGRID).
//	 *
//	 * @param v The 2D cartesian coordinate vector.
//	 * @param h The ijk+ coordinates of the containing hex.
//	 */
func X_hex2dToCoordIJK(tls *libc.TLS, v uintptr, h uintptr) {
	var a1, a2, r1, r2, x1, x2 float64
	var axisi, axisi1, diff, diff1 int64
	var m1, m2 int32
	_, _, _, _, _, _, _, _, _, _, _, _ = a1, a2, axisi, axisi1, diff, diff1, m1, m2, r1, r2, x1, x2
	// quantize into the ij system and then normalize
	(*TCoordIJK)(unsafe.Pointer(h)).Fk = 0
	a1 = float64(libc.Xfabsl(tls, (*TVec2d)(unsafe.Pointer(v)).Fx))
	a2 = float64(libc.Xfabsl(tls, (*TVec2d)(unsafe.Pointer(v)).Fy))
	// first do a reverse conversion
	x2 = a2 * float64(1.1547005383792515)
	x1 = a1 + x2/float64(2)
	// check if we have the center of a hex
	m1 = int32(x1)
	m2 = int32(x2)
	// otherwise round correctly
	r1 = x1 - float64(m1)
	r2 = x2 - float64(m2)
	if r1 < float64(0.5) {
		if r1 < libc.Float64FromFloat64(1)/libc.Float64FromFloat64(3) {
			if r2 < (float64(1)+r1)/float64(2) {
				(*TCoordIJK)(unsafe.Pointer(h)).Fi = m1
				(*TCoordIJK)(unsafe.Pointer(h)).Fj = m2
			} else {
				(*TCoordIJK)(unsafe.Pointer(h)).Fi = m1
				(*TCoordIJK)(unsafe.Pointer(h)).Fj = m2 + int32(1)
			}
		} else {
			if r2 < float64(1)-r1 {
				(*TCoordIJK)(unsafe.Pointer(h)).Fj = m2
			} else {
				(*TCoordIJK)(unsafe.Pointer(h)).Fj = m2 + int32(1)
			}
			if float64(1)-r1 <= r2 && r2 < float64(2)*r1 {
				(*TCoordIJK)(unsafe.Pointer(h)).Fi = m1 + int32(1)
			} else {
				(*TCoordIJK)(unsafe.Pointer(h)).Fi = m1
			}
		}
	} else {
		if r1 < libc.Float64FromFloat64(2)/libc.Float64FromFloat64(3) {
			if r2 < float64(1)-r1 {
				(*TCoordIJK)(unsafe.Pointer(h)).Fj = m2
			} else {
				(*TCoordIJK)(unsafe.Pointer(h)).Fj = m2 + int32(1)
			}
			if float64(2)*r1-float64(1) < r2 && r2 < float64(1)-r1 {
				(*TCoordIJK)(unsafe.Pointer(h)).Fi = m1
			} else {
				(*TCoordIJK)(unsafe.Pointer(h)).Fi = m1 + int32(1)
			}
		} else {
			if r2 < r1/float64(2) {
				(*TCoordIJK)(unsafe.Pointer(h)).Fi = m1 + int32(1)
				(*TCoordIJK)(unsafe.Pointer(h)).Fj = m2
			} else {
				(*TCoordIJK)(unsafe.Pointer(h)).Fi = m1 + int32(1)
				(*TCoordIJK)(unsafe.Pointer(h)).Fj = m2 + int32(1)
			}
		}
	}
	// now fold across the axes if necessary
	if (*TVec2d)(unsafe.Pointer(v)).Fx < float64(0) {
		if (*TCoordIJK)(unsafe.Pointer(h)).Fj%int32(2) == 0 { // even
			axisi = int64((*TCoordIJK)(unsafe.Pointer(h)).Fj / int32(2))
			diff = int64((*TCoordIJK)(unsafe.Pointer(h)).Fi) - axisi
			(*TCoordIJK)(unsafe.Pointer(h)).Fi = int32(float64((*TCoordIJK)(unsafe.Pointer(h)).Fi) - libc.Float64FromFloat64(2)*float64(diff))
		} else {
			axisi1 = int64(((*TCoordIJK)(unsafe.Pointer(h)).Fj + int32(1)) / int32(2))
			diff1 = int64((*TCoordIJK)(unsafe.Pointer(h)).Fi) - axisi1
			(*TCoordIJK)(unsafe.Pointer(h)).Fi = int32(float64((*TCoordIJK)(unsafe.Pointer(h)).Fi) - (libc.Float64FromFloat64(2)*float64(diff1) + libc.Float64FromInt32(1)))
		}
	}
	if (*TVec2d)(unsafe.Pointer(v)).Fy < float64(0) {
		(*TCoordIJK)(unsafe.Pointer(h)).Fi = (*TCoordIJK)(unsafe.Pointer(h)).Fi - (int32(2)*(*TCoordIJK)(unsafe.Pointer(h)).Fj+int32(1))/int32(2)
		(*TCoordIJK)(unsafe.Pointer(h)).Fj = -int32(1) * (*TCoordIJK)(unsafe.Pointer(h)).Fj
	}
	X_ijkNormalize(tls, h)
}

// C documentation
//
//	/**
//	 * Find the center point in 2D cartesian coordinates of a hex.
//	 *
//	 * @param h The ijk coordinates of the hex.
//	 * @param v The 2D cartesian coordinates of the hex center point.
//	 */
func X_ijkToHex2d(tls *libc.TLS, h uintptr, v uintptr) {
	var i, j int32
	_, _ = i, j
	i = (*TCoordIJK)(unsafe.Pointer(h)).Fi - (*TCoordIJK)(unsafe.Pointer(h)).Fk
	j = (*TCoordIJK)(unsafe.Pointer(h)).Fj - (*TCoordIJK)(unsafe.Pointer(h)).Fk
	(*TVec2d)(unsafe.Pointer(v)).Fx = float64(i) - float64(0.5)*float64(j)
	(*TVec2d)(unsafe.Pointer(v)).Fy = float64(j) * float64(0.8660254037844386)
}

// C documentation
//
//	/**
//	 * Returns whether or not two ijk coordinates contain exactly the same
//	 * component values.
//	 *
//	 * @param c1 The first set of ijk coordinates.
//	 * @param c2 The second set of ijk coordinates.
//	 * @return 1 if the two addresses match, 0 if they do not.
//	 */
func X_ijkMatches(tls *libc.TLS, c1 uintptr, c2 uintptr) (r int32) {
	return libc.BoolInt32((*TCoordIJK)(unsafe.Pointer(c1)).Fi == (*TCoordIJK)(unsafe.Pointer(c2)).Fi && (*TCoordIJK)(unsafe.Pointer(c1)).Fj == (*TCoordIJK)(unsafe.Pointer(c2)).Fj && (*TCoordIJK)(unsafe.Pointer(c1)).Fk == (*TCoordIJK)(unsafe.Pointer(c2)).Fk)
}

// C documentation
//
//	/**
//	 * Add two ijk coordinates.
//	 *
//	 * @param h1 The first set of ijk coordinates.
//	 * @param h2 The second set of ijk coordinates.
//	 * @param sum The sum of the two sets of ijk coordinates.
//	 */
func X_ijkAdd(tls *libc.TLS, h1 uintptr, h2 uintptr, sum uintptr) {
	(*TCoordIJK)(unsafe.Pointer(sum)).Fi = (*TCoordIJK)(unsafe.Pointer(h1)).Fi + (*TCoordIJK)(unsafe.Pointer(h2)).Fi
	(*TCoordIJK)(unsafe.Pointer(sum)).Fj = (*TCoordIJK)(unsafe.Pointer(h1)).Fj + (*TCoordIJK)(unsafe.Pointer(h2)).Fj
	(*TCoordIJK)(unsafe.Pointer(sum)).Fk = (*TCoordIJK)(unsafe.Pointer(h1)).Fk + (*TCoordIJK)(unsafe.Pointer(h2)).Fk
}

// C documentation
//
//	/**
//	 * Subtract two ijk coordinates.
//	 *
//	 * @param h1 The first set of ijk coordinates.
//	 * @param h2 The second set of ijk coordinates.
//	 * @param diff The difference of the two sets of ijk coordinates (h1 - h2).
//	 */
func X_ijkSub(tls *libc.TLS, h1 uintptr, h2 uintptr, diff uintptr) {
	(*TCoordIJK)(unsafe.Pointer(diff)).Fi = (*TCoordIJK)(unsafe.Pointer(h1)).Fi - (*TCoordIJK)(unsafe.Pointer(h2)).Fi
	(*TCoordIJK)(unsafe.Pointer(diff)).Fj = (*TCoordIJK)(unsafe.Pointer(h1)).Fj - (*TCoordIJK)(unsafe.Pointer(h2)).Fj
	(*TCoordIJK)(unsafe.Pointer(diff)).Fk = (*TCoordIJK)(unsafe.Pointer(h1)).Fk - (*TCoordIJK)(unsafe.Pointer(h2)).Fk
}

// C documentation
//
//	/**
//	 * Uniformly scale ijk coordinates by a scalar. Works in place.
//	 *
//	 * @param c The ijk coordinates to scale.
//	 * @param factor The scaling factor.
//	 */
func X_ijkScale(tls *libc.TLS, c uintptr, factor int32) {
	*(*int32)(unsafe.Pointer(c)) *= factor
	*(*int32)(unsafe.Pointer(c + 4)) *= factor
	*(*int32)(unsafe.Pointer(c + 8)) *= factor
}

// C documentation
//
//	/**
//	 * Returns true if _ijkNormalize with the given input could have a signed
//	 * integer overflow. Assumes k is set to 0.
//	 */
func X_ijkNormalizeCouldOverflow(tls *libc.TLS, ijk uintptr) (r uint8) {
	var max, min int32
	var v1, v10, v2, v5, v6, v9 Tint32_t
	var v11, v3, v7 uint8
	_, _, _, _, _, _, _, _, _, _, _ = max, min, v1, v10, v11, v2, v3, v5, v6, v7, v9
	if (*TCoordIJK)(unsafe.Pointer(ijk)).Fi > (*TCoordIJK)(unsafe.Pointer(ijk)).Fj {
		max = (*TCoordIJK)(unsafe.Pointer(ijk)).Fi
		min = (*TCoordIJK)(unsafe.Pointer(ijk)).Fj
	} else {
		max = (*TCoordIJK)(unsafe.Pointer(ijk)).Fj
		min = (*TCoordIJK)(unsafe.Pointer(ijk)).Fi
	}
	if min < 0 {
		// Only if the min is less than 0 will the resulting number be larger
		// than max. If min is positive, then max is also positive, and a
		// positive signed integer minus another positive signed integer will
		// not overflow.
		v1 = max
		v2 = min
		if v1 > libc.Int32FromInt32(0) {
			v3 = libc.BoolUint8(int32(DINT32_MAX1)-v1 < v2)
			goto _4
		} else {
			v3 = libc.BoolUint8(-libc.Int32FromInt32(1)-libc.Int32FromInt32(0x7fffffff)-v1 > v2)
			goto _4
		}
	_4:
		if v3 != 0 {
			// max + min would overflow
			return uint8(Dtrue)
		}
		v5 = 0
		v6 = min
		if v5 >= libc.Int32FromInt32(0) {
			v7 = libc.BoolUint8(-libc.Int32FromInt32(1)-libc.Int32FromInt32(0x7fffffff)+v5 >= v6)
			goto _8
		} else {
			v7 = libc.BoolUint8(int32(DINT32_MAX1)+v5+int32(1) < v6)
			goto _8
		}
	_8:
		if v7 != 0 {
			// 0 - INT32_MIN would overflow
			return uint8(Dtrue)
		}
		v9 = max
		v10 = min
		if v9 >= libc.Int32FromInt32(0) {
			v11 = libc.BoolUint8(-libc.Int32FromInt32(1)-libc.Int32FromInt32(0x7fffffff)+v9 >= v10)
			goto _12
		} else {
			v11 = libc.BoolUint8(int32(DINT32_MAX1)+v9+int32(1) < v10)
			goto _12
		}
	_12:
		if v11 != 0 {
			// max - min would overflow
			return uint8(Dtrue)
		}
	}
	return uint8(Dfalse)
}

// C documentation
//
//	/**
//	 * Normalizes ijk coordinates by setting the components to the smallest possible
//	 * values. Works in place.
//	 *
//	 * This function does not protect against signed integer overflow. The caller
//	 * must ensure that none of (i - j), (i - k), (j - i), (j - k), (k - i), (k - j)
//	 * will overflow. This function may be changed in the future to make that check
//	 * itself and return an error code.
//	 *
//	 * @param c The ijk coordinates to normalize.
//	 */
func X_ijkNormalize(tls *libc.TLS, c uintptr) {
	var min int32
	_ = min
	// remove any negative values
	if (*TCoordIJK)(unsafe.Pointer(c)).Fi < 0 {
		*(*int32)(unsafe.Pointer(c + 4)) -= (*TCoordIJK)(unsafe.Pointer(c)).Fi
		*(*int32)(unsafe.Pointer(c + 8)) -= (*TCoordIJK)(unsafe.Pointer(c)).Fi
		(*TCoordIJK)(unsafe.Pointer(c)).Fi = 0
	}
	if (*TCoordIJK)(unsafe.Pointer(c)).Fj < 0 {
		*(*int32)(unsafe.Pointer(c)) -= (*TCoordIJK)(unsafe.Pointer(c)).Fj
		*(*int32)(unsafe.Pointer(c + 8)) -= (*TCoordIJK)(unsafe.Pointer(c)).Fj
		(*TCoordIJK)(unsafe.Pointer(c)).Fj = 0
	}
	if (*TCoordIJK)(unsafe.Pointer(c)).Fk < 0 {
		*(*int32)(unsafe.Pointer(c)) -= (*TCoordIJK)(unsafe.Pointer(c)).Fk
		*(*int32)(unsafe.Pointer(c + 4)) -= (*TCoordIJK)(unsafe.Pointer(c)).Fk
		(*TCoordIJK)(unsafe.Pointer(c)).Fk = 0
	}
	// remove the min value if needed
	min = (*TCoordIJK)(unsafe.Pointer(c)).Fi
	if (*TCoordIJK)(unsafe.Pointer(c)).Fj < min {
		min = (*TCoordIJK)(unsafe.Pointer(c)).Fj
	}
	if (*TCoordIJK)(unsafe.Pointer(c)).Fk < min {
		min = (*TCoordIJK)(unsafe.Pointer(c)).Fk
	}
	if min > 0 {
		*(*int32)(unsafe.Pointer(c)) -= min
		*(*int32)(unsafe.Pointer(c + 4)) -= min
		*(*int32)(unsafe.Pointer(c + 8)) -= min
	}
}

// C documentation
//
//	/**
//	 * Determines the H3 digit corresponding to a unit vector or the zero vector
//	 * in ijk coordinates.
//	 *
//	 * @param ijk The ijk coordinates; must be a unit vector or zero vector.
//	 * @return The H3 digit (0-6) corresponding to the ijk unit vector, zero vector,
//	 * or INVALID_DIGIT (7) on failure.
//	 */
func X_unitIjkToDigit(tls *libc.TLS, ijk uintptr) (r TDirection) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var digit, i TDirection
	var _ /* c at bp+0 */ TCoordIJK
	_, _ = digit, i
	*(*TCoordIJK)(unsafe.Pointer(bp)) = *(*TCoordIJK)(unsafe.Pointer(ijk))
	X_ijkNormalize(tls, bp)
	digit = int32(_INVALID_DIGIT)
	i = int32(_CENTER_DIGIT)
	for {
		if !(i < int32(_NUM_DIGITS)) {
			break
		}
		if X_ijkMatches(tls, bp, uintptr(unsafe.Pointer(&_UNIT_VECS3))+uintptr(i)*12) != 0 {
			digit = i
			break
		}
		goto _1
	_1:
		;
		i++
	}
	return digit
}

// C documentation
//
//	/**
//	 * Returns non-zero if _upAp7 with the given input could have a signed integer
//	 * overflow.
//	 *
//	 * Assumes ijk is IJK+ coordinates (no negative numbers).
//	 */
func X_upAp7Checked(tls *libc.TLS, ijk uintptr) (r TH3Error) {
	var i, i2, i3, j, j2, v21 int32
	var v1, v10, v13, v14, v17, v18, v2, v5, v6, v9 Tint32_t
	var v11, v15, v19, v3, v7 uint8
	var v22 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = i, i2, i3, j, j2, v1, v10, v11, v13, v14, v15, v17, v18, v19, v2, v21, v22, v3, v5, v6, v7, v9
	// Doesn't need to be checked because i, j, and k must all be non-negative
	i = (*TCoordIJK)(unsafe.Pointer(ijk)).Fi - (*TCoordIJK)(unsafe.Pointer(ijk)).Fk
	j = (*TCoordIJK)(unsafe.Pointer(ijk)).Fj - (*TCoordIJK)(unsafe.Pointer(ijk)).Fk
	// <0 is checked because the input must all be non-negative, but some
	// negative inputs are used in unit tests to exercise the below.
	if i >= libc.Int32FromInt32(DINT32_MAX1)/libc.Int32FromInt32(3) || j >= libc.Int32FromInt32(DINT32_MAX1)/libc.Int32FromInt32(3) || i < 0 || j < 0 {
		v1 = i
		v2 = i
		if v1 > libc.Int32FromInt32(0) {
			v3 = libc.BoolUint8(int32(DINT32_MAX1)-v1 < v2)
			goto _4
		} else {
			v3 = libc.BoolUint8(-libc.Int32FromInt32(1)-libc.Int32FromInt32(0x7fffffff)-v1 > v2)
			goto _4
		}
	_4:
		if v3 != 0 {
			return uint32(_E_FAILED)
		}
		i2 = i + i
		v5 = i2
		v6 = i
		if v5 > libc.Int32FromInt32(0) {
			v7 = libc.BoolUint8(int32(DINT32_MAX1)-v5 < v6)
			goto _8
		} else {
			v7 = libc.BoolUint8(-libc.Int32FromInt32(1)-libc.Int32FromInt32(0x7fffffff)-v5 > v6)
			goto _8
		}
	_8:
		if v7 != 0 {
			return uint32(_E_FAILED)
		}
		i3 = i2 + i
		v9 = j
		v10 = j
		if v9 > libc.Int32FromInt32(0) {
			v11 = libc.BoolUint8(int32(DINT32_MAX1)-v9 < v10)
			goto _12
		} else {
			v11 = libc.BoolUint8(-libc.Int32FromInt32(1)-libc.Int32FromInt32(0x7fffffff)-v9 > v10)
			goto _12
		}
	_12:
		if v11 != 0 {
			return uint32(_E_FAILED)
		}
		j2 = j + j
		v13 = i3
		v14 = j
		if v13 >= libc.Int32FromInt32(0) {
			v15 = libc.BoolUint8(-libc.Int32FromInt32(1)-libc.Int32FromInt32(0x7fffffff)+v13 >= v14)
			goto _16
		} else {
			v15 = libc.BoolUint8(int32(DINT32_MAX1)+v13+int32(1) < v14)
			goto _16
		}
	_16:
		if v15 != 0 {
			return uint32(_E_FAILED)
		}
		v17 = i
		v18 = j2
		if v17 > libc.Int32FromInt32(0) {
			v19 = libc.BoolUint8(int32(DINT32_MAX1)-v17 < v18)
			goto _20
		} else {
			v19 = libc.BoolUint8(-libc.Int32FromInt32(1)-libc.Int32FromInt32(0x7fffffff)-v17 > v18)
			goto _20
		}
	_20:
		if v19 != 0 {
			return uint32(_E_FAILED)
		}
	}
	(*TCoordIJK)(unsafe.Pointer(ijk)).Fi = int32(libc.X__builtin_lround(tls, float64(i*libc.Int32FromInt32(3)-j)*float64(0.14285714285714285)))
	(*TCoordIJK)(unsafe.Pointer(ijk)).Fj = int32(libc.X__builtin_lround(tls, float64(i+j*libc.Int32FromInt32(2))*float64(0.14285714285714285)))
	(*TCoordIJK)(unsafe.Pointer(ijk)).Fk = 0
	// Expected not to be reachable, because max + min or max - min would need
	// to overflow.
	if X_ijkNormalizeCouldOverflow(tls, ijk) != 0 {
		if v22 = libc.Bool(0 != 0); !v22 {
			libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+27, int32(354), uintptr(unsafe.Pointer(&___func__2)))
		}
		_ = v22 || libc.Bool(libc.Int32FromInt32(0) != 0)
		v21 = libc.Int32FromInt32(1)
	} else {
		v21 = 0
	}
	if v21 != 0 {
		return uint32(_E_FAILED)
	}
	X_ijkNormalize(tls, ijk)
	return uint32(_E_SUCCESS)
}

var ___func__2 = [14]uint8{'_', 'u', 'p', 'A', 'p', '7', 'C', 'h', 'e', 'c', 'k', 'e', 'd'}

// C documentation
//
//	/**
//	 * Returns non-zero if _upAp7r with the given input could have a signed integer
//	 * overflow.
//	 *
//	 * Assumes ijk is IJK+ coordinates (no negative numbers).
//	 */
func X_upAp7rChecked(tls *libc.TLS, ijk uintptr) (r TH3Error) {
	var i, i2, j, j2, j3, v21 int32
	var v1, v10, v13, v14, v17, v18, v2, v5, v6, v9 Tint32_t
	var v11, v15, v19, v3, v7 uint8
	var v22 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = i, i2, j, j2, j3, v1, v10, v11, v13, v14, v15, v17, v18, v19, v2, v21, v22, v3, v5, v6, v7, v9
	// Doesn't need to be checked because i, j, and k must all be non-negative
	i = (*TCoordIJK)(unsafe.Pointer(ijk)).Fi - (*TCoordIJK)(unsafe.Pointer(ijk)).Fk
	j = (*TCoordIJK)(unsafe.Pointer(ijk)).Fj - (*TCoordIJK)(unsafe.Pointer(ijk)).Fk
	// <0 is checked because the input must all be non-negative, but some
	// negative inputs are used in unit tests to exercise the below.
	if i >= libc.Int32FromInt32(DINT32_MAX1)/libc.Int32FromInt32(3) || j >= libc.Int32FromInt32(DINT32_MAX1)/libc.Int32FromInt32(3) || i < 0 || j < 0 {
		v1 = i
		v2 = i
		if v1 > libc.Int32FromInt32(0) {
			v3 = libc.BoolUint8(int32(DINT32_MAX1)-v1 < v2)
			goto _4
		} else {
			v3 = libc.BoolUint8(-libc.Int32FromInt32(1)-libc.Int32FromInt32(0x7fffffff)-v1 > v2)
			goto _4
		}
	_4:
		if v3 != 0 {
			return uint32(_E_FAILED)
		}
		i2 = i + i
		v5 = j
		v6 = j
		if v5 > libc.Int32FromInt32(0) {
			v7 = libc.BoolUint8(int32(DINT32_MAX1)-v5 < v6)
			goto _8
		} else {
			v7 = libc.BoolUint8(-libc.Int32FromInt32(1)-libc.Int32FromInt32(0x7fffffff)-v5 > v6)
			goto _8
		}
	_8:
		if v7 != 0 {
			return uint32(_E_FAILED)
		}
		j2 = j + j
		v9 = j2
		v10 = j
		if v9 > libc.Int32FromInt32(0) {
			v11 = libc.BoolUint8(int32(DINT32_MAX1)-v9 < v10)
			goto _12
		} else {
			v11 = libc.BoolUint8(-libc.Int32FromInt32(1)-libc.Int32FromInt32(0x7fffffff)-v9 > v10)
			goto _12
		}
	_12:
		if v11 != 0 {
			return uint32(_E_FAILED)
		}
		j3 = j2 + j
		v13 = i2
		v14 = j
		if v13 > libc.Int32FromInt32(0) {
			v15 = libc.BoolUint8(int32(DINT32_MAX1)-v13 < v14)
			goto _16
		} else {
			v15 = libc.BoolUint8(-libc.Int32FromInt32(1)-libc.Int32FromInt32(0x7fffffff)-v13 > v14)
			goto _16
		}
	_16:
		if v15 != 0 {
			return uint32(_E_FAILED)
		}
		v17 = j3
		v18 = i
		if v17 >= libc.Int32FromInt32(0) {
			v19 = libc.BoolUint8(-libc.Int32FromInt32(1)-libc.Int32FromInt32(0x7fffffff)+v17 >= v18)
			goto _20
		} else {
			v19 = libc.BoolUint8(int32(DINT32_MAX1)+v17+int32(1) < v18)
			goto _20
		}
	_20:
		if v19 != 0 {
			return uint32(_E_FAILED)
		}
	}
	(*TCoordIJK)(unsafe.Pointer(ijk)).Fi = int32(libc.X__builtin_lround(tls, float64(i*libc.Int32FromInt32(2)+j)*float64(0.14285714285714285)))
	(*TCoordIJK)(unsafe.Pointer(ijk)).Fj = int32(libc.X__builtin_lround(tls, float64(j*libc.Int32FromInt32(3)-i)*float64(0.14285714285714285)))
	(*TCoordIJK)(unsafe.Pointer(ijk)).Fk = 0
	// Expected not to be reachable, because max + min or max - min would need
	// to overflow.
	if X_ijkNormalizeCouldOverflow(tls, ijk) != 0 {
		if v22 = libc.Bool(0 != 0); !v22 {
			libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+27, int32(402), uintptr(unsafe.Pointer(&___func__21)))
		}
		_ = v22 || libc.Bool(libc.Int32FromInt32(0) != 0)
		v21 = libc.Int32FromInt32(1)
	} else {
		v21 = 0
	}
	if v21 != 0 {
		return uint32(_E_FAILED)
	}
	X_ijkNormalize(tls, ijk)
	return uint32(_E_SUCCESS)
}

var ___func__21 = [15]uint8{'_', 'u', 'p', 'A', 'p', '7', 'r', 'C', 'h', 'e', 'c', 'k', 'e', 'd'}

// C documentation
//
//	/**
//	 * Find the normalized ijk coordinates of the indexing parent of a cell in a
//	 * counter-clockwise aperture 7 grid. Works in place.
//	 *
//	 * @param ijk The ijk coordinates.
//	 */
func X_upAp7(tls *libc.TLS, ijk uintptr) {
	var i, j int32
	_, _ = i, j
	// convert to CoordIJ
	i = (*TCoordIJK)(unsafe.Pointer(ijk)).Fi - (*TCoordIJK)(unsafe.Pointer(ijk)).Fk
	j = (*TCoordIJK)(unsafe.Pointer(ijk)).Fj - (*TCoordIJK)(unsafe.Pointer(ijk)).Fk
	(*TCoordIJK)(unsafe.Pointer(ijk)).Fi = int32(libc.X__builtin_lround(tls, float64(libc.Int32FromInt32(3)*i-j)*float64(0.14285714285714285)))
	(*TCoordIJK)(unsafe.Pointer(ijk)).Fj = int32(libc.X__builtin_lround(tls, float64(i+libc.Int32FromInt32(2)*j)*float64(0.14285714285714285)))
	(*TCoordIJK)(unsafe.Pointer(ijk)).Fk = 0
	X_ijkNormalize(tls, ijk)
}

// C documentation
//
//	/**
//	 * Find the normalized ijk coordinates of the indexing parent of a cell in a
//	 * clockwise aperture 7 grid. Works in place.
//	 *
//	 * @param ijk The ijk coordinates.
//	 */
func X_upAp7r(tls *libc.TLS, ijk uintptr) {
	var i, j int32
	_, _ = i, j
	// convert to CoordIJ
	i = (*TCoordIJK)(unsafe.Pointer(ijk)).Fi - (*TCoordIJK)(unsafe.Pointer(ijk)).Fk
	j = (*TCoordIJK)(unsafe.Pointer(ijk)).Fj - (*TCoordIJK)(unsafe.Pointer(ijk)).Fk
	(*TCoordIJK)(unsafe.Pointer(ijk)).Fi = int32(libc.X__builtin_lround(tls, float64(libc.Int32FromInt32(2)*i+j)*float64(0.14285714285714285)))
	(*TCoordIJK)(unsafe.Pointer(ijk)).Fj = int32(libc.X__builtin_lround(tls, float64(libc.Int32FromInt32(3)*j-i)*float64(0.14285714285714285)))
	(*TCoordIJK)(unsafe.Pointer(ijk)).Fk = 0
	X_ijkNormalize(tls, ijk)
}

// C documentation
//
//	/**
//	 * Find the normalized ijk coordinates of the hex centered on the indicated
//	 * hex at the next finer aperture 7 counter-clockwise resolution. Works in
//	 * place.
//	 *
//	 * @param ijk The ijk coordinates.
//	 */
func X_downAp7(tls *libc.TLS, ijk uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var _ /* iVec at bp+0 */ TCoordIJK
	var _ /* jVec at bp+12 */ TCoordIJK
	var _ /* kVec at bp+24 */ TCoordIJK
	// res r unit vectors in res r+1
	*(*TCoordIJK)(unsafe.Pointer(bp)) = TCoordIJK{
		Fi: int32(3),
		Fk: int32(1),
	}
	*(*TCoordIJK)(unsafe.Pointer(bp + 12)) = TCoordIJK{
		Fi: int32(1),
		Fj: int32(3),
	}
	*(*TCoordIJK)(unsafe.Pointer(bp + 24)) = TCoordIJK{
		Fj: int32(1),
		Fk: int32(3),
	}
	X_ijkScale(tls, bp, (*TCoordIJK)(unsafe.Pointer(ijk)).Fi)
	X_ijkScale(tls, bp+12, (*TCoordIJK)(unsafe.Pointer(ijk)).Fj)
	X_ijkScale(tls, bp+24, (*TCoordIJK)(unsafe.Pointer(ijk)).Fk)
	X_ijkAdd(tls, bp, bp+12, ijk)
	X_ijkAdd(tls, ijk, bp+24, ijk)
	X_ijkNormalize(tls, ijk)
}

// C documentation
//
//	/**
//	 * Find the normalized ijk coordinates of the hex centered on the indicated
//	 * hex at the next finer aperture 7 clockwise resolution. Works in place.
//	 *
//	 * @param ijk The ijk coordinates.
//	 */
func X_downAp7r(tls *libc.TLS, ijk uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var _ /* iVec at bp+0 */ TCoordIJK
	var _ /* jVec at bp+12 */ TCoordIJK
	var _ /* kVec at bp+24 */ TCoordIJK
	// res r unit vectors in res r+1
	*(*TCoordIJK)(unsafe.Pointer(bp)) = TCoordIJK{
		Fi: int32(3),
		Fj: int32(1),
	}
	*(*TCoordIJK)(unsafe.Pointer(bp + 12)) = TCoordIJK{
		Fj: int32(3),
		Fk: int32(1),
	}
	*(*TCoordIJK)(unsafe.Pointer(bp + 24)) = TCoordIJK{
		Fi: int32(1),
		Fk: int32(3),
	}
	X_ijkScale(tls, bp, (*TCoordIJK)(unsafe.Pointer(ijk)).Fi)
	X_ijkScale(tls, bp+12, (*TCoordIJK)(unsafe.Pointer(ijk)).Fj)
	X_ijkScale(tls, bp+24, (*TCoordIJK)(unsafe.Pointer(ijk)).Fk)
	X_ijkAdd(tls, bp, bp+12, ijk)
	X_ijkAdd(tls, ijk, bp+24, ijk)
	X_ijkNormalize(tls, ijk)
}

// C documentation
//
//	/**
//	 * Find the normalized ijk coordinates of the hex in the specified digit
//	 * direction from the specified ijk coordinates. Works in place.
//	 *
//	 * @param ijk The ijk coordinates.
//	 * @param digit The digit direction from the original ijk coordinates.
//	 */
func X_neighbor(tls *libc.TLS, ijk uintptr, digit TDirection) {
	if digit > int32(_CENTER_DIGIT) && digit < int32(_NUM_DIGITS) {
		X_ijkAdd(tls, ijk, uintptr(unsafe.Pointer(&_UNIT_VECS3))+uintptr(digit)*12, ijk)
		X_ijkNormalize(tls, ijk)
	}
}

// C documentation
//
//	/**
//	 * Rotates ijk coordinates 60 degrees counter-clockwise. Works in place.
//	 *
//	 * @param ijk The ijk coordinates.
//	 */
func X_ijkRotate60ccw(tls *libc.TLS, ijk uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var _ /* iVec at bp+0 */ TCoordIJK
	var _ /* jVec at bp+12 */ TCoordIJK
	var _ /* kVec at bp+24 */ TCoordIJK
	// unit vector rotations
	*(*TCoordIJK)(unsafe.Pointer(bp)) = TCoordIJK{
		Fi: int32(1),
		Fj: int32(1),
	}
	*(*TCoordIJK)(unsafe.Pointer(bp + 12)) = TCoordIJK{
		Fj: int32(1),
		Fk: int32(1),
	}
	*(*TCoordIJK)(unsafe.Pointer(bp + 24)) = TCoordIJK{
		Fi: int32(1),
		Fk: int32(1),
	}
	X_ijkScale(tls, bp, (*TCoordIJK)(unsafe.Pointer(ijk)).Fi)
	X_ijkScale(tls, bp+12, (*TCoordIJK)(unsafe.Pointer(ijk)).Fj)
	X_ijkScale(tls, bp+24, (*TCoordIJK)(unsafe.Pointer(ijk)).Fk)
	X_ijkAdd(tls, bp, bp+12, ijk)
	X_ijkAdd(tls, ijk, bp+24, ijk)
	X_ijkNormalize(tls, ijk)
}

// C documentation
//
//	/**
//	 * Rotates ijk coordinates 60 degrees clockwise. Works in place.
//	 *
//	 * @param ijk The ijk coordinates.
//	 */
func X_ijkRotate60cw(tls *libc.TLS, ijk uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var _ /* iVec at bp+0 */ TCoordIJK
	var _ /* jVec at bp+12 */ TCoordIJK
	var _ /* kVec at bp+24 */ TCoordIJK
	// unit vector rotations
	*(*TCoordIJK)(unsafe.Pointer(bp)) = TCoordIJK{
		Fi: int32(1),
		Fk: int32(1),
	}
	*(*TCoordIJK)(unsafe.Pointer(bp + 12)) = TCoordIJK{
		Fi: int32(1),
		Fj: int32(1),
	}
	*(*TCoordIJK)(unsafe.Pointer(bp + 24)) = TCoordIJK{
		Fj: int32(1),
		Fk: int32(1),
	}
	X_ijkScale(tls, bp, (*TCoordIJK)(unsafe.Pointer(ijk)).Fi)
	X_ijkScale(tls, bp+12, (*TCoordIJK)(unsafe.Pointer(ijk)).Fj)
	X_ijkScale(tls, bp+24, (*TCoordIJK)(unsafe.Pointer(ijk)).Fk)
	X_ijkAdd(tls, bp, bp+12, ijk)
	X_ijkAdd(tls, ijk, bp+24, ijk)
	X_ijkNormalize(tls, ijk)
}

// C documentation
//
//	/**
//	 * Rotates indexing digit 60 degrees counter-clockwise. Returns result.
//	 *
//	 * @param digit Indexing digit (between 1 and 6 inclusive)
//	 */
func X_rotate60ccw(tls *libc.TLS, digit TDirection) (r TDirection) {
	switch digit {
	case int32(_K_AXES_DIGIT):
		return int32(_IK_AXES_DIGIT)
	case int32(_IK_AXES_DIGIT):
		return int32(_I_AXES_DIGIT)
	case int32(_I_AXES_DIGIT):
		return int32(_IJ_AXES_DIGIT)
	case int32(_IJ_AXES_DIGIT):
		return int32(_J_AXES_DIGIT)
	case int32(_J_AXES_DIGIT):
		return int32(_JK_AXES_DIGIT)
	case int32(_JK_AXES_DIGIT):
		return int32(_K_AXES_DIGIT)
	default:
		return digit
	}
	return r
}

// C documentation
//
//	/**
//	 * Rotates indexing digit 60 degrees clockwise. Returns result.
//	 *
//	 * @param digit Indexing digit (between 1 and 6 inclusive)
//	 */
func X_rotate60cw(tls *libc.TLS, digit TDirection) (r TDirection) {
	switch digit {
	case int32(_K_AXES_DIGIT):
		return int32(_JK_AXES_DIGIT)
	case int32(_JK_AXES_DIGIT):
		return int32(_J_AXES_DIGIT)
	case int32(_J_AXES_DIGIT):
		return int32(_IJ_AXES_DIGIT)
	case int32(_IJ_AXES_DIGIT):
		return int32(_I_AXES_DIGIT)
	case int32(_I_AXES_DIGIT):
		return int32(_IK_AXES_DIGIT)
	case int32(_IK_AXES_DIGIT):
		return int32(_K_AXES_DIGIT)
	default:
		return digit
	}
	return r
}

// C documentation
//
//	/**
//	 * Find the normalized ijk coordinates of the hex centered on the indicated
//	 * hex at the next finer aperture 3 counter-clockwise resolution. Works in
//	 * place.
//	 *
//	 * @param ijk The ijk coordinates.
//	 */
func X_downAp3(tls *libc.TLS, ijk uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var _ /* iVec at bp+0 */ TCoordIJK
	var _ /* jVec at bp+12 */ TCoordIJK
	var _ /* kVec at bp+24 */ TCoordIJK
	// res r unit vectors in res r+1
	*(*TCoordIJK)(unsafe.Pointer(bp)) = TCoordIJK{
		Fi: int32(2),
		Fk: int32(1),
	}
	*(*TCoordIJK)(unsafe.Pointer(bp + 12)) = TCoordIJK{
		Fi: int32(1),
		Fj: int32(2),
	}
	*(*TCoordIJK)(unsafe.Pointer(bp + 24)) = TCoordIJK{
		Fj: int32(1),
		Fk: int32(2),
	}
	X_ijkScale(tls, bp, (*TCoordIJK)(unsafe.Pointer(ijk)).Fi)
	X_ijkScale(tls, bp+12, (*TCoordIJK)(unsafe.Pointer(ijk)).Fj)
	X_ijkScale(tls, bp+24, (*TCoordIJK)(unsafe.Pointer(ijk)).Fk)
	X_ijkAdd(tls, bp, bp+12, ijk)
	X_ijkAdd(tls, ijk, bp+24, ijk)
	X_ijkNormalize(tls, ijk)
}

// C documentation
//
//	/**
//	 * Find the normalized ijk coordinates of the hex centered on the indicated
//	 * hex at the next finer aperture 3 clockwise resolution. Works in place.
//	 *
//	 * @param ijk The ijk coordinates.
//	 */
func X_downAp3r(tls *libc.TLS, ijk uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var _ /* iVec at bp+0 */ TCoordIJK
	var _ /* jVec at bp+12 */ TCoordIJK
	var _ /* kVec at bp+24 */ TCoordIJK
	// res r unit vectors in res r+1
	*(*TCoordIJK)(unsafe.Pointer(bp)) = TCoordIJK{
		Fi: int32(2),
		Fj: int32(1),
	}
	*(*TCoordIJK)(unsafe.Pointer(bp + 12)) = TCoordIJK{
		Fj: int32(2),
		Fk: int32(1),
	}
	*(*TCoordIJK)(unsafe.Pointer(bp + 24)) = TCoordIJK{
		Fi: int32(1),
		Fk: int32(2),
	}
	X_ijkScale(tls, bp, (*TCoordIJK)(unsafe.Pointer(ijk)).Fi)
	X_ijkScale(tls, bp+12, (*TCoordIJK)(unsafe.Pointer(ijk)).Fj)
	X_ijkScale(tls, bp+24, (*TCoordIJK)(unsafe.Pointer(ijk)).Fk)
	X_ijkAdd(tls, bp, bp+12, ijk)
	X_ijkAdd(tls, ijk, bp+24, ijk)
	X_ijkNormalize(tls, ijk)
}

// C documentation
//
//	/**
//	 * Finds the distance between the two coordinates. Returns result.
//	 *
//	 * @param c1 The first set of ijk coordinates.
//	 * @param c2 The second set of ijk coordinates.
//	 */
func XijkDistance(tls *libc.TLS, c1 uintptr, c2 uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var absDiff TCoordIJK
	var v1, v2, v3 int32
	var _ /* diff at bp+0 */ TCoordIJK
	_, _, _, _ = absDiff, v1, v2, v3
	X_ijkSub(tls, c1, c2, bp)
	X_ijkNormalize(tls, bp)
	absDiff = TCoordIJK{
		Fi: libc.Xabs(tls, (*(*TCoordIJK)(unsafe.Pointer(bp))).Fi),
		Fj: libc.Xabs(tls, (*(*TCoordIJK)(unsafe.Pointer(bp))).Fj),
		Fk: libc.Xabs(tls, (*(*TCoordIJK)(unsafe.Pointer(bp))).Fk),
	}
	if absDiff.Fj > absDiff.Fk {
		v2 = absDiff.Fj
	} else {
		v2 = absDiff.Fk
	}
	if absDiff.Fi > v2 {
		v1 = absDiff.Fi
	} else {
		if absDiff.Fj > absDiff.Fk {
			v3 = absDiff.Fj
		} else {
			v3 = absDiff.Fk
		}
		v1 = v3
	}
	return v1
}

// C documentation
//
//	/**
//	 * Transforms coordinates from the IJK+ coordinate system to the IJ coordinate
//	 * system.
//	 *
//	 * @param ijk The input IJK+ coordinates
//	 * @param ij The output IJ coordinates
//	 */
func XijkToIj(tls *libc.TLS, ijk uintptr, ij uintptr) {
	(*TCoordIJ)(unsafe.Pointer(ij)).Fi = (*TCoordIJK)(unsafe.Pointer(ijk)).Fi - (*TCoordIJK)(unsafe.Pointer(ijk)).Fk
	(*TCoordIJ)(unsafe.Pointer(ij)).Fj = (*TCoordIJK)(unsafe.Pointer(ijk)).Fj - (*TCoordIJK)(unsafe.Pointer(ijk)).Fk
}

// C documentation
//
//	/**
//	 * Transforms coordinates from the IJ coordinate system to the IJK+ coordinate
//	 * system.
//	 *
//	 * @param ij The input IJ coordinates
//	 * @param ijk The output IJK+ coordinates
//	 * @returns E_SUCCESS on success, E_FAILED if signed integer overflow would have
//	 * occurred.
//	 */
func XijToIjk(tls *libc.TLS, ij uintptr, ijk uintptr) (r TH3Error) {
	(*TCoordIJK)(unsafe.Pointer(ijk)).Fi = (*TCoordIJ)(unsafe.Pointer(ij)).Fi
	(*TCoordIJK)(unsafe.Pointer(ijk)).Fj = (*TCoordIJ)(unsafe.Pointer(ij)).Fj
	(*TCoordIJK)(unsafe.Pointer(ijk)).Fk = 0
	if X_ijkNormalizeCouldOverflow(tls, ijk) != 0 {
		return uint32(_E_FAILED)
	}
	X_ijkNormalize(tls, ijk)
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * Convert IJK coordinates to cube coordinates, in place
//	 * @param ijk Coordinate to convert
//	 */
func XijkToCube(tls *libc.TLS, ijk uintptr) {
	(*TCoordIJK)(unsafe.Pointer(ijk)).Fi = -(*TCoordIJK)(unsafe.Pointer(ijk)).Fi + (*TCoordIJK)(unsafe.Pointer(ijk)).Fk
	(*TCoordIJK)(unsafe.Pointer(ijk)).Fj = (*TCoordIJK)(unsafe.Pointer(ijk)).Fj - (*TCoordIJK)(unsafe.Pointer(ijk)).Fk
	(*TCoordIJK)(unsafe.Pointer(ijk)).Fk = -(*TCoordIJK)(unsafe.Pointer(ijk)).Fi - (*TCoordIJK)(unsafe.Pointer(ijk)).Fj
}

// C documentation
//
//	/**
//	 * Convert cube coordinates to IJK coordinates, in place
//	 * @param ijk Coordinate to convert
//	 */
func XcubeToIjk(tls *libc.TLS, ijk uintptr) {
	(*TCoordIJK)(unsafe.Pointer(ijk)).Fi = -(*TCoordIJK)(unsafe.Pointer(ijk)).Fi
	(*TCoordIJK)(unsafe.Pointer(ijk)).Fk = 0
	X_ijkNormalize(tls, ijk)
}

const DINT32_MAX2 = 0x7fffffff
const DMAX_BASE_CELL_FACES = 5
const DM_ONESEVENTH2 = 0.14285714285714285714285714285714285
const DM_RSIN602 = 1.1547005383792515290182975610039149112953
const DM_SQRT3_22 = 0.8660254037844386467637231707529361834714
const DPRIX16 = "X"
const DPRIX32 = "X"
const DPRIX8 = "X"
const DPRIXFAST16 = "X"
const DPRIXFAST32 = "X"
const DPRIXFAST8 = "X"
const DPRIXLEAST16 = "X"
const DPRIXLEAST32 = "X"
const DPRIXLEAST8 = "X"
const DPRId16 = "d"
const DPRId32 = "d"
const DPRId8 = "d"
const DPRIdFAST16 = "d"
const DPRIdFAST32 = "d"
const DPRIdFAST8 = "d"
const DPRIdLEAST16 = "d"
const DPRIdLEAST32 = "d"
const DPRIdLEAST8 = "d"
const DPRIi16 = "i"
const DPRIi32 = "i"
const DPRIi8 = "i"
const DPRIiFAST16 = "i"
const DPRIiFAST32 = "i"
const DPRIiFAST8 = "i"
const DPRIiLEAST16 = "i"
const DPRIiLEAST32 = "i"
const DPRIiLEAST8 = "i"
const DPRIo16 = "o"
const DPRIo32 = "o"
const DPRIo8 = "o"
const DPRIoFAST16 = "o"
const DPRIoFAST32 = "o"
const DPRIoFAST8 = "o"
const DPRIoLEAST16 = "o"
const DPRIoLEAST32 = "o"
const DPRIoLEAST8 = "o"
const DPRIu16 = "u"
const DPRIu32 = "u"
const DPRIu8 = "u"
const DPRIuFAST16 = "u"
const DPRIuFAST32 = "u"
const DPRIuFAST8 = "u"
const DPRIuLEAST16 = "u"
const DPRIuLEAST32 = "u"
const DPRIuLEAST8 = "u"
const DPRIx16 = "x"
const DPRIx32 = "x"
const DPRIx8 = "x"
const DPRIxFAST16 = "x"
const DPRIxFAST32 = "x"
const DPRIxFAST8 = "x"
const DPRIxLEAST16 = "x"
const DPRIxLEAST32 = "x"
const DPRIxLEAST8 = "x"
const DSCNd16 = "hd"
const DSCNd32 = "d"
const DSCNd8 = "hhd"
const DSCNdFAST16 = "d"
const DSCNdFAST32 = "d"
const DSCNdFAST8 = "hhd"
const DSCNdLEAST16 = "hd"
const DSCNdLEAST32 = "d"
const DSCNdLEAST8 = "hhd"
const DSCNi16 = "hi"
const DSCNi32 = "i"
const DSCNi8 = "hhi"
const DSCNiFAST16 = "i"
const DSCNiFAST32 = "i"
const DSCNiFAST8 = "hhi"
const DSCNiLEAST16 = "hi"
const DSCNiLEAST32 = "i"
const DSCNiLEAST8 = "hhi"
const DSCNo16 = "ho"
const DSCNo32 = "o"
const DSCNo8 = "hho"
const DSCNoFAST16 = "o"
const DSCNoFAST32 = "o"
const DSCNoFAST8 = "hho"
const DSCNoLEAST16 = "ho"
const DSCNoLEAST32 = "o"
const DSCNoLEAST8 = "hho"
const DSCNu16 = "hu"
const DSCNu32 = "u"
const DSCNu8 = "hhu"
const DSCNuFAST16 = "u"
const DSCNuFAST32 = "u"
const DSCNuFAST8 = "hhu"
const DSCNuLEAST16 = "hu"
const DSCNuLEAST32 = "u"
const DSCNuLEAST8 = "hhu"
const DSCNx16 = "hx"
const DSCNx32 = "x"
const DSCNx8 = "hhx"
const DSCNxFAST16 = "x"
const DSCNxFAST32 = "x"
const DSCNxFAST8 = "hhx"
const DSCNxLEAST16 = "hx"
const DSCNxLEAST32 = "x"
const DSCNxLEAST8 = "hhx"
const D__PRI64 = "l"
const D__PRIPTR = "l"

type Timaxdiv_t = struct {
	Fquot Tintmax_t
	Frem  Tintmax_t
}

var _UNIT_VECS4 = [7]TCoordIJK{
	0: {},
	1: {
		Fk: int32(1),
	},
	2: {
		Fj: int32(1),
	},
	3: {
		Fj: int32(1),
		Fk: int32(1),
	},
	4: {
		Fi: int32(1),
	},
	5: {
		Fi: int32(1),
		Fk: int32(1),
	},
	6: {
		Fi: int32(1),
		Fj: int32(1),
	},
}

type TPentagonDirectionFaces = struct {
	FbaseCell int32
	Ffaces    [5]int32
}

// C documentation
//
//	/**
//	 * Returns whether or not the provided H3Indexes are neighbors.
//	 * @param origin The origin H3 index.
//	 * @param destination The destination H3 index.
//	 * @param out Set to 1 if the indexes are neighbors, 0 otherwise
//	 * @return Error code if the origin or destination are invalid or incomparable.
//	 */
func XareNeighborCells(tls *libc.TLS, origin TH3Index, destination TH3Index, out uintptr) (r TH3Error) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var destinationResDigit, originResDigit TDirection
	var i, parentRes int32
	var neighborSetClockwise, neighborSetCounterclockwise [7]TDirection
	var _ /* destinationParent at bp+8 */ TH3Index
	var _ /* neighborRing at bp+16 */ [7]TH3Index
	var _ /* originParent at bp+0 */ TH3Index
	_, _, _, _, _, _ = destinationResDigit, i, neighborSetClockwise, neighborSetCounterclockwise, originResDigit, parentRes
	// Make sure they're hexagon indexes
	if libc.Int32FromUint64(origin&(libc.Uint64FromInt32(libc.Int32FromInt32(15))<<libc.Int32FromInt32(DH3_MODE_OFFSET))>>libc.Int32FromInt32(DH3_MODE_OFFSET)) != int32(DH3_CELL_MODE) || libc.Int32FromUint64(destination&(libc.Uint64FromInt32(libc.Int32FromInt32(15))<<libc.Int32FromInt32(DH3_MODE_OFFSET))>>libc.Int32FromInt32(DH3_MODE_OFFSET)) != int32(DH3_CELL_MODE) {
		return uint32(_E_CELL_INVALID)
	}
	// Hexagons cannot be neighbors with themselves
	if origin == destination {
		*(*int32)(unsafe.Pointer(out)) = 0
		return uint32(_E_SUCCESS)
	}
	// Only hexagons in the same resolution can be neighbors
	if libc.Int32FromUint64(origin&(libc.Uint64FromUint64(15)<<libc.Int32FromInt32(DH3_RES_OFFSET))>>libc.Int32FromInt32(DH3_RES_OFFSET)) != libc.Int32FromUint64(destination&(libc.Uint64FromUint64(15)<<libc.Int32FromInt32(DH3_RES_OFFSET))>>libc.Int32FromInt32(DH3_RES_OFFSET)) {
		return uint32(_E_RES_MISMATCH)
	}
	// H3 Indexes that share the same parent are very likely to be neighbors
	// Child 0 is neighbor with all of its parent's 'offspring', the other
	// children are neighbors with 3 of the 7 children. So a simple comparison
	// of origin and destination parents and then a lookup table of the children
	// is a super-cheap way to possibly determine they are neighbors.
	parentRes = libc.Int32FromUint64(origin&(libc.Uint64FromUint64(15)<<libc.Int32FromInt32(DH3_RES_OFFSET))>>libc.Int32FromInt32(DH3_RES_OFFSET)) - int32(1)
	if parentRes > 0 {
		XcellToParent(tls, origin, parentRes, bp)
		XcellToParent(tls, destination, parentRes, bp+8)
		if *(*TH3Index)(unsafe.Pointer(bp)) == *(*TH3Index)(unsafe.Pointer(bp + 8)) {
			originResDigit = libc.Int32FromUint64(origin >> ((libc.Int32FromInt32(DMAX_H3_RES) - (parentRes + libc.Int32FromInt32(1))) * libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET)) & libc.Uint64FromInt32(libc.Int32FromInt32(7)))
			destinationResDigit = libc.Int32FromUint64(destination >> ((libc.Int32FromInt32(DMAX_H3_RES) - (parentRes + libc.Int32FromInt32(1))) * libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET)) & libc.Uint64FromInt32(libc.Int32FromInt32(7)))
			if originResDigit == int32(_CENTER_DIGIT) || destinationResDigit == int32(_CENTER_DIGIT) {
				*(*int32)(unsafe.Pointer(out)) = int32(1)
				return uint32(_E_SUCCESS)
			}
			if originResDigit >= int32(_INVALID_DIGIT) {
				// Prevent indexing off the end of the array below
				return uint32(_E_CELL_INVALID)
			}
			if (originResDigit == int32(_K_AXES_DIGIT) || destinationResDigit == int32(_K_AXES_DIGIT)) && XisPentagon(tls, *(*TH3Index)(unsafe.Pointer(bp))) != 0 {
				// If these are invalid cells, fail rather than incorrectly
				// reporting neighbors. For pentagon cells that are actually
				// neighbors across the deleted subsequence, they will fail the
				// optimized check below, but they will be accepted by the
				// gridDisk check below that.
				return uint32(_E_CELL_INVALID)
			}
			// These sets are the relevant neighbors in the clockwise
			// and counter-clockwise
			neighborSetClockwise = [7]TDirection{
				1: int32(_JK_AXES_DIGIT),
				2: int32(_IJ_AXES_DIGIT),
				3: int32(_J_AXES_DIGIT),
				4: int32(_IK_AXES_DIGIT),
				5: int32(_K_AXES_DIGIT),
				6: int32(_I_AXES_DIGIT),
			}
			neighborSetCounterclockwise = [7]TDirection{
				1: int32(_IK_AXES_DIGIT),
				2: int32(_JK_AXES_DIGIT),
				3: int32(_K_AXES_DIGIT),
				4: int32(_IJ_AXES_DIGIT),
				5: int32(_I_AXES_DIGIT),
				6: int32(_J_AXES_DIGIT),
			}
			if neighborSetClockwise[originResDigit] == destinationResDigit || neighborSetCounterclockwise[originResDigit] == destinationResDigit {
				*(*int32)(unsafe.Pointer(out)) = int32(1)
				return uint32(_E_SUCCESS)
			}
		}
	}
	// Otherwise, we have to determine the neighbor relationship the "hard" way.
	*(*[7]TH3Index)(unsafe.Pointer(bp + 16)) = [7]TH3Index{}
	XgridDisk(tls, origin, int32(1), bp+16)
	i = 0
	for {
		if !(i < int32(7)) {
			break
		}
		if (*(*[7]TH3Index)(unsafe.Pointer(bp + 16)))[i] == destination {
			*(*int32)(unsafe.Pointer(out)) = int32(1)
			return uint32(_E_SUCCESS)
		}
		goto _1
	_1:
		;
		i++
	}
	// Made it here, they definitely aren't neighbors
	*(*int32)(unsafe.Pointer(out)) = 0
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * Returns a directed edge H3 index based on the provided origin and
//	 * destination
//	 * @param origin The origin H3 hexagon index
//	 * @param destination The destination H3 hexagon index
//	 * @param out Output: The directed edge H3Index.
//	 */
func XcellsToDirectedEdge(tls *libc.TLS, origin TH3Index, destination TH3Index, out uintptr) (r TH3Error) {
	var direction TDirection
	var output TH3Index
	_, _ = direction, output
	// Determine the IJK direction from the origin to the destination
	direction = XdirectionForNeighbor(tls, origin, destination)
	// The direction will be invalid if the cells are not neighbors
	if direction == int32(_INVALID_DIGIT) {
		return uint32(_E_NOT_NEIGHBORS)
	}
	// Create the edge index for the neighbor direction
	output = origin
	output = output & ^(libc.Uint64FromInt32(libc.Int32FromInt32(15))<<libc.Int32FromInt32(DH3_MODE_OFFSET)) | libc.Uint64FromInt32(libc.Int32FromInt32(DH3_DIRECTEDEDGE_MODE))<<libc.Int32FromInt32(DH3_MODE_OFFSET)
	output = output & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<libc.Int32FromInt32(DH3_RESERVED_OFFSET)) | libc.Uint64FromInt32(direction)<<libc.Int32FromInt32(DH3_RESERVED_OFFSET)
	*(*TH3Index)(unsafe.Pointer(out)) = output
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * Returns the origin hexagon from the directed edge H3Index
//	 * @param edge The edge H3 index
//	 * @param out Output: The origin H3 hexagon index
//	 */
func XgetDirectedEdgeOrigin(tls *libc.TLS, edge TH3Index, out uintptr) (r TH3Error) {
	var origin TH3Index
	_ = origin
	if libc.Int32FromUint64(edge&(libc.Uint64FromInt32(libc.Int32FromInt32(15))<<libc.Int32FromInt32(DH3_MODE_OFFSET))>>libc.Int32FromInt32(DH3_MODE_OFFSET)) != int32(DH3_DIRECTEDEDGE_MODE) {
		return uint32(_E_DIR_EDGE_INVALID)
	}
	origin = edge
	origin = origin & ^(libc.Uint64FromInt32(libc.Int32FromInt32(15))<<libc.Int32FromInt32(DH3_MODE_OFFSET)) | libc.Uint64FromInt32(libc.Int32FromInt32(DH3_CELL_MODE))<<libc.Int32FromInt32(DH3_MODE_OFFSET)
	origin = origin & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<libc.Int32FromInt32(DH3_RESERVED_OFFSET)) | libc.Uint64FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(DH3_RESERVED_OFFSET)
	*(*TH3Index)(unsafe.Pointer(out)) = origin
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * Returns the destination hexagon from the directed edge H3Index
//	 * @param edge The edge H3 index
//	 * @param out Output: The destination H3 hexagon index
//	 */
func XgetDirectedEdgeDestination(tls *libc.TLS, edge TH3Index, out uintptr) (r TH3Error) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var direction TDirection
	var originResult TH3Error
	var _ /* origin at bp+8 */ TH3Index
	var _ /* rotations at bp+0 */ int32
	_, _ = direction, originResult
	direction = libc.Int32FromUint64(edge & (libc.Uint64FromInt32(libc.Int32FromInt32(7)) << libc.Int32FromInt32(DH3_RESERVED_OFFSET)) >> libc.Int32FromInt32(DH3_RESERVED_OFFSET))
	*(*int32)(unsafe.Pointer(bp)) = 0
	// Note: This call is also checking for H3_DIRECTEDEDGE_MODE
	originResult = XgetDirectedEdgeOrigin(tls, edge, bp+8)
	if originResult != 0 {
		return originResult
	}
	return Xh3NeighborRotations(tls, *(*TH3Index)(unsafe.Pointer(bp + 8)), direction, bp, out)
}

// C documentation
//
//	/**
//	 * Determines if the provided H3Index is a valid directed edge index
//	 * @param edge The directed edge H3Index
//	 * @return 1 if it is a directed edge H3Index, otherwise 0.
//	 */
func XisValidDirectedEdge(tls *libc.TLS, edge TH3Index) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var neighborDirection TDirection
	var originResult TH3Error
	var _ /* origin at bp+0 */ TH3Index
	_, _ = neighborDirection, originResult
	neighborDirection = libc.Int32FromUint64(edge & (libc.Uint64FromInt32(libc.Int32FromInt32(7)) << libc.Int32FromInt32(DH3_RESERVED_OFFSET)) >> libc.Int32FromInt32(DH3_RESERVED_OFFSET))
	if neighborDirection <= int32(_CENTER_DIGIT) || neighborDirection >= int32(_NUM_DIGITS) {
		return 0
	}
	// Note: This call is also checking for H3_DIRECTEDEDGE_MODE
	originResult = XgetDirectedEdgeOrigin(tls, edge, bp)
	if originResult != 0 {
		return 0
	}
	if XisPentagon(tls, *(*TH3Index)(unsafe.Pointer(bp))) != 0 && neighborDirection == int32(_K_AXES_DIGIT) {
		return 0
	}
	return XisValidCell(tls, *(*TH3Index)(unsafe.Pointer(bp)))
}

// C documentation
//
//	/**
//	 * Returns the origin, destination pair of hexagon IDs for the given edge ID
//	 * @param edge The directed edge H3Index
//	 * @param originDestination Pointer to memory to store origin and destination
//	 * IDs
//	 */
func XdirectedEdgeToCells(tls *libc.TLS, edge TH3Index, originDestination uintptr) (r TH3Error) {
	var destinationResult, originResult TH3Error
	_, _ = destinationResult, originResult
	originResult = XgetDirectedEdgeOrigin(tls, edge, originDestination)
	if originResult != 0 {
		return originResult
	}
	destinationResult = XgetDirectedEdgeDestination(tls, edge, originDestination+1*8)
	if destinationResult != 0 {
		return destinationResult
	}
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * Provides all of the directed edges from the current H3Index.
//	 * @param origin The origin hexagon H3Index to find edges for.
//	 * @param edges The memory to store all of the edges inside.
//	 */
func XoriginToDirectedEdges(tls *libc.TLS, origin TH3Index, edges uintptr) (r TH3Error) {
	var i, isPent int32
	_, _ = i, isPent
	// Determine if the origin is a pentagon and special treatment needed.
	isPent = XisPentagon(tls, origin)
	// This is actually quite simple. Just modify the bits of the origin
	// slightly for each direction, except the 'k' direction in pentagons,
	// which is zeroed.
	i = 0
	for {
		if !(i < int32(6)) {
			break
		}
		if isPent != 0 && i == 0 {
			*(*TH3Index)(unsafe.Pointer(edges + uintptr(i)*8)) = uint64(DH3_NULL)
		} else {
			*(*TH3Index)(unsafe.Pointer(edges + uintptr(i)*8)) = origin
			*(*TH3Index)(unsafe.Pointer(edges + uintptr(i)*8)) = *(*TH3Index)(unsafe.Pointer(edges + uintptr(i)*8)) & ^(libc.Uint64FromInt32(libc.Int32FromInt32(15))<<libc.Int32FromInt32(DH3_MODE_OFFSET)) | libc.Uint64FromInt32(libc.Int32FromInt32(DH3_DIRECTEDEDGE_MODE))<<libc.Int32FromInt32(DH3_MODE_OFFSET)
			*(*TH3Index)(unsafe.Pointer(edges + uintptr(i)*8)) = *(*TH3Index)(unsafe.Pointer(edges + uintptr(i)*8)) & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<libc.Int32FromInt32(DH3_RESERVED_OFFSET)) | libc.Uint64FromInt32(i+libc.Int32FromInt32(1))<<libc.Int32FromInt32(DH3_RESERVED_OFFSET)
		}
		goto _1
	_1:
		;
		i++
	}
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * Provides the coordinates defining the directed edge.
//	 * @param edge The directed edge H3Index
//	 * @param cb The cellboundary object to store the edge coordinates.
//	 */
func XdirectedEdgeToBoundary(tls *libc.TLS, edge TH3Index, cb uintptr) (r TH3Error) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var direction TDirection
	var fijkResult, originResult TH3Error
	var isPent, res, startVertex, v1 int32
	var v2 bool
	var _ /* fijk at bp+8 */ TFaceIJK
	var _ /* origin at bp+0 */ TH3Index
	_, _, _, _, _, _, _, _ = direction, fijkResult, isPent, originResult, res, startVertex, v1, v2
	// Get the origin and neighbor direction from the edge
	direction = libc.Int32FromUint64(edge & (libc.Uint64FromInt32(libc.Int32FromInt32(7)) << libc.Int32FromInt32(DH3_RESERVED_OFFSET)) >> libc.Int32FromInt32(DH3_RESERVED_OFFSET))
	originResult = XgetDirectedEdgeOrigin(tls, edge, bp)
	if originResult != 0 {
		return originResult
	}
	// Get the start vertex for the edge
	startVertex = XvertexNumForDirection(tls, *(*TH3Index)(unsafe.Pointer(bp)), direction)
	if startVertex == -int32(1) {
		// This is not actually an edge (i.e. no valid direction),
		// so return no vertices.
		(*TCellBoundary)(unsafe.Pointer(cb)).FnumVerts = 0
		return uint32(_E_DIR_EDGE_INVALID)
	}
	fijkResult = X_h3ToFaceIjk(tls, *(*TH3Index)(unsafe.Pointer(bp)), bp+8)
	if fijkResult != 0 {
		if v2 = libc.Bool(0 != 0); !v2 {
			libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+55, int32(282), uintptr(unsafe.Pointer(&___func__1)))
		}
		_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
		v1 = libc.Int32FromInt32(1)
	} else {
		v1 = 0
	}
	if v1 != 0 {
		return fijkResult
	}
	res = libc.Int32FromUint64(*(*TH3Index)(unsafe.Pointer(bp)) & (libc.Uint64FromUint64(15) << libc.Int32FromInt32(DH3_RES_OFFSET)) >> libc.Int32FromInt32(DH3_RES_OFFSET))
	isPent = XisPentagon(tls, *(*TH3Index)(unsafe.Pointer(bp)))
	if isPent != 0 {
		X_faceIjkPentToCellBoundary(tls, bp+8, res, startVertex, int32(2), cb)
	} else {
		X_faceIjkToCellBoundary(tls, bp+8, res, startVertex, int32(2), cb)
	}
	return uint32(_E_SUCCESS)
}

var ___func__1 = [23]uint8{'d', 'i', 'r', 'e', 'c', 't', 'e', 'd', 'E', 'd', 'g', 'e', 'T', 'o', 'B', 'o', 'u', 'n', 'd', 'a', 'r', 'y'}

const DEPSILON1 = 1e-16
const DINV_RES0_U_GNOMONIC1 = 2.618033988749896
const DM_AP7_ROT_RADS1 = 0.3334731722518321
const DM_ONETHIRD1 = 0.3333333333333333
const DM_RSQRT7 = 0.37796447300922725
const DM_SQRT3_23 = 0.8660254037844386
const DM_SQRT7 = 2.6457513110645907
const DRES0_U_GNOMONIC1 = 0.381966011250105

var _UNIT_VECS5 = [7]TCoordIJK{
	0: {},
	1: {
		Fk: int32(1),
	},
	2: {
		Fj: int32(1),
	},
	3: {
		Fj: int32(1),
		Fk: int32(1),
	},
	4: {
		Fi: int32(1),
	},
	5: {
		Fi: int32(1),
		Fk: int32(1),
	},
	6: {
		Fi: int32(1),
		Fj: int32(1),
	},
}

type TVec3d = struct {
	Fx float64
	Fy float64
	Fz float64
}

// C documentation
//
//	/** @brief icosahedron face centers in x/y/z on the unit sphere */
var _faceCenterPoint = [20]TVec3d{
	0: {
		Fx: float64(0.2199307791404606),
		Fy: float64(0.6583691780274996),
		Fz: float64(0.7198475378926182),
	},
	1: {
		Fx: -libc.Float64FromFloat64(0.2139234834501421),
		Fy: float64(0.1478171829550703),
		Fz: float64(0.9656017935214205),
	},
	2: {
		Fx: float64(0.1092625278784797),
		Fy: -libc.Float64FromFloat64(0.481195157287321),
		Fz: float64(0.8697775121287253),
	},
	3: {
		Fx: float64(0.7428567301586791),
		Fy: -libc.Float64FromFloat64(0.3593941678278028),
		Fz: float64(0.5648005936517033),
	},
	4: {
		Fx: float64(0.8112534709140969),
		Fy: float64(0.3448953237639384),
		Fz: float64(0.472138773641393),
	},
	5: {
		Fx: -libc.Float64FromFloat64(0.1055498149613921),
		Fy: float64(0.9794457296411413),
		Fz: float64(0.1718874610009365),
	},
	6: {
		Fx: -libc.Float64FromFloat64(0.8075407579970092),
		Fy: float64(0.1533552485898818),
		Fz: float64(0.5695261994882688),
	},
	7: {
		Fx: -libc.Float64FromFloat64(0.2846148069787907),
		Fy: -libc.Float64FromFloat64(0.8644080972654206),
		Fz: float64(0.4144792552473539),
	},
	8: {
		Fx: float64(0.7405621473854482),
		Fy: -libc.Float64FromFloat64(0.6673299564565524),
		Fz: -libc.Float64FromFloat64(0.0789837646326737),
	},
	9: {
		Fx: float64(0.8512303986474293),
		Fy: float64(0.4722343788582681),
		Fz: -libc.Float64FromFloat64(0.2289137388687808),
	},
	10: {
		Fx: -libc.Float64FromFloat64(0.7405621473854481),
		Fy: float64(0.6673299564565524),
		Fz: float64(0.0789837646326737),
	},
	11: {
		Fx: -libc.Float64FromFloat64(0.8512303986474292),
		Fy: -libc.Float64FromFloat64(0.4722343788582682),
		Fz: float64(0.2289137388687808),
	},
	12: {
		Fx: float64(0.1055498149613919),
		Fy: -libc.Float64FromFloat64(0.9794457296411413),
		Fz: -libc.Float64FromFloat64(0.1718874610009365),
	},
	13: {
		Fx: float64(0.8075407579970092),
		Fy: -libc.Float64FromFloat64(0.1533552485898819),
		Fz: -libc.Float64FromFloat64(0.5695261994882688),
	},
	14: {
		Fx: float64(0.2846148069787908),
		Fy: float64(0.8644080972654204),
		Fz: -libc.Float64FromFloat64(0.4144792552473539),
	},
	15: {
		Fx: -libc.Float64FromFloat64(0.7428567301586791),
		Fy: float64(0.3593941678278027),
		Fz: -libc.Float64FromFloat64(0.5648005936517033),
	},
	16: {
		Fx: -libc.Float64FromFloat64(0.811253470914097),
		Fy: -libc.Float64FromFloat64(0.3448953237639382),
		Fz: -libc.Float64FromFloat64(0.472138773641393),
	},
	17: {
		Fx: -libc.Float64FromFloat64(0.2199307791404607),
		Fy: -libc.Float64FromFloat64(0.6583691780274996),
		Fz: -libc.Float64FromFloat64(0.7198475378926182),
	},
	18: {
		Fx: float64(0.213923483450142),
		Fy: -libc.Float64FromFloat64(0.1478171829550704),
		Fz: -libc.Float64FromFloat64(0.9656017935214205),
	},
	19: {
		Fx: -libc.Float64FromFloat64(0.1092625278784796),
		Fy: float64(0.481195157287321),
		Fz: -libc.Float64FromFloat64(0.8697775121287253),
	},
}

// C documentation
//
//	/** @brief icosahedron face ijk axes as azimuth in radians from face center to
//	 * vertex 0/1/2 respectively
//	 */
var _faceAxesAzRadsCII = [20][3]float64{
	0: {
		0: float64(5.6199582685239395),
		1: float64(3.5255631661307447),
		2: float64(1.4311680637375488),
	},
	1: {
		0: float64(5.7603390817141875),
		1: float64(3.665943979320992),
		2: float64(1.571548876927796),
	},
	2: {
		0: float64(0.78021365439343),
		1: float64(4.969003859179821),
		2: float64(2.8746087567866256),
	},
	3: {
		0: float64(0.4304693639799999),
		1: float64(4.619259568766391),
		2: float64(2.5248644663731956),
	},
	4: {
		0: float64(6.130269123335111),
		1: float64(4.0358740209419155),
		2: float64(1.9414789185487202),
	},
	5: {
		0: float64(2.692877706530643),
		1: float64(0.5984826041374471),
		2: float64(4.787272808923838),
	},
	6: {
		0: float64(2.982963003477244),
		1: float64(0.8885679010840484),
		2: float64(5.07735810587044),
	},
	7: {
		0: float64(3.532912002790141),
		1: float64(1.4385169003969456),
		2: float64(5.627307105183337),
	},
	8: {
		0: float64(3.494305004259568),
		1: float64(1.3999099018663728),
		2: float64(5.588700106652764),
	},
	9: {
		0: float64(3.0032141694995382),
		1: float64(0.908819067106343),
		2: float64(5.0976092718927335),
	},
	10: {
		0: float64(5.930472956509812),
		1: float64(3.836077854116616),
		2: float64(1.7416827517234204),
	},
	11: {
		0: float64(0.13837848409025486),
		1: float64(4.327168688876646),
		2: float64(2.23277358648345),
	},
	12: {
		0: float64(0.4487149470591504),
		1: float64(4.6375051518455415),
		2: float64(2.543110049452346),
	},
	13: {
		0: float64(0.15862965011254937),
		1: float64(4.3474198548989405),
		2: float64(2.2530247525057447),
	},
	14: {
		0: float64(5.891865957979238),
		1: float64(3.797470855586043),
		2: float64(1.7030757531928475),
	},
	15: {
		0: float64(2.711123289609793),
		1: float64(0.6167281872165977),
		2: float64(4.8055183920029885),
	},
	16: {
		0: float64(3.294508837434268),
		1: float64(1.2001137350410729),
		2: float64(5.388903939827464),
	},
	17: {
		0: float64(3.80481969224544),
		1: float64(1.7104245898522445),
		2: float64(5.8992147946386355),
	},
	18: {
		0: float64(3.6644388790551923),
		1: float64(1.570043776661997),
		2: float64(5.758833981448388),
	},
	19: {
		0: float64(2.361378999196363),
		1: float64(0.2669838968031676),
		2: float64(4.455774101589559),
	},
}

// C documentation
//
//	/** @brief Definition of which faces neighbor each other. */
var _faceNeighbors = [20][4]TFaceOrientIJK{
	0: {
		0: {},
		1: {
			Fface: int32(4),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(1),
		},
		2: {
			Fface: int32(1),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fj: int32(2),
			},
			FccwRot60: int32(5),
		},
		3: {
			Fface: int32(5),
			Ftranslate: TCoordIJK{
				Fj: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
	},
	1: {
		0: {
			Fface: int32(1),
		},
		1: {
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(1),
		},
		2: {
			Fface: int32(2),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fj: int32(2),
			},
			FccwRot60: int32(5),
		},
		3: {
			Fface: int32(6),
			Ftranslate: TCoordIJK{
				Fj: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
	},
	2: {
		0: {
			Fface: int32(2),
		},
		1: {
			Fface: int32(1),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(1),
		},
		2: {
			Fface: int32(3),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fj: int32(2),
			},
			FccwRot60: int32(5),
		},
		3: {
			Fface: int32(7),
			Ftranslate: TCoordIJK{
				Fj: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
	},
	3: {
		0: {
			Fface: int32(3),
		},
		1: {
			Fface: int32(2),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(1),
		},
		2: {
			Fface: int32(4),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fj: int32(2),
			},
			FccwRot60: int32(5),
		},
		3: {
			Fface: int32(8),
			Ftranslate: TCoordIJK{
				Fj: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
	},
	4: {
		0: {
			Fface: int32(4),
		},
		1: {
			Fface: int32(3),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(1),
		},
		2: {
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fj: int32(2),
			},
			FccwRot60: int32(5),
		},
		3: {
			Fface: int32(9),
			Ftranslate: TCoordIJK{
				Fj: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
	},
	5: {
		0: {
			Fface: int32(5),
		},
		1: {
			Fface: int32(10),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fj: int32(2),
			},
			FccwRot60: int32(3),
		},
		2: {
			Fface: int32(14),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
		3: {
			Ftranslate: TCoordIJK{
				Fj: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
	},
	6: {
		0: {
			Fface: int32(6),
		},
		1: {
			Fface: int32(11),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fj: int32(2),
			},
			FccwRot60: int32(3),
		},
		2: {
			Fface: int32(10),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
		3: {
			Fface: int32(1),
			Ftranslate: TCoordIJK{
				Fj: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
	},
	7: {
		0: {
			Fface: int32(7),
		},
		1: {
			Fface: int32(12),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fj: int32(2),
			},
			FccwRot60: int32(3),
		},
		2: {
			Fface: int32(11),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
		3: {
			Fface: int32(2),
			Ftranslate: TCoordIJK{
				Fj: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
	},
	8: {
		0: {
			Fface: int32(8),
		},
		1: {
			Fface: int32(13),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fj: int32(2),
			},
			FccwRot60: int32(3),
		},
		2: {
			Fface: int32(12),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
		3: {
			Fface: int32(3),
			Ftranslate: TCoordIJK{
				Fj: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
	},
	9: {
		0: {
			Fface: int32(9),
		},
		1: {
			Fface: int32(14),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fj: int32(2),
			},
			FccwRot60: int32(3),
		},
		2: {
			Fface: int32(13),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
		3: {
			Fface: int32(4),
			Ftranslate: TCoordIJK{
				Fj: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
	},
	10: {
		0: {
			Fface: int32(10),
		},
		1: {
			Fface: int32(5),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fj: int32(2),
			},
			FccwRot60: int32(3),
		},
		2: {
			Fface: int32(6),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
		3: {
			Fface: int32(15),
			Ftranslate: TCoordIJK{
				Fj: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
	},
	11: {
		0: {
			Fface: int32(11),
		},
		1: {
			Fface: int32(6),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fj: int32(2),
			},
			FccwRot60: int32(3),
		},
		2: {
			Fface: int32(7),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
		3: {
			Fface: int32(16),
			Ftranslate: TCoordIJK{
				Fj: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
	},
	12: {
		0: {
			Fface: int32(12),
		},
		1: {
			Fface: int32(7),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fj: int32(2),
			},
			FccwRot60: int32(3),
		},
		2: {
			Fface: int32(8),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
		3: {
			Fface: int32(17),
			Ftranslate: TCoordIJK{
				Fj: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
	},
	13: {
		0: {
			Fface: int32(13),
		},
		1: {
			Fface: int32(8),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fj: int32(2),
			},
			FccwRot60: int32(3),
		},
		2: {
			Fface: int32(9),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
		3: {
			Fface: int32(18),
			Ftranslate: TCoordIJK{
				Fj: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
	},
	14: {
		0: {
			Fface: int32(14),
		},
		1: {
			Fface: int32(9),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fj: int32(2),
			},
			FccwRot60: int32(3),
		},
		2: {
			Fface: int32(5),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
		3: {
			Fface: int32(19),
			Ftranslate: TCoordIJK{
				Fj: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
	},
	15: {
		0: {
			Fface: int32(15),
		},
		1: {
			Fface: int32(16),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(1),
		},
		2: {
			Fface: int32(19),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fj: int32(2),
			},
			FccwRot60: int32(5),
		},
		3: {
			Fface: int32(10),
			Ftranslate: TCoordIJK{
				Fj: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
	},
	16: {
		0: {
			Fface: int32(16),
		},
		1: {
			Fface: int32(17),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(1),
		},
		2: {
			Fface: int32(15),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fj: int32(2),
			},
			FccwRot60: int32(5),
		},
		3: {
			Fface: int32(11),
			Ftranslate: TCoordIJK{
				Fj: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
	},
	17: {
		0: {
			Fface: int32(17),
		},
		1: {
			Fface: int32(18),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(1),
		},
		2: {
			Fface: int32(16),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fj: int32(2),
			},
			FccwRot60: int32(5),
		},
		3: {
			Fface: int32(12),
			Ftranslate: TCoordIJK{
				Fj: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
	},
	18: {
		0: {
			Fface: int32(18),
		},
		1: {
			Fface: int32(19),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(1),
		},
		2: {
			Fface: int32(17),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fj: int32(2),
			},
			FccwRot60: int32(5),
		},
		3: {
			Fface: int32(13),
			Ftranslate: TCoordIJK{
				Fj: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
	},
	19: {
		0: {
			Fface: int32(19),
		},
		1: {
			Fface: int32(15),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(1),
		},
		2: {
			Fface: int32(18),
			Ftranslate: TCoordIJK{
				Fi: int32(2),
				Fj: int32(2),
			},
			FccwRot60: int32(5),
		},
		3: {
			Fface: int32(14),
			Ftranslate: TCoordIJK{
				Fj: int32(2),
				Fk: int32(2),
			},
			FccwRot60: int32(3),
		},
	},
}

// C documentation
//
//	/** @brief direction from the origin face to the destination face, relative to
//	 * the origin face's coordinate system, or -1 if not adjacent.
//	 */
var _adjacentFaceDir = [20][20]int32{
	0: {
		1:  int32(DKI),
		2:  -int32(1),
		3:  -int32(1),
		4:  int32(DIJ),
		5:  int32(DJK),
		6:  -int32(1),
		7:  -int32(1),
		8:  -int32(1),
		9:  -int32(1),
		10: -int32(1),
		11: -int32(1),
		12: -int32(1),
		13: -int32(1),
		14: -int32(1),
		15: -int32(1),
		16: -int32(1),
		17: -int32(1),
		18: -int32(1),
		19: -int32(1),
	},
	1: {
		0:  int32(DIJ),
		2:  int32(DKI),
		3:  -int32(1),
		4:  -int32(1),
		5:  -int32(1),
		6:  int32(DJK),
		7:  -int32(1),
		8:  -int32(1),
		9:  -int32(1),
		10: -int32(1),
		11: -int32(1),
		12: -int32(1),
		13: -int32(1),
		14: -int32(1),
		15: -int32(1),
		16: -int32(1),
		17: -int32(1),
		18: -int32(1),
		19: -int32(1),
	},
	2: {
		0:  -int32(1),
		1:  int32(DIJ),
		3:  int32(DKI),
		4:  -int32(1),
		5:  -int32(1),
		6:  -int32(1),
		7:  int32(DJK),
		8:  -int32(1),
		9:  -int32(1),
		10: -int32(1),
		11: -int32(1),
		12: -int32(1),
		13: -int32(1),
		14: -int32(1),
		15: -int32(1),
		16: -int32(1),
		17: -int32(1),
		18: -int32(1),
		19: -int32(1),
	},
	3: {
		0:  -int32(1),
		1:  -int32(1),
		2:  int32(DIJ),
		4:  int32(DKI),
		5:  -int32(1),
		6:  -int32(1),
		7:  -int32(1),
		8:  int32(DJK),
		9:  -int32(1),
		10: -int32(1),
		11: -int32(1),
		12: -int32(1),
		13: -int32(1),
		14: -int32(1),
		15: -int32(1),
		16: -int32(1),
		17: -int32(1),
		18: -int32(1),
		19: -int32(1),
	},
	4: {
		0:  int32(DKI),
		1:  -int32(1),
		2:  -int32(1),
		3:  int32(DIJ),
		5:  -int32(1),
		6:  -int32(1),
		7:  -int32(1),
		8:  -int32(1),
		9:  int32(DJK),
		10: -int32(1),
		11: -int32(1),
		12: -int32(1),
		13: -int32(1),
		14: -int32(1),
		15: -int32(1),
		16: -int32(1),
		17: -int32(1),
		18: -int32(1),
		19: -int32(1),
	},
	5: {
		0:  int32(DJK),
		1:  -int32(1),
		2:  -int32(1),
		3:  -int32(1),
		4:  -int32(1),
		6:  -int32(1),
		7:  -int32(1),
		8:  -int32(1),
		9:  -int32(1),
		10: int32(DIJ),
		11: -int32(1),
		12: -int32(1),
		13: -int32(1),
		14: int32(DKI),
		15: -int32(1),
		16: -int32(1),
		17: -int32(1),
		18: -int32(1),
		19: -int32(1),
	},
	6: {
		0:  -int32(1),
		1:  int32(DJK),
		2:  -int32(1),
		3:  -int32(1),
		4:  -int32(1),
		5:  -int32(1),
		7:  -int32(1),
		8:  -int32(1),
		9:  -int32(1),
		10: int32(DKI),
		11: int32(DIJ),
		12: -int32(1),
		13: -int32(1),
		14: -int32(1),
		15: -int32(1),
		16: -int32(1),
		17: -int32(1),
		18: -int32(1),
		19: -int32(1),
	},
	7: {
		0:  -int32(1),
		1:  -int32(1),
		2:  int32(DJK),
		3:  -int32(1),
		4:  -int32(1),
		5:  -int32(1),
		6:  -int32(1),
		8:  -int32(1),
		9:  -int32(1),
		10: -int32(1),
		11: int32(DKI),
		12: int32(DIJ),
		13: -int32(1),
		14: -int32(1),
		15: -int32(1),
		16: -int32(1),
		17: -int32(1),
		18: -int32(1),
		19: -int32(1),
	},
	8: {
		0:  -int32(1),
		1:  -int32(1),
		2:  -int32(1),
		3:  int32(DJK),
		4:  -int32(1),
		5:  -int32(1),
		6:  -int32(1),
		7:  -int32(1),
		9:  -int32(1),
		10: -int32(1),
		11: -int32(1),
		12: int32(DKI),
		13: int32(DIJ),
		14: -int32(1),
		15: -int32(1),
		16: -int32(1),
		17: -int32(1),
		18: -int32(1),
		19: -int32(1),
	},
	9: {
		0:  -int32(1),
		1:  -int32(1),
		2:  -int32(1),
		3:  -int32(1),
		4:  int32(DJK),
		5:  -int32(1),
		6:  -int32(1),
		7:  -int32(1),
		8:  -int32(1),
		10: -int32(1),
		11: -int32(1),
		12: -int32(1),
		13: int32(DKI),
		14: int32(DIJ),
		15: -int32(1),
		16: -int32(1),
		17: -int32(1),
		18: -int32(1),
		19: -int32(1),
	},
	10: {
		0:  -int32(1),
		1:  -int32(1),
		2:  -int32(1),
		3:  -int32(1),
		4:  -int32(1),
		5:  int32(DIJ),
		6:  int32(DKI),
		7:  -int32(1),
		8:  -int32(1),
		9:  -int32(1),
		11: -int32(1),
		12: -int32(1),
		13: -int32(1),
		14: -int32(1),
		15: int32(DJK),
		16: -int32(1),
		17: -int32(1),
		18: -int32(1),
		19: -int32(1),
	},
	11: {
		0:  -int32(1),
		1:  -int32(1),
		2:  -int32(1),
		3:  -int32(1),
		4:  -int32(1),
		5:  -int32(1),
		6:  int32(DIJ),
		7:  int32(DKI),
		8:  -int32(1),
		9:  -int32(1),
		10: -int32(1),
		12: -int32(1),
		13: -int32(1),
		14: -int32(1),
		15: -int32(1),
		16: int32(DJK),
		17: -int32(1),
		18: -int32(1),
		19: -int32(1),
	},
	12: {
		0:  -int32(1),
		1:  -int32(1),
		2:  -int32(1),
		3:  -int32(1),
		4:  -int32(1),
		5:  -int32(1),
		6:  -int32(1),
		7:  int32(DIJ),
		8:  int32(DKI),
		9:  -int32(1),
		10: -int32(1),
		11: -int32(1),
		13: -int32(1),
		14: -int32(1),
		15: -int32(1),
		16: -int32(1),
		17: int32(DJK),
		18: -int32(1),
		19: -int32(1),
	},
	13: {
		0:  -int32(1),
		1:  -int32(1),
		2:  -int32(1),
		3:  -int32(1),
		4:  -int32(1),
		5:  -int32(1),
		6:  -int32(1),
		7:  -int32(1),
		8:  int32(DIJ),
		9:  int32(DKI),
		10: -int32(1),
		11: -int32(1),
		12: -int32(1),
		14: -int32(1),
		15: -int32(1),
		16: -int32(1),
		17: -int32(1),
		18: int32(DJK),
		19: -int32(1),
	},
	14: {
		0:  -int32(1),
		1:  -int32(1),
		2:  -int32(1),
		3:  -int32(1),
		4:  -int32(1),
		5:  int32(DKI),
		6:  -int32(1),
		7:  -int32(1),
		8:  -int32(1),
		9:  int32(DIJ),
		10: -int32(1),
		11: -int32(1),
		12: -int32(1),
		13: -int32(1),
		15: -int32(1),
		16: -int32(1),
		17: -int32(1),
		18: -int32(1),
		19: int32(DJK),
	},
	15: {
		0:  -int32(1),
		1:  -int32(1),
		2:  -int32(1),
		3:  -int32(1),
		4:  -int32(1),
		5:  -int32(1),
		6:  -int32(1),
		7:  -int32(1),
		8:  -int32(1),
		9:  -int32(1),
		10: int32(DJK),
		11: -int32(1),
		12: -int32(1),
		13: -int32(1),
		14: -int32(1),
		16: int32(DIJ),
		17: -int32(1),
		18: -int32(1),
		19: int32(DKI),
	},
	16: {
		0:  -int32(1),
		1:  -int32(1),
		2:  -int32(1),
		3:  -int32(1),
		4:  -int32(1),
		5:  -int32(1),
		6:  -int32(1),
		7:  -int32(1),
		8:  -int32(1),
		9:  -int32(1),
		10: -int32(1),
		11: int32(DJK),
		12: -int32(1),
		13: -int32(1),
		14: -int32(1),
		15: int32(DKI),
		17: int32(DIJ),
		18: -int32(1),
		19: -int32(1),
	},
	17: {
		0:  -int32(1),
		1:  -int32(1),
		2:  -int32(1),
		3:  -int32(1),
		4:  -int32(1),
		5:  -int32(1),
		6:  -int32(1),
		7:  -int32(1),
		8:  -int32(1),
		9:  -int32(1),
		10: -int32(1),
		11: -int32(1),
		12: int32(DJK),
		13: -int32(1),
		14: -int32(1),
		15: -int32(1),
		16: int32(DKI),
		18: int32(DIJ),
		19: -int32(1),
	},
	18: {
		0:  -int32(1),
		1:  -int32(1),
		2:  -int32(1),
		3:  -int32(1),
		4:  -int32(1),
		5:  -int32(1),
		6:  -int32(1),
		7:  -int32(1),
		8:  -int32(1),
		9:  -int32(1),
		10: -int32(1),
		11: -int32(1),
		12: -int32(1),
		13: int32(DJK),
		14: -int32(1),
		15: -int32(1),
		16: -int32(1),
		17: int32(DKI),
		19: int32(DIJ),
	},
	19: {
		0:  -int32(1),
		1:  -int32(1),
		2:  -int32(1),
		3:  -int32(1),
		4:  -int32(1),
		5:  -int32(1),
		6:  -int32(1),
		7:  -int32(1),
		8:  -int32(1),
		9:  -int32(1),
		10: -int32(1),
		11: -int32(1),
		12: -int32(1),
		13: -int32(1),
		14: int32(DJK),
		15: int32(DIJ),
		16: -int32(1),
		17: -int32(1),
		18: int32(DKI),
	},
}

// C documentation
//
//	/** @brief overage distance table */
var _maxDimByCIIres = [17]int32{
	0:  int32(2),
	1:  -int32(1),
	2:  int32(14),
	3:  -int32(1),
	4:  int32(98),
	5:  -int32(1),
	6:  int32(686),
	7:  -int32(1),
	8:  int32(4802),
	9:  -int32(1),
	10: int32(33614),
	11: -int32(1),
	12: int32(235298),
	13: -int32(1),
	14: int32(1647086),
	15: -int32(1),
	16: int32(11529602),
}

// C documentation
//
//	/** @brief unit scale distance table */
var _unitScaleByCIIres = [17]int32{
	0:  int32(1),
	1:  -int32(1),
	2:  int32(7),
	3:  -int32(1),
	4:  int32(49),
	5:  -int32(1),
	6:  int32(343),
	7:  -int32(1),
	8:  int32(2401),
	9:  -int32(1),
	10: int32(16807),
	11: -int32(1),
	12: int32(117649),
	13: -int32(1),
	14: int32(823543),
	15: -int32(1),
	16: int32(5764801),
}

// C documentation
//
//	/**
//	 * Encodes a coordinate on the sphere to the FaceIJK address of the containing
//	 * cell at the specified resolution.
//	 *
//	 * @param g The spherical coordinates to encode.
//	 * @param res The desired H3 resolution for the encoding.
//	 * @param h The FaceIJK address of the containing cell at resolution res.
//	 */
func X_geoToFaceIjk(tls *libc.TLS, g uintptr, res int32, h uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* v at bp+0 */ TVec2d
	X_geoToHex2d(tls, g, res, h, bp)
	// then convert to ijk+
	X_hex2dToCoordIJK(tls, bp, h+4)
}

// C documentation
//
//	/**
//	 * Encodes a coordinate on the sphere to the corresponding icosahedral face and
//	 * containing 2D hex coordinates relative to that face center.
//	 *
//	 * @param g The spherical coordinates to encode.
//	 * @param res The desired H3 resolution for the encoding.
//	 * @param face The icosahedral face containing the spherical coordinates.
//	 * @param v The 2D hex coordinates of the cell containing the point.
//	 */
func X_geoToHex2d(tls *libc.TLS, g uintptr, res int32, face uintptr, v uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i int32
	var r, theta, v1 float64
	var _ /* sqd at bp+0 */ float64
	_, _, _, _ = i, r, theta, v1
	X_geoToClosestFace(tls, g, face, bp)
	// cos(r) = 1 - 2 * sin^2(r/2) = 1 - 2 * (sqd / 4) = 1 - sqd/2
	r = libc.Xacos(tls, libc.Float64FromInt32(1)-*(*float64)(unsafe.Pointer(bp))*float64(0.5))
	if r < float64(1e-16) {
		v1 = libc.Float64FromFloat64(0)
		(*TVec2d)(unsafe.Pointer(v)).Fy = v1
		(*TVec2d)(unsafe.Pointer(v)).Fx = v1
		return
	}
	// now have face and r, now find CCW theta from CII i-axis
	theta = X_posAngleRads(tls, *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(&_faceAxesAzRadsCII)) + uintptr(*(*int32)(unsafe.Pointer(face)))*24))-X_posAngleRads(tls, X_geoAzimuthRads(tls, uintptr(unsafe.Pointer(&XfaceCenterGeo))+uintptr(*(*int32)(unsafe.Pointer(face)))*16, g)))
	// adjust theta for Class III (odd resolutions)
	if XisResolutionClassIII(tls, res) != 0 {
		theta = X_posAngleRads(tls, theta-float64(0.3334731722518321))
	}
	// perform gnomonic scaling of r
	r = libc.Xtan(tls, r)
	// scale for current resolution length u
	r *= float64(2.618033988749896)
	i = 0
	for {
		if !(i < res) {
			break
		}
		r *= float64(2.6457513110645907)
		goto _2
	_2:
		;
		i++
	}
	// we now have (r, theta) in hex2d with theta ccw from x-axes
	// convert to local x,y
	(*TVec2d)(unsafe.Pointer(v)).Fx = r * libc.Xcos(tls, theta)
	(*TVec2d)(unsafe.Pointer(v)).Fy = r * libc.Xsin(tls, theta)
}

// C documentation
//
//	/**
//	 * Determines the center point in spherical coordinates of a cell given by 2D
//	 * hex coordinates on a particular icosahedral face.
//	 *
//	 * @param v The 2D hex coordinates of the cell.
//	 * @param face The icosahedral face upon which the 2D hex coordinate system is
//	 *             centered.
//	 * @param res The H3 resolution of the cell.
//	 * @param substrate Indicates whether or not this grid is actually a substrate
//	 *        grid relative to the specified resolution.
//	 * @param g The spherical coordinates of the cell center point.
//	 */
func X_hex2dToGeo(tls *libc.TLS, v uintptr, face int32, res int32, substrate int32, g uintptr) {
	var i int32
	var r, theta float64
	_, _, _ = i, r, theta
	// calculate (r, theta) in hex2d
	r = X_v2dMag(tls, v)
	if r < float64(1e-16) {
		*(*TLatLng)(unsafe.Pointer(g)) = XfaceCenterGeo[face]
		return
	}
	theta = libc.Xatan2(tls, (*TVec2d)(unsafe.Pointer(v)).Fy, (*TVec2d)(unsafe.Pointer(v)).Fx)
	// scale for current resolution length u
	i = 0
	for {
		if !(i < res) {
			break
		}
		r *= float64(0.37796447300922725)
		goto _1
	_1:
		;
		i++
	}
	// scale accordingly if this is a substrate grid
	if substrate != 0 {
		r *= float64(0.3333333333333333)
		if XisResolutionClassIII(tls, res) != 0 {
			r *= float64(0.37796447300922725)
		}
	}
	r *= float64(0.381966011250105)
	// perform inverse gnomonic scaling of r
	r = libc.Xatan(tls, r)
	// adjust theta for Class III
	// if a substrate grid, then it's already been adjusted for Class III
	if !(substrate != 0) && XisResolutionClassIII(tls, res) != 0 {
		theta = X_posAngleRads(tls, theta+float64(0.3334731722518321))
	}
	// find theta as an azimuth
	theta = X_posAngleRads(tls, *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(&_faceAxesAzRadsCII)) + uintptr(face)*24))-theta)
	// now find the point at (r,theta) from the face center
	X_geoAzDistanceRads(tls, uintptr(unsafe.Pointer(&XfaceCenterGeo))+uintptr(face)*16, theta, r, g)
}

// C documentation
//
//	/**
//	 * Determines the center point in spherical coordinates of a cell given by
//	 * a FaceIJK address at a specified resolution.
//	 *
//	 * @param h The FaceIJK address of the cell.
//	 * @param res The H3 resolution of the cell.
//	 * @param g The spherical coordinates of the cell center point.
//	 */
func X_faceIjkToGeo(tls *libc.TLS, h uintptr, res int32, g uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* v at bp+0 */ TVec2d
	X_ijkToHex2d(tls, h+4, bp)
	X_hex2dToGeo(tls, bp, (*TFaceIJK)(unsafe.Pointer(h)).Fface, res, 0, g)
}

// C documentation
//
//	/**
//	 * Generates the cell boundary in spherical coordinates for a pentagonal cell
//	 * given by a FaceIJK address at a specified resolution.
//	 *
//	 * @param h The FaceIJK address of the pentagonal cell.
//	 * @param res The H3 resolution of the cell.
//	 * @param start The first topological vertex to return.
//	 * @param length The number of topological vertexes to return.
//	 * @param g The spherical coordinates of the cell boundary.
//	 */
func X_faceIjkPentToCellBoundary(tls *libc.TLS, h uintptr, res int32, start int32, length int32, g uintptr) {
	bp := tls.Alloc(288)
	defer tls.Free(288)
	var additionalIteration, currentToLastDir, i, maxDim, v, vert, v1 int32
	var edge0, edge1, fijkOrient, ijk uintptr
	var v4 bool
	var _ /* adjRes at bp+0 */ int32
	var _ /* centerIJK at bp+4 */ TFaceIJK
	var _ /* fijk at bp+116 */ TFaceIJK
	var _ /* fijkVerts at bp+20 */ [5]TFaceIJK
	var _ /* inter at bp+248 */ TVec2d
	var _ /* lastFijk at bp+100 */ TFaceIJK
	var _ /* orig2d0 at bp+152 */ TVec2d
	var _ /* orig2d1 at bp+184 */ TVec2d
	var _ /* tmpFijk at bp+132 */ TFaceIJK
	var _ /* transVec at bp+168 */ TCoordIJK
	var _ /* v0 at bp+200 */ TVec2d
	var _ /* v1 at bp+216 */ TVec2d
	var _ /* v2 at bp+232 */ TVec2d
	var _ /* vec at bp+264 */ TVec2d
	_, _, _, _, _, _, _, _, _, _, _, _ = additionalIteration, currentToLastDir, edge0, edge1, fijkOrient, i, ijk, maxDim, v, vert, v1, v4
	*(*int32)(unsafe.Pointer(bp)) = res
	*(*TFaceIJK)(unsafe.Pointer(bp + 4)) = *(*TFaceIJK)(unsafe.Pointer(h))
	X_faceIjkPentToVerts(tls, bp+4, bp, bp+20)
	if length == int32(DNUM_PENT_VERTS) {
		v1 = int32(1)
	} else {
		v1 = 0
	}
	// If we're returning the entire loop, we need one more iteration in case
	// of a distortion vertex on the last edge
	additionalIteration = v1
	// convert each vertex to lat/lng
	// adjust the face of each vertex as appropriate and introduce
	// edge-crossing vertices as needed
	(*TCellBoundary)(unsafe.Pointer(g)).FnumVerts = 0
	vert = start
	for {
		if !(vert < start+length+additionalIteration) {
			break
		}
		v = vert % int32(DNUM_PENT_VERTS)
		*(*TFaceIJK)(unsafe.Pointer(bp + 116)) = (*(*[5]TFaceIJK)(unsafe.Pointer(bp + 20)))[v]
		X_adjustPentVertOverage(tls, bp+116, *(*int32)(unsafe.Pointer(bp)))
		// all Class III pentagon edges cross icosa edges
		// note that Class II pentagons have vertices on the edge,
		// not edge intersections
		if XisResolutionClassIII(tls, res) != 0 && vert > start {
			// find hex2d of the two vertexes on the last face
			*(*TFaceIJK)(unsafe.Pointer(bp + 132)) = *(*TFaceIJK)(unsafe.Pointer(bp + 116))
			X_ijkToHex2d(tls, bp+100+4, bp+152)
			currentToLastDir = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_adjacentFaceDir)) + uintptr((*(*TFaceIJK)(unsafe.Pointer(bp + 132))).Fface)*80 + uintptr((*(*TFaceIJK)(unsafe.Pointer(bp + 100))).Fface)*4))
			fijkOrient = uintptr(unsafe.Pointer(&_faceNeighbors)) + uintptr((*(*TFaceIJK)(unsafe.Pointer(bp + 132))).Fface)*80 + uintptr(currentToLastDir)*20
			(*(*TFaceIJK)(unsafe.Pointer(bp + 132))).Fface = (*TFaceOrientIJK)(unsafe.Pointer(fijkOrient)).Fface
			ijk = bp + 132 + 4
			// rotate and translate for adjacent face
			i = 0
			for {
				if !(i < (*TFaceOrientIJK)(unsafe.Pointer(fijkOrient)).FccwRot60) {
					break
				}
				X_ijkRotate60ccw(tls, ijk)
				goto _3
			_3:
				;
				i++
			}
			*(*TCoordIJK)(unsafe.Pointer(bp + 168)) = (*TFaceOrientIJK)(unsafe.Pointer(fijkOrient)).Ftranslate
			X_ijkScale(tls, bp+168, _unitScaleByCIIres[*(*int32)(unsafe.Pointer(bp))]*int32(3))
			X_ijkAdd(tls, ijk, bp+168, ijk)
			X_ijkNormalize(tls, ijk)
			X_ijkToHex2d(tls, ijk, bp+184)
			// find the appropriate icosa face edge vertexes
			maxDim = _maxDimByCIIres[*(*int32)(unsafe.Pointer(bp))]
			*(*TVec2d)(unsafe.Pointer(bp + 200)) = TVec2d{
				Fx: float64(3) * float64(maxDim),
			}
			*(*TVec2d)(unsafe.Pointer(bp + 216)) = TVec2d{
				Fx: -libc.Float64FromFloat64(1.5) * float64(maxDim),
				Fy: libc.Float64FromFloat64(3) * libc.Float64FromFloat64(0.8660254037844386) * float64(maxDim),
			}
			*(*TVec2d)(unsafe.Pointer(bp + 232)) = TVec2d{
				Fx: -libc.Float64FromFloat64(1.5) * float64(maxDim),
				Fy: -libc.Float64FromFloat64(3) * float64(0.8660254037844386) * float64(maxDim),
			}
			switch *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_adjacentFaceDir)) + uintptr((*(*TFaceIJK)(unsafe.Pointer(bp + 132))).Fface)*80 + uintptr((*(*TFaceIJK)(unsafe.Pointer(bp + 116))).Fface)*4)) {
			case int32(DIJ):
				edge0 = bp + 200
				edge1 = bp + 216
			case int32(DJK):
				edge0 = bp + 216
				edge1 = bp + 232
			case int32(DKI):
				fallthrough
			default:
				if v4 = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_adjacentFaceDir)) + uintptr((*(*TFaceIJK)(unsafe.Pointer(bp + 132))).Fface)*80 + uintptr((*(*TFaceIJK)(unsafe.Pointer(bp + 116))).Fface)*4)) == int32(DKI); !v4 {
					libc.X__assert_fail(tls, __ccgo_ts+87, __ccgo_ts+134, int32(572), uintptr(unsafe.Pointer(&___func__3)))
				}
				_ = v4 || libc.Bool(libc.Int32FromInt32(0) != 0)
				edge0 = bp + 232
				edge1 = bp + 200
				break
			}
			X_v2dIntersect(tls, bp+152, bp+184, edge0, edge1, bp+248)
			X_hex2dToGeo(tls, bp+248, (*(*TFaceIJK)(unsafe.Pointer(bp + 132))).Fface, *(*int32)(unsafe.Pointer(bp)), int32(1), g+8+uintptr((*TCellBoundary)(unsafe.Pointer(g)).FnumVerts)*16)
			(*TCellBoundary)(unsafe.Pointer(g)).FnumVerts++
		}
		// convert vertex to lat/lng and add to the result
		// vert == start + NUM_PENT_VERTS is only used to test for possible
		// intersection on last edge
		if vert < start+int32(DNUM_PENT_VERTS) {
			X_ijkToHex2d(tls, bp+116+4, bp+264)
			X_hex2dToGeo(tls, bp+264, (*(*TFaceIJK)(unsafe.Pointer(bp + 116))).Fface, *(*int32)(unsafe.Pointer(bp)), int32(1), g+8+uintptr((*TCellBoundary)(unsafe.Pointer(g)).FnumVerts)*16)
			(*TCellBoundary)(unsafe.Pointer(g)).FnumVerts++
		}
		*(*TFaceIJK)(unsafe.Pointer(bp + 100)) = *(*TFaceIJK)(unsafe.Pointer(bp + 116))
		goto _2
	_2:
		;
		vert++
	}
}

var ___func__3 = [27]uint8{'_', 'f', 'a', 'c', 'e', 'I', 'j', 'k', 'P', 'e', 'n', 't', 'T', 'o', 'C', 'e', 'l', 'l', 'B', 'o', 'u', 'n', 'd', 'a', 'r', 'y'}

// C documentation
//
//	/**
//	 * Get the vertices of a pentagon cell as substrate FaceIJK addresses
//	 *
//	 * @param fijk The FaceIJK address of the cell.
//	 * @param res The H3 resolution of the cell. This may be adjusted if
//	 *            necessary for the substrate grid resolution.
//	 * @param fijkVerts Output array for the vertices
//	 */
func X_faceIjkPentToVerts(tls *libc.TLS, fijk uintptr, res uintptr, fijkVerts uintptr) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var v int32
	var verts uintptr
	var _ /* vertsCII at bp+0 */ [5]TCoordIJK
	var _ /* vertsCIII at bp+60 */ [5]TCoordIJK
	_, _ = v, verts
	// the vertexes of an origin-centered pentagon in a Class II resolution on a
	// substrate grid with aperture sequence 33r. The aperture 3 gets us the
	// vertices, and the 3r gets us back to Class II.
	// vertices listed ccw from the i-axes
	*(*[5]TCoordIJK)(unsafe.Pointer(bp)) = [5]TCoordIJK{
		0: {
			Fi: int32(2),
			Fj: int32(1),
		},
		1: {
			Fi: int32(1),
			Fj: int32(2),
		},
		2: {
			Fj: int32(2),
			Fk: int32(1),
		},
		3: {
			Fj: int32(1),
			Fk: int32(2),
		},
		4: {
			Fi: int32(1),
			Fk: int32(2),
		},
	}
	// the vertexes of an origin-centered pentagon in a Class III resolution on
	// a substrate grid with aperture sequence 33r7r. The aperture 3 gets us the
	// vertices, and the 3r7r gets us to Class II. vertices listed ccw from the
	// i-axes
	*(*[5]TCoordIJK)(unsafe.Pointer(bp + 60)) = [5]TCoordIJK{
		0: {
			Fi: int32(5),
			Fj: int32(4),
		},
		1: {
			Fi: int32(1),
			Fj: int32(5),
		},
		2: {
			Fj: int32(5),
			Fk: int32(4),
		},
		3: {
			Fj: int32(1),
			Fk: int32(5),
		},
		4: {
			Fi: int32(4),
			Fk: int32(5),
		},
	}
	if XisResolutionClassIII(tls, *(*int32)(unsafe.Pointer(res))) != 0 {
		verts = bp + 60
	} else {
		verts = bp
	}
	// adjust the center point to be in an aperture 33r substrate grid
	// these should be composed for speed
	X_downAp3(tls, fijk+4)
	X_downAp3r(tls, fijk+4)
	// if res is Class III we need to add a cw aperture 7 to get to
	// icosahedral Class II
	if XisResolutionClassIII(tls, *(*int32)(unsafe.Pointer(res))) != 0 {
		X_downAp7r(tls, fijk+4)
		*(*int32)(unsafe.Pointer(res)) += int32(1)
	}
	// The center point is now in the same substrate grid as the origin
	// cell vertices. Add the center point substate coordinates
	// to each vertex to translate the vertices to that cell.
	v = 0
	for {
		if !(v < int32(DNUM_PENT_VERTS)) {
			break
		}
		(*(*TFaceIJK)(unsafe.Pointer(fijkVerts + uintptr(v)*16))).Fface = (*TFaceIJK)(unsafe.Pointer(fijk)).Fface
		X_ijkAdd(tls, fijk+4, verts+uintptr(v)*12, fijkVerts+uintptr(v)*16+4)
		X_ijkNormalize(tls, fijkVerts+uintptr(v)*16+4)
		goto _1
	_1:
		;
		v++
	}
}

// C documentation
//
//	/**
//	 * Generates the cell boundary in spherical coordinates for a cell given by a
//	 * FaceIJK address at a specified resolution.
//	 *
//	 * @param h The FaceIJK address of the cell.
//	 * @param res The H3 resolution of the cell.
//	 * @param start The first topological vertex to return.
//	 * @param length The number of topological vertexes to return.
//	 * @param g The spherical coordinates of the cell boundary.
//	 */
func X_faceIjkToCellBoundary(tls *libc.TLS, h uintptr, res int32, start int32, length int32, g uintptr) {
	bp := tls.Alloc(256)
	defer tls.Free(256)
	var additionalIteration, face2, lastFace, lastV, maxDim, pentLeading4, v, vert, v1, v3 int32
	var edge0, edge1 uintptr
	var isIntersectionAtVertex uint8
	var lastOverage, overage TOverage
	var v4 bool
	var _ /* adjRes at bp+0 */ int32
	var _ /* centerIJK at bp+4 */ TFaceIJK
	var _ /* fijk at bp+116 */ TFaceIJK
	var _ /* fijkVerts at bp+20 */ [6]TFaceIJK
	var _ /* inter at bp+216 */ TVec2d
	var _ /* orig2d0 at bp+136 */ TVec2d
	var _ /* orig2d1 at bp+152 */ TVec2d
	var _ /* v0 at bp+168 */ TVec2d
	var _ /* v1 at bp+184 */ TVec2d
	var _ /* v2 at bp+200 */ TVec2d
	var _ /* vec at bp+232 */ TVec2d
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = additionalIteration, edge0, edge1, face2, isIntersectionAtVertex, lastFace, lastOverage, lastV, maxDim, overage, pentLeading4, v, vert, v1, v3, v4
	*(*int32)(unsafe.Pointer(bp)) = res
	*(*TFaceIJK)(unsafe.Pointer(bp + 4)) = *(*TFaceIJK)(unsafe.Pointer(h))
	X_faceIjkToVerts(tls, bp+4, bp, bp+20)
	if length == int32(DNUM_HEX_VERTS) {
		v1 = int32(1)
	} else {
		v1 = 0
	}
	// If we're returning the entire loop, we need one more iteration in case
	// of a distortion vertex on the last edge
	additionalIteration = v1
	// convert each vertex to lat/lng
	// adjust the face of each vertex as appropriate and introduce
	// edge-crossing vertices as needed
	(*TCellBoundary)(unsafe.Pointer(g)).FnumVerts = 0
	lastFace = -int32(1)
	lastOverage = int32(_NO_OVERAGE)
	vert = start
	for {
		if !(vert < start+length+additionalIteration) {
			break
		}
		v = vert % int32(DNUM_HEX_VERTS)
		*(*TFaceIJK)(unsafe.Pointer(bp + 116)) = (*(*[6]TFaceIJK)(unsafe.Pointer(bp + 20)))[v]
		pentLeading4 = 0
		overage = X_adjustOverageClassII(tls, bp+116, *(*int32)(unsafe.Pointer(bp)), pentLeading4, int32(1))
		/*
		   Check for edge-crossing. Each face of the underlying icosahedron is a
		   different projection plane. So if an edge of the hexagon crosses an
		   icosahedron edge, an additional vertex must be introduced at that
		   intersection point. Then each half of the cell edge can be projected
		   to geographic coordinates using the appropriate icosahedron face
		   projection. Note that Class II cell edges have vertices on the face
		   edge, with no edge line intersections.
		*/
		if XisResolutionClassIII(tls, res) != 0 && vert > start && (*(*TFaceIJK)(unsafe.Pointer(bp + 116))).Fface != lastFace && lastOverage != int32(_FACE_EDGE) {
			// find hex2d of the two vertexes on original face
			lastV = (v + int32(5)) % int32(DNUM_HEX_VERTS)
			X_ijkToHex2d(tls, bp+20+uintptr(lastV)*16+4, bp+136)
			X_ijkToHex2d(tls, bp+20+uintptr(v)*16+4, bp+152)
			// find the appropriate icosa face edge vertexes
			maxDim = _maxDimByCIIres[*(*int32)(unsafe.Pointer(bp))]
			*(*TVec2d)(unsafe.Pointer(bp + 168)) = TVec2d{
				Fx: float64(3) * float64(maxDim),
			}
			*(*TVec2d)(unsafe.Pointer(bp + 184)) = TVec2d{
				Fx: -libc.Float64FromFloat64(1.5) * float64(maxDim),
				Fy: libc.Float64FromFloat64(3) * libc.Float64FromFloat64(0.8660254037844386) * float64(maxDim),
			}
			*(*TVec2d)(unsafe.Pointer(bp + 200)) = TVec2d{
				Fx: -libc.Float64FromFloat64(1.5) * float64(maxDim),
				Fy: -libc.Float64FromFloat64(3) * float64(0.8660254037844386) * float64(maxDim),
			}
			if lastFace == (*(*TFaceIJK)(unsafe.Pointer(bp + 4))).Fface {
				v3 = (*(*TFaceIJK)(unsafe.Pointer(bp + 116))).Fface
			} else {
				v3 = lastFace
			}
			face2 = v3
			switch *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_adjacentFaceDir)) + uintptr((*(*TFaceIJK)(unsafe.Pointer(bp + 4))).Fface)*80 + uintptr(face2)*4)) {
			case int32(DIJ):
				edge0 = bp + 168
				edge1 = bp + 184
			case int32(DJK):
				edge0 = bp + 184
				edge1 = bp + 200
				break
				// case KI:
				fallthrough
			default:
				if v4 = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_adjacentFaceDir)) + uintptr((*(*TFaceIJK)(unsafe.Pointer(bp + 4))).Fface)*80 + uintptr(face2)*4)) == int32(DKI); !v4 {
					libc.X__assert_fail(tls, __ccgo_ts+161, __ccgo_ts+134, int32(737), uintptr(unsafe.Pointer(&___func__4)))
				}
				_ = v4 || libc.Bool(libc.Int32FromInt32(0) != 0)
				edge0 = bp + 200
				edge1 = bp + 168
				break
			}
			X_v2dIntersect(tls, bp+136, bp+152, edge0, edge1, bp+216)
			/*
			   If a point of intersection occurs at a hexagon vertex, then each
			   adjacent hexagon edge will lie completely on a single icosahedron
			   face, and no additional vertex is required.
			*/
			isIntersectionAtVertex = libc.BoolUint8(X_v2dAlmostEquals(tls, bp+136, bp+216) != 0 || X_v2dAlmostEquals(tls, bp+152, bp+216) != 0)
			if !(isIntersectionAtVertex != 0) {
				X_hex2dToGeo(tls, bp+216, (*(*TFaceIJK)(unsafe.Pointer(bp + 4))).Fface, *(*int32)(unsafe.Pointer(bp)), int32(1), g+8+uintptr((*TCellBoundary)(unsafe.Pointer(g)).FnumVerts)*16)
				(*TCellBoundary)(unsafe.Pointer(g)).FnumVerts++
			}
		}
		// convert vertex to lat/lng and add to the result
		// vert == start + NUM_HEX_VERTS is only used to test for possible
		// intersection on last edge
		if vert < start+int32(DNUM_HEX_VERTS) {
			X_ijkToHex2d(tls, bp+116+4, bp+232)
			X_hex2dToGeo(tls, bp+232, (*(*TFaceIJK)(unsafe.Pointer(bp + 116))).Fface, *(*int32)(unsafe.Pointer(bp)), int32(1), g+8+uintptr((*TCellBoundary)(unsafe.Pointer(g)).FnumVerts)*16)
			(*TCellBoundary)(unsafe.Pointer(g)).FnumVerts++
		}
		lastFace = (*(*TFaceIJK)(unsafe.Pointer(bp + 116))).Fface
		lastOverage = overage
		goto _2
	_2:
		;
		vert++
	}
}

var ___func__4 = [23]uint8{'_', 'f', 'a', 'c', 'e', 'I', 'j', 'k', 'T', 'o', 'C', 'e', 'l', 'l', 'B', 'o', 'u', 'n', 'd', 'a', 'r', 'y'}

// C documentation
//
//	/**
//	 * Get the vertices of a cell as substrate FaceIJK addresses
//	 *
//	 * @param fijk The FaceIJK address of the cell.
//	 * @param res The H3 resolution of the cell. This may be adjusted if
//	 *            necessary for the substrate grid resolution.
//	 * @param fijkVerts Output array for the vertices
//	 */
func X_faceIjkToVerts(tls *libc.TLS, fijk uintptr, res uintptr, fijkVerts uintptr) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var v int32
	var verts uintptr
	var _ /* vertsCII at bp+0 */ [6]TCoordIJK
	var _ /* vertsCIII at bp+72 */ [6]TCoordIJK
	_, _ = v, verts
	// the vertexes of an origin-centered cell in a Class II resolution on a
	// substrate grid with aperture sequence 33r. The aperture 3 gets us the
	// vertices, and the 3r gets us back to Class II.
	// vertices listed ccw from the i-axes
	*(*[6]TCoordIJK)(unsafe.Pointer(bp)) = [6]TCoordIJK{
		0: {
			Fi: int32(2),
			Fj: int32(1),
		},
		1: {
			Fi: int32(1),
			Fj: int32(2),
		},
		2: {
			Fj: int32(2),
			Fk: int32(1),
		},
		3: {
			Fj: int32(1),
			Fk: int32(2),
		},
		4: {
			Fi: int32(1),
			Fk: int32(2),
		},
		5: {
			Fi: int32(2),
			Fk: int32(1),
		},
	}
	// the vertexes of an origin-centered cell in a Class III resolution on a
	// substrate grid with aperture sequence 33r7r. The aperture 3 gets us the
	// vertices, and the 3r7r gets us to Class II.
	// vertices listed ccw from the i-axes
	*(*[6]TCoordIJK)(unsafe.Pointer(bp + 72)) = [6]TCoordIJK{
		0: {
			Fi: int32(5),
			Fj: int32(4),
		},
		1: {
			Fi: int32(1),
			Fj: int32(5),
		},
		2: {
			Fj: int32(5),
			Fk: int32(4),
		},
		3: {
			Fj: int32(1),
			Fk: int32(5),
		},
		4: {
			Fi: int32(4),
			Fk: int32(5),
		},
		5: {
			Fi: int32(5),
			Fk: int32(1),
		},
	}
	if XisResolutionClassIII(tls, *(*int32)(unsafe.Pointer(res))) != 0 {
		verts = bp + 72
	} else {
		verts = bp
	}
	// adjust the center point to be in an aperture 33r substrate grid
	// these should be composed for speed
	X_downAp3(tls, fijk+4)
	X_downAp3r(tls, fijk+4)
	// if res is Class III we need to add a cw aperture 7 to get to
	// icosahedral Class II
	if XisResolutionClassIII(tls, *(*int32)(unsafe.Pointer(res))) != 0 {
		X_downAp7r(tls, fijk+4)
		*(*int32)(unsafe.Pointer(res)) += int32(1)
	}
	// The center point is now in the same substrate grid as the origin
	// cell vertices. Add the center point substate coordinates
	// to each vertex to translate the vertices to that cell.
	v = 0
	for {
		if !(v < int32(DNUM_HEX_VERTS)) {
			break
		}
		(*(*TFaceIJK)(unsafe.Pointer(fijkVerts + uintptr(v)*16))).Fface = (*TFaceIJK)(unsafe.Pointer(fijk)).Fface
		X_ijkAdd(tls, fijk+4, verts+uintptr(v)*12, fijkVerts+uintptr(v)*16+4)
		X_ijkNormalize(tls, fijkVerts+uintptr(v)*16+4)
		goto _1
	_1:
		;
		v++
	}
}

// C documentation
//
//	/**
//	 * Adjusts a FaceIJK address in place so that the resulting cell address is
//	 * relative to the correct icosahedral face.
//	 *
//	 * @param fijk The FaceIJK address of the cell.
//	 * @param res The H3 resolution of the cell.
//	 * @param pentLeading4 Whether or not the cell is a pentagon with a leading
//	 *        digit 4.
//	 * @param substrate Whether or not the cell is in a substrate grid.
//	 * @return 0 if on original face (no overage); 1 if on face edge (only occurs
//	 *         on substrate grids); 2 if overage on new face interior
//	 */
func X_adjustOverageClassII(tls *libc.TLS, fijk uintptr, res int32, pentLeading4 int32, substrate int32) (r TOverage) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var fijkOrient, ijk uintptr
	var i, maxDim, unitScale int32
	var overage TOverage
	var _ /* origin at bp+0 */ TCoordIJK
	var _ /* tmp at bp+12 */ TCoordIJK
	var _ /* transVec at bp+24 */ TCoordIJK
	_, _, _, _, _, _ = fijkOrient, i, ijk, maxDim, overage, unitScale
	overage = int32(_NO_OVERAGE)
	ijk = fijk + 4
	// get the maximum dimension value; scale if a substrate grid
	maxDim = _maxDimByCIIres[res]
	if substrate != 0 {
		maxDim *= int32(3)
	}
	// check for overage
	if substrate != 0 && (*TCoordIJK)(unsafe.Pointer(ijk)).Fi+(*TCoordIJK)(unsafe.Pointer(ijk)).Fj+(*TCoordIJK)(unsafe.Pointer(ijk)).Fk == maxDim { // on edge
		overage = int32(_FACE_EDGE)
	} else {
		if (*TCoordIJK)(unsafe.Pointer(ijk)).Fi+(*TCoordIJK)(unsafe.Pointer(ijk)).Fj+(*TCoordIJK)(unsafe.Pointer(ijk)).Fk > maxDim { // overage
			overage = int32(_NEW_FACE)
			if (*TCoordIJK)(unsafe.Pointer(ijk)).Fk > 0 {
				if (*TCoordIJK)(unsafe.Pointer(ijk)).Fj > 0 { // jk "quadrant"
					fijkOrient = uintptr(unsafe.Pointer(&_faceNeighbors)) + uintptr((*TFaceIJK)(unsafe.Pointer(fijk)).Fface)*80 + 3*20
				} else { // ik "quadrant"
					fijkOrient = uintptr(unsafe.Pointer(&_faceNeighbors)) + uintptr((*TFaceIJK)(unsafe.Pointer(fijk)).Fface)*80 + 2*20
					// adjust for the pentagonal missing sequence
					if pentLeading4 != 0 {
						X_setIJK(tls, bp, maxDim, 0, 0)
						X_ijkSub(tls, ijk, bp, bp+12)
						// rotate to adjust for the missing sequence
						X_ijkRotate60cw(tls, bp+12)
						// translate the origin back to the center of the triangle
						X_ijkAdd(tls, bp+12, bp, ijk)
					}
				}
			} else { // ij "quadrant"
				fijkOrient = uintptr(unsafe.Pointer(&_faceNeighbors)) + uintptr((*TFaceIJK)(unsafe.Pointer(fijk)).Fface)*80 + 1*20
			}
			(*TFaceIJK)(unsafe.Pointer(fijk)).Fface = (*TFaceOrientIJK)(unsafe.Pointer(fijkOrient)).Fface
			// rotate and translate for adjacent face
			i = 0
			for {
				if !(i < (*TFaceOrientIJK)(unsafe.Pointer(fijkOrient)).FccwRot60) {
					break
				}
				X_ijkRotate60ccw(tls, ijk)
				goto _1
			_1:
				;
				i++
			}
			*(*TCoordIJK)(unsafe.Pointer(bp + 24)) = (*TFaceOrientIJK)(unsafe.Pointer(fijkOrient)).Ftranslate
			unitScale = _unitScaleByCIIres[res]
			if substrate != 0 {
				unitScale *= int32(3)
			}
			X_ijkScale(tls, bp+24, unitScale)
			X_ijkAdd(tls, ijk, bp+24, ijk)
			X_ijkNormalize(tls, ijk)
			// overage points on pentagon boundaries can end up on edges
			if substrate != 0 && (*TCoordIJK)(unsafe.Pointer(ijk)).Fi+(*TCoordIJK)(unsafe.Pointer(ijk)).Fj+(*TCoordIJK)(unsafe.Pointer(ijk)).Fk == maxDim { // on edge
				overage = int32(_FACE_EDGE)
			}
		}
	}
	return overage
}

// C documentation
//
//	/**
//	 * Adjusts a FaceIJK address for a pentagon vertex in a substrate grid in
//	 * place so that the resulting cell address is relative to the correct
//	 * icosahedral face.
//	 *
//	 * @param fijk The FaceIJK address of the cell.
//	 * @param res The H3 resolution of the cell.
//	 */
func X_adjustPentVertOverage(tls *libc.TLS, fijk uintptr, res int32) (r TOverage) {
	var overage TOverage
	var pentLeading4 int32
	_, _ = overage, pentLeading4
	pentLeading4 = 0
	for cond := true; cond; cond = overage == int32(_NEW_FACE) {
		overage = X_adjustOverageClassII(tls, fijk, res, pentLeading4, int32(1))
	}
	return overage
}

// C documentation
//
//	/**
//	 * Encodes a coordinate on the sphere to the corresponding icosahedral face and
//	 * containing the squared euclidean distance to that face center.
//	 *
//	 * @param g The spherical coordinates to encode.
//	 * @param face The icosahedral face containing the spherical coordinates.
//	 * @param sqd The squared euclidean distance to its icosahedral face center.
//	 */
func X_geoToClosestFace(tls *libc.TLS, g uintptr, face uintptr, sqd uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var f int32
	var sqdT float64
	var _ /* v3d at bp+0 */ TVec3d
	_, _ = f, sqdT
	X_geoToVec3d(tls, g, bp)
	// determine the icosahedron face
	*(*int32)(unsafe.Pointer(face)) = 0
	// The distance between two farthest points is 2.0, therefore the square of
	// the distance between two points should always be less or equal than 4.0 .
	*(*float64)(unsafe.Pointer(sqd)) = float64(5)
	f = 0
	for {
		if !(f < int32(DNUM_ICOSA_FACES)) {
			break
		}
		sqdT = X_pointSquareDist(tls, uintptr(unsafe.Pointer(&_faceCenterPoint))+uintptr(f)*24, bp)
		if sqdT < *(*float64)(unsafe.Pointer(sqd)) {
			*(*int32)(unsafe.Pointer(face)) = f
			*(*float64)(unsafe.Pointer(sqd)) = sqdT
		}
		goto _1
	_1:
		;
		f++
	}
}

const DEPSILON2 = 0.0000000000000001
const DINT32_MAX3 = 2147483647
const DINV_RES0_U_GNOMONIC2 = 2.61803398874989588842
const DM_AP7_ROT_RADS2 = 0.333473172251832115336090755351601070065900389
const DM_ONETHIRD2 = 0.333333333333333333333333333333333333333
const DM_SQRT3_24 = 0.8660254037844386467637231707529361834714
const DRES0_U_GNOMONIC2 = 0.38196601125010500003

var _UNIT_VECS6 = [7]TCoordIJK{
	0: {},
	1: {
		Fk: int32(1),
	},
	2: {
		Fj: int32(1),
	},
	3: {
		Fj: int32(1),
		Fk: int32(1),
	},
	4: {
		Fi: int32(1),
	},
	5: {
		Fi: int32(1),
		Fk: int32(1),
	},
	6: {
		Fi: int32(1),
		Fj: int32(1),
	},
}

type TIterCellsChildren = struct {
	Fh          TH3Index
	F_parentRes int32
	F_skipDigit int32
}

type TIterCellsResolution = struct {
	Fh            TH3Index
	F_baseCellNum int32
	F_res         int32
	F_itC         TIterCellsChildren
}

// C documentation
//
//	/** @var H3ErrorDescriptions
//	 *  @brief An array of strings describing each of the H3ErrorCodes enum values
//	 */
var _H3ErrorDescriptions = [16]uintptr{
	0:  __ccgo_ts + 206,
	1:  __ccgo_ts + 214,
	2:  __ccgo_ts + 278,
	3:  __ccgo_ts + 319,
	4:  __ccgo_ts + 384,
	5:  __ccgo_ts + 436,
	6:  __ccgo_ts + 464,
	7:  __ccgo_ts + 501,
	8:  __ccgo_ts + 540,
	9:  __ccgo_ts + 570,
	10: __ccgo_ts + 606,
	11: __ccgo_ts + 622,
	12: __ccgo_ts + 656,
	13: __ccgo_ts + 700,
	14: __ccgo_ts + 725,
	15: __ccgo_ts + 769,
}

// C documentation
//
//	/**
//	 * Returns the string describing the H3Error. This string is internally
//	 * allocated and should not be `free`d.
//	 * @param err The H3 error.
//	 * @return The string describing the H3Error
//	 */
func XdescribeH3Error(tls *libc.TLS, err TH3Error) (r uintptr) {
	if err >= uint32(0) && err <= uint32(15) { // TODO: Better way to bounds check here?
		return _H3ErrorDescriptions[err]
	} else {
		return __ccgo_ts + 806
	}
	return r
}

// C documentation
//
//	/**
//	 * Returns the H3 resolution of an H3 index.
//	 * @param h The H3 index.
//	 * @return The resolution of the H3 index argument.
//	 */
func XgetResolution(tls *libc.TLS, h TH3Index) (r int32) {
	return libc.Int32FromUint64(h & (libc.Uint64FromUint64(15) << libc.Int32FromInt32(DH3_RES_OFFSET)) >> libc.Int32FromInt32(DH3_RES_OFFSET))
}

// C documentation
//
//	/**
//	 * Returns the H3 base cell "number" of an H3 cell (hexagon or pentagon).
//	 *
//	 * Note: Technically works on H3 edges, but will return base cell of the
//	 * origin cell.
//	 *
//	 * @param h The H3 cell.
//	 * @return The base cell "number" of the H3 cell argument.
//	 */
func XgetBaseCellNumber(tls *libc.TLS, h TH3Index) (r int32) {
	return libc.Int32FromUint64(h & (libc.Uint64FromInt32(libc.Int32FromInt32(127)) << libc.Int32FromInt32(DH3_BC_OFFSET)) >> libc.Int32FromInt32(DH3_BC_OFFSET))
}

// C documentation
//
//	/**
//	 * Converts a string representation of an H3 index into an H3 index.
//	 * @param str The string representation of an H3 index.
//	 * @param out Output: The H3 index corresponding to the string argument
//	 */
func XstringToH3(tls *libc.TLS, str uintptr, out uintptr) (r TH3Error) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var read int32
	var _ /* h at bp+0 */ TH3Index
	_ = read
	*(*TH3Index)(unsafe.Pointer(bp)) = uint64(DH3_NULL)
	// If failed, h will be unmodified and we should return H3_NULL anyways.
	read = libc.Xsscanf(tls, str, __ccgo_ts+825, libc.VaList(bp+16, bp))
	if read != int32(1) {
		return uint32(_E_FAILED)
	}
	*(*TH3Index)(unsafe.Pointer(out)) = *(*TH3Index)(unsafe.Pointer(bp))
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * Converts an H3 index into a string representation.
//	 * @param h The H3 index to convert.
//	 * @param str The string representation of the H3 index.
//	 * @param sz Size of the buffer `str`
//	 */
func Xh3ToString(tls *libc.TLS, h TH3Index, str uintptr, sz Tsize_t) (r TH3Error) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	// An unsigned 64 bit integer will be expressed in at most
	// 16 digits plus 1 for the null terminator.
	if sz < uint64(17) {
		// Buffer is potentially not large enough.
		return uint32(_E_MEMORY_BOUNDS)
	}
	libc.Xsprintf(tls, str, __ccgo_ts+825, libc.VaList(bp+8, h))
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/*
//	The top 8 bits of any cell should be a specific constant:
//
//	- The 1 high bit should be `0`
//	- The 4 mode bits should be `0001` (H3_CELL_MODE)
//	- The 3 reserved bits should be `000`
//
//	In total, the top 8 bits should be `0_0001_000`
//	*/
func __hasGoodTopBits(tls *libc.TLS, h TH3Index) (r uint8) {
	h >>= libc.Uint64FromInt32(libc.Int32FromInt32(64) - libc.Int32FromInt32(8))
	return libc.BoolUint8(h == uint64(0b00001000))
}

// C documentation
//
//	/* Check that no digit from 1 to `res` is 7 (INVALID_DIGIT).
//
//	MHI = 0b100100100100100100100100100100100100100100100;
//	MLO = MHI >> 2;
//
//	|  d  | d & MHI |  ~d | ~d - MLO | d & MHI & (~d - MLO) |  result |
//	|-----|---------|-----|----------|----------------------|---------|
//	| 000 |     000 |     |          |                  000 | OK      |
//	| 001 |     000 |     |          |                  000 | OK      |
//	| 010 |     000 |     |          |                  000 | OK      |
//	| 011 |     000 |     |          |                  000 | OK      |
//	| 100 |     100 | 011 | 010      |                  000 | OK      |
//	| 101 |     100 | 010 | 001      |                  000 | OK      |
//	| 110 |     100 | 001 | 000      |                  000 | OK      |
//	| 111 |     100 | 000 | 111*     |                  100 | invalid |
//
//	  *: carry happened
//
//
//	Note: only care about identifying the *lowest* 7.
//
//	Examples with multiple digits:
//
//	|    d    | d & MHI |    ~d   | ~d - MLO | d & MHI & (~d - MLO) |  result |
//	|---------|---------|---------|----------|----------------------|---------|
//	| 111.111 | 100.100 | 000.000 | 110.111* |              100.100 | invalid |
//	| 110.111 | 100.100 | 001.000 | 111.111* |              100.100 | invalid |
//	| 110.110 | 100.100 | 001.001 | 000.000  |              000.000 | OK      |
//
//	  *: carry happened
//
//	In the second example with 110.111, we "misidentify" the 110 as a 7, due
//	to a carry from the lower bits. But this is OK because we correctly
//	identify the lowest (only, in this example) 7 just before it.
//
//	We only have to worry about carries affecting higher bits in the case of
//	a 7; all other digits (0--6) don't cause a carry when computing ~d - MLO.
//	So even though a 7 can affect the results of higher bits, this is OK
//	because we will always correctly identify the lowest 7.
//
//	For further notes, see the discussion here:
//	https://github.com/uber/h3/pull/496#discussion_r795851046
//	*/
func __hasAny7UptoRes(tls *libc.TLS, h TH3Index, res int32) (r uint8) {
	var MHI, MLO Tuint64_t
	var shift int32
	_, _, _ = MHI, MLO, shift
	MHI = uint64(0b100100100100100100100100100100100100100100100)
	MLO = MHI >> int32(2)
	shift = int32(3) * (int32(15) - res)
	h >>= libc.Uint64FromInt32(shift)
	h <<= libc.Uint64FromInt32(shift)
	h = h & MHI & (^h - MLO)
	return libc.BoolUint8(h != uint64(0))
}

// C documentation
//
//	/* Check that all unused digits after `res` are set to 7 (INVALID_DIGIT).
//
//	Bit shift to avoid looping through digits.
//	*/
func __hasAll7AfterRes(tls *libc.TLS, h TH3Index, res int32) (r uint8) {
	var shift int32
	_ = shift
	// NOTE: res check is needed because we can't shift by 64
	if res < int32(15) {
		shift = int32(19) + int32(3)*res
		h = ^h
		h <<= libc.Uint64FromInt32(shift)
		h >>= libc.Uint64FromInt32(shift)
		return libc.BoolUint8(h == uint64(0))
	}
	return uint8(Dtrue)
}

// C documentation
//
//	/*
//	Get index of first nonzero bit of an H3Index.
//
//	When available, use compiler intrinsics, which should be fast.
//	If not available, fall back to a loop.
//	*/
func __firstOneIndex(tls *libc.TLS, h TH3Index) (r int32) {
	return int32(63) - libc.X__builtin_clzll(tls, h)
}

// C documentation
//
//	/*
//	One final validation just for cells whose base cell (res 0)
//	is a pentagon.
//
//	Pentagon cells start with a sequence of 0's (CENTER_DIGIT's).
//	The first nonzero digit can't be a 1 (i.e., "deleted subsequence",
//	PENTAGON_SKIPPED_DIGIT, or K_AXES_DIGIT).
//
//	We can check that (in the lower 45 = 15*3 bits) the position of the
//	first 1 bit isn't divisible by 3.
//	*/
func __hasDeletedSubsequence(tls *libc.TLS, h TH3Index, base_cell int32) (r uint8) {
	if _isBaseCellPentagonArr[base_cell] != 0 {
		h <<= uint64(19)
		h >>= uint64(19)
		if h == uint64(0) {
			return uint8(Dfalse)
		} // all zeros: res 15 pentagon
		return libc.BoolUint8(__firstOneIndex(tls, h)%int32(3) == 0)
	}
	return uint8(Dfalse)
}

// TODO: https://github.com/uber/h3/issues/984
var _isBaseCellPentagonArr = [128]uint8{
	4:   uint8(1),
	14:  uint8(1),
	24:  uint8(1),
	38:  uint8(1),
	49:  uint8(1),
	58:  uint8(1),
	63:  uint8(1),
	72:  uint8(1),
	83:  uint8(1),
	97:  uint8(1),
	107: uint8(1),
	117: uint8(1),
}

// C documentation
//
//	/**
//	 * Returns whether or not an H3 index is a valid cell (hexagon or pentagon).
//	 * @param h The H3 index to validate.
//	 * @return 1 if the H3 index if valid, and 0 if it is not.
//	 */
func XisValidCell(tls *libc.TLS, h TH3Index) (r int32) {
	var bc, res int32
	_, _ = bc, res
	/*
	   Look for bit patterns that would disqualify an H3Index from
	   being valid. If identified, exit early.
	   For reference the H3 index bit layout:
	   |   Region   | # bits |
	   |------------|--------|
	   | High       |      1 |
	   | Mode       |      4 |
	   | Reserved   |      3 |
	   | Resolution |      4 |
	   | Base Cell  |      7 |
	   | Digit 1    |      3 |
	   | Digit 2    |      3 |
	   | ...        |    ... |
	   | Digit 15   |      3 |
	   Speed benefits come from using bit manipulation instead of loops,
	   whenever possible.
	*/
	if !(__hasGoodTopBits(tls, h) != 0) {
		return Dfalse
	}
	// No need to check resolution; any 4 bits give a valid resolution.
	res = libc.Int32FromUint64(h & (libc.Uint64FromUint64(15) << libc.Int32FromInt32(DH3_RES_OFFSET)) >> libc.Int32FromInt32(DH3_RES_OFFSET))
	// Get base cell number and check that it is valid.
	bc = libc.Int32FromUint64(h & (libc.Uint64FromInt32(libc.Int32FromInt32(127)) << libc.Int32FromInt32(DH3_BC_OFFSET)) >> libc.Int32FromInt32(DH3_BC_OFFSET))
	if bc >= int32(DNUM_BASE_CELLS) {
		return Dfalse
	}
	if __hasAny7UptoRes(tls, h, res) != 0 {
		return Dfalse
	}
	if !(__hasAll7AfterRes(tls, h, res) != 0) {
		return Dfalse
	}
	if __hasDeletedSubsequence(tls, h, bc) != 0 {
		return Dfalse
	}
	// If no disqualifications were identified, the index is a valid H3 cell.
	return int32(Dtrue)
}

// C documentation
//
//	/**
//	 * Initializes an H3 index.
//	 * @param hp The H3 index to initialize.
//	 * @param res The H3 resolution to initialize the index to.
//	 * @param baseCell The H3 base cell to initialize the index to.
//	 * @param initDigit The H3 digit (0-7) to initialize all of the index digits to.
//	 */
func XsetH3Index(tls *libc.TLS, hp uintptr, res int32, baseCell int32, initDigit TDirection) {
	var h TH3Index
	var r int32
	_, _ = h, r
	h = libc.Uint64FromUint64(35184372088831)
	h = h & ^(libc.Uint64FromInt32(libc.Int32FromInt32(15))<<libc.Int32FromInt32(DH3_MODE_OFFSET)) | libc.Uint64FromInt32(libc.Int32FromInt32(DH3_CELL_MODE))<<libc.Int32FromInt32(DH3_MODE_OFFSET)
	h = h & ^(libc.Uint64FromUint64(15)<<libc.Int32FromInt32(DH3_RES_OFFSET)) | libc.Uint64FromInt32(res)<<libc.Int32FromInt32(DH3_RES_OFFSET)
	h = h & ^(libc.Uint64FromInt32(libc.Int32FromInt32(127))<<libc.Int32FromInt32(DH3_BC_OFFSET)) | libc.Uint64FromInt32(baseCell)<<libc.Int32FromInt32(DH3_BC_OFFSET)
	r = int32(1)
	for {
		if !(r <= res) {
			break
		}
		h = h & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<((libc.Int32FromInt32(DMAX_H3_RES)-r)*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))) | libc.Uint64FromInt32(initDigit)<<((libc.Int32FromInt32(DMAX_H3_RES)-r)*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))
		goto _1
	_1:
		;
		r++
	}
	*(*TH3Index)(unsafe.Pointer(hp)) = h
}

// C documentation
//
//	/**
//	 * cellToParent produces the parent index for a given H3 index
//	 *
//	 * @param h H3Index to find parent of
//	 * @param parentRes The resolution to switch to (parent, grandparent, etc)
//	 * @param out Output: H3Index of the parent
//	 */
func XcellToParent(tls *libc.TLS, h TH3Index, parentRes int32, out uintptr) (r TH3Error) {
	var childRes, i int32
	var parentH, v1 TH3Index
	_, _, _, _ = childRes, i, parentH, v1
	childRes = libc.Int32FromUint64(h & (libc.Uint64FromUint64(15) << libc.Int32FromInt32(DH3_RES_OFFSET)) >> libc.Int32FromInt32(DH3_RES_OFFSET))
	if parentRes < 0 || parentRes > int32(DMAX_H3_RES) {
		return uint32(_E_RES_DOMAIN)
	} else {
		if parentRes > childRes {
			return uint32(_E_RES_MISMATCH)
		} else {
			if parentRes == childRes {
				*(*TH3Index)(unsafe.Pointer(out)) = h
				return uint32(_E_SUCCESS)
			}
		}
	}
	v1 = h & ^(libc.Uint64FromUint64(15)<<libc.Int32FromInt32(DH3_RES_OFFSET)) | libc.Uint64FromInt32(parentRes)<<libc.Int32FromInt32(DH3_RES_OFFSET)
	h = v1
	parentH = v1
	i = parentRes + int32(1)
	for {
		if !(i <= childRes) {
			break
		}
		parentH = parentH & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<((libc.Int32FromInt32(DMAX_H3_RES)-i)*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))) | libc.Uint64FromInt32(libc.Int32FromInt32(7))<<((libc.Int32FromInt32(DMAX_H3_RES)-i)*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))
		goto _2
	_2:
		;
		i++
	}
	*(*TH3Index)(unsafe.Pointer(out)) = parentH
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * Determines whether one resolution is a valid child resolution for a cell.
//	 * Each resolution is considered a valid child resolution of itself.
//	 *
//	 * @param h         h3Index  parent cell
//	 * @param childRes  int      resolution of the child
//	 *
//	 * @return The validity of the child resolution
//	 */
func __hasChildAtRes(tls *libc.TLS, h TH3Index, childRes int32) (r uint8) {
	var parentRes int32
	_ = parentRes
	parentRes = libc.Int32FromUint64(h & (libc.Uint64FromUint64(15) << libc.Int32FromInt32(DH3_RES_OFFSET)) >> libc.Int32FromInt32(DH3_RES_OFFSET))
	if childRes < parentRes || childRes > int32(DMAX_H3_RES) {
		return uint8(Dfalse)
	}
	return uint8(Dtrue)
}

// C documentation
//
//	/**
//	 * cellToChildrenSize returns the exact number of children for a cell at a
//	 * given child resolution.
//	 *
//	 * @param h         H3Index to find the number of children of
//	 * @param childRes  The child resolution you're interested in
//	 * @param out      Output: exact number of children (handles hexagons and
//	 * pentagons correctly)
//	 */
func XcellToChildrenSize(tls *libc.TLS, h TH3Index, childRes int32, out uintptr) (r TH3Error) {
	var n int32
	_ = n
	if !(__hasChildAtRes(tls, h, childRes) != 0) {
		return uint32(_E_RES_DOMAIN)
	}
	n = childRes - libc.Int32FromUint64(h&(libc.Uint64FromUint64(15)<<libc.Int32FromInt32(DH3_RES_OFFSET))>>libc.Int32FromInt32(DH3_RES_OFFSET))
	if XisPentagon(tls, h) != 0 {
		*(*Tint64_t)(unsafe.Pointer(out)) = int64(1) + int64(5)*(X_ipow(tls, int64(7), int64(n))-int64(1))/int64(6)
	} else {
		*(*Tint64_t)(unsafe.Pointer(out)) = X_ipow(tls, int64(7), int64(n))
	}
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * makeDirectChild takes an index and immediately returns the immediate child
//	 * index based on the specified cell number. Bit operations only, could generate
//	 * invalid indexes if not careful (deleted cell under a pentagon).
//	 *
//	 * @param h H3Index to find the direct child of
//	 * @param cellNumber int id of the direct child (0-6)
//	 *
//	 * @return The new H3Index for the child
//	 */
func XmakeDirectChild(tls *libc.TLS, h TH3Index, cellNumber int32) (r TH3Index) {
	var childH, v1 TH3Index
	var childRes int32
	_, _, _ = childH, childRes, v1
	childRes = libc.Int32FromUint64(h&(libc.Uint64FromUint64(15)<<libc.Int32FromInt32(DH3_RES_OFFSET))>>libc.Int32FromInt32(DH3_RES_OFFSET)) + int32(1)
	v1 = h & ^(libc.Uint64FromUint64(15)<<libc.Int32FromInt32(DH3_RES_OFFSET)) | libc.Uint64FromInt32(childRes)<<libc.Int32FromInt32(DH3_RES_OFFSET)
	h = v1
	childH = v1
	childH = childH & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<((libc.Int32FromInt32(DMAX_H3_RES)-childRes)*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))) | libc.Uint64FromInt32(cellNumber)<<((libc.Int32FromInt32(DMAX_H3_RES)-childRes)*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))
	return childH
}

// C documentation
//
//	/**
//	 * cellToChildren takes the given hexagon id and generates all of the children
//	 * at the specified resolution storing them into the provided memory pointer.
//	 * It's assumed that cellToChildrenSize was used to determine the allocation.
//	 *
//	 * @param h H3Index to find the children of
//	 * @param childRes int the child level to produce
//	 * @param children H3Index* the memory to store the resulting addresses in
//	 */
func XcellToChildren(tls *libc.TLS, h TH3Index, childRes int32, children uintptr) (r TH3Error) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i Tint64_t
	var _ /* iter at bp+0 */ TIterCellsChildren
	_ = i
	i = 0
	*(*TIterCellsChildren)(unsafe.Pointer(bp)) = XiterInitParent(tls, h, childRes)
	for {
		if !((*(*TIterCellsChildren)(unsafe.Pointer(bp))).Fh != 0) {
			break
		}
		*(*TH3Index)(unsafe.Pointer(children + uintptr(i)*8)) = (*(*TIterCellsChildren)(unsafe.Pointer(bp))).Fh
		i++
		goto _1
	_1:
		;
		XiterStepChild(tls, bp)
	}
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * Zero out index digits from start to end, inclusive.
//	 * No-op if start > end.
//	 */
func X_zeroIndexDigits(tls *libc.TLS, h TH3Index, start int32, end int32) (r TH3Index) {
	var m TH3Index
	_ = m
	if start > end {
		return h
	}
	m = uint64(0)
	m = ^m
	m <<= libc.Uint64FromInt32(int32(DH3_PER_DIGIT_OFFSET) * (end - start + int32(1)))
	m = ^m
	m <<= libc.Uint64FromInt32(int32(DH3_PER_DIGIT_OFFSET) * (int32(DMAX_H3_RES) - end))
	m = ^m
	return h & m
}

// C documentation
//
//	/**
//	 * cellToCenterChild produces the center child index for a given H3 index at
//	 * the specified resolution
//	 *
//	 * @param h H3Index to find center child of
//	 * @param childRes The resolution to switch to
//	 * @param child H3Index of the center child
//	 * @return 0 (E_SUCCESS) on success
//	 */
func XcellToCenterChild(tls *libc.TLS, h TH3Index, childRes int32, child uintptr) (r TH3Error) {
	if !(__hasChildAtRes(tls, h, childRes) != 0) {
		return uint32(_E_RES_DOMAIN)
	}
	h = X_zeroIndexDigits(tls, h, libc.Int32FromUint64(h&(libc.Uint64FromUint64(15)<<libc.Int32FromInt32(DH3_RES_OFFSET))>>libc.Int32FromInt32(DH3_RES_OFFSET))+int32(1), childRes)
	h = h & ^(libc.Uint64FromUint64(15)<<libc.Int32FromInt32(DH3_RES_OFFSET)) | libc.Uint64FromInt32(childRes)<<libc.Int32FromInt32(DH3_RES_OFFSET)
	*(*TH3Index)(unsafe.Pointer(child)) = h
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * compactCells takes a set of hexagons all at the same resolution and
//	 * compresses them by pruning full child branches to the parent level. This is
//	 * also done for all parents recursively to get the minimum number of hex
//	 * addresses that perfectly cover the defined space.
//	 * @param h3Set Set of hexagons
//	 * @param compactedSet The output array of compressed hexagons (preallocated)
//	 * @param numHexes The size of the input and output arrays (possible that no
//	 * contiguous regions exist in the set at all and no compression possible)
//	 * @return an error code on bad input data
//	 */
func XcompactCells(tls *libc.TLS, h3Set uintptr, compactedSet uintptr, numHexes Tint64_t) (r TH3Error) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var compactableCount, i, i1, i2, i3, loc, loc1, loopCount, loopCount1, maxCompactableCount, numRemainingHexes, uncompactableCount Tint64_t
	var compactableHexes, compactedSetOffset, hashSetArray, remainingHexes uintptr
	var count, count1, count2, limitCount, parentRes, res, v3, v7, v9 int32
	var currIndex, currIndex1, tempIndex, tempIndex1 TH3Index
	var isUncompactable uint8
	var parentError, parentError1 TH3Error
	var v10, v4, v8 bool
	var _ /* parent at bp+0 */ TH3Index
	var _ /* parent at bp+8 */ TH3Index
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = compactableCount, compactableHexes, compactedSetOffset, count, count1, count2, currIndex, currIndex1, hashSetArray, i, i1, i2, i3, isUncompactable, limitCount, loc, loc1, loopCount, loopCount1, maxCompactableCount, numRemainingHexes, parentError, parentError1, parentRes, remainingHexes, res, tempIndex, tempIndex1, uncompactableCount, v10, v3, v4, v7, v8, v9
	if numHexes == 0 {
		return uint32(_E_SUCCESS)
	}
	res = libc.Int32FromUint64(*(*TH3Index)(unsafe.Pointer(h3Set)) & (libc.Uint64FromUint64(15) << libc.Int32FromInt32(DH3_RES_OFFSET)) >> libc.Int32FromInt32(DH3_RES_OFFSET))
	if res == 0 {
		// No compaction possible, just copy the set to output
		i = 0
		for {
			if !(i < numHexes) {
				break
			}
			*(*TH3Index)(unsafe.Pointer(compactedSet + uintptr(i)*8)) = *(*TH3Index)(unsafe.Pointer(h3Set + uintptr(i)*8))
			goto _1
		_1:
			;
			i++
		}
		return uint32(_E_SUCCESS)
	}
	remainingHexes = libc.Xmalloc(tls, libc.Uint64FromInt64(numHexes)*uint64(8))
	if !(remainingHexes != 0) {
		return uint32(_E_MEMORY_ALLOC)
	}
	libc.Xmemcpy(tls, remainingHexes, h3Set, libc.Uint64FromInt64(numHexes)*uint64(8))
	hashSetArray = libc.Xcalloc(tls, libc.Uint64FromInt64(numHexes), uint64(8))
	if !(hashSetArray != 0) {
		libc.Xfree(tls, remainingHexes)
		return uint32(_E_MEMORY_ALLOC)
	}
	compactedSetOffset = compactedSet
	numRemainingHexes = numHexes
	for numRemainingHexes != 0 {
		res = libc.Int32FromUint64(*(*TH3Index)(unsafe.Pointer(remainingHexes)) & (libc.Uint64FromUint64(15) << libc.Int32FromInt32(DH3_RES_OFFSET)) >> libc.Int32FromInt32(DH3_RES_OFFSET))
		parentRes = res - int32(1)
		// If parentRes is less than zero, we've compacted all the way up to the
		// base cells. Time to process the remaining cells.
		if parentRes >= 0 {
			// Put the parents of the hexagons into the temp array
			// via a hashing mechanism, and use the reserved bits
			// to track how many times a parent is duplicated
			i1 = 0
			for {
				if !(i1 < numRemainingHexes) {
					break
				}
				currIndex = *(*TH3Index)(unsafe.Pointer(remainingHexes + uintptr(i1)*8))
				// TODO: This case is coverable (reachable by fuzzer)
				if currIndex != uint64(0) {
					// If the reserved bits were set by the caller, the
					// algorithm below may encounter undefined behavior
					// because it expects to have set the reserved bits
					// itself.
					if libc.Int32FromUint64(currIndex&(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<libc.Int32FromInt32(DH3_RESERVED_OFFSET))>>libc.Int32FromInt32(DH3_RESERVED_OFFSET)) != 0 {
						libc.Xfree(tls, remainingHexes)
						libc.Xfree(tls, hashSetArray)
						return uint32(_E_CELL_INVALID)
					}
					parentError = XcellToParent(tls, currIndex, parentRes, bp)
					// Should never be reachable as a result of the compact
					// algorithm. Can happen if cellToParent errors e.g.
					// because of incompatible resolutions.
					if parentError != 0 {
						libc.Xfree(tls, remainingHexes)
						libc.Xfree(tls, hashSetArray)
						return parentError
					}
					// Modulus hash the parent into the temp array
					loc = libc.Int64FromUint64(*(*TH3Index)(unsafe.Pointer(bp)) % libc.Uint64FromInt64(numRemainingHexes))
					loopCount = 0
					for *(*TH3Index)(unsafe.Pointer(hashSetArray + uintptr(loc)*8)) != uint64(0) {
						if loopCount > numRemainingHexes {
							if v4 = libc.Bool(0 != 0); !v4 {
								libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+829, int32(529), uintptr(unsafe.Pointer(&___func__5)))
							}
							_ = v4 || libc.Bool(libc.Int32FromInt32(0) != 0)
							v3 = libc.Int32FromInt32(1)
						} else {
							v3 = 0
						}
						if v3 != 0 {
							// This case should not be possible because at
							// most one index is placed into hashSetArray
							// per numRemainingHexes.
							libc.Xfree(tls, remainingHexes)
							libc.Xfree(tls, hashSetArray)
							return uint32(_E_FAILED)
						}
						tempIndex = *(*TH3Index)(unsafe.Pointer(hashSetArray + uintptr(loc)*8)) & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7)) << libc.Int32FromInt32(DH3_RESERVED_OFFSET))
						if tempIndex == *(*TH3Index)(unsafe.Pointer(bp)) {
							count = libc.Int32FromUint64(*(*TH3Index)(unsafe.Pointer(hashSetArray + uintptr(loc)*8))&(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<libc.Int32FromInt32(DH3_RESERVED_OFFSET))>>libc.Int32FromInt32(DH3_RESERVED_OFFSET)) + int32(1)
							limitCount = int32(7)
							if XisPentagon(tls, tempIndex & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<libc.Int32FromInt32(DH3_RESERVED_OFFSET))) != 0 {
								limitCount--
							}
							// One is added to count for this check to match
							// one being added to count later in this
							// function when checking for all children being
							// present.
							if count+int32(1) > limitCount {
								// Only possible on duplicate input
								libc.Xfree(tls, remainingHexes)
								libc.Xfree(tls, hashSetArray)
								return uint32(_E_DUPLICATE_INPUT)
							}
							*(*TH3Index)(unsafe.Pointer(bp)) = *(*TH3Index)(unsafe.Pointer(bp)) & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<libc.Int32FromInt32(DH3_RESERVED_OFFSET)) | libc.Uint64FromInt32(count)<<libc.Int32FromInt32(DH3_RESERVED_OFFSET)
							*(*TH3Index)(unsafe.Pointer(hashSetArray + uintptr(loc)*8)) = uint64(DH3_NULL)
						} else {
							loc = (loc + int64(1)) % numRemainingHexes
						}
						loopCount++
					}
					*(*TH3Index)(unsafe.Pointer(hashSetArray + uintptr(loc)*8)) = *(*TH3Index)(unsafe.Pointer(bp))
				}
				goto _2
			_2:
				;
				i1++
			}
		}
		// Determine which parent hexagons have a complete set
		// of children and put them in the compactableHexes array
		compactableCount = 0
		maxCompactableCount = numRemainingHexes / int64(6) // Somehow all pentagons; conservative
		if maxCompactableCount == 0 {
			libc.Xmemcpy(tls, compactedSetOffset, remainingHexes, libc.Uint64FromInt64(numRemainingHexes)*uint64(8))
			break
		}
		compactableHexes = libc.Xcalloc(tls, libc.Uint64FromInt64(maxCompactableCount), uint64(8))
		if !(compactableHexes != 0) {
			libc.Xfree(tls, remainingHexes)
			libc.Xfree(tls, hashSetArray)
			return uint32(_E_MEMORY_ALLOC)
		}
		i2 = 0
		for {
			if !(i2 < numRemainingHexes) {
				break
			}
			if *(*TH3Index)(unsafe.Pointer(hashSetArray + uintptr(i2)*8)) == uint64(0) {
				goto _5
			}
			count1 = libc.Int32FromUint64(*(*TH3Index)(unsafe.Pointer(hashSetArray + uintptr(i2)*8))&(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<libc.Int32FromInt32(DH3_RESERVED_OFFSET))>>libc.Int32FromInt32(DH3_RESERVED_OFFSET)) + int32(1)
			// Include the deleted direction for pentagons as implicitly "there"
			if XisPentagon(tls, *(*TH3Index)(unsafe.Pointer(hashSetArray + uintptr(i2)*8)) & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<libc.Int32FromInt32(DH3_RESERVED_OFFSET))) != 0 {
				// We need this later on, no need to recalculate
				*(*TH3Index)(unsafe.Pointer(hashSetArray + uintptr(i2)*8)) = *(*TH3Index)(unsafe.Pointer(hashSetArray + uintptr(i2)*8)) & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<libc.Int32FromInt32(DH3_RESERVED_OFFSET)) | libc.Uint64FromInt32(count1)<<libc.Int32FromInt32(DH3_RESERVED_OFFSET)
				// Increment count after setting the reserved bits,
				// since count is already incremented above, so it
				// will be the expected value for a complete hexagon.
				count1++
			}
			if count1 == int32(7) {
				// Bingo! Full set!
				*(*TH3Index)(unsafe.Pointer(compactableHexes + uintptr(compactableCount)*8)) = *(*TH3Index)(unsafe.Pointer(hashSetArray + uintptr(i2)*8)) & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7)) << libc.Int32FromInt32(DH3_RESERVED_OFFSET))
				compactableCount++
			}
			goto _5
		_5:
			;
			i2++
		}
		// Uncompactable hexes are immediately copied into the
		// output compactedSetOffset
		uncompactableCount = 0
		i3 = 0
		for {
			if !(i3 < numRemainingHexes) {
				break
			}
			currIndex1 = *(*TH3Index)(unsafe.Pointer(remainingHexes + uintptr(i3)*8))
			// TODO: This case is coverable (reachable by fuzzer)
			if currIndex1 != uint64(DH3_NULL) {
				isUncompactable = uint8(Dtrue)
				// Resolution 0 cells always uncompactable, and trying to take
				// the res -1 parent of a cell is invalid.
				if parentRes >= 0 {
					parentError1 = XcellToParent(tls, currIndex1, parentRes, bp+8)
					if parentError1 != 0 {
						if v8 = libc.Bool(0 != 0); !v8 {
							libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+829, int32(620), uintptr(unsafe.Pointer(&___func__5)))
						}
						_ = v8 || libc.Bool(libc.Int32FromInt32(0) != 0)
						v7 = libc.Int32FromInt32(1)
					} else {
						v7 = 0
					}
					if v7 != 0 {
						libc.Xfree(tls, compactableHexes)
						libc.Xfree(tls, remainingHexes)
						libc.Xfree(tls, hashSetArray)
						return parentError1
					}
					// Modulus hash the parent into the temp array
					// to determine if this index was included in
					// the compactableHexes array
					loc1 = libc.Int64FromUint64(*(*TH3Index)(unsafe.Pointer(bp + 8)) % libc.Uint64FromInt64(numRemainingHexes))
					loopCount1 = 0
					for cond := true; cond; cond = *(*TH3Index)(unsafe.Pointer(hashSetArray + uintptr(loc1)*8)) != *(*TH3Index)(unsafe.Pointer(bp + 8)) {
						if loopCount1 > numRemainingHexes {
							if v10 = libc.Bool(0 != 0); !v10 {
								libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+829, int32(632), uintptr(unsafe.Pointer(&___func__5)))
							}
							_ = v10 || libc.Bool(libc.Int32FromInt32(0) != 0)
							v9 = libc.Int32FromInt32(1)
						} else {
							v9 = 0
						}
						if v9 != 0 {
							// This case should not be possible because at most
							// one index is placed into hashSetArray per input
							// hexagon.
							libc.Xfree(tls, compactableHexes)
							libc.Xfree(tls, remainingHexes)
							libc.Xfree(tls, hashSetArray)
							return uint32(_E_FAILED)
						}
						tempIndex1 = *(*TH3Index)(unsafe.Pointer(hashSetArray + uintptr(loc1)*8)) & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7)) << libc.Int32FromInt32(DH3_RESERVED_OFFSET))
						if tempIndex1 == *(*TH3Index)(unsafe.Pointer(bp + 8)) {
							count2 = libc.Int32FromUint64(*(*TH3Index)(unsafe.Pointer(hashSetArray + uintptr(loc1)*8))&(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<libc.Int32FromInt32(DH3_RESERVED_OFFSET))>>libc.Int32FromInt32(DH3_RESERVED_OFFSET)) + int32(1)
							if count2 == int32(7) {
								isUncompactable = uint8(Dfalse)
							}
							break
						} else {
							loc1 = (loc1 + int64(1)) % numRemainingHexes
						}
						loopCount1++
					}
				}
				if isUncompactable != 0 {
					*(*TH3Index)(unsafe.Pointer(compactedSetOffset + uintptr(uncompactableCount)*8)) = *(*TH3Index)(unsafe.Pointer(remainingHexes + uintptr(i3)*8))
					uncompactableCount++
				}
			}
			goto _6
		_6:
			;
			i3++
		}
		// Set up for the next loop
		libc.Xmemset(tls, hashSetArray, 0, libc.Uint64FromInt64(numHexes)*uint64(8))
		compactedSetOffset += uintptr(uncompactableCount) * 8
		libc.Xmemcpy(tls, remainingHexes, compactableHexes, libc.Uint64FromInt64(compactableCount)*uint64(8))
		numRemainingHexes = compactableCount
		libc.Xfree(tls, compactableHexes)
	}
	libc.Xfree(tls, remainingHexes)
	libc.Xfree(tls, hashSetArray)
	return uint32(_E_SUCCESS)
}

var ___func__5 = [13]uint8{'c', 'o', 'm', 'p', 'a', 'c', 't', 'C', 'e', 'l', 'l', 's'}

// C documentation
//
//	/**
//	 * uncompactCells takes a compressed set of cells and expands back to the
//	 * original set of cells.
//	 *
//	 * Skips elements that are H3_NULL (i.e., 0).
//	 *
//	 * @param   compactedSet  Set of compacted cells
//	 * @param   numCompacted  The number of cells in the input compacted set
//	 * @param   outSet      Output array for decompressed cells (preallocated)
//	 * @param   numOut      The size of the output array to bound check against
//	 * @param   res         The H3 resolution to decompress to
//	 * @return              An error code if output array is too small or any cell
//	 *                      is smaller than the output resolution.
//	 */
func XuncompactCells(tls *libc.TLS, compactedSet uintptr, numCompacted Tint64_t, outSet uintptr, numOut Tint64_t, res int32) (r TH3Error) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i, j Tint64_t
	var _ /* iter at bp+0 */ TIterCellsChildren
	_, _ = i, j
	i = 0
	j = 0
	for {
		if !(j < numCompacted) {
			break
		}
		if !(__hasChildAtRes(tls, *(*TH3Index)(unsafe.Pointer(compactedSet + uintptr(j)*8)), res) != 0) {
			return uint32(_E_RES_MISMATCH)
		}
		*(*TIterCellsChildren)(unsafe.Pointer(bp)) = XiterInitParent(tls, *(*TH3Index)(unsafe.Pointer(compactedSet + uintptr(j)*8)), res)
		for {
			if !((*(*TIterCellsChildren)(unsafe.Pointer(bp))).Fh != 0) {
				break
			}
			if i >= numOut {
				return uint32(_E_MEMORY_BOUNDS)
			} // went too far; abort!
			*(*TH3Index)(unsafe.Pointer(outSet + uintptr(i)*8)) = (*(*TIterCellsChildren)(unsafe.Pointer(bp))).Fh
			goto _2
		_2:
			;
			i++
			XiterStepChild(tls, bp)
		}
		goto _1
	_1:
		;
		j++
	}
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * uncompactCellsSize takes a compacted set of hexagons and provides
//	 * the exact size of the uncompacted set of hexagons.
//	 *
//	 * @param   compactedSet  Set of hexagons
//	 * @param   numCompacted  The number of hexes in the input set
//	 * @param   res           The hexagon resolution to decompress to
//	 * @param   out           The number of hexagons to allocate memory for
//	 * @returns E_SUCCESS on success, or another value on error
//	 */
func XuncompactCellsSize(tls *libc.TLS, compactedSet uintptr, numCompacted Tint64_t, res int32, out uintptr) (r TH3Error) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var childrenError TH3Error
	var i, numOut Tint64_t
	var _ /* childrenSize at bp+0 */ Tint64_t
	_, _, _ = childrenError, i, numOut
	numOut = 0
	i = 0
	for {
		if !(i < numCompacted) {
			break
		}
		if *(*TH3Index)(unsafe.Pointer(compactedSet + uintptr(i)*8)) == uint64(DH3_NULL) {
			goto _1
		}
		childrenError = XcellToChildrenSize(tls, *(*TH3Index)(unsafe.Pointer(compactedSet + uintptr(i)*8)), res, bp)
		if childrenError != 0 {
			// The parent res does not contain `res`.
			return uint32(_E_RES_MISMATCH)
		}
		numOut += *(*Tint64_t)(unsafe.Pointer(bp))
		goto _1
	_1:
		;
		i++
	}
	*(*Tint64_t)(unsafe.Pointer(out)) = numOut
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * isResClassIII takes a hexagon ID and determines if it is in a
//	 * Class III resolution (rotated versus the icosahedron and subject
//	 * to shape distortion adding extra points on icosahedron edges, making
//	 * them not true hexagons).
//	 * @param h The H3Index to check.
//	 * @return Returns 1 if the hexagon is class III, otherwise 0.
//	 */
func XisResClassIII(tls *libc.TLS, h TH3Index) (r int32) {
	return libc.Int32FromUint64(h&(libc.Uint64FromUint64(15)<<libc.Int32FromInt32(DH3_RES_OFFSET))>>libc.Int32FromInt32(DH3_RES_OFFSET)) % int32(2)
}

// C documentation
//
//	/**
//	 * isPentagon takes an H3Index and determines if it is actually a pentagon.
//	 * @param h The H3Index to check.
//	 * @return Returns 1 if it is a pentagon, otherwise 0.
//	 */
func XisPentagon(tls *libc.TLS, h TH3Index) (r int32) {
	return libc.BoolInt32(X_isBaseCellPentagon(tls, libc.Int32FromUint64(h&(libc.Uint64FromInt32(libc.Int32FromInt32(127))<<libc.Int32FromInt32(DH3_BC_OFFSET))>>libc.Int32FromInt32(DH3_BC_OFFSET))) != 0 && !(X_h3LeadingNonZeroDigit(tls, h) != 0))
}

// C documentation
//
//	/**
//	 * Returns the highest resolution non-zero digit in an H3Index.
//	 * @param h The H3Index.
//	 * @return The highest resolution non-zero digit in the H3Index.
//	 */
func X_h3LeadingNonZeroDigit(tls *libc.TLS, h TH3Index) (r1 TDirection) {
	var r int32
	_ = r
	r = int32(1)
	for {
		if !(r <= libc.Int32FromUint64(h&(libc.Uint64FromUint64(15)<<libc.Int32FromInt32(DH3_RES_OFFSET))>>libc.Int32FromInt32(DH3_RES_OFFSET))) {
			break
		}
		if libc.Int32FromUint64(h>>((libc.Int32FromInt32(DMAX_H3_RES)-r)*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))&libc.Uint64FromInt32(libc.Int32FromInt32(7))) != 0 {
			return libc.Int32FromUint64(h >> ((libc.Int32FromInt32(DMAX_H3_RES) - r) * libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET)) & libc.Uint64FromInt32(libc.Int32FromInt32(7)))
		}
		goto _1
	_1:
		;
		r++
	}
	// if we're here it's all 0's
	return int32(_CENTER_DIGIT)
}

// C documentation
//
//	/**
//	 * Rotate an H3Index 60 degrees counter-clockwise about a pentagonal center.
//	 * @param h The H3Index.
//	 */
func X_h3RotatePent60ccw(tls *libc.TLS, h TH3Index) (r1 TH3Index) {
	var foundFirstNonZeroDigit, r, res int32
	_, _, _ = foundFirstNonZeroDigit, r, res
	// rotate in place; skips any leading 1 digits (k-axis)
	foundFirstNonZeroDigit = 0
	r = int32(1)
	res = libc.Int32FromUint64(h & (libc.Uint64FromUint64(15) << libc.Int32FromInt32(DH3_RES_OFFSET)) >> libc.Int32FromInt32(DH3_RES_OFFSET))
	for {
		if !(r <= res) {
			break
		}
		// rotate this digit
		h = h & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<((libc.Int32FromInt32(DMAX_H3_RES)-r)*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))) | libc.Uint64FromInt32(X_rotate60ccw(tls, libc.Int32FromUint64(h>>((libc.Int32FromInt32(DMAX_H3_RES)-r)*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))&libc.Uint64FromInt32(libc.Int32FromInt32(7)))))<<((libc.Int32FromInt32(DMAX_H3_RES)-r)*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))
		// look for the first non-zero digit so we
		// can adjust for deleted k-axes sequence
		// if necessary
		if !(foundFirstNonZeroDigit != 0) && libc.Int32FromUint64(h>>((libc.Int32FromInt32(DMAX_H3_RES)-r)*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))&libc.Uint64FromInt32(libc.Int32FromInt32(7))) != 0 {
			foundFirstNonZeroDigit = int32(1)
			// adjust for deleted k-axes sequence
			if X_h3LeadingNonZeroDigit(tls, h) == int32(_K_AXES_DIGIT) {
				h = X_h3Rotate60ccw(tls, h)
			}
		}
		goto _1
	_1:
		;
		r++
	}
	return h
}

// C documentation
//
//	/**
//	 * Rotate an H3Index 60 degrees clockwise about a pentagonal center.
//	 * @param h The H3Index.
//	 */
func X_h3RotatePent60cw(tls *libc.TLS, h TH3Index) (r1 TH3Index) {
	var foundFirstNonZeroDigit, r, res int32
	_, _, _ = foundFirstNonZeroDigit, r, res
	// rotate in place; skips any leading 1 digits (k-axis)
	foundFirstNonZeroDigit = 0
	r = int32(1)
	res = libc.Int32FromUint64(h & (libc.Uint64FromUint64(15) << libc.Int32FromInt32(DH3_RES_OFFSET)) >> libc.Int32FromInt32(DH3_RES_OFFSET))
	for {
		if !(r <= res) {
			break
		}
		// rotate this digit
		h = h & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<((libc.Int32FromInt32(DMAX_H3_RES)-r)*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))) | libc.Uint64FromInt32(X_rotate60cw(tls, libc.Int32FromUint64(h>>((libc.Int32FromInt32(DMAX_H3_RES)-r)*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))&libc.Uint64FromInt32(libc.Int32FromInt32(7)))))<<((libc.Int32FromInt32(DMAX_H3_RES)-r)*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))
		// look for the first non-zero digit so we
		// can adjust for deleted k-axes sequence
		// if necessary
		if !(foundFirstNonZeroDigit != 0) && libc.Int32FromUint64(h>>((libc.Int32FromInt32(DMAX_H3_RES)-r)*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))&libc.Uint64FromInt32(libc.Int32FromInt32(7))) != 0 {
			foundFirstNonZeroDigit = int32(1)
			// adjust for deleted k-axes sequence
			if X_h3LeadingNonZeroDigit(tls, h) == int32(_K_AXES_DIGIT) {
				h = X_h3Rotate60cw(tls, h)
			}
		}
		goto _1
	_1:
		;
		r++
	}
	return h
}

// C documentation
//
//	/**
//	 * Rotate an H3Index 60 degrees counter-clockwise.
//	 * @param h The H3Index.
//	 */
func X_h3Rotate60ccw(tls *libc.TLS, h TH3Index) (r1 TH3Index) {
	var oldDigit TDirection
	var r, res int32
	_, _, _ = oldDigit, r, res
	r = int32(1)
	res = libc.Int32FromUint64(h & (libc.Uint64FromUint64(15) << libc.Int32FromInt32(DH3_RES_OFFSET)) >> libc.Int32FromInt32(DH3_RES_OFFSET))
	for {
		if !(r <= res) {
			break
		}
		oldDigit = libc.Int32FromUint64(h >> ((libc.Int32FromInt32(DMAX_H3_RES) - r) * libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET)) & libc.Uint64FromInt32(libc.Int32FromInt32(7)))
		h = h & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<((libc.Int32FromInt32(DMAX_H3_RES)-r)*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))) | libc.Uint64FromInt32(X_rotate60ccw(tls, oldDigit))<<((libc.Int32FromInt32(DMAX_H3_RES)-r)*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))
		goto _1
	_1:
		;
		r++
	}
	return h
}

// C documentation
//
//	/**
//	 * Rotate an H3Index 60 degrees clockwise.
//	 * @param h The H3Index.
//	 */
func X_h3Rotate60cw(tls *libc.TLS, h TH3Index) (r1 TH3Index) {
	var r, res int32
	_, _ = r, res
	r = int32(1)
	res = libc.Int32FromUint64(h & (libc.Uint64FromUint64(15) << libc.Int32FromInt32(DH3_RES_OFFSET)) >> libc.Int32FromInt32(DH3_RES_OFFSET))
	for {
		if !(r <= res) {
			break
		}
		h = h & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<((libc.Int32FromInt32(DMAX_H3_RES)-r)*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))) | libc.Uint64FromInt32(X_rotate60cw(tls, libc.Int32FromUint64(h>>((libc.Int32FromInt32(DMAX_H3_RES)-r)*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))&libc.Uint64FromInt32(libc.Int32FromInt32(7)))))<<((libc.Int32FromInt32(DMAX_H3_RES)-r)*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))
		goto _1
	_1:
		;
		r++
	}
	return h
}

// C documentation
//
//	/**
//	 * Convert an FaceIJK address to the corresponding H3Index.
//	 * @param fijk The FaceIJK address.
//	 * @param res The cell resolution.
//	 * @return The encoded H3Index (or H3_NULL on failure).
//	 */
func X_faceIjkToH3(tls *libc.TLS, fijk uintptr, res int32) (r1 TH3Index) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var baseCell, i, i1, numRots, r int32
	var h TH3Index
	var ijk uintptr
	var _ /* diff at bp+40 */ TCoordIJK
	var _ /* fijkBC at bp+0 */ TFaceIJK
	var _ /* lastCenter at bp+28 */ TCoordIJK
	var _ /* lastIJK at bp+16 */ TCoordIJK
	_, _, _, _, _, _, _ = baseCell, h, i, i1, ijk, numRots, r
	// initialize the index
	h = libc.Uint64FromUint64(35184372088831)
	h = h & ^(libc.Uint64FromInt32(libc.Int32FromInt32(15))<<libc.Int32FromInt32(DH3_MODE_OFFSET)) | libc.Uint64FromInt32(libc.Int32FromInt32(DH3_CELL_MODE))<<libc.Int32FromInt32(DH3_MODE_OFFSET)
	h = h & ^(libc.Uint64FromUint64(15)<<libc.Int32FromInt32(DH3_RES_OFFSET)) | libc.Uint64FromInt32(res)<<libc.Int32FromInt32(DH3_RES_OFFSET)
	// check for res 0/base cell
	if res == 0 {
		if (*TFaceIJK)(unsafe.Pointer(fijk)).Fcoord.Fi > int32(DMAX_FACE_COORD) || (*TFaceIJK)(unsafe.Pointer(fijk)).Fcoord.Fj > int32(DMAX_FACE_COORD) || (*TFaceIJK)(unsafe.Pointer(fijk)).Fcoord.Fk > int32(DMAX_FACE_COORD) {
			// out of range input
			return uint64(DH3_NULL)
		}
		h = h & ^(libc.Uint64FromInt32(libc.Int32FromInt32(127))<<libc.Int32FromInt32(DH3_BC_OFFSET)) | libc.Uint64FromInt32(X_faceIjkToBaseCell(tls, fijk))<<libc.Int32FromInt32(DH3_BC_OFFSET)
		return h
	}
	// we need to find the correct base cell FaceIJK for this H3 index;
	// start with the passed in face and resolution res ijk coordinates
	// in that face's coordinate system
	*(*TFaceIJK)(unsafe.Pointer(bp)) = *(*TFaceIJK)(unsafe.Pointer(fijk))
	// build the H3Index from finest res up
	// adjust r for the fact that the res 0 base cell offsets the indexing
	// digits
	ijk = bp + 4
	r = res - int32(1)
	for {
		if !(r >= 0) {
			break
		}
		*(*TCoordIJK)(unsafe.Pointer(bp + 16)) = *(*TCoordIJK)(unsafe.Pointer(ijk))
		if XisResolutionClassIII(tls, r+int32(1)) != 0 {
			// rotate ccw
			X_upAp7(tls, ijk)
			*(*TCoordIJK)(unsafe.Pointer(bp + 28)) = *(*TCoordIJK)(unsafe.Pointer(ijk))
			X_downAp7(tls, bp+28)
		} else {
			// rotate cw
			X_upAp7r(tls, ijk)
			*(*TCoordIJK)(unsafe.Pointer(bp + 28)) = *(*TCoordIJK)(unsafe.Pointer(ijk))
			X_downAp7r(tls, bp+28)
		}
		X_ijkSub(tls, bp+16, bp+28, bp+40)
		X_ijkNormalize(tls, bp+40)
		h = h & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<((libc.Int32FromInt32(DMAX_H3_RES)-(r+libc.Int32FromInt32(1)))*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))) | libc.Uint64FromInt32(X_unitIjkToDigit(tls, bp+40))<<((libc.Int32FromInt32(DMAX_H3_RES)-(r+libc.Int32FromInt32(1)))*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))
		goto _1
	_1:
		;
		r--
	}
	// fijkBC should now hold the IJK of the base cell in the
	// coordinate system of the current face
	if (*(*TFaceIJK)(unsafe.Pointer(bp))).Fcoord.Fi > int32(DMAX_FACE_COORD) || (*(*TFaceIJK)(unsafe.Pointer(bp))).Fcoord.Fj > int32(DMAX_FACE_COORD) || (*(*TFaceIJK)(unsafe.Pointer(bp))).Fcoord.Fk > int32(DMAX_FACE_COORD) {
		// out of range input
		return uint64(DH3_NULL)
	}
	// lookup the correct base cell
	baseCell = X_faceIjkToBaseCell(tls, bp)
	h = h & ^(libc.Uint64FromInt32(libc.Int32FromInt32(127))<<libc.Int32FromInt32(DH3_BC_OFFSET)) | libc.Uint64FromInt32(baseCell)<<libc.Int32FromInt32(DH3_BC_OFFSET)
	// rotate if necessary to get canonical base cell orientation
	// for this base cell
	numRots = X_faceIjkToBaseCellCCWrot60(tls, bp)
	if X_isBaseCellPentagon(tls, baseCell) != 0 {
		// force rotation out of missing k-axes sub-sequence
		if X_h3LeadingNonZeroDigit(tls, h) == int32(_K_AXES_DIGIT) {
			// check for a cw/ccw offset face; default is ccw
			if X_baseCellIsCwOffset(tls, baseCell, (*(*TFaceIJK)(unsafe.Pointer(bp))).Fface) != 0 {
				h = X_h3Rotate60cw(tls, h)
			} else {
				h = X_h3Rotate60ccw(tls, h)
			}
		}
		i = 0
		for {
			if !(i < numRots) {
				break
			}
			h = X_h3RotatePent60ccw(tls, h)
			goto _2
		_2:
			;
			i++
		}
	} else {
		i1 = 0
		for {
			if !(i1 < numRots) {
				break
			}
			h = X_h3Rotate60ccw(tls, h)
			goto _3
		_3:
			;
			i1++
		}
	}
	return h
}

// C documentation
//
//	/**
//	 * Encodes a coordinate on the sphere to the H3 index of the containing cell at
//	 * the specified resolution.
//	 *
//	 * Returns 0 on invalid input.
//	 *
//	 * @param g The spherical coordinates to encode.
//	 * @param res The desired H3 resolution for the encoding.
//	 * @param out The encoded H3Index.
//	 * @returns E_SUCCESS (0) on success, another value otherwise
//	 */
func XlatLngToCell(tls *libc.TLS, g uintptr, res int32, out uintptr) (r TH3Error) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var v1, v3 uint64
	var v5, v7 bool
	var v6 int32
	var _ /* __u at bp+0 */ struct {
		F__i [0]uint64
		F__f float64
	}
	var _ /* fijk at bp+8 */ TFaceIJK
	_, _, _, _, _ = v1, v3, v5, v6, v7
	if res < 0 || res > int32(DMAX_H3_RES) {
		return uint32(_E_RES_DOMAIN)
	}
	*(*float64)(unsafe.Pointer(bp)) = (*TLatLng)(unsafe.Pointer(g)).Flat
	v1 = *(*uint64)(unsafe.Pointer(bp))
	goto _2
_2:
	;
	if v5 = !(libc.BoolInt32(v1&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) < libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0); !v5 {
		*(*float64)(unsafe.Pointer(bp)) = (*TLatLng)(unsafe.Pointer(g)).Flng
		v3 = *(*uint64)(unsafe.Pointer(bp))
		goto _4
	_4:
	}
	if v5 || !(libc.BoolInt32(v3&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) < libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0) {
		return uint32(_E_LATLNG_DOMAIN)
	}
	X_geoToFaceIjk(tls, g, res, bp+8)
	*(*TH3Index)(unsafe.Pointer(out)) = X_faceIjkToH3(tls, bp+8, res)
	if *(*TH3Index)(unsafe.Pointer(out)) != 0 {
		v6 = int32(1)
	} else {
		if v7 = libc.Bool(0 != 0); !v7 {
			libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+829, int32(959), uintptr(unsafe.Pointer(&___func__11)))
		}
		_ = v7 || libc.Bool(libc.Int32FromInt32(0) != 0)
		v6 = libc.Int32FromInt32(0)
	}
	if v6 != 0 {
		return uint32(_E_SUCCESS)
	} else {
		return uint32(_E_FAILED)
	}
	return r
}

var ___func__11 = [13]uint8{'l', 'a', 't', 'L', 'n', 'g', 'T', 'o', 'C', 'e', 'l', 'l'}

// C documentation
//
//	/**
//	 * Convert an H3Index to the FaceIJK address on a specified icosahedral face.
//	 * @param h The H3Index.
//	 * @param fijk The FaceIJK address, initialized with the desired face
//	 *        and normalized base cell coordinates.
//	 * @return Returns 1 if the possibility of overage exists, otherwise 0.
//	 */
func X_h3ToFaceIjkWithInitializedFijk(tls *libc.TLS, h TH3Index, fijk uintptr) (r1 int32) {
	var ijk uintptr
	var possibleOverage, r, res int32
	_, _, _, _ = ijk, possibleOverage, r, res
	ijk = fijk + 4
	res = libc.Int32FromUint64(h & (libc.Uint64FromUint64(15) << libc.Int32FromInt32(DH3_RES_OFFSET)) >> libc.Int32FromInt32(DH3_RES_OFFSET))
	// center base cell hierarchy is entirely on this face
	possibleOverage = int32(1)
	if !(X_isBaseCellPentagon(tls, libc.Int32FromUint64(h&(libc.Uint64FromInt32(libc.Int32FromInt32(127))<<libc.Int32FromInt32(DH3_BC_OFFSET))>>libc.Int32FromInt32(DH3_BC_OFFSET))) != 0) && (res == 0 || (*TFaceIJK)(unsafe.Pointer(fijk)).Fcoord.Fi == 0 && (*TFaceIJK)(unsafe.Pointer(fijk)).Fcoord.Fj == 0 && (*TFaceIJK)(unsafe.Pointer(fijk)).Fcoord.Fk == 0) {
		possibleOverage = 0
	}
	r = int32(1)
	for {
		if !(r <= res) {
			break
		}
		if XisResolutionClassIII(tls, r) != 0 {
			// Class III == rotate ccw
			X_downAp7(tls, ijk)
		} else {
			// Class II == rotate cw
			X_downAp7r(tls, ijk)
		}
		X_neighbor(tls, ijk, libc.Int32FromUint64(h>>((libc.Int32FromInt32(DMAX_H3_RES)-r)*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))&libc.Uint64FromInt32(libc.Int32FromInt32(7))))
		goto _1
	_1:
		;
		r++
	}
	return possibleOverage
}

// C documentation
//
//	/**
//	 * Convert an H3Index to a FaceIJK address.
//	 * @param h The H3Index.
//	 * @param fijk The corresponding FaceIJK address.
//	 */
func X_h3ToFaceIjk(tls *libc.TLS, h TH3Index, fijk uintptr) (r TH3Error) {
	var baseCell, pentLeading4, res, v1, v3, v4 int32
	var origIJK TCoordIJK
	var v2 bool
	_, _, _, _, _, _, _, _ = baseCell, origIJK, pentLeading4, res, v1, v2, v3, v4
	baseCell = libc.Int32FromUint64(h & (libc.Uint64FromInt32(libc.Int32FromInt32(127)) << libc.Int32FromInt32(DH3_BC_OFFSET)) >> libc.Int32FromInt32(DH3_BC_OFFSET))
	if baseCell < 0 {
		if v2 = libc.Bool(0 != 0); !v2 {
			libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+829, int32(1006), uintptr(unsafe.Pointer(&___func__6)))
		}
		_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
		v1 = libc.Int32FromInt32(1)
	} else {
		v1 = 0
	}
	if v1 != 0 || baseCell >= int32(DNUM_BASE_CELLS) {
		// Base cells less than zero can not be represented in an index
		// To prevent reading uninitialized memory, we zero the output.
		(*TFaceIJK)(unsafe.Pointer(fijk)).Fface = 0
		v4 = libc.Int32FromInt32(0)
		(*TFaceIJK)(unsafe.Pointer(fijk)).Fcoord.Fk = v4
		v3 = v4
		(*TFaceIJK)(unsafe.Pointer(fijk)).Fcoord.Fj = v3
		(*TFaceIJK)(unsafe.Pointer(fijk)).Fcoord.Fi = v3
		return uint32(_E_CELL_INVALID)
	}
	// adjust for the pentagonal missing sequence; all of sub-sequence 5 needs
	// to be adjusted (and some of sub-sequence 4 below)
	if X_isBaseCellPentagon(tls, baseCell) != 0 && X_h3LeadingNonZeroDigit(tls, h) == int32(5) {
		h = X_h3Rotate60cw(tls, h)
	}
	// start with the "home" face and ijk+ coordinates for the base cell of c
	*(*TFaceIJK)(unsafe.Pointer(fijk)) = XbaseCellData[baseCell].FhomeFijk
	if !(X_h3ToFaceIjkWithInitializedFijk(tls, h, fijk) != 0) {
		return uint32(_E_SUCCESS)
	} // no overage is possible; h lies on this face
	// if we're here we have the potential for an "overage"; i.e., it is
	// possible that c lies on an adjacent face
	origIJK = (*TFaceIJK)(unsafe.Pointer(fijk)).Fcoord
	// if we're in Class III, drop into the next finer Class II grid
	res = libc.Int32FromUint64(h & (libc.Uint64FromUint64(15) << libc.Int32FromInt32(DH3_RES_OFFSET)) >> libc.Int32FromInt32(DH3_RES_OFFSET))
	if XisResolutionClassIII(tls, res) != 0 {
		// Class III
		X_downAp7r(tls, fijk+4)
		res++
	}
	// adjust for overage if needed
	// a pentagon base cell with a leading 4 digit requires special handling
	pentLeading4 = libc.BoolInt32(X_isBaseCellPentagon(tls, baseCell) != 0 && X_h3LeadingNonZeroDigit(tls, h) == int32(4))
	if X_adjustOverageClassII(tls, fijk, res, pentLeading4, 0) != int32(_NO_OVERAGE) {
		// if the base cell is a pentagon we have the potential for secondary
		// overages
		if X_isBaseCellPentagon(tls, baseCell) != 0 {
			for X_adjustOverageClassII(tls, fijk, res, 0, 0) != int32(_NO_OVERAGE) {
				continue
			}
		}
		if res != libc.Int32FromUint64(h&(libc.Uint64FromUint64(15)<<libc.Int32FromInt32(DH3_RES_OFFSET))>>libc.Int32FromInt32(DH3_RES_OFFSET)) {
			X_upAp7r(tls, fijk+4)
		}
	} else {
		if res != libc.Int32FromUint64(h&(libc.Uint64FromUint64(15)<<libc.Int32FromInt32(DH3_RES_OFFSET))>>libc.Int32FromInt32(DH3_RES_OFFSET)) {
			(*TFaceIJK)(unsafe.Pointer(fijk)).Fcoord = origIJK
		}
	}
	return uint32(_E_SUCCESS)
}

var ___func__6 = [13]uint8{'_', 'h', '3', 'T', 'o', 'F', 'a', 'c', 'e', 'I', 'j', 'k'}

// C documentation
//
//	/**
//	 * Determines the spherical coordinates of the center point of an H3 index.
//	 *
//	 * @param h3 The H3 index.
//	 * @param g The spherical coordinates of the H3 cell center.
//	 */
func XcellToLatLng(tls *libc.TLS, h3 TH3Index, g uintptr) (r TH3Error) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var e TH3Error
	var _ /* fijk at bp+0 */ TFaceIJK
	_ = e
	e = X_h3ToFaceIjk(tls, h3, bp)
	if e != 0 {
		return e
	}
	X_faceIjkToGeo(tls, bp, libc.Int32FromUint64(h3&(libc.Uint64FromUint64(15)<<libc.Int32FromInt32(DH3_RES_OFFSET))>>libc.Int32FromInt32(DH3_RES_OFFSET)), g)
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * Determines the cell boundary in spherical coordinates for an H3 index.
//	 *
//	 * @param h3 The H3 index.
//	 * @param cb The boundary of the H3 cell in spherical coordinates.
//	 */
func XcellToBoundary(tls *libc.TLS, h3 TH3Index, cb uintptr) (r TH3Error) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var e TH3Error
	var _ /* fijk at bp+0 */ TFaceIJK
	_ = e
	e = X_h3ToFaceIjk(tls, h3, bp)
	if e != 0 {
		return e
	}
	if XisPentagon(tls, h3) != 0 {
		X_faceIjkPentToCellBoundary(tls, bp, libc.Int32FromUint64(h3&(libc.Uint64FromUint64(15)<<libc.Int32FromInt32(DH3_RES_OFFSET))>>libc.Int32FromInt32(DH3_RES_OFFSET)), 0, int32(DNUM_PENT_VERTS), cb)
	} else {
		X_faceIjkToCellBoundary(tls, bp, libc.Int32FromUint64(h3&(libc.Uint64FromUint64(15)<<libc.Int32FromInt32(DH3_RES_OFFSET))>>libc.Int32FromInt32(DH3_RES_OFFSET)), 0, int32(DNUM_HEX_VERTS), cb)
	}
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * Returns the max number of possible icosahedron faces an H3 index
//	 * may intersect.
//	 *
//	 * @return int count of faces
//	 */
func XmaxFaceCount(tls *libc.TLS, h3 TH3Index, out uintptr) (r TH3Error) {
	var v1 int32
	_ = v1
	// a pentagon always intersects 5 faces, a hexagon never intersects more
	// than 2 (but may only intersect 1)
	if XisPentagon(tls, h3) != 0 {
		v1 = int32(5)
	} else {
		v1 = int32(2)
	}
	*(*int32)(unsafe.Pointer(out)) = v1
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * Find all icosahedron faces intersected by a given H3 index, represented
//	 * as integers from 0-19. The array is sparse; since 0 is a valid value,
//	 * invalid array values are represented as -1. It is the responsibility of
//	 * the caller to filter out invalid values.
//	 *
//	 * @param h3 The H3 index
//	 * @param out Output array. Must be of size maxFaceCount(h3).
//	 */
func XgetIcosahedronFaces(tls *libc.TLS, h3 TH3Index, out uintptr) (r TH3Error) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var childPentagon TH3Index
	var err, maxFaceCountError TH3Error
	var face, i, i1, isPent, pos, vertexCount, v1 int32
	var vert uintptr
	var v2 bool
	var _ /* faceCount at bp+116 */ int32
	var _ /* fijk at bp+4 */ TFaceIJK
	var _ /* fijkVerts at bp+20 */ [6]TFaceIJK
	var _ /* res at bp+0 */ int32
	_, _, _, _, _, _, _, _, _, _, _, _ = childPentagon, err, face, i, i1, isPent, maxFaceCountError, pos, vert, vertexCount, v1, v2
	*(*int32)(unsafe.Pointer(bp)) = libc.Int32FromUint64(h3 & (libc.Uint64FromUint64(15) << libc.Int32FromInt32(DH3_RES_OFFSET)) >> libc.Int32FromInt32(DH3_RES_OFFSET))
	isPent = XisPentagon(tls, h3)
	// We can't use the vertex-based approach here for class II pentagons,
	// because all their vertices are on the icosahedron edges. Their
	// direct child pentagons cross the same faces, so use those instead.
	if isPent != 0 && !(XisResolutionClassIII(tls, *(*int32)(unsafe.Pointer(bp))) != 0) {
		// Note that this would not work for res 15, but this is only run on
		// Class II pentagons, it should never be invoked for a res 15 index.
		childPentagon = XmakeDirectChild(tls, h3, 0)
		return XgetIcosahedronFaces(tls, childPentagon, out)
	}
	err = X_h3ToFaceIjk(tls, h3, bp+4)
	if err != 0 {
		return err
	}
	if isPent != 0 {
		vertexCount = int32(DNUM_PENT_VERTS)
		X_faceIjkPentToVerts(tls, bp+4, bp, bp+20)
	} else {
		vertexCount = int32(DNUM_HEX_VERTS)
		X_faceIjkToVerts(tls, bp+4, bp, bp+20)
	}
	maxFaceCountError = XmaxFaceCount(tls, h3, bp+116)
	if maxFaceCountError != uint32(_E_SUCCESS) {
		if v2 = libc.Bool(0 != 0); !v2 {
			libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+829, int32(1153), uintptr(unsafe.Pointer(&___func__7)))
		}
		_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
		v1 = libc.Int32FromInt32(1)
	} else {
		v1 = 0
	}
	if v1 != 0 {
		return maxFaceCountError
	}
	i = 0
	for {
		if !(i < *(*int32)(unsafe.Pointer(bp + 116))) {
			break
		}
		*(*int32)(unsafe.Pointer(out + uintptr(i)*4)) = -int32(1)
		goto _3
	_3:
		;
		i++
	}
	// add each vertex face, using the output array as a hash set
	i1 = 0
	for {
		if !(i1 < vertexCount) {
			break
		}
		vert = bp + 20 + uintptr(i1)*16
		// Adjust overage, determining whether this vertex is
		// on another face
		if isPent != 0 {
			X_adjustPentVertOverage(tls, vert, *(*int32)(unsafe.Pointer(bp)))
		} else {
			X_adjustOverageClassII(tls, vert, *(*int32)(unsafe.Pointer(bp)), 0, int32(1))
		}
		// Save the face to the output array
		face = (*TFaceIJK)(unsafe.Pointer(vert)).Fface
		pos = 0
		// Find the first empty output position, or the first position
		// matching the current face
		for *(*int32)(unsafe.Pointer(out + uintptr(pos)*4)) != -int32(1) && *(*int32)(unsafe.Pointer(out + uintptr(pos)*4)) != face {
			pos++
			if pos >= *(*int32)(unsafe.Pointer(bp + 116)) {
				// Mismatch between the heuristic used in maxFaceCount and
				// calculation here - indicates an invalid index.
				return uint32(_E_FAILED)
			}
		}
		*(*int32)(unsafe.Pointer(out + uintptr(pos)*4)) = face
		goto _4
	_4:
		;
		i1++
	}
	return uint32(_E_SUCCESS)
}

var ___func__7 = [20]uint8{'g', 'e', 't', 'I', 'c', 'o', 's', 'a', 'h', 'e', 'd', 'r', 'o', 'n', 'F', 'a', 'c', 'e', 's'}

// C documentation
//
//	/**
//	 * pentagonCount returns the number of pentagons (same at any resolution)
//	 *
//	 * @return int count of pentagon indexes
//	 */
func XpentagonCount(tls *libc.TLS) (r int32) {
	return int32(DNUM_PENTAGONS)
}

// C documentation
//
//	/**
//	 * Generates all pentagons at the specified resolution
//	 *
//	 * @param res The resolution to produce pentagons at.
//	 * @param out Output array. Must be of size pentagonCount().
//	 */
func XgetPentagons(tls *libc.TLS, res int32, out uintptr) (r TH3Error) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var bc, i, v2 int32
	var _ /* pentagon at bp+0 */ TH3Index
	_, _, _ = bc, i, v2
	if res < 0 || res > int32(DMAX_H3_RES) {
		return uint32(_E_RES_DOMAIN)
	}
	i = 0
	bc = 0
	for {
		if !(bc < int32(DNUM_BASE_CELLS)) {
			break
		}
		if X_isBaseCellPentagon(tls, bc) != 0 {
			XsetH3Index(tls, bp, res, bc, 0)
			v2 = i
			i++
			*(*TH3Index)(unsafe.Pointer(out + uintptr(v2)*8)) = *(*TH3Index)(unsafe.Pointer(bp))
		}
		goto _1
	_1:
		;
		bc++
	}
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * Returns whether or not a resolution is a Class III grid. Note that odd
//	 * resolutions are Class III and even resolutions are Class II.
//	 * @param res The H3 resolution.
//	 * @return 1 if the resolution is a Class III grid, and 0 if the resolution is
//	 *         a Class II grid.
//	 */
func XisResolutionClassIII(tls *libc.TLS, res int32) (r int32) {
	return res % int32(2)
}

// C documentation
//
//	/**
//	 * Validate a child position in the context of a given parent, returning
//	 * an error if validation fails.
//	 */
func _validateChildPos(tls *libc.TLS, childPos Tint64_t, parent TH3Index, childRes int32) (r TH3Error) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var sizeError TH3Error
	var v1 int32
	var v2 bool
	var _ /* maxChildCount at bp+0 */ Tint64_t
	_, _, _ = sizeError, v1, v2
	sizeError = XcellToChildrenSize(tls, parent, childRes, bp)
	if sizeError != 0 {
		if v2 = libc.Bool(0 != 0); !v2 {
			libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+829, int32(1236), uintptr(unsafe.Pointer(&___func__8)))
		}
		_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
		v1 = libc.Int32FromInt32(1)
	} else {
		v1 = 0
	}
	if v1 != 0 {
		return sizeError
	}
	if childPos < 0 || childPos >= *(*Tint64_t)(unsafe.Pointer(bp)) {
		return uint32(_E_DOMAIN)
	}
	return uint32(_E_SUCCESS)
}

var ___func__8 = [17]uint8{'v', 'a', 'l', 'i', 'd', 'a', 't', 'e', 'C', 'h', 'i', 'l', 'd', 'P', 'o', 's'}

// C documentation
//
//	/**
//	 * Returns the position of the cell within an ordered list of all children of
//	 * the cell's parent at the specified resolution
//	 * @param child Child cell index
//	 * @param parentRes Resolution of the parent cell to find the position within
//	 * @param out Output: The position of the child cell within its parents cell
//	 * list of children
//	 */
func XcellToChildPos(tls *libc.TLS, child TH3Index, parentRes int32, out uintptr) (r TH3Error) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var childRes, digit, digit1, parentIsPentagon, rawDigit, res, res1, v2, v4, v7 int32
	var hexChildCount Tint64_t
	var parentError, parentError1 TH3Error
	var v3, v8 bool
	var v5 int64
	var _ /* originalParent at bp+0 */ TH3Index
	var _ /* parent at bp+8 */ TH3Index
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = childRes, digit, digit1, hexChildCount, parentError, parentError1, parentIsPentagon, rawDigit, res, res1, v2, v3, v4, v5, v7, v8
	childRes = libc.Int32FromUint64(child & (libc.Uint64FromUint64(15) << libc.Int32FromInt32(DH3_RES_OFFSET)) >> libc.Int32FromInt32(DH3_RES_OFFSET))
	parentError = XcellToParent(tls, child, parentRes, bp)
	if parentError != 0 {
		return parentError
	}
	// Define the initial parent. Note that these variables are reassigned
	// within the loop.
	*(*TH3Index)(unsafe.Pointer(bp + 8)) = *(*TH3Index)(unsafe.Pointer(bp))
	parentIsPentagon = XisPentagon(tls, *(*TH3Index)(unsafe.Pointer(bp + 8)))
	// Walk up the resolution digits, incrementing the index
	*(*Tint64_t)(unsafe.Pointer(out)) = 0
	if parentIsPentagon != 0 {
		// Pentagon logic. Pentagon parents skip the 1 digit, so the offsets are
		// different from hexagons
		res = childRes
		for {
			if !(res > parentRes) {
				break
			}
			parentError1 = XcellToParent(tls, child, res-int32(1), bp+8)
			if parentError1 != 0 {
				if v3 = libc.Bool(0 != 0); !v3 {
					libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+829, int32(1276), uintptr(unsafe.Pointer(&___func__9)))
				}
				_ = v3 || libc.Bool(libc.Int32FromInt32(0) != 0)
				v2 = libc.Int32FromInt32(1)
			} else {
				v2 = 0
			}
			if v2 != 0 {
				return parentError1
			}
			parentIsPentagon = XisPentagon(tls, *(*TH3Index)(unsafe.Pointer(bp + 8)))
			rawDigit = libc.Int32FromUint64(child >> ((libc.Int32FromInt32(DMAX_H3_RES) - res) * libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET)) & libc.Uint64FromInt32(libc.Int32FromInt32(7)))
			// Validate the digit before proceeding
			if rawDigit == int32(_INVALID_DIGIT) || parentIsPentagon != 0 && rawDigit == int32(_K_AXES_DIGIT) {
				return uint32(_E_CELL_INVALID)
			}
			if parentIsPentagon != 0 && rawDigit > 0 {
				v4 = rawDigit - int32(1)
			} else {
				v4 = rawDigit
			}
			digit = v4
			if digit != int32(_CENTER_DIGIT) {
				hexChildCount = X_ipow(tls, int64(7), int64(childRes-res))
				// The offset for the 0-digit slot depends on whether the
				// current index is the child of a pentagon. If so, the offset
				// is based on the count of pentagon children, otherwise,
				// hexagon children.
				if parentIsPentagon != 0 {
					v5 = int64(1) + int64(5)*(hexChildCount-int64(1))/int64(6)
				} else {
					v5 = hexChildCount
				}
				*(*Tint64_t)(unsafe.Pointer(out)) += v5 + int64(digit-libc.Int32FromInt32(1))*hexChildCount
			}
			goto _1
		_1:
			;
			res--
		}
	} else {
		// Hexagon logic. Offsets are simple powers of 7
		res1 = childRes
		for {
			if !(res1 > parentRes) {
				break
			}
			digit1 = libc.Int32FromUint64(child >> ((libc.Int32FromInt32(DMAX_H3_RES) - res1) * libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET)) & libc.Uint64FromInt32(libc.Int32FromInt32(7)))
			if digit1 == int32(_INVALID_DIGIT) {
				return uint32(_E_CELL_INVALID)
			}
			*(*Tint64_t)(unsafe.Pointer(out)) += int64(digit1) * X_ipow(tls, int64(7), int64(childRes-res1))
			goto _6
		_6:
			;
			res1--
		}
	}
	if _validateChildPos(tls, *(*Tint64_t)(unsafe.Pointer(out)), *(*TH3Index)(unsafe.Pointer(bp)), childRes) != 0 {
		if v8 = libc.Bool(0 != 0); !v8 {
			libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+829, int32(1316), uintptr(unsafe.Pointer(&___func__9)))
		}
		_ = v8 || libc.Bool(libc.Int32FromInt32(0) != 0)
		v7 = libc.Int32FromInt32(1)
	} else {
		v7 = 0
	}
	if v7 != 0 {
		// This is the result of an internal error, so return E_FAILED
		// instead of the validation error
		return uint32(_E_FAILED)
	}
	return uint32(_E_SUCCESS)
}

var ___func__9 = [15]uint8{'c', 'e', 'l', 'l', 'T', 'o', 'C', 'h', 'i', 'l', 'd', 'P', 'o', 's'}

// C documentation
//
//	/**
//	 * Returns the child cell at a given position within an ordered list of all
//	 * children at the specified resolution
//	 * @param childPos Position within the ordered list
//	 * @param parent Parent cell of the cell index to find
//	 * @param childRes Resolution of the child cell index
//	 * @param child Output: child cell index
//	 */
func XchildPosToCell(tls *libc.TLS, childPos Tint64_t, parent TH3Index, childRes int32, child uintptr) (r TH3Error) {
	var childPosErr TH3Error
	var idx, pentWidth, resWidth, resWidth1 Tint64_t
	var inPent uint8
	var parentRes, res, res1, resOffset int32
	_, _, _, _, _, _, _, _, _, _ = childPosErr, idx, inPent, parentRes, pentWidth, res, res1, resOffset, resWidth, resWidth1
	// Validate resolution
	if childRes < 0 || childRes > int32(DMAX_H3_RES) {
		return uint32(_E_RES_DOMAIN)
	}
	// Validate parent resolution
	parentRes = libc.Int32FromUint64(parent & (libc.Uint64FromUint64(15) << libc.Int32FromInt32(DH3_RES_OFFSET)) >> libc.Int32FromInt32(DH3_RES_OFFSET))
	if childRes < parentRes {
		return uint32(_E_RES_MISMATCH)
	}
	// Validate child pos
	childPosErr = _validateChildPos(tls, childPos, parent, childRes)
	if childPosErr != 0 {
		return childPosErr
	}
	resOffset = childRes - parentRes
	*(*TH3Index)(unsafe.Pointer(child)) = parent
	idx = childPos
	*(*TH3Index)(unsafe.Pointer(child)) = *(*TH3Index)(unsafe.Pointer(child)) & ^(libc.Uint64FromUint64(15)<<libc.Int32FromInt32(DH3_RES_OFFSET)) | libc.Uint64FromInt32(childRes)<<libc.Int32FromInt32(DH3_RES_OFFSET)
	if XisPentagon(tls, parent) != 0 {
		// Pentagon tile logic. Pentagon tiles skip the 1 digit, so the offsets
		// are different
		inPent = uint8(Dtrue)
		res = int32(1)
		for {
			if !(res <= resOffset) {
				break
			}
			resWidth = X_ipow(tls, int64(7), int64(resOffset-res))
			if inPent != 0 {
				// While we are inside a parent pentagon, we need to check if
				// this cell is a pentagon, and if not, we need to offset its
				// digit to account for the skipped direction
				pentWidth = int64(1) + int64(5)*(resWidth-int64(1))/int64(6)
				if idx < pentWidth {
					*(*TH3Index)(unsafe.Pointer(child)) = *(*TH3Index)(unsafe.Pointer(child)) & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<((libc.Int32FromInt32(DMAX_H3_RES)-(parentRes+res))*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))) | libc.Uint64FromInt32(libc.Int32FromInt32(0))<<((libc.Int32FromInt32(DMAX_H3_RES)-(parentRes+res))*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))
				} else {
					idx -= pentWidth
					inPent = uint8(Dfalse)
					*(*TH3Index)(unsafe.Pointer(child)) = *(*TH3Index)(unsafe.Pointer(child)) & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<((libc.Int32FromInt32(DMAX_H3_RES)-(parentRes+res))*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))) | libc.Uint64FromInt64(idx/resWidth+libc.Int64FromInt32(2))<<((libc.Int32FromInt32(DMAX_H3_RES)-(parentRes+res))*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))
					idx %= resWidth
				}
			} else {
				// We're no longer inside a pentagon, continue as for hex
				*(*TH3Index)(unsafe.Pointer(child)) = *(*TH3Index)(unsafe.Pointer(child)) & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<((libc.Int32FromInt32(DMAX_H3_RES)-(parentRes+res))*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))) | libc.Uint64FromInt64(idx/resWidth)<<((libc.Int32FromInt32(DMAX_H3_RES)-(parentRes+res))*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))
				idx %= resWidth
			}
			goto _1
		_1:
			;
			res++
		}
	} else {
		// Hexagon tile logic. Offsets are simple powers of 7
		res1 = int32(1)
		for {
			if !(res1 <= resOffset) {
				break
			}
			resWidth1 = X_ipow(tls, int64(7), int64(resOffset-res1))
			*(*TH3Index)(unsafe.Pointer(child)) = *(*TH3Index)(unsafe.Pointer(child)) & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<((libc.Int32FromInt32(DMAX_H3_RES)-(parentRes+res1))*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))) | libc.Uint64FromInt64(idx/resWidth1)<<((libc.Int32FromInt32(DMAX_H3_RES)-(parentRes+res1))*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))
			idx %= resWidth1
			goto _2
		_2:
			;
			res1++
		}
	}
	return uint32(_E_SUCCESS)
}

const DINT32_MAX4 = 0x7fffffff

var _UNIT_VECS7 = [7]TCoordIJK{
	0: {},
	1: {
		Fk: int32(1),
	},
	2: {
		Fj: int32(1),
	},
	3: {
		Fj: int32(1),
		Fk: int32(1),
	},
	4: {
		Fi: int32(1),
	},
	5: {
		Fi: int32(1),
		Fk: int32(1),
	},
	6: {
		Fi: int32(1),
		Fj: int32(1),
	},
}

// C documentation
//
//	// extract the `res` digit (0--7) of the current cell
func __getResDigit(tls *libc.TLS, it uintptr, res int32) (r int32) {
	return libc.Int32FromUint64((*TIterCellsChildren)(unsafe.Pointer(it)).Fh >> ((libc.Int32FromInt32(DMAX_H3_RES) - res) * libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET)) & libc.Uint64FromInt32(libc.Int32FromInt32(7)))
}

// C documentation
//
//	// increment the digit (0--7) at location `res`
//	// H3_PER_DIGIT_OFFSET == 3
func __incrementResDigit(tls *libc.TLS, it uintptr, res int32) {
	var val TH3Index
	_ = val
	val = uint64(1)
	val <<= libc.Uint64FromInt32(int32(DH3_PER_DIGIT_OFFSET) * (int32(DMAX_H3_RES) - res))
	*(*TH3Index)(unsafe.Pointer(it)) += val
}

// C documentation
//
//	/**
//	 * Create a fully nulled-out child iterator for when an iterator is exhausted.
//	 * This helps minimize the chance that a user will depend on the iterator
//	 * internal state after it's exhausted, like the child resolution, for
//	 * example.
//	 */
func __null_iter(tls *libc.TLS) (r TIterCellsChildren) {
	return TIterCellsChildren{
		F_parentRes: -int32(1),
		F_skipDigit: -int32(1),
	}
}

/**

## Logic for iterating through the children of a cell

We'll describe the logic for ....

- normal (non pentagon iteration)
- pentagon iteration. define "pentagon digit"


### Cell Index Component Diagrams

The lower 56 bits of an H3 Cell Index describe the following index components:

- the cell resolution (4 bits)
- the base cell number (7 bits)
- the child cell digit for each resolution from 1 to 15 (3*15 = 45 bits)

These are the bits we'll be focused on when iterating through child cells.
To help describe the iteration logic, we'll use diagrams displaying the
(decimal) values for each component like:

                            child digit for resolution 2
                           /
| res | base cell # | 1 | 2 | 3 | 4 | 5 | 6 | ... |
|-----|-------------|---|---|---|---|---|---|-----|
|   9 |          17 | 5 | 3 | 0 | 6 | 2 | 1 | ... |


### Iteration through children of a hexagon (but not a pentagon)

Iteration through the children of a *hexagon* (but not a pentagon)
simply involves iterating through all the children values (0--6)
for each child digit (up to the child's resolution).

For example, suppose a resolution 3 hexagon index has the following
components:
                                parent resolution
                               /
| res | base cell # | 1 | 2 | 3 | 4 | 5 | 6 | ... |
|-----|-------------|---|---|---|---|---|---|-----|
|   3 |          17 | 3 | 5 | 1 | 7 | 7 | 7 | ... |

The iteration through all children of resolution 6 would look like:


                                parent res  child res
                               /           /
| res | base cell # | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | ... |
|-----|-------------|---|---|---|---|---|---|---|---|-----|
| 6   |          17 | 3 | 5 | 1 | 0 | 0 | 0 | 7 | 7 | ... |
| 6   |          17 | 3 | 5 | 1 | 0 | 0 | 1 | 7 | 7 | ... |
| ... |             |   |   |   |   |   |   |   |   |     |
| 6   |          17 | 3 | 5 | 1 | 0 | 0 | 6 | 7 | 7 | ... |
| 6   |          17 | 3 | 5 | 1 | 0 | 1 | 0 | 7 | 7 | ... |
| 6   |          17 | 3 | 5 | 1 | 0 | 1 | 1 | 7 | 7 | ... |
| ... |             |   |   |   |   |   |   |   |   |     |
| 6   |          17 | 3 | 5 | 1 | 6 | 6 | 6 | 7 | 7 | ... |


### Step sequence on a *pentagon* cell

Pentagon cells have a base cell number (e.g., 97) corresponding to a
resolution 0 pentagon, and have all zeros from digit 1 to the digit
corresponding to the cell's resolution.
(We'll drop the ellipses from now on, knowing that digits should contain
7's beyond the cell resolution.)

                            parent res      child res
                           /               /
| res | base cell # | 1 | 2 | 3 | 4 | 5 | 6 |
|-----|-------------|---|---|---|---|---|---|
|   6 |          97 | 0 | 0 | 0 | 0 | 0 | 0 |

Iteration through children of a *pentagon* is almost the same
as *hexagon* iteration, except that we skip the *first* 1 value
that appears in the "skip digit". This corresponds to the fact
that a pentagon only has 6 children, which are denoted with
the numbers {0,2,3,4,5,6}.

The skip digit starts at the child resolution position.
When iterating through children more than one resolution below
the parent, we move the skip digit to the left
(up to the next coarser resolution) each time we skip the 1 value
in that digit.

Iteration would start like:

                            parent res      child res
                           /               /
| res | base cell # | 1 | 2 | 3 | 4 | 5 | 6 |
|-----|-------------|---|---|---|---|---|---|
|   6 |          97 | 0 | 0 | 0 | 0 | 0 | 0 |
                                                                                       skip digit

Noticing we skip the 1 value and move the skip digit,
the next iterate would be:


| res | base cell # | 1 | 2 | 3 | 4 | 5 | 6 |
|-----|-------------|---|---|---|---|---|---|
|   6 |          97 | 0 | 0 | 0 | 0 | 0 | 2 |
                                                                               skip digit

Iteration continues normally until we get to:


| res | base cell # | 1 | 2 | 3 | 4 | 5 | 6 |
|-----|-------------|---|---|---|---|---|---|
|   6 |          97 | 0 | 0 | 0 | 0 | 0 | 6 |
                                                                               skip digit

which is followed by (skipping the 1):


| res | base cell # | 1 | 2 | 3 | 4 | 5 | 6 |
|-----|-------------|---|---|---|---|---|---|
|   6 |          97 | 0 | 0 | 0 | 0 | 2 | 0 |
                                                                       skip digit

For the next iterate, we won't skip the `1` in the previous digit
because it is no longer the skip digit:

| res | base cell # | 1 | 2 | 3 | 4 | 5 | 6 |
|-----|-------------|---|---|---|---|---|---|
|   6 |          97 | 0 | 0 | 0 | 0 | 2 | 1 |
                                                                       skip digit

Iteration continues normally until we're right before the next skip
digit:

| res | base cell # | 1 | 2 | 3 | 4 | 5 | 6 |
|-----|-------------|---|---|---|---|---|---|
|   6 |          97 | 0 | 0 | 0 | 0 | 6 | 6 |
                                                                       skip digit

Which is followed by

| res | base cell # | 1 | 2 | 3 | 4 | 5 | 6 |
|-----|-------------|---|---|---|---|---|---|
|   6 |          97 | 0 | 0 | 0 | 2 | 0 | 0 |
                                                               skip digit

and so on.

*/

// C documentation
//
//	/**
//	 * Initialize a IterCellsChildren struct representing the sequence giving
//	 * the children of cell `h` at resolution `childRes`.
//	 *
//	 * At any point in the iteration, starting once
//	 * the struct is initialized, IterCellsChildren.h gives the current child.
//	 *
//	 * Also, IterCellsChildren.h == H3_NULL when all the children have been iterated
//	 * through, or if the input to `iterInitParent` was invalid.
//	 */
func XiterInitParent(tls *libc.TLS, h TH3Index, childRes int32) (r TIterCellsChildren) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* it at bp+0 */ TIterCellsChildren
	X_iterInitParent(tls, h, childRes, bp)
	return *(*TIterCellsChildren)(unsafe.Pointer(bp))
}

// C documentation
//
//	/**
//	 * Internal function - initialize a parent iterator in-place
//	 */
func X_iterInitParent(tls *libc.TLS, h TH3Index, childRes int32, iter uintptr) {
	(*TIterCellsChildren)(unsafe.Pointer(iter)).F_parentRes = libc.Int32FromUint64(h & (libc.Uint64FromUint64(15) << libc.Int32FromInt32(DH3_RES_OFFSET)) >> libc.Int32FromInt32(DH3_RES_OFFSET))
	if childRes < (*TIterCellsChildren)(unsafe.Pointer(iter)).F_parentRes || childRes > int32(DMAX_H3_RES) || h == uint64(DH3_NULL) {
		*(*TIterCellsChildren)(unsafe.Pointer(iter)) = __null_iter(tls)
		return
	}
	(*TIterCellsChildren)(unsafe.Pointer(iter)).Fh = X_zeroIndexDigits(tls, h, (*TIterCellsChildren)(unsafe.Pointer(iter)).F_parentRes+int32(1), childRes)
	(*TIterCellsChildren)(unsafe.Pointer(iter)).Fh = (*TIterCellsChildren)(unsafe.Pointer(iter)).Fh & ^(libc.Uint64FromUint64(15)<<libc.Int32FromInt32(DH3_RES_OFFSET)) | libc.Uint64FromInt32(childRes)<<libc.Int32FromInt32(DH3_RES_OFFSET)
	if XisPentagon(tls, (*TIterCellsChildren)(unsafe.Pointer(iter)).Fh) != 0 {
		// The skip digit skips `1` for pentagons.
		// The "_skipDigit" moves to the left as we count up from the
		// child resolution to the parent resolution.
		(*TIterCellsChildren)(unsafe.Pointer(iter)).F_skipDigit = childRes
	} else {
		// if not a pentagon, we can ignore "skip digit" logic
		(*TIterCellsChildren)(unsafe.Pointer(iter)).F_skipDigit = -int32(1)
	}
}

// C documentation
//
//	/**
//	 * Step a IterCellsChildren to the next child cell.
//	 * When the iteration is over, IterCellsChildren.h will be H3_NULL.
//	 * Handles iterating through hexagon and pentagon cells.
//	 */
func XiterStepChild(tls *libc.TLS, it uintptr) {
	var childRes, i int32
	_, _ = childRes, i
	// once h == H3_NULL, the iterator returns an infinite sequence of H3_NULL
	if (*TIterCellsChildren)(unsafe.Pointer(it)).Fh == uint64(DH3_NULL) {
		return
	}
	childRes = libc.Int32FromUint64((*TIterCellsChildren)(unsafe.Pointer(it)).Fh & (libc.Uint64FromUint64(15) << libc.Int32FromInt32(DH3_RES_OFFSET)) >> libc.Int32FromInt32(DH3_RES_OFFSET))
	__incrementResDigit(tls, it, childRes)
	i = childRes
	for {
		if !(i >= (*TIterCellsChildren)(unsafe.Pointer(it)).F_parentRes) {
			break
		}
		if i == (*TIterCellsChildren)(unsafe.Pointer(it)).F_parentRes {
			// if we're modifying the parent resolution digit, then we're done
			*(*TIterCellsChildren)(unsafe.Pointer(it)) = __null_iter(tls)
			return
		}
		// PENTAGON_SKIPPED_DIGIT == 1
		if i == (*TIterCellsChildren)(unsafe.Pointer(it)).F_skipDigit && __getResDigit(tls, it, i) == int32(_PENTAGON_SKIPPED_DIGIT) {
			// Then we are iterating through the children of a pentagon cell.
			// All children of a pentagon have the property that the first
			// nonzero digit between the parent and child resolutions is
			// not 1.
			// I.e., we never see a sequence like 00001.
			// Thus, we skip the `1` in this digit.
			__incrementResDigit(tls, it, i)
			*(*int32)(unsafe.Pointer(it + 12)) -= int32(1)
			return
		}
		// INVALID_DIGIT == 7
		if __getResDigit(tls, it, i) == int32(_INVALID_DIGIT) {
			__incrementResDigit(tls, it, i) // zeros out it[i] and increments it[i-1] by 1
		} else {
			break
		}
		goto _1
	_1:
		;
		i--
	}
}

// C documentation
//
//	// create iterator for children of base cell at given resolution
func XiterInitBaseCellNum(tls *libc.TLS, baseCellNum int32, childRes int32) (r TIterCellsChildren) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* baseCell at bp+0 */ TH3Index
	if baseCellNum < 0 || baseCellNum >= int32(DNUM_BASE_CELLS) || childRes < 0 || childRes > int32(DMAX_H3_RES) {
		return __null_iter(tls)
	}
	XsetH3Index(tls, bp, 0, baseCellNum, 0)
	return XiterInitParent(tls, *(*TH3Index)(unsafe.Pointer(bp)), childRes)
}

// C documentation
//
//	// create iterator for all cells at given resolution
func XiterInitRes(tls *libc.TLS, res int32) (r TIterCellsResolution) {
	var itC TIterCellsChildren
	var itR TIterCellsResolution
	_, _ = itC, itR
	itC = XiterInitBaseCellNum(tls, 0, res)
	itR = TIterCellsResolution{
		Fh:    itC.Fh,
		F_res: res,
		F_itC: itC,
	}
	return itR
}

func XiterStepRes(tls *libc.TLS, itR uintptr) {
	// reached the end of over iterator; emits H3_NULL from now on
	if (*TIterCellsResolution)(unsafe.Pointer(itR)).Fh == uint64(DH3_NULL) {
		return
	}
	// step child iterator
	XiterStepChild(tls, itR+16)
	// If the child iterator is exhausted and there are still
	// base cells remaining, we initialize the next base cell child iterator
	if (*TIterCellsResolution)(unsafe.Pointer(itR)).F_itC.Fh == uint64(DH3_NULL) && (*TIterCellsResolution)(unsafe.Pointer(itR)).F_baseCellNum+int32(1) < int32(DNUM_BASE_CELLS) {
		*(*int32)(unsafe.Pointer(itR + 8)) += int32(1)
		(*TIterCellsResolution)(unsafe.Pointer(itR)).F_itC = XiterInitBaseCellNum(tls, (*TIterCellsResolution)(unsafe.Pointer(itR)).F_baseCellNum, (*TIterCellsResolution)(unsafe.Pointer(itR)).F_res)
	}
	// This overall iterator reflects the next cell in the child iterator.
	// Note: This sets itR->h = H3_NULL if the base cells were
	// exhausted in the check above.
	(*TIterCellsResolution)(unsafe.Pointer(itR)).Fh = (*TIterCellsResolution)(unsafe.Pointer(itR)).F_itC.Fh
}

const DEPSILON3 = 1e-16
const DEPSILON_DEG1 = 1e-09
const DINT32_MAX5 = 2147483647
const DM_180_PI1 = 57.29577951308232
const DM_2PI3 = 6.283185307179586
const DM_PI3 = 3.141592653589793
const DM_PI_1801 = 0.017453292519943295

// C documentation
//
//	/**
//	 * Normalizes radians to a value between 0.0 and two PI.
//	 *
//	 * @param rads The input radians value.
//	 * @return The normalized radians value.
//	 */
func X_posAngleRads(tls *libc.TLS, rads float64) (r float64) {
	var tmp, v1 float64
	_, _ = tmp, v1
	if rads < float64(0) {
		v1 = rads + float64(6.283185307179586)
	} else {
		v1 = rads
	}
	tmp = v1
	if rads >= float64(6.283185307179586) {
		tmp -= float64(6.283185307179586)
	}
	return tmp
}

// C documentation
//
//	/**
//	 * Determines if the components of two spherical coordinates are within some
//	 * threshold distance of each other.
//	 *
//	 * @param p1 The first spherical coordinates.
//	 * @param p2 The second spherical coordinates.
//	 * @param threshold The threshold distance.
//	 * @return Whether or not the two coordinates are within the threshold distance
//	 *         of each other.
//	 */
func XgeoAlmostEqualThreshold(tls *libc.TLS, p1 uintptr, p2 uintptr, threshold float64) (r uint8) {
	return libc.BoolUint8(libc.Xfabs(tls, (*TLatLng)(unsafe.Pointer(p1)).Flat-(*TLatLng)(unsafe.Pointer(p2)).Flat) < threshold && libc.Xfabs(tls, (*TLatLng)(unsafe.Pointer(p1)).Flng-(*TLatLng)(unsafe.Pointer(p2)).Flng) < threshold)
}

// C documentation
//
//	/**
//	 * Determines if the components of two spherical coordinates are within our
//	 * standard epsilon distance of each other.
//	 *
//	 * @param p1 The first spherical coordinates.
//	 * @param p2 The second spherical coordinates.
//	 * @return Whether or not the two coordinates are within the epsilon distance
//	 *         of each other.
//	 */
func XgeoAlmostEqual(tls *libc.TLS, p1 uintptr, p2 uintptr) (r uint8) {
	return XgeoAlmostEqualThreshold(tls, p1, p2, libc.Float64FromFloat64(1e-09)*libc.Float64FromFloat64(0.017453292519943295))
}

// C documentation
//
//	/**
//	 * Set the components of spherical coordinates in decimal degrees.
//	 *
//	 * @param p The spherical coordinates.
//	 * @param latDegs The desired latitude in decimal degrees.
//	 * @param lngDegs The desired longitude in decimal degrees.
//	 */
func XsetGeoDegs(tls *libc.TLS, p uintptr, latDegs float64, lngDegs float64) {
	X_setGeoRads(tls, p, XdegsToRads(tls, latDegs), XdegsToRads(tls, lngDegs))
}

// C documentation
//
//	/**
//	 * Set the components of spherical coordinates in radians.
//	 *
//	 * @param p The spherical coordinates.
//	 * @param latRads The desired latitude in decimal radians.
//	 * @param lngRads The desired longitude in decimal radians.
//	 */
func X_setGeoRads(tls *libc.TLS, p uintptr, latRads float64, lngRads float64) {
	(*TLatLng)(unsafe.Pointer(p)).Flat = latRads
	(*TLatLng)(unsafe.Pointer(p)).Flng = lngRads
}

// C documentation
//
//	/**
//	 * Convert from decimal degrees to radians.
//	 *
//	 * @param degrees The decimal degrees.
//	 * @return The corresponding radians.
//	 */
func XdegsToRads(tls *libc.TLS, degrees float64) (r float64) {
	return degrees * float64(0.017453292519943295)
}

// C documentation
//
//	/**
//	 * Convert from radians to decimal degrees.
//	 *
//	 * @param radians The radians.
//	 * @return The corresponding decimal degrees.
//	 */
func XradsToDegs(tls *libc.TLS, radians float64) (r float64) {
	return radians * float64(57.29577951308232)
}

// C documentation
//
//	/**
//	 * constrainLat makes sure latitudes are in the proper bounds
//	 *
//	 * @param lat The original lat value
//	 * @return The corrected lat value
//	 */
func XconstrainLat(tls *libc.TLS, lat float64) (r float64) {
	for lat > float64(1.5707963267948966) {
		lat = lat - float64(3.141592653589793)
	}
	return lat
}

// C documentation
//
//	/**
//	 * constrainLng makes sure longitudes are in the proper bounds
//	 *
//	 * @param lng The origin lng value
//	 * @return The corrected lng value
//	 */
func XconstrainLng(tls *libc.TLS, lng float64) (r float64) {
	for lng > float64(3.141592653589793) {
		lng = lng - libc.Float64FromInt32(2)*libc.Float64FromFloat64(3.141592653589793)
	}
	for lng < -libc.Float64FromFloat64(3.141592653589793) {
		lng = lng + libc.Float64FromInt32(2)*libc.Float64FromFloat64(3.141592653589793)
	}
	return lng
}

// C documentation
//
//	/**
//	 * Normalize an input longitude according to the specified normalization
//	 * @param  lng           Input longitude
//	 * @param  normalization Longitude normalization strategy
//	 * @return               Normalized longitude
//	 */
func XnormalizeLng(tls *libc.TLS, lng float64, normalization TLongitudeNormalization) (r float64) {
	var v1, v2 float64
	_, _ = v1, v2
	switch normalization {
	case int32(_NORMALIZE_EAST):
		if lng < libc.Float64FromInt32(0) {
			v1 = lng + libc.Float64FromFloat64(6.283185307179586)
		} else {
			v1 = lng
		}
		return v1
	case int32(_NORMALIZE_WEST):
		if lng > libc.Float64FromInt32(0) {
			v2 = lng - libc.Float64FromFloat64(6.283185307179586)
		} else {
			v2 = lng
		}
		return v2
	default:
		return lng
	}
	return r
}

// C documentation
//
//	/**
//	 * The great circle distance in radians between two spherical coordinates.
//	 *
//	 * This function uses the Haversine formula.
//	 * For math details, see:
//	 *     https://en.wikipedia.org/wiki/Haversine_formula
//	 *     https://www.movable-type.co.uk/scripts/latlong.html
//	 *
//	 * @param  a  the first lat/lng pair (in radians)
//	 * @param  b  the second lat/lng pair (in radians)
//	 *
//	 * @return    the great circle distance in radians between a and b
//	 */
func XgreatCircleDistanceRads(tls *libc.TLS, a uintptr, b uintptr) (r float64) {
	var A, sinLat, sinLng float64
	_, _, _ = A, sinLat, sinLng
	sinLat = libc.Xsin(tls, ((*TLatLng)(unsafe.Pointer(b)).Flat-(*TLatLng)(unsafe.Pointer(a)).Flat)*float64(0.5))
	sinLng = libc.Xsin(tls, ((*TLatLng)(unsafe.Pointer(b)).Flng-(*TLatLng)(unsafe.Pointer(a)).Flng)*float64(0.5))
	A = sinLat*sinLat + libc.Xcos(tls, (*TLatLng)(unsafe.Pointer(a)).Flat)*libc.Xcos(tls, (*TLatLng)(unsafe.Pointer(b)).Flat)*sinLng*sinLng
	return libc.Float64FromInt32(2) * libc.Xatan2(tls, libc.Xsqrt(tls, A), libc.Xsqrt(tls, libc.Float64FromInt32(1)-A))
}

// C documentation
//
//	/**
//	 * The great circle distance in kilometers between two spherical coordinates.
//	 *
//	 * @param  a  the first lat/lng pair (in radians)
//	 * @param  b  the second lat/lng pair (in radians)
//	 *
//	 * @return    the great circle distance in kilometers between a and b
//	 */
func XgreatCircleDistanceKm(tls *libc.TLS, a uintptr, b uintptr) (r float64) {
	return XgreatCircleDistanceRads(tls, a, b) * float64(DEARTH_RADIUS_KM)
}

// C documentation
//
//	/**
//	 * The great circle distance in meters between two spherical coordinates.
//	 *
//	 * @param  a  the first lat/lng pair (in radians)
//	 * @param  b  the second lat/lng pair (in radians)
//	 *
//	 * @return    the great circle distance in meters between a and b
//	 */
func XgreatCircleDistanceM(tls *libc.TLS, a uintptr, b uintptr) (r float64) {
	return XgreatCircleDistanceKm(tls, a, b) * libc.Float64FromInt32(1000)
}

// C documentation
//
//	/**
//	 * Determines the azimuth to p2 from p1 in radians.
//	 *
//	 * @param p1 The first spherical coordinates.
//	 * @param p2 The second spherical coordinates.
//	 * @return The azimuth in radians from p1 to p2.
//	 */
func X_geoAzimuthRads(tls *libc.TLS, p1 uintptr, p2 uintptr) (r float64) {
	return libc.Xatan2(tls, libc.Xcos(tls, (*TLatLng)(unsafe.Pointer(p2)).Flat)*libc.Xsin(tls, (*TLatLng)(unsafe.Pointer(p2)).Flng-(*TLatLng)(unsafe.Pointer(p1)).Flng), libc.Xcos(tls, (*TLatLng)(unsafe.Pointer(p1)).Flat)*libc.Xsin(tls, (*TLatLng)(unsafe.Pointer(p2)).Flat)-libc.Xsin(tls, (*TLatLng)(unsafe.Pointer(p1)).Flat)*libc.Xcos(tls, (*TLatLng)(unsafe.Pointer(p2)).Flat)*libc.Xcos(tls, (*TLatLng)(unsafe.Pointer(p2)).Flng-(*TLatLng)(unsafe.Pointer(p1)).Flng))
}

// C documentation
//
//	/**
//	 * Computes the point on the sphere a specified azimuth and distance from
//	 * another point.
//	 *
//	 * @param p1 The first spherical coordinates.
//	 * @param az The desired azimuth from p1.
//	 * @param distance The desired distance from p1, must be non-negative.
//	 * @param p2 The spherical coordinates at the desired azimuth and distance from
//	 * p1.
//	 */
func X_geoAzDistanceRads(tls *libc.TLS, p1 uintptr, az float64, distance float64, p2 uintptr) {
	var coslng, invcosp2lat, sinlat, sinlng float64
	_, _, _, _ = coslng, invcosp2lat, sinlat, sinlng
	if distance < float64(1e-16) {
		*(*TLatLng)(unsafe.Pointer(p2)) = *(*TLatLng)(unsafe.Pointer(p1))
		return
	}
	az = X_posAngleRads(tls, az)
	// check for due north/south azimuth
	if az < float64(1e-16) || libc.Xfabs(tls, az-float64(3.141592653589793)) < float64(1e-16) {
		if az < float64(1e-16) { // due north
			(*TLatLng)(unsafe.Pointer(p2)).Flat = (*TLatLng)(unsafe.Pointer(p1)).Flat + distance
		} else { // due south
			(*TLatLng)(unsafe.Pointer(p2)).Flat = (*TLatLng)(unsafe.Pointer(p1)).Flat - distance
		}
		if libc.Xfabs(tls, (*TLatLng)(unsafe.Pointer(p2)).Flat-float64(1.5707963267948966)) < float64(1e-16) { // north pole
			(*TLatLng)(unsafe.Pointer(p2)).Flat = float64(1.5707963267948966)
			(*TLatLng)(unsafe.Pointer(p2)).Flng = float64(0)
		} else {
			if libc.Xfabs(tls, (*TLatLng)(unsafe.Pointer(p2)).Flat+float64(1.5707963267948966)) < float64(1e-16) { // south pole
				(*TLatLng)(unsafe.Pointer(p2)).Flat = -libc.Float64FromFloat64(1.5707963267948966)
				(*TLatLng)(unsafe.Pointer(p2)).Flng = float64(0)
			} else {
				(*TLatLng)(unsafe.Pointer(p2)).Flng = XconstrainLng(tls, (*TLatLng)(unsafe.Pointer(p1)).Flng)
			}
		}
	} else { // not due north or south
		sinlat = libc.Xsin(tls, (*TLatLng)(unsafe.Pointer(p1)).Flat)*libc.Xcos(tls, distance) + libc.Xcos(tls, (*TLatLng)(unsafe.Pointer(p1)).Flat)*libc.Xsin(tls, distance)*libc.Xcos(tls, az)
		if sinlat > float64(1) {
			sinlat = float64(1)
		}
		if sinlat < -libc.Float64FromFloat64(1) {
			sinlat = -libc.Float64FromFloat64(1)
		}
		(*TLatLng)(unsafe.Pointer(p2)).Flat = libc.Xasin(tls, sinlat)
		if libc.Xfabs(tls, (*TLatLng)(unsafe.Pointer(p2)).Flat-float64(1.5707963267948966)) < float64(1e-16) { // north pole
			(*TLatLng)(unsafe.Pointer(p2)).Flat = float64(1.5707963267948966)
			(*TLatLng)(unsafe.Pointer(p2)).Flng = float64(0)
		} else {
			if libc.Xfabs(tls, (*TLatLng)(unsafe.Pointer(p2)).Flat+float64(1.5707963267948966)) < float64(1e-16) { // south pole
				(*TLatLng)(unsafe.Pointer(p2)).Flat = -libc.Float64FromFloat64(1.5707963267948966)
				(*TLatLng)(unsafe.Pointer(p2)).Flng = float64(0)
			} else {
				invcosp2lat = float64(1) / libc.Xcos(tls, (*TLatLng)(unsafe.Pointer(p2)).Flat)
				sinlng = libc.Xsin(tls, az) * libc.Xsin(tls, distance) * invcosp2lat
				coslng = (libc.Xcos(tls, distance) - libc.Xsin(tls, (*TLatLng)(unsafe.Pointer(p1)).Flat)*libc.Xsin(tls, (*TLatLng)(unsafe.Pointer(p2)).Flat)) / libc.Xcos(tls, (*TLatLng)(unsafe.Pointer(p1)).Flat) * invcosp2lat
				if sinlng > float64(1) {
					sinlng = float64(1)
				}
				if sinlng < -libc.Float64FromFloat64(1) {
					sinlng = -libc.Float64FromFloat64(1)
				}
				if coslng > float64(1) {
					coslng = float64(1)
				}
				if coslng < -libc.Float64FromFloat64(1) {
					coslng = -libc.Float64FromFloat64(1)
				}
				(*TLatLng)(unsafe.Pointer(p2)).Flng = XconstrainLng(tls, (*TLatLng)(unsafe.Pointer(p1)).Flng+libc.Xatan2(tls, sinlng, coslng))
			}
		}
	}
}

/*
 * The following functions provide meta information about the H3 hexagons at
 * each zoom level. Since there are only 16 total levels, these are current
 * handled with hardwired static values, but it may be worthwhile to put these
 * static values into another file that can be autogenerated by source code in
 * the future.
 */

func XgetHexagonAreaAvgKm2(tls *libc.TLS, res int32, out uintptr) (r TH3Error) {
	if res < 0 || res > int32(DMAX_H3_RES) {
		return uint32(_E_RES_DOMAIN)
	}
	*(*float64)(unsafe.Pointer(out)) = _areas[res]
	return uint32(_E_SUCCESS)
}

var _areas = [16]float64{
	0:  float64(4.357449416078383e+06),
	1:  float64(609788.4417941332),
	2:  float64(86801.7803989972),
	3:  float64(12393.43465508816),
	4:  float64(1770.347654491307),
	5:  float64(252.9038581819449),
	6:  float64(36.12906216441245),
	7:  float64(5.161293359717191),
	8:  float64(0.7373275975944177),
	9:  float64(0.1053325134272067),
	10: float64(0.01504750190766435),
	11: float64(0.002149643129451879),
	12: float64(0.000307091875631606),
	13: float64(4.387026794728296e-05),
	14: float64(6.267181135324313e-06),
	15: float64(8.95311590760579e-07),
}

func XgetHexagonAreaAvgM2(tls *libc.TLS, res int32, out uintptr) (r TH3Error) {
	if res < 0 || res > int32(DMAX_H3_RES) {
		return uint32(_E_RES_DOMAIN)
	}
	*(*float64)(unsafe.Pointer(out)) = _areas1[res]
	return uint32(_E_SUCCESS)
}

var _areas1 = [16]float64{
	0:  float64(4.35744941607839e+12),
	1:  float64(6.097884417941339e+11),
	2:  float64(8.680178039899731e+10),
	3:  float64(1.239343465508818e+10),
	4:  float64(1.770347654491309e+09),
	5:  float64(2.529038581819452e+08),
	6:  float64(3.61290621644125e+07),
	7:  float64(5.161293359717198e+06),
	8:  float64(737327.5975944188),
	9:  float64(105332.5134272069),
	10: float64(15047.50190766437),
	11: float64(2149.643129451882),
	12: float64(307.0918756316063),
	13: float64(43.87026794728301),
	14: float64(6.267181135324322),
	15: float64(0.8953115907605802),
}

func XgetHexagonEdgeLengthAvgKm(tls *libc.TLS, res int32, out uintptr) (r TH3Error) {
	if res < 0 || res > int32(DMAX_H3_RES) {
		return uint32(_E_RES_DOMAIN)
	}
	*(*float64)(unsafe.Pointer(out)) = _lens[res]
	return uint32(_E_SUCCESS)
}

var _lens = [16]float64{
	0:  float64(1281.256011),
	1:  float64(483.0568391),
	2:  float64(182.5129565),
	3:  float64(68.97922179),
	4:  float64(26.07175968),
	5:  float64(9.85409099),
	6:  float64(3.724532667),
	7:  float64(1.406475763),
	8:  float64(0.53141401),
	9:  float64(0.200786148),
	10: float64(0.075863783),
	11: float64(0.028663897),
	12: float64(0.010830188),
	13: float64(0.00409201),
	14: float64(0.0015461),
	15: float64(0.000584169),
}

func XgetHexagonEdgeLengthAvgM(tls *libc.TLS, res int32, out uintptr) (r TH3Error) {
	if res < 0 || res > int32(DMAX_H3_RES) {
		return uint32(_E_RES_DOMAIN)
	}
	*(*float64)(unsafe.Pointer(out)) = _lens1[res]
	return uint32(_E_SUCCESS)
}

var _lens1 = [16]float64{
	0:  float64(1.281256011e+06),
	1:  float64(483056.8391),
	2:  float64(182512.9565),
	3:  float64(68979.22179),
	4:  float64(26071.75968),
	5:  float64(9854.09099),
	6:  float64(3724.532667),
	7:  float64(1406.475763),
	8:  float64(531.4140101),
	9:  float64(200.7861476),
	10: float64(75.86378287),
	11: float64(28.66389748),
	12: float64(10.83018784),
	13: float64(4.092010473),
	14: float64(1.546099657),
	15: float64(0.58416863),
}

func XgetNumCells(tls *libc.TLS, res int32, out uintptr) (r TH3Error) {
	if res < 0 || res > int32(DMAX_H3_RES) {
		return uint32(_E_RES_DOMAIN)
	}
	*(*Tint64_t)(unsafe.Pointer(out)) = int64(2) + int64(120)*X_ipow(tls, int64(7), int64(res))
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * Surface area in radians^2 of spherical triangle on unit sphere.
//	 *
//	 * For the math, see:
//	 * https://en.wikipedia.org/wiki/Spherical_trigonometry#Area_and_spherical_excess
//	 *
//	 * @param   a  length of triangle side A in radians
//	 * @param   b  length of triangle side B in radians
//	 * @param   c  length of triangle side C in radians
//	 *
//	 * @return     area in radians^2 of triangle on unit sphere
//	 */
func XtriangleEdgeLengthsToArea(tls *libc.TLS, a float64, b float64, c float64) (r float64) {
	var s float64
	_ = s
	s = (a + b + c) * float64(0.5)
	a = (s - a) * float64(0.5)
	b = (s - b) * float64(0.5)
	c = (s - c) * float64(0.5)
	s = s * float64(0.5)
	return libc.Float64FromInt32(4) * libc.Xatan(tls, libc.Xsqrt(tls, libc.Xtan(tls, s)*libc.Xtan(tls, a)*libc.Xtan(tls, b)*libc.Xtan(tls, c)))
}

// C documentation
//
//	/**
//	 * Compute area in radians^2 of a spherical triangle, given its vertices.
//	 *
//	 * @param   a  vertex lat/lng in radians
//	 * @param   b  vertex lat/lng in radians
//	 * @param   c  vertex lat/lng in radians
//	 *
//	 * @return     area of triangle on unit sphere, in radians^2
//	 */
func XtriangleArea(tls *libc.TLS, a uintptr, b uintptr, c uintptr) (r float64) {
	return XtriangleEdgeLengthsToArea(tls, XgreatCircleDistanceRads(tls, a, b), XgreatCircleDistanceRads(tls, b, c), XgreatCircleDistanceRads(tls, c, a))
}

// C documentation
//
//	/**
//	 * Area of H3 cell in radians^2.
//	 *
//	 * The area is calculated by breaking the cell into spherical triangles and
//	 * summing up their areas. Note that some H3 cells (hexagons and pentagons)
//	 * are irregular, and have more than 6 or 5 sides.
//	 *
//	 * todo: optimize the computation by re-using the edges shared between triangles
//	 *
//	 * @param   cell  H3 cell
//	 * @param    out  cell area in radians^2
//	 * @return        E_SUCCESS on success, or an error code otherwise
//	 */
func XcellAreaRads2(tls *libc.TLS, cell TH3Index, out uintptr) (r TH3Error) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var area float64
	var err TH3Error
	var i, j, v1 int32
	var v2 bool
	var _ /* c at bp+0 */ TLatLng
	var _ /* cb at bp+16 */ TCellBoundary
	_, _, _, _, _, _ = area, err, i, j, v1, v2
	err = XcellToLatLng(tls, cell, bp)
	if err != 0 {
		return err
	}
	err = XcellToBoundary(tls, cell, bp+16)
	if err != 0 {
		if v2 = libc.Bool(0 != 0); !v2 {
			libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+856, int32(415), uintptr(unsafe.Pointer(&___func__10)))
		}
		_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
		v1 = libc.Int32FromInt32(1)
	} else {
		v1 = 0
	}
	if v1 != 0 {
		// Uncoverable because cellToLatLng will have returned an error already
		return err
	}
	area = float64(0)
	i = 0
	for {
		if !(i < (*(*TCellBoundary)(unsafe.Pointer(bp + 16))).FnumVerts) {
			break
		}
		j = (i + int32(1)) % (*(*TCellBoundary)(unsafe.Pointer(bp + 16))).FnumVerts
		area += XtriangleArea(tls, bp+16+8+uintptr(i)*16, bp+16+8+uintptr(j)*16, bp)
		goto _3
	_3:
		;
		i++
	}
	*(*float64)(unsafe.Pointer(out)) = area
	return uint32(_E_SUCCESS)
}

var ___func__10 = [14]uint8{'c', 'e', 'l', 'l', 'A', 'r', 'e', 'a', 'R', 'a', 'd', 's', '2'}

// C documentation
//
//	/**
//	 * Area of H3 cell in kilometers^2.
//	 *
//	 * @param   cell  H3 cell
//	 * @param    out  cell area in kilometers^2
//	 * @return        E_SUCCESS on success, or an error code otherwise
//	 */
func XcellAreaKm2(tls *libc.TLS, cell TH3Index, out uintptr) (r TH3Error) {
	var err TH3Error
	_ = err
	err = XcellAreaRads2(tls, cell, out)
	if !(err != 0) {
		*(*float64)(unsafe.Pointer(out)) = *(*float64)(unsafe.Pointer(out)) * float64(DEARTH_RADIUS_KM) * float64(DEARTH_RADIUS_KM)
	}
	return err
}

// C documentation
//
//	/**
//	 * Area of H3 cell in meters^2.
//	 *
//	 * @param   cell  H3 cell
//	 * @param    out  cell area in meters^2
//	 * @return        E_SUCCESS on success, or an error code otherwise
//	 */
func XcellAreaM2(tls *libc.TLS, cell TH3Index, out uintptr) (r TH3Error) {
	var err TH3Error
	_ = err
	err = XcellAreaKm2(tls, cell, out)
	if !(err != 0) {
		*(*float64)(unsafe.Pointer(out)) = *(*float64)(unsafe.Pointer(out)) * libc.Float64FromInt32(1000) * libc.Float64FromInt32(1000)
	}
	return err
}

// C documentation
//
//	/**
//	 * Length of a directed edge in radians.
//	 *
//	 * @param   edge  H3 directed edge
//	 * @param    length  length in radians
//	 * @return        E_SUCCESS on success, or an error code otherwise
//	 */
func XedgeLengthRads(tls *libc.TLS, edge TH3Index, length uintptr) (r TH3Error) {
	bp := tls.Alloc(176)
	defer tls.Free(176)
	var err TH3Error
	var i int32
	var _ /* cb at bp+0 */ TCellBoundary
	_, _ = err, i
	err = XdirectedEdgeToBoundary(tls, edge, bp)
	if err != 0 {
		return err
	}
	*(*float64)(unsafe.Pointer(length)) = float64(0)
	i = 0
	for {
		if !(i < (*(*TCellBoundary)(unsafe.Pointer(bp))).FnumVerts-int32(1)) {
			break
		}
		*(*float64)(unsafe.Pointer(length)) += XgreatCircleDistanceRads(tls, bp+8+uintptr(i)*16, bp+8+uintptr(i+int32(1))*16)
		goto _1
	_1:
		;
		i++
	}
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * Length of a directed edge in kilometers.
//	 *
//	 * @param   edge  H3 directed edge
//	 * @param    length  length in kilometers
//	 * @return        E_SUCCESS on success, or an error code otherwise
//	 */
func XedgeLengthKm(tls *libc.TLS, edge TH3Index, length uintptr) (r TH3Error) {
	var err TH3Error
	_ = err
	err = XedgeLengthRads(tls, edge, length)
	*(*float64)(unsafe.Pointer(length)) = *(*float64)(unsafe.Pointer(length)) * float64(DEARTH_RADIUS_KM)
	return err
}

// C documentation
//
//	/**
//	 * Length of a directed edge in meters.
//	 *
//	 * @param   edge  H3 directed edge
//	 * @param    length  length in meters
//	 * @return        E_SUCCESS on success, or an error code otherwise
//	 */
func XedgeLengthM(tls *libc.TLS, edge TH3Index, length uintptr) (r TH3Error) {
	var err TH3Error
	_ = err
	err = XedgeLengthKm(tls, edge, length)
	*(*float64)(unsafe.Pointer(length)) = *(*float64)(unsafe.Pointer(length)) * libc.Float64FromInt32(1000)
	return err
}

const DDBL_EPSILON1 = 2.220446049250313e-16
const DDBL_MAX1 = 1.7976931348623157e+308
const DEPSILON4 = 0.0000000000000001
const DEPSILON_DEG2 = .000000001
const DINT32_MAX6 = 0x7fffffff
const DM_180_PI2 = 57.29577951308232087679815481410517033240547
const DM_PI_1802 = 0.0174532925199432957692369076848861271111

/*
 * Copyright 2020 Uber Technologies, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/** @file alloc.h
 * @brief   Memory management functions
 *
 * This file contains macros and the necessary declarations to be able
 * to point H3 at different memory management functions than the standard
 * malloc/free/etc functions.
 */

/*
 * Copyright 2016-2021 Uber Technologies, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/** @file h3api.h
 * @brief   Primary H3 core library entry points.
 *
 * This file defines the public API of the H3 library. Incompatible changes to
 * these functions require the library's major version be increased.
 */

/*
 * Copyright 2016-2021 Uber Technologies, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/** @file h3api.h
 * @brief   Primary H3 core library entry points.
 *
 * This file defines the public API of the H3 library. Incompatible changes to
 * these functions require the library's major version be increased.
 */

/*
 * Copyright 2016-2021 Uber Technologies, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/** @file latLng.h
 * @brief   Geodetic (lat/lng) functions.
 */

// C documentation
//
//	/**
//	 * Add a linked polygon to the current polygon
//	 * @param  polygon Polygon to add link to
//	 * @return         Pointer to new polygon
//	 */
func XaddNewLinkedPolygon(tls *libc.TLS, polygon uintptr) (r uintptr) {
	var next uintptr
	var v1, v2 bool
	_, _, _ = next, v1, v2
	if v1 = (*TLinkedGeoPolygon)(unsafe.Pointer(polygon)).Fnext == libc.UintptrFromInt32(0); !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+882, __ccgo_ts+904, int32(35), uintptr(unsafe.Pointer(&___func__12)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
	next = libc.Xcalloc(tls, uint64(1), uint64(24))
	if v2 = next != libc.UintptrFromInt32(0); !v2 {
		libc.X__assert_fail(tls, __ccgo_ts+933, __ccgo_ts+904, int32(37), uintptr(unsafe.Pointer(&___func__12)))
	}
	_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
	(*TLinkedGeoPolygon)(unsafe.Pointer(polygon)).Fnext = next
	return next
}

var ___func__12 = [20]uint8{'a', 'd', 'd', 'N', 'e', 'w', 'L', 'i', 'n', 'k', 'e', 'd', 'P', 'o', 'l', 'y', 'g', 'o', 'n'}

// C documentation
//
//	/**
//	 * Add a new linked loop to the current polygon
//	 * @param  polygon Polygon to add loop to
//	 * @return         Pointer to loop
//	 */
func XaddNewLinkedLoop(tls *libc.TLS, polygon uintptr) (r uintptr) {
	var loop uintptr
	var v1 bool
	_, _ = loop, v1
	loop = libc.Xcalloc(tls, uint64(1), uint64(24))
	if v1 = loop != libc.UintptrFromInt32(0); !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+946, __ccgo_ts+904, int32(49), uintptr(unsafe.Pointer(&___func__13)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
	return XaddLinkedLoop(tls, polygon, loop)
}

var ___func__13 = [17]uint8{'a', 'd', 'd', 'N', 'e', 'w', 'L', 'i', 'n', 'k', 'e', 'd', 'L', 'o', 'o', 'p'}

// C documentation
//
//	/**
//	 * Add an existing linked loop to the current polygon
//	 * @param  polygon Polygon to add loop to
//	 * @return         Pointer to loop
//	 */
func XaddLinkedLoop(tls *libc.TLS, polygon uintptr, loop uintptr) (r uintptr) {
	var last uintptr
	var v1 bool
	_, _ = last, v1
	last = (*TLinkedGeoPolygon)(unsafe.Pointer(polygon)).Flast
	if last == libc.UintptrFromInt32(0) {
		if v1 = (*TLinkedGeoPolygon)(unsafe.Pointer(polygon)).Ffirst == libc.UintptrFromInt32(0); !v1 {
			libc.X__assert_fail(tls, __ccgo_ts+959, __ccgo_ts+904, int32(61), uintptr(unsafe.Pointer(&___func__14)))
		}
		_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
		(*TLinkedGeoPolygon)(unsafe.Pointer(polygon)).Ffirst = loop
	} else {
		(*TLinkedGeoLoop)(unsafe.Pointer(last)).Fnext = loop
	}
	(*TLinkedGeoPolygon)(unsafe.Pointer(polygon)).Flast = loop
	return loop
}

var ___func__14 = [14]uint8{'a', 'd', 'd', 'L', 'i', 'n', 'k', 'e', 'd', 'L', 'o', 'o', 'p'}

// C documentation
//
//	/**
//	 * Add a new linked coordinate to the current loop
//	 * @param  loop   Loop to add coordinate to
//	 * @param  vertex Coordinate to add
//	 * @return        Pointer to the coordinate
//	 */
func XaddLinkedCoord(tls *libc.TLS, loop uintptr, vertex uintptr) (r uintptr) {
	var coord, last uintptr
	var v1, v2 bool
	_, _, _, _ = coord, last, v1, v2
	coord = libc.Xmalloc(tls, uint64(24))
	if v1 = coord != libc.UintptrFromInt32(0); !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+982, __ccgo_ts+904, int32(78), uintptr(unsafe.Pointer(&___func__15)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
	*(*TLinkedLatLng)(unsafe.Pointer(coord)) = TLinkedLatLng{
		Fvertex: *(*TLatLng)(unsafe.Pointer(vertex)),
	}
	last = (*TLinkedGeoLoop)(unsafe.Pointer(loop)).Flast
	if last == libc.UintptrFromInt32(0) {
		if v2 = (*TLinkedGeoLoop)(unsafe.Pointer(loop)).Ffirst == libc.UintptrFromInt32(0); !v2 {
			libc.X__assert_fail(tls, __ccgo_ts+996, __ccgo_ts+904, int32(82), uintptr(unsafe.Pointer(&___func__15)))
		}
		_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
		(*TLinkedGeoLoop)(unsafe.Pointer(loop)).Ffirst = coord
	} else {
		(*TLinkedLatLng)(unsafe.Pointer(last)).Fnext = coord
	}
	(*TLinkedGeoLoop)(unsafe.Pointer(loop)).Flast = coord
	return coord
}

var ___func__15 = [15]uint8{'a', 'd', 'd', 'L', 'i', 'n', 'k', 'e', 'd', 'C', 'o', 'o', 'r', 'd'}

// C documentation
//
//	/**
//	 * Free all allocated memory for a linked geo loop. The caller is
//	 * responsible for freeing memory allocated to input loop struct.
//	 * @param loop Loop to free
//	 */
func XdestroyLinkedGeoLoop(tls *libc.TLS, loop uintptr) {
	var currentCoord, nextCoord uintptr
	_, _ = currentCoord, nextCoord
	currentCoord = (*TLinkedGeoLoop)(unsafe.Pointer(loop)).Ffirst
	for {
		if !(currentCoord != libc.UintptrFromInt32(0)) {
			break
		}
		nextCoord = (*TLinkedLatLng)(unsafe.Pointer(currentCoord)).Fnext
		libc.Xfree(tls, currentCoord)
		goto _1
	_1:
		;
		currentCoord = nextCoord
	}
}

// C documentation
//
//	/**
//	 * Free all allocated memory for a linked geo structure. The caller is
//	 * responsible for freeing memory allocated to input polygon struct.
//	 * @param polygon Pointer to the first polygon in the structure
//	 */
func XdestroyLinkedMultiPolygon(tls *libc.TLS, polygon uintptr) {
	var currentLoop, currentPolygon, nextLoop, nextPolygon uintptr
	var skip uint8
	_, _, _, _, _ = currentLoop, currentPolygon, nextLoop, nextPolygon, skip
	// flag to skip the input polygon
	skip = uint8(Dtrue)
	currentPolygon = polygon
	for {
		if !(currentPolygon != libc.UintptrFromInt32(0)) {
			break
		}
		currentLoop = (*TLinkedGeoPolygon)(unsafe.Pointer(currentPolygon)).Ffirst
		for {
			if !(currentLoop != libc.UintptrFromInt32(0)) {
				break
			}
			XdestroyLinkedGeoLoop(tls, currentLoop)
			nextLoop = (*TLinkedGeoLoop)(unsafe.Pointer(currentLoop)).Fnext
			libc.Xfree(tls, currentLoop)
			goto _2
		_2:
			;
			currentLoop = nextLoop
		}
		nextPolygon = (*TLinkedGeoPolygon)(unsafe.Pointer(currentPolygon)).Fnext
		if skip != 0 {
			// do not free the input polygon
			skip = uint8(Dfalse)
		} else {
			libc.Xfree(tls, currentPolygon)
		}
		goto _1
	_1:
		;
		currentPolygon = nextPolygon
	}
}

// C documentation
//
//	/**
//	 * Count the number of polygons in a linked list
//	 * @param  polygon Starting polygon
//	 * @return         Count
//	 */
func XcountLinkedPolygons(tls *libc.TLS, polygon uintptr) (r int32) {
	var count int32
	_ = count
	count = 0
	for polygon != libc.UintptrFromInt32(0) {
		count++
		polygon = (*TLinkedGeoPolygon)(unsafe.Pointer(polygon)).Fnext
	}
	return count
}

// C documentation
//
//	/**
//	 * Count the number of linked loops in a polygon
//	 * @param  polygon Polygon to count loops for
//	 * @return         Count
//	 */
func XcountLinkedLoops(tls *libc.TLS, polygon uintptr) (r int32) {
	var count int32
	var loop uintptr
	_, _ = count, loop
	loop = (*TLinkedGeoPolygon)(unsafe.Pointer(polygon)).Ffirst
	count = 0
	for loop != libc.UintptrFromInt32(0) {
		count++
		loop = (*TLinkedGeoLoop)(unsafe.Pointer(loop)).Fnext
	}
	return count
}

// C documentation
//
//	/**
//	 * Count the number of coordinates in a loop
//	 * @param  loop Loop to count coordinates for
//	 * @return      Count
//	 */
func XcountLinkedCoords(tls *libc.TLS, loop uintptr) (r int32) {
	var coord uintptr
	var count int32
	_, _ = coord, count
	coord = (*TLinkedGeoLoop)(unsafe.Pointer(loop)).Ffirst
	count = 0
	for coord != libc.UintptrFromInt32(0) {
		count++
		coord = (*TLinkedLatLng)(unsafe.Pointer(coord)).Fnext
	}
	return count
}

// C documentation
//
//	/**
//	 * Count the number of polygons containing a given loop.
//	 * @param  loop         Loop to count containers for
//	 * @param  polygons     Polygons to test
//	 * @param  bboxes       Bounding boxes for polygons, used in point-in-poly check
//	 * @param  polygonCount Number of polygons in the test array
//	 * @return              Number of polygons containing the loop
//	 */
func _countContainers(tls *libc.TLS, loop uintptr, polygons uintptr, bboxes uintptr, polygonCount int32) (r int32) {
	var containerCount, i int32
	_, _ = containerCount, i
	containerCount = 0
	i = 0
	for {
		if !(i < polygonCount) {
			break
		}
		if loop != (*TLinkedGeoPolygon)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(polygons + uintptr(i)*8)))).Ffirst && XpointInsideLinkedGeoLoop(tls, (*TLinkedGeoPolygon)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(polygons + uintptr(i)*8)))).Ffirst, *(*uintptr)(unsafe.Pointer(bboxes + uintptr(i)*8)), (*TLinkedGeoLoop)(unsafe.Pointer(loop)).Ffirst) != 0 {
			containerCount++
		}
		goto _1
	_1:
		;
		i++
	}
	return containerCount
}

// C documentation
//
//	/**
//	 * Given a list of nested containers, find the one most deeply nested.
//	 * @param  polygons     Polygon containers to check
//	 * @param  bboxes       Bounding boxes for polygons, used in point-in-poly check
//	 * @param  polygonCount Number of polygons in the list
//	 * @return              Deepest container, or null if list is empty
//	 */
func _findDeepestContainer(tls *libc.TLS, polygons uintptr, bboxes uintptr, polygonCount int32) (r uintptr) {
	var count, i, max int32
	var parent, v1 uintptr
	_, _, _, _, _ = count, i, max, parent, v1
	if polygonCount > 0 {
		v1 = *(*uintptr)(unsafe.Pointer(polygons))
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	// Set the initial return value to the first candidate
	parent = v1
	// If we have multiple polygons, they must be nested inside each other.
	// Find the innermost polygon by taking the one with the most containers
	// in the list.
	if polygonCount > int32(1) {
		max = -int32(1)
		i = 0
		for {
			if !(i < polygonCount) {
				break
			}
			count = _countContainers(tls, (*TLinkedGeoPolygon)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(polygons + uintptr(i)*8)))).Ffirst, polygons, bboxes, polygonCount)
			if count > max {
				parent = *(*uintptr)(unsafe.Pointer(polygons + uintptr(i)*8))
				max = count
			}
			goto _2
		_2:
			;
			i++
		}
	}
	return parent
}

// C documentation
//
//	/**
//	 * Find the polygon to which a given hole should be allocated. Note that this
//	 * function will return null if no parent is found.
//	 * @param  loop         Inner loop describing a hole
//	 * @param  polygon      Head of a linked list of polygons to check
//	 * @param  bboxes       Bounding boxes for polygons, used in point-in-poly check
//	 * @param  polygonCount Number of polygons to check
//	 * @return              Pointer to parent polygon, or null if not found
//	 */
func _findPolygonForHole(tls *libc.TLS, loop uintptr, polygon uintptr, bboxes uintptr, polygonCount int32) (r uintptr) {
	var candidateBBoxes, candidates, parent uintptr
	var candidateCount, index int32
	var v1, v2 bool
	_, _, _, _, _, _, _ = candidateBBoxes, candidateCount, candidates, index, parent, v1, v2
	// Early exit with no polygons
	if polygonCount == 0 {
		return libc.UintptrFromInt32(0)
	}
	// Initialize arrays for candidate loops and their bounding boxes
	candidates = libc.Xmalloc(tls, libc.Uint64FromInt32(polygonCount)*uint64(8))
	if v1 = candidates != libc.UintptrFromInt32(0); !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+1016, __ccgo_ts+904, int32(249), uintptr(unsafe.Pointer(&___func__16)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
	candidateBBoxes = libc.Xmalloc(tls, libc.Uint64FromInt32(polygonCount)*uint64(8))
	if v2 = candidateBBoxes != libc.UintptrFromInt32(0); !v2 {
		libc.X__assert_fail(tls, __ccgo_ts+1035, __ccgo_ts+904, int32(252), uintptr(unsafe.Pointer(&___func__16)))
	}
	_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
	// Find all polygons that contain the loop
	candidateCount = 0
	index = 0
	for polygon != 0 {
		// We are guaranteed not to overlap, so just test the first point
		if XpointInsideLinkedGeoLoop(tls, (*TLinkedGeoPolygon)(unsafe.Pointer(polygon)).Ffirst, bboxes+uintptr(index)*32, (*TLinkedGeoLoop)(unsafe.Pointer(loop)).Ffirst) != 0 {
			*(*uintptr)(unsafe.Pointer(candidates + uintptr(candidateCount)*8)) = polygon
			*(*uintptr)(unsafe.Pointer(candidateBBoxes + uintptr(candidateCount)*8)) = bboxes + uintptr(index)*32
			candidateCount++
		}
		polygon = (*TLinkedGeoPolygon)(unsafe.Pointer(polygon)).Fnext
		index++
	}
	// The most deeply nested container is the immediate parent
	parent = _findDeepestContainer(tls, candidates, candidateBBoxes, candidateCount)
	// Free allocated memory
	libc.Xfree(tls, candidates)
	libc.Xfree(tls, candidateBBoxes)
	return parent
}

var ___func__16 = [19]uint8{'f', 'i', 'n', 'd', 'P', 'o', 'l', 'y', 'g', 'o', 'n', 'F', 'o', 'r', 'H', 'o', 'l', 'e'}

// C documentation
//
//	/**
//	 * Normalize a LinkedGeoPolygon in-place into a structure following GeoJSON
//	 * MultiPolygon rules: Each polygon must have exactly one outer loop, which
//	 * must be first in the list, followed by any holes. Holes in this algorithm
//	 * are identified by winding order (holes are clockwise), which is guaranteed
//	 * by the h3SetToVertexGraph algorithm.
//	 *
//	 * Input to this function is assumed to be a single polygon including all
//	 * loops to normalize. It's assumed that a valid arrangement is possible.
//	 *
//	 * @param root Root polygon including all loops
//	 * @return     0 on success, or an error code > 0 for invalid input
//	 */
func XnormalizeMultiPolygon(tls *libc.TLS, root uintptr) (r TH3Error) {
	var bboxes, innerLoops, loop, next, polygon, v3 uintptr
	var i, innerCount, loopCount, outerCount int32
	var resultCode TH3Error
	var v1, v2 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _ = bboxes, i, innerCount, innerLoops, loop, loopCount, next, outerCount, polygon, resultCode, v1, v2, v3
	// We assume that the input is a single polygon with loops;
	// if it has multiple polygons, don't touch it
	if (*TLinkedGeoPolygon)(unsafe.Pointer(root)).Fnext != 0 {
		return uint32(_E_FAILED)
	}
	// Count loops, exiting early if there's only one
	loopCount = XcountLinkedLoops(tls, root)
	if loopCount <= int32(1) {
		return uint32(_E_SUCCESS)
	}
	resultCode = uint32(_E_SUCCESS)
	polygon = libc.UintptrFromInt32(0)
	innerCount = 0
	outerCount = 0
	// Create an array to hold all of the inner loops. Note that
	// this array will never be full, as there will always be fewer
	// inner loops than outer loops.
	innerLoops = libc.Xmalloc(tls, libc.Uint64FromInt32(loopCount)*uint64(8))
	if v1 = innerLoops != libc.UintptrFromInt32(0); !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+1059, __ccgo_ts+904, int32(317), uintptr(unsafe.Pointer(&___func__17)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
	// Create an array to hold the bounding boxes for the outer loops
	bboxes = libc.Xmalloc(tls, libc.Uint64FromInt32(loopCount)*uint64(32))
	if v2 = bboxes != libc.UintptrFromInt32(0); !v2 {
		libc.X__assert_fail(tls, __ccgo_ts+1078, __ccgo_ts+904, int32(321), uintptr(unsafe.Pointer(&___func__17)))
	}
	_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
	// Get the first loop and unlink it from root
	loop = (*TLinkedGeoPolygon)(unsafe.Pointer(root)).Ffirst
	*(*TLinkedGeoPolygon)(unsafe.Pointer(root)) = TLinkedGeoPolygon{}
	// Iterate over all loops, moving inner loops into an array and
	// assigning outer loops to new polygons
	for loop != 0 {
		if XisClockwiseLinkedGeoLoop(tls, loop) != 0 {
			*(*uintptr)(unsafe.Pointer(innerLoops + uintptr(innerCount)*8)) = loop
			innerCount++
		} else {
			if polygon == libc.UintptrFromInt32(0) {
				v3 = root
			} else {
				v3 = XaddNewLinkedPolygon(tls, polygon)
			}
			polygon = v3
			XaddLinkedLoop(tls, polygon, loop)
			XbboxFromLinkedGeoLoop(tls, loop, bboxes+uintptr(outerCount)*32)
			outerCount++
		}
		// get the next loop and unlink it from this one
		next = (*TLinkedGeoLoop)(unsafe.Pointer(loop)).Fnext
		(*TLinkedGeoLoop)(unsafe.Pointer(loop)).Fnext = libc.UintptrFromInt32(0)
		loop = next
	}
	// Find polygon for each inner loop and assign the hole to it
	i = 0
	for {
		if !(i < innerCount) {
			break
		}
		polygon = _findPolygonForHole(tls, *(*uintptr)(unsafe.Pointer(innerLoops + uintptr(i)*8)), root, bboxes, outerCount)
		if polygon != 0 {
			XaddLinkedLoop(tls, polygon, *(*uintptr)(unsafe.Pointer(innerLoops + uintptr(i)*8)))
		} else {
			// If we can't find a polygon (possible with invalid input), then
			// we need to release the memory for the hole, because the loop has
			// been unlinked from the root and the caller will no longer have
			// a way to destroy it with destroyLinkedMultiPolygon.
			XdestroyLinkedGeoLoop(tls, *(*uintptr)(unsafe.Pointer(innerLoops + uintptr(i)*8)))
			libc.Xfree(tls, *(*uintptr)(unsafe.Pointer(innerLoops + uintptr(i)*8)))
			resultCode = uint32(_E_FAILED)
		}
		goto _4
	_4:
		;
		i++
	}
	// Free allocated memory
	libc.Xfree(tls, innerLoops)
	libc.Xfree(tls, bboxes)
	return resultCode
}

var ___func__17 = [22]uint8{'n', 'o', 'r', 'm', 'a', 'l', 'i', 'z', 'e', 'M', 'u', 'l', 't', 'i', 'P', 'o', 'l', 'y', 'g', 'o', 'n'}

func XpointInsideLinkedGeoLoop(tls *libc.TLS, loop uintptr, bbox uintptr, coord uintptr) (r uint8) {
	var a, b, tmp TLatLng
	var aLng, bLng, lat, lng, ratio, testLng, v1, v4, v5, v6 float64
	var contains, isTransmeridian uint8
	var currentCoord, nextCoord, v2, v3 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a, aLng, b, bLng, contains, currentCoord, isTransmeridian, lat, lng, nextCoord, ratio, testLng, tmp, v1, v2, v3, v4, v5, v6
	if !(XbboxContains(tls, bbox, coord) != 0) {
		return uint8(Dfalse)
	}
	isTransmeridian = XbboxIsTransmeridian(tls, bbox)
	contains = uint8(Dfalse)
	lat = (*TLatLng)(unsafe.Pointer(coord)).Flat
	if isTransmeridian != 0 && (*TLatLng)(unsafe.Pointer(coord)).Flng < libc.Float64FromInt32(0) {
		v1 = (*TLatLng)(unsafe.Pointer(coord)).Flng + libc.Float64FromFloat64(6.283185307179586)
	} else {
		v1 = (*TLatLng)(unsafe.Pointer(coord)).Flng
	}
	lng = v1
	currentCoord = libc.UintptrFromInt32(0)
	nextCoord = libc.UintptrFromInt32(0)
	for int32(Dtrue) != 0 {
		if currentCoord == libc.UintptrFromInt32(0) {
			v2 = (*TLinkedGeoLoop)(unsafe.Pointer(loop)).Ffirst
		} else {
			v2 = (*TLinkedLatLng)(unsafe.Pointer(currentCoord)).Fnext
		}
		currentCoord = v2
		if currentCoord == libc.UintptrFromInt32(0) {
			break
		}
		a = (*TLinkedLatLng)(unsafe.Pointer(currentCoord)).Fvertex
		if (*TLinkedLatLng)(unsafe.Pointer(currentCoord)).Fnext == libc.UintptrFromInt32(0) {
			v3 = (*TLinkedGeoLoop)(unsafe.Pointer(loop)).Ffirst
		} else {
			v3 = (*TLinkedLatLng)(unsafe.Pointer(currentCoord)).Fnext
		}
		nextCoord = v3
		b = (*TLinkedLatLng)(unsafe.Pointer(nextCoord)).Fvertex
		if a.Flat > b.Flat {
			tmp = a
			a = b
			b = tmp
		}
		if lat == a.Flat || lat == b.Flat {
			lat += float64(2.220446049250313e-16)
		}
		if lat < a.Flat || lat > b.Flat {
			continue
		}
		if isTransmeridian != 0 && a.Flng < libc.Float64FromInt32(0) {
			v4 = a.Flng + libc.Float64FromFloat64(6.283185307179586)
		} else {
			v4 = a.Flng
		}
		aLng = v4
		if isTransmeridian != 0 && b.Flng < libc.Float64FromInt32(0) {
			v5 = b.Flng + libc.Float64FromFloat64(6.283185307179586)
		} else {
			v5 = b.Flng
		}
		bLng = v5
		if aLng == lng || bLng == lng {
			lng -= float64(2.220446049250313e-16)
		}
		ratio = (lat - a.Flat) / (b.Flat - a.Flat)
		if isTransmeridian != 0 && aLng+(bLng-aLng)*ratio < libc.Float64FromInt32(0) {
			v6 = aLng + (bLng-aLng)*ratio + libc.Float64FromFloat64(6.283185307179586)
		} else {
			v6 = aLng + (bLng-aLng)*ratio
		}
		testLng = v6
		if testLng > lng {
			contains = libc.BoolUint8(!(contains != 0))
		}
	}
	return contains
}

func XbboxFromLinkedGeoLoop(tls *libc.TLS, loop uintptr, bbox uintptr) {
	var coord, next TLatLng
	var currentCoord, nextCoord, v1, v2 uintptr
	var isTransmeridian uint8
	var lat, lng, maxNegLng, minPosLng float64
	_, _, _, _, _, _, _, _, _, _, _ = coord, currentCoord, isTransmeridian, lat, lng, maxNegLng, minPosLng, next, nextCoord, v1, v2
	if (*TLinkedGeoLoop)(unsafe.Pointer(loop)).Ffirst == libc.UintptrFromInt32(0) {
		*(*TBBox)(unsafe.Pointer(bbox)) = TBBox{}
		return
	}
	(*TBBox)(unsafe.Pointer(bbox)).Fsouth = float64(1.7976931348623157e+308)
	(*TBBox)(unsafe.Pointer(bbox)).Fwest = float64(1.7976931348623157e+308)
	(*TBBox)(unsafe.Pointer(bbox)).Fnorth = -libc.Float64FromFloat64(1.7976931348623157e+308)
	(*TBBox)(unsafe.Pointer(bbox)).Feast = -libc.Float64FromFloat64(1.7976931348623157e+308)
	minPosLng = float64(1.7976931348623157e+308)
	maxNegLng = -libc.Float64FromFloat64(1.7976931348623157e+308)
	isTransmeridian = uint8(Dfalse)
	currentCoord = libc.UintptrFromInt32(0)
	nextCoord = libc.UintptrFromInt32(0)
	for int32(Dtrue) != 0 {
		if currentCoord == libc.UintptrFromInt32(0) {
			v1 = (*TLinkedGeoLoop)(unsafe.Pointer(loop)).Ffirst
		} else {
			v1 = (*TLinkedLatLng)(unsafe.Pointer(currentCoord)).Fnext
		}
		currentCoord = v1
		if currentCoord == libc.UintptrFromInt32(0) {
			break
		}
		coord = (*TLinkedLatLng)(unsafe.Pointer(currentCoord)).Fvertex
		if (*TLinkedLatLng)(unsafe.Pointer(currentCoord)).Fnext == libc.UintptrFromInt32(0) {
			v2 = (*TLinkedGeoLoop)(unsafe.Pointer(loop)).Ffirst
		} else {
			v2 = (*TLinkedLatLng)(unsafe.Pointer(currentCoord)).Fnext
		}
		nextCoord = v2
		next = (*TLinkedLatLng)(unsafe.Pointer(nextCoord)).Fvertex
		lat = coord.Flat
		lng = coord.Flng
		if lat < (*TBBox)(unsafe.Pointer(bbox)).Fsouth {
			(*TBBox)(unsafe.Pointer(bbox)).Fsouth = lat
		}
		if lng < (*TBBox)(unsafe.Pointer(bbox)).Fwest {
			(*TBBox)(unsafe.Pointer(bbox)).Fwest = lng
		}
		if lat > (*TBBox)(unsafe.Pointer(bbox)).Fnorth {
			(*TBBox)(unsafe.Pointer(bbox)).Fnorth = lat
		}
		if lng > (*TBBox)(unsafe.Pointer(bbox)).Feast {
			(*TBBox)(unsafe.Pointer(bbox)).Feast = lng
		}
		if lng > libc.Float64FromInt32(0) && lng < minPosLng {
			minPosLng = lng
		}
		if lng < libc.Float64FromInt32(0) && lng > maxNegLng {
			maxNegLng = lng
		}
		if libc.Xfabs(tls, lng-next.Flng) > float64(3.141592653589793) {
			isTransmeridian = uint8(Dtrue)
		}
	}
	if isTransmeridian != 0 {
		(*TBBox)(unsafe.Pointer(bbox)).Feast = maxNegLng
		(*TBBox)(unsafe.Pointer(bbox)).Fwest = minPosLng
	}
}

func _isClockwiseNormalizedLinkedGeoLoop(tls *libc.TLS, loop uintptr, isTransmeridian uint8) (r uint8) {
	var a, b TLatLng
	var currentCoord, nextCoord, v1, v2 uintptr
	var sum, v3, v4 float64
	_, _, _, _, _, _, _, _, _ = a, b, currentCoord, nextCoord, sum, v1, v2, v3, v4
	sum = libc.Float64FromInt32(0)
	currentCoord = libc.UintptrFromInt32(0)
	nextCoord = libc.UintptrFromInt32(0)
	for int32(Dtrue) != 0 {
		if currentCoord == libc.UintptrFromInt32(0) {
			v1 = (*TLinkedGeoLoop)(unsafe.Pointer(loop)).Ffirst
		} else {
			v1 = (*TLinkedLatLng)(unsafe.Pointer(currentCoord)).Fnext
		}
		currentCoord = v1
		if currentCoord == libc.UintptrFromInt32(0) {
			break
		}
		a = (*TLinkedLatLng)(unsafe.Pointer(currentCoord)).Fvertex
		if (*TLinkedLatLng)(unsafe.Pointer(currentCoord)).Fnext == libc.UintptrFromInt32(0) {
			v2 = (*TLinkedGeoLoop)(unsafe.Pointer(loop)).Ffirst
		} else {
			v2 = (*TLinkedLatLng)(unsafe.Pointer(currentCoord)).Fnext
		}
		nextCoord = v2
		b = (*TLinkedLatLng)(unsafe.Pointer(nextCoord)).Fvertex
		if !(isTransmeridian != 0) && libc.Xfabs(tls, a.Flng-b.Flng) > float64(3.141592653589793) {
			return _isClockwiseNormalizedLinkedGeoLoop(tls, loop, uint8(Dtrue))
		}
		if isTransmeridian != 0 && b.Flng < libc.Float64FromInt32(0) {
			v3 = b.Flng + libc.Float64FromFloat64(6.283185307179586)
		} else {
			v3 = b.Flng
		}
		if isTransmeridian != 0 && a.Flng < libc.Float64FromInt32(0) {
			v4 = a.Flng + libc.Float64FromFloat64(6.283185307179586)
		} else {
			v4 = a.Flng
		}
		sum += (v3 - v4) * (b.Flat + a.Flat)
	}
	return libc.BoolUint8(sum > libc.Float64FromInt32(0))
}

func XisClockwiseLinkedGeoLoop(tls *libc.TLS, loop uintptr) (r uint8) {
	return _isClockwiseNormalizedLinkedGeoLoop(tls, loop, uint8(Dfalse))
}

const DINT32_MAX7 = 2147483647
const DM_2PI4 = 6.28318530717958647692528676655900576839433
const DM_PI4 = 3.14159265358979323846
const DM_PI_21 = 1.57079632679489661923

var _UNIT_VECS8 = [7]TCoordIJK{
	0: {},
	1: {
		Fk: int32(1),
	},
	2: {
		Fj: int32(1),
	},
	3: {
		Fj: int32(1),
		Fk: int32(1),
	},
	4: {
		Fi: int32(1),
	},
	5: {
		Fi: int32(1),
		Fk: int32(1),
	},
	6: {
		Fi: int32(1),
		Fj: int32(1),
	},
}

// C documentation
//
//	/**
//	 * Produces ijk+ coordinates for an index anchored by an origin.
//	 *
//	 * The coordinate space used by this function may have deleted
//	 * regions or warping due to pentagonal distortion.
//	 *
//	 * Coordinates are only comparable if they come from the same
//	 * origin index.
//	 *
//	 * Failure may occur if the index is too far away from the origin
//	 * or if the index is on the other side of a pentagon.
//	 *
//	 * @param origin An anchoring index for the ijk+ coordinate system.
//	 * @param index Index to find the coordinates of
//	 * @param out ijk+ coordinates of the index will be placed here on success
//	 * @return 0 on success, or another value on failure.
//	 */
func XcellToLocalIjk(tls *libc.TLS, origin TH3Index, h3 TH3Index, out uintptr) (r1 TH3Error) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var baseCell, baseCellRotations, directionRotations, i, i1, i2, i3, i4, indexLeadingDigit, indexLeadingDigit1, indexOnPent, originBaseCell, originLeadingDigit, originLeadingDigit1, originOnPent, pentagonRotations, r, res, withinPentagonRotations, v1, v3 int32
	var dir, revDir TDirection
	var v13, v2, v4, v5, v8, v9 bool
	var _ /* indexFijk at bp+0 */ TFaceIJK
	var _ /* offset at bp+16 */ TCoordIJK
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = baseCell, baseCellRotations, dir, directionRotations, i, i1, i2, i3, i4, indexLeadingDigit, indexLeadingDigit1, indexOnPent, originBaseCell, originLeadingDigit, originLeadingDigit1, originOnPent, pentagonRotations, r, res, revDir, withinPentagonRotations, v1, v13, v2, v3, v4, v5, v8, v9
	res = libc.Int32FromUint64(origin & (libc.Uint64FromUint64(15) << libc.Int32FromInt32(DH3_RES_OFFSET)) >> libc.Int32FromInt32(DH3_RES_OFFSET))
	if res != libc.Int32FromUint64(h3&(libc.Uint64FromUint64(15)<<libc.Int32FromInt32(DH3_RES_OFFSET))>>libc.Int32FromInt32(DH3_RES_OFFSET)) {
		return uint32(_E_RES_MISMATCH)
	}
	originBaseCell = libc.Int32FromUint64(origin & (libc.Uint64FromInt32(libc.Int32FromInt32(127)) << libc.Int32FromInt32(DH3_BC_OFFSET)) >> libc.Int32FromInt32(DH3_BC_OFFSET))
	baseCell = libc.Int32FromUint64(h3 & (libc.Uint64FromInt32(libc.Int32FromInt32(127)) << libc.Int32FromInt32(DH3_BC_OFFSET)) >> libc.Int32FromInt32(DH3_BC_OFFSET))
	if originBaseCell < 0 {
		if v2 = libc.Bool(0 != 0); !v2 {
			libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+1093, int32(142), uintptr(unsafe.Pointer(&___func__18)))
		}
		_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
		v1 = libc.Int32FromInt32(1)
	} else {
		v1 = 0
	}
	if v1 != 0 || originBaseCell >= int32(DNUM_BASE_CELLS) {
		// Base cells less than zero can not be represented in an index
		return uint32(_E_CELL_INVALID)
	}
	if baseCell < 0 {
		if v4 = libc.Bool(0 != 0); !v4 {
			libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+1093, int32(146), uintptr(unsafe.Pointer(&___func__18)))
		}
		_ = v4 || libc.Bool(libc.Int32FromInt32(0) != 0)
		v3 = libc.Int32FromInt32(1)
	} else {
		v3 = 0
	}
	if v3 != 0 || baseCell >= int32(DNUM_BASE_CELLS) {
		// Base cells less than zero can not be represented in an index
		return uint32(_E_CELL_INVALID)
	}
	// Direction from origin base cell to index base cell
	dir = int32(_CENTER_DIGIT)
	revDir = int32(_CENTER_DIGIT)
	if originBaseCell != baseCell {
		dir = X_getBaseCellDirection(tls, originBaseCell, baseCell)
		if dir == int32(_INVALID_DIGIT) {
			// Base cells are not neighbors, can't unfold.
			return uint32(_E_FAILED)
		}
		revDir = X_getBaseCellDirection(tls, baseCell, originBaseCell)
		if v5 = revDir != int32(_INVALID_DIGIT); !v5 {
			libc.X__assert_fail(tls, __ccgo_ts+1120, __ccgo_ts+1093, int32(161), uintptr(unsafe.Pointer(&___func__18)))
		}
		_ = v5 || libc.Bool(libc.Int32FromInt32(0) != 0)
	}
	originOnPent = X_isBaseCellPentagon(tls, originBaseCell)
	indexOnPent = X_isBaseCellPentagon(tls, baseCell)
	*(*TFaceIJK)(unsafe.Pointer(bp)) = TFaceIJK{}
	if dir != int32(_CENTER_DIGIT) {
		// Rotate index into the orientation of the origin base cell.
		// cw because we are undoing the rotation into that base cell.
		baseCellRotations = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&XbaseCellNeighbor60CCWRots)) + uintptr(originBaseCell)*28 + uintptr(dir)*4))
		if indexOnPent != 0 {
			i = 0
			for {
				if !(i < baseCellRotations) {
					break
				}
				h3 = X_h3RotatePent60cw(tls, h3)
				revDir = X_rotate60cw(tls, revDir)
				if revDir == int32(_K_AXES_DIGIT) {
					revDir = X_rotate60cw(tls, revDir)
				}
				goto _6
			_6:
				;
				i++
			}
		} else {
			i1 = 0
			for {
				if !(i1 < baseCellRotations) {
					break
				}
				h3 = X_h3Rotate60cw(tls, h3)
				revDir = X_rotate60cw(tls, revDir)
				goto _7
			_7:
				;
				i1++
			}
		}
	}
	// Face is unused. This produces coordinates in base cell coordinate space.
	X_h3ToFaceIjkWithInitializedFijk(tls, h3, bp)
	if dir != int32(_CENTER_DIGIT) {
		if v8 = baseCell != originBaseCell; !v8 {
			libc.X__assert_fail(tls, __ccgo_ts+1144, __ccgo_ts+1093, int32(191), uintptr(unsafe.Pointer(&___func__18)))
		}
		_ = v8 || libc.Bool(libc.Int32FromInt32(0) != 0)
		if v9 = !(originOnPent != 0 && indexOnPent != 0); !v9 {
			libc.X__assert_fail(tls, __ccgo_ts+1171, __ccgo_ts+1093, int32(192), uintptr(unsafe.Pointer(&___func__18)))
		}
		_ = v9 || libc.Bool(libc.Int32FromInt32(0) != 0)
		pentagonRotations = 0
		directionRotations = 0
		if originOnPent != 0 {
			originLeadingDigit = X_h3LeadingNonZeroDigit(tls, origin)
			if originLeadingDigit == int32(_INVALID_DIGIT) {
				return uint32(_E_CELL_INVALID)
			}
			if *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&XFAILED_DIRECTIONS)) + uintptr(originLeadingDigit)*7 + uintptr(dir))) != 0 {
				// TODO: We may be unfolding the pentagon incorrectly in this
				// case; return an error code until this is guaranteed to be
				// correct.
				return uint32(_E_FAILED)
			}
			directionRotations = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&XPENTAGON_ROTATIONS)) + uintptr(originLeadingDigit)*28 + uintptr(dir)*4))
			pentagonRotations = directionRotations
		} else {
			if indexOnPent != 0 {
				indexLeadingDigit = X_h3LeadingNonZeroDigit(tls, h3)
				if indexLeadingDigit == int32(_INVALID_DIGIT) {
					return uint32(_E_CELL_INVALID)
				}
				if *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&XFAILED_DIRECTIONS)) + uintptr(indexLeadingDigit)*7 + uintptr(revDir))) != 0 {
					// TODO: We may be unfolding the pentagon incorrectly in this
					// case; return an error code until this is guaranteed to be
					// correct.
					return uint32(_E_FAILED)
				}
				pentagonRotations = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&XPENTAGON_ROTATIONS)) + uintptr(revDir)*28 + uintptr(indexLeadingDigit)*4))
			}
		}
		if pentagonRotations < 0 || directionRotations < 0 {
			// This occurs when an invalid K axis digit is present
			return uint32(_E_CELL_INVALID)
		}
		i2 = 0
		for {
			if !(i2 < pentagonRotations) {
				break
			}
			X_ijkRotate60cw(tls, bp+4)
			goto _10
		_10:
			;
			i2++
		}
		*(*TCoordIJK)(unsafe.Pointer(bp + 16)) = TCoordIJK{}
		X_neighbor(tls, bp+16, dir)
		// Scale offset based on resolution
		r = res - int32(1)
		for {
			if !(r >= 0) {
				break
			}
			if XisResolutionClassIII(tls, r+int32(1)) != 0 {
				// rotate ccw
				X_downAp7(tls, bp+16)
			} else {
				// rotate cw
				X_downAp7r(tls, bp+16)
			}
			goto _11
		_11:
			;
			r--
		}
		i3 = 0
		for {
			if !(i3 < directionRotations) {
				break
			}
			X_ijkRotate60cw(tls, bp+16)
			goto _12
		_12:
			;
			i3++
		}
		// Perform necessary translation
		X_ijkAdd(tls, bp+4, bp+16, bp+4)
		X_ijkNormalize(tls, bp+4)
	} else {
		if originOnPent != 0 && indexOnPent != 0 {
			// If the origin and index are on pentagon, and we checked that the base
			// cells are the same or neighboring, then they must be the same base
			// cell.
			if v13 = baseCell == originBaseCell; !v13 {
				libc.X__assert_fail(tls, __ccgo_ts+1202, __ccgo_ts+1093, int32(261), uintptr(unsafe.Pointer(&___func__18)))
			}
			_ = v13 || libc.Bool(libc.Int32FromInt32(0) != 0)
			originLeadingDigit1 = X_h3LeadingNonZeroDigit(tls, origin)
			indexLeadingDigit1 = X_h3LeadingNonZeroDigit(tls, h3)
			if originLeadingDigit1 == int32(_INVALID_DIGIT) || indexLeadingDigit1 == int32(_INVALID_DIGIT) {
				return uint32(_E_CELL_INVALID)
			}
			if *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&XFAILED_DIRECTIONS)) + uintptr(originLeadingDigit1)*7 + uintptr(indexLeadingDigit1))) != 0 {
				// TODO: We may be unfolding the pentagon incorrectly in this case;
				// return an error code until this is guaranteed to be correct.
				return uint32(_E_FAILED)
			}
			withinPentagonRotations = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&XPENTAGON_ROTATIONS)) + uintptr(originLeadingDigit1)*28 + uintptr(indexLeadingDigit1)*4))
			i4 = 0
			for {
				if !(i4 < withinPentagonRotations) {
					break
				}
				X_ijkRotate60cw(tls, bp+4)
				goto _14
			_14:
				;
				i4++
			}
		}
	}
	*(*TCoordIJK)(unsafe.Pointer(out)) = (*(*TFaceIJK)(unsafe.Pointer(bp))).Fcoord
	return uint32(_E_SUCCESS)
}

var ___func__18 = [15]uint8{'c', 'e', 'l', 'l', 'T', 'o', 'L', 'o', 'c', 'a', 'l', 'I', 'j', 'k'}

// C documentation
//
//	/**
//	 * Produces an index for ijk+ coordinates anchored by an origin.
//	 *
//	 * The coordinate space used by this function may have deleted
//	 * regions or warping due to pentagonal distortion.
//	 *
//	 * Failure may occur if the coordinates are too far away from the origin
//	 * or if the index is on the other side of a pentagon.
//	 *
//	 * @param origin An anchoring index for the ijk+ coordinate system.
//	 * @param ijk IJK+ Coordinates to find the index of
//	 * @param out The index will be placed here on success
//	 * @return 0 on success, or another value on failure.
//	 */
func XlocalIjkToCell(tls *libc.TLS, origin TH3Index, ijk uintptr, out uintptr) (r1 TH3Error) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var baseCell, baseCellRotations, i, i1, i2, i3, i4, i5, indexLeadingDigit1, indexOnPent, newBaseCell, originBaseCell, originLeadingDigit1, originOnPent, pentagonRotations, r, res, withinPentagonRotations, v1, v11, v13, v4 int32
	var dir, dir1, indexLeadingDigit, originLeadingDigit, revDir TDirection
	var upAp7Error, upAp7rError TH3Error
	var v12, v14, v2, v6, v7, v8, v9 bool
	var _ /* diff at bp+36 */ TCoordIJK
	var _ /* ijkCopy at bp+0 */ TCoordIJK
	var _ /* lastCenter at bp+24 */ TCoordIJK
	var _ /* lastIJK at bp+12 */ TCoordIJK
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = baseCell, baseCellRotations, dir, dir1, i, i1, i2, i3, i4, i5, indexLeadingDigit, indexLeadingDigit1, indexOnPent, newBaseCell, originBaseCell, originLeadingDigit, originLeadingDigit1, originOnPent, pentagonRotations, r, res, revDir, upAp7Error, upAp7rError, withinPentagonRotations, v1, v11, v12, v13, v14, v2, v4, v6, v7, v8, v9
	res = libc.Int32FromUint64(origin & (libc.Uint64FromUint64(15) << libc.Int32FromInt32(DH3_RES_OFFSET)) >> libc.Int32FromInt32(DH3_RES_OFFSET))
	originBaseCell = libc.Int32FromUint64(origin & (libc.Uint64FromInt32(libc.Int32FromInt32(127)) << libc.Int32FromInt32(DH3_BC_OFFSET)) >> libc.Int32FromInt32(DH3_BC_OFFSET))
	if originBaseCell < 0 {
		if v2 = libc.Bool(0 != 0); !v2 {
			libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+1093, int32(305), uintptr(unsafe.Pointer(&___func__19)))
		}
		_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
		v1 = libc.Int32FromInt32(1)
	} else {
		v1 = 0
	}
	if v1 != 0 || originBaseCell >= int32(DNUM_BASE_CELLS) {
		// Base cells less than zero can not be represented in an index
		return uint32(_E_CELL_INVALID)
	}
	originOnPent = X_isBaseCellPentagon(tls, originBaseCell)
	// This logic is very similar to faceIjkToH3
	// initialize the index
	*(*TH3Index)(unsafe.Pointer(out)) = libc.Uint64FromUint64(35184372088831)
	*(*TH3Index)(unsafe.Pointer(out)) = *(*TH3Index)(unsafe.Pointer(out)) & ^(libc.Uint64FromInt32(libc.Int32FromInt32(15))<<libc.Int32FromInt32(DH3_MODE_OFFSET)) | libc.Uint64FromInt32(libc.Int32FromInt32(DH3_CELL_MODE))<<libc.Int32FromInt32(DH3_MODE_OFFSET)
	*(*TH3Index)(unsafe.Pointer(out)) = *(*TH3Index)(unsafe.Pointer(out)) & ^(libc.Uint64FromUint64(15)<<libc.Int32FromInt32(DH3_RES_OFFSET)) | libc.Uint64FromInt32(res)<<libc.Int32FromInt32(DH3_RES_OFFSET)
	// check for res 0/base cell
	if res == 0 {
		dir = X_unitIjkToDigit(tls, ijk)
		if dir == int32(_INVALID_DIGIT) {
			// out of range input - not a unit vector or zero vector
			return uint32(_E_FAILED)
		}
		newBaseCell = X_getBaseCellNeighbor(tls, originBaseCell, dir)
		if newBaseCell == int32(DINVALID_BASE_CELL) {
			// Moving in an invalid direction off a pentagon.
			return uint32(_E_FAILED)
		}
		*(*TH3Index)(unsafe.Pointer(out)) = *(*TH3Index)(unsafe.Pointer(out)) & ^(libc.Uint64FromInt32(libc.Int32FromInt32(127))<<libc.Int32FromInt32(DH3_BC_OFFSET)) | libc.Uint64FromInt32(newBaseCell)<<libc.Int32FromInt32(DH3_BC_OFFSET)
		return uint32(_E_SUCCESS)
	}
	// we need to find the correct base cell offset (if any) for this H3 index;
	// start with the passed in base cell and resolution res ijk coordinates
	// in that base cell's coordinate system
	*(*TCoordIJK)(unsafe.Pointer(bp)) = *(*TCoordIJK)(unsafe.Pointer(ijk))
	// build the H3Index from finest res up
	// adjust r for the fact that the res 0 base cell offsets the indexing
	// digits
	r = res - int32(1)
	for {
		if !(r >= 0) {
			break
		}
		*(*TCoordIJK)(unsafe.Pointer(bp + 12)) = *(*TCoordIJK)(unsafe.Pointer(bp))
		if XisResolutionClassIII(tls, r+int32(1)) != 0 {
			// rotate ccw
			upAp7Error = X_upAp7Checked(tls, bp)
			if upAp7Error != 0 {
				return upAp7Error
			}
			*(*TCoordIJK)(unsafe.Pointer(bp + 24)) = *(*TCoordIJK)(unsafe.Pointer(bp))
			X_downAp7(tls, bp+24)
		} else {
			// rotate cw
			upAp7rError = X_upAp7rChecked(tls, bp)
			if upAp7rError != 0 {
				return upAp7rError
			}
			*(*TCoordIJK)(unsafe.Pointer(bp + 24)) = *(*TCoordIJK)(unsafe.Pointer(bp))
			X_downAp7r(tls, bp+24)
		}
		X_ijkSub(tls, bp+12, bp+24, bp+36)
		X_ijkNormalize(tls, bp+36)
		*(*TH3Index)(unsafe.Pointer(out)) = *(*TH3Index)(unsafe.Pointer(out)) & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<((libc.Int32FromInt32(DMAX_H3_RES)-(r+libc.Int32FromInt32(1)))*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))) | libc.Uint64FromInt32(X_unitIjkToDigit(tls, bp+36))<<((libc.Int32FromInt32(DMAX_H3_RES)-(r+libc.Int32FromInt32(1)))*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))
		goto _3
	_3:
		;
		r--
	}
	// ijkCopy should now hold the IJK of the base cell in the
	// coordinate system of the current base cell
	if (*(*TCoordIJK)(unsafe.Pointer(bp))).Fi > int32(1) || (*(*TCoordIJK)(unsafe.Pointer(bp))).Fj > int32(1) || (*(*TCoordIJK)(unsafe.Pointer(bp))).Fk > int32(1) {
		// out of range input
		return uint32(_E_FAILED)
	}
	// lookup the correct base cell
	dir1 = X_unitIjkToDigit(tls, bp)
	baseCell = X_getBaseCellNeighbor(tls, originBaseCell, dir1)
	if baseCell == int32(DINVALID_BASE_CELL) {
		v4 = 0
	} else {
		v4 = X_isBaseCellPentagon(tls, baseCell)
	}
	// If baseCell is invalid, it must be because the origin base cell is a
	// pentagon, and because pentagon base cells do not border each other,
	// baseCell must not be a pentagon.
	indexOnPent = v4
	if dir1 != int32(_CENTER_DIGIT) {
		// If the index is in a warped direction, we need to unwarp the base
		// cell direction. There may be further need to rotate the index digits.
		pentagonRotations = 0
		if originOnPent != 0 {
			originLeadingDigit = X_h3LeadingNonZeroDigit(tls, origin)
			if originLeadingDigit == int32(_INVALID_DIGIT) {
				return uint32(_E_CELL_INVALID)
			}
			pentagonRotations = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&XPENTAGON_ROTATIONS_REVERSE)) + uintptr(originLeadingDigit)*28 + uintptr(dir1)*4))
			i = 0
			for {
				if !(i < pentagonRotations) {
					break
				}
				dir1 = X_rotate60ccw(tls, dir1)
				goto _5
			_5:
				;
				i++
			}
			// The pentagon rotations are being chosen so that dir is not the
			// deleted direction. If it still happens, it means we're moving
			// into a deleted subsequence, so there is no index here.
			if dir1 == int32(_K_AXES_DIGIT) {
				return uint32(_E_PENTAGON)
			}
			baseCell = X_getBaseCellNeighbor(tls, originBaseCell, dir1)
			// indexOnPent does not need to be checked again since no pentagon
			// base cells border each other.
			if v6 = baseCell != int32(DINVALID_BASE_CELL); !v6 {
				libc.X__assert_fail(tls, __ccgo_ts+1229, __ccgo_ts+1093, int32(411), uintptr(unsafe.Pointer(&___func__19)))
			}
			_ = v6 || libc.Bool(libc.Int32FromInt32(0) != 0)
			if v7 = !(X_isBaseCellPentagon(tls, baseCell) != 0); !v7 {
				libc.X__assert_fail(tls, __ccgo_ts+1259, __ccgo_ts+1093, int32(412), uintptr(unsafe.Pointer(&___func__19)))
			}
			_ = v7 || libc.Bool(libc.Int32FromInt32(0) != 0)
		}
		// Now we can determine the relation between the origin and target base
		// cell.
		baseCellRotations = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&XbaseCellNeighbor60CCWRots)) + uintptr(originBaseCell)*28 + uintptr(dir1)*4))
		if v8 = baseCellRotations >= 0; !v8 {
			libc.X__assert_fail(tls, __ccgo_ts+1290, __ccgo_ts+1093, int32(419), uintptr(unsafe.Pointer(&___func__19)))
		}
		_ = v8 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// Adjust for pentagon warping within the base cell. The base cell
		// should be in the right location, so now we need to rotate the index
		// back. We might not need to check for errors since we would just be
		// double mapping.
		if indexOnPent != 0 {
			revDir = X_getBaseCellDirection(tls, baseCell, originBaseCell)
			if v9 = revDir != int32(_INVALID_DIGIT); !v9 {
				libc.X__assert_fail(tls, __ccgo_ts+1120, __ccgo_ts+1093, int32(428), uintptr(unsafe.Pointer(&___func__19)))
			}
			_ = v9 || libc.Bool(libc.Int32FromInt32(0) != 0)
			// Adjust for the different coordinate space in the two base cells.
			// This is done first because we need to do the pentagon rotations
			// based on the leading digit in the pentagon's coordinate system.
			i1 = 0
			for {
				if !(i1 < baseCellRotations) {
					break
				}
				*(*TH3Index)(unsafe.Pointer(out)) = X_h3Rotate60ccw(tls, *(*TH3Index)(unsafe.Pointer(out)))
				goto _10
			_10:
				;
				i1++
			}
			indexLeadingDigit = X_h3LeadingNonZeroDigit(tls, *(*TH3Index)(unsafe.Pointer(out)))
			// This case should be unreachable because this function is building
			// *out, and should never generate an invalid digit, above.
			if indexLeadingDigit == int32(_INVALID_DIGIT) {
				if v12 = libc.Bool(0 != 0); !v12 {
					libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+1093, int32(440), uintptr(unsafe.Pointer(&___func__19)))
				}
				_ = v12 || libc.Bool(libc.Int32FromInt32(0) != 0)
				v11 = libc.Int32FromInt32(1)
			} else {
				v11 = 0
			}
			if v11 != 0 {
				return uint32(_E_CELL_INVALID)
			}
			if X_isBaseCellPolarPentagon(tls, baseCell) != 0 {
				pentagonRotations = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&XPENTAGON_ROTATIONS_REVERSE_POLAR)) + uintptr(revDir)*28 + uintptr(indexLeadingDigit)*4))
			} else {
				pentagonRotations = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&XPENTAGON_ROTATIONS_REVERSE_NONPOLAR)) + uintptr(revDir)*28 + uintptr(indexLeadingDigit)*4))
			}
			// For this to occur, revDir would need to be 1. Since revDir is
			// from the index base cell (which is a pentagon) towards the
			// origin, this should never be the case.
			if pentagonRotations < 0 {
				if v14 = libc.Bool(0 != 0); !v14 {
					libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+1093, int32(454), uintptr(unsafe.Pointer(&___func__19)))
				}
				_ = v14 || libc.Bool(libc.Int32FromInt32(0) != 0)
				v13 = libc.Int32FromInt32(1)
			} else {
				v13 = 0
			}
			if v13 != 0 {
				return uint32(_E_CELL_INVALID)
			}
			i2 = 0
			for {
				if !(i2 < pentagonRotations) {
					break
				}
				*(*TH3Index)(unsafe.Pointer(out)) = X_h3RotatePent60ccw(tls, *(*TH3Index)(unsafe.Pointer(out)))
				goto _15
			_15:
				;
				i2++
			}
		} else {
			if pentagonRotations < 0 {
				return uint32(_E_CELL_INVALID)
			}
			i3 = 0
			for {
				if !(i3 < pentagonRotations) {
					break
				}
				*(*TH3Index)(unsafe.Pointer(out)) = X_h3Rotate60ccw(tls, *(*TH3Index)(unsafe.Pointer(out)))
				goto _16
			_16:
				;
				i3++
			}
			// Adjust for the different coordinate space in the two base cells.
			i4 = 0
			for {
				if !(i4 < baseCellRotations) {
					break
				}
				*(*TH3Index)(unsafe.Pointer(out)) = X_h3Rotate60ccw(tls, *(*TH3Index)(unsafe.Pointer(out)))
				goto _17
			_17:
				;
				i4++
			}
		}
	} else {
		if originOnPent != 0 && indexOnPent != 0 {
			originLeadingDigit1 = X_h3LeadingNonZeroDigit(tls, origin)
			indexLeadingDigit1 = X_h3LeadingNonZeroDigit(tls, *(*TH3Index)(unsafe.Pointer(out)))
			if originLeadingDigit1 == int32(_INVALID_DIGIT) || indexLeadingDigit1 == int32(_INVALID_DIGIT) {
				return uint32(_E_CELL_INVALID)
			}
			withinPentagonRotations = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&XPENTAGON_ROTATIONS_REVERSE)) + uintptr(originLeadingDigit1)*28 + uintptr(indexLeadingDigit1)*4))
			if withinPentagonRotations < 0 {
				// This occurs when an invalid K axis digit is present
				return uint32(_E_CELL_INVALID)
			}
			i5 = 0
			for {
				if !(i5 < withinPentagonRotations) {
					break
				}
				*(*TH3Index)(unsafe.Pointer(out)) = X_h3Rotate60ccw(tls, *(*TH3Index)(unsafe.Pointer(out)))
				goto _18
			_18:
				;
				i5++
			}
		}
	}
	if indexOnPent != 0 {
		// TODO: There are cases in cellToLocalIjk which are failed but not
		// accounted for here - instead just fail if the recovered index is
		// invalid.
		if X_h3LeadingNonZeroDigit(tls, *(*TH3Index)(unsafe.Pointer(out))) == int32(_K_AXES_DIGIT) {
			return uint32(_E_PENTAGON)
		}
	}
	*(*TH3Index)(unsafe.Pointer(out)) = *(*TH3Index)(unsafe.Pointer(out)) & ^(libc.Uint64FromInt32(libc.Int32FromInt32(127))<<libc.Int32FromInt32(DH3_BC_OFFSET)) | libc.Uint64FromInt32(baseCell)<<libc.Int32FromInt32(DH3_BC_OFFSET)
	return uint32(_E_SUCCESS)
}

var ___func__19 = [15]uint8{'l', 'o', 'c', 'a', 'l', 'I', 'j', 'k', 'T', 'o', 'C', 'e', 'l', 'l'}

// C documentation
//
//	/**
//	 * Produces ij coordinates for an index anchored by an origin.
//	 *
//	 * The coordinate space used by this function may have deleted
//	 * regions or warping due to pentagonal distortion.
//	 *
//	 * Coordinates are only comparable if they come from the same
//	 * origin index.
//	 *
//	 * Failure may occur if the index is too far away from the origin
//	 * or if the index is on the other side of a pentagon.
//	 *
//	 * This function's output is not guaranteed
//	 * to be compatible across different versions of H3.
//	 *
//	 * @param origin An anchoring index for the ij coordinate system.
//	 * @param index Index to find the coordinates of
//	 * @param mode Mode, must be 0
//	 * @param out ij coordinates of the index will be placed here on success
//	 * @return 0 on success, or another value on failure.
//	 */
func XcellToLocalIj(tls *libc.TLS, origin TH3Index, index TH3Index, mode Tuint32_t, out uintptr) (r TH3Error) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var failed TH3Error
	var _ /* ijk at bp+0 */ TCoordIJK
	_ = failed
	if mode != uint32(0) {
		return uint32(_E_OPTION_INVALID)
	}
	failed = XcellToLocalIjk(tls, origin, index, bp)
	if failed != 0 {
		return failed
	}
	XijkToIj(tls, bp, out)
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * Produces an index for ij coordinates anchored by an origin.
//	 *
//	 * The coordinate space used by this function may have deleted
//	 * regions or warping due to pentagonal distortion.
//	 *
//	 * Failure may occur if the index is too far away from the origin
//	 * or if the index is on the other side of a pentagon.
//	 *
//	 * This function's output is not guaranteed
//	 * to be compatible across different versions of H3.
//	 *
//	 * @param origin An anchoring index for the ij coordinate system.
//	 * @param ij ij coordinates to index.
//	 * @param mode Mode, must be 0
//	 * @param out Index will be placed here on success.
//	 * @return 0 on success, or another value on failure.
//	 */
func XlocalIjToCell(tls *libc.TLS, origin TH3Index, ij uintptr, mode Tuint32_t, out uintptr) (r TH3Error) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var ijToIjkError TH3Error
	var _ /* ijk at bp+0 */ TCoordIJK
	_ = ijToIjkError
	if mode != uint32(0) {
		return uint32(_E_OPTION_INVALID)
	}
	ijToIjkError = XijToIjk(tls, ij, bp)
	if ijToIjkError != 0 {
		return ijToIjkError
	}
	return XlocalIjkToCell(tls, origin, bp, out)
}

// C documentation
//
//	/**
//	 * Produces the grid distance between the two indexes.
//	 *
//	 * This function may fail to find the distance between two indexes, for
//	 * example if they are very far apart. It may also fail when finding
//	 * distances for indexes on opposite sides of a pentagon.
//	 *
//	 * @param origin Index to find the distance from.
//	 * @param index Index to find the distance to.
//	 * @param out The distance in cells
//	 * @returns E_SUCCESS on success, or another value if the library cannot compute
//	 * the distance.
//	 */
func XgridDistance(tls *libc.TLS, origin TH3Index, index TH3Index, out uintptr) (r TH3Error) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var destError, originError TH3Error
	var _ /* h3Ijk at bp+12 */ TCoordIJK
	var _ /* originIjk at bp+0 */ TCoordIJK
	_, _ = destError, originError
	originError = XcellToLocalIjk(tls, origin, origin, bp)
	if originError != 0 {
		return originError
	}
	destError = XcellToLocalIjk(tls, origin, index, bp+12)
	if destError != 0 {
		return destError
	}
	*(*Tint64_t)(unsafe.Pointer(out)) = int64(XijkDistance(tls, bp, bp+12))
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * Number of indexes in a line from the start index to the end index,
//	 * to be used for allocating memory. Returns a negative number if the
//	 * line cannot be computed.
//	 *
//	 * @param start Start index of the line
//	 * @param end End index of the line
//	 * @param size Size of the line
//	 * @returns 0 on success, or another value on error
//	 */
func XgridPathCellsSize(tls *libc.TLS, start TH3Index, end TH3Index, size uintptr) (r TH3Error) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var distanceError TH3Error
	var _ /* distance at bp+0 */ Tint64_t
	_ = distanceError
	distanceError = XgridDistance(tls, start, end, bp)
	if distanceError != 0 {
		return distanceError
	}
	*(*Tint64_t)(unsafe.Pointer(size)) = *(*Tint64_t)(unsafe.Pointer(bp)) + int64(1)
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * Given cube coords as doubles, round to valid integer coordinates. Algorithm
//	 * from https://www.redblobgames.com/grids/hexagons/#rounding
//	 * @param i   Floating-point I coord
//	 * @param j   Floating-point J coord
//	 * @param k   Floating-point K coord
//	 * @param ijk IJK coord struct, modified in place
//	 */
func _cubeRound(tls *libc.TLS, i float64, j float64, k float64, ijk uintptr) {
	var iDiff, jDiff, kDiff float64
	var ri, rj, rk int32
	_, _, _, _, _, _ = iDiff, jDiff, kDiff, ri, rj, rk
	ri = int32(libc.X__builtin_round(tls, i))
	rj = int32(libc.X__builtin_round(tls, j))
	rk = int32(libc.X__builtin_round(tls, k))
	iDiff = libc.Xfabs(tls, float64(ri)-i)
	jDiff = libc.Xfabs(tls, float64(rj)-j)
	kDiff = libc.Xfabs(tls, float64(rk)-k)
	// Round, maintaining valid cube coords
	if iDiff > jDiff && iDiff > kDiff {
		ri = -rj - rk
	} else {
		if jDiff > kDiff {
			rj = -ri - rk
		} else {
			rk = -ri - rj
		}
	}
	(*TCoordIJK)(unsafe.Pointer(ijk)).Fi = ri
	(*TCoordIJK)(unsafe.Pointer(ijk)).Fj = rj
	(*TCoordIJK)(unsafe.Pointer(ijk)).Fk = rk
}

// C documentation
//
//	/**
//	 * Given two H3 indexes, return the line of indexes between them (inclusive).
//	 *
//	 * This function may fail to find the line between two indexes, for
//	 * example if they are very far apart. It may also fail when finding
//	 * distances for indexes on opposite sides of a pentagon.
//	 *
//	 * Notes:
//	 *
//	 *  - The specific output of this function should not be considered stable
//	 *    across library versions. The only guarantees the library provides are
//	 *    that the line length will be `gridDistance(start, end) + 1` and that
//	 *    every index in the line will be a neighbor of the preceding index.
//	 *  - Lines are drawn in grid space, and may not correspond exactly to either
//	 *    Cartesian lines or great arcs.
//	 *
//	 * @param start Start index of the line
//	 * @param end End index of the line
//	 * @param out Output array, which must be of size gridPathCellsSize(start, end)
//	 * @return 0 on success, or another value on failure.
//	 */
func XgridPathCells(tls *libc.TLS, start TH3Index, end TH3Index, out uintptr) (r TH3Error) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var currentError, distanceError, endError, startError TH3Error
	var iStep, invDistance, jStep, kStep, v5 float64
	var n Tint64_t
	var v1, v3 int32
	var v2, v4 bool
	var _ /* currentIjk at bp+32 */ TCoordIJK
	var _ /* distance at bp+0 */ Tint64_t
	var _ /* endIjk at bp+20 */ TCoordIJK
	var _ /* startIjk at bp+8 */ TCoordIJK
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = currentError, distanceError, endError, iStep, invDistance, jStep, kStep, n, startError, v1, v2, v3, v4, v5
	distanceError = XgridDistance(tls, start, end, bp)
	// Early exit if we can't calculate the line
	if distanceError != 0 {
		return distanceError
	}
	// Get IJK coords for the start and end. We've already confirmed
	// that these can be calculated with the distance check above.
	*(*TCoordIJK)(unsafe.Pointer(bp + 8)) = TCoordIJK{}
	*(*TCoordIJK)(unsafe.Pointer(bp + 20)) = TCoordIJK{}
	// Convert H3 addresses to IJK coords
	startError = XcellToLocalIjk(tls, start, start, bp+8)
	if startError != 0 {
		if v2 = libc.Bool(0 != 0); !v2 {
			libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+1093, int32(692), uintptr(unsafe.Pointer(&___func__20)))
		}
		_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
		v1 = libc.Int32FromInt32(1)
	} else {
		v1 = 0
	}
	if v1 != 0 {
		// Unreachable because this was called as part of gridDistance
		return startError
	}
	endError = XcellToLocalIjk(tls, start, end, bp+20)
	if endError != 0 {
		if v4 = libc.Bool(0 != 0); !v4 {
			libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+1093, int32(697), uintptr(unsafe.Pointer(&___func__20)))
		}
		_ = v4 || libc.Bool(libc.Int32FromInt32(0) != 0)
		v3 = libc.Int32FromInt32(1)
	} else {
		v3 = 0
	}
	if v3 != 0 {
		// Unreachable because this was called as part of gridDistance
		return endError
	}
	// Convert IJK to cube coordinates suitable for linear interpolation
	XijkToCube(tls, bp+8)
	XijkToCube(tls, bp+20)
	if *(*Tint64_t)(unsafe.Pointer(bp)) != 0 {
		v5 = float64(1) / float64(*(*Tint64_t)(unsafe.Pointer(bp)))
	} else {
		v5 = libc.Float64FromInt32(0)
	}
	invDistance = v5
	iStep = float64((*(*TCoordIJK)(unsafe.Pointer(bp + 20))).Fi-(*(*TCoordIJK)(unsafe.Pointer(bp + 8))).Fi) * invDistance
	jStep = float64((*(*TCoordIJK)(unsafe.Pointer(bp + 20))).Fj-(*(*TCoordIJK)(unsafe.Pointer(bp + 8))).Fj) * invDistance
	kStep = float64((*(*TCoordIJK)(unsafe.Pointer(bp + 20))).Fk-(*(*TCoordIJK)(unsafe.Pointer(bp + 8))).Fk) * invDistance
	*(*TCoordIJK)(unsafe.Pointer(bp + 32)) = TCoordIJK{
		Fi: (*(*TCoordIJK)(unsafe.Pointer(bp + 8))).Fi,
		Fj: (*(*TCoordIJK)(unsafe.Pointer(bp + 8))).Fj,
		Fk: (*(*TCoordIJK)(unsafe.Pointer(bp + 8))).Fk,
	}
	n = 0
	for {
		if !(n <= *(*Tint64_t)(unsafe.Pointer(bp))) {
			break
		}
		_cubeRound(tls, float64((*(*TCoordIJK)(unsafe.Pointer(bp + 8))).Fi)+iStep*float64(n), float64((*(*TCoordIJK)(unsafe.Pointer(bp + 8))).Fj)+jStep*float64(n), float64((*(*TCoordIJK)(unsafe.Pointer(bp + 8))).Fk)+kStep*float64(n), bp+32)
		// Convert cube -> ijk -> h3 index
		XcubeToIjk(tls, bp+32)
		currentError = XlocalIjkToCell(tls, start, bp+32, out+uintptr(n)*8)
		if currentError != 0 {
			// The cells between `start` and `end` may fall in pentagon
			// distortion.
			return currentError
		}
		goto _6
	_6:
		;
		n++
	}
	return uint32(_E_SUCCESS)
}

var ___func__20 = [14]uint8{'g', 'r', 'i', 'd', 'P', 'a', 't', 'h', 'C', 'e', 'l', 'l', 's'}

// C documentation
//
//	/**
//	 * _ipow does integer exponentiation efficiently. Taken from StackOverflow.
//	 *
//	 * @param base the integer base (can be positive or negative)
//	 * @param exp the integer exponent (should be nonnegative)
//	 *
//	 * @return the exponentiated value
//	 */
func X_ipow(tls *libc.TLS, base Tint64_t, exp Tint64_t) (r Tint64_t) {
	var result Tint64_t
	_ = result
	result = int64(1)
	for exp != 0 {
		if exp&int64(1) != 0 {
			result *= base
		}
		exp >>= int64(1)
		base *= base
	}
	return result
}

const DCELL_SCALE_FACTOR = 1.1
const DCHILD_SCALE_FACTOR = 1.4
const DINT32_MAX8 = 0x7fffffff
const DM_PI5 = 3.141592653589793

type TIterCellsPolygonCompact = struct {
	Fcell     TH3Index
	Ferror1   TH3Error
	F_res     int32
	F_flags   Tuint32_t
	F_polygon uintptr
	F_bboxes  uintptr
	F_started uint8
}

type TIterCellsPolygon = struct {
	Fcell       TH3Index
	Ferror1     TH3Error
	F_cellIter  TIterCellsPolygonCompact
	F_childIter TIterCellsChildren
}

var _UNIT_VECS9 = [7]TCoordIJK{
	0: {},
	1: {
		Fk: int32(1),
	},
	2: {
		Fj: int32(1),
	},
	3: {
		Fj: int32(1),
		Fk: int32(1),
	},
	4: {
		Fi: int32(1),
	},
	5: {
		Fi: int32(1),
		Fk: int32(1),
	},
	6: {
		Fi: int32(1),
		Fj: int32(1),
	},
}

// Factor by which to scale the cell bounding box to include all cells.
// This was determined empirically by finding the smallest factor that
// passed exhaustive tests.

// Factor by which to scale the cell bounding box to include all children.
// This was determined empirically by finding the smallest factor that
// passed exhaustive tests.

// C documentation
//
//	/**
//	 * Max cell edge length, in radians, for each resolution. This was computed
//	 * by taking the max exact edge length for cells at the center of each base
//	 * cell at that resolution.
//	 */
var _MAX_EDGE_LENGTH_RADS = [16]float64{
	0:  float64(0.2157720626513),
	1:  float64(0.08308767068495),
	2:  float64(0.03148970436439),
	3:  float64(0.01190662871439),
	4:  float64(0.00450053330908),
	5:  float64(0.00170105523619),
	6:  float64(0.00064293917678),
	7:  float64(0.00024300820659),
	8:  float64(9.184847087e-05),
	9:  float64(3.471545901e-05),
	10: float64(1.312121017e-05),
	11: float64(4.95935129e-06),
	12: float64(1.8744586e-06),
	13: float64(7.0847876e-07),
	14: float64(2.677798e-07),
	15: float64(1.0121125e-07),
}

// C documentation
//
//	/** All cells that contain the north pole, by res */
var _NORTH_POLE_CELLS = [16]TH3Index{
	0:  uint64(0x8001fffffffffff),
	1:  uint64(0x81033ffffffffff),
	2:  uint64(0x820327fffffffff),
	3:  uint64(0x830326fffffffff),
	4:  uint64(0x8403263ffffffff),
	5:  uint64(0x85032623fffffff),
	6:  uint64(0x860326237ffffff),
	7:  uint64(0x870326233ffffff),
	8:  uint64(0x880326233bfffff),
	9:  uint64(0x890326233abffff),
	10: uint64(0x8a0326233ab7fff),
	11: uint64(0x8b0326233ab0fff),
	12: uint64(0x8c0326233ab03ff),
	13: uint64(0x8d0326233ab03bf),
	14: uint64(0x8e0326233ab039f),
	15: uint64(0x8f0326233ab0399),
}

// C documentation
//
//	/** All cells that contain the south pole, by res */
var _SOUTH_POLE_CELLS = [16]TH3Index{
	0:  uint64(0x80f3fffffffffff),
	1:  uint64(0x81f2bffffffffff),
	2:  uint64(0x82f297fffffffff),
	3:  uint64(0x83f293fffffffff),
	4:  uint64(0x84f2939ffffffff),
	5:  uint64(0x85f29383fffffff),
	6:  uint64(0x86f29380fffffff),
	7:  uint64(0x87f29380effffff),
	8:  uint64(0x88f29380e1fffff),
	9:  uint64(0x89f29380e0fffff),
	10: uint64(0x8af29380e0d7fff),
	11: uint64(0x8bf29380e0d0fff),
	12: uint64(0x8cf29380e0d0dff),
	13: uint64(0x8df29380e0d0cff),
	14: uint64(0x8ef29380e0d0cc7),
	15: uint64(0x8ff29380e0d0cc4),
}

// C documentation
//
//	/** Pre-calculated bounding boxes for all res 0 cells */
var _RES0_BBOXES = [122]TBBox{
	0: {
		Fnorth: float64(1.52480158339146),
		Fsouth: float64(1.20305471830087),
		Feast:  -libc.Float64FromFloat64(0.60664883654036),
		Fwest:  float64(0.00568297271999),
	},
	1: {
		Fnorth: float64(1.52480158339146),
		Fsouth: float64(1.17872424267511),
		Feast:  -libc.Float64FromFloat64(0.60664883654036),
		Fwest:  float64(2.54046980298264),
	},
	2: {
		Fnorth: float64(1.52480158339146),
		Fsouth: float64(1.09069387298096),
		Feast:  -libc.Float64FromFloat64(2.85286053297673),
		Fwest:  float64(1.64310689027893),
	},
	3: {
		Fnorth: float64(1.41845302535151),
		Fsouth: float64(1.01285145697208),
		Feast:  float64(0.00568297271999),
		Fwest:  -libc.Float64FromFloat64(1.16770379632602),
	},
	4: {
		Fnorth: float64(1.27950477868453),
		Fsouth: float64(0.97226652536306),
		Feast:  float64(0.55556064983494),
		Fwest:  -libc.Float64FromFloat64(0.18229924845326),
	},
	5: {
		Fnorth: float64(1.32929586572429),
		Fsouth: float64(0.91898920750071),
		Feast:  float64(2.05622344943192),
		Fwest:  float64(1.08813154278274),
	},
	6: {
		Fnorth: float64(1.32899086063916),
		Fsouth: float64(0.9427181537636),
		Feast:  -libc.Float64FromFloat64(2.29875289606378),
		Fwest:  float64(3.01700008041993),
	},
	7: {
		Fnorth: float64(1.26020983864103),
		Fsouth: float64(0.84291228415618),
		Feast:  -libc.Float64FromFloat64(0.89971867664861),
		Fwest:  -libc.Float64FromFloat64(1.75967359310997),
	},
	8: {
		Fnorth: float64(1.21114673854945),
		Fsouth: float64(0.86170600921069),
		Feast:  float64(1.19129757609455),
		Fwest:  float64(0.43777608996454),
	},
	9: {
		Fnorth: float64(1.21075831414294),
		Fsouth: float64(0.83795331049498),
		Feast:  -libc.Float64FromFloat64(1.72022875779891),
		Fwest:  -libc.Float64FromFloat64(2.43793861727138),
	},
	10: {
		Fnorth: float64(1.15546530929588),
		Fsouth: float64(0.78982455384253),
		Feast:  float64(2.53659412229266),
		Fwest:  float64(1.85709133451243),
	},
	11: {
		Fnorth: float64(1.15528445067052),
		Fsouth: float64(0.76641428724335),
		Feast:  -libc.Float64FromFloat64(3.06738507202411),
		Fwest:  float64(2.53646110244042),
	},
	12: {
		Fnorth: float64(1.10121643537669),
		Fsouth: float64(0.71330093663066),
		Feast:  float64(0.09640581900154),
		Fwest:  -libc.Float64FromFloat64(0.52154514518248),
	},
	13: {
		Fnorth: float64(1.07042472765165),
		Fsouth: float64(0.67603948819406),
		Feast:  -libc.Float64FromFloat64(0.47984202840088),
		Fwest:  -libc.Float64FromFloat64(1.1030615960309),
	},
	14: {
		Fnorth: float64(1.0327022874896),
		Fsouth: float64(0.72356358827215),
		Feast:  -libc.Float64FromFloat64(2.24990138725146),
		Fwest:  -libc.Float64FromFloat64(2.74510220919157),
	},
	15: {
		Fnorth: float64(1.01929924623886),
		Fsouth: float64(0.65491232835426),
		Feast:  float64(0.63035574240731),
		Fwest:  float64(0.0353703009647),
	},
	16: {
		Fnorth: float64(1.01786037568858),
		Fsouth: float64(0.58827636737638),
		Feast:  float64(1.53192721817065),
		Fwest:  float64(0.93672682511233),
	},
	17: {
		Fnorth: float64(0.9808143413602),
		Fsouth: float64(0.61076063532947),
		Feast:  -libc.Float64FromFloat64(2.67100636598529),
		Fwest:  float64(3.06516463008733),
	},
	18: {
		Fnorth: float64(0.98106023192774),
		Fsouth: float64(0.5867983657157),
		Feast:  float64(2.02829766214461),
		Fwest:  float64(1.5133437497028),
	},
	19: {
		Fnorth: float64(0.96374551790056),
		Fsouth: float64(0.55186491737474),
		Feast:  -libc.Float64FromFloat64(1.42976721313659),
		Fwest:  -libc.Float64FromFloat64(1.96852202530104),
	},
	20: {
		Fnorth: float64(0.87536136210723),
		Fsouth: float64(0.50008952762292),
		Feast:  -libc.Float64FromFloat64(1.9243561357143),
		Fwest:  -libc.Float64FromFloat64(2.41641343219793),
	},
	21: {
		Fnorth: float64(0.88611243445554),
		Fsouth: float64(0.52742963716774),
		Feast:  -libc.Float64FromFloat64(0.95781946324194),
		Fwest:  -libc.Float64FromFloat64(1.4762896630593),
	},
	22: {
		Fnorth: float64(0.86881343251986),
		Fsouth: float64(0.50770567021439),
		Feast:  float64(1.03236795495839),
		Fwest:  float64(0.50347284027426),
	},
	23: {
		Fnorth: float64(0.89235638181782),
		Fsouth: float64(0.48781264892508),
		Feast:  float64(2.7643030211915),
		Fwest:  float64(2.29989716697031),
	},
	24: {
		Fnorth: float64(0.82570569254601),
		Fsouth: float64(0.52173101741059),
		Feast:  float64(2.30921681461428),
		Fwest:  float64(1.9319854182898),
	},
	25: {
		Fnorth: float64(0.80599330438546),
		Fsouth: float64(0.40150819579319),
		Feast:  -libc.Float64FromFloat64(3.0641755940324),
		Fwest:  float64(2.70079300784409),
	},
	26: {
		Fnorth: float64(0.81612079704781),
		Fsouth: float64(0.38396800633226),
		Feast:  -libc.Float64FromFloat64(0.21614378891839),
		Fwest:  -libc.Float64FromFloat64(0.70420149722178),
	},
	27: {
		Fnorth: float64(0.75822779851431),
		Fsouth: float64(0.39943555383751),
		Feast:  -libc.Float64FromFloat64(2.34059978084699),
		Fwest:  -libc.Float64FromFloat64(2.82127373822444),
	},
	28: {
		Fnorth: float64(0.78861390967531),
		Fsouth: float64(0.38742018303868),
		Feast:  float64(0.23115687731652),
		Fwest:  -libc.Float64FromFloat64(0.22599491086066),
	},
	29: {
		Fnorth: float64(0.71515840341957),
		Fsouth: float64(0.33012478438475),
		Feast:  -libc.Float64FromFloat64(0.64847976163163),
		Fwest:  -libc.Float64FromFloat64(1.08249728121219),
	},
	30: {
		Fnorth: float64(0.70359051048414),
		Fsouth: float64(0.29148673180722),
		Feast:  float64(1.71441081857246),
		Fwest:  float64(1.28443348381696),
	},
	31: {
		Fnorth: float64(0.69190629544818),
		Fsouth: float64(0.28808313184381),
		Feast:  float64(0.64863909244647),
		Fwest:  float64(0.16372369282557),
	},
	32: {
		Fnorth: float64(0.64863235654749),
		Fsouth: float64(0.26290420067147),
		Feast:  float64(2.10318098268379),
		Fwest:  float64(1.69556122548344),
	},
	33: {
		Fnorth: float64(0.65722892279906),
		Fsouth: float64(0.28222653310929),
		Feast:  float64(1.30918693285466),
		Fwest:  float64(0.87594416271685),
	},
	34: {
		Fnorth: float64(0.64750997738584),
		Fsouth: float64(0.2414986570985),
		Feast:  -libc.Float64FromFloat64(1.30272192474556),
		Fwest:  -libc.Float64FromFloat64(1.68708570163242),
	},
	35: {
		Fnorth: float64(0.62380174028378),
		Fsouth: float64(0.25522080363509),
		Feast:  -libc.Float64FromFloat64(2.72428423026826),
		Fwest:  float64(3.1040147323763),
	},
	36: {
		Fnorth: float64(0.64228460410023),
		Fsouth: float64(0.21206753429148),
		Feast:  -libc.Float64FromFloat64(1.67639240992071),
		Fwest:  -libc.Float64FromFloat64(2.11772366767341),
	},
	37: {
		Fnorth: float64(0.59919175361146),
		Fsouth: float64(0.2162046083657),
		Feast:  float64(2.4859286838769),
		Fwest:  float64(2.07350353893591),
	},
	38: {
		Fnorth: float64(0.55637406851384),
		Fsouth: float64(0.2527655743723),
		Feast:  -libc.Float64FromFloat64(0.99885388505694),
		Fwest:  -libc.Float64FromFloat64(1.32642489358939),
	},
	39: {
		Fnorth: float64(0.55648013300665),
		Fsouth: float64(0.15187401321019),
		Feast:  float64(2.87032088421324),
		Fwest:  float64(2.44642320475367),
	},
	40: {
		Fnorth: float64(0.5460368797045),
		Fsouth: float64(0.15589091511369),
		Feast:  -libc.Float64FromFloat64(2.0678986606706),
		Fwest:  -libc.Float64FromFloat64(2.49091419631961),
	},
	41: {
		Fnorth: float64(0.51206347752697),
		Fsouth: float64(0.15522020377124),
		Feast:  float64(0.95446767315996),
		Fwest:  float64(0.54443262110414),
	},
	42: {
		Fnorth: float64(0.49767951537101),
		Fsouth: float64(0.10944898890579),
		Feast:  -libc.Float64FromFloat64(0.04335162263358),
		Fwest:  -libc.Float64FromFloat64(0.42900268178569),
	},
	43: {
		Fnorth: float64(0.46538045483671),
		Fsouth: float64(0.0602996863772),
		Feast:  -libc.Float64FromFloat64(0.41240613713421),
		Fwest:  -libc.Float64FromFloat64(0.80603623808166),
	},
	44: {
		Fnorth: float64(0.44686891066946),
		Fsouth: float64(0.06926857458503),
		Feast:  float64(0.32053284794952),
		Fwest:  -libc.Float64FromFloat64(0.07005748900849),
	},
	45: {
		Fnorth: float64(0.43208958202064),
		Fsouth: float64(0.0779644093814),
		Feast:  -libc.Float64FromFloat64(3.0623245307966),
		Fwest:  float64(2.80602499990282),
	},
	46: {
		Fnorth: float64(0.43103892586713),
		Fsouth: float64(0.02927431919853),
		Feast:  -libc.Float64FromFloat64(2.41589238618422),
		Fwest:  -libc.Float64FromFloat64(2.85735809951951),
	},
	47: {
		Fnorth: float64(0.38073727558986),
		Fsouth: -libc.Float64FromFloat64(0.00297016159959),
		Feast:  -libc.Float64FromFloat64(0.77039553861218),
		Fwest:  -libc.Float64FromFloat64(1.14788248745028),
	},
	48: {
		Fnorth: float64(0.39113816687141),
		Fsouth: -libc.Float64FromFloat64(0.01518764903038),
		Feast:  float64(1.4913024695829),
		Fwest:  float64(1.14714731736311),
	},
	49: {
		Fnorth: float64(0.33421063142418),
		Fsouth: float64(0.02526613430348),
		Feast:  float64(1.15141032578749),
		Fwest:  float64(0.85000706261644),
	},
	50: {
		Fnorth: float64(0.38915669778582),
		Fsouth: -libc.Float64FromFloat64(0.04371359825454),
		Feast:  float64(1.88046353933242),
		Fwest:  float64(1.48230231380717),
	},
	51: {
		Fnorth: float64(0.33787520825987),
		Fsouth: -libc.Float64FromFloat64(0.04835090128296),
		Feast:  -libc.Float64FromFloat64(1.12274014380603),
		Fwest:  -libc.Float64FromFloat64(1.49454408844749),
	},
	52: {
		Fnorth: float64(0.33601418932337),
		Fsouth: -libc.Float64FromFloat64(0.06675068178541),
		Feast:  float64(2.23792354204464),
		Fwest:  float64(1.85723423013211),
	},
	53: {
		Fnorth: float64(0.31838318078049),
		Fsouth: -libc.Float64FromFloat64(0.05821955623722),
		Feast:  float64(0.66058854060373),
		Fwest:  float64(0.25452572938783),
	},
	54: {
		Fnorth: float64(0.33630761471457),
		Fsouth: -libc.Float64FromFloat64(0.07589541031521),
		Feast:  -libc.Float64FromFloat64(1.47957331741818),
		Fwest:  -libc.Float64FromFloat64(1.85981735718264),
	},
	55: {
		Fnorth: float64(0.2892481732287),
		Fsouth: -libc.Float64FromFloat64(0.09150638064667),
		Feast:  -libc.Float64FromFloat64(1.83561930288569),
		Fwest:  -libc.Float64FromFloat64(2.21855897384292),
	},
	56: {
		Fnorth: float64(0.26678632252475),
		Fsouth: -libc.Float64FromFloat64(0.10058088990867),
		Feast:  -libc.Float64FromFloat64(2.76808651991421),
		Fwest:  float64(3.12792953247061),
	},
	57: {
		Fnorth: float64(0.29285254112587),
		Fsouth: -libc.Float64FromFloat64(0.13483165093783),
		Feast:  float64(2.61406468380434),
		Fwest:  float64(2.20466422911705),
	},
	58: {
		Fnorth: float64(0.20150342788824),
		Fsouth: -libc.Float64FromFloat64(0.10279852729762),
		Feast:  float64(0.06881896344365),
		Fwest:  -libc.Float64FromFloat64(0.23925229432978),
	},
	59: {
		Fnorth: float64(0.21283813275258),
		Fsouth: -libc.Float64FromFloat64(0.18626835417891),
		Feast:  float64(2.93800440256577),
		Fwest:  float64(2.57470747655623),
	},
	60: {
		Fnorth: float64(0.19587614179884),
		Fsouth: -libc.Float64FromFloat64(0.17237030304155),
		Feast:  -libc.Float64FromFloat64(2.16941795427335),
		Fwest:  -libc.Float64FromFloat64(2.55405165906601),
	},
	61: {
		Fnorth: float64(0.17237030304155),
		Fsouth: -libc.Float64FromFloat64(0.19587614179884),
		Feast:  float64(0.97217469931645),
		Fwest:  float64(0.58754099452378),
	},
	62: {
		Fnorth: float64(0.18626835417891),
		Fsouth: -libc.Float64FromFloat64(0.21283813275258),
		Feast:  -libc.Float64FromFloat64(0.20358825102402),
		Fwest:  -libc.Float64FromFloat64(0.56688517703356),
	},
	63: {
		Fnorth: float64(0.10279852729762),
		Fsouth: -libc.Float64FromFloat64(0.20150342788824),
		Feast:  -libc.Float64FromFloat64(3.07277369014614),
		Fwest:  float64(2.90234035926002),
	},
	64: {
		Fnorth: float64(0.13483165093783),
		Fsouth: -libc.Float64FromFloat64(0.29285254112587),
		Feast:  -libc.Float64FromFloat64(0.52752796978545),
		Fwest:  -libc.Float64FromFloat64(0.93692842447275),
	},
	65: {
		Fnorth: float64(0.10058088990867),
		Fsouth: -libc.Float64FromFloat64(0.26678632252475),
		Feast:  float64(0.37350613367558),
		Fwest:  -libc.Float64FromFloat64(0.01366312111919),
	},
	66: {
		Fnorth: float64(0.09150638064667),
		Fsouth: -libc.Float64FromFloat64(0.2892481732287),
		Feast:  float64(1.3059733507041),
		Fwest:  float64(0.92303367974687),
	},
	67: {
		Fnorth: float64(0.07589541031521),
		Fsouth: -libc.Float64FromFloat64(0.33630761471457),
		Feast:  float64(1.66201933617161),
		Fwest:  float64(1.28177529640715),
	},
	68: {
		Fnorth: float64(0.05821955623722),
		Fsouth: -libc.Float64FromFloat64(0.31838318078049),
		Feast:  -libc.Float64FromFloat64(2.48100411298606),
		Fwest:  -libc.Float64FromFloat64(2.88706692420196),
	},
	69: {
		Fnorth: float64(0.06675068178541),
		Fsouth: -libc.Float64FromFloat64(0.33601418932337),
		Feast:  -libc.Float64FromFloat64(0.90366911154516),
		Fwest:  -libc.Float64FromFloat64(1.28435842345769),
	},
	70: {
		Fnorth: float64(0.04835090128296),
		Fsouth: -libc.Float64FromFloat64(0.33787520825987),
		Feast:  float64(2.01885250978376),
		Fwest:  float64(1.6470485651423),
	},
	71: {
		Fnorth: float64(0.04371359825454),
		Fsouth: -libc.Float64FromFloat64(0.38915669778582),
		Feast:  -libc.Float64FromFloat64(1.26112911425737),
		Fwest:  -libc.Float64FromFloat64(1.65929033978262),
	},
	72: {
		Fnorth: -libc.Float64FromFloat64(0.02526613430348),
		Fsouth: -libc.Float64FromFloat64(0.33421063142418),
		Feast:  -libc.Float64FromFloat64(1.99018232780231),
		Fwest:  -libc.Float64FromFloat64(2.29158559097336),
	},
	73: {
		Fnorth: float64(0.01518764903038),
		Fsouth: -libc.Float64FromFloat64(0.3911381668714),
		Feast:  -libc.Float64FromFloat64(1.6502901840069),
		Fwest:  -libc.Float64FromFloat64(1.99444533622668),
	},
	74: {
		Fnorth: float64(0.00297016159959),
		Fsouth: -libc.Float64FromFloat64(0.38073727558986),
		Feast:  float64(2.37119711497761),
		Fwest:  float64(1.99371016613951),
	},
	75: {
		Fnorth: -libc.Float64FromFloat64(0.02927431919853),
		Fsouth: -libc.Float64FromFloat64(0.43103892586713),
		Feast:  float64(0.72570026740558),
		Fwest:  float64(0.28423455407029),
	},
	76: {
		Fnorth: -libc.Float64FromFloat64(0.0779644093814),
		Fsouth: -libc.Float64FromFloat64(0.43208958202064),
		Feast:  float64(0.07926812279319),
		Fwest:  -libc.Float64FromFloat64(0.33556765368697),
	},
	77: {
		Fnorth: -libc.Float64FromFloat64(0.06926857458503),
		Fsouth: -libc.Float64FromFloat64(0.44686891066946),
		Feast:  -libc.Float64FromFloat64(2.82105980564027),
		Fwest:  float64(3.07153516458131),
	},
	78: {
		Fnorth: -libc.Float64FromFloat64(0.0602996863772),
		Fsouth: -libc.Float64FromFloat64(0.46538045483671),
		Feast:  float64(2.72918651645558),
		Fwest:  float64(2.33555641550814),
	},
	79: {
		Fnorth: -libc.Float64FromFloat64(0.10944898890579),
		Fsouth: -libc.Float64FromFloat64(0.49767951537101),
		Feast:  float64(3.09824103095621),
		Fwest:  float64(2.7125899718041),
	},
	80: {
		Fnorth: -libc.Float64FromFloat64(0.15522020377124),
		Fsouth: -libc.Float64FromFloat64(0.51206347752697),
		Feast:  -libc.Float64FromFloat64(2.18712498042983),
		Fwest:  -libc.Float64FromFloat64(2.59716003248565),
	},
	81: {
		Fnorth: -libc.Float64FromFloat64(0.15589091511369),
		Fsouth: -libc.Float64FromFloat64(0.5460368797045),
		Feast:  float64(1.07369399291919),
		Fwest:  float64(0.65067845727018),
	},
	82: {
		Fnorth: -libc.Float64FromFloat64(0.15187401321019),
		Fsouth: -libc.Float64FromFloat64(0.55648013300665),
		Feast:  -libc.Float64FromFloat64(0.27127176937655),
		Fwest:  -libc.Float64FromFloat64(0.69516944883612),
	},
	83: {
		Fnorth: -libc.Float64FromFloat64(0.2527655743723),
		Fsouth: -libc.Float64FromFloat64(0.55637406851385),
		Feast:  float64(2.14273876853285),
		Fwest:  float64(1.81516776000041),
	},
	84: {
		Fnorth: -libc.Float64FromFloat64(0.2162046083657),
		Fsouth: -libc.Float64FromFloat64(0.59919175361146),
		Feast:  -libc.Float64FromFloat64(0.6556639697129),
		Fwest:  -libc.Float64FromFloat64(1.06808911465388),
	},
	85: {
		Fnorth: -libc.Float64FromFloat64(0.21206753429148),
		Fsouth: -libc.Float64FromFloat64(0.64228460410023),
		Feast:  float64(1.46520024366909),
		Fwest:  float64(1.02386898591638),
	},
	86: {
		Fnorth: -libc.Float64FromFloat64(0.25522080363509),
		Fsouth: -libc.Float64FromFloat64(0.62380174028378),
		Feast:  float64(0.41730842332153),
		Fwest:  -libc.Float64FromFloat64(0.0375779212135),
	},
	87: {
		Fnorth: -libc.Float64FromFloat64(0.2414986570985),
		Fsouth: -libc.Float64FromFloat64(0.64750997738584),
		Feast:  float64(1.83887072884423),
		Fwest:  float64(1.45450695195737),
	},
	88: {
		Fnorth: -libc.Float64FromFloat64(0.28222653310929),
		Fsouth: -libc.Float64FromFloat64(0.65722892279906),
		Feast:  -libc.Float64FromFloat64(1.83240572073513),
		Fwest:  -libc.Float64FromFloat64(2.26564849087294),
	},
	89: {
		Fnorth: -libc.Float64FromFloat64(0.26290420067147),
		Fsouth: -libc.Float64FromFloat64(0.64863235654749),
		Feast:  -libc.Float64FromFloat64(1.03841167090601),
		Fwest:  -libc.Float64FromFloat64(1.44603142810635),
	},
	90: {
		Fnorth: -libc.Float64FromFloat64(0.28808313184381),
		Fsouth: -libc.Float64FromFloat64(0.69190629544818),
		Feast:  -libc.Float64FromFloat64(2.49295356114332),
		Fwest:  -libc.Float64FromFloat64(2.97786896076422),
	},
	91: {
		Fnorth: -libc.Float64FromFloat64(0.29148673180722),
		Fsouth: -libc.Float64FromFloat64(0.70359051048414),
		Feast:  -libc.Float64FromFloat64(1.42718183501734),
		Fwest:  -libc.Float64FromFloat64(1.85715916977284),
	},
	92: {
		Fnorth: -libc.Float64FromFloat64(0.33012478438475),
		Fsouth: -libc.Float64FromFloat64(0.71515840341957),
		Feast:  float64(2.49311289195816),
		Fwest:  float64(2.05909537237761),
	},
	93: {
		Fnorth: -libc.Float64FromFloat64(0.38742018303868),
		Fsouth: -libc.Float64FromFloat64(0.78861390967531),
		Feast:  -libc.Float64FromFloat64(2.91043577627328),
		Fwest:  float64(2.91559774272914),
	},
	94: {
		Fnorth: -libc.Float64FromFloat64(0.39943555383751),
		Fsouth: -libc.Float64FromFloat64(0.75822779851431),
		Feast:  float64(0.8009928727428),
		Fwest:  float64(0.32031891536535),
	},
	95: {
		Fnorth: -libc.Float64FromFloat64(0.38396800633226),
		Fsouth: -libc.Float64FromFloat64(0.81612079704781),
		Feast:  float64(2.9254488646714),
		Fwest:  float64(2.43739115636801),
	},
	96: {
		Fnorth: -libc.Float64FromFloat64(0.40150819579319),
		Fsouth: -libc.Float64FromFloat64(0.80599330438546),
		Feast:  float64(0.07741705955739),
		Fwest:  -libc.Float64FromFloat64(0.4407996457457),
	},
	97: {
		Fnorth: -libc.Float64FromFloat64(0.52173101741059),
		Fsouth: -libc.Float64FromFloat64(0.82570569254601),
		Feast:  -libc.Float64FromFloat64(0.83237583897551),
		Fwest:  -libc.Float64FromFloat64(1.20960723529999),
	},
	98: {
		Fnorth: -libc.Float64FromFloat64(0.48781264892508),
		Fsouth: -libc.Float64FromFloat64(0.89235638181782),
		Feast:  -libc.Float64FromFloat64(0.3772896323983),
		Fwest:  -libc.Float64FromFloat64(0.84169548661948),
	},
	99: {
		Fnorth: -libc.Float64FromFloat64(0.50770567021439),
		Fsouth: -libc.Float64FromFloat64(0.86881343251986),
		Feast:  -libc.Float64FromFloat64(2.10922469863141),
		Fwest:  -libc.Float64FromFloat64(2.63811981331554),
	},
	100: {
		Fnorth: -libc.Float64FromFloat64(0.52742963716774),
		Fsouth: -libc.Float64FromFloat64(0.88611243445554),
		Feast:  float64(2.18377319034785),
		Fwest:  float64(1.6653029905305),
	},
	101: {
		Fnorth: -libc.Float64FromFloat64(0.50008952762292),
		Fsouth: -libc.Float64FromFloat64(0.87536136210723),
		Feast:  float64(1.21723651787549),
		Fwest:  float64(0.72517922139186),
	},
	102: {
		Fnorth: -libc.Float64FromFloat64(0.55186491737474),
		Fsouth: -libc.Float64FromFloat64(0.96374551790056),
		Feast:  float64(1.7118254404532),
		Fwest:  float64(1.17307062828876),
	},
	103: {
		Fnorth: -libc.Float64FromFloat64(0.5867983657157),
		Fsouth: -libc.Float64FromFloat64(0.98106023192774),
		Feast:  -libc.Float64FromFloat64(1.11329499144518),
		Fwest:  -libc.Float64FromFloat64(1.62824890388699),
	},
	104: {
		Fnorth: -libc.Float64FromFloat64(0.61076063532947),
		Fsouth: -libc.Float64FromFloat64(0.9808143413602),
		Feast:  float64(0.4705862876045),
		Fwest:  -libc.Float64FromFloat64(0.07642802350246),
	},
	105: {
		Fnorth: -libc.Float64FromFloat64(0.58827636737638),
		Fsouth: -libc.Float64FromFloat64(1.01786037568858),
		Feast:  -libc.Float64FromFloat64(1.60966543541914),
		Fwest:  -libc.Float64FromFloat64(2.20486582847747),
	},
	106: {
		Fnorth: -libc.Float64FromFloat64(0.65491232835426),
		Fsouth: -libc.Float64FromFloat64(1.01929924623886),
		Feast:  -libc.Float64FromFloat64(2.51123691118248),
		Fwest:  -libc.Float64FromFloat64(3.1062223526251),
	},
	107: {
		Fnorth: -libc.Float64FromFloat64(0.72356358827215),
		Fsouth: -libc.Float64FromFloat64(1.0327022874896),
		Feast:  float64(0.89169126633833),
		Fwest:  float64(0.39649044439822),
	},
	108: {
		Fnorth: -libc.Float64FromFloat64(0.67603948819406),
		Fsouth: -libc.Float64FromFloat64(1.07042472765165),
		Feast:  float64(2.66175062518892),
		Fwest:  float64(2.03853105755889),
	},
	109: {
		Fnorth: -libc.Float64FromFloat64(0.71330093663066),
		Fsouth: -libc.Float64FromFloat64(1.10121643537669),
		Feast:  -libc.Float64FromFloat64(3.04518683458825),
		Fwest:  float64(2.62004750840731),
	},
	110: {
		Fnorth: -libc.Float64FromFloat64(0.76641428724335),
		Fsouth: -libc.Float64FromFloat64(1.15528445067052),
		Feast:  float64(0.07420758156568),
		Fwest:  -libc.Float64FromFloat64(0.60513155114938),
	},
	111: {
		Fnorth: -libc.Float64FromFloat64(0.78982455384253),
		Fsouth: -libc.Float64FromFloat64(1.15546530929588),
		Feast:  -libc.Float64FromFloat64(0.60499853129713),
		Fwest:  -libc.Float64FromFloat64(1.28450131907736),
	},
	112: {
		Fnorth: -libc.Float64FromFloat64(0.83795331049498),
		Fsouth: -libc.Float64FromFloat64(1.21075831414294),
		Feast:  float64(1.42136389579088),
		Fwest:  float64(0.70365403631841),
	},
	113: {
		Fnorth: -libc.Float64FromFloat64(0.86170600921069),
		Fsouth: -libc.Float64FromFloat64(1.21114673854945),
		Feast:  -libc.Float64FromFloat64(1.95029507749525),
		Fwest:  -libc.Float64FromFloat64(2.70381656362525),
	},
	114: {
		Fnorth: -libc.Float64FromFloat64(0.84291228415618),
		Fsouth: -libc.Float64FromFloat64(1.26020983864103),
		Feast:  float64(2.24187397694118),
		Fwest:  float64(1.38191906047983),
	},
	115: {
		Fnorth: -libc.Float64FromFloat64(0.9427181537636),
		Fsouth: -libc.Float64FromFloat64(1.32899086063916),
		Feast:  float64(0.84283975752601),
		Fwest:  -libc.Float64FromFloat64(0.12459257316986),
	},
	116: {
		Fnorth: -libc.Float64FromFloat64(0.91898920750071),
		Fsouth: -libc.Float64FromFloat64(1.32929586572429),
		Feast:  -libc.Float64FromFloat64(1.08536920415787),
		Fwest:  -libc.Float64FromFloat64(2.05346111080706),
	},
	117: {
		Fnorth: -libc.Float64FromFloat64(0.97226652536306),
		Fsouth: -libc.Float64FromFloat64(1.27950477868453),
		Feast:  -libc.Float64FromFloat64(2.58603200375485),
		Fwest:  float64(2.95929340513654),
	},
	118: {
		Fnorth: -libc.Float64FromFloat64(1.01285145697208),
		Fsouth: -libc.Float64FromFloat64(1.41845302535151),
		Feast:  -libc.Float64FromFloat64(3.13590968086981),
		Fwest:  float64(1.97388885726377),
	},
	119: {
		Fnorth: -libc.Float64FromFloat64(1.09069387298096),
		Fsouth: -libc.Float64FromFloat64(1.52480158339146),
		Feast:  float64(0.28873212061306),
		Fwest:  -libc.Float64FromFloat64(1.49848576331087),
	},
	120: {
		Fnorth: -libc.Float64FromFloat64(1.17872424267511),
		Fsouth: -libc.Float64FromFloat64(1.52480158339146),
		Feast:  float64(2.53494381704943),
		Fwest:  -libc.Float64FromFloat64(0.60112285060716),
	},
	121: {
		Fnorth: -libc.Float64FromFloat64(1.20305471830087),
		Fsouth: -libc.Float64FromFloat64(1.52480158339146),
		Feast:  -libc.Float64FromFloat64(0.60112285060716),
		Fwest:  float64(2.53494381704943),
	},
}

var _VALID_RANGE_BBOX = TBBox{
	Fnorth: float64(1.5707963267948966),
	Fsouth: -libc.Float64FromFloat64(1.5707963267948966),
	Feast:  float64(3.141592653589793),
	Fwest:  -libc.Float64FromFloat64(3.141592653589793),
}

// C documentation
//
//	/**
//	 * For a given cell, return its bounding box. If coverChildren is true, the bbox
//	 * will be guaranteed to contain its children at any finer resolution. Note that
//	 * no guarantee is provided as to the level of accuracy, and the bounding box
//	 * may have a significant margin of error.
//	 * @param cell Cell to calculate bbox for
//	 * @param out  BBox to hold output
//	 * @param coverChildren Whether the bounding box should cover all children
//	 */
func XcellToBBox(tls *libc.TLS, cell TH3Index, out uintptr, coverChildren uint8) (r TH3Error) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var baseCell, res, v1 int32
	var centerErr TH3Error
	var lngRatio, v3 float64
	var v2 bool
	var _ /* center at bp+0 */ TLatLng
	_, _, _, _, _, _, _ = baseCell, centerErr, lngRatio, res, v1, v2, v3
	// Adjust the BBox to handle poles, if needed
	res = libc.Int32FromUint64(cell & (libc.Uint64FromUint64(15) << libc.Int32FromInt32(DH3_RES_OFFSET)) >> libc.Int32FromInt32(DH3_RES_OFFSET))
	if res == 0 {
		baseCell = libc.Int32FromUint64(cell & (libc.Uint64FromInt32(libc.Int32FromInt32(127)) << libc.Int32FromInt32(DH3_BC_OFFSET)) >> libc.Int32FromInt32(DH3_BC_OFFSET))
		if baseCell < 0 {
			if v2 = libc.Bool(0 != 0); !v2 {
				libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+1313, int32(227), uintptr(unsafe.Pointer(&___func__22)))
			}
			_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
			v1 = libc.Int32FromInt32(1)
		} else {
			v1 = 0
		}
		if v1 != 0 || baseCell >= int32(DNUM_BASE_CELLS) {
			return uint32(_E_CELL_INVALID)
		}
		*(*TBBox)(unsafe.Pointer(out)) = _RES0_BBOXES[baseCell]
	} else {
		centerErr = XcellToLatLng(tls, cell, bp)
		if centerErr != uint32(_E_SUCCESS) {
			return centerErr
		}
		lngRatio = libc.Float64FromInt32(1) / libc.Xcos(tls, (*(*TLatLng)(unsafe.Pointer(bp))).Flat)
		(*TBBox)(unsafe.Pointer(out)).Fnorth = (*(*TLatLng)(unsafe.Pointer(bp))).Flat + _MAX_EDGE_LENGTH_RADS[res]
		(*TBBox)(unsafe.Pointer(out)).Fsouth = (*(*TLatLng)(unsafe.Pointer(bp))).Flat - _MAX_EDGE_LENGTH_RADS[res]
		(*TBBox)(unsafe.Pointer(out)).Feast = (*(*TLatLng)(unsafe.Pointer(bp))).Flng + _MAX_EDGE_LENGTH_RADS[res]*lngRatio
		(*TBBox)(unsafe.Pointer(out)).Fwest = (*(*TLatLng)(unsafe.Pointer(bp))).Flng - _MAX_EDGE_LENGTH_RADS[res]*lngRatio
	}
	// Buffer the bounding box to cover children. Call this even if no buffering
	// is required in order to normalize the bbox to lat/lng bounds
	if coverChildren != 0 {
		v3 = float64(DCHILD_SCALE_FACTOR)
	} else {
		v3 = float64(DCELL_SCALE_FACTOR)
	}
	XscaleBBox(tls, out, v3)
	// Cell that contains the north pole
	if cell == _NORTH_POLE_CELLS[res] {
		(*TBBox)(unsafe.Pointer(out)).Fnorth = float64(1.5707963267948966)
	}
	// Cell that contains the south pole
	if cell == _SOUTH_POLE_CELLS[res] {
		(*TBBox)(unsafe.Pointer(out)).Fsouth = -libc.Float64FromFloat64(1.5707963267948966)
	}
	// If we contain a pole, expand the longitude to include the full domain,
	// effectively making the bbox a circle around the pole.
	if (*TBBox)(unsafe.Pointer(out)).Fnorth == float64(1.5707963267948966) || (*TBBox)(unsafe.Pointer(out)).Fsouth == -libc.Float64FromFloat64(1.5707963267948966) {
		(*TBBox)(unsafe.Pointer(out)).Feast = float64(3.141592653589793)
		(*TBBox)(unsafe.Pointer(out)).Fwest = -libc.Float64FromFloat64(3.141592653589793)
	}
	return uint32(_E_SUCCESS)
}

var ___func__22 = [11]uint8{'c', 'e', 'l', 'l', 'T', 'o', 'B', 'B', 'o', 'x'}

// C documentation
//
//	/**
//	 * Get a base cell by number, or H3_NULL if out of bounds
//	 */
func XbaseCellNumToCell(tls *libc.TLS, baseCellNum int32) (r TH3Index) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* baseCell at bp+0 */ TH3Index
	if baseCellNum < 0 || baseCellNum >= int32(DNUM_BASE_CELLS) {
		return uint64(DH3_NULL)
	}
	XsetH3Index(tls, bp, 0, baseCellNum, 0)
	return *(*TH3Index)(unsafe.Pointer(bp))
}

func _iterErrorPolygonCompact(tls *libc.TLS, iter uintptr, error1 TH3Error) {
	XiterDestroyPolygonCompact(tls, iter)
	(*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).Ferror1 = error1
}

// C documentation
//
//	/**
//	 * Given a cell, find the next cell in the sequence of all cells
//	 * to check in the iteration.
//	 */
func _nextCell(tls *libc.TLS, cell TH3Index) (r TH3Index) {
	var digit TDirection
	var parent TH3Index
	var res, v1 int32
	_, _, _, _ = digit, parent, res, v1
	res = libc.Int32FromUint64(cell & (libc.Uint64FromUint64(15) << libc.Int32FromInt32(DH3_RES_OFFSET)) >> libc.Int32FromInt32(DH3_RES_OFFSET))
	for int32(Dtrue) != 0 {
		// If this is a base cell, set to next base cell (or H3_NULL if done)
		if res == 0 {
			return XbaseCellNumToCell(tls, libc.Int32FromUint64(cell&(libc.Uint64FromInt32(libc.Int32FromInt32(127))<<libc.Int32FromInt32(DH3_BC_OFFSET))>>libc.Int32FromInt32(DH3_BC_OFFSET))+int32(1))
		}
		// Faster cellToParent when we know the resolution is valid
		// and we're only moving up one level
		parent = cell
		parent = parent & ^(libc.Uint64FromUint64(15)<<libc.Int32FromInt32(DH3_RES_OFFSET)) | libc.Uint64FromInt32(res-libc.Int32FromInt32(1))<<libc.Int32FromInt32(DH3_RES_OFFSET)
		parent = parent & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<((libc.Int32FromInt32(DMAX_H3_RES)-res)*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))) | libc.Uint64FromInt32(libc.Int32FromInt32(7))<<((libc.Int32FromInt32(DMAX_H3_RES)-res)*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))
		// If not the last sibling of parent, return next sibling
		digit = libc.Int32FromUint64(cell >> ((libc.Int32FromInt32(DMAX_H3_RES) - res) * libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET)) & libc.Uint64FromInt32(libc.Int32FromInt32(7)))
		if digit < int32(_INVALID_DIGIT)-libc.Int32FromInt32(1) {
			if XisPentagon(tls, parent) != 0 && digit == int32(_CENTER_DIGIT) {
				v1 = int32(2)
			} else {
				v1 = int32(1)
			}
			cell = cell & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<((libc.Int32FromInt32(DMAX_H3_RES)-res)*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))) | libc.Uint64FromInt32(digit+v1)<<((libc.Int32FromInt32(DMAX_H3_RES)-res)*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))
			return cell
		}
		// Move up to the parent for the next loop iteration
		res--
		cell = parent
	}
	return r
}

// C documentation
//
//	/**
//	 * Internal function - initialize the iterator without stepping to the first
//	 * value
//	 */
func __iterInitPolygonCompact(tls *libc.TLS, polygon uintptr, res int32, flags Tuint32_t) (r TIterCellsPolygonCompact) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var flagErr TH3Error
	var _ /* iter at bp+0 */ TIterCellsPolygonCompact
	_ = flagErr
	*(*TIterCellsPolygonCompact)(unsafe.Pointer(bp)) = TIterCellsPolygonCompact{
		Fcell:     XbaseCellNumToCell(tls, 0),
		F_res:     res,
		F_flags:   flags,
		F_polygon: polygon,
	}
	if res < 0 || res > int32(DMAX_H3_RES) {
		_iterErrorPolygonCompact(tls, bp, uint32(_E_RES_DOMAIN))
		return *(*TIterCellsPolygonCompact)(unsafe.Pointer(bp))
	}
	flagErr = XvalidatePolygonFlags(tls, flags)
	if flagErr != 0 {
		_iterErrorPolygonCompact(tls, bp, flagErr)
		return *(*TIterCellsPolygonCompact)(unsafe.Pointer(bp))
	}
	// Initialize bounding boxes for polygon and any holes. Memory allocated
	// here must be released through iterDestroyPolygonCompact
	(*(*TIterCellsPolygonCompact)(unsafe.Pointer(bp))).F_bboxes = libc.Xcalloc(tls, libc.Uint64FromInt32((*TGeoPolygon)(unsafe.Pointer(polygon)).FnumHoles+libc.Int32FromInt32(1)), uint64(32))
	if !((*(*TIterCellsPolygonCompact)(unsafe.Pointer(bp))).F_bboxes != 0) {
		_iterErrorPolygonCompact(tls, bp, uint32(_E_MEMORY_ALLOC))
		return *(*TIterCellsPolygonCompact)(unsafe.Pointer(bp))
	}
	XbboxesFromGeoPolygon(tls, polygon, (*(*TIterCellsPolygonCompact)(unsafe.Pointer(bp))).F_bboxes)
	return *(*TIterCellsPolygonCompact)(unsafe.Pointer(bp))
}

// C documentation
//
//	/**
//	 * Initialize a IterCellsPolygonCompact struct representing the sequence of
//	 * compact cells within the target polygon. The test for including edge cells is
//	 * defined by the polyfill mode passed in the `flags` argument.
//	 *
//	 * Initialization of this object may fail, in which case the `error` property
//	 * will be set and all iteration will return H3_NULL. It is the responsibility
//	 * of the caller to check the error property after initialization.
//	 *
//	 * At any point in the iteration, starting once the struct is initialized, the
//	 * output value can be accessed through the `cell` property.
//	 *
//	 * Note that initializing the iterator allocates memory. If an iterator is
//	 * exhausted or returns an error that memory is released; otherwise it must be
//	 * released manually with iterDestroyPolygonCompact.
//	 *
//	 * @param  polygon Polygon to fill with compact cells
//	 * @param  res     Finest resolution for output cells
//	 * @param  flags   Bit mask of option flags
//	 * @return         Initialized iterator, with the first value available
//	 */
func XiterInitPolygonCompact(tls *libc.TLS, polygon uintptr, res int32, flags Tuint32_t) (r TIterCellsPolygonCompact) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var _ /* iter at bp+0 */ TIterCellsPolygonCompact
	*(*TIterCellsPolygonCompact)(unsafe.Pointer(bp)) = __iterInitPolygonCompact(tls, polygon, res, flags)
	// Start the iterator by taking the first step.
	// This is necessary to have a valid value after initialization.
	XiterStepPolygonCompact(tls, bp)
	return *(*TIterCellsPolygonCompact)(unsafe.Pointer(bp))
}

// C documentation
//
//	/**
//	 * Increment the polyfill iterator, running the polygon to cells algorithm.
//	 *
//	 * Briefly, the algorithm checks every cell in the global grid hierarchically,
//	 * starting with the base cells. Cells coarser than the target resolution are
//	 * checked for complete child inclusion using a bounding box guaranteed to
//	 * contain all children.
//	 * - If the bounding box is contained by the polygon, output is set to the cell
//	 * - If the bounding box intersects, recurse into the first child
//	 * - Otherwise, continue with the next cell in sequence
//	 *
//	 * For cells at the target resolution, a finer-grained check is used according
//	 * to the inclusion criteria set in flags.
//	 *
//	 * @param  iter Iterator to increment
//	 */
func XiterStepPolygonCompact(tls *libc.TLS, iter uintptr) {
	bp := tls.Alloc(656)
	defer tls.Free(656)
	var bboxErr, bboxErr1, bboxErr2, boundaryErr, centerErr, childErr, polygonCellErr TH3Error
	var cell TH3Index
	var cellRes, v1, v3 int32
	var mode TContainmentMode
	var v2, v4 bool
	var _ /* bbox at bp+208 */ TBBox
	var _ /* bbox at bp+240 */ TBBox
	var _ /* bbox at bp+440 */ TBBox
	var _ /* bboxBoundary at bp+272 */ TCellBoundary
	var _ /* bboxBoundary at bp+472 */ TCellBoundary
	var _ /* boundary at bp+40 */ TCellBoundary
	var _ /* center at bp+0 */ TLatLng
	var _ /* child at bp+640 */ TH3Index
	var _ /* firstVertex at bp+16 */ TLatLng
	var _ /* polygonCell at bp+32 */ TH3Index
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = bboxErr, bboxErr1, bboxErr2, boundaryErr, cell, cellRes, centerErr, childErr, mode, polygonCellErr, v1, v2, v3, v4
	cell = (*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).Fcell
	// once the cell is H3_NULL, the iterator returns an infinite sequence of
	// H3_NULL
	if cell == uint64(DH3_NULL) {
		return
	}
	// For the first step, we need to evaluate the current cell; after that, we
	// should start with the next cell.
	if (*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_started != 0 {
		cell = _nextCell(tls, cell)
	} else {
		(*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_started = uint8(Dtrue)
	}
	// Short-circuit iteration for 0-vert polygon
	if (*TGeoPolygon)(unsafe.Pointer((*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_polygon)).Fgeoloop.FnumVerts == 0 {
		XiterDestroyPolygonCompact(tls, iter)
		return
	}
	mode = libc.Int32FromUint32((*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_flags & libc.Uint32FromInt32(libc.Int32FromInt32(15)))
	for cell != 0 {
		cellRes = libc.Int32FromUint64(cell & (libc.Uint64FromUint64(15) << libc.Int32FromInt32(DH3_RES_OFFSET)) >> libc.Int32FromInt32(DH3_RES_OFFSET))
		// Target res: Do a fine-grained check
		if cellRes == (*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_res {
			if mode == int32(_CONTAINMENT_CENTER) || mode == int32(_CONTAINMENT_OVERLAPPING) || mode == int32(_CONTAINMENT_OVERLAPPING_BBOX) {
				centerErr = XcellToLatLng(tls, cell, bp)
				if centerErr != uint32(_E_SUCCESS) {
					_iterErrorPolygonCompact(tls, iter, centerErr)
					return
				}
				if XpointInsidePolygon(tls, (*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_polygon, (*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_bboxes, bp) != 0 {
					// Set to next output
					(*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).Fcell = cell
					return
				}
			}
			if mode == int32(_CONTAINMENT_OVERLAPPING) || mode == int32(_CONTAINMENT_OVERLAPPING_BBOX) {
				// For overlapping, we need to do a quick check to determine
				// whether the polygon is wholly contained by the cell. We
				// check the first polygon vertex, which if it is contained
				// could also mean we simply intersect.
				// Deferencing verts[0] is safe because we check numVerts above
				*(*TLatLng)(unsafe.Pointer(bp + 16)) = *(*TLatLng)(unsafe.Pointer((*TGeoPolygon)(unsafe.Pointer((*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_polygon)).Fgeoloop.Fverts))
				// We have to check whether the point is in the expected range
				// first, because out-of-bounds values will yield false
				// positives with latLngToCell
				if XbboxContains(tls, uintptr(unsafe.Pointer(&_VALID_RANGE_BBOX)), bp+16) != 0 {
					polygonCellErr = XlatLngToCell(tls, (*TGeoPolygon)(unsafe.Pointer((*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_polygon)).Fgeoloop.Fverts, cellRes, bp+32)
					if polygonCellErr != uint32(_E_SUCCESS) {
						if v2 = libc.Bool(0 != 0); !v2 {
							libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+1313, int32(470), uintptr(unsafe.Pointer(&___func__23)))
						}
						_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
						v1 = libc.Int32FromInt32(1)
					} else {
						v1 = 0
					}
					if v1 != 0 {
						// This should be unreachable with the bbox check
						_iterErrorPolygonCompact(tls, iter, polygonCellErr)
						return
					}
					if *(*TH3Index)(unsafe.Pointer(bp + 32)) == cell {
						// Set to next output
						(*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).Fcell = cell
						return
					}
				}
			}
			if mode == int32(_CONTAINMENT_FULL) || mode == int32(_CONTAINMENT_OVERLAPPING) || mode == int32(_CONTAINMENT_OVERLAPPING_BBOX) {
				boundaryErr = XcellToBoundary(tls, cell, bp+40)
				if boundaryErr != uint32(_E_SUCCESS) {
					_iterErrorPolygonCompact(tls, iter, boundaryErr)
					return
				}
				bboxErr = XcellToBBox(tls, cell, bp+208, uint8(Dfalse))
				if bboxErr != uint32(_E_SUCCESS) {
					if v4 = libc.Bool(0 != 0); !v4 {
						libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+1313, int32(493), uintptr(unsafe.Pointer(&___func__23)))
					}
					_ = v4 || libc.Bool(libc.Int32FromInt32(0) != 0)
					v3 = libc.Int32FromInt32(1)
				} else {
					v3 = 0
				}
				if v3 != 0 {
					// Should be unreachable - invalid cells would be caught in
					// the previous boundaryErr
					_iterErrorPolygonCompact(tls, iter, bboxErr)
					return
				}
				// Check if the cell is fully contained by the polygon
				if (mode == int32(_CONTAINMENT_FULL) || mode == int32(_CONTAINMENT_OVERLAPPING_BBOX)) && XcellBoundaryInsidePolygon(tls, (*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_polygon, (*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_bboxes, bp+40, bp+208) != 0 {
					// Set to next output
					(*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).Fcell = cell
					return
				} else {
					if (mode == int32(_CONTAINMENT_OVERLAPPING) || mode == int32(_CONTAINMENT_OVERLAPPING_BBOX)) && XcellBoundaryCrossesPolygon(tls, (*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_polygon, (*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_bboxes, bp+40, bp+208) != 0 {
						// Set to next output
						(*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).Fcell = cell
						return
					}
				}
			}
			if mode == int32(_CONTAINMENT_OVERLAPPING_BBOX) {
				bboxErr1 = XcellToBBox(tls, cell, bp+240, uint8(Dtrue))
				if bboxErr1 != 0 {
					_iterErrorPolygonCompact(tls, iter, bboxErr1)
					return
				}
				if XbboxOverlapsBBox(tls, (*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_bboxes, bp+240) != 0 {
					*(*TCellBoundary)(unsafe.Pointer(bp + 272)) = XbboxToCellBoundary(tls, bp+240)
					if XbboxContainsBBox(tls, bp+240, (*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_bboxes) != 0 || XpointInsidePolygon(tls, (*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_polygon, (*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_bboxes, bp+272+8) != 0 || XcellBoundaryCrossesPolygon(tls, (*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_polygon, (*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_bboxes, bp+272, bp+240) != 0 {
						(*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).Fcell = cell
						return
					}
				}
			}
		}
		// Coarser cell: Check the bounding box
		if cellRes < (*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_res {
			bboxErr2 = XcellToBBox(tls, cell, bp+440, uint8(Dtrue))
			if bboxErr2 != 0 {
				_iterErrorPolygonCompact(tls, iter, bboxErr2)
				return
			}
			if XbboxOverlapsBBox(tls, (*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_bboxes, bp+440) != 0 {
				// Quick check for possible containment
				if XbboxContainsBBox(tls, (*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_bboxes, bp+440) != 0 {
					*(*TCellBoundary)(unsafe.Pointer(bp + 472)) = XbboxToCellBoundary(tls, bp+440)
					// Do a fine-grained, more expensive check on the polygon
					if XcellBoundaryInsidePolygon(tls, (*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_polygon, (*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_bboxes, bp+472, bp+440) != 0 {
						// Bounding box is fully contained, so all children are
						// included. Set to next output.
						(*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).Fcell = cell
						return
					}
				}
				childErr = XcellToCenterChild(tls, cell, cellRes+int32(1), bp+640)
				if childErr != 0 {
					_iterErrorPolygonCompact(tls, iter, childErr)
					return
				}
				// Restart the loop with the child cell
				cell = *(*TH3Index)(unsafe.Pointer(bp + 640))
				continue
			}
		}
		// Find the next cell in the sequence of all cells and continue
		cell = _nextCell(tls, cell)
	}
	// If we make it out of the loop, we're done
	XiterDestroyPolygonCompact(tls, iter)
}

var ___func__23 = [23]uint8{'i', 't', 'e', 'r', 'S', 't', 'e', 'p', 'P', 'o', 'l', 'y', 'g', 'o', 'n', 'C', 'o', 'm', 'p', 'a', 'c', 't'}

// C documentation
//
//	/**
//	 * Destroy an iterator, releasing any allocated memory. Iterators destroyed in
//	 * this manner are safe to use but will always return H3_NULL.
//	 * @param  iter Iterator to destroy
//	 */
func XiterDestroyPolygonCompact(tls *libc.TLS, iter uintptr) {
	if (*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_bboxes != 0 {
		libc.Xfree(tls, (*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_bboxes)
	}
	(*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).Fcell = uint64(DH3_NULL)
	(*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).Ferror1 = uint32(_E_SUCCESS)
	(*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_polygon = libc.UintptrFromInt32(0)
	(*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_res = -int32(1)
	(*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_flags = uint32(0)
	(*TIterCellsPolygonCompact)(unsafe.Pointer(iter)).F_bboxes = libc.UintptrFromInt32(0)
}

// C documentation
//
//	/**
//	 * Initialize a IterCellsPolygon struct representing the sequence of
//	 * cells within the target polygon. The test for including edge cells is defined
//	 * by the polyfill mode passed in the `flags` argument.
//	 *
//	 * Initialization of this object may fail, in which case the `error` property
//	 * will be set and all iteration will return H3_NULL. It is the responsibility
//	 * of the caller to check the error property after initialization.
//	 *
//	 * At any point in the iteration, starting once the struct is initialized, the
//	 * output value can be accessed through the `cell` property.
//	 *
//	 * Note that initializing the iterator allocates memory. If an iterator is
//	 * exhausted or returns an error that memory is released; otherwise it must be
//	 * released manually with iterDestroyPolygon.
//	 *
//	 * @param  polygon Polygon to fill with cells
//	 * @param  res     Resolution for output cells
//	 * @param  flags   Bit mask of option flags
//	 * @return         Initialized iterator, with the first value available
//	 */
func XiterInitPolygon(tls *libc.TLS, polygon uintptr, res int32, flags Tuint32_t) (r TIterCellsPolygon) {
	var cellIter TIterCellsPolygonCompact
	var childIter TIterCellsChildren
	var iter TIterCellsPolygon
	_, _, _ = cellIter, childIter, iter
	// Create the sub-iterator for compact cells
	cellIter = XiterInitPolygonCompact(tls, polygon, res, flags)
	// Create the sub-iterator for children
	childIter = XiterInitParent(tls, cellIter.Fcell, res)
	iter = TIterCellsPolygon{
		Fcell:       childIter.Fh,
		Ferror1:     cellIter.Ferror1,
		F_cellIter:  cellIter,
		F_childIter: childIter,
	}
	return iter
}

// C documentation
//
//	/**
//	 * Increment the polyfill iterator, outputting the latest cell at the
//	 * desired resolution.
//	 *
//	 * @param  iter Iterator to increment
//	 */
func XiterStepPolygon(tls *libc.TLS, iter uintptr) {
	if (*TIterCellsPolygon)(unsafe.Pointer(iter)).Fcell == uint64(DH3_NULL) {
		return
	}
	// See if there are more children to output
	XiterStepChild(tls, iter+64)
	if (*TIterCellsPolygon)(unsafe.Pointer(iter)).F_childIter.Fh != 0 {
		(*TIterCellsPolygon)(unsafe.Pointer(iter)).Fcell = (*TIterCellsPolygon)(unsafe.Pointer(iter)).F_childIter.Fh
		return
	}
	// Otherwise, increment the polyfill iterator
	XiterStepPolygonCompact(tls, iter+16)
	if (*TIterCellsPolygon)(unsafe.Pointer(iter)).F_cellIter.Fcell != 0 {
		X_iterInitParent(tls, (*TIterCellsPolygon)(unsafe.Pointer(iter)).F_cellIter.Fcell, (*TIterCellsPolygon)(unsafe.Pointer(iter)).F_cellIter.F_res, iter+64)
		(*TIterCellsPolygon)(unsafe.Pointer(iter)).Fcell = (*TIterCellsPolygon)(unsafe.Pointer(iter)).F_childIter.Fh
		return
	}
	// All done, set to null and report errors if any
	(*TIterCellsPolygon)(unsafe.Pointer(iter)).Fcell = uint64(DH3_NULL)
	(*TIterCellsPolygon)(unsafe.Pointer(iter)).Ferror1 = (*TIterCellsPolygon)(unsafe.Pointer(iter)).F_cellIter.Ferror1
}

// C documentation
//
//	/**
//	 * Destroy an iterator, releasing any allocated memory. Iterators destroyed in
//	 * this manner are safe to use but will always return H3_NULL.
//	 * @param  iter Iterator to destroy
//	 */
func XiterDestroyPolygon(tls *libc.TLS, iter uintptr) {
	XiterDestroyPolygonCompact(tls, iter+16)
	// null out the child iterator by passing H3_NULL
	X_iterInitParent(tls, uint64(DH3_NULL), 0, iter+64)
	(*TIterCellsPolygon)(unsafe.Pointer(iter)).Fcell = uint64(DH3_NULL)
	(*TIterCellsPolygon)(unsafe.Pointer(iter)).Ferror1 = uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * polygonToCells takes a given GeoJSON-like data structure and preallocated,
//	 * zeroed memory, and fills it with the hexagons that are contained by
//	 * the GeoJSON-like data structure. Polygons are considered in Cartesian space.
//	 *
//	 * @param polygon The geoloop and holes defining the relevant area
//	 * @param res The Hexagon resolution (0-15)
//	 * @param flags Algorithm flags such as containment mode
//	 * @param size Maximum number of indexes to write to `out`.
//	 * @param out The slab of zeroed memory to write to. Must be at least of size
//	 * `size`.
//	 */
func XpolygonToCellsExperimental(tls *libc.TLS, polygon uintptr, res int32, flags Tuint32_t, size Tint64_t, out uintptr) (r TH3Error) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var i, v2 Tint64_t
	var _ /* iter at bp+0 */ TIterCellsPolygon
	_, _ = i, v2
	*(*TIterCellsPolygon)(unsafe.Pointer(bp)) = XiterInitPolygon(tls, polygon, res, flags)
	i = 0
	for {
		if !((*(*TIterCellsPolygon)(unsafe.Pointer(bp))).Fcell != 0) {
			break
		}
		if i >= size {
			XiterDestroyPolygon(tls, bp)
			return uint32(_E_MEMORY_BOUNDS)
		}
		v2 = i
		i++
		*(*TH3Index)(unsafe.Pointer(out + uintptr(v2)*8)) = (*(*TIterCellsPolygon)(unsafe.Pointer(bp))).Fcell
		goto _1
	_1:
		;
		XiterStepPolygon(tls, bp)
	}
	return (*(*TIterCellsPolygon)(unsafe.Pointer(bp))).Ferror1
}

var _MAX_SIZE_CELL_THRESHOLD = int32(10)

func _getAverageCellArea(tls *libc.TLS, res int32) (r float64) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* hexAreaKm2 at bp+0 */ float64
	XgetHexagonAreaAvgKm2(tls, res, bp)
	return *(*float64)(unsafe.Pointer(bp))
}

// C documentation
//
//	/**
//	 * maxPolygonToCellsSize returns the number of cells to allocate space for
//	 * when performing a polygonToCells on the given GeoJSON-like data structure.
//	 * @param polygon A GeoJSON-like data structure indicating the poly to fill
//	 * @param res Hexagon resolution (0-15)
//	 * @param flags Bit mask of option flags
//	 * @param out number of cells to allocate for
//	 * @return 0 (E_SUCCESS) on success.
//	 */
func XmaxPolygonToCellsSizeExperimental(tls *libc.TLS, polygon uintptr, res int32, flags Tuint32_t, out uintptr) (r TH3Error) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var polygonBBox uintptr
	var polygonBBoxAreaKm2 float64
	var _ /* childrenSize at bp+48 */ Tint64_t
	var _ /* iter at bp+0 */ TIterCellsPolygonCompact
	_, _ = polygonBBox, polygonBBoxAreaKm2
	// Special case: 0-vertex polygon
	if (*TGeoPolygon)(unsafe.Pointer(polygon)).Fgeoloop.FnumVerts == 0 {
		*(*Tint64_t)(unsafe.Pointer(out)) = 0
		return uint32(_E_SUCCESS)
	}
	// Initialize the iterator without stepping, so we can adjust the res and
	// flags (after they are validated by the initialization) before we start
	*(*TIterCellsPolygonCompact)(unsafe.Pointer(bp)) = __iterInitPolygonCompact(tls, polygon, res, flags)
	if (*(*TIterCellsPolygonCompact)(unsafe.Pointer(bp))).Ferror1 != 0 {
		return (*(*TIterCellsPolygonCompact)(unsafe.Pointer(bp))).Ferror1
	}
	// Ignore the requested flags and use the faster overlapping-bbox mode
	(*(*TIterCellsPolygonCompact)(unsafe.Pointer(bp))).F_flags = uint32(_CONTAINMENT_OVERLAPPING_BBOX)
	// Get a (very) rough area of the polygon bounding box
	polygonBBox = (*(*TIterCellsPolygonCompact)(unsafe.Pointer(bp))).F_bboxes
	polygonBBoxAreaKm2 = XbboxHeightRads(tls, polygonBBox) * XbboxWidthRads(tls, polygonBBox) / libc.Xcos(tls, libc.X__builtin_fmin(tls, libc.Xfabs(tls, (*TBBox)(unsafe.Pointer(polygonBBox)).Fnorth), libc.Xfabs(tls, (*TBBox)(unsafe.Pointer(polygonBBox)).Fsouth))) * float64(DEARTH_RADIUS_KM) * float64(DEARTH_RADIUS_KM)
	// Determine the res for the size estimate, based on a (very) rough estimate
	// of the number of cells at various resolutions that would fit in the
	// polygon. All we need here is a general order of magnitude.
	for (*(*TIterCellsPolygonCompact)(unsafe.Pointer(bp))).F_res > 0 && polygonBBoxAreaKm2/_getAverageCellArea(tls, (*(*TIterCellsPolygonCompact)(unsafe.Pointer(bp))).F_res-int32(1)) > float64(_MAX_SIZE_CELL_THRESHOLD) {
		(*(*TIterCellsPolygonCompact)(unsafe.Pointer(bp))).F_res--
	}
	// Now run the polyfill, counting the output in the target res.
	// We have to take the first step outside the loop, to get the first
	// valid output cell
	XiterStepPolygonCompact(tls, bp)
	*(*Tint64_t)(unsafe.Pointer(out)) = 0
	for {
		if !((*(*TIterCellsPolygonCompact)(unsafe.Pointer(bp))).Fcell != 0) {
			break
		}
		XcellToChildrenSize(tls, (*(*TIterCellsPolygonCompact)(unsafe.Pointer(bp))).Fcell, res, bp+48)
		*(*Tint64_t)(unsafe.Pointer(out)) += *(*Tint64_t)(unsafe.Pointer(bp + 48))
		goto _1
	_1:
		;
		XiterStepPolygonCompact(tls, bp)
	}
	return (*(*TIterCellsPolygonCompact)(unsafe.Pointer(bp))).Ferror1
}

const DM_2PI5 = 6.283185307179586

func XpointInsideGeoLoop(tls *libc.TLS, loop uintptr, bbox uintptr, coord uintptr) (r uint8) {
	var a, b, tmp TLatLng
	var aLng, bLng, lat, lng, ratio, testLng, v1, v3, v4, v5 float64
	var contains, isTransmeridian uint8
	var loopIndex, v2 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a, aLng, b, bLng, contains, isTransmeridian, lat, lng, loopIndex, ratio, testLng, tmp, v1, v2, v3, v4, v5
	if !(XbboxContains(tls, bbox, coord) != 0) {
		return uint8(Dfalse)
	}
	isTransmeridian = XbboxIsTransmeridian(tls, bbox)
	contains = uint8(Dfalse)
	lat = (*TLatLng)(unsafe.Pointer(coord)).Flat
	if isTransmeridian != 0 && (*TLatLng)(unsafe.Pointer(coord)).Flng < libc.Float64FromInt32(0) {
		v1 = (*TLatLng)(unsafe.Pointer(coord)).Flng + libc.Float64FromFloat64(6.283185307179586)
	} else {
		v1 = (*TLatLng)(unsafe.Pointer(coord)).Flng
	}
	lng = v1
	loopIndex = -int32(1)
	for int32(Dtrue) != 0 {
		loopIndex++
		v2 = loopIndex
		if v2 >= (*TGeoLoop)(unsafe.Pointer(loop)).FnumVerts {
			break
		}
		a = *(*TLatLng)(unsafe.Pointer((*TGeoLoop)(unsafe.Pointer(loop)).Fverts + uintptr(loopIndex)*16))
		b = *(*TLatLng)(unsafe.Pointer((*TGeoLoop)(unsafe.Pointer(loop)).Fverts + uintptr((loopIndex+int32(1))%(*TGeoLoop)(unsafe.Pointer(loop)).FnumVerts)*16))
		if a.Flat > b.Flat {
			tmp = a
			a = b
			b = tmp
		}
		if lat == a.Flat || lat == b.Flat {
			lat += float64(2.220446049250313e-16)
		}
		if lat < a.Flat || lat > b.Flat {
			continue
		}
		if isTransmeridian != 0 && a.Flng < libc.Float64FromInt32(0) {
			v3 = a.Flng + libc.Float64FromFloat64(6.283185307179586)
		} else {
			v3 = a.Flng
		}
		aLng = v3
		if isTransmeridian != 0 && b.Flng < libc.Float64FromInt32(0) {
			v4 = b.Flng + libc.Float64FromFloat64(6.283185307179586)
		} else {
			v4 = b.Flng
		}
		bLng = v4
		if aLng == lng || bLng == lng {
			lng -= float64(2.220446049250313e-16)
		}
		ratio = (lat - a.Flat) / (b.Flat - a.Flat)
		if isTransmeridian != 0 && aLng+(bLng-aLng)*ratio < libc.Float64FromInt32(0) {
			v5 = aLng + (bLng-aLng)*ratio + libc.Float64FromFloat64(6.283185307179586)
		} else {
			v5 = aLng + (bLng-aLng)*ratio
		}
		testLng = v5
		if testLng > lng {
			contains = libc.BoolUint8(!(contains != 0))
		}
	}
	return contains
}

func XbboxFromGeoLoop(tls *libc.TLS, loop uintptr, bbox uintptr) {
	var coord, next TLatLng
	var isTransmeridian uint8
	var lat, lng, maxNegLng, minPosLng float64
	var loopIndex, v1 int32
	_, _, _, _, _, _, _, _, _ = coord, isTransmeridian, lat, lng, loopIndex, maxNegLng, minPosLng, next, v1
	if (*TGeoLoop)(unsafe.Pointer(loop)).FnumVerts == 0 {
		*(*TBBox)(unsafe.Pointer(bbox)) = TBBox{}
		return
	}
	(*TBBox)(unsafe.Pointer(bbox)).Fsouth = float64(1.7976931348623157e+308)
	(*TBBox)(unsafe.Pointer(bbox)).Fwest = float64(1.7976931348623157e+308)
	(*TBBox)(unsafe.Pointer(bbox)).Fnorth = -libc.Float64FromFloat64(1.7976931348623157e+308)
	(*TBBox)(unsafe.Pointer(bbox)).Feast = -libc.Float64FromFloat64(1.7976931348623157e+308)
	minPosLng = float64(1.7976931348623157e+308)
	maxNegLng = -libc.Float64FromFloat64(1.7976931348623157e+308)
	isTransmeridian = uint8(Dfalse)
	loopIndex = -int32(1)
	for int32(Dtrue) != 0 {
		loopIndex++
		v1 = loopIndex
		if v1 >= (*TGeoLoop)(unsafe.Pointer(loop)).FnumVerts {
			break
		}
		coord = *(*TLatLng)(unsafe.Pointer((*TGeoLoop)(unsafe.Pointer(loop)).Fverts + uintptr(loopIndex)*16))
		next = *(*TLatLng)(unsafe.Pointer((*TGeoLoop)(unsafe.Pointer(loop)).Fverts + uintptr((loopIndex+int32(1))%(*TGeoLoop)(unsafe.Pointer(loop)).FnumVerts)*16))
		lat = coord.Flat
		lng = coord.Flng
		if lat < (*TBBox)(unsafe.Pointer(bbox)).Fsouth {
			(*TBBox)(unsafe.Pointer(bbox)).Fsouth = lat
		}
		if lng < (*TBBox)(unsafe.Pointer(bbox)).Fwest {
			(*TBBox)(unsafe.Pointer(bbox)).Fwest = lng
		}
		if lat > (*TBBox)(unsafe.Pointer(bbox)).Fnorth {
			(*TBBox)(unsafe.Pointer(bbox)).Fnorth = lat
		}
		if lng > (*TBBox)(unsafe.Pointer(bbox)).Feast {
			(*TBBox)(unsafe.Pointer(bbox)).Feast = lng
		}
		if lng > libc.Float64FromInt32(0) && lng < minPosLng {
			minPosLng = lng
		}
		if lng < libc.Float64FromInt32(0) && lng > maxNegLng {
			maxNegLng = lng
		}
		if libc.Xfabs(tls, lng-next.Flng) > float64(3.141592653589793) {
			isTransmeridian = uint8(Dtrue)
		}
	}
	if isTransmeridian != 0 {
		(*TBBox)(unsafe.Pointer(bbox)).Feast = maxNegLng
		(*TBBox)(unsafe.Pointer(bbox)).Fwest = minPosLng
	}
}

func _isClockwiseNormalizedGeoLoop(tls *libc.TLS, loop uintptr, isTransmeridian uint8) (r uint8) {
	var a, b TLatLng
	var loopIndex, v1 int32
	var sum, v2, v3 float64
	_, _, _, _, _, _, _ = a, b, loopIndex, sum, v1, v2, v3
	sum = libc.Float64FromInt32(0)
	loopIndex = -int32(1)
	for int32(Dtrue) != 0 {
		loopIndex++
		v1 = loopIndex
		if v1 >= (*TGeoLoop)(unsafe.Pointer(loop)).FnumVerts {
			break
		}
		a = *(*TLatLng)(unsafe.Pointer((*TGeoLoop)(unsafe.Pointer(loop)).Fverts + uintptr(loopIndex)*16))
		b = *(*TLatLng)(unsafe.Pointer((*TGeoLoop)(unsafe.Pointer(loop)).Fverts + uintptr((loopIndex+int32(1))%(*TGeoLoop)(unsafe.Pointer(loop)).FnumVerts)*16))
		if !(isTransmeridian != 0) && libc.Xfabs(tls, a.Flng-b.Flng) > float64(3.141592653589793) {
			return _isClockwiseNormalizedGeoLoop(tls, loop, uint8(Dtrue))
		}
		if isTransmeridian != 0 && b.Flng < libc.Float64FromInt32(0) {
			v2 = b.Flng + libc.Float64FromFloat64(6.283185307179586)
		} else {
			v2 = b.Flng
		}
		if isTransmeridian != 0 && a.Flng < libc.Float64FromInt32(0) {
			v3 = a.Flng + libc.Float64FromFloat64(6.283185307179586)
		} else {
			v3 = a.Flng
		}
		sum += (v2 - v3) * (b.Flat + a.Flat)
	}
	return libc.BoolUint8(sum > libc.Float64FromInt32(0))
}

func XisClockwiseGeoLoop(tls *libc.TLS, loop uintptr) (r uint8) {
	return _isClockwiseNormalizedGeoLoop(tls, loop, uint8(Dfalse))
}

// C documentation
//
//	/**
//	 * Whether the flags for the polyfill operation are valid
//	 * TODO: Move to polyfill.c when the old algo is removed
//	 * @param  flags Flags to validate
//	 * @return       Whether the flags are valid
//	 */
func XvalidatePolygonFlags(tls *libc.TLS, flags Tuint32_t) (r TH3Error) {
	if flags & ^libc.Uint32FromInt32(libc.Int32FromInt32(15)) != 0 || flags&libc.Uint32FromInt32(libc.Int32FromInt32(15)) >= uint32(_CONTAINMENT_INVALID) {
		return uint32(_E_OPTION_INVALID)
	}
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * Create a bounding box from a GeoPolygon
//	 * @param polygon Input GeoPolygon
//	 * @param bboxes  Output bboxes, one for the outer loop and one for each hole
//	 */
func XbboxesFromGeoPolygon(tls *libc.TLS, polygon uintptr, bboxes uintptr) {
	var i int32
	_ = i
	XbboxFromGeoLoop(tls, polygon, bboxes)
	i = 0
	for {
		if !(i < (*TGeoPolygon)(unsafe.Pointer(polygon)).FnumHoles) {
			break
		}
		XbboxFromGeoLoop(tls, (*TGeoPolygon)(unsafe.Pointer(polygon)).Fholes+uintptr(i)*16, bboxes+uintptr(i+int32(1))*32)
		goto _1
	_1:
		;
		i++
	}
}

// C documentation
//
//	/**
//	 * pointInsidePolygon takes a given GeoPolygon data structure and
//	 * checks if it contains a given geo coordinate.
//	 *
//	 * @param geoPolygon The geoloop and holes defining the relevant area
//	 * @param bboxes     The bboxes for the main geoloop and each of its holes
//	 * @param coord      The coordinate to check
//	 * @return           Whether the point is contained
//	 */
func XpointInsidePolygon(tls *libc.TLS, geoPolygon uintptr, bboxes uintptr, coord uintptr) (r uint8) {
	var contains uint8
	var i int32
	_, _ = contains, i
	// Start with contains state of primary geoloop
	contains = XpointInsideGeoLoop(tls, geoPolygon, bboxes, coord)
	// If the point is contained in the primary geoloop, but there are holes in
	// the geoloop iterate through all holes and return false if the point is
	// contained in any hole
	if contains != 0 && (*TGeoPolygon)(unsafe.Pointer(geoPolygon)).FnumHoles > 0 {
		i = 0
		for {
			if !(i < (*TGeoPolygon)(unsafe.Pointer(geoPolygon)).FnumHoles) {
				break
			}
			if XpointInsideGeoLoop(tls, (*TGeoPolygon)(unsafe.Pointer(geoPolygon)).Fholes+uintptr(i)*16, bboxes+uintptr(i+int32(1))*32, coord) != 0 {
				return uint8(Dfalse)
			}
			goto _1
		_1:
			;
			i++
		}
	}
	return contains
}

// C documentation
//
//	/**
//	 * Whether a cell boundary is completely contained by a polygon.
//	 * @param  geoPolygon The polygon to test
//	 * @param  bboxes     The bboxes for the main geoloop and each of its holes
//	 * @param  boundary   The cell boundary to test
//	 * @return            Whether the cell boundary is contained
//	 */
func XcellBoundaryInsidePolygon(tls *libc.TLS, geoPolygon uintptr, bboxes uintptr, boundary uintptr, boundaryBBox uintptr) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i int32
	var _ /* boundaryLoop at bp+0 */ TGeoLoop
	_ = i
	// First test a single point. Note that this fails fast if point is outside
	// bounding box.
	if !(XpointInsidePolygon(tls, geoPolygon, bboxes, boundary+8) != 0) {
		return uint8(Dfalse)
	}
	// If a point is contained, check for any line intersections
	if XcellBoundaryCrossesGeoLoop(tls, geoPolygon, bboxes, boundary, boundaryBBox) != 0 {
		return uint8(Dfalse)
	}
	// Convert boundary to geoloop for point-inside check
	*(*TGeoLoop)(unsafe.Pointer(bp)) = TGeoLoop{
		FnumVerts: (*TCellBoundary)(unsafe.Pointer(boundary)).FnumVerts,
		Fverts:    boundary + 8,
	}
	// Check for line intersections with, or containment of, any hole
	i = 0
	for {
		if !(i < (*TGeoPolygon)(unsafe.Pointer(geoPolygon)).FnumHoles) {
			break
		}
		// If the hole has no verts, it is not possible to intersect with it.
		if (*(*TGeoLoop)(unsafe.Pointer((*TGeoPolygon)(unsafe.Pointer(geoPolygon)).Fholes + uintptr(i)*16))).FnumVerts > 0 && (XpointInsideGeoLoop(tls, bp, boundaryBBox, (*(*TGeoLoop)(unsafe.Pointer((*TGeoPolygon)(unsafe.Pointer(geoPolygon)).Fholes + uintptr(i)*16))).Fverts) != 0 || XcellBoundaryCrossesGeoLoop(tls, (*TGeoPolygon)(unsafe.Pointer(geoPolygon)).Fholes+uintptr(i)*16, bboxes+uintptr(i+int32(1))*32, boundary, boundaryBBox) != 0) {
			return uint8(Dfalse)
		}
		goto _1
	_1:
		;
		i++
	}
	return uint8(Dtrue)
}

// C documentation
//
//	/**
//	 * Whether any part of a cell boundary crosses a polygon. Crossing in this case
//	 * means whether any line segments intersect; it does not include containment.
//	 * @param  geoPolygon The polygon to test
//	 * @param  bboxes     The bboxes for the main geoloop and each of its holes
//	 * @param  boundary   The cell boundary to test
//	 * @return            Whether the cell boundary is contained
//	 */
func XcellBoundaryCrossesPolygon(tls *libc.TLS, geoPolygon uintptr, bboxes uintptr, boundary uintptr, boundaryBBox uintptr) (r uint8) {
	var i int32
	_ = i
	// Check for line intersections with outer loop
	if XcellBoundaryCrossesGeoLoop(tls, geoPolygon, bboxes, boundary, boundaryBBox) != 0 {
		return uint8(Dtrue)
	}
	// Check for line intersections with any hole
	i = 0
	for {
		if !(i < (*TGeoPolygon)(unsafe.Pointer(geoPolygon)).FnumHoles) {
			break
		}
		if XcellBoundaryCrossesGeoLoop(tls, (*TGeoPolygon)(unsafe.Pointer(geoPolygon)).Fholes+uintptr(i)*16, bboxes+uintptr(i+int32(1))*32, boundary, boundaryBBox) != 0 {
			return uint8(Dtrue)
		}
		goto _1
	_1:
		;
		i++
	}
	return uint8(Dfalse)
}

// C documentation
//
//	/**
//	 * Whether a cell boundary crosses a geo loop. Crossing in this case means
//	 * whether any line segments intersect; it does not include containment.
//	 * @param  geoloop  Geo loop to test
//	 * @param  boundary Cell boundary to test
//	 * @return          Whether any line segments in the boundary intersect any line
//	 * segments in the geo loop
//	 */
func XcellBoundaryCrossesGeoLoop(tls *libc.TLS, geoloop uintptr, loopBBox uintptr, boundary uintptr, boundaryBBox uintptr) (r uint8) {
	bp := tls.Alloc(208)
	defer tls.Free(208)
	var i, i1, j int32
	var normalBoundaryBBox TBBox
	var _ /* boundaryNormalization at bp+4 */ TLongitudeNormalization
	var _ /* loop1 at bp+176 */ TLatLng
	var _ /* loop2 at bp+192 */ TLatLng
	var _ /* loopNormalization at bp+0 */ TLongitudeNormalization
	var _ /* normalBoundary at bp+8 */ TCellBoundary
	_, _, _, _ = i, i1, j, normalBoundaryBBox
	if !(XbboxOverlapsBBox(tls, loopBBox, boundaryBBox) != 0) {
		return uint8(Dfalse)
	}
	XbboxNormalization(tls, loopBBox, boundaryBBox, bp, bp+4)
	*(*TCellBoundary)(unsafe.Pointer(bp + 8)) = *(*TCellBoundary)(unsafe.Pointer(boundary))
	i = 0
	for {
		if !(i < (*TCellBoundary)(unsafe.Pointer(boundary)).FnumVerts) {
			break
		}
		(*(*TLatLng)(unsafe.Pointer(bp + 8 + 8 + uintptr(i)*16))).Flng = XnormalizeLng(tls, (*(*TLatLng)(unsafe.Pointer(bp + 8 + 8 + uintptr(i)*16))).Flng, *(*TLongitudeNormalization)(unsafe.Pointer(bp + 4)))
		goto _1
	_1:
		;
		i++
	}
	normalBoundaryBBox = TBBox{
		Fnorth: (*TBBox)(unsafe.Pointer(boundaryBBox)).Fnorth,
		Fsouth: (*TBBox)(unsafe.Pointer(boundaryBBox)).Fsouth,
		Feast:  XnormalizeLng(tls, (*TBBox)(unsafe.Pointer(boundaryBBox)).Feast, *(*TLongitudeNormalization)(unsafe.Pointer(bp + 4))),
		Fwest:  XnormalizeLng(tls, (*TBBox)(unsafe.Pointer(boundaryBBox)).Fwest, *(*TLongitudeNormalization)(unsafe.Pointer(bp + 4))),
	}
	i1 = 0
	for {
		if !(i1 < (*TGeoLoop)(unsafe.Pointer(geoloop)).FnumVerts) {
			break
		}
		*(*TLatLng)(unsafe.Pointer(bp + 176)) = *(*TLatLng)(unsafe.Pointer((*TGeoLoop)(unsafe.Pointer(geoloop)).Fverts + uintptr(i1)*16))
		(*(*TLatLng)(unsafe.Pointer(bp + 176))).Flng = XnormalizeLng(tls, (*(*TLatLng)(unsafe.Pointer(bp + 176))).Flng, *(*TLongitudeNormalization)(unsafe.Pointer(bp)))
		*(*TLatLng)(unsafe.Pointer(bp + 192)) = *(*TLatLng)(unsafe.Pointer((*TGeoLoop)(unsafe.Pointer(geoloop)).Fverts + uintptr((i1+int32(1))%(*TGeoLoop)(unsafe.Pointer(geoloop)).FnumVerts)*16))
		(*(*TLatLng)(unsafe.Pointer(bp + 192))).Flng = XnormalizeLng(tls, (*(*TLatLng)(unsafe.Pointer(bp + 192))).Flng, *(*TLongitudeNormalization)(unsafe.Pointer(bp)))
		// Quick check if the line segment overlaps our bbox
		if (*(*TLatLng)(unsafe.Pointer(bp + 176))).Flat >= normalBoundaryBBox.Fnorth && (*(*TLatLng)(unsafe.Pointer(bp + 192))).Flat >= normalBoundaryBBox.Fnorth || (*(*TLatLng)(unsafe.Pointer(bp + 176))).Flat <= normalBoundaryBBox.Fsouth && (*(*TLatLng)(unsafe.Pointer(bp + 192))).Flat <= normalBoundaryBBox.Fsouth || (*(*TLatLng)(unsafe.Pointer(bp + 176))).Flng <= normalBoundaryBBox.Fwest && (*(*TLatLng)(unsafe.Pointer(bp + 192))).Flng <= normalBoundaryBBox.Fwest || (*(*TLatLng)(unsafe.Pointer(bp + 176))).Flng >= normalBoundaryBBox.Feast && (*(*TLatLng)(unsafe.Pointer(bp + 192))).Flng >= normalBoundaryBBox.Feast {
			goto _2
		}
		j = 0
		for {
			if !(j < (*(*TCellBoundary)(unsafe.Pointer(bp + 8))).FnumVerts) {
				break
			}
			if XlineCrossesLine(tls, bp+176, bp+192, bp+8+8+uintptr(j)*16, bp+8+8+uintptr((j+int32(1))%(*(*TCellBoundary)(unsafe.Pointer(bp + 8))).FnumVerts)*16) != 0 {
				return uint8(Dtrue)
			}
			goto _3
		_3:
			;
			j++
		}
		goto _2
	_2:
		;
		i1++
	}
	return uint8(Dfalse)
}

// C documentation
//
//	/**
//	 * Whether two lines intersect. This is a purely Cartesian implementation
//	 * and does not consider anti-meridian wrapping, poles, etc. Based on
//	 * http://www.jeffreythompson.org/collision-detection/line-line.php
//	 * @param  a1 Start of line A
//	 * @param  a2 End of line A
//	 * @param  b1 Start of line B
//	 * @param  b2 End of line B
//	 * @return    Whether the lines intersect
//	 */
func XlineCrossesLine(tls *libc.TLS, a1 uintptr, a2 uintptr, b1 uintptr, b2 uintptr) (r uint8) {
	var denom, test float64
	_, _ = denom, test
	denom = ((*TLatLng)(unsafe.Pointer(b2)).Flng-(*TLatLng)(unsafe.Pointer(b1)).Flng)*((*TLatLng)(unsafe.Pointer(a2)).Flat-(*TLatLng)(unsafe.Pointer(a1)).Flat) - ((*TLatLng)(unsafe.Pointer(b2)).Flat-(*TLatLng)(unsafe.Pointer(b1)).Flat)*((*TLatLng)(unsafe.Pointer(a2)).Flng-(*TLatLng)(unsafe.Pointer(a1)).Flng)
	if !(denom != 0) {
		return uint8(Dfalse)
	}
	test = (((*TLatLng)(unsafe.Pointer(b2)).Flat-(*TLatLng)(unsafe.Pointer(b1)).Flat)*((*TLatLng)(unsafe.Pointer(a1)).Flng-(*TLatLng)(unsafe.Pointer(b1)).Flng) - ((*TLatLng)(unsafe.Pointer(b2)).Flng-(*TLatLng)(unsafe.Pointer(b1)).Flng)*((*TLatLng)(unsafe.Pointer(a1)).Flat-(*TLatLng)(unsafe.Pointer(b1)).Flat)) / denom
	if test < libc.Float64FromInt32(0) || test > libc.Float64FromInt32(1) {
		return uint8(Dfalse)
	}
	test = (((*TLatLng)(unsafe.Pointer(a2)).Flat-(*TLatLng)(unsafe.Pointer(a1)).Flat)*((*TLatLng)(unsafe.Pointer(a1)).Flng-(*TLatLng)(unsafe.Pointer(b1)).Flng) - ((*TLatLng)(unsafe.Pointer(a2)).Flng-(*TLatLng)(unsafe.Pointer(a1)).Flng)*((*TLatLng)(unsafe.Pointer(a1)).Flat-(*TLatLng)(unsafe.Pointer(b1)).Flat)) / denom
	return libc.BoolUint8(test >= libc.Float64FromInt32(0) && test <= libc.Float64FromInt32(1))
}

const DDBL_EPSILON2 = 2.22044604925031308085e-16
const DDBL_MAX2 = 1.79769313486231570815e+308
const DM_PI6 = 3.14159265358979323846

// C documentation
//
//	/**
//	 * Calculates the magnitude of a 2D cartesian vector.
//	 * @param v The 2D cartesian vector.
//	 * @return The magnitude of the vector.
//	 */
func X_v2dMag(tls *libc.TLS, v uintptr) (r float64) {
	return libc.Xsqrt(tls, (*TVec2d)(unsafe.Pointer(v)).Fx*(*TVec2d)(unsafe.Pointer(v)).Fx+(*TVec2d)(unsafe.Pointer(v)).Fy*(*TVec2d)(unsafe.Pointer(v)).Fy)
}

// C documentation
//
//	/**
//	 * Finds the intersection between two lines. Assumes that the lines intersect
//	 * and that the intersection is not at an endpoint of either line.
//	 * @param p0 The first endpoint of the first line.
//	 * @param p1 The second endpoint of the first line.
//	 * @param p2 The first endpoint of the second line.
//	 * @param p3 The second endpoint of the second line.
//	 * @param inter The intersection point.
//	 */
func X_v2dIntersect(tls *libc.TLS, p0 uintptr, p1 uintptr, p2 uintptr, p3 uintptr, inter uintptr) {
	var s1, s2 TVec2d
	var t float64
	_, _, _ = s1, s2, t
	s1.Fx = (*TVec2d)(unsafe.Pointer(p1)).Fx - (*TVec2d)(unsafe.Pointer(p0)).Fx
	s1.Fy = (*TVec2d)(unsafe.Pointer(p1)).Fy - (*TVec2d)(unsafe.Pointer(p0)).Fy
	s2.Fx = (*TVec2d)(unsafe.Pointer(p3)).Fx - (*TVec2d)(unsafe.Pointer(p2)).Fx
	s2.Fy = (*TVec2d)(unsafe.Pointer(p3)).Fy - (*TVec2d)(unsafe.Pointer(p2)).Fy
	t = (s2.Fx*((*TVec2d)(unsafe.Pointer(p0)).Fy-(*TVec2d)(unsafe.Pointer(p2)).Fy) - s2.Fy*((*TVec2d)(unsafe.Pointer(p0)).Fx-(*TVec2d)(unsafe.Pointer(p2)).Fx)) / (-s2.Fx*s1.Fy + s1.Fx*s2.Fy)
	(*TVec2d)(unsafe.Pointer(inter)).Fx = (*TVec2d)(unsafe.Pointer(p0)).Fx + t*s1.Fx
	(*TVec2d)(unsafe.Pointer(inter)).Fy = (*TVec2d)(unsafe.Pointer(p0)).Fy + t*s1.Fy
}

// C documentation
//
//	/**
//	 * Whether two 2D vectors are almost equal, within some threshold
//	 * @param v1 First vector to compare
//	 * @param v2 Second vector to compare
//	 * @return Whether the vectors are almost equal
//	 */
func X_v2dAlmostEquals(tls *libc.TLS, v1 uintptr, v2 uintptr) (r uint8) {
	return libc.BoolUint8(libc.Xfabs(tls, (*TVec2d)(unsafe.Pointer(v1)).Fx-(*TVec2d)(unsafe.Pointer(v2)).Fx) < libc.Float64FromFloat32(1.1920928955078125e-07) && libc.Xfabs(tls, (*TVec2d)(unsafe.Pointer(v1)).Fy-(*TVec2d)(unsafe.Pointer(v2)).Fy) < libc.Float64FromFloat32(1.1920928955078125e-07))
}

const DM_2PI6 = 6.28318530717958647692528676655900576839433

// C documentation
//
//	/**
//	 * Square of a number
//	 *
//	 * @param x The input number.
//	 * @return The square of the input number.
//	 */
func X_square(tls *libc.TLS, x float64) (r float64) {
	return x * x
}

// C documentation
//
//	/**
//	 * Calculate the square of the distance between two 3D coordinates.
//	 *
//	 * @param v1 The first 3D coordinate.
//	 * @param v2 The second 3D coordinate.
//	 * @return The square of the distance between the given points.
//	 */
func X_pointSquareDist(tls *libc.TLS, v1 uintptr, v2 uintptr) (r float64) {
	return X_square(tls, (*TVec3d)(unsafe.Pointer(v1)).Fx-(*TVec3d)(unsafe.Pointer(v2)).Fx) + X_square(tls, (*TVec3d)(unsafe.Pointer(v1)).Fy-(*TVec3d)(unsafe.Pointer(v2)).Fy) + X_square(tls, (*TVec3d)(unsafe.Pointer(v1)).Fz-(*TVec3d)(unsafe.Pointer(v2)).Fz)
}

// C documentation
//
//	/**
//	 * Calculate the 3D coordinate on unit sphere from the latitude and longitude.
//	 *
//	 * @param geo The latitude and longitude of the point.
//	 * @param v The 3D coordinate of the point.
//	 */
func X_geoToVec3d(tls *libc.TLS, geo uintptr, v uintptr) {
	var r float64
	_ = r
	r = libc.Xcos(tls, (*TLatLng)(unsafe.Pointer(geo)).Flat)
	(*TVec3d)(unsafe.Pointer(v)).Fz = libc.Xsin(tls, (*TLatLng)(unsafe.Pointer(geo)).Flat)
	(*TVec3d)(unsafe.Pointer(v)).Fx = libc.Xcos(tls, (*TLatLng)(unsafe.Pointer(geo)).Flng) * r
	(*TVec3d)(unsafe.Pointer(v)).Fy = libc.Xsin(tls, (*TLatLng)(unsafe.Pointer(geo)).Flng) * r
}

const DDIRECTION_INDEX_OFFSET = 2
const DM_PI_22 = 1.5707963267948966

var _UNIT_VECS10 = [7]TCoordIJK{
	0: {},
	1: {
		Fk: int32(1),
	},
	2: {
		Fj: int32(1),
	},
	3: {
		Fj: int32(1),
		Fk: int32(1),
	},
	4: {
		Fi: int32(1),
	},
	5: {
		Fi: int32(1),
		Fk: int32(1),
	},
	6: {
		Fi: int32(1),
		Fj: int32(1),
	},
}

/*
** The testcase() macro is used to aid in coverage testing.  When
** doing coverage testing, the condition inside the argument to
** testcase() must be evaluated both true and false in order to
** get full branch coverage.  The testcase() macro is inserted
** to help ensure adequate test coverage in places where simple
** condition/decision coverage is inadequate.  For example, testcase()
** can be used to make sure boundary values are tested.  For
** bitmask tests, testcase() can be used to make sure each bit
** is significant and used at least once.  On switch statements
** where multiple cases go to the same block of code, testcase()
** can insure that all cases are evaluated.
 */

/*
** Disable ALWAYS() and NEVER() (make them pass-throughs) for coverage
** and mutation testing
 */

/*
** The TESTONLY macro is used to enclose variable declarations or
** other bits of code that are needed to support the arguments
** within testcase() and assert() macros.
 */

/*
** The DEFENSEONLY macro is used to enclose variable declarations or
** other bits of code that are needed to support the arguments
** within ALWAYS() or NEVER() macros.
 */

/*
** The ALWAYS and NEVER macros surround boolean expressions which
** are intended to always be true or false, respectively.  Such
** expressions could be omitted from the code completely.  But they
** are included in a few cases in order to enhance the resilience
** of the H3 library to unexpected behavior - to make the code "self-healing"
** or "ductile" rather than being "brittle" and crashing at the first
** hint of unplanned behavior.
**
** In other words, ALWAYS and NEVER are added for defensive code.
**
** When doing coverage testing ALWAYS and NEVER are hard-coded to
** be true and false so that the unreachable code they specify will
** not be counted as untested code.
 */

/*
 * Copyright 2016-2018, 2020 Uber Technologies, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/** @file h3Index.h
 * @brief   H3Index functions.
 */

/*
 * Copyright 2016-2021 Uber Technologies, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/** @file latLng.h
 * @brief   Geodetic (lat/lng) functions.
 */

// C documentation
//
//	/** @brief Table of direction-to-face mapping for each pentagon
//	 *
//	 * Note that faces are in directional order, starting at J_AXES_DIGIT.
//	 * This table is generated by the generatePentagonDirectionFaces script.
//	 */
var _pentagonDirectionFaces = [12]TPentagonDirectionFaces{
	0: {
		FbaseCell: int32(4),
		Ffaces: [5]int32{
			0: int32(4),
			2: int32(2),
			3: int32(1),
			4: int32(3),
		},
	},
	1: {
		FbaseCell: int32(14),
		Ffaces: [5]int32{
			0: int32(6),
			1: int32(11),
			2: int32(2),
			3: int32(7),
			4: int32(1),
		},
	},
	2: {
		FbaseCell: int32(24),
		Ffaces: [5]int32{
			0: int32(5),
			1: int32(10),
			2: int32(1),
			3: int32(6),
		},
	},
	3: {
		FbaseCell: int32(38),
		Ffaces: [5]int32{
			0: int32(7),
			1: int32(12),
			2: int32(3),
			3: int32(8),
			4: int32(2),
		},
	},
	4: {
		FbaseCell: int32(49),
		Ffaces: [5]int32{
			0: int32(9),
			1: int32(14),
			3: int32(5),
			4: int32(4),
		},
	},
	5: {
		FbaseCell: int32(58),
		Ffaces: [5]int32{
			0: int32(8),
			1: int32(13),
			2: int32(4),
			3: int32(9),
			4: int32(3),
		},
	},
	6: {
		FbaseCell: int32(63),
		Ffaces: [5]int32{
			0: int32(11),
			1: int32(6),
			2: int32(15),
			3: int32(10),
			4: int32(16),
		},
	},
	7: {
		FbaseCell: int32(72),
		Ffaces: [5]int32{
			0: int32(12),
			1: int32(7),
			2: int32(16),
			3: int32(11),
			4: int32(17),
		},
	},
	8: {
		FbaseCell: int32(83),
		Ffaces: [5]int32{
			0: int32(10),
			1: int32(5),
			2: int32(19),
			3: int32(14),
			4: int32(15),
		},
	},
	9: {
		FbaseCell: int32(97),
		Ffaces: [5]int32{
			0: int32(13),
			1: int32(8),
			2: int32(17),
			3: int32(12),
			4: int32(18),
		},
	},
	10: {
		FbaseCell: int32(107),
		Ffaces: [5]int32{
			0: int32(14),
			1: int32(9),
			2: int32(18),
			3: int32(13),
			4: int32(19),
		},
	},
	11: {
		FbaseCell: int32(117),
		Ffaces: [5]int32{
			0: int32(15),
			1: int32(19),
			2: int32(17),
			3: int32(18),
			4: int32(16),
		},
	},
}

// C documentation
//
//	/**
//	 * Get the number of CCW rotations of the cell's vertex numbers
//	 * compared to the directional layout of its neighbors.
//	 * @param out Number of CCW rotations for the cell
//	 */
func _vertexRotations(tls *libc.TLS, cell TH3Index, out uintptr) (r TH3Error) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var baseCell, ccwRot60, cellLeadingDigit, p int32
	var err TH3Error
	var _ /* baseFijk at bp+16 */ TFaceIJK
	var _ /* dirFaces at bp+32 */ TPentagonDirectionFaces
	var _ /* fijk at bp+0 */ TFaceIJK
	_, _, _, _, _ = baseCell, ccwRot60, cellLeadingDigit, err, p
	err = X_h3ToFaceIjk(tls, cell, bp)
	if err != 0 {
		return err
	}
	baseCell = XgetBaseCellNumber(tls, cell)
	cellLeadingDigit = X_h3LeadingNonZeroDigit(tls, cell)
	X_baseCellToFaceIjk(tls, baseCell, bp+16)
	ccwRot60 = X_baseCellToCCWrot60(tls, baseCell, (*(*TFaceIJK)(unsafe.Pointer(bp))).Fface)
	if X_isBaseCellPentagon(tls, baseCell) != 0 {
		// We never hit the end condition
		p = 0
		for {
			if !(p < int32(DNUM_PENTAGONS)) {
				break
			}
			if _pentagonDirectionFaces[p].FbaseCell == baseCell {
				*(*TPentagonDirectionFaces)(unsafe.Pointer(bp + 32)) = _pentagonDirectionFaces[p]
				break
			}
			goto _1
		_1:
			;
			p++
		}
		if p == int32(DNUM_PENTAGONS) {
			return uint32(_E_FAILED)
		}
		// additional CCW rotation for polar neighbors or IK neighbors
		if (*(*TFaceIJK)(unsafe.Pointer(bp))).Fface != (*(*TFaceIJK)(unsafe.Pointer(bp + 16))).Fface && (X_isBaseCellPolarPentagon(tls, baseCell) != 0 || (*(*TFaceIJK)(unsafe.Pointer(bp))).Fface == *(*int32)(unsafe.Pointer(bp + 32 + 4 + uintptr(int32(_IK_AXES_DIGIT)-libc.Int32FromInt32(DDIRECTION_INDEX_OFFSET))*4))) {
			ccwRot60 = (ccwRot60 + int32(1)) % int32(6)
		}
		// Check whether the cell crosses a deleted pentagon subsequence
		if cellLeadingDigit == int32(_JK_AXES_DIGIT) && (*(*TFaceIJK)(unsafe.Pointer(bp))).Fface == *(*int32)(unsafe.Pointer(bp + 32 + 4 + uintptr(int32(_IK_AXES_DIGIT)-libc.Int32FromInt32(DDIRECTION_INDEX_OFFSET))*4)) {
			// Crosses from JK to IK: Rotate CW
			ccwRot60 = (ccwRot60 + int32(5)) % int32(6)
		} else {
			if cellLeadingDigit == int32(_IK_AXES_DIGIT) && (*(*TFaceIJK)(unsafe.Pointer(bp))).Fface == *(*int32)(unsafe.Pointer(bp + 32 + 4 + uintptr(int32(_JK_AXES_DIGIT)-libc.Int32FromInt32(DDIRECTION_INDEX_OFFSET))*4)) {
				// Crosses from IK to JK: Rotate CCW
				ccwRot60 = (ccwRot60 + int32(1)) % int32(6)
			}
		}
	}
	*(*int32)(unsafe.Pointer(out)) = ccwRot60
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/** @brief Hexagon direction to vertex number relationships (same face).
//	 *         Note that we don't use direction 0 (center).
//	 */
var _directionToVertexNumHex = [7]int32{
	0: int32(_INVALID_DIGIT),
	1: int32(3),
	2: int32(1),
	3: int32(2),
	4: int32(5),
	5: int32(4),
}

// C documentation
//
//	/** @brief Pentagon direction to vertex number relationships (same face).
//	 *         Note that we don't use directions 0 (center) or 1 (deleted K axis).
//	 */
var _directionToVertexNumPent = [7]int32{
	0: int32(_INVALID_DIGIT),
	1: int32(_INVALID_DIGIT),
	2: int32(1),
	3: int32(2),
	4: int32(4),
	5: int32(3),
}

// C documentation
//
//	/**
//	 * Get the first vertex number for a given direction. The neighbor in this
//	 * direction is located between this vertex number and the next number in
//	 * sequence.
//	 * @returns The number for the first topological vertex, or INVALID_VERTEX_NUM
//	 *          if the direction is not valid for this cell
//	 */
func XvertexNumForDirection(tls *libc.TLS, origin TH3Index, direction TDirection) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var err TH3Error
	var isPent int32
	var _ /* rotations at bp+0 */ int32
	_, _ = err, isPent
	isPent = XisPentagon(tls, origin)
	// Check for invalid directions
	if direction == int32(_CENTER_DIGIT) || direction >= int32(_INVALID_DIGIT) || isPent != 0 && direction == int32(_K_AXES_DIGIT) {
		return -int32(1)
	}
	err = _vertexRotations(tls, origin, bp)
	if err != 0 {
		return -int32(1)
	}
	// Find the appropriate vertex, rotating CCW if necessary
	if isPent != 0 {
		return (_directionToVertexNumPent[direction] + int32(DNUM_PENT_VERTS) - *(*int32)(unsafe.Pointer(bp))) % int32(DNUM_PENT_VERTS)
	} else {
		return (_directionToVertexNumHex[direction] + int32(DNUM_HEX_VERTS) - *(*int32)(unsafe.Pointer(bp))) % int32(DNUM_HEX_VERTS)
	}
	return r
}

// C documentation
//
//	/** @brief Vertex number to hexagon direction relationships (same face).
//	 */
var _vertexNumToDirectionHex = [6]TDirection{
	0: int32(_IJ_AXES_DIGIT),
	1: int32(_J_AXES_DIGIT),
	2: int32(_JK_AXES_DIGIT),
	3: int32(_K_AXES_DIGIT),
	4: int32(_IK_AXES_DIGIT),
	5: int32(_I_AXES_DIGIT),
}

// C documentation
//
//	/** @brief Vertex number to pentagon direction relationships (same face).
//	 */
var _vertexNumToDirectionPent = [5]TDirection{
	0: int32(_IJ_AXES_DIGIT),
	1: int32(_J_AXES_DIGIT),
	2: int32(_JK_AXES_DIGIT),
	3: int32(_IK_AXES_DIGIT),
	4: int32(_I_AXES_DIGIT),
}

// C documentation
//
//	/**
//	 * Get the direction for a given vertex number. This returns the direction for
//	 * the neighbor between the given vertex number and the next number in sequence.
//	 * @returns The direction for this vertex, or INVALID_DIGIT if the vertex
//	 * number is invalid.
//	 */
func XdirectionForVertexNum(tls *libc.TLS, origin TH3Index, vertexNum int32) (r TDirection) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var err TH3Error
	var isPent, v1, v3 int32
	var v2 bool
	var _ /* rotations at bp+0 */ int32
	_, _, _, _, _ = err, isPent, v1, v2, v3
	isPent = XisPentagon(tls, origin)
	// Check for invalid vertexes
	if v2 = vertexNum < 0; !v2 {
		if isPent != 0 {
			v1 = int32(DNUM_PENT_VERTS)
		} else {
			v1 = int32(DNUM_HEX_VERTS)
		}
	}
	if v2 || vertexNum > v1-int32(1) {
		return int32(_INVALID_DIGIT)
	}
	err = _vertexRotations(tls, origin, bp)
	if err != 0 {
		return int32(_INVALID_DIGIT)
	}
	// Find the appropriate direction, rotating CW if necessary
	if isPent != 0 {
		v3 = _vertexNumToDirectionPent[(vertexNum+*(*int32)(unsafe.Pointer(bp)))%int32(DNUM_PENT_VERTS)]
	} else {
		v3 = _vertexNumToDirectionHex[(vertexNum+*(*int32)(unsafe.Pointer(bp)))%int32(DNUM_HEX_VERTS)]
	}
	return v3
}

// C documentation
//
//	/** @brief Directions in CCW order */
var _DIRECTIONS1 = [6]TDirection{
	0: int32(_J_AXES_DIGIT),
	1: int32(_JK_AXES_DIGIT),
	2: int32(_K_AXES_DIGIT),
	3: int32(_IK_AXES_DIGIT),
	4: int32(_I_AXES_DIGIT),
	5: int32(_IJ_AXES_DIGIT),
}

// C documentation
//
//	/** @brief Reverse direction from neighbor in each direction,
//	 *         given as an index into DIRECTIONS to facilitate rotation
//	 */
var _revNeighborDirectionsHex = [7]int32{
	0: int32(_INVALID_DIGIT),
	1: int32(5),
	2: int32(3),
	3: int32(4),
	4: int32(1),
	6: int32(2),
}

// C documentation
//
//	/**
//	 * Get a single vertex for a given cell, as an H3 index, or
//	 * H3_NULL if the vertex is invalid
//	 * @param cell    Cell to get the vertex for
//	 * @param vertexNum Number (index) of the vertex to calculate
//	 * @param out Output: The vertex index
//	 */
func XcellToVertex(tls *libc.TLS, cell TH3Index, vertexNum int32, out uintptr) (r TH3Error) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var cellIsPentagon, cellNumVerts, ownerIsPentagon, ownerVertexNum, res, v1, v2, v4, v5 int32
	var dir, dir1, left, right TDirection
	var leftNeighborError, rightNeighborError TH3Error
	var owner, vertex TH3Index
	var v3 bool
	var _ /* lRotations at bp+0 */ int32
	var _ /* leftNeighbor at bp+8 */ TH3Index
	var _ /* rRotations at bp+16 */ int32
	var _ /* rightNeighbor at bp+24 */ TH3Index
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = cellIsPentagon, cellNumVerts, dir, dir1, left, leftNeighborError, owner, ownerIsPentagon, ownerVertexNum, res, right, rightNeighborError, vertex, v1, v2, v3, v4, v5
	cellIsPentagon = XisPentagon(tls, cell)
	if cellIsPentagon != 0 {
		v1 = int32(DNUM_PENT_VERTS)
	} else {
		v1 = int32(DNUM_HEX_VERTS)
	}
	cellNumVerts = v1
	res = libc.Int32FromUint64(cell & (libc.Uint64FromUint64(15) << libc.Int32FromInt32(DH3_RES_OFFSET)) >> libc.Int32FromInt32(DH3_RES_OFFSET))
	// Check for invalid vertexes
	if vertexNum < 0 || vertexNum > cellNumVerts-int32(1) {
		return uint32(_E_DOMAIN)
	}
	// Default the owner and vertex number to the input cell
	owner = cell
	ownerVertexNum = vertexNum
	// Determine the owner, looking at the three cells that share the vertex.
	// By convention, the owner is the cell with the lowest numerical index.
	// If the cell is the center child of its parent, it will always have
	// the lowest index of any neighbor, so we can skip determining the owner
	if res == 0 || libc.Int32FromUint64(cell>>((libc.Int32FromInt32(DMAX_H3_RES)-res)*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))&libc.Uint64FromInt32(libc.Int32FromInt32(7))) != int32(_CENTER_DIGIT) {
		// Get the left neighbor of the vertex, with its rotations
		left = XdirectionForVertexNum(tls, cell, vertexNum)
		if left == int32(_INVALID_DIGIT) {
			return uint32(_E_FAILED)
		}
		*(*int32)(unsafe.Pointer(bp)) = 0
		leftNeighborError = Xh3NeighborRotations(tls, cell, left, bp, bp+8)
		if leftNeighborError != 0 {
			return leftNeighborError
		}
		// Set to owner if lowest index
		if *(*TH3Index)(unsafe.Pointer(bp + 8)) < owner {
			owner = *(*TH3Index)(unsafe.Pointer(bp + 8))
		}
		// As above, skip the right neighbor if the left is known lowest
		if res == 0 || libc.Int32FromUint64(*(*TH3Index)(unsafe.Pointer(bp + 8))>>((libc.Int32FromInt32(DMAX_H3_RES)-res)*libc.Int32FromInt32(DH3_PER_DIGIT_OFFSET))&libc.Uint64FromInt32(libc.Int32FromInt32(7))) != int32(_CENTER_DIGIT) {
			// Get the right neighbor of the vertex, with its rotations
			// Note that vertex - 1 is the right side, as vertex numbers are CCW
			right = XdirectionForVertexNum(tls, cell, (vertexNum-int32(1)+cellNumVerts)%cellNumVerts)
			// This case should be unreachable; invalid verts fail earlier
			if right == int32(_INVALID_DIGIT) {
				if v3 = libc.Bool(0 != 0); !v3 {
					libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+1341, int32(246), uintptr(unsafe.Pointer(&___func__24)))
				}
				_ = v3 || libc.Bool(libc.Int32FromInt32(0) != 0)
				v2 = libc.Int32FromInt32(1)
			} else {
				v2 = 0
			}
			if v2 != 0 {
				return uint32(_E_FAILED)
			}
			*(*int32)(unsafe.Pointer(bp + 16)) = 0
			rightNeighborError = Xh3NeighborRotations(tls, cell, right, bp+16, bp+24)
			if rightNeighborError != 0 {
				return rightNeighborError
			}
			// Set to owner if lowest index
			if *(*TH3Index)(unsafe.Pointer(bp + 24)) < owner {
				owner = *(*TH3Index)(unsafe.Pointer(bp + 24))
				if XisPentagon(tls, owner) != 0 {
					v4 = XdirectionForNeighbor(tls, owner, cell)
				} else {
					v4 = _DIRECTIONS1[(_revNeighborDirectionsHex[right]+*(*int32)(unsafe.Pointer(bp + 16)))%int32(DNUM_HEX_VERTS)]
				}
				dir = v4
				ownerVertexNum = XvertexNumForDirection(tls, owner, dir)
			}
		}
		// Determine the vertex number for the left neighbor
		if owner == *(*TH3Index)(unsafe.Pointer(bp + 8)) {
			ownerIsPentagon = XisPentagon(tls, owner)
			if ownerIsPentagon != 0 {
				v5 = XdirectionForNeighbor(tls, owner, cell)
			} else {
				v5 = _DIRECTIONS1[(_revNeighborDirectionsHex[left]+*(*int32)(unsafe.Pointer(bp)))%int32(DNUM_HEX_VERTS)]
			}
			dir1 = v5
			// For the left neighbor, we need the second vertex of the
			// edge, which may involve looping around the vertex nums
			ownerVertexNum = XvertexNumForDirection(tls, owner, dir1) + int32(1)
			if ownerVertexNum == int32(DNUM_HEX_VERTS) || ownerIsPentagon != 0 && ownerVertexNum == int32(DNUM_PENT_VERTS) {
				ownerVertexNum = 0
			}
		}
	}
	// Create the vertex index
	vertex = owner
	vertex = vertex & ^(libc.Uint64FromInt32(libc.Int32FromInt32(15))<<libc.Int32FromInt32(DH3_MODE_OFFSET)) | libc.Uint64FromInt32(libc.Int32FromInt32(DH3_VERTEX_MODE))<<libc.Int32FromInt32(DH3_MODE_OFFSET)
	vertex = vertex & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<libc.Int32FromInt32(DH3_RESERVED_OFFSET)) | libc.Uint64FromInt32(ownerVertexNum)<<libc.Int32FromInt32(DH3_RESERVED_OFFSET)
	*(*TH3Index)(unsafe.Pointer(out)) = vertex
	return uint32(_E_SUCCESS)
}

var ___func__24 = [13]uint8{'c', 'e', 'l', 'l', 'T', 'o', 'V', 'e', 'r', 't', 'e', 'x'}

// C documentation
//
//	/**
//	 * Get all vertexes for the given cell
//	 * @param cell      Cell to get the vertexes for
//	 * @param vertexes  Array to hold vertex output. Must have length >= 6.
//	 */
func XcellToVertexes(tls *libc.TLS, cell TH3Index, vertexes uintptr) (r TH3Error) {
	var cellError TH3Error
	var i int32
	var isPent uint8
	_, _, _ = cellError, i, isPent
	// Get all vertexes. If the cell is a pentagon, will fill the final slot
	// with H3_NULL.
	isPent = libc.Uint8FromInt32(XisPentagon(tls, cell))
	i = 0
	for {
		if !(i < int32(DNUM_HEX_VERTS)) {
			break
		}
		if i == int32(5) && isPent != 0 {
			*(*TH3Index)(unsafe.Pointer(vertexes + uintptr(i)*8)) = uint64(DH3_NULL)
		} else {
			cellError = XcellToVertex(tls, cell, i, vertexes+uintptr(i)*8)
			if cellError != 0 {
				return cellError
			}
		}
		goto _1
	_1:
		;
		i++
	}
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * Get the geocoordinates of an H3 vertex
//	 * @param vertex H3 index describing a vertex
//	 * @param coord  Output geo coordinate
//	 */
func XvertexToLatLng(tls *libc.TLS, vertex TH3Index, coord uintptr) (r TH3Error) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var fijkError TH3Error
	var owner TH3Index
	var res, vertexNum int32
	var _ /* fijk at bp+168 */ TFaceIJK
	var _ /* gb at bp+0 */ TCellBoundary
	_, _, _, _ = fijkError, owner, res, vertexNum
	// Get the vertex number and owner from the vertex
	vertexNum = libc.Int32FromUint64(vertex & (libc.Uint64FromInt32(libc.Int32FromInt32(7)) << libc.Int32FromInt32(DH3_RESERVED_OFFSET)) >> libc.Int32FromInt32(DH3_RESERVED_OFFSET))
	owner = vertex
	owner = owner & ^(libc.Uint64FromInt32(libc.Int32FromInt32(15))<<libc.Int32FromInt32(DH3_MODE_OFFSET)) | libc.Uint64FromInt32(libc.Int32FromInt32(DH3_CELL_MODE))<<libc.Int32FromInt32(DH3_MODE_OFFSET)
	owner = owner & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<libc.Int32FromInt32(DH3_RESERVED_OFFSET)) | libc.Uint64FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(DH3_RESERVED_OFFSET)
	fijkError = X_h3ToFaceIjk(tls, owner, bp+168)
	if fijkError != 0 {
		return fijkError
	}
	res = libc.Int32FromUint64(owner & (libc.Uint64FromUint64(15) << libc.Int32FromInt32(DH3_RES_OFFSET)) >> libc.Int32FromInt32(DH3_RES_OFFSET))
	if XisPentagon(tls, owner) != 0 {
		X_faceIjkPentToCellBoundary(tls, bp+168, res, vertexNum, int32(1), bp)
	} else {
		X_faceIjkToCellBoundary(tls, bp+168, res, vertexNum, int32(1), bp)
	}
	// Copy from boundary to output coord
	*(*TLatLng)(unsafe.Pointer(coord)) = *(*TLatLng)(unsafe.Pointer(bp + 8))
	return uint32(_E_SUCCESS)
}

// C documentation
//
//	/**
//	 * Whether the input is a valid H3 vertex
//	 * @param  vertex H3 index possibly describing a vertex
//	 * @return        Whether the input is valid
//	 */
func XisValidVertex(tls *libc.TLS, vertex TH3Index) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var owner TH3Index
	var vertexNum, v1 int32
	var _ /* canonical at bp+0 */ TH3Index
	_, _, _ = owner, vertexNum, v1
	if libc.Int32FromUint64(vertex&(libc.Uint64FromInt32(libc.Int32FromInt32(15))<<libc.Int32FromInt32(DH3_MODE_OFFSET))>>libc.Int32FromInt32(DH3_MODE_OFFSET)) != int32(DH3_VERTEX_MODE) {
		return 0
	}
	vertexNum = libc.Int32FromUint64(vertex & (libc.Uint64FromInt32(libc.Int32FromInt32(7)) << libc.Int32FromInt32(DH3_RESERVED_OFFSET)) >> libc.Int32FromInt32(DH3_RESERVED_OFFSET))
	owner = vertex
	owner = owner & ^(libc.Uint64FromInt32(libc.Int32FromInt32(15))<<libc.Int32FromInt32(DH3_MODE_OFFSET)) | libc.Uint64FromInt32(libc.Int32FromInt32(DH3_CELL_MODE))<<libc.Int32FromInt32(DH3_MODE_OFFSET)
	owner = owner & ^(libc.Uint64FromInt32(libc.Int32FromInt32(7))<<libc.Int32FromInt32(DH3_RESERVED_OFFSET)) | libc.Uint64FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(DH3_RESERVED_OFFSET)
	if !(XisValidCell(tls, owner) != 0) {
		return 0
	}
	if XcellToVertex(tls, owner, vertexNum, bp) != 0 {
		return 0
	}
	if vertex == *(*TH3Index)(unsafe.Pointer(bp)) {
		v1 = int32(1)
	} else {
		v1 = 0
	}
	return v1
}

/*
 * Copyright 2020 Uber Technologies, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/** @file alloc.h
 * @brief   Memory management functions
 *
 * This file contains macros and the necessary declarations to be able
 * to point H3 at different memory management functions than the standard
 * malloc/free/etc functions.
 */

/*
 * Copyright 2016-2021 Uber Technologies, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/** @file h3api.h
 * @brief   Primary H3 core library entry points.
 *
 * This file defines the public API of the H3 library. Incompatible changes to
 * these functions require the library's major version be increased.
 */

/*
 * Copyright 2016-2021 Uber Technologies, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/** @file latLng.h
 * @brief   Geodetic (lat/lng) functions.
 */

// C documentation
//
//	/**
//	 * Initialize a new VertexGraph
//	 * @param graph       Graph to initialize
//	 * @param  numBuckets Number of buckets to include in the graph
//	 * @param  res        Resolution of the hexagons whose vertices we're storing
//	 */
func XinitVertexGraph(tls *libc.TLS, graph uintptr, numBuckets int32, res int32) {
	var v1 bool
	_ = v1
	if numBuckets > 0 {
		(*TVertexGraph)(unsafe.Pointer(graph)).Fbuckets = libc.Xcalloc(tls, libc.Uint64FromInt32(numBuckets), uint64(8))
		if v1 = (*TVertexGraph)(unsafe.Pointer(graph)).Fbuckets != libc.UintptrFromInt32(0); !v1 {
			libc.X__assert_fail(tls, __ccgo_ts+1367, __ccgo_ts+1390, int32(40), uintptr(unsafe.Pointer(&___func__25)))
		}
		_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
	} else {
		(*TVertexGraph)(unsafe.Pointer(graph)).Fbuckets = libc.UintptrFromInt32(0)
	}
	(*TVertexGraph)(unsafe.Pointer(graph)).FnumBuckets = numBuckets
	(*TVertexGraph)(unsafe.Pointer(graph)).Fsize = 0
	(*TVertexGraph)(unsafe.Pointer(graph)).Fres = res
}

var ___func__25 = [16]uint8{'i', 'n', 'i', 't', 'V', 'e', 'r', 't', 'e', 'x', 'G', 'r', 'a', 'p', 'h'}

// C documentation
//
//	/**
//	 * Destroy a VertexGraph's sub-objects, freeing their memory. The caller is
//	 * responsible for freeing memory allocated to the VertexGraph struct itself.
//	 * @param graph Graph to destroy
//	 */
func XdestroyVertexGraph(tls *libc.TLS, graph uintptr) {
	var node, v1 uintptr
	_, _ = node, v1
	for {
		v1 = XfirstVertexNode(tls, graph)
		node = v1
		if !(v1 != libc.UintptrFromInt32(0)) {
			break
		}
		XremoveVertexNode(tls, graph, node)
	}
	libc.Xfree(tls, (*TVertexGraph)(unsafe.Pointer(graph)).Fbuckets)
}

// C documentation
//
//	/**
//	 * Get an integer hash for a lat/lng point, at a precision determined
//	 * by the current hexagon resolution.
//	 * TODO: Light testing suggests this might not be sufficient at resolutions
//	 * finer than 10. Design a better hash function if performance and collisions
//	 * seem to be an issue here.
//	 * @param  vertex     Lat/lng vertex to hash
//	 * @param  res        Resolution of the hexagon the vertex belongs to
//	 * @param  numBuckets Number of buckets in the graph
//	 * @return            Integer hash
//	 */
func X_hashVertex(tls *libc.TLS, vertex uintptr, res int32, numBuckets int32) (r Tuint32_t) {
	// Simple hash: Take the sum of the lat and lng with a precision level
	// determined by the resolution, converted to int, modulo bucket count.
	return uint32(libc.Xfmod(tls, libc.Xfabs(tls, ((*TLatLng)(unsafe.Pointer(vertex)).Flat+(*TLatLng)(unsafe.Pointer(vertex)).Flng)*libc.Xpow(tls, libc.Float64FromInt32(10), float64(int32(15)-res))), float64(numBuckets)))
}

func X_initVertexNode(tls *libc.TLS, node uintptr, fromVtx uintptr, toVtx uintptr) {
	(*TVertexNode)(unsafe.Pointer(node)).Ffrom = *(*TLatLng)(unsafe.Pointer(fromVtx))
	(*TVertexNode)(unsafe.Pointer(node)).Fto = *(*TLatLng)(unsafe.Pointer(toVtx))
	(*TVertexNode)(unsafe.Pointer(node)).Fnext = libc.UintptrFromInt32(0)
}

// C documentation
//
//	/**
//	 * Add a edge to the graph
//	 * @param graph   Graph to add node to
//	 * @param fromVtx Start vertex
//	 * @param toVtx   End vertex
//	 * @return        Pointer to the new node
//	 */
func XaddVertexNode(tls *libc.TLS, graph uintptr, fromVtx uintptr, toVtx uintptr) (r uintptr) {
	var currentNode, node uintptr
	var index Tuint32_t
	var v1 bool
	_, _, _, _ = currentNode, index, node, v1
	// Make the new node
	node = libc.Xmalloc(tls, uint64(40))
	if v1 = node != libc.UintptrFromInt32(0); !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+1421, __ccgo_ts+1390, int32(98), uintptr(unsafe.Pointer(&___func__26)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
	X_initVertexNode(tls, node, fromVtx, toVtx)
	// Determine location
	index = X_hashVertex(tls, fromVtx, (*TVertexGraph)(unsafe.Pointer(graph)).Fres, (*TVertexGraph)(unsafe.Pointer(graph)).FnumBuckets)
	// Check whether there's an existing node in that spot
	currentNode = *(*uintptr)(unsafe.Pointer((*TVertexGraph)(unsafe.Pointer(graph)).Fbuckets + uintptr(index)*8))
	if currentNode == libc.UintptrFromInt32(0) {
		// Set bucket to the new node
		*(*uintptr)(unsafe.Pointer((*TVertexGraph)(unsafe.Pointer(graph)).Fbuckets + uintptr(index)*8)) = node
	} else {
		// Find the end of the list
		for cond := true; cond; cond = (*TVertexNode)(unsafe.Pointer(currentNode)).Fnext != libc.UintptrFromInt32(0) {
			// Check the the edge we're adding doesn't already exist
			if XgeoAlmostEqual(tls, currentNode, fromVtx) != 0 && XgeoAlmostEqual(tls, currentNode+16, toVtx) != 0 {
				// already exists, bail
				libc.Xfree(tls, node)
				return currentNode
			}
			if (*TVertexNode)(unsafe.Pointer(currentNode)).Fnext != libc.UintptrFromInt32(0) {
				currentNode = (*TVertexNode)(unsafe.Pointer(currentNode)).Fnext
			}
		}
		// Add the new node to the end of the list
		(*TVertexNode)(unsafe.Pointer(currentNode)).Fnext = node
	}
	(*TVertexGraph)(unsafe.Pointer(graph)).Fsize++
	return node
}

var ___func__26 = [14]uint8{'a', 'd', 'd', 'V', 'e', 'r', 't', 'e', 'x', 'N', 'o', 'd', 'e'}

// C documentation
//
//	/**
//	 * Remove a node from the graph. The input node will be freed, and should
//	 * not be used after removal.
//	 * @param graph Graph to mutate
//	 * @param node  Node to remove
//	 * @return      0 on success, 1 on failure (node not found)
//	 */
func XremoveVertexNode(tls *libc.TLS, graph uintptr, node uintptr) (r int32) {
	var currentNode uintptr
	var found int32
	var index Tuint32_t
	_, _, _ = currentNode, found, index
	// Determine location
	index = X_hashVertex(tls, node, (*TVertexGraph)(unsafe.Pointer(graph)).Fres, (*TVertexGraph)(unsafe.Pointer(graph)).FnumBuckets)
	currentNode = *(*uintptr)(unsafe.Pointer((*TVertexGraph)(unsafe.Pointer(graph)).Fbuckets + uintptr(index)*8))
	found = 0
	if currentNode != libc.UintptrFromInt32(0) {
		if currentNode == node {
			*(*uintptr)(unsafe.Pointer((*TVertexGraph)(unsafe.Pointer(graph)).Fbuckets + uintptr(index)*8)) = (*TVertexNode)(unsafe.Pointer(node)).Fnext
			found = int32(1)
		}
		// Look through the list
		for !(found != 0) && (*TVertexNode)(unsafe.Pointer(currentNode)).Fnext != libc.UintptrFromInt32(0) {
			if (*TVertexNode)(unsafe.Pointer(currentNode)).Fnext == node {
				// splice the node out
				(*TVertexNode)(unsafe.Pointer(currentNode)).Fnext = (*TVertexNode)(unsafe.Pointer(node)).Fnext
				found = int32(1)
			}
			currentNode = (*TVertexNode)(unsafe.Pointer(currentNode)).Fnext
		}
	}
	if found != 0 {
		libc.Xfree(tls, node)
		(*TVertexGraph)(unsafe.Pointer(graph)).Fsize--
		return 0
	}
	// Failed to find the node
	return int32(1)
}

// C documentation
//
//	/**
//	 * Find the Vertex node for a given edge, if it exists
//	 * @param  graph   Graph to look in
//	 * @param  fromVtx Start vertex
//	 * @param  toVtx   End vertex, or NULL if we don't care
//	 * @return         Pointer to the vertex node, if found
//	 */
func XfindNodeForEdge(tls *libc.TLS, graph uintptr, fromVtx uintptr, toVtx uintptr) (r uintptr) {
	var index Tuint32_t
	var node uintptr
	_, _ = index, node
	// Determine location
	index = X_hashVertex(tls, fromVtx, (*TVertexGraph)(unsafe.Pointer(graph)).Fres, (*TVertexGraph)(unsafe.Pointer(graph)).FnumBuckets)
	// Check whether there's an existing node in that spot
	node = *(*uintptr)(unsafe.Pointer((*TVertexGraph)(unsafe.Pointer(graph)).Fbuckets + uintptr(index)*8))
	if node != libc.UintptrFromInt32(0) {
		// Look through the list and see if we find the edge
		for cond := true; cond; cond = node != libc.UintptrFromInt32(0) {
			if XgeoAlmostEqual(tls, node, fromVtx) != 0 && (toVtx == libc.UintptrFromInt32(0) || XgeoAlmostEqual(tls, node+16, toVtx) != 0) {
				return node
			}
			node = (*TVertexNode)(unsafe.Pointer(node)).Fnext
		}
	}
	// Iteration lookup fail
	return libc.UintptrFromInt32(0)
}

// C documentation
//
//	/**
//	 * Find a Vertex node starting at the given vertex
//	 * @param  graph   Graph to look in
//	 * @param  fromVtx Start vertex
//	 * @return         Pointer to the vertex node, if found
//	 */
func XfindNodeForVertex(tls *libc.TLS, graph uintptr, fromVtx uintptr) (r uintptr) {
	return XfindNodeForEdge(tls, graph, fromVtx, libc.UintptrFromInt32(0))
}

// C documentation
//
//	/**
//	 * Get the next vertex node in the graph.
//	 * @param  graph Graph to iterate
//	 * @return       Vertex node, or NULL if at the end
//	 */
func XfirstVertexNode(tls *libc.TLS, graph uintptr) (r uintptr) {
	var currentIndex int32
	var node uintptr
	_, _ = currentIndex, node
	node = libc.UintptrFromInt32(0)
	currentIndex = 0
	for node == libc.UintptrFromInt32(0) {
		if currentIndex < (*TVertexGraph)(unsafe.Pointer(graph)).FnumBuckets {
			// find the first node in the next bucket
			node = *(*uintptr)(unsafe.Pointer((*TVertexGraph)(unsafe.Pointer(graph)).Fbuckets + uintptr(currentIndex)*8))
		} else {
			// end of iteration
			return libc.UintptrFromInt32(0)
		}
		currentIndex++
	}
	return node
}

// C documentation
//
//	/**
//	 * Prohibited directions when unfolding a pentagon.
//	 *
//	 * Indexes by two directions, both relative to the pentagon base cell. The first
//	 * is the direction of the origin index and the second is the direction of the
//	 * index to unfold. Direction refers to the direction from base cell to base
//	 * cell if the indexes are on different base cells, or the leading digit if
//	 * within the pentagon base cell.
//	 *
//	 * This previously included a Class II/Class III check but these were removed
//	 * due to failure cases. It's possible this could be restricted to a narrower
//	 * set of a failure cases. Currently, the logic is any unfolding across more
//	 * than one icosahedron face is not permitted.
//	 */
var XFAILED_DIRECTIONS = [7][7]uint8{
	0: {},
	1: {},
	2: {
		4: uint8(Dtrue),
		5: uint8(Dtrue),
	},
	3: {
		4: uint8(Dtrue),
		6: uint8(Dtrue),
	},
	4: {
		2: uint8(Dtrue),
		3: uint8(Dtrue),
	},
	5: {
		2: uint8(Dtrue),
		6: uint8(Dtrue),
	},
	6: {
		3: uint8(Dtrue),
		5: uint8(Dtrue),
	},
}

// C documentation
//
//	/**
//	 * Origin leading digit -> index leading digit -> rotations 60 cw
//	 * Either being 1 (K axis) is invalid.
//	 * No good default at 0.
//	 */
var XPENTAGON_ROTATIONS = [7][7]int32{
	0: {
		1: -int32(1),
	},
	1: {
		0: -int32(1),
		1: -int32(1),
		2: -int32(1),
		3: -int32(1),
		4: -int32(1),
		5: -int32(1),
		6: -int32(1),
	},
	2: {
		1: -int32(1),
		5: int32(1),
	},
	3: {
		1: -int32(1),
		4: int32(1),
		5: int32(1),
	},
	4: {
		1: -int32(1),
		3: int32(5),
	},
	5: {
		1: -int32(1),
		2: int32(5),
		3: int32(5),
	},
	6: {
		1: -int32(1),
	},
}

// C documentation
//
//	/**
//	 * Reverse base cell direction -> leading index digit -> rotations 60 ccw.
//	 * For reversing the rotation introduced in PENTAGON_ROTATIONS when
//	 * the origin is on a pentagon (regardless of the base cell of the index.)
//	 */
var XPENTAGON_ROTATIONS_REVERSE = [7][7]int32{
	0: {},
	1: {
		0: -int32(1),
		1: -int32(1),
		2: -int32(1),
		3: -int32(1),
		4: -int32(1),
		5: -int32(1),
		6: -int32(1),
	},
	2: {
		1: int32(1),
	},
	3: {
		1: int32(1),
		5: int32(1),
	},
	4: {
		1: int32(5),
	},
	5: {
		1: int32(5),
		3: int32(5),
	},
	6: {},
}

// C documentation
//
//	/**
//	 * Reverse base cell direction -> leading index digit -> rotations 60 ccw.
//	 * For reversing the rotation introduced in PENTAGON_ROTATIONS when the index is
//	 * on a pentagon and the origin is not.
//	 */
var XPENTAGON_ROTATIONS_REVERSE_NONPOLAR = [7][7]int32{
	0: {},
	1: {
		0: -int32(1),
		1: -int32(1),
		2: -int32(1),
		3: -int32(1),
		4: -int32(1),
		5: -int32(1),
		6: -int32(1),
	},
	2: {
		1: int32(1),
	},
	3: {
		1: int32(1),
		5: int32(1),
	},
	4: {
		1: int32(5),
	},
	5: {
		1: int32(1),
		3: int32(5),
		4: int32(1),
		5: int32(1),
	},
	6: {},
}

// C documentation
//
//	/**
//	 * Reverse base cell direction -> leading index digit -> rotations 60 ccw.
//	 * For reversing the rotation introduced in PENTAGON_ROTATIONS when the index is
//	 * on a polar pentagon and the origin is not.
//	 */
var XPENTAGON_ROTATIONS_REVERSE_POLAR = [7][7]int32{
	0: {},
	1: {
		0: -int32(1),
		1: -int32(1),
		2: -int32(1),
		3: -int32(1),
		4: -int32(1),
		5: -int32(1),
		6: -int32(1),
	},
	2: {
		1: int32(1),
		2: int32(1),
		3: int32(1),
		4: int32(1),
		5: int32(1),
		6: int32(1),
	},
	3: {
		1: int32(1),
		5: int32(1),
	},
	4: {
		1: int32(1),
		4: int32(1),
		5: int32(1),
		6: int32(1),
	},
	5: {
		1: int32(1),
		3: int32(5),
		4: int32(1),
		5: int32(1),
	},
	6: {
		1: int32(1),
		2: int32(1),
		4: int32(1),
		5: int32(1),
		6: int32(1),
	},
}

// C documentation
//
//	/** @brief Resolution 0 base cell data table.
//	 *
//	 * For each base cell, gives the "home" face and ijk+ coordinates on that face,
//	 * whether or not the base cell is a pentagon. Additionally, if the base cell
//	 * is a pentagon, the two cw offset rotation adjacent faces are given (-1
//	 * indicates that no cw offset rotation faces exist for this base cell).
//	 */
var XbaseCellData = [122]TBaseCellData{
	0: {
		FhomeFijk: TFaceIJK{
			Fface: int32(1),
			Fcoord: TCoordIJK{
				Fi: int32(1),
			},
		},
	},
	1: {
		FhomeFijk: TFaceIJK{
			Fface: int32(2),
			Fcoord: TCoordIJK{
				Fi: int32(1),
				Fj: int32(1),
			},
		},
	},
	2: {
		FhomeFijk: TFaceIJK{
			Fface: int32(1),
		},
	},
	3: {
		FhomeFijk: TFaceIJK{
			Fface: int32(2),
			Fcoord: TCoordIJK{
				Fi: int32(1),
			},
		},
	},
	4: {
		FhomeFijk: TFaceIJK{
			Fcoord: TCoordIJK{
				Fi: int32(2),
			},
		},
		FisPentagon: int32(1),
		FcwOffsetPent: [2]int32{
			0: -int32(1),
			1: -int32(1),
		},
	},
	5: {
		FhomeFijk: TFaceIJK{
			Fface: int32(1),
			Fcoord: TCoordIJK{
				Fi: int32(1),
				Fj: int32(1),
			},
		},
	},
	6: {
		FhomeFijk: TFaceIJK{
			Fface: int32(1),
			Fcoord: TCoordIJK{
				Fk: int32(1),
			},
		},
	},
	7: {
		FhomeFijk: TFaceIJK{
			Fface: int32(2),
		},
	},
	8: {
		FhomeFijk: TFaceIJK{
			Fcoord: TCoordIJK{
				Fi: int32(1),
			},
		},
	},
	9: {
		FhomeFijk: TFaceIJK{
			Fface: int32(2),
			Fcoord: TCoordIJK{
				Fj: int32(1),
			},
		},
	},
	10: {
		FhomeFijk: TFaceIJK{
			Fface: int32(1),
			Fcoord: TCoordIJK{
				Fj: int32(1),
			},
		},
	},
	11: {
		FhomeFijk: TFaceIJK{
			Fface: int32(1),
			Fcoord: TCoordIJK{
				Fj: int32(1),
				Fk: int32(1),
			},
		},
	},
	12: {
		FhomeFijk: TFaceIJK{
			Fface: int32(3),
			Fcoord: TCoordIJK{
				Fi: int32(1),
			},
		},
	},
	13: {
		FhomeFijk: TFaceIJK{
			Fface: int32(3),
			Fcoord: TCoordIJK{
				Fi: int32(1),
				Fj: int32(1),
			},
		},
	},
	14: {
		FhomeFijk: TFaceIJK{
			Fface: int32(11),
			Fcoord: TCoordIJK{
				Fi: int32(2),
			},
		},
		FisPentagon: int32(1),
		FcwOffsetPent: [2]int32{
			0: int32(2),
			1: int32(6),
		},
	},
	15: {
		FhomeFijk: TFaceIJK{
			Fface: int32(4),
			Fcoord: TCoordIJK{
				Fi: int32(1),
			},
		},
	},
	16: {},
	17: {
		FhomeFijk: TFaceIJK{
			Fface: int32(6),
			Fcoord: TCoordIJK{
				Fj: int32(1),
			},
		},
	},
	18: {
		FhomeFijk: TFaceIJK{
			Fcoord: TCoordIJK{
				Fk: int32(1),
			},
		},
	},
	19: {
		FhomeFijk: TFaceIJK{
			Fface: int32(2),
			Fcoord: TCoordIJK{
				Fj: int32(1),
				Fk: int32(1),
			},
		},
	},
	20: {
		FhomeFijk: TFaceIJK{
			Fface: int32(7),
			Fcoord: TCoordIJK{
				Fk: int32(1),
			},
		},
	},
	21: {
		FhomeFijk: TFaceIJK{
			Fface: int32(2),
			Fcoord: TCoordIJK{
				Fk: int32(1),
			},
		},
	},
	22: {
		FhomeFijk: TFaceIJK{
			Fcoord: TCoordIJK{
				Fi: int32(1),
				Fj: int32(1),
			},
		},
	},
	23: {
		FhomeFijk: TFaceIJK{
			Fface: int32(6),
			Fcoord: TCoordIJK{
				Fk: int32(1),
			},
		},
	},
	24: {
		FhomeFijk: TFaceIJK{
			Fface: int32(10),
			Fcoord: TCoordIJK{
				Fi: int32(2),
			},
		},
		FisPentagon: int32(1),
		FcwOffsetPent: [2]int32{
			0: int32(1),
			1: int32(5),
		},
	},
	25: {
		FhomeFijk: TFaceIJK{
			Fface: int32(6),
		},
	},
	26: {
		FhomeFijk: TFaceIJK{
			Fface: int32(3),
		},
	},
	27: {
		FhomeFijk: TFaceIJK{
			Fface: int32(11),
			Fcoord: TCoordIJK{
				Fi: int32(1),
			},
		},
	},
	28: {
		FhomeFijk: TFaceIJK{
			Fface: int32(4),
			Fcoord: TCoordIJK{
				Fi: int32(1),
				Fj: int32(1),
			},
		},
	},
	29: {
		FhomeFijk: TFaceIJK{
			Fface: int32(3),
			Fcoord: TCoordIJK{
				Fj: int32(1),
			},
		},
	},
	30: {
		FhomeFijk: TFaceIJK{
			Fcoord: TCoordIJK{
				Fj: int32(1),
				Fk: int32(1),
			},
		},
	},
	31: {
		FhomeFijk: TFaceIJK{
			Fface: int32(4),
		},
	},
	32: {
		FhomeFijk: TFaceIJK{
			Fface: int32(5),
			Fcoord: TCoordIJK{
				Fj: int32(1),
			},
		},
	},
	33: {
		FhomeFijk: TFaceIJK{
			Fcoord: TCoordIJK{
				Fj: int32(1),
			},
		},
	},
	34: {
		FhomeFijk: TFaceIJK{
			Fface: int32(7),
			Fcoord: TCoordIJK{
				Fj: int32(1),
			},
		},
	},
	35: {
		FhomeFijk: TFaceIJK{
			Fface: int32(11),
			Fcoord: TCoordIJK{
				Fi: int32(1),
				Fj: int32(1),
			},
		},
	},
	36: {
		FhomeFijk: TFaceIJK{
			Fface: int32(7),
		},
	},
	37: {
		FhomeFijk: TFaceIJK{
			Fface: int32(10),
			Fcoord: TCoordIJK{
				Fi: int32(1),
			},
		},
	},
	38: {
		FhomeFijk: TFaceIJK{
			Fface: int32(12),
			Fcoord: TCoordIJK{
				Fi: int32(2),
			},
		},
		FisPentagon: int32(1),
		FcwOffsetPent: [2]int32{
			0: int32(3),
			1: int32(7),
		},
	},
	39: {
		FhomeFijk: TFaceIJK{
			Fface: int32(6),
			Fcoord: TCoordIJK{
				Fi: int32(1),
				Fk: int32(1),
			},
		},
	},
	40: {
		FhomeFijk: TFaceIJK{
			Fface: int32(7),
			Fcoord: TCoordIJK{
				Fi: int32(1),
				Fk: int32(1),
			},
		},
	},
	41: {
		FhomeFijk: TFaceIJK{
			Fface: int32(4),
			Fcoord: TCoordIJK{
				Fk: int32(1),
			},
		},
	},
	42: {
		FhomeFijk: TFaceIJK{
			Fface: int32(3),
			Fcoord: TCoordIJK{
				Fk: int32(1),
			},
		},
	},
	43: {
		FhomeFijk: TFaceIJK{
			Fface: int32(3),
			Fcoord: TCoordIJK{
				Fj: int32(1),
				Fk: int32(1),
			},
		},
	},
	44: {
		FhomeFijk: TFaceIJK{
			Fface: int32(4),
			Fcoord: TCoordIJK{
				Fj: int32(1),
			},
		},
	},
	45: {
		FhomeFijk: TFaceIJK{
			Fface: int32(6),
			Fcoord: TCoordIJK{
				Fi: int32(1),
			},
		},
	},
	46: {
		FhomeFijk: TFaceIJK{
			Fface: int32(11),
		},
	},
	47: {
		FhomeFijk: TFaceIJK{
			Fface: int32(8),
			Fcoord: TCoordIJK{
				Fk: int32(1),
			},
		},
	},
	48: {
		FhomeFijk: TFaceIJK{
			Fface: int32(5),
			Fcoord: TCoordIJK{
				Fk: int32(1),
			},
		},
	},
	49: {
		FhomeFijk: TFaceIJK{
			Fface: int32(14),
			Fcoord: TCoordIJK{
				Fi: int32(2),
			},
		},
		FisPentagon: int32(1),
		FcwOffsetPent: [2]int32{
			1: int32(9),
		},
	},
	50: {
		FhomeFijk: TFaceIJK{
			Fface: int32(5),
		},
	},
	51: {
		FhomeFijk: TFaceIJK{
			Fface: int32(12),
			Fcoord: TCoordIJK{
				Fi: int32(1),
			},
		},
	},
	52: {
		FhomeFijk: TFaceIJK{
			Fface: int32(10),
			Fcoord: TCoordIJK{
				Fi: int32(1),
				Fj: int32(1),
			},
		},
	},
	53: {
		FhomeFijk: TFaceIJK{
			Fface: int32(4),
			Fcoord: TCoordIJK{
				Fj: int32(1),
				Fk: int32(1),
			},
		},
	},
	54: {
		FhomeFijk: TFaceIJK{
			Fface: int32(12),
			Fcoord: TCoordIJK{
				Fi: int32(1),
				Fj: int32(1),
			},
		},
	},
	55: {
		FhomeFijk: TFaceIJK{
			Fface: int32(7),
			Fcoord: TCoordIJK{
				Fi: int32(1),
			},
		},
	},
	56: {
		FhomeFijk: TFaceIJK{
			Fface: int32(11),
			Fcoord: TCoordIJK{
				Fj: int32(1),
			},
		},
	},
	57: {
		FhomeFijk: TFaceIJK{
			Fface: int32(10),
		},
	},
	58: {
		FhomeFijk: TFaceIJK{
			Fface: int32(13),
			Fcoord: TCoordIJK{
				Fi: int32(2),
			},
		},
		FisPentagon: int32(1),
		FcwOffsetPent: [2]int32{
			0: int32(4),
			1: int32(8),
		},
	},
	59: {
		FhomeFijk: TFaceIJK{
			Fface: int32(10),
			Fcoord: TCoordIJK{
				Fk: int32(1),
			},
		},
	},
	60: {
		FhomeFijk: TFaceIJK{
			Fface: int32(11),
			Fcoord: TCoordIJK{
				Fk: int32(1),
			},
		},
	},
	61: {
		FhomeFijk: TFaceIJK{
			Fface: int32(9),
			Fcoord: TCoordIJK{
				Fj: int32(1),
			},
		},
	},
	62: {
		FhomeFijk: TFaceIJK{
			Fface: int32(8),
			Fcoord: TCoordIJK{
				Fj: int32(1),
			},
		},
	},
	63: {
		FhomeFijk: TFaceIJK{
			Fface: int32(6),
			Fcoord: TCoordIJK{
				Fi: int32(2),
			},
		},
		FisPentagon: int32(1),
		FcwOffsetPent: [2]int32{
			0: int32(11),
			1: int32(15),
		},
	},
	64: {
		FhomeFijk: TFaceIJK{
			Fface: int32(8),
		},
	},
	65: {
		FhomeFijk: TFaceIJK{
			Fface: int32(9),
			Fcoord: TCoordIJK{
				Fk: int32(1),
			},
		},
	},
	66: {
		FhomeFijk: TFaceIJK{
			Fface: int32(14),
			Fcoord: TCoordIJK{
				Fi: int32(1),
			},
		},
	},
	67: {
		FhomeFijk: TFaceIJK{
			Fface: int32(5),
			Fcoord: TCoordIJK{
				Fi: int32(1),
				Fk: int32(1),
			},
		},
	},
	68: {
		FhomeFijk: TFaceIJK{
			Fface: int32(16),
			Fcoord: TCoordIJK{
				Fj: int32(1),
				Fk: int32(1),
			},
		},
	},
	69: {
		FhomeFijk: TFaceIJK{
			Fface: int32(8),
			Fcoord: TCoordIJK{
				Fi: int32(1),
				Fk: int32(1),
			},
		},
	},
	70: {
		FhomeFijk: TFaceIJK{
			Fface: int32(5),
			Fcoord: TCoordIJK{
				Fi: int32(1),
			},
		},
	},
	71: {
		FhomeFijk: TFaceIJK{
			Fface: int32(12),
		},
	},
	72: {
		FhomeFijk: TFaceIJK{
			Fface: int32(7),
			Fcoord: TCoordIJK{
				Fi: int32(2),
			},
		},
		FisPentagon: int32(1),
		FcwOffsetPent: [2]int32{
			0: int32(12),
			1: int32(16),
		},
	},
	73: {
		FhomeFijk: TFaceIJK{
			Fface: int32(12),
			Fcoord: TCoordIJK{
				Fj: int32(1),
			},
		},
	},
	74: {
		FhomeFijk: TFaceIJK{
			Fface: int32(10),
			Fcoord: TCoordIJK{
				Fj: int32(1),
			},
		},
	},
	75: {
		FhomeFijk: TFaceIJK{
			Fface: int32(9),
		},
	},
	76: {
		FhomeFijk: TFaceIJK{
			Fface: int32(13),
			Fcoord: TCoordIJK{
				Fi: int32(1),
			},
		},
	},
	77: {
		FhomeFijk: TFaceIJK{
			Fface: int32(16),
			Fcoord: TCoordIJK{
				Fk: int32(1),
			},
		},
	},
	78: {
		FhomeFijk: TFaceIJK{
			Fface: int32(15),
			Fcoord: TCoordIJK{
				Fj: int32(1),
				Fk: int32(1),
			},
		},
	},
	79: {
		FhomeFijk: TFaceIJK{
			Fface: int32(15),
			Fcoord: TCoordIJK{
				Fj: int32(1),
			},
		},
	},
	80: {
		FhomeFijk: TFaceIJK{
			Fface: int32(16),
			Fcoord: TCoordIJK{
				Fj: int32(1),
			},
		},
	},
	81: {
		FhomeFijk: TFaceIJK{
			Fface: int32(14),
			Fcoord: TCoordIJK{
				Fi: int32(1),
				Fj: int32(1),
			},
		},
	},
	82: {
		FhomeFijk: TFaceIJK{
			Fface: int32(13),
			Fcoord: TCoordIJK{
				Fi: int32(1),
				Fj: int32(1),
			},
		},
	},
	83: {
		FhomeFijk: TFaceIJK{
			Fface: int32(5),
			Fcoord: TCoordIJK{
				Fi: int32(2),
			},
		},
		FisPentagon: int32(1),
		FcwOffsetPent: [2]int32{
			0: int32(10),
			1: int32(19),
		},
	},
	84: {
		FhomeFijk: TFaceIJK{
			Fface: int32(8),
			Fcoord: TCoordIJK{
				Fi: int32(1),
			},
		},
	},
	85: {
		FhomeFijk: TFaceIJK{
			Fface: int32(14),
		},
	},
	86: {
		FhomeFijk: TFaceIJK{
			Fface: int32(9),
			Fcoord: TCoordIJK{
				Fi: int32(1),
				Fk: int32(1),
			},
		},
	},
	87: {
		FhomeFijk: TFaceIJK{
			Fface: int32(14),
			Fcoord: TCoordIJK{
				Fk: int32(1),
			},
		},
	},
	88: {
		FhomeFijk: TFaceIJK{
			Fface: int32(17),
			Fcoord: TCoordIJK{
				Fk: int32(1),
			},
		},
	},
	89: {
		FhomeFijk: TFaceIJK{
			Fface: int32(12),
			Fcoord: TCoordIJK{
				Fk: int32(1),
			},
		},
	},
	90: {
		FhomeFijk: TFaceIJK{
			Fface: int32(16),
		},
	},
	91: {
		FhomeFijk: TFaceIJK{
			Fface: int32(17),
			Fcoord: TCoordIJK{
				Fj: int32(1),
				Fk: int32(1),
			},
		},
	},
	92: {
		FhomeFijk: TFaceIJK{
			Fface: int32(15),
			Fcoord: TCoordIJK{
				Fk: int32(1),
			},
		},
	},
	93: {
		FhomeFijk: TFaceIJK{
			Fface: int32(16),
			Fcoord: TCoordIJK{
				Fi: int32(1),
				Fk: int32(1),
			},
		},
	},
	94: {
		FhomeFijk: TFaceIJK{
			Fface: int32(9),
			Fcoord: TCoordIJK{
				Fi: int32(1),
			},
		},
	},
	95: {
		FhomeFijk: TFaceIJK{
			Fface: int32(15),
		},
	},
	96: {
		FhomeFijk: TFaceIJK{
			Fface: int32(13),
		},
	},
	97: {
		FhomeFijk: TFaceIJK{
			Fface: int32(8),
			Fcoord: TCoordIJK{
				Fi: int32(2),
			},
		},
		FisPentagon: int32(1),
		FcwOffsetPent: [2]int32{
			0: int32(13),
			1: int32(17),
		},
	},
	98: {
		FhomeFijk: TFaceIJK{
			Fface: int32(13),
			Fcoord: TCoordIJK{
				Fj: int32(1),
			},
		},
	},
	99: {
		FhomeFijk: TFaceIJK{
			Fface: int32(17),
			Fcoord: TCoordIJK{
				Fi: int32(1),
				Fk: int32(1),
			},
		},
	},
	100: {
		FhomeFijk: TFaceIJK{
			Fface: int32(19),
			Fcoord: TCoordIJK{
				Fj: int32(1),
			},
		},
	},
	101: {
		FhomeFijk: TFaceIJK{
			Fface: int32(14),
			Fcoord: TCoordIJK{
				Fj: int32(1),
			},
		},
	},
	102: {
		FhomeFijk: TFaceIJK{
			Fface: int32(19),
			Fcoord: TCoordIJK{
				Fj: int32(1),
				Fk: int32(1),
			},
		},
	},
	103: {
		FhomeFijk: TFaceIJK{
			Fface: int32(17),
			Fcoord: TCoordIJK{
				Fj: int32(1),
			},
		},
	},
	104: {
		FhomeFijk: TFaceIJK{
			Fface: int32(13),
			Fcoord: TCoordIJK{
				Fk: int32(1),
			},
		},
	},
	105: {
		FhomeFijk: TFaceIJK{
			Fface: int32(17),
		},
	},
	106: {
		FhomeFijk: TFaceIJK{
			Fface: int32(16),
			Fcoord: TCoordIJK{
				Fi: int32(1),
			},
		},
	},
	107: {
		FhomeFijk: TFaceIJK{
			Fface: int32(9),
			Fcoord: TCoordIJK{
				Fi: int32(2),
			},
		},
		FisPentagon: int32(1),
		FcwOffsetPent: [2]int32{
			0: int32(14),
			1: int32(18),
		},
	},
	108: {
		FhomeFijk: TFaceIJK{
			Fface: int32(15),
			Fcoord: TCoordIJK{
				Fi: int32(1),
				Fk: int32(1),
			},
		},
	},
	109: {
		FhomeFijk: TFaceIJK{
			Fface: int32(15),
			Fcoord: TCoordIJK{
				Fi: int32(1),
			},
		},
	},
	110: {
		FhomeFijk: TFaceIJK{
			Fface: int32(18),
			Fcoord: TCoordIJK{
				Fj: int32(1),
				Fk: int32(1),
			},
		},
	},
	111: {
		FhomeFijk: TFaceIJK{
			Fface: int32(18),
			Fcoord: TCoordIJK{
				Fk: int32(1),
			},
		},
	},
	112: {
		FhomeFijk: TFaceIJK{
			Fface: int32(19),
			Fcoord: TCoordIJK{
				Fk: int32(1),
			},
		},
	},
	113: {
		FhomeFijk: TFaceIJK{
			Fface: int32(17),
			Fcoord: TCoordIJK{
				Fi: int32(1),
			},
		},
	},
	114: {
		FhomeFijk: TFaceIJK{
			Fface: int32(19),
		},
	},
	115: {
		FhomeFijk: TFaceIJK{
			Fface: int32(18),
			Fcoord: TCoordIJK{
				Fj: int32(1),
			},
		},
	},
	116: {
		FhomeFijk: TFaceIJK{
			Fface: int32(18),
			Fcoord: TCoordIJK{
				Fi: int32(1),
				Fk: int32(1),
			},
		},
	},
	117: {
		FhomeFijk: TFaceIJK{
			Fface: int32(19),
			Fcoord: TCoordIJK{
				Fi: int32(2),
			},
		},
		FisPentagon: int32(1),
		FcwOffsetPent: [2]int32{
			0: -int32(1),
			1: -int32(1),
		},
	},
	118: {
		FhomeFijk: TFaceIJK{
			Fface: int32(19),
			Fcoord: TCoordIJK{
				Fi: int32(1),
			},
		},
	},
	119: {
		FhomeFijk: TFaceIJK{
			Fface: int32(18),
		},
	},
	120: {
		FhomeFijk: TFaceIJK{
			Fface: int32(19),
			Fcoord: TCoordIJK{
				Fi: int32(1),
				Fk: int32(1),
			},
		},
	},
	121: {
		FhomeFijk: TFaceIJK{
			Fface: int32(18),
			Fcoord: TCoordIJK{
				Fi: int32(1),
			},
		},
	},
}

// C documentation
//
//	/** @brief Neighboring base cell rotations in each IJK direction.
//	 *
//	 * For each base cell, for each direction, the number of 60 degree
//	 * CCW rotations to the coordinate system of the neighbor is given.
//	 * -1 indicates there is no neighbor in that direction.
//	 */
var XbaseCellNeighbor60CCWRots = [122][7]int32{
	0: {
		1: int32(5),
		4: int32(1),
		5: int32(5),
		6: int32(1),
	},
	1: {
		2: int32(1),
		4: int32(1),
		6: int32(1),
	},
	2: {
		5: int32(5),
	},
	3: {
		1: int32(5),
		4: int32(2),
		5: int32(5),
		6: int32(1),
	},
	4: {
		1: -int32(1),
		2: int32(1),
		4: int32(3),
		5: int32(4),
		6: int32(2),
	},
	5: {
		2: int32(1),
		4: int32(1),
		6: int32(1),
	},
	6: {
		3: int32(3),
		4: int32(5),
		5: int32(5),
	},
	7: {
		5: int32(5),
	},
	8: {
		1: int32(5),
		5: int32(5),
		6: int32(1),
	},
	9: {
		2: int32(1),
		3: int32(3),
		6: int32(1),
	},
	10: {
		2: int32(1),
		3: int32(3),
		6: int32(1),
	},
	11: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
	},
	12: {
		1: int32(5),
		4: int32(3),
		5: int32(5),
		6: int32(1),
	},
	13: {
		2: int32(1),
		4: int32(1),
		6: int32(1),
	},
	14: {
		1: -int32(1),
		2: int32(3),
		4: int32(5),
		5: int32(2),
	},
	15: {
		1: int32(5),
		4: int32(4),
		5: int32(5),
		6: int32(1),
	},
	16: {
		5: int32(5),
	},
	17: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
		4: int32(3),
		6: int32(3),
	},
	18: {
		3: int32(3),
		4: int32(5),
		5: int32(5),
	},
	19: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
	},
	20: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
		5: int32(3),
	},
	21: {
		3: int32(3),
		4: int32(5),
		5: int32(5),
	},
	22: {
		2: int32(1),
		4: int32(1),
		6: int32(1),
	},
	23: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
		5: int32(3),
	},
	24: {
		1: -int32(1),
		2: int32(3),
		4: int32(5),
		5: int32(2),
	},
	25: {
		3: int32(3),
		6: int32(3),
	},
	26: {
		5: int32(5),
	},
	27: {
		1: int32(3),
		5: int32(3),
		6: int32(3),
	},
	28: {
		2: int32(1),
		4: int32(1),
		6: int32(1),
	},
	29: {
		2: int32(1),
		3: int32(3),
		6: int32(1),
	},
	30: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
	},
	31: {
		5: int32(5),
	},
	32: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
		4: int32(3),
		6: int32(3),
	},
	33: {
		2: int32(1),
		3: int32(3),
		6: int32(1),
	},
	34: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
		4: int32(3),
		6: int32(3),
	},
	35: {
		2: int32(3),
		4: int32(3),
		6: int32(3),
	},
	36: {
		3: int32(3),
		6: int32(3),
	},
	37: {
		1: int32(3),
		5: int32(3),
		6: int32(3),
	},
	38: {
		1: -int32(1),
		2: int32(3),
		4: int32(5),
		5: int32(2),
	},
	39: {
		1: int32(3),
		4: int32(3),
		5: int32(3),
	},
	40: {
		1: int32(3),
		4: int32(3),
		5: int32(3),
	},
	41: {
		3: int32(3),
		4: int32(5),
		5: int32(5),
	},
	42: {
		3: int32(3),
		4: int32(5),
		5: int32(5),
	},
	43: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
	},
	44: {
		2: int32(1),
		3: int32(3),
		6: int32(1),
	},
	45: {
		2: int32(3),
		5: int32(3),
		6: int32(3),
	},
	46: {
		3: int32(3),
		5: int32(3),
	},
	47: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
		5: int32(3),
	},
	48: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
		5: int32(3),
	},
	49: {
		1: -int32(1),
		2: int32(3),
		4: int32(5),
		5: int32(2),
	},
	50: {
		3: int32(3),
		6: int32(3),
	},
	51: {
		1: int32(3),
		5: int32(3),
		6: int32(3),
	},
	52: {
		2: int32(3),
		4: int32(3),
		6: int32(3),
	},
	53: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
	},
	54: {
		2: int32(3),
		4: int32(3),
		6: int32(3),
	},
	55: {
		2: int32(3),
		5: int32(3),
		6: int32(3),
	},
	56: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
		6: int32(3),
	},
	57: {
		3: int32(3),
		5: int32(3),
	},
	58: {
		1: -int32(1),
		2: int32(3),
		4: int32(5),
		5: int32(2),
	},
	59: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
		4: int32(3),
		5: int32(3),
	},
	60: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
		4: int32(3),
		5: int32(3),
	},
	61: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
		4: int32(3),
		6: int32(3),
	},
	62: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
		4: int32(3),
		6: int32(3),
	},
	63: {
		1: -int32(1),
		2: int32(3),
		4: int32(5),
		5: int32(2),
	},
	64: {
		3: int32(3),
		6: int32(3),
	},
	65: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
		5: int32(3),
	},
	66: {
		1: int32(3),
		5: int32(3),
		6: int32(3),
	},
	67: {
		1: int32(3),
		4: int32(3),
		5: int32(3),
	},
	68: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
	},
	69: {
		1: int32(3),
		4: int32(3),
		5: int32(3),
	},
	70: {
		2: int32(3),
		5: int32(3),
		6: int32(3),
	},
	71: {
		3: int32(3),
		5: int32(3),
	},
	72: {
		1: -int32(1),
		2: int32(3),
		4: int32(5),
		5: int32(2),
	},
	73: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
		6: int32(3),
	},
	74: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
		6: int32(3),
	},
	75: {
		3: int32(3),
		6: int32(3),
	},
	76: {
		1: int32(3),
		5: int32(3),
		6: int32(3),
	},
	77: {
		3: int32(3),
		5: int32(5),
	},
	78: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
	},
	79: {
		2: int32(1),
		3: int32(3),
		4: int32(1),
		6: int32(1),
	},
	80: {
		2: int32(1),
		3: int32(3),
		4: int32(1),
		6: int32(1),
	},
	81: {
		2: int32(3),
		4: int32(3),
		6: int32(3),
	},
	82: {
		2: int32(3),
		4: int32(3),
		6: int32(3),
	},
	83: {
		1: -int32(1),
		2: int32(3),
		4: int32(5),
		5: int32(2),
	},
	84: {
		2: int32(3),
		5: int32(3),
		6: int32(3),
	},
	85: {
		3: int32(3),
		5: int32(3),
	},
	86: {
		1: int32(3),
		4: int32(3),
		5: int32(3),
	},
	87: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
		4: int32(3),
		5: int32(3),
	},
	88: {
		3: int32(3),
		5: int32(5),
	},
	89: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
		4: int32(3),
		5: int32(3),
	},
	90: {
		6: int32(1),
	},
	91: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
	},
	92: {
		3: int32(3),
		5: int32(5),
	},
	93: {
		1: int32(5),
		4: int32(5),
		5: int32(5),
	},
	94: {
		2: int32(3),
		5: int32(3),
		6: int32(3),
	},
	95: {
		6: int32(1),
	},
	96: {
		3: int32(3),
		5: int32(3),
	},
	97: {
		1: -int32(1),
		2: int32(3),
		4: int32(5),
		5: int32(2),
	},
	98: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
		6: int32(3),
	},
	99: {
		1: int32(5),
		4: int32(5),
		5: int32(5),
	},
	100: {
		2: int32(1),
		3: int32(3),
		4: int32(1),
		6: int32(1),
	},
	101: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
		6: int32(3),
	},
	102: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
	},
	103: {
		2: int32(1),
		3: int32(3),
		4: int32(1),
		6: int32(1),
	},
	104: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
		4: int32(3),
		5: int32(3),
	},
	105: {
		6: int32(1),
	},
	106: {
		2: int32(1),
		4: int32(3),
		5: int32(5),
		6: int32(1),
	},
	107: {
		1: -int32(1),
		2: int32(3),
		4: int32(5),
		5: int32(2),
	},
	108: {
		1: int32(5),
		4: int32(5),
		5: int32(5),
	},
	109: {
		2: int32(1),
		4: int32(4),
		5: int32(5),
		6: int32(1),
	},
	110: {
		1: int32(3),
		2: int32(3),
		3: int32(3),
	},
	111: {
		3: int32(3),
		5: int32(5),
	},
	112: {
		3: int32(3),
		5: int32(5),
	},
	113: {
		2: int32(1),
		4: int32(2),
		5: int32(5),
		6: int32(1),
	},
	114: {
		6: int32(1),
	},
	115: {
		2: int32(1),
		3: int32(3),
		4: int32(1),
		6: int32(1),
	},
	116: {
		1: int32(5),
		4: int32(5),
		5: int32(5),
	},
	117: {
		1: -int32(1),
		2: int32(1),
		4: int32(3),
		5: int32(4),
		6: int32(2),
	},
	118: {
		2: int32(1),
		5: int32(5),
		6: int32(1),
	},
	119: {
		6: int32(1),
	},
	120: {
		1: int32(5),
		4: int32(5),
		5: int32(5),
	},
	121: {
		2: int32(1),
		4: int32(1),
		5: int32(5),
		6: int32(1),
	},
}

// C documentation
//
//	/** @brief Neighboring base cell ID in each IJK direction.
//	 *
//	 * For each base cell, for each direction, the neighboring base
//	 * cell ID is given. 127 indicates there is no neighbor in that direction.
//	 */
var XbaseCellNeighbors = [122][7]int32{
	0: {
		1: int32(1),
		2: int32(5),
		3: int32(2),
		4: int32(4),
		5: int32(3),
		6: int32(8),
	},
	1: {
		0: int32(1),
		1: int32(7),
		2: int32(6),
		3: int32(9),
		5: int32(3),
		6: int32(2),
	},
	2: {
		0: int32(2),
		1: int32(6),
		2: int32(10),
		3: int32(11),
		5: int32(1),
		6: int32(5),
	},
	3: {
		0: int32(3),
		1: int32(13),
		2: int32(1),
		3: int32(7),
		4: int32(4),
		5: int32(12),
	},
	4: {
		0: int32(4),
		1: int32(DINVALID_BASE_CELL),
		2: int32(15),
		3: int32(8),
		4: int32(3),
		6: int32(12),
	},
	5: {
		0: int32(5),
		1: int32(2),
		2: int32(18),
		3: int32(10),
		4: int32(8),
		6: int32(16),
	},
	6: {
		0: int32(6),
		1: int32(14),
		2: int32(11),
		3: int32(17),
		4: int32(1),
		5: int32(9),
		6: int32(2),
	},
	7: {
		0: int32(7),
		1: int32(21),
		2: int32(9),
		3: int32(19),
		4: int32(3),
		5: int32(13),
		6: int32(1),
	},
	8: {
		0: int32(8),
		1: int32(5),
		2: int32(22),
		3: int32(16),
		4: int32(4),
		6: int32(15),
	},
	9: {
		0: int32(9),
		1: int32(19),
		2: int32(14),
		3: int32(20),
		4: int32(1),
		5: int32(7),
		6: int32(6),
	},
	10: {
		0: int32(10),
		1: int32(11),
		2: int32(24),
		3: int32(23),
		4: int32(5),
		5: int32(2),
		6: int32(18),
	},
	11: {
		0: int32(11),
		1: int32(17),
		2: int32(23),
		3: int32(25),
		4: int32(2),
		5: int32(6),
		6: int32(10),
	},
	12: {
		0: int32(12),
		1: int32(28),
		2: int32(13),
		3: int32(26),
		4: int32(4),
		5: int32(15),
		6: int32(3),
	},
	13: {
		0: int32(13),
		1: int32(26),
		2: int32(21),
		3: int32(29),
		4: int32(3),
		5: int32(12),
		6: int32(7),
	},
	14: {
		0: int32(14),
		1: int32(DINVALID_BASE_CELL),
		2: int32(17),
		3: int32(27),
		4: int32(9),
		5: int32(20),
		6: int32(6),
	},
	15: {
		0: int32(15),
		1: int32(22),
		2: int32(28),
		3: int32(31),
		4: int32(4),
		5: int32(8),
		6: int32(12),
	},
	16: {
		0: int32(16),
		1: int32(18),
		2: int32(33),
		3: int32(30),
		4: int32(8),
		5: int32(5),
		6: int32(22),
	},
	17: {
		0: int32(17),
		1: int32(11),
		2: int32(14),
		3: int32(6),
		4: int32(35),
		5: int32(25),
		6: int32(27),
	},
	18: {
		0: int32(18),
		1: int32(24),
		2: int32(30),
		3: int32(32),
		4: int32(5),
		5: int32(10),
		6: int32(16),
	},
	19: {
		0: int32(19),
		1: int32(34),
		2: int32(20),
		3: int32(36),
		4: int32(7),
		5: int32(21),
		6: int32(9),
	},
	20: {
		0: int32(20),
		1: int32(14),
		2: int32(19),
		3: int32(9),
		4: int32(40),
		5: int32(27),
		6: int32(36),
	},
	21: {
		0: int32(21),
		1: int32(38),
		2: int32(19),
		3: int32(34),
		4: int32(13),
		5: int32(29),
		6: int32(7),
	},
	22: {
		0: int32(22),
		1: int32(16),
		2: int32(41),
		3: int32(33),
		4: int32(15),
		5: int32(8),
		6: int32(31),
	},
	23: {
		0: int32(23),
		1: int32(24),
		2: int32(11),
		3: int32(10),
		4: int32(39),
		5: int32(37),
		6: int32(25),
	},
	24: {
		0: int32(24),
		1: int32(DINVALID_BASE_CELL),
		2: int32(32),
		3: int32(37),
		4: int32(10),
		5: int32(23),
		6: int32(18),
	},
	25: {
		0: int32(25),
		1: int32(23),
		2: int32(17),
		3: int32(11),
		4: int32(45),
		5: int32(39),
		6: int32(35),
	},
	26: {
		0: int32(26),
		1: int32(42),
		2: int32(29),
		3: int32(43),
		4: int32(12),
		5: int32(28),
		6: int32(13),
	},
	27: {
		0: int32(27),
		1: int32(40),
		2: int32(35),
		3: int32(46),
		4: int32(14),
		5: int32(20),
		6: int32(17),
	},
	28: {
		0: int32(28),
		1: int32(31),
		2: int32(42),
		3: int32(44),
		4: int32(12),
		5: int32(15),
		6: int32(26),
	},
	29: {
		0: int32(29),
		1: int32(43),
		2: int32(38),
		3: int32(47),
		4: int32(13),
		5: int32(26),
		6: int32(21),
	},
	30: {
		0: int32(30),
		1: int32(32),
		2: int32(48),
		3: int32(50),
		4: int32(16),
		5: int32(18),
		6: int32(33),
	},
	31: {
		0: int32(31),
		1: int32(41),
		2: int32(44),
		3: int32(53),
		4: int32(15),
		5: int32(22),
		6: int32(28),
	},
	32: {
		0: int32(32),
		1: int32(30),
		2: int32(24),
		3: int32(18),
		4: int32(52),
		5: int32(50),
		6: int32(37),
	},
	33: {
		0: int32(33),
		1: int32(30),
		2: int32(49),
		3: int32(48),
		4: int32(22),
		5: int32(16),
		6: int32(41),
	},
	34: {
		0: int32(34),
		1: int32(19),
		2: int32(38),
		3: int32(21),
		4: int32(54),
		5: int32(36),
		6: int32(51),
	},
	35: {
		0: int32(35),
		1: int32(46),
		2: int32(45),
		3: int32(56),
		4: int32(17),
		5: int32(27),
		6: int32(25),
	},
	36: {
		0: int32(36),
		1: int32(20),
		2: int32(34),
		3: int32(19),
		4: int32(55),
		5: int32(40),
		6: int32(54),
	},
	37: {
		0: int32(37),
		1: int32(39),
		2: int32(52),
		3: int32(57),
		4: int32(24),
		5: int32(23),
		6: int32(32),
	},
	38: {
		0: int32(38),
		1: int32(DINVALID_BASE_CELL),
		2: int32(34),
		3: int32(51),
		4: int32(29),
		5: int32(47),
		6: int32(21),
	},
	39: {
		0: int32(39),
		1: int32(37),
		2: int32(25),
		3: int32(23),
		4: int32(59),
		5: int32(57),
		6: int32(45),
	},
	40: {
		0: int32(40),
		1: int32(27),
		2: int32(36),
		3: int32(20),
		4: int32(60),
		5: int32(46),
		6: int32(55),
	},
	41: {
		0: int32(41),
		1: int32(49),
		2: int32(53),
		3: int32(61),
		4: int32(22),
		5: int32(33),
		6: int32(31),
	},
	42: {
		0: int32(42),
		1: int32(58),
		2: int32(43),
		3: int32(62),
		4: int32(28),
		5: int32(44),
		6: int32(26),
	},
	43: {
		0: int32(43),
		1: int32(62),
		2: int32(47),
		3: int32(64),
		4: int32(26),
		5: int32(42),
		6: int32(29),
	},
	44: {
		0: int32(44),
		1: int32(53),
		2: int32(58),
		3: int32(65),
		4: int32(28),
		5: int32(31),
		6: int32(42),
	},
	45: {
		0: int32(45),
		1: int32(39),
		2: int32(35),
		3: int32(25),
		4: int32(63),
		5: int32(59),
		6: int32(56),
	},
	46: {
		0: int32(46),
		1: int32(60),
		2: int32(56),
		3: int32(68),
		4: int32(27),
		5: int32(40),
		6: int32(35),
	},
	47: {
		0: int32(47),
		1: int32(38),
		2: int32(43),
		3: int32(29),
		4: int32(69),
		5: int32(51),
		6: int32(64),
	},
	48: {
		0: int32(48),
		1: int32(49),
		2: int32(30),
		3: int32(33),
		4: int32(67),
		5: int32(66),
		6: int32(50),
	},
	49: {
		0: int32(49),
		1: int32(DINVALID_BASE_CELL),
		2: int32(61),
		3: int32(66),
		4: int32(33),
		5: int32(48),
		6: int32(41),
	},
	50: {
		0: int32(50),
		1: int32(48),
		2: int32(32),
		3: int32(30),
		4: int32(70),
		5: int32(67),
		6: int32(52),
	},
	51: {
		0: int32(51),
		1: int32(69),
		2: int32(54),
		3: int32(71),
		4: int32(38),
		5: int32(47),
		6: int32(34),
	},
	52: {
		0: int32(52),
		1: int32(57),
		2: int32(70),
		3: int32(74),
		4: int32(32),
		5: int32(37),
		6: int32(50),
	},
	53: {
		0: int32(53),
		1: int32(61),
		2: int32(65),
		3: int32(75),
		4: int32(31),
		5: int32(41),
		6: int32(44),
	},
	54: {
		0: int32(54),
		1: int32(71),
		2: int32(55),
		3: int32(73),
		4: int32(34),
		5: int32(51),
		6: int32(36),
	},
	55: {
		0: int32(55),
		1: int32(40),
		2: int32(54),
		3: int32(36),
		4: int32(72),
		5: int32(60),
		6: int32(73),
	},
	56: {
		0: int32(56),
		1: int32(68),
		2: int32(63),
		3: int32(77),
		4: int32(35),
		5: int32(46),
		6: int32(45),
	},
	57: {
		0: int32(57),
		1: int32(59),
		2: int32(74),
		3: int32(78),
		4: int32(37),
		5: int32(39),
		6: int32(52),
	},
	58: {
		0: int32(58),
		1: int32(DINVALID_BASE_CELL),
		2: int32(62),
		3: int32(76),
		4: int32(44),
		5: int32(65),
		6: int32(42),
	},
	59: {
		0: int32(59),
		1: int32(63),
		2: int32(78),
		3: int32(79),
		4: int32(39),
		5: int32(45),
		6: int32(57),
	},
	60: {
		0: int32(60),
		1: int32(72),
		2: int32(68),
		3: int32(80),
		4: int32(40),
		5: int32(55),
		6: int32(46),
	},
	61: {
		0: int32(61),
		1: int32(53),
		2: int32(49),
		3: int32(41),
		4: int32(81),
		5: int32(75),
		6: int32(66),
	},
	62: {
		0: int32(62),
		1: int32(43),
		2: int32(58),
		3: int32(42),
		4: int32(82),
		5: int32(64),
		6: int32(76),
	},
	63: {
		0: int32(63),
		1: int32(DINVALID_BASE_CELL),
		2: int32(56),
		3: int32(45),
		4: int32(79),
		5: int32(59),
		6: int32(77),
	},
	64: {
		0: int32(64),
		1: int32(47),
		2: int32(62),
		3: int32(43),
		4: int32(84),
		5: int32(69),
		6: int32(82),
	},
	65: {
		0: int32(65),
		1: int32(58),
		2: int32(53),
		3: int32(44),
		4: int32(86),
		5: int32(76),
		6: int32(75),
	},
	66: {
		0: int32(66),
		1: int32(67),
		2: int32(81),
		3: int32(85),
		4: int32(49),
		5: int32(48),
		6: int32(61),
	},
	67: {
		0: int32(67),
		1: int32(66),
		2: int32(50),
		3: int32(48),
		4: int32(87),
		5: int32(85),
		6: int32(70),
	},
	68: {
		0: int32(68),
		1: int32(56),
		2: int32(60),
		3: int32(46),
		4: int32(90),
		5: int32(77),
		6: int32(80),
	},
	69: {
		0: int32(69),
		1: int32(51),
		2: int32(64),
		3: int32(47),
		4: int32(89),
		5: int32(71),
		6: int32(84),
	},
	70: {
		0: int32(70),
		1: int32(67),
		2: int32(52),
		3: int32(50),
		4: int32(83),
		5: int32(87),
		6: int32(74),
	},
	71: {
		0: int32(71),
		1: int32(89),
		2: int32(73),
		3: int32(91),
		4: int32(51),
		5: int32(69),
		6: int32(54),
	},
	72: {
		0: int32(72),
		1: int32(DINVALID_BASE_CELL),
		2: int32(73),
		3: int32(55),
		4: int32(80),
		5: int32(60),
		6: int32(88),
	},
	73: {
		0: int32(73),
		1: int32(91),
		2: int32(72),
		3: int32(88),
		4: int32(54),
		5: int32(71),
		6: int32(55),
	},
	74: {
		0: int32(74),
		1: int32(78),
		2: int32(83),
		3: int32(92),
		4: int32(52),
		5: int32(57),
		6: int32(70),
	},
	75: {
		0: int32(75),
		1: int32(65),
		2: int32(61),
		3: int32(53),
		4: int32(94),
		5: int32(86),
		6: int32(81),
	},
	76: {
		0: int32(76),
		1: int32(86),
		2: int32(82),
		3: int32(96),
		4: int32(58),
		5: int32(65),
		6: int32(62),
	},
	77: {
		0: int32(77),
		1: int32(63),
		2: int32(68),
		3: int32(56),
		4: int32(93),
		5: int32(79),
		6: int32(90),
	},
	78: {
		0: int32(78),
		1: int32(74),
		2: int32(59),
		3: int32(57),
		4: int32(95),
		5: int32(92),
		6: int32(79),
	},
	79: {
		0: int32(79),
		1: int32(78),
		2: int32(63),
		3: int32(59),
		4: int32(93),
		5: int32(95),
		6: int32(77),
	},
	80: {
		0: int32(80),
		1: int32(68),
		2: int32(72),
		3: int32(60),
		4: int32(99),
		5: int32(90),
		6: int32(88),
	},
	81: {
		0: int32(81),
		1: int32(85),
		2: int32(94),
		3: int32(101),
		4: int32(61),
		5: int32(66),
		6: int32(75),
	},
	82: {
		0: int32(82),
		1: int32(96),
		2: int32(84),
		3: int32(98),
		4: int32(62),
		5: int32(76),
		6: int32(64),
	},
	83: {
		0: int32(83),
		1: int32(DINVALID_BASE_CELL),
		2: int32(74),
		3: int32(70),
		4: int32(100),
		5: int32(87),
		6: int32(92),
	},
	84: {
		0: int32(84),
		1: int32(69),
		2: int32(82),
		3: int32(64),
		4: int32(97),
		5: int32(89),
		6: int32(98),
	},
	85: {
		0: int32(85),
		1: int32(87),
		2: int32(101),
		3: int32(102),
		4: int32(66),
		5: int32(67),
		6: int32(81),
	},
	86: {
		0: int32(86),
		1: int32(76),
		2: int32(75),
		3: int32(65),
		4: int32(104),
		5: int32(96),
		6: int32(94),
	},
	87: {
		0: int32(87),
		1: int32(83),
		2: int32(102),
		3: int32(100),
		4: int32(67),
		5: int32(70),
		6: int32(85),
	},
	88: {
		0: int32(88),
		1: int32(72),
		2: int32(91),
		3: int32(73),
		4: int32(99),
		5: int32(80),
		6: int32(105),
	},
	89: {
		0: int32(89),
		1: int32(97),
		2: int32(91),
		3: int32(103),
		4: int32(69),
		5: int32(84),
		6: int32(71),
	},
	90: {
		0: int32(90),
		1: int32(77),
		2: int32(80),
		3: int32(68),
		4: int32(106),
		5: int32(93),
		6: int32(99),
	},
	91: {
		0: int32(91),
		1: int32(73),
		2: int32(89),
		3: int32(71),
		4: int32(105),
		5: int32(88),
		6: int32(103),
	},
	92: {
		0: int32(92),
		1: int32(83),
		2: int32(78),
		3: int32(74),
		4: int32(108),
		5: int32(100),
		6: int32(95),
	},
	93: {
		0: int32(93),
		1: int32(79),
		2: int32(90),
		3: int32(77),
		4: int32(109),
		5: int32(95),
		6: int32(106),
	},
	94: {
		0: int32(94),
		1: int32(86),
		2: int32(81),
		3: int32(75),
		4: int32(107),
		5: int32(104),
		6: int32(101),
	},
	95: {
		0: int32(95),
		1: int32(92),
		2: int32(79),
		3: int32(78),
		4: int32(109),
		5: int32(108),
		6: int32(93),
	},
	96: {
		0: int32(96),
		1: int32(104),
		2: int32(98),
		3: int32(110),
		4: int32(76),
		5: int32(86),
		6: int32(82),
	},
	97: {
		0: int32(97),
		1: int32(DINVALID_BASE_CELL),
		2: int32(98),
		3: int32(84),
		4: int32(103),
		5: int32(89),
		6: int32(111),
	},
	98: {
		0: int32(98),
		1: int32(110),
		2: int32(97),
		3: int32(111),
		4: int32(82),
		5: int32(96),
		6: int32(84),
	},
	99: {
		0: int32(99),
		1: int32(80),
		2: int32(105),
		3: int32(88),
		4: int32(106),
		5: int32(90),
		6: int32(113),
	},
	100: {
		0: int32(100),
		1: int32(102),
		2: int32(83),
		3: int32(87),
		4: int32(108),
		5: int32(114),
		6: int32(92),
	},
	101: {
		0: int32(101),
		1: int32(102),
		2: int32(107),
		3: int32(112),
		4: int32(81),
		5: int32(85),
		6: int32(94),
	},
	102: {
		0: int32(102),
		1: int32(101),
		2: int32(87),
		3: int32(85),
		4: int32(114),
		5: int32(112),
		6: int32(100),
	},
	103: {
		0: int32(103),
		1: int32(91),
		2: int32(97),
		3: int32(89),
		4: int32(116),
		5: int32(105),
		6: int32(111),
	},
	104: {
		0: int32(104),
		1: int32(107),
		2: int32(110),
		3: int32(115),
		4: int32(86),
		5: int32(94),
		6: int32(96),
	},
	105: {
		0: int32(105),
		1: int32(88),
		2: int32(103),
		3: int32(91),
		4: int32(113),
		5: int32(99),
		6: int32(116),
	},
	106: {
		0: int32(106),
		1: int32(93),
		2: int32(99),
		3: int32(90),
		4: int32(117),
		5: int32(109),
		6: int32(113),
	},
	107: {
		0: int32(107),
		1: int32(DINVALID_BASE_CELL),
		2: int32(101),
		3: int32(94),
		4: int32(115),
		5: int32(104),
		6: int32(112),
	},
	108: {
		0: int32(108),
		1: int32(100),
		2: int32(95),
		3: int32(92),
		4: int32(118),
		5: int32(114),
		6: int32(109),
	},
	109: {
		0: int32(109),
		1: int32(108),
		2: int32(93),
		3: int32(95),
		4: int32(117),
		5: int32(118),
		6: int32(106),
	},
	110: {
		0: int32(110),
		1: int32(98),
		2: int32(104),
		3: int32(96),
		4: int32(119),
		5: int32(111),
		6: int32(115),
	},
	111: {
		0: int32(111),
		1: int32(97),
		2: int32(110),
		3: int32(98),
		4: int32(116),
		5: int32(103),
		6: int32(119),
	},
	112: {
		0: int32(112),
		1: int32(107),
		2: int32(102),
		3: int32(101),
		4: int32(120),
		5: int32(115),
		6: int32(114),
	},
	113: {
		0: int32(113),
		1: int32(99),
		2: int32(116),
		3: int32(105),
		4: int32(117),
		5: int32(106),
		6: int32(121),
	},
	114: {
		0: int32(114),
		1: int32(112),
		2: int32(100),
		3: int32(102),
		4: int32(118),
		5: int32(120),
		6: int32(108),
	},
	115: {
		0: int32(115),
		1: int32(110),
		2: int32(107),
		3: int32(104),
		4: int32(120),
		5: int32(119),
		6: int32(112),
	},
	116: {
		0: int32(116),
		1: int32(103),
		2: int32(119),
		3: int32(111),
		4: int32(113),
		5: int32(105),
		6: int32(121),
	},
	117: {
		0: int32(117),
		1: int32(DINVALID_BASE_CELL),
		2: int32(109),
		3: int32(118),
		4: int32(113),
		5: int32(121),
		6: int32(106),
	},
	118: {
		0: int32(118),
		1: int32(120),
		2: int32(108),
		3: int32(114),
		4: int32(117),
		5: int32(121),
		6: int32(109),
	},
	119: {
		0: int32(119),
		1: int32(111),
		2: int32(115),
		3: int32(110),
		4: int32(121),
		5: int32(116),
		6: int32(120),
	},
	120: {
		0: int32(120),
		1: int32(115),
		2: int32(114),
		3: int32(112),
		4: int32(121),
		5: int32(119),
		6: int32(118),
	},
	121: {
		0: int32(121),
		1: int32(116),
		2: int32(120),
		3: int32(119),
		4: int32(117),
		5: int32(113),
		6: int32(118),
	},
}

/** square root of 7 and inverse square root of 7 */

// C documentation
//
//	/** @brief icosahedron face centers in lat/lng radians */
var XfaceCenterGeo = [20]TLatLng{
	0: {
		Flat: float64(0.80358264971899),
		Flng: float64(1.2483974196173961),
	},
	1: {
		Flat: float64(1.3077478834556382),
		Flng: float64(2.5369450098779214),
	},
	2: {
		Flat: float64(1.054751253523952),
		Flng: -libc.Float64FromFloat64(1.3475173589003966),
	},
	3: {
		Flat: float64(0.6001915955381868),
		Flng: -libc.Float64FromFloat64(0.45060390946975576),
	},
	4: {
		Flat: float64(0.49171542819877384),
		Flng: float64(0.40198820291130694),
	},
	5: {
		Flat: float64(0.1727453274156187),
		Flng: float64(1.6781468852804338),
	},
	6: {
		Flat: float64(0.6059293215713507),
		Flng: float64(2.9539233298124117),
	},
	7: {
		Flat: float64(0.42737051832897965),
		Flng: -libc.Float64FromFloat64(1.8888762003362853),
	},
	8: {
		Flat: -libc.Float64FromFloat64(0.07906611854921283),
		Flng: -libc.Float64FromFloat64(0.7334295133808677),
	},
	9: {
		Flat: -libc.Float64FromFloat64(0.23096164445538364),
		Flng: float64(0.506495587332349),
	},
	10: {
		Flat: float64(0.07906611854921283),
		Flng: float64(2.4081631402089254),
	},
	11: {
		Flat: float64(0.23096164445538364),
		Flng: -libc.Float64FromFloat64(2.635097066257444),
	},
	12: {
		Flat: -libc.Float64FromFloat64(0.1727453274156187),
		Flng: -libc.Float64FromFloat64(1.4634457683093596),
	},
	13: {
		Flat: -libc.Float64FromFloat64(0.6059293215713507),
		Flng: -libc.Float64FromFloat64(0.18766932377738163),
	},
	14: {
		Flat: -libc.Float64FromFloat64(0.42737051832897965),
		Flng: float64(1.2527164532535078),
	},
	15: {
		Flat: -libc.Float64FromFloat64(0.6001915955381868),
		Flng: float64(2.6909887441200375),
	},
	16: {
		Flat: -libc.Float64FromFloat64(0.49171542819877384),
		Flng: -libc.Float64FromFloat64(2.7396044506784865),
	},
	17: {
		Flat: -libc.Float64FromFloat64(0.80358264971899),
		Flng: -libc.Float64FromFloat64(1.8931952339723972),
	},
	18: {
		Flat: -libc.Float64FromFloat64(1.3077478834556382),
		Flng: -libc.Float64FromFloat64(0.6046476437118721),
	},
	19: {
		Flat: -libc.Float64FromFloat64(1.054751253523952),
		Flng: float64(1.7940752946893965),
	},
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "0\x00../src/h3lib/lib/algos.c\x00../src/h3lib/lib/coordijk.c\x00../src/h3lib/lib/directedEdge.c\x00adjacentFaceDir[tmpFijk.face][fijk.face] == KI\x00../src/h3lib/lib/faceijk.c\x00adjacentFaceDir[centerIJK.face][face2] == KI\x00Success\x00The operation failed but a more specific error is not available\x00Argument was outside of acceptable range\x00Latitude or longitude arguments were outside of acceptable range\x00Resolution argument was outside of acceptable range\x00Cell argument was not valid\x00Directed edge argument was not valid\x00Undirected edge argument was not valid\x00Vertex argument was not valid\x00Pentagon distortion was encountered\x00Duplicate input\x00Cell arguments were not neighbors\x00Cell arguments had incompatible resolutions\x00Memory allocation failed\x00Bounds of provided memory were insufficient\x00Mode or flags argument was not valid\x00Invalid error code\x00%lx\x00../src/h3lib/lib/h3Index.c\x00../src/h3lib/lib/latLng.c\x00polygon->next == NULL\x00../src/h3lib/lib/linkedGeo.c\x00next != NULL\x00loop != NULL\x00polygon->first == NULL\x00coord != NULL\x00loop->first == NULL\x00candidates != NULL\x00candidateBBoxes != NULL\x00innerLoops != NULL\x00bboxes != NULL\x00../src/h3lib/lib/localij.c\x00revDir != INVALID_DIGIT\x00baseCell != originBaseCell\x00!(originOnPent && indexOnPent)\x00baseCell == originBaseCell\x00baseCell != INVALID_BASE_CELL\x00!_isBaseCellPentagon(baseCell)\x00baseCellRotations >= 0\x00../src/h3lib/lib/polyfill.c\x00../src/h3lib/lib/vertex.c\x00graph->buckets != NULL\x00../src/h3lib/lib/vertexGraph.c\x00node != NULL\x00"
