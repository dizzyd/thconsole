package main

import "crypto/rand"
import "crypto/rsa"
import "fmt"
import "github.com/GeertJohan/go.linenoise"
import "github.com/dizzyd/gogotelehash"
import "log"
import "os"
import "path"

func main() {
	// Initialize data dir, if it doesn't already exist
	configDir := path.Join(os.ExpandEnv("$HOME"), ".thconsole")
	os.Mkdir(configDir, 0700)

	// Setup logging subsystem
	logFilename := path.Join(configDir, "log")
	logFile, err := os.OpenFile(logFilename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("Could not open %s: %s\n", logFilename, err)
		os.Exit(1)
	}
	defer logFile.Close()

	// Direct all logging output to the log file
	log.SetOutput(logFile)
	log.Println("Started thconsole!")

	key, _ := rsa.GenerateKey(rand.Reader, 1024)
	telehash.NewSwitch("0.0.0.0:0", key)

	// Load command line history
	historyFilename := path.Join(configDir, "history")
	linenoise.LoadHistory(historyFilename)
	defer linenoise.SaveHistory(historyFilename)

	defer log.Println("Shutting down...")

	// Start processing commands
	for {
		str, err := linenoise.Line("th> ")
		if err != nil {
			if err == linenoise.KillSignalError {
				return
			}
			fmt.Println("Unexpected error: %s", err)
			return
		}

		switch str {
		case "quit":
			return
		}

		linenoise.AddHistory(str)
		fmt.Printf("Got: %s\n", str)
	}
}
