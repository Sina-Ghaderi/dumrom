package sysexit

import (
	"fmt"
	"log"
	"os"
)

const (
	reset string = "\033[0m"
	erred        = "\033[31m"
	yello        = "\033[33m"
)

// Errsig ...
type Errsig struct {
	Why error
	Cod int
	Pid uintptr
}

//HandlePan ...
func HandlePan() {
	if hap := recover(); hap != nil {
		if ms, owkey := hap.(Errsig); owkey {
			fmt.Println(erred+"fatal:"+reset, ms.Why, "\nprocess", ms.Pid, "exit with ststus", ms.Cod)
			os.Exit(ms.Cod)
		}
		panic(hap)
	}
}

// Inform ...
func Inform(h ...interface{}) {
	h = append([]interface{}{(yello + "info:" + reset)}, h...)
	log.Println(h...)
}
