package ptrace

import "golang.org/x/sys/unix"

func waitpid(pid int) int {
	var s unix.WaitStatus
	wpid, _ := unix.Wait4(pid, &s, unix.WALL, nil)
	return wpid
}
