package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"

	"github.com/kelseyhightower/journal-2-logentries/logentries"
)

func main() {
	url := os.Getenv("LOGENTRIES_URL")
	if url == "" {
		url = logentries.DefaultUrl
	}
	token := os.Getenv("LOGENTRIES_TOKEN")
	if token == "" {
		log.Fatal("non-empty input token (LOGENTRIES_TOKEN) is required. See https://logentries.com/doc/input-token")
	}
	logs, err := followJournal()
	if err != nil {
		log.Fatal(err.Error())
	}
	le, err := logentries.New(url, token)
	if err != nil {
		log.Fatal(err.Error())
	}
	for {
		select {
		case data := <-logs:
			if _, err := le.Write(data); err != nil {
				log.Print(err.Error())
			}
		}
	}
}

func followJournal() (<-chan []byte, error) {
	cmd := exec.Command("/usr/bin/journalctl", "--follow")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	logs := make(chan []byte)
	scanner := bufio.NewScanner(stdout)
	go func() {
		for scanner.Scan() {
			logs <- scanner.Bytes()
		}
		if err := scanner.Err(); err != nil {
			log.Println(err.Error())
			close(logs)
		}
	}()
	return logs, nil
}
