package main

import (
      "encoding/csv"
      "fmt"
      "io"
      "os"
      "strconv"
      "strings"
)

type County struct {
      name string
      numBeds int
      pop int
      timeline [96]int
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
            if counties[i].name == name {
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
      counties = append(counties, newCounty("albany"))
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
                  fmt.Printf("\n%s\n",err)
                  os.Exit(3)
            } else {
                  name := strings.ToLower(strings.Trim(line[14], " "))
                  if name != "not available" {
                        index := findCounty(counties, name)
                        if index < 0 { //if the county isn't already on record
                              counties = append(counties, newCounty(name))
                              index = len(counties) - 1
                              if line[13] != "-999" {
                                   //s fmt.Printf("%s\n", line[13])
                                    popNum, err := strconv.Atoi(line[13])
                                    if err != nil {
                                          fmt.Printf("%s\n", err)
                                          os.Exit(3)
                                    }
                                    if popNum >= 0 {
                                          counties[index].pop = popNum
                                    }
                              } //county population, not likely to change significantly enough to track
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
      dateString := "2020-01-21"
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
      fmt.Printf("%d\n", len(counties))
      for i := len(counties) - 1; i > 0; i-- {
            if counties[i].numBeds == 0 {
                  counties[i] = counties[len(counties) - 1]
                  //counties[len(counties) - 1] = null
                  counties = counties[:len(counties) - 1]
                  i--
            }
            //fmt.Printf("%s: population: %d. number of beds: %d. Number of infected people on December 10th: %d\n", counties[i].name, counties[i].pop, counties[i].numBeds, counties[i].timeline[95          ])
      }
      return counties
}
