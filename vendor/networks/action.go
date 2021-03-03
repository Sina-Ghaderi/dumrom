package networks

import (
	"bufio"
	"fmt"
	"net"
	"somecall"
	"sysexit"
)

// NetListen ...
func NetListen(listenaddr string) {
	l, err := net.Listen("tcp", listenaddr)
	if err != nil {
		panic(sysexit.Errsig{Why: err, Cod: 1, Pid: somecall.Getpid()})
	}

	defer l.Close()
	sysexit.Inform("application is listening on address", listenaddr)
	for {
		conn, err := l.Accept()
		if err != nil {
			sysexit.Inform(err.Error())
			continue
		}

		go cachethecon(conn)
	}

}

func cachethecon(con net.Conn) {
	defer func() {
		con.Close()
		sysexit.Inform(con.RemoteAddr(), "connection closed by peer")
	}()

	sysexit.Inform(con.RemoteAddr(), "connection established")
	if _, err := con.Write([]byte("help banner for this terminal: system cdrom --> type 1 to eject and 0 to rollin \n")); err != nil {
		sysexit.Inform(err.Error())
		return
	}
	scanner := bufio.NewScanner(con)
	for scanner.Scan() {
		if makeanaction, ok := somecall.TakeAction[scanner.Text()]; ok {
			if err := makeanaction(somecall.UDEVLINK); err != nil {
				fmt.Fprintln(con, "fatal error: "+err.Error())
				continue
			}
			fmt.Fprintln(con, "operation completed successfully, check the cdrom.")
			continue
		}
		fmt.Fprintln(con, "fatal error: unknow input parameters, please input either 1 or 0")
	}
	if err := scanner.Err(); err != nil {
		sysexit.Inform(err.Error())
		return
	}
}
