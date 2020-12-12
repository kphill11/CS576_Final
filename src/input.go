package main

import(
      "encoding/csv"
      "fmt"
      "io"
      "os"
      "strings"
      "strconv"
)

type County struct {
      name string
      numBeds int
      pop int
      timeline [95]int
}

func newCounty(newName string) County {
      county := County{}
      county.name = newName
      county.numBeds = 0
      county.pop = 0
      for i := 0; i < len(county.timeline); i++ {
            county.timeline[i] = 0
      }
      return county
}

func findCounty(counties []County, name string) int {
      for i := 0; i < len(counties); i++ {
            if(counties[i].name == name) {
                  return i
            }
      }
      //if you've searched through the entire array and haven't found it
      return -1
}

func readInput(counties []County) []County {
      //hospitals.csv:
            //field 13 is population
            //field 14 is county
            //field 32 is beds
            //field 33 is trauma (likely not going to use)
      //var counties []County
      csvHospFile, errHosp := os.Open("..\\data\\hospitals.csv")
      if errHosp != nil {
            fmt.Printf("Error while loading hospitals.csv:\n")
            fmt.Printf("\n%s\n", errHosp)
            os.Exit(3)
      }
      reader := csv.NewReader(csvHospFile)
      for {
            line, err := reader.Read()
            if err == io.EOF {
                  break
            } else if err != nil {
                  fmt.Printf("\nerror while reading hospitals.csv:\n")
                  fmt.Printf("\n%s\n",err)
                  os.Exit(3)
            } else {
                  name := strings.ToLower(strings.Trim(line[14], " "))
                  index := findCounty(counties, name)
                  if(index < 0) { //if the county isn't already on record
                        counties = append(counties, newCounty(name))
                        index := len(counties) - 1
                        popNum, err := strconv.Atoi(line[13]) //county population, not likely to change significantly enough to track
                        counties[index].pop = popNum
                  }
                  numBeds, err := strconv.Atoi(line[32])
                  counties[index].numBeds += numBeds //number of regular beds
            }
      }
      //covid-latest.csv
            //field 0 is date
            //field 1 is County
            //field 4 is cases
      csvCovidFile, errCov := os.Open("..\\data\\covid-latest.csv")
      if errCov != nil {
            fmt.Printf("\nError while loading covid-latest.csv:\n")
            fmt.Printf("%s\n", errCov)
            os.Exit(3)
      }
      reader = csv.NewReader(csvCovidFile)
      date := 0
      dateString := "2020-01-21"
      for {
            line, err := reader.Read()
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
                        index := len(counties) - 1
                        //counties[index].pop = strconv.Atoi(line[13]) //county population, not likely to change significantly enough to track
                  }
                  if dateString != line[0] {
                        date++;
                        dateString = line[0]
                  }
                  caseNum, err := strconv.Atoi(line[4])
                  counties[index].timeline[date] += caseNum
            }
      }
      fmt.Printf("%d\n", len(counties))
      for i := 0; i < len(counties); i++ {
            fmt.Printf("%s: population: %d. number of beds: %d. Number of infected people on september 7th: %d\n", counties[i].name, counties[i].pop, counties[i].numBeds, counties[i].timeline[0])
      }
      return counties
}
