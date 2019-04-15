package adbt

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

// Process defines a basic process
type Process struct {
	Command string
	Params  []string
}

// Run starts the process
func (p *Process) Run(identifier string) bool {
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd := exec.Command(p.Command, p.Params...)
	fmt.Println(cmd.Args)
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	var backlog BackLog
	backlog.Identifier = identifier
	backlog.Time = time.Now().String()
	err := cmd.Start()
	if err != nil {
		backlog.Success = false
		backlog.Log = err.Error()
		backlog.writeLog()
		log.Fatalf("ADBT failed with '%s'\n", err)

	}
	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()
	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()
	err = cmd.Wait()
	if err != nil {
		log.Fatalf("ADBT failed with %s\n", err)
		backlog.Success = false
		backlog.Log = err.Error()
		backlog.writeLog()
	}
	if errStdout != nil || errStderr != nil {
		log.Fatal("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
	backlog.Success = true
	backlog.Log = outStr
	backlog.writeLog()
	return true
}
