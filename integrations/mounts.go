package integrations

import (
	"fmt"
	"os"
	"syscall"
)

type MountInfo struct {
	MountLocation  string
	FilesystemType string
	Source         string
	IsRemote       bool
}

func arrayToString(buf []int8) string {
	var newBuf []byte

	for _, c := range buf {
		if c == 0 {
			break
		}
		newBuf = append(newBuf, byte(c))
	}

	return string(newBuf)
}

func checkFlag(val uint32, mask uint32) bool {
	return val&mask == mask
}

func GetMountInfo() {

	count, err := syscall.Getfsstat(nil, MNT_NOWAIT)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("fs count: %d\n", count)

	buf := make([]syscall.Statfs_t, count)

	count, err = syscall.Getfsstat(buf, MNT_NOWAIT)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}

	/*
		Bsize       uint32
		Iosize      int32
		Blocks      uint64
		Bfree       uint64
		Bavail      uint64
		Files       uint64
		Ffree       uint64
		Fsid        Fsid
		Owner       uint32
		Type        uint32
		Flags       uint32
		Fssubtype   uint32
		Fstypename  [16]int8
		Mntonname   [1024]int8
		Mntfromname [1024]int8
		Reserved    [8]uint32
	*/

	for i := 0; i < count; i++ {
		b := buf[i]
		fmt.Printf("Bsize: %d\n", b.Bsize)
		fmt.Printf("Iosize: %d\n", b.Iosize)
		fmt.Printf("Blocks: %d\n", b.Blocks)
		fmt.Printf("Bfree: %d\n", b.Bfree)
		fmt.Printf("Bavail: %d\n", b.Bavail)
		fmt.Printf("Files: %d\n", b.Files)
		fmt.Printf("Ffree: %d\n", b.Ffree)
		fmt.Printf("Owner: %d\n", b.Owner)
		fmt.Printf("Type: %d\n", b.Type)
		fmt.Printf("Flags: %d\n", b.Flags)
		fmt.Printf("Fssubtype: %d\n", b.Fssubtype)

		fmt.Printf("Fstypename: %s\n", arrayToString(b.Fstypename[:]))
		fmt.Printf("Mntonname: %s\n", arrayToString(b.Mntonname[:]))
		fmt.Printf("Mntfromname: %s\n", arrayToString(b.Mntfromname[:]))

		fmt.Printf("flag readonly: %v\n", checkFlag(b.Flags, MNT_RDONLY))
		fmt.Printf("flag removable: %v\n", checkFlag(b.Flags, MNT_REMOVABLE))
		fmt.Printf("flag is_local: %v\n", checkFlag(b.Flags, MNT_LOCAL))

		fmt.Println()
	}

}
