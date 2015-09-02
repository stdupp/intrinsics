package prefetchwt1

import . "github.com/klauspost/intrinsics/x86"


// Prefetch: Fetch the line of data from memory that contains address 'p' to a
// location in the cache heirarchy specified by the locality hint 'i'. 
//
//		
//
// Instruction: 'PREFETCHWT1'. Intrinsic: '_mm_prefetch'.
// Requires PREFETCHWT1.
func Prefetch(p byte, i int)  {
	prefetch(p, i)
}

func prefetch(p byte, i int) 
