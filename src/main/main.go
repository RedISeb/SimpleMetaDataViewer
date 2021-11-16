package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {

	//path := os.Args[1]
	var path string

	flag.StringVar(&path, "F", "", "specify the file path")

	flag.Parse()

	if path == "" {
		os.Exit(0)
	}

	fileStat, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}

	//powershell magic for getting creation date
	out, err2 := exec.Command("powershell", "Get-ItemProperty", path, "|", "Select-Object", "-Property", "'CreationTime'",
		"|", "Out-String", "-stream", "|", "Select-String", "-Pattern", "'CreationTime'", "-NotMatch").CombinedOutput()
	if err2 != nil {
		log.Fatal(err2)
	}
	output := string(out[:])
	output2 := strings.Split(output, "\n")[3]

	fmt.Println("File Name:", fileStat.Name())
	fmt.Println("Size:", fileStat.Size())
	fmt.Println("Permissions:", fileStat.Mode())
	fmt.Println("Created: ", output2)
	fmt.Println("Last Modified:", fileStat.ModTime())
	fmt.Println("Is Directory: ", fileStat.IsDir())

}
