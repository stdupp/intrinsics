// THESE PACKAGES ARE FOR DEMONSTRATION PURPOSES ONLY!
//
// THEY DO NOT NOT CONTAIN WORKING INTRINSICS!
//
// See https://github.com/klauspost/intrinsics
package clflushopt

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// Skipped: _mm_clflushopt. Contains pointer parameter.

