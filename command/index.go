package command

import (
	"bytes"
	"context"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Command struct {
	command string
	path    string

	cancel context.CancelFunc
}

func Run(path string, command string, toOS bool) (string, error) {
	cmd := New(path, command)
	return cmd.Run(toOS)
}

func New(path string, command string) Command {
	return Command{command: command, path: path}
}

func (c *Command) Cancel() {
	c.cancel()
}

func cretaeMultiWritter(toOS bool) (io.Writer, bytes.Buffer) {
	var buf bytes.Buffer
	if toOS {
		return io.MultiWriter(os.Stdout, &buf), buf
	}
	return io.MultiWriter(&buf), buf
}

func (c *Command) Run(toOS bool) (string, error) {
	parts := strings.Fields(c.command)
	cmdName := parts[0]
	cmdArgs := parts[1:]

	ctx, cancel := context.WithCancel(context.Background())
	c.cancel = cancel
	cmd := exec.CommandContext(ctx, cmdName, cmdArgs...)
	cmd.Dir = c.path

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Printf("StderrPipe not created because: %v\n", err)
		return "", err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Printf("StdoutPipe not created because: %v\n", err)
		return "", err
	}

	if err := cmd.Start(); err != nil {
		log.Printf("command not started because: %v\n", err)
		return "", err
	}

	mw, buf := cretaeMultiWritter(toOS)

	if _, err := io.Copy(mw, stdout); err != nil {
		log.Printf("stdout not copied because: %v\n", err)
		return "", err
	}

	if _, err := io.Copy(mw, stderr); err != nil {
		log.Printf("stderr not copied because: %v\n", err)
		return "", err
	}

	if err := cmd.Wait(); err != nil {
		log.Printf("error waiting for command to end: %v\n", err)
		return buf.String(), err
	}

	return buf.String(), nil
}
