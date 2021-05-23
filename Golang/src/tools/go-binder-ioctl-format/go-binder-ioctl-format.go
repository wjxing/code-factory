package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
    _IOC_NRBITS     = 8
    _IOC_TYPEBITS   = 8
    _IOC_SIZEBITS   = 14
    _IOC_DIRBITS    = 2
    _IOC_NRMASK     = (1 << _IOC_NRBITS) - 1
    _IOC_TYPEMASK   = (1 << _IOC_TYPEBITS) - 1
    _IOC_SIZEMASK   = (1 << _IOC_SIZEBITS) - 1
    _IOC_DIRMASK    = (1 << _IOC_DIRBITS) - 1
    _IOC_NRSHIFT    = 0
    _IOC_TYPESHIFT  = _IOC_NRSHIFT + _IOC_NRBITS
    _IOC_SIZESHIFT  = _IOC_TYPESHIFT + _IOC_TYPEBITS
    _IOC_DIRSHIFT   = _IOC_SIZESHIFT + _IOC_SIZEBITS
    _IOC_NONE       = 0
    _IOC_WRITE      = 1
    _IOC_READ       = 2
)

type CmdType int64

func _IOC(d int, t int, n int, s int) CmdType {
    return CmdType(d << _IOC_DIRSHIFT | t << _IOC_TYPESHIFT | n << _IOC_NRSHIFT | s << _IOC_SIZESHIFT)
}

func _IO(t int, n int) CmdType {
    return _IOC(_IOC_NONE, t, n, 0)
}

func _IOR(t int, n int, s int) CmdType {
    return _IOC(_IOC_READ, t, n, s)
}

func _IOW(t int, n int, s int) CmdType {
    return _IOC(_IOC_WRITE, t, n, s)
}

var (
    BR_ERROR                            CmdType = _IOR('r', 0, 0x4)
    BR_OK                               CmdType = _IO('r', 1)
    BR_TRANSACTION                      CmdType = _IOR('r', 2, 0x40)
    BR_REPLY                            CmdType = _IOR('r', 3, 0x40)
    BR_ACQUIRE_RESULT                   CmdType = _IOR('r', 4, 0x4)
    BR_DEAD_REPLY                       CmdType = _IO('r', 5)
    BR_TRANSACTION_COMPLETE             CmdType = _IO('r', 6)
    BR_INCREFS                          CmdType = _IOR('r', 7, 0x10)
    BR_ACQUIRE                          CmdType = _IOR('r', 8, 0x10)
    BR_RELEASE                          CmdType = _IOR('r', 9, 0x10)
    BR_DECREFS                          CmdType = _IOR('r', 10, 0x10)
    BR_INCREFS_32                       CmdType = _IOR('r', 7, 0x8)
    BR_ACQUIRE_32                       CmdType = _IOR('r', 8, 0x8)
    BR_RELEASE_32                       CmdType = _IOR('r', 9, 0x8)
    BR_DECREFS_32                       CmdType = _IOR('r', 10, 0x8)
    BR_ATTEMPT_ACQUIRE                  CmdType = _IOR('r', 11, 0x18)
    BR_ATTEMPT_ACQUIRE_32               CmdType = _IOR('r', 11, 0xc)
    BR_NOOP                             CmdType = _IO('r', 12)
    BR_SPAWN_LOOPER                     CmdType = _IO('r', 13)
    BR_FINISHED                         CmdType = _IO('r', 14)
    BR_DEAD_BINDER                      CmdType = _IOR('r', 15, 0x8)
    BR_DEAD_BINDER_32                   CmdType = _IOR('r', 15, 0x4)
    BR_CLEAR_DEATH_NOTIFICATION_DONE    CmdType = _IOR('r', 16, 0x8)
    BR_CLEAR_DEATH_NOTIFICATION_DONE_32 CmdType = _IOR('r', 16, 0x4)
    BR_FAILED_REPLY                     CmdType = _IO('r', 17)
)

var (
    BC_TRANSACTION                  CmdType = _IOW('c', 0, 0x40)
    BC_REPLY                        CmdType = _IOW('c', 1, 0x40)
    BC_ACQUIRE_RESULT               CmdType = _IOW('c', 2, 0x4)
    BC_FREE_BUFFER                  CmdType = _IOW('c', 3, 0x8)
    BC_FREE_BUFFER_32               CmdType = _IOW('c', 3, 0x4)
    BC_INCREFS                      CmdType = _IOW('c', 4, 0x4)
    BC_ACQUIRE                      CmdType = _IOW('c', 5, 0x4)
    BC_RELEASE                      CmdType = _IOW('c', 6, 0x4)
    BC_DECREFS                      CmdType = _IOW('c', 7, 0x4)
    BC_INCREFS_DONE                 CmdType = _IOW('c', 8, 0x10)
    BC_ACQUIRE_DONE                 CmdType = _IOW('c', 9, 0x10)
    BC_INCREFS_DONE_32              CmdType = _IOW('c', 8, 0x8)
    BC_ACQUIRE_DONE_32              CmdType = _IOW('c', 9, 0x8)
    BC_ATTEMPT_ACQUIRE              CmdType = _IOW('c', 10, 0x8)
    BC_REGISTER_LOOPER              CmdType = _IO('c', 11)
    BC_ENTER_LOOPER                 CmdType = _IO('c', 12)
    BC_EXIT_LOOPER                  CmdType = _IO('c', 13)
    BC_REQUEST_DEATH_NOTIFICATION   CmdType = _IOW('c', 14, 0x10)
    BC_CLEAR_DEATH_NOTIFICATION     CmdType = _IOW('c', 15, 0x10)
    BC_DEAD_BINDER_DONE             CmdType = _IOW('c', 16, 0x8)
    BC_REQUEST_DEATH_NOTIFICATION_32 CmdType = _IOW('c', 14, 0x8)
    BC_CLEAR_DEATH_NOTIFICATION_32  CmdType = _IOW('c', 15, 0x8)
    BC_DEAD_BINDER_DONE_32          CmdType = _IOW('c', 16, 0x4)
    BC_TRANSACTION_SG               CmdType = _IOW('c', 17, 0x48)
    BC_REPLY_SG                     CmdType = _IOW('c', 18, 0x48)
)

