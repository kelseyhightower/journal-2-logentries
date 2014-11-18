package journal

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
)

const DefaultSocket = "/run/journald.sock"

func Follow(socket string) (<-chan []byte, error) {
	if socket == "" {
		socket = DefaultSocket
	}
	c := &http.Client{
		Transport: &http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				return net.Dial("unix", socket)
			},
		},
	}
	req, err := http.NewRequest("GET", "http://journal/entries?follow", nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non 200 response: %d", resp.StatusCode)
	}
	logs := make(chan []byte)
	scanner := bufio.NewScanner(resp.Body)
	go func() {
		for scanner.Scan() {
			logs <- scanner.Bytes()
		}
		if err := scanner.Err(); err != nil {
			log.Println(err.Error())
		}
	}()
	return logs, nil
}
