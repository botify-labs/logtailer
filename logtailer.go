package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

//TODO: handle signals, especially SIGINT!
func main() {
	//args parsing
	var servers []string
	var files []string
	for _, elem := range os.Args[1:] {
		if strings.HasPrefix(elem, "/") {
			files = append(files, elem)
		} else {
			servers = append(servers, elem)
		}
	}

	//fail if no server or no file
	if len(files) == 0 || len(servers) == 0 {
		fmt.Fprintln(os.Stderr, "Please specify at least ONE server and ONE file beginning with '/'")
		os.Exit(1)
	}

	//this will help us wait instead of exiting directly
	var wg sync.WaitGroup

	//for each server, let's tail the logs (async!)
	for _, server := range servers {
		wg.Add(1)
		go func(server string) {
			tailServerLogs(server, files)
			wg.Done()
		}(server)
	}

	//wait for all goroutines to finish
	wg.Wait()
}

//tail logs on a remote server
func tailServerLogs(server string, files []string) {
	//for later use in output
	coloredServer := ColorHostname(server)

	//build command
	cmdName := "ssh"
	tailCmd := "sudo tail -n 0 -F " + strings.Join(files, " ")
	cmdArgs := []string{server, tailCmd}

	//prepare command
	cmd := exec.Command(cmdName, cmdArgs...)

	//handle stdout
	cmdStdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating stdout pipe for command: ", err)
		return
	}
	scannerStdout := bufio.NewScanner(cmdStdout)
	go func() {
		for scannerStdout.Scan() {
			now := time.Now().Format(time.RFC3339)
			fmt.Printf(
				"%s %s %s %s\n",
				coloredServer,
				now,
				ColorStream("out", "stdout"),
				scannerStdout.Text(),
			)
		}
	}()

	//handle stderr
	cmdStderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating stderr pipe for command: ", err)
		return
	}
	scannerStderr := bufio.NewScanner(cmdStderr)
	go func() {
		for scannerStderr.Scan() {
			now := time.Now().Format(time.RFC3339)
			fmt.Printf(
				"%s %s %s %s\n",
				coloredServer,
				now,
				ColorStream("ERR", "stderr"),
				scannerStderr.Text(),
			)
		}
	}()

	//launch command
	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting command: ", err)
		return
	}

	//wait for it to end
	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for command: ", err)
		return
	}
}

// Returns a colorized version of the hostname
// The color shouldn't vary for a given hostname so we calculate
// a numeric hash for the hostname and reduce it to a list of colors
func ColorHostname(hostname string) string {
	if !strings.HasPrefix(os.Getenv("TERM"), "xterm") {
		return hostname
	}
	validColors := []string{
		"33", "34", "35", "36", "37",
	}
	colorIndex := HashHostnameToInt(hostname) % len(validColors)
	color := validColors[colorIndex]
	colorReset := "\x1b[0m"
	colorHost := "\x1b[" + color + "m"
	return colorHost + hostname + colorReset
}

// Generates a numeric hash out of a string
func HashHostnameToInt(str string) int {
	h := fnv.New32a()
	h.Write([]byte(str))
	return int(h.Sum32())
}

// Helper to colorize stdout/stdin markers (red=err, green=out)
func ColorStream(message string, stream string) string {
	colorOkay := ""
	colorFail := ""
	colorReset := ""
	if strings.HasPrefix(os.Getenv("TERM"), "xterm") {
		colorOkay = "\x1b[32m"
		colorFail = "\x1b[31m"
		colorReset = "\x1b[0m"
	}
	if stream == "stdout" {
		return colorOkay + message + colorReset
	} else if stream == "stderr" {
		return colorFail + message + colorReset
	} else {
		//turn it into an error?
		return message
	}
}
