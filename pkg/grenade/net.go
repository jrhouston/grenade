package grenade

import (
	"net"
	"time"

	"github.com/icrowley/fake"
)

// ConnectLeak will steadily open TCP connections and never close them
func ConnectLeak(address string, delay time.Duration) {
	conns := []net.Conn{}

	for {
		conn, _ := net.Dial("tcp", address)
		conns = append(conns, conn)
		time.Sleep(delay)
	}
}

// DNSLookupSpike will cause a spike in DNS lookups to randomly generated domains
func DNSLookupSpike() {
	for {
		domain := fake.DomainName()
		net.LookupIP(domain)
		time.Sleep(10 * time.Millisecond)
	}
}
