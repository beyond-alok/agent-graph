package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	ctx,cancel := context.WithTimeout(context.Background(),time.Minute)
	defer cancel()
	cmd := exec.CommandContext(ctx,"agy","-p","create a file named hello.txt containing the word hii", "--mode","plan")

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	cmd.Dir = "/home/alok/building/agent-graph"
	err := cmd.Run()

	exitcode := 0

	var exitErr *exec.ExitError
	if errors.As(err,&exitErr) {
		exitcode = exitErr.ExitCode()
	}

	fmt.Println("stdout",stdout.String())

	fmt.Println("stderr",stderr.String())

	fmt.Println("exit code",exitcode)

	fmt.Println("error",err)

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("TIME OUT")
	}
	
}