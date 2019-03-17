package grenade

import "time"

var blockSize = 1073741824 // 1 MB

// MemoryLeak steadily allocates BlockSize of memory never freeing it
func MemoryLeak(delay time.Duration) {
	data := [][]byte{}

	for {
		data = append(data, make([]byte, blockSize))
		time.Sleep(delay)
	}
}
