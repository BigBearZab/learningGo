package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Try to use a reciever
func (c cards) toString() string {
	fmt.Println(c)
	s := strings.Join(c, ",")
	return s
}

func SaveFile(b []byte, f string) error {
	err := ioutil.WriteFile(f, b, 0666)
	if err != nil {
		return err
	}
	return nil
}

func strFromFile(f string) []byte {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}
	return b
}

func BytesToStringList(b []byte) []string {
	s := string(b)
	ss := strings.Split(s, ",")
	return ss
}
