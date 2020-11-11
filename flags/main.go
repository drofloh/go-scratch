/*
Example to set some variables first by flag and then by environment variables
with the flags taking precendent.
*/
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// set some variables for username and password
	var uname, pass string

	// flags
	flag.StringVar(&uname, "u", "", "Specify username.")
	flag.StringVar(&pass, "p", "", "Specify pass.")
	flag.Parse() // after declaring flags we need to call it

	val, ok := os.LookupEnv("FUSER")
	if ok && uname == "" {
		uname = val
	}

	val, ok = os.LookupEnv("FPASS")
	if ok && pass == "" {
		pass = val
	}

	if uname == "" || pass == "" {
		fmt.Println("Must ensure both username and password are set")
		os.Exit(1)
	}

	// check if cli params match
	if uname == "root" && pass == "password" {
		fmt.Printf("Logging in")
	} else {
		fmt.Printf("Invalid credentials!")
	}
}