func (c CmdType) String() string {
    switch (c) {
    case BR_ERROR: return "BR_ERROR"
    case BR_OK: return "BR_OK"
    case BR_TRANSACTION: return "BR_TRANSACTION"
    case BR_REPLY: return "BR_REPLY"
    case BR_ACQUIRE_RESULT: return "BR_ACQUIRE_RESULT"
    case BR_DEAD_REPLY: return "BR_DEAD_REPLY"
    case BR_TRANSACTION_COMPLETE: return "BR_TRANSACTION_COMPLETE"
    case BR_INCREFS: return "BR_INCREFS"
    case BR_ACQUIRE: return "BR_ACQUIRE"
    case BR_RELEASE: return "BR_RELEASE"
    case BR_DECREFS: return "BR_DECREFS"
    case BR_INCREFS_32: return "BR_INCREFS"
    case BR_ACQUIRE_32: return "BR_ACQUIRE"
    case BR_RELEASE_32: return "BR_RELEASE"
    case BR_DECREFS_32: return "BR_DECREFS"
    case BR_ATTEMPT_ACQUIRE_32: return "BR_ATTEMPT_ACQUIRE"
    case BR_DEAD_BINDER_32: return "BR_DEAD_BINDER_32"
    case BR_ATTEMPT_ACQUIRE: return "BR_ATTEMPT_ACQUIRE"
    case BR_NOOP: return "BR_NOOP"
    case BR_SPAWN_LOOPER: return "BR_SPAWN_LOOPER"
    case BR_FINISHED: return "BR_FINISHED"
    case BR_DEAD_BINDER: return "BR_DEAD_BINDER"
    case BR_CLEAR_DEATH_NOTIFICATION_DONE: return "BR_CLEAR_DEATH_NOTIFICATION_DONE"
    case BR_FAILED_REPLY: return "BR_FAILED_REPLY"
    case BC_TRANSACTION: return "BC_TRANSACTION"
    case BC_REPLY: return "BC_REPLY"
    case BC_ACQUIRE_RESULT: return "BC_ACQUIRE_RESULT"
    case BC_FREE_BUFFER: return "BC_FREE_BUFFER"
    case BC_FREE_BUFFER_32: return "BC_FREE_BUFFER"
    case BC_INCREFS: return "BC_INCREFS"
    case BC_ACQUIRE: return "BC_ACQUIRE"
    case BC_RELEASE: return "BC_RELEASE"
    case BC_DECREFS: return "BC_DECREFS"
    case BC_INCREFS_DONE: return "BC_INCREFS_DONE"
    case BC_ACQUIRE_DONE: return "BC_ACQUIRE_DONE"
    case BC_INCREFS_DONE_32: return "BC_INCREFS_DONE"
    case BC_ACQUIRE_DONE_32: return "BC_ACQUIRE_DONE"
    case BC_ATTEMPT_ACQUIRE: return "BC_ATTEMPT_ACQUIRE"
    case BC_REGISTER_LOOPER: return "BC_REGISTER_LOOPER"
    case BC_ENTER_LOOPER: return "BC_ENTER_LOOPER"
    case BC_EXIT_LOOPER: return "BC_EXIT_LOOPER"
    case BC_REQUEST_DEATH_NOTIFICATION: return "BC_REQUEST_DEATH_NOTIFICATION"
    case BC_CLEAR_DEATH_NOTIFICATION: return "BC_CLEAR_DEATH_NOTIFICATION"
    case BC_DEAD_BINDER_DONE: return "BC_DEAD_BINDER_DONE"
    case BC_REQUEST_DEATH_NOTIFICATION_32: return "BC_REQUEST_DEATH_NOTIFICATION"
    case BC_CLEAR_DEATH_NOTIFICATION_32: return "BC_CLEAR_DEATH_NOTIFICATION"
    case BC_DEAD_BINDER_DONE_32: return "BC_DEAD_BINDER_DONE"
    case BC_TRANSACTION_SG: return "BC_TRANSACTION_SG"
    case BC_REPLY_SG: return "BC_REPLY_SG"
    default: return "UNKNOWN"
    }
}

func get_sub(num int64, shift uint, mask int) int {
	return (int(num) >> shift) & mask
}

func show_banner() {
	fmt.Println("ION ioctl format:")
	fmt.Println("          |31    30|29    16|15     8|7      0")
	fmt.Println("----------|--------|--------|--------|--------")
	fmt.Println("Origin    |Dir     |Size    |Type    |Number")
	fmt.Println("----------|--------|--------|--------|--------")
}

func show_format(num int64) {
	n := get_sub(num, 0, 0xFF)
	t := get_sub(num, 8, 0xFF)
	s := get_sub(num, 16, 0x3FFF)
	d := get_sub(num, 30, 0x3)
	fmt.Printf("0x%-8X|%-8d|%-8d|%-8d|%-8d\n", num, d, s, t, n)
	fmt.Printf("Cmd: %v\n", CmdType(num))
}

func main() {
	show_banner()
	for i := 1; i < len(os.Args); i++ {
		str := strings.TrimPrefix(os.Args[i], "0x")
		num, _ := strconv.ParseInt(str, 16, 64)
		show_format(num)
	}
}
