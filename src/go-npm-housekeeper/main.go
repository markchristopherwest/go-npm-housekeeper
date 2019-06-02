// Main
package main

// Import Modules for handling buffer, os, etc.
import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// Create Main within Path
func main() {

	// Define The Path
	var goWalkPath = "/Users/mark/SynologyDrive/Documents/Code/"

	// Declare Function as Variable
	var goWalkFunction = func(thisPath string, thisFileInfo os.FileInfo, thisErrorHandler error) error {

		// first thing to do, check error. and decide what to do about it
		if thisErrorHandler != nil {
			fmt.Printf("ERROR: 「%v」 @ 「%q」\n", thisErrorHandler, thisPath)
			return thisErrorHandler
		}
		// Skip .DS_STORE
		if thisFileInfo.IsDir() && thisFileInfo.Name() == ".DS_STORE" {
			return filepath.SkipDir
		}
		// Skip .git
		if thisFileInfo.IsDir() && thisFileInfo.Name() == ".git" {
			return filepath.SkipDir
		}
		// Skip node_modules
		if thisFileInfo.IsDir() && thisFileInfo.Name() == "node_modules" {
			return filepath.SkipDir
		}
		// Skip node_modules
		if thisFileInfo.IsDir() && thisFileInfo.Name() == "truetlen" {
			return filepath.SkipDir
		}
		// Catch package.json		
		if thisFileInfo.Name() == "package.json" {
			// Define the Struct
			cmd:= exec.Command("npm", "audit", "fix")
			// Set the Path
			cmd.Dir = filepath.Dir(thisPath)
			// Define the Output
			out, err := cmd.CombinedOutput()
			// Handle the Errors
			if err != nil {
				fmt.Printf("cmd.Run() failed with %s\n", err)
			}
			// Print the Output
			fmt.Printf("combined out:\n%s\n", string(out))

		}
		return nil
	}
	err := filepath.Walk(goWalkPath, goWalkFunction)

	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", goWalkPath, err)
	}
}