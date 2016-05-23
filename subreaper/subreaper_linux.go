package subreaper

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/opencontainers/runc/libcontainer/system"
)

func Subreaper(pids chan int) {
	system.SetSubreaper(os.Getpid())

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGCHLD)

	var status syscall.WaitStatus
	for range signals {
		for {
			// -1, the call waits for any child process.
			wpid, err := syscall.Wait4(-1, &status, syscall.WNOHANG, nil)
			if err != nil || wpid <= 0 {
				break
			}

			pids <- wpid
		}
	}

}
