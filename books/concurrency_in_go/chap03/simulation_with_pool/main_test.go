package main

import (
	"io/ioutil"
	"net"
	"testing"
)

func init() {
	daemonStarted := startNetworkDaemon()
	daemonStarted.Wait()
}

func BenchmarkNetworkRequest(b *testing.B) {
	for range b.N {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("can't dial host: %v", err)
		}
		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("can't read: %v", err)
		}
		conn.Close()
	}
}
