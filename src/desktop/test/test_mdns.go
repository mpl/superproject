package main

import (
	"context"
	"flag"
	"log"
	"net"
	"time"

	"github.com/grandcat/zeroconf"
)

var browse = false

func init() {
	flag.BoolVar(&browse, "browse", false, "whether to browse")
}

func withZero(appName string) {
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		log.Fatalln("Failed to initialize resolver:", err.Error())
	}

	entries := make(chan *zeroconf.ServiceEntry)
	go func(results <-chan *zeroconf.ServiceEntry) {
		for entry := range results {
			log.Println(entry)
		}
		log.Println("No more entries.")
	}(entries)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if browse {
		err = resolver.Browse(ctx, "_http._tcp", "local.", entries)
		if err != nil {
			log.Fatalln("Failed to browse:", err.Error())
		}
	} else {
		if err := resolver.Lookup(ctx, appName, "_http._tcp", "local.", entries); err != nil {
			log.Fatal(err)
		}
		time.Sleep(5 * time.Second)
	}
	<-ctx.Done()
}

func main() {
	appName := "ftcDesktop"

	flag.Parse()

	withZero(appName)

	// test that we actually don't even need a lookup, as things just resolve by themselves.
	addr, err := net.LookupHost(appName + ".local")
	if err != nil {
		log.Fatal(err)
	}
	println("IMPLICIT SERVICE IP:", addr[0])
}

/*
	var iface *net.Interface
	if runtime.GOOS == "windows" {
		ifaces, _ := net.Interfaces()
		for _, i := range ifaces {
			fmt.Println(i.Name)
			if strings.Contains(i.Name, "Ethernet") {
				iface = &net.Interface{}
				*iface = i
				addrs, _ := i.Addrs()
				for _, a := range addrs {
					fmt.Println(a)
				}
				break
			}
		}
		if iface == nil {
			log.Fatal("SUITABLE INTERFACE NOT FOUND")
		}
	}
*/
