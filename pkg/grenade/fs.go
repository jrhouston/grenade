package grenade

import (
	"fmt"
	"os"
	"time"

	"github.com/icrowley/fake"
)

// OpenFileLeak steadily opens files and never closes them
func OpenFileLeak(delay time.Duration) {
	files := []*os.File{}

	for {
		randomFile := fmt.Sprintf("/tmp/%s%s.txt", fake.Word(), fake.Digits())
		f, err := os.Create(randomFile)
		if err != nil {
			continue
		}
		files = append(files, f)
		time.Sleep(delay)
	}
}

// DiskUsageLeak opens a file and steadily eats disk space
func DiskUsageLeak() {

}
