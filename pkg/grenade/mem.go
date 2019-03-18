package grenade

import "time"

var memoryBlockSize = 1073741824 // 1 MB

// MemoryLeak steadily allocates memoryBlockSize of memory never freeing it
func MemoryLeak(delay time.Duration) {
	data := [][]byte{}

	for {
		data = append(data, make([]byte, memoryBlockSize))
		time.Sleep(delay)
	}
}
