package vm;

import (
	"math"
)


/*
 * INSTRUCTION SET
 *
 * WIP, evolving as satunna is developed.
 * This file contains the instruction set
 * for the satunnaVM.
 */

//TODO: fix this shit
/* IADD_M
 * adds array to itself
 */
func IADD_M(src []uint64) []uint64 {
	ret := make([]uint64, len(src));
	for i := 0; i < len(src); i++ {
		ret[i] = src[i] + src[i]; 
	}

	return ret;
}


/* ISUB_R
 * sub with IMUL_M array w/ overflow
 */
func ISUB_R(src []uint64, dst []uint64) []uint64 {
	ret := make([]uint64, len(src));
	for i := 0; i < len(src); i++ {
		ret[i] = src[i] - dst[i]; 
	}

	return ret;
}


/* IMUL_M
 * adds array to itself
 */
func IMUL_M(src []uint64) []uint64 {
	ret := make([]uint64, len(src));
	/*we dont want to mulyiply by 0 here*/
	for i := 0; i < len(src); i++ {
		if src[i] != 0 {
			ret[i] = src[i] * src[i];
		}
	}

	return ret;
}

// TODO: replace this with something that doesnt divide
/* IMULH_M
 * adds array to itself and then divided by 2 2 times
 */
func IMULH_M(src []uint64) []uint64 {
	ret := make([]uint64, len(src));
	/*we dont want to mulyiply by 0 here*/
	for i := 0; i < len(src); i++ {
		if src[i] != 0 {
			ret[i] = (src[i] * src[i]) >> 2;
		}
	}

	return ret;
}


/* IMULS_M
 * multiplies array with seed
 */
func IMULS_M(src []uint64, seed uint32) []uint64 {
	ret := make([]uint64, len(src));
	for i := 0; i < len(src); i++ {
		ret[i] = (src[i] * src[i]) * uint64(seed);
	}

	return ret;
}


/* IXOR_R
 * XOR with IMUL_M array
 */
func IXOR_R(src []uint64, dst []uint64) []uint64 {
	ret := make([]uint64, len(src));
	for i := 0; i < len(src); i++ {
		ret[i] = src[i] ^ dst[i]; 
	}

	return ret;

}


/* IOR_R
 * OR with IMUL_M array
 */
func IOR_R(src []uint64, dst []uint64) []uint64 {
	ret := make([]uint64, len(src));
	for i := 0; i < len(src); i++ {
		ret[i] = src[i] | dst[i]; 
	}

	return ret;

}


/* ISQR_R
 * square root of itself
 */
func ISQR_R(src []uint64) []uint64 {
	ret := make([]uint64, len(src));
	for i := 0; i < len(src); i++ {
		ret[i] = uint64(math.Sqrt(float64(src[i]))); 
	}

	return ret;
}


/* ISIN_R
 * sine of itself
 */
func ISIN_R(src []uint64) []uint64 {
	ret := make([]uint64, len(src));
	for i := 0; i < len(src); i++ {
		ret[i] = uint64(math.Sin(float64(src[i]))); 
	}

	return ret;
}


