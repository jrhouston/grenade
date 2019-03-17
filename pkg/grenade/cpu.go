package grenade

import (
	"runtime"
	"time"
)

// CPUSpike spawns a goroutine with an empty for loop for every CPU
func CPUSpike() {
	for i := 0; i <= runtime.NumCPU(); i++ {
		go func() {
			for {
			}
		}()
	}

	select {}
}

// GoroutineLeak will steadily create goroutines that never finish
func GoroutineLeak() {
	for {
		go func() {
			select {}
		}()

		time.Sleep(1 * time.Minute)
	}
}
