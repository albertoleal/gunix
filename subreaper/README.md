## Subreaper

In effect, a subreaper fulfills the role
of init(1) for its descendant processes.  Upon termination of
a process that is orphaned (i.e., its immediate parent has
already terminated) and marked as having a subreaper, the
nearest still living ancestor subreaper will receive a SIGCHLD
signal and be able to wait(2) on the process to discover its
termination status. 

More information: http://man7.org/linux/man-pages/man2/prctl.2.html


To see the zombie process you should comment out the following line:
```bash
go func() {
  // Subreaper(pids)
}()

```

```
# ps -ef | grep -C 2 sleep
root      3688  3652  0 20:48 pts/0    00:00:00 /tmp/ginkgo521694956/subreaper.test --test.timeout=24h --ginkgo.seed=1464036490 --ginkgo.slowSpecThreshold=5.00000
root      3693  3688  0 20:48 pts/0    00:00:00 [sleep] <defunct>
```
