// MIT License
// 
// Copyright (c) 2017 Nobi Hoang <nobi.hoa@gmail.com>
// 

package cmdutil

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadLinePmt prompts the given message and reads a line from the command line.
// It returns the line and an error if that fails.
func ReadLinePmt(prompt string) (string, error) {
	fmt.Printf("%s : ", prompt)
	line, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(line)), nil
}

// ReadLine reads a line from the command line.
// It returns the line and an error if that fails.
func ReadLine() (string, error) {
	line, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(line)), nil
}