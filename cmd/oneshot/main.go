package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	cmd := exec.CommandContext(ctx, "agy", "-p", "implement Slugify in slug.go so that all test in slug_test.go pass. Only edit slug.go", "--mode", "accept-edits", "--model", "gemini-3.8-flash-medium")

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	cmd.Dir = "/home/alok/building/fixture/"

	err := cmd.Run()

	exitCode := 0

	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		exitCode = exitErr.ExitCode()
	}

	werr := os.WriteFile("out.txt", stdout.Bytes(), 0o644)
	if werr != nil {
		fmt.Println("error while writing stdout ", werr)
	}

	fmt.Println("stdout", stdout.String())
	fmt.Println("stderr", stderr.String())
	fmt.Println("exicode", exitCode)
	fmt.Println("error", err)

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("TIMEOUT")
	}
}
