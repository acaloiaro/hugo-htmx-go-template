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

	_, err := exec.LookPath("hugo")
	if err != nil {
		log.Fatal("can't find 'hugo' in your $PATH. you probably need to install hugo: `go install -tags extended github.com/gohugoio/hugo@latest`")
	}

	// build hugo site static content
	cmd := exec.Command("hugo")

	outPipe, _ := cmd.StdoutPipe()
	errPipe, _ := cmd.StderrPipe()

	err = cmd.Start()
	if err != nil {
		log.Printf("can't start hugo server: %v", err)
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
		log.Fatalf("unable to build hugo site: %q", err)
	}

	cmd = exec.Command("go", "build", "-o", "build/server", "server.go")

	outPipe, _ = cmd.StdoutPipe()
	errPipe, _ = cmd.StderrPipe()

	err = cmd.Start()
	if err != nil {
		log.Printf("can't build API server: %v", err)
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
		log.Fatalf("ERROR: unable to build server executable: %q", err)
	}

	log.Println("build complete")
}
