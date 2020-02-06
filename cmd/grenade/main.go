package main

import (
	"fmt"
	"time"

	"github.com/jrhouston/grenade/pkg/grenade"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/jrhouston/grenade/internal/pkg/webserver"
)

var (
	webserverPort = kingpin.Flag("port", "Webserver port").Default("8080").Int()

	delay = kingpin.Flag("delay", "trigger the grenade after specified delay e.g 5m30s").Duration()

	// NOTE kingpin doesn't yet support mutually exlusive flags
	crash = kingpin.Flag("crash", "Process crash").Bool()
	exit  = kingpin.Flag("exit", "Exit process with non-zero exit code").Bool()

	memSpike = kingpin.Flag("mem-spike", "Cause a memory spike").Bool()
	memLeak  = kingpin.Flag("mem-leak", "Cause a memory leak").Bool()

	cpuSpike      = kingpin.Flag("cpu-spike", "Cause a CPU spike").Bool()
	goroutineLeak = kingpin.Flag("goroutine-leak", "Cause goroutines to be leaked").Bool()

	connSpike = kingpin.Flag("conn-spike", "Cause a connect spike to an address").String()
	connLeak  = kingpin.Flag("conn-leak", "Cause a connect leak to an address").String()
	dnsSpike  = kingpin.Flag("dns-spike", "Cause a spike in DNS lookups").Bool()

	logSpike = kingpin.Flag("log-spike", "Cause a spike in log volume").Bool()

	fdSpike        = kingpin.Flag("fd-spike", "Cause a spike in open file descriptors").Bool()
	diskUsageSpike = kingpin.Flag("disk-spike", "Cause a spike in disk space usage").Bool()
	diskUsageLeak  = kingpin.Flag("disk-leak", "Cause a leak in disk space usage").Bool()
)

func main() {
	kingpin.Parse()

	go webserver.Run(*webserverPort)

	fmt.Println("💣 ... 🔥")

	if *delay != 0 {
		time.Sleep(*delay)
	}

	if *crash {
		grenade.Crash()
	}

	if *exit {
		grenade.Exit()
	}

	if *memSpike {
		grenade.MemoryLeak(0)
	}

	if *memLeak {
		grenade.MemoryLeak(1 * time.Second)
	}

	if *cpuSpike {
		grenade.CPUSpike()
	}

	if *connSpike != "" {
		grenade.ConnectLeak(*connSpike, 0)
	}

	if *connLeak != "" {
		grenade.ConnectLeak(*connLeak, 10*time.Second)
	}

	if *goroutineLeak {
		grenade.GoroutineLeak()
	}

	if *dnsSpike {
		grenade.DNSLookupSpike()
	}

	if *logSpike {
		grenade.LogVolumeSpike()
	}

	if *fdSpike {
		grenade.OpenFileLeak(0)
	}

	if *diskUsageSpike {
		grenade.DiskUsageLeak(0)
	}

	if *diskUsageLeak {
		grenade.DiskUsageLeak(500 * time.Millisecond)
	}
}
