package subreaper_test

import (
	"os/exec"
	"syscall"

	"github.com/albertoleal/gunix/subreaper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Subreaper", func() {

	It("detects when a child process is killed", func() {
		pids := make(chan int, 1)
		go func() {
			subreaper.Subreaper(pids)
		}()

		cmd := exec.Command("sleep", "10")
		// It is not needed to call Wait()
		err := cmd.Start()
		Expect(err).NotTo(HaveOccurred())

		err = syscall.Kill(cmd.Process.Pid, syscall.SIGTERM)
		Expect(err).NotTo(HaveOccurred())

		//It cannot be init process :P
		Expect(<-pids).To(BeNumerically(">", 1))
	})

})
