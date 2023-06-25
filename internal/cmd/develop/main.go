package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	var haveAir = false
	_, err := exec.LookPath("hugo")
	if err != nil {
		log.Fatal("can't find 'hugo' in your $PATH. you probably need to install hugo: `go install -tags extended github.com/gohugoio/hugo@latest`")
	}

	_, err = exec.LookPath("air")
	if err == nil {
		haveAir = true
	}

	// start hugo server
	go func() {
		cmd := exec.Command("hugo", "server")

		outPipe, _ := cmd.StdoutPipe()
		errPipe, _ := cmd.StderrPipe()

		err = cmd.Start()
		if err != nil {
			log.Fatalf("can't start hugo server: %v", err)
		}

		go func() {
			// print the output from the command in real time
			scanner := bufio.NewScanner(outPipe)
			for scanner.Scan() {
				m := scanner.Text()
				fmt.Println(m)
			}
		}()

		go func() {
			// print err output from the command in real time
			scanner := bufio.NewScanner(errPipe)
			for scanner.Scan() {
				m := scanner.Text()
				fmt.Println(m)
			}
		}()

		err = cmd.Wait()
		if err != nil {
			log.Fatalf("ERROR: hugo server stopped: %q", err)
		}
	}()

	// start the API server
	go func() {
		var cmd *exec.Cmd
		if haveAir {
			cmd = exec.Command("air")
		} else {
			cmd = exec.Command("go", "run", "server.go")
		}
		outPipe, _ := cmd.StdoutPipe()
		errPipe, _ := cmd.StderrPipe()

		err = cmd.Start()
		if err != nil {
			log.Printf("can't start API server: %v", err)
		}

		go func() {
			// print the output from the command in real time
			scanner := bufio.NewScanner(outPipe)
			for scanner.Scan() {
				m := scanner.Text()
				fmt.Println(m)
			}
		}()

		go func() {
			// print err output from the command in real time
			scanner := bufio.NewScanner(errPipe)
			for scanner.Scan() {
				m := scanner.Text()
				fmt.Println(m)
			}
		}()

		err = cmd.Wait()
		if err != nil {
			log.Fatalf("ERROR: API server stopped: %q", err)
		}

		c <- syscall.SIGINT
	}()

	<-c
}
