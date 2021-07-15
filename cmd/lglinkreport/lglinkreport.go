// linkreport.go traverse all the fields that have links and reports
// where they are found. Output is JSON.
//
// Author: R. S. Doiel <rsdoiel@caltech.edu>
//
// Copyright (c) 2021, Caltech
// All rights not granted herein are expressly reserved by Caltech.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
// this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
// this list of conditions and the following disclaimer in the documentation
// and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors
// may be used to endorse or promote products derived from this software without
// specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.
//
package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	// Caltech Library Package
	"github.com/caltechlibrary/springytools"
)

func usage(appName string, exitCode int) {
	out := os.Stderr
	if exitCode == 0 {
		out = os.Stdout
	}
	fmt.Fprintf(out, `
USAGE: %s

    %s SOURCE_FILE DESTINATION_FILE

Reads a LibGuides' XML export and generates JSON reporting
on links founds and where they were found.

OPTIONS

    -h, -help      display help
    -fmt FORMAT    set the output format, 
                   i.e. csv (defualt), json, xml

EXAMPLE

    %s LibGuides_export_221133.xml links.json

springytools v%s
`, appName, appName, appName, springytools.Version)
	os.Exit(exitCode)
}

func main() {
	// command line name and options support
	appName := path.Base(os.Args[0])
	help, version := false, false
	format := "csv"
	args := []string{}
	// Setup to parse command line
	flag.BoolVar(&help, "h", false, "display help")
	flag.BoolVar(&help, "help", false, "display help")
	flag.BoolVar(&version, "version", false, "display version")
	flag.StringVar(&format, "format", format, "output report using format (i.e. csv, json, xml)")
	flag.Parse()

	args = flag.Args()

	// Process options and run report
	if help {
		usage(appName, 0)
	}
	if version {
		fmt.Printf("springytools, %s v%s\n", appName, springytools.Version)
		os.Exit(0)
	}
	if len(args) != 2 {
		fmt.Printf("Missing source or destination names\n\n")
		usage(appName, 1)
	}
	err := springytools.LinkReport(args[0], args[1], format)
	if err != nil {
		fmt.Printf("ERROR: %s", err)
		os.Exit(1)
	}
}
