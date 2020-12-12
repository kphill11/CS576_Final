package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

type County struct {
	name     string
	numBeds  int
	timeline [95]int
}

func newCounty(newName string) County {
	county := County{}
	county.name = newName
	county.numBeds = 0
	for i := 0; i < len(county.timeline); i++ {
		county.timeline[i] = 0
	}
	return county
}

func findCounty(counties []County, name string) int {
	for i := 0; i < len(counties); i++ {
		if counties[i].name == name {
			return i
		}
	}
	//if you've searched through the entire array and haven't found it
	return -1
}

//func readInput(counties []County) []County {
func main() {
	//hospitals.csv:
	//field 13 is population
	//field 14 is county
	//field 32 is beds
	//field 33 is trauma (likely not going to use)
	var counties []County
	//counties = append(counties, newCounty("albany"))
	csvHospFile, errHosp := os.Open("..\\CS576_Final\\data\\hospitals.csv")
	if errHosp != nil {
		fmt.Printf("Error while loading hospitals.csv:\n")
		fmt.Printf("\n%s\n", errHosp)
		os.Exit(3)
	}
	reader := csv.NewReader(csvHospFile)
	line, err := reader.Read() //just getting the definitional line out of the way
	//fmt.Printf("%s, %s, %s\n", line[14], line[13], line[31])
	for {
		line, err = reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("\nerror while reading hospitals.csv:\n")
			fmt.Printf("\n%s\n", err)
			os.Exit(3)
		} else {
			name := strings.ToLower(strings.Trim(line[14], " "))
			if name != "not available" {
				index := findCounty(counties, name)
				if index < 0 { //if the county isn't already on record
					counties = append(counties, newCounty(name))
					index = len(counties) - 1
				}
				if line[31] != "-999" {
					//fmt.Printf("%s\n", line[31])
					numBeds, err := strconv.Atoi(line[31])
					if err != nil {
						fmt.Printf("%s\n", err)
						os.Exit(3)
					}
					if numBeds >= 0 {
						counties[index].numBeds += numBeds //number of regular beds
					}
				}
			}
		}
	}
	//covid-latest.csv
	//field 0 is date
	//field 1 is County
	//field 4 is cases
	csvCovidFile, errCov := os.Open("..\\CS576_Final\\data\\covid-latest.csv")
	if errCov != nil {
		fmt.Printf("\nError while loading covid-latest.csv:\n")
		fmt.Printf("%s\n", errCov)
		os.Exit(3)
	}
	reader = csv.NewReader(csvCovidFile)
	line, err = reader.Read()
	date := 0
	dateString := "2020-09-07"
	for line[0] != dateString {
		line, err = reader.Read()
	}
	for {
		line, err = reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error while reading covid-latest.csv:\n")
			fmt.Printf("%s\n", err)
			os.Exit(3)
		} else {
			name := strings.ToLower(strings.Trim(line[1], " "))
			index := findCounty(counties, name)
			if index < 0 { //if the county isn't already on record - bit more or a problem here
				counties = append(counties, newCounty(name))
				index = len(counties) - 1
				//counties[index].pop = strconv.Atoi(line[13]) //county population, not likely to change significantly enough to track
			}
			if dateString != line[0] {
				date++
				dateString = line[0]
			}
			caseNum, err := strconv.Atoi(line[4])
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(3)
			}
			counties[index].timeline[date] += caseNum
		}
	}
	//fmt.Printf("%d\n", len(counties))
	for i := len(counties) - 1; i > 0; i-- {
		if counties[i].numBeds == 0 {
			counties[i] = counties[len(counties)-1]
			//counties[len(counties) - 1] = null
			counties = counties[:len(counties)-1]
			i--
		}

	}
	//fmt.Printf("%d\n", len(counties))
	//fmt.Printf("%s: population: %d. number of beds: %d. Number of infected people on December 10th: %d\n", counties[i].name, counties[i].pop, counties[i].numBeds, counties[i].timeline[95          ])

	//county := counties[0]
	//predict(county.name, convertToJSON(county))

	//bufferSize := 10

	/*
		- IPC
		- Single threaded server: single connection, multiple connections, limited individual connections
	*/
	workerThreads := 4

	processed := make(chan string)
	queue := make(chan County)
	done := make(chan bool)
	for i := 0; i < workerThreads; i++ {
		go func() {
			for {
				select {
				case county := <-queue:
					result := run(county)
					go func() {
						processed <- result
					}()
				case <-done:
					break
				}
			}
		}()
	}

	//c := make(chan string)
	for _, county := range counties {
		queue <- county
		//county := county
		//go func() {
		//fmt.Println(i, string(buff))
		//}()

		//time.Sleep(100 * time.Millisecond)
	}
	for range counties {
		prediction := <-processed
		fmt.Println("FINAL GOT PRED: " + prediction)
	}

	//command := exec.Command("python", "src/test.py")
	////command.Stderr = os.Stderr
	//
	//buffer := bytes.Buffer{}
	//fmt.Println(convertToJSON(county))
	//buffer.Write([]byte(convertToJSON(county)))
	//command.Stdin = &buffer
	//
	//out, err := command.Output()
	//
	//if err != nil {
	//	//fmt.Println(convertToJSON(county))
	//	println("FUCK:", err.Error())
	//	return
	//}
	//println("F: ", out)

	//result, err := strconv.Atoi(strings.TrimSpace(string(out)))
	//fmt.Println(result)

	//c := make(chan Entry)
	//for _, county := range counties {
	//	county := county
	//	go func() {
	//		c <- predict(county.name, convertToJSON(county))
	//	}()
	//}
	//for range counties {
	//	prediction := <-c
	//	fmt.Println("FINAL GOT PRED: "+prediction.key, prediction.value)
	//}
}

//func convertToJSON(arr []County) string {
//	var buff bytes.Buffer
//	for i := 0; i < len(arr); i++ {
//		j, _ := json.Marshal(arr[i])
//		buff.WriteString(string(j))
//	}
//	fmt.Println(buff.String())
//	return buff.String()
//}

func run(county County) string {
	conn, err := net.Dial("tcp", "127.0.0.1:5000")
	if err != nil {
		println("CE:", err)
	}
	//c <- predict(county.name, convertToJSON(county))
	if _, err = conn.Write([]byte(convertToJSON(county))); err != nil {
		println("E:", err.Error())
	}

	buff := make([]byte, 128)
	if _, err = conn.Read(buff); err != nil {
		println("EF:", err.Error())
	}
	conn.Close()
	//c <- string(buff)
	return string(buff)
}

func convertToJSON(county County) string {
	var buff bytes.Buffer
	buff.WriteString("[14, [")

	last := len(county.timeline) - 1
	for j, cases := range county.timeline {
		buff.WriteString(strconv.Itoa(cases))
		if j < last {
			buff.WriteString(",")
		}
	}
	buff.WriteString("]")
	buff.WriteString("]")
	return buff.String()
}
