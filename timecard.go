package main

import (
	"strconv"
	"strings"
	"fmt"
	"time"
	"os"
	"io/ioutil"
)

func main() {
	createProject("hello")
	checkIn("hello")
	time.Sleep(1000 * time.Millisecond)
	checkOut("hello")
	listTime("hello")
}

func createProject(projectName string) {
	if _, err := os.Stat(projectName); os.IsNotExist(err) {
		fmt.Println("creating project", projectName)
		os.Create(projectName)
	}
}

func checkIn(projectName string) {
	createProject(projectName)
	file, _ := os.OpenFile(projectName, os.O_RDWR | os.O_APPEND, 0666)
	file.Write([]byte(strconv.FormatInt(time.Now().Unix(), 10) + "\n"))
	file.Close()
}

func checkOut(projectName string) {
	createProject(projectName)
	file, _ := os.OpenFile(projectName, os.O_RDWR | os.O_APPEND, 0666)
	file.Write([]byte(strconv.FormatInt(time.Now().Unix(), 10) + "\n"))
	file.Close()
}

func listTime(projectName string) {
	bytes, _ := ioutil.ReadFile(projectName)
	lines := strings.Split(string(bytes), "\n")
	totalTime := int64(0)
	for i := 0; i < len(lines) - 1; i += 2 {
		startTime, _ := strconv.ParseInt(lines[i], 10, 64)
		stopTime, _ := strconv.ParseInt(lines[i+1], 10, 64)
		totalTime += (stopTime - startTime)
	}
	hours := float64(totalTime) / 60 / 60
	fmt.Println(hours)
}
