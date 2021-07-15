// tables_test.go provides tests for tables.go
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
package springytools

import (
	"bufio"
	"encoding/xml"
	"os"
	"testing"
)

func TestTableBuilding(t *testing.T) {
	tbl := Table{
		XMLName: xml.Name{},
		Caption: "",
		Head:    THead{},
		Body:    TBody{},
	}
	tbl.SetCaption("This is a table")
	tbl.AppendHeadings([]string{"One", "Two", "Three"}...)
	tbl.AppendRow([]string{"1", "2", "3"}...)
	tbl.AppendRow([]string{"4", "5", "6"}...)
	if tbl.Caption != "This is a table" {
		expectedString(t, "This is a table", tbl.Caption)
	}
	expectedString(t, "One", tbl.Head.Row[0])
	expectedString(t, "Two", tbl.Head.Row[1])
	expectedString(t, "Three", tbl.Head.Row[2])
	expectedString(t, "1", tbl.Body.Rows[0][0])
	expectedString(t, "2", tbl.Body.Rows[0][1])
	expectedString(t, "3", tbl.Body.Rows[0][2])
	expectedString(t, "4", tbl.Body.Rows[1][0])
	expectedString(t, "5", tbl.Body.Rows[1][1])
	expectedString(t, "6", tbl.Body.Rows[1][2])
}

func TestTableWriting(t *testing.T) {
	// tbl := Table{
	// 	XMLName: xml.Name{},
	// 	Caption: "This is a table",
	// 	Head:    THead{Row{"One","two", "Three"}},
	// 	Body:    TBody{Rows[][]string{
	// 		[]string{"1","2","3"},
	// 		[]string{"4","5","6"},
	// 	}},
	// }

	tbl := Table{
		XMLName: xml.Name{},
		Caption: "",
		Head:    THead{},
		Body:    TBody{},
	}
	tbl.SetCaption("This is a table")
	tbl.AppendHeadings([]string{"One", "Two", "Three"}...)
	tbl.AppendRow([]string{"1", "2", "3"}...)
	tbl.AppendRow([]string{"4", "5", "6"}...)

	fName := "testout/table.csv"
	if err := tbl.ToCSVFile("testout/table.csv", true); err != nil {
		t.Errorf("Write fail for %q: %s", fName, err)
	}
	// I need to read the table output back in and make sure it wrote correctly.
	if fp, err := os.Open(fName); err != nil {
		t.Errorf("Failed to read %q: %s", fName, err)
	} else {
		defer fp.Close()
		scanner := bufio.NewScanner(fp)
		scanner.Split(bufio.ScanLines)
		expectedLines := []string{
			`One,Two,Three`,
			`1,2,3`,
			`4,5,6`,
		}
		i := 0
		for scanner.Scan() {
			expectedString(t, expectedLines[i], scanner.Text())
			i++
		}
	}
	//FIXME: Need to test xml (HTML) output and JSON output
	t.Errorf("Testing for XML (HTML) output not implemented.")
	t.Errorf("Testing for JSON output not implemented.")
}
