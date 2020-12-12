package main

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"
)

type Entry struct {
	key   string
	value int
}

func predict(county, json string) chan Entry {
	var c = make(chan Entry)
	//println("PREDICTING: " + county)

	go func() {
		//fmt.Println("RUNNING " + county + " :: " + json)
		command := exec.Command("python", "src/test.py")

		buffer := bytes.Buffer{}
		buffer.Write([]byte(json))
		command.Stdin = &buffer

		out, err := command.Output()

		if err != nil {
			println(err.Error())
			return
		}

		result, err := strconv.Atoi(strings.TrimSpace(string(out)))
		if err != nil {
			println(err)
		}
		//fmt.Println(string(out)+" :: ", result)
		e := Entry{county, result}
		c <- e
	}()

	return c
}

func main2() {
	//counties := readInput([]County{})
	//fmt.Println(counties[0])

	//a := newCounty("A")
	//b := newCounty("B")
	//c := newCounty("B")
	//
	//data := []County{a, b, c}

	//m := map[string]int{
	//	"something": 10,
	//	"yo":        20,
	//	"blah":      20,
	//}
	//
	//var sortedCounties []Entry
	//for k, v := range m {
	//	sortedCounties = append(sortedCounties, Entry{k, v})
	//}
	//
	//sort.Slice(sortedCounties, func(i, j int) bool {
	//	return sortedCounties[i].Value > sortedCounties[j].Value
	//})
	//
	//fmt.Println(sortedCounties)

	//command := exec.Command("python", "src/test.py")
	//
	//buffer := bytes.Buffer{}
	//buffer.Write([]byte("[2,[1,2,4,8]]"))
	//command.Stdin = &buffer
	//
	//out, err := command.Output()
	//
	//if err != nil {
	//	println(err.Error())
	//	return
	//}
	//
	//fmt.Println(string(out))
}
