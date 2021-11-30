package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGABRT, syscall.SIGALRM, syscall.SIGBUS, syscall.SIGCHLD, syscall.SIGCLD,
		syscall.SIGCONT, syscall.SIGFPE, syscall.SIGHUP, syscall.SIGILL, syscall.SIGINT, syscall.SIGIO,
		syscall.SIGIOT, syscall.SIGKILL, syscall.SIGPIPE, syscall.SIGPOLL, syscall.SIGPROF, syscall.SIGPWR,
		syscall.SIGQUIT, syscall.SIGSEGV, syscall.SIGSTKFLT, syscall.SIGSTOP, syscall.SIGSYS, syscall.SIGTERM,
		syscall.SIGTRAP, syscall.SIGTSTP, syscall.SIGTTIN, syscall.SIGTTOU, syscall.SIGUNUSED, syscall.SIGUSR1,
		syscall.SIGUSR2, syscall.SIGVTALRM, syscall.SIGWINCH, syscall.SIGILL, syscall.SIGINT, syscall.SIGIO,
		syscall.SIGCONT, syscall.SIGFPE, syscall.SIGXCPU, syscall.SIGXFSZ,
	)

	fmt.Printf("Testing 'kill' command - LPIC1 - Chapter 02\n\tM. Javad Mohebi\n\nRead 'man kill' and send signal to this app. You might need to use 'screen' or 'tmux' to split the screen to different windows or panes!\n\n\n--------------------\n\n")
	fmt.Printf("PID: %v, PPID: %v\n", os.Getpid(), os.Getppid())
	fmt.Printf("\n--------------------\n")

	go func() {
		intr := 0
		for {
			sig := <-sigs
			fmt.Println()
			fmt.Printf("Signal: [ %d ] %v\n", sig, sig)
			if sig == os.Interrupt {
				intr += 1
				if intr == 1 {
					fmt.Printf("\t '%v' recevied %v time, press CTRL+C again to exit!\n", sig, intr)
				}
				if intr == 2 {
					fmt.Printf("\t %v recevied %v time, we are going to exit!\n", sig, intr)
					done <- true
				}
			} else {
				intr = 0
			}
		}

	}()

	fmt.Println("Wating for signals")
	<-done
	fmt.Println("exiting")
}
