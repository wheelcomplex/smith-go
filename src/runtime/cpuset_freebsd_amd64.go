// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

func getncpu() int32 {
	var mask [_CPU_SETSIZE_MAX]uintptr
	var size int
	for size = _CPU_SETSIZE_MAX; size >= _CPU_SETSIZE_MIN; size -= 8 {
		if cpuset_getaffinity(_CPU_LEVEL_WHICH, _CPU_WHICH_PID, _CPU_CURRENT_PID, size, &mask[0]) == 0 {
			break
		}
	}
	if size < _CPU_SETSIZE_MIN {
		// probe CPU_SETSIZE failed
		return int32(1)
	}
	n := int32(0)
	for _, v := range mask[:size] {
		for v != 0 {
			n += int32(v & 1)
			v >>= 1
		}
	}
	if n == 0 {
		n = 1
	}
	return n
}

// id_t is 64 bit
//go:noescape
func cpuset_getaffinity(level int, which int, id int, len int, mask *uintptr) int32
