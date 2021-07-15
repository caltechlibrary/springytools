// tables.go provides a XML, JSON and CSV rendering of the Table datastructure.
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
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"os"
)

type THead struct {
	XMLName xml.Name `xml:"thead" json:"-"`
	Row     []string `xml:"tr>th" json:"columns,omitempty"`
}

type TBody struct {
	XMLName xml.Name   `xml:"tbody" json:"-"`
	Rows    [][]string `xml:"tr>td" json:"rows,omitempty"`
}

type Table struct {
	XMLName xml.Name `xml:"table" json:"-"`
	Caption string   `xml:"caption" json:"caption,omitempty"`
	Head    THead    `xml:"thead" json:"head,omitempty"`
	Body    TBody    `xml:"tbody" json:"body,omitempty"`
}

func (t *Table) SetCaption(caption string) {
	t.Caption = caption
}

func (t *Table) AppendHeadings(cells ...string) {
	t.Head.Row = append(t.Head.Row, cells...)
}

func (t *Table) AppendRow(cells ...string) {
	t.Body.Rows = append(t.Body.Rows, cells)
}

func (t *Table) ToXML() ([]byte, error) {
	return xml.MarshalIndent(t, "", "\t")
}

func (t *Table) ToJSON() ([]byte, error) {
	return json.MarshalIndent(t, "", "\t")
}

// ToXMLFile will creates an XML (HTML) version of Table, it is a destructive write.
// A file with the same name will be replaced. Accepts the filename and Returns an error
// if one is encountered.
func (t *Table) ToXMLFile(destName string) error {
	src, err := t.ToXML()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(destName, src, 0777)
}

// ToJSONFile will creates a JSON version of Table, it is a destructive write.
// A file with the same name will be replaced. Accepts the filename and Returns an error
// if one is encountered.
func (t *Table) ToJSONFile(destName string) error {
	src, err := t.ToJSON()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(destName, src, 0777)
}

// ToCSVFile will create a CSV version of Table, it is a destructive write.
// A file with the same name will be replaced. Accepts the filename and header boolean.
// if header is true and the table's header is populated it will render a header row at
// start of the CSV output. Returns an error if one is encountered.
func (t *Table) ToCSVFile(destName string, header bool) error {
	fp, err := os.Create(destName)
	if err != nil {
		return err
	}
	w := csv.NewWriter(fp)
	if header {
		if (t.Head.Row != nil) && (len(t.Head.Row) > 0) {
			if err = w.Write(t.Head.Row); err != nil {
				return err
			}
		}
	}
	for _, row := range t.Body.Rows {
		if err = w.Write(row); err != nil {
			return err
		}
	}
	w.Flush()
	err = w.Error()
	return nil
}
