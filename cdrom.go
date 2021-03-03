package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"networks"
	"os"
	"runtime"
	"somecall"
	"sysexit"
)

func main() {

	if runtime.GOOS != "linux" {
		fmt.Println("this application is supposed to run on linux, can't do anything here... \nexiting with status 0")
		return
	}
	defer sysexit.HandlePan()
	log.SetFlags(0)
	flag.Usage = flagusg
	laddrs := flag.String("net", "127.0.0.1:9940", "server tcp listen address")
	udevpath := flag.String("dev", "/dev/cdrom", "rom block device to mess with")
	flag.Parse()
	if somecall.Getuid() != 0 {
		panic(sysexit.Errsig{Why: errors.New("without superuser access this application cant do anything"), Cod: 1, Pid: somecall.Getpid()})
	}
	somecall.UDEVLINK = *udevpath
	networks.NetListen(*laddrs)
}

func flagusg() {
	fmt.Printf(`usage of dummy cdrom-dvdrom ejector server:
%v --net [ipv4:port] --dev <rom block device>
options:
	--net string     tcp network to listen on <ipv4:port> (default "127.0.0.1:9940")
	--dev string	 rom block device to mess with. (default is /dev/cdrom)
	--h              print this banner and exit
example: 
	%v --net 0.0.0.0:9940 --dev /dev/cdrom

Copyright (c) 2021 blg.snix.ir, All rights reserved.
Developed BY sina@snix.ir --> FYI: this is just for fun and totally useless :D
This work is licensed under the terms of the MIT license.
`, os.Args[0], os.Args[0])
}
