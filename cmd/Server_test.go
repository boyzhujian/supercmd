package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/mholt/archiver"
	"github.com/parnurzeal/gorequest"
)

const hostname = "https://127.0.0.1:8443"

func TestServer(t *testing.T) {
	servername, _ := os.Hostname()
	log.Println(servername)

	file := "/server/fileexist?filename=/Users/zhujian/goinstall/src/github.com/boyzhujian/supercmd/cmd/main.go"
	query := hostname + file
	fmt.Println(query)
	_, body, _ := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true}).Get(query).End()
	if !strings.Contains(body, "okokok") {
		t.Errorf("file not exist" + file)
	}
	fmt.Println(body)

}

func TestArchinve(t *testing.T) {
	err := archiver.Archive([]string{"/Users/zhujian/goinstall/src/github.com/boyzhujian/supercmd", "main.go"}, "test.tar.gz")
	fmt.Println(err)

}
