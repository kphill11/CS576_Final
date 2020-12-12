package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	command := exec.Command("python", "src/test.py")

	buffer := bytes.Buffer{}
	buffer.Write([]byte("[2,[1,2,4,8]]"))
	command.Stdin = &buffer

	out, err := command.Output()

	if err != nil {
		println(err.Error())
		return
	}

	fmt.Println(string(out))
}
