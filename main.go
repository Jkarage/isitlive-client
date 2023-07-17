/*
Package client
This is a client that would be run continuously every few minutes. This should be run separately.
This tries to connect to a server and if fails sends the result to the mail & sms server

- Author Jkarage
*/
package main

import (
	"fmt"
	"io"
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
	fmt.Println(targets)

	ports := []int{22, 80}

	for _, v := range targets {
		_, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", v, ports[0]), 5*time.Second)
		if err != nil {
			fmt.Println(err)
			if err = NewGetRequest(v); err != nil {
				fmt.Println("Got an error from request")
				return err
			}
		}

		_, err = net.DialTimeout("tcp", fmt.Sprintf("%s:%d", v, ports[1]), 5*time.Second)
		if err != nil {
			return NewGetRequest(v)
		}
	}

	return nil
}

func NewGetRequest(v string) error {
	fmt.Println("Started new get request")

	// Create a new Get request
	req, err := http.NewRequest("GET", fmt.Sprintf("http://127.0.0.1:3000/%s", v), nil)
	if err != nil {
		fmt.Println("Started here")
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("No It started over here")
		return err
	}

	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
