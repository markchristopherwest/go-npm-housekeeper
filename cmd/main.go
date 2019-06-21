// Main
package main

// Import Modules
import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

type executiveClass struct {
	Command   string
	Directory string
}

func (class executiveClass) StructMethod() {
	fmt.Println(class.Command)

}

func pick(cli executiveClass, cwd string) {

	fmt.Println("\nRunning in Directory: ", cwd)
	fmt.Println("\nExecute this command: ", cli.Command)
	// Split up the command line arguments using space as a delimiter
	parts := strings.Split(cli.Command, " ")

	// The first part is the command, the rest are the args:
	head := parts[0]
	args := parts[1:len(parts)]

	// Format the command
	cmd := exec.Command(head, args...)
	cmd.Dir = cwd

	// Handle the Output
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalln("\ncmd.Run() our with out: ", string(out))
		log.Fatalln("\ncmd.Run() err with err: ", err)
	}
	fmt.Println("\nSuccess @ cmd.Run() command: ", cli.Command)
	fmt.Println("\nSuccess @ cmd.Run() outputs: ", string(out))

}

// Create Main within Path
func main() {

	// Define The Path
	var goWalkPath = "/Users/mark/Sites"

	// Declare Function as Variable
	var goWalkFunction = func(thisPath string, thisFileInfo os.FileInfo, thisErrorHandler error) error {

		// first thing to do, check error. and decide what to do about it
		if thisErrorHandler != nil {
			fmt.Printf("ERROR: 「%v」 @ 「%q」\n", thisErrorHandler, thisPath)
			return thisErrorHandler
		}

		if thisFileInfo.IsDir() && thisFileInfo.Name() == ".git" {
			// Ignore git
			return filepath.SkipDir
		}

		if thisFileInfo.IsDir() && thisFileInfo.Name() == ".npm" {
			// Ignore git
			return filepath.SkipDir
		}

		if thisFileInfo.IsDir() && thisFileInfo.Name() == "phpmyadmin" {
			// Ignore git
			return filepath.SkipDir
		}

		// Skip node_modules
		if thisFileInfo.IsDir() && thisFileInfo.Name() == "node_modules" {
			// Loop is deleting node_modules before generating them
			return filepath.SkipDir
		}

		if thisFileInfo.Name() == "package.json" {
			path := filepath.Dir(thisPath)
			//	Set these Command
			cmd0 := executiveClass{Command: "rm -rfv package-lock.json"}
			cmd1 := executiveClass{Command: "rm -rf node_modules"}
			cmd2 := executiveClass{Command: "npm update --save-dev"}
			cmd3 := executiveClass{Command: "npm update --save"}
			cmd4 := executiveClass{Command: "npm i --package-lock-only"}
			cmd5 := executiveClass{Command: "npm audit fix"}
			cmd6 := executiveClass{Command: "npm i"}
			// Run these commands
			pick(cmd0, path)
			pick(cmd1, path)
			pick(cmd2, path)
			pick(cmd3, path)
			pick(cmd4, path)
			pick(cmd5, path)
			pick(cmd6, path)
			pick(cmd1, path)
			time.Sleep(5 * time.Second)

		}
		return nil
	}
	err := filepath.Walk(goWalkPath, goWalkFunction)

	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", goWalkPath, err)
	}

	cmd := executiveClass{Command: "rsync -avh /Users/user/Sites/ dude@dude.local:/home/dude/"}
	// Run these commands
	pick(cmd, "")

	fmt.Printf(InfoColor, "Notice")
	fmt.Printf("\nOperation complete.\n")

}
