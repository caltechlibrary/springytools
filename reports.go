// reports.go provides the functions that work in an input filename and output filename
// generating reports or data conversions.
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
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

// LibGuidesXMLFileToJSONFile reads in a LibGuides XML export file and writes
// a JSON version of the file. It expects the name of the XML file in srcName
// the name of the JSON file in destName. It will return an error if any
// encountered.
func LibGuidesXMLFileToJSONFile(srcName, destName string) error {
	var (
		src []byte
		err error
	)
	lg := LibGuides{
		XMLName:  xml.Name{},
		Customer: &Customer{},
		Site:     &Site{},
		Accounts: []*Account{},
		Groups:   []*Group{},
		Subjects: []*Subject{},
		Tags:     []*Tag{},
		Vendors:  []*Vendor{},
		Guides:   []*Guide{},
	}
	src, err = ioutil.ReadFile(srcName)
	if err != nil {
		return err
	}
	err = lg.FromXML(src)
	if err != nil {
		return err
	}
	src, err = lg.ToJSON()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(destName, src, 0777)
	if err != nil {
		return err
	}
	return nil
}

// LinkReport reads in a LibGuides XML export and generates a link report
// encoded in JSON. Accepts a srcName (LibGuides XML export), destName, format
// (i.e. csv, json, xml). Returns an error if any encountered.
func LinkReport(srcName, destName, format string) error {
	var (
		src    []byte
		err    error
		rptFmt string
	)
	// This maps variations for the three supported formats of CSV, JSON and XML (as HTML)
	formats := map[string]string{
		"CSV":   "csv",
		"csv":   "csv",
		"JSON":  "json",
		"json":  "json",
		"HTML":  "xml",
		"html":  "xml",
		"XML":   "xml",
		"xml":   "xml",
		".csv":  "csv",
		".json": "json",
		".html": "xml",
		".xml":  "xml",
	}
	rptFmt = "csv"
	if val, ok := formats[format]; ok {
		rptFmt = val
	} else {
		return fmt.Errorf("%q is not a supported format", format)
	}

	lg := LibGuides{
		XMLName:  xml.Name{},
		Customer: &Customer{},
		Site:     &Site{},
		Accounts: []*Account{},
		Groups:   []*Group{},
		Subjects: []*Subject{},
		Tags:     []*Tag{},
		Vendors:  []*Vendor{},
		Guides:   []*Guide{},
	}
	src, err = ioutil.ReadFile(srcName)
	if err != nil {
		return err
	}
	err = lg.FromXML(src)
	if err != nil {
		return err
	}
	sitePrefix := "https://libguides.example.edu"
	if lg.Site != nil {
		sitePrefix = fmt.Sprintf("https://%s", lg.Site.Domain)
	}

	// Prep our reporting datastructure
	tbl := new(Table)
	tbl.SetCaption(fmt.Sprintf("Link report for %q", srcName))
	tbl.AppendHeadings([]string{"URL (partial)", "Object Type", "Id", "LibGuides Link", "Embedded"}...)

	// Traverse over each section of the export before finally analyziing the
	// data in the "guides" element. (tags and vendors are skipped, no URL data)
	for _, account := range lg.Accounts {
		if account.Website != "" {
			tbl.AppendRow(account.Website,
				"Account",
				fmt.Sprintf("%d", account.Id),
				"", "false")
		}
	}
	for _, group := range lg.Groups {
		if group.Url != "" {
			tbl.AppendRow(group.Url, "Group",
				fmt.Sprintf("%d", group.Id),
				group.Url, "false")
		}
	}
	for _, subject := range lg.Subjects {
		if subject.Url != "" {
			tbl.AppendRow(subject.Url, "Subject",
				fmt.Sprintf("%d", subject.Id),
				fmt.Sprintf("%s/sb.php?subject_id=%d", sitePrefix, subject.Id),
				"false")
		}
	}
	// Now process the guides, pages and assets
	for _, guide := range lg.Guides {
		if guide.Url != "" {
			// Note this is the Lib Guide URL
			tbl.AppendRow(guide.Url, "Guide",
				fmt.Sprintf("%d", guide.Id),
				guide.Url,
				"false")
		}
		group := guide.Group
		if group.Url != "" {
			tbl.AppendRow(group.Url, fmt.Sprintf("Guide (%d) Group", guide.Id),
				fmt.Sprintf("%d", group.Id),
				group.Url, "false")
		}

		for _, subject := range guide.Subjects {
			if subject.Url != "" {
				tbl.AppendRow(subject.Url, fmt.Sprintf("Guide (%d) Subject", guide.Id),
					fmt.Sprintf("%d", subject.Id),
					subject.Url, "false")
			}
		}
		for _, page := range guide.Pages {
			// NOTE: Only report the page if it is not hidden
			if page.Hidden == 0 {
				if page.Url != "" {
					tbl.AppendRow(page.Url,
						fmt.Sprintf("Guide (%d) Page", guide.Id),
						fmt.Sprintf("%d", page.Id),
						page.Url, "false")
				}
				if page.Description != "" {
					// NOTE: Scan for embedded URLs in the description
					if urlList, cnt := ExtractHTTPLinks(page.Description); cnt > 0 {
						for i := 0; i < cnt; i++ {
							tbl.AppendRow(urlList[i],
								fmt.Sprintf("Guide (%d) Page (%d) Description", guide.Id, page.Id),
								fmt.Sprintf("%d of %d", i+1, cnt),
								fmt.Sprintf("%s/c.php?g=%d&p=%d", sitePrefix, guide.Id, page.Id),
								"true")
						}
					}
				}
				for _, box := range page.Boxes {
					// Process box Assets
					if box.Hidden == 0 {
						for _, asset := range box.Assets {
							if asset.Url != "" {
								tbl.AppendRow(asset.Url,
									fmt.Sprintf("Guide (%d) Page (%d) Boxes (%d) Asset", guide.Id, page.Id, box.Id),
									fmt.Sprintf("%d", asset.Id),
									fmt.Sprintf("%s/c.php?g=%d&p=%d", sitePrefix, guide.Id, page.Id),
									"false")
							}
							if asset.Description != "" {
								// NOTE: Scan for embedded URLs in the description
								if urlList, cnt := ExtractHTTPLinks(asset.Description); cnt > 0 {
									for i := 0; i < cnt; i++ {
										tbl.AppendRow(urlList[i],
											fmt.Sprintf("Guide (%d) Page (%d) Boxes (%d) Asset (%d) Description", guide.Id, page.Id, box.Id, asset.Id),
											fmt.Sprintf("%d of %d", i+1, cnt),
											fmt.Sprintf("%s/c.php?g=%d&p=%d", sitePrefix, guide.Id, page.Id),
											"true")
									}
								}
							}
						}
						for i, pane := range box.Panes {
							for _, asset := range pane.Assets {
								if asset.Url != "" {
									tbl.AppendRow(asset.Url,
										fmt.Sprintf("Guide (%d) Page (%d) Boxes (%d) Pane (%d) Asset", guide.Id, page.Id, box.Id, i+1),
										fmt.Sprintf("%d", asset.Id),
										fmt.Sprintf("%s/c.php?g=%d&p=%d", sitePrefix, guide.Id, page.Id),
										"false")
								}
								if asset.Description != "" {
									// NOTE: Scan for embedded URLs in the description
									if urlList, cnt := ExtractHTTPLinks(asset.Description); cnt > 0 {
										for k := 0; k < cnt; k++ {
											tbl.AppendRow(urlList[k],
												fmt.Sprintf("Guide (%d) Page (%d) Boxes (%d) Pane (%d) Asset (%d) Description", guide.Id, page.Id, box.Id, i+1, asset.Id),
												fmt.Sprintf("%d of %d", k+1, cnt),
												fmt.Sprintf("%s/c.php?g=%d&p=%d", sitePrefix, guide.Id, page.Id),
												"true")
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	// Format output
	switch rptFmt {
	case "csv":
		err = tbl.ToCSVFile(destName, true)
	case "xml":
		err = tbl.ToXMLFile(destName)
	case "json":
		err = tbl.ToJSONFile(destName)
	}
	return err
}
