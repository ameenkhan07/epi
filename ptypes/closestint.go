// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package ptypes

// IntWeight returns number of bits set to 1 in x.
func IntWeight(x uint64) (w int) {
	for x > 0 {
		w++
		x &= (x - 1)
	}
	return w
}

// ClosestInt returns integer closest to x with the same weight.
// The time complexity is O(n), where n is the integer width.
// The O(1) additional space is needed.
func ClosestInt(x uint64) (ci uint64, ok bool) {
	for i := uint(0); i < 63; i++ {
		if (x>>i)&1 != x>>(i+1)&1 {
			x ^= 1<<i | 1<<(i+1)
			return x, true
		}
	}
	return 0, false // If all bits are 0 or 1.
}
