package main

import (
	"flag"
	"fmt"
	"net"
	"net/url"
	"time"
)

func main() {
	flag.Parse()
	endpoints := flag.Args()
	endpoint := endpoints[0]

	parsedEndpoint, err := url.Parse(endpoint)
	if err != nil {
		panic(err)
	}

	startTime := time.Now()
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", parsedEndpoint.Host, parsedEndpoint.Scheme))
	elapsedTime := time.Now().Sub(startTime)


	defer conn.Close()

	if err != nil {
		panic(err)
	}

	localAddr := conn.LocalAddr()
	remoteAddr := conn.RemoteAddr()

	ip, err := net.LookupIP(parsedEndpoint.Host)
	if err != nil {
		panic(err)
	}

	lookupNS, err := net.LookupNS(parsedEndpoint.Host)
	if err != nil {
		panic(err)
	}
	nameServers := make([]string, len(lookupNS))
	for i, ns := range lookupNS {
		nameServers[i] = ns.Host
	}

	fmt.Printf("gurl Response:\n" +
		"\tLocal Address: %s\n"+
		"\tRemote Address: %s\n"+
		"\tLookup Address: %s\n"+
		"\tLookup NameServers: %s\n" +
		"\tRound Trip Time (RTT): %s\n", localAddr, remoteAddr, ip, nameServers, elapsedTime.String())
}
