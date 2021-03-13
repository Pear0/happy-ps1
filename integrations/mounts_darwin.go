package integrations

import (
	"syscall"
)

// Constants from https://github.com/apple/darwin-xnu/blob/main/bsd/sys/mount.h

// Flags for getfsstat syscall
//goland:noinspection all
const (
	MNT_WAIT   int = 1 /* synchronized I/O file integrity completion */
	MNT_NOWAIT int = 2 /* start all I/O, but do not wait for it */
	MNT_DWAIT  int = 4 /* synchronized I/O data integrity completion */
)

// Flags for Statfs_t.Flags
//goland:noinspection all
const (
	/*
	 * User specifiable flags.
	 *
	 * Unmount uses MNT_FORCE flag.
	 */
	MNT_RDONLY      uint32 = 0x00000001 /* read only filesystem */
	MNT_SYNCHRONOUS uint32 = 0x00000002 /* file system written synchronously */
	MNT_NOEXEC      uint32 = 0x00000004 /* can't exec from filesystem */
	MNT_NOSUID      uint32 = 0x00000008 /* don't honor setuid bits on fs */
	MNT_NODEV       uint32 = 0x00000010 /* don't interpret special files */
	MNT_UNION       uint32 = 0x00000020 /* union with underlying filesystem */
	MNT_ASYNC       uint32 = 0x00000040 /* file system written asynchronously */
	MNT_CPROTECT    uint32 = 0x00000080 /* file system supports content protection */

	/*
	 * NFS export related mount flags.
	 */
	MNT_EXPORTED uint32 = 0x00000100 /* file system is exported */

	/*
	 * Denotes storage which can be removed from the system by the user.
	 */

	MNT_REMOVABLE uint32 = 0x00000200

	/*
	 * MAC labeled / "quarantined" flag
	 */
	MNT_QUARANTINE uint32 = 0x00000400 /* file system is quarantined */

	/*
	 * Flags set by internal operations.
	 */
	MNT_LOCAL   uint32 = 0x00001000 /* filesystem is stored locally */
	MNT_QUOTA   uint32 = 0x00002000 /* quotas are enabled on filesystem */
	MNT_ROOTFS  uint32 = 0x00004000 /* identifies the root filesystem */
	MNT_DOVOLFS uint32 = 0x00008000 /* FS supports volfs (deprecated flag in Mac OS X 10.5) */

	MNT_DONTBROWSE       uint32 = 0x00100000 /* file system is not appropriate path to user data */
	MNT_IGNORE_OWNERSHIP uint32 = 0x00200000 /* VFS will ignore ownership information on filesystem objects */
	MNT_AUTOMOUNTED      uint32 = 0x00400000 /* filesystem was mounted by automounter */
	MNT_JOURNALED        uint32 = 0x00800000 /* filesystem is journaled */
	MNT_NOUSERXATTR      uint32 = 0x01000000 /* Don't allow user extended attributes */
	MNT_DEFWRITE         uint32 = 0x02000000 /* filesystem should defer writes */
	MNT_MULTILABEL       uint32 = 0x04000000 /* MAC support for individual labels */
	MNT_NOATIME          uint32 = 0x10000000 /* disable update of file access time */
	MNT_SNAPSHOT         uint32 = 0x40000000 /* The mount is a snapshot */
	MNT_STRICTATIME      uint32 = 0x80000000 /* enable strict update of file access time */
)

func toMountInfo(fs syscall.Statfs_t) MountInfo {
	return MountInfo{
		MountLocation:  arrayToString(fs.Mntonname[:]),
		FilesystemType: arrayToString(fs.Fstypename[:]),
		Source:         arrayToString(fs.Mntfromname[:]),
		IsRemote:       !checkFlag(fs.Flags, MNT_LOCAL),
	}
}

func getMounts() ([]syscall.Statfs_t, error) {
	count, err := syscall.Getfsstat(nil, MNT_NOWAIT)
	if err != nil {
		return nil, err
	}

	buf := make([]syscall.Statfs_t, count)

	count, err = syscall.Getfsstat(buf, MNT_NOWAIT)
	if err != nil {
		return nil, err
	}

	return buf[:count], nil
}

func GetMountForPath(path string) (MountInfo, error) {
	var fs syscall.Statfs_t

	err := syscall.Statfs(path, &fs)
	if err != nil {
		return MountInfo{}, err
	}

	return toMountInfo(fs), nil
}
