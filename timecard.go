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
	args := os.Args
	if len(args) < 2 || len(args) > 3 {
		usage(args)
		return
	}
	if len(args) == 2 {
		if args[1] == "help" {
			fmt.Println("show help")
		} else {
			usage(args)
		}
	} else {
		projectName := args[1]
		command := args[2]
		switch command {
		case "in":
			checkIn(projectName)
		case "out":
			checkOut(projectName)
		case "hours":
			listTime(projectName)
		default:
			fmt.Printf("Error: invalid command [%s]\n", command)
			usage(args)
		}
	}
}

func usage(args []string) {
	fmt.Printf("Usage: %s <project_name> <command>\n", args[0])
	fmt.Printf("For help type: %s help\n", args[0])
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
