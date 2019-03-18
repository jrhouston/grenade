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

var fileBlockSize = 1073741824 // 1 MB
var leakFile = "/tmp/leak.txt"

// DiskUsageLeak opens a file and steadily eats disk space
func DiskUsageLeak(delay time.Duration) {
	f, err := os.Create(leakFile)
	defer f.Close()

	if err != nil {
		panic(err)
	}

	for {
		f.Write(make([]byte, fileBlockSize))
		time.Sleep(delay)
	}
}
