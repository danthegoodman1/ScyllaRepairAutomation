package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func repair(execScript, host, logDir string) (int, string) {
	repairCommand := "nodetool repair -pr"

	cmd := exec.Command("bash", execScript, host, repairCommand)
	logFilePath := ""
	if logDir != "" {
		logFilePath = filepath.Join(logDir, fmt.Sprintf("%s_repair.log", host))

		logFile, err := os.Create(logFilePath)
		if err != nil {
			log.Fatal(err)
		}
		defer logFile.Close()
		cmd.Stdout = logFile
		cmd.Stderr = logFile
	} else {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	err := cmd.Run()

	if err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if ok {
			return exitErr.ExitCode(), logFilePath
		}
		log.Fatal(err)
	}

	return 0, logFilePath
}

func main() {
	execScript := flag.String("exec", "", "Path to the bash script to execute")
	hosts := flag.String("hosts", "", "CSV file containing a list of hosts")
	logOutputDir := flag.String("logdir", "", "Log output directory, if does not exist then logs will come out to console. Useful for making one log file per host, otherwise use normal redirects.")
	failureScript := flag.String("f", "", "Script to execute on failure, first parameter passed is the host, second is the log output file (will be blank if logging to console)")
	successScript := flag.String("s", "", "Script to execute on success")

	flag.Parse()

	if _, err := os.Stat(*logOutputDir); os.IsNotExist(err) {
		os.MkdirAll(*logOutputDir, os.ModePerm)
	}

	startTime := time.Now()

	for _, host := range strings.Split(*hosts, ",") {
		resultCode, logPath := repair(*execScript, host, *logOutputDir)
		if resultCode != 0 {
			fmt.Printf("Failed on host %s, check log output\n", host)
			if *failureScript != "" {
				absPath, err := filepath.Abs(logPath)
				if err != nil {
					panic(err)
				}
				exec.Command("bash", *failureScript, host, absPath).Run()
			}
			os.Exit(1)
		}
	}

	if *successScript != "" {
		exec.Command("bash", *successScript, time.Since(startTime).String()).Run()
	}
	log.Println("Success!")
}
