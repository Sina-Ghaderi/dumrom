package somecall

import (
	"os"
	"syscall"
	"sysexit"
)

// lets make some calls for you ;) pull a few strings.. and see whats happened
const (
	SYSCDROMEJECTSW = 0x530f
	SYSCDROMEJECT   = 0x5309
	CDROMCLOSETRAY  = 0x5319
)

// UDEVLINK --> change this if u want mess with another device
var UDEVLINK string

// Getpid ...
func Getpid() uintptr {
	pid, _, err := syscall.Syscall(syscall.SYS_GETPID, 0, 0, 0)
	if err != 0 {
		panic(sysexit.Errsig{Why: os.NewSyscallError("SYS_GETPID", err), Cod: 1})
	}
	return pid
}

// Getuid ...
func Getuid() uintptr {
	uid, _, err := syscall.Syscall(syscall.SYS_GETUID, 0, 0, 0)
	if err != 0 {
		panic(sysexit.Errsig{Why: os.NewSyscallError("SYS_GETPID", err), Cod: 1})
	}
	return uid
}

// TakeAction ...
var TakeAction = map[string]func(device string) error{
	"1": func(device string) error {
		rom, err := os.OpenFile(device, syscall.O_RDONLY|syscall.O_NONBLOCK, 0666)
		if err != nil {
			return err
		}
		defer rom.Close()

		_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, rom.Fd(), uintptr(SYSCDROMEJECT), 0)
		if errno != 0 {
			return os.NewSyscallError("SYS_IOCTL", errno)
		}
		_, _, errno = syscall.Syscall(syscall.SYS_IOCTL, rom.Fd(), uintptr(SYSCDROMEJECT), 0)
		if errno != 0 {
			return os.NewSyscallError("SYS_IOCTL", errno)
		}

		return nil

	},
	"0": func(device string) error {
		rom, err := os.OpenFile(device, syscall.O_RDONLY|syscall.O_NONBLOCK, 0666)
		if err != nil {
			return err
		}
		defer rom.Close()

		_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, rom.Fd(), uintptr(CDROMCLOSETRAY), 0)
		if errno != 0 {
			return os.NewSyscallError("SYS_IOCTL", errno)
		}

		return nil
	},
}
