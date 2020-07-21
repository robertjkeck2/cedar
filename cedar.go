package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

const (
	branchClimbError = "Unable to climb branch. Try again."
	branchGrowError  = "Unable to grow branch. Try again."
	cedarDir         = "/.cedar/"
	dateFormat       = "01-02-2006"
	githubError      = "Unable to connect GitHub. Try again."
	leafGrowError    = "Unable to grow leaf. Try again."
	timeFormat       = "15:04:05"
)

// Branch is a collection of Leaves for a given day
type Branch struct {
	Date     string
	Filepath string
	Leaves   []Leaf
}

// Leaf is a low-level, text log
type Leaf struct {
	Time string
	Text string
}

func main() {
	var branch Branch
	var leaf Leaf

	// Initialize date and time for timestamping
	currentDate, currentTime := getCurrentDateAndTime()

	// Get user home directory for use in filepath creation
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("unable to access user homeDir")
	}

	// Grow a Branch if one does not exist
	branch.Grow(currentDate, home)

	// Parse CLI arguments
	flag.Parse()
	args := flag.Args()
	switch len(args) {
	case 0:
		branch.Climb()
	default:
		matched, err := regexp.Match(`^https:\/\/github.com\/.+.git`, []byte(args[0]))
		if err != nil {
			log.Fatal(githubError)
		}
		if matched && len(args) == 1 {
			connectGitHub(args[0], home)
		} else {
			leaf.Grow(branch, currentTime, args)
			syncGitHub(branch, home)
		}
	}
}

// Climb outputs all the Leaves for a given Branch
func (b *Branch) Climb() {
	b.Read()
	if len(b.Leaves) > 0 {
		for _, leaf := range b.Leaves {
			fmt.Println(leaf.Time + "->" + leaf.Text)
		}
	}
}

// Grow adds a new branch on a given date if one does not already exist
func (b *Branch) Grow(date string, homeDir string) {
	b.Date = date
	b.Filepath = homeDir + cedarDir + b.Date
	createDirectoryIfNotExist(homeDir + cedarDir)
	createFileIfNotExist(b.Filepath)
}

// Read formats the contents of a file into a Branch
func (b *Branch) Read() {
	f, err := os.Open(b.Filepath)
	if err != nil {
		log.Fatal(branchClimbError)
	}
	defer f.Close()
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		log.Fatal(branchClimbError)
	}
	for _, l := range lines {
		leaf := Leaf{
			Time: strings.Split(l, "->")[0],
			Text: strings.Split(l, "->")[1],
		}
		b.Leaves = append(b.Leaves, leaf)
	}
}

// Grow adds a new Leaf on to a given Branch
func (l *Leaf) Grow(b Branch, time string, args []string) {
	l.Text = strings.Join(args[:], " ")
	l.Time = time
	f, err := os.OpenFile(b.Filepath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(branchClimbError)
	}
	defer f.Close()
	if _, err := f.WriteString(l.Time + " -> " + l.Text + "\n"); err != nil {
		log.Fatal(leafGrowError)
	}
}

// Set up GitHub repo for cedar log syncing
func connectGitHub(repo string, homeDir string) {
	err := exec.Command("git", "-C", homeDir+cedarDir, "init", ".").Run()
	if err != nil {
		log.Fatal(githubError)
	}
	err = exec.Command("git", "-C", homeDir+cedarDir, "remote", "add", "origin", repo).Run()
	if err != nil {
		log.Fatal(githubError)
	}
	err = exec.Command("git", "-C", homeDir+cedarDir, "add", ".").Run()
	if err != nil {
		log.Fatal(githubError)
	}
	err = exec.Command("git", "-C", homeDir+cedarDir, "commit", "-m", "initial cedar repo setup").Run()
	if err != nil {
		log.Fatal(githubError)
	}
	err = exec.Command("git", "-C", homeDir+cedarDir, "push", "origin", "master").Run()
	if err != nil {
		log.Fatal(githubError)
	}
}

// Sync the current cedar logs to GitHub
func syncGitHub(b Branch, homeDir string) {
	err := exec.Command("git", "-C", homeDir+cedarDir, "add", ".").Run()
	if err != nil {
		log.Fatal(githubError)
	}
	err = exec.Command("git", "-C", homeDir+cedarDir, "commit", "-m", "new log entry for "+b.Date).Run()
	if err != nil {
		log.Fatal(githubError)
	}
	err = exec.Command("git", "-C", homeDir+cedarDir, "pull", "origin", "master").Run()
	if err != nil {
		log.Fatal(githubError)
	}
	err = exec.Command("git", "-C", homeDir+cedarDir, "push", "origin", "master").Run()
	if err != nil {
		log.Fatal(githubError)
	}
}

// Get string formatted current date and time
func getCurrentDateAndTime() (currentDate, currentTime string) {
	today := time.Now()
	currentDate = today.Format(dateFormat)
	currentTime = today.Format(timeFormat)
	return
}

// Create a directory if it does not already exist
func createDirectoryIfNotExist(dirFilepath string) {
	if _, err := os.Stat(dirFilepath); os.IsNotExist(err) {
		err = os.Mkdir(dirFilepath, 0777)
		if err != nil {
			log.Fatal(branchGrowError)
		}
	}
}

// Create a file if it does not already exist
func createFileIfNotExist(filepath string) {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		_, err = os.Create(filepath)
		if err != nil {
			log.Fatal(branchGrowError)
		}
	}
}
