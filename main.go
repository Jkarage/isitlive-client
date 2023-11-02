/*
This is a client that would be run continuously every few minutes. This and the server should be run separately.
This tries to connect to a server and if fails sends the result to the mail & sms server

- Author Jkarage
*/
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	if err := Dial(); err != nil {
		log.Fatal(err)
	}
}

func Dial() error {
	// Adding endpoints of the machines using the terminal command line arguments
	// targets := []string{"", ""}
	targets := os.Args[1:]

	port := 80

	for _, v := range targets {
		_, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", v, port), 50*time.Second)
		if err != nil {
			if err = NewRequest(v); err != nil {
				return err
			}
		}

	}

	return nil
}

func NewRequest(v string) error {

	// Create a new Get request
	req, err := http.NewRequest("GET", fmt.Sprintf("http://127.0.0.1:3001/%s", v), nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if err != nil {
		return err
	}
	return nil
}
