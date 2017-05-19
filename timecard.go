package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

var path string = os.Getenv("GOPATH") + "/bin/"

func main() {
	args := os.Args
	if len(args) < 2 || len(args) > 3 {
		usage(args)
		return
	}
	if len(args) == 2 {
		if args[1] == "help" {
			help()
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
	if _, err := os.Stat(path + projectName); os.IsNotExist(err) {
		fmt.Println("creating project", projectName)
		os.Create(path + projectName)
	}
}

func checkIn(projectName string) {
	num := numEntries(projectName)
	if num%2 == 1 {
		fmt.Println("You have already checked into", projectName)
		return
	}
	addTimestamp(projectName)
}

func checkOut(projectName string) {
	num := numEntries(projectName)
	if num == -1 {
		fmt.Println("Error:", projectName, "does not exist.")
		return
	} else if num%2 == 0 {
		fmt.Println("You have already checked out of", projectName)
		return
	}
	addTimestamp(projectName)
}

func numEntries(projectName string) int {
	contents, err := ioutil.ReadFile(path + projectName)
	if err != nil {
		return -1
	}
	lines := strings.Split(string(contents), "\n")
	var realLines []string
	for _, line := range lines {
		if line != "" {
			realLines = append(realLines, line)
		}
	}
	return len(realLines)
}

func addTimestamp(projectName string) {
	createProject(projectName)
	file, _ := os.OpenFile(path+projectName, os.O_RDWR|os.O_APPEND, 0666)
	file.Write([]byte(strconv.FormatInt(time.Now().Unix(), 10) + "\n"))
	file.Close()
}

func listTime(projectName string) {
	bytes, _ := ioutil.ReadFile(path + projectName)
	lines := strings.Split(string(bytes), "\n")
	totalTime := int64(0)
	for i := 0; i < len(lines)-1; i += 2 {
		startTime, _ := strconv.ParseInt(lines[i], 10, 64)
		stopTime, _ := strconv.ParseInt(lines[i+1], 10, 64)
		totalTime += (stopTime - startTime)
	}
	hours := float64(totalTime) / 60 / 60
	fmt.Println(hours)
}

func help() {
	helpMsg := `Commands:
in 	: check into a project
out 	: check out of a project 
hours 	: list total hours worked on project`
	fmt.Println(helpMsg)
}
