package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

var showVersion bool
var outputPath string
var defaultVersion string
var verbose bool
var VERSION = "v0"

func init() {
	flag.BoolVar(&showVersion, "version", false, "Show version details")
	flag.BoolVar(&verbose, "verbose", false, "Verbose logging")
	flag.StringVar(&outputPath, "output", "", "File to be generated")
	flag.StringVar(&defaultVersion, "defaultVersion", "[unversioned]", "Default version string")
	flag.Parse()

	if showVersion {
		fmt.Println(VERSION)
		os.Exit(0)
	}
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	gitVers, err := getGitVersion(cwd)
	if err != nil {
		if verbose {
			log.Printf("Failed to find git version information: %s", err.Error())
		}
		gitVers = defaultVersion
	}
	version := fmt.Sprintf("%s (%s)", gitVers, time.Now().Format(time.RFC3339))

	if outputPath != "" {
		file, err := os.Create(outputPath)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer file.Close()
		fmt.Fprintf(file, generateFileData(version))
	} else {
		fmt.Print(version)
	}
}

func generateFileData(vers string) string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "package main\n\n")
	fmt.Fprintf(&buf, "const VERSION = \"%s\"", vers)
	return buf.String()
}

func getGitVersion(p string) (string, error) {
	cmd := exec.Command("git", "describe", "--tags", "--long")
	cmd.Dir = p
	data, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}
