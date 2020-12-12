package main

import (
	"fmt"
	"net"
	"os/exec"
	"time"
)

type Entry struct {
	key   string
	value int
}

type FloatEntry struct {
	key   string
	value float32
}

func predict(county, json string) Entry {
	subprocess := exec.Command("python", "src/test.py")

	subprocess.Start()

	stdin, err := subprocess.StdinPipe()
	if err != nil {
		println("ERROR:", err)
	}
	fmt.Println(json)
	subprocess.Start()
	stdin.Write([]byte(json))
	out, _ := subprocess.Output()
	fmt.Println("OUT:", string(out))
	//stdin, err := subprocess.StdinPipe()
	//if err != nil {
	//	println("ERROR:", err)
	//}
	//defer stdin.Close()
	//fmt.Println(json)

	//if err = subprocess.Start(); err != nil { //Use start, not run
	//	fmt.Println("An error occured: ", err) //replace with logger, or anything you want
	//}

	//buffer := bytes.Buffer{}
	//buffer.Write([]byte(json))
	//subprocess.Stdin = &buffer
	//
	////subprocess.Stdout = os.Stdout
	////subprocess.Stderr = os.Stderr
	//
	//out, _ := subprocess.Output()
	//fmt.Println("OUT:", string(out))
	//
	//buffer.Write([]byte("[2,[1,2,4,8]]"))
	//fmt.Println("OUT:", string(out))

	//io.WriteString(stdin, json)
	//stdin.Write([]byte(json))
	//if _, err = io.WriteString(stdin, json); err != nil {
	//	println(err)
	//}

	//subprocess.Wait()

	//buffer := bytes.Buffer{}
	//buffer.Write([]byte(json))
	//subprocess.Stdin = &buffer
	//
	//out, err := subprocess.Output()
	//
	//if err != nil {
	//	println(err.Error())
	//	return Entry{county, -1}
	//}
	//
	//result, err := strconv.Atoi(strings.TrimSpace(string(out)))
	//if err != nil {
	//	println(err)
	//}
	//return Entry{county, result}
	return Entry{county, -1}
}

func main2() {
	conn, err := net.Dial("tcp", "127.0.0.1:5000")
	if err != nil {
		println(err)
		return
	}
	defer conn.Close()
	if _, err = conn.Write([]byte("[12,[1,2,4,8]]")); err != nil {
		println(err)
	}

	buff := make([]byte, 128)
	if _, err = conn.Read(buff); err != nil {
		println(err)
	}
	fmt.Println(string(buff))

	time.Sleep(3000 * time.Millisecond)
	if _, err = conn.Write([]byte("[11,[1,2,4,8]]")); err != nil {
		println(err)
	}

	if _, err = conn.Read(buff); err != nil {
		println(err)
	}
	fmt.Println(string(buff))
	//fmt.Println(ioutil.ReadAll(conn))

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
