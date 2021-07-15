// libguides.go implements the data structures for working with
// with LibGuides exported XML.
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
	"encoding/json"
	"encoding/xml"
)

type Customer struct {
	XMLName  xml.Name `xml:"customer" json:"-"`
	Id       int      `xml:"id" json:"id"`
	Type     string   `xml:"type" json:"type"`
	Name     string   `xml:"name" json:"name"`
	Url      string   `xml:"url" json:"url"`
	City     string   `xml:"city" json:"city"`
	State    string   `xml:"state" json:"state"`
	Country  string   `xml:"country" json:"country"`
	TimeZone string   `xml:"time_zone" json:"time_zone"`
	Created  string   `xml:"created" json:"created"`
	Updated  string   `xml:"updated" json:"updated"`
}

type Site struct {
	XMLName xml.Name `xml:"site" json:"-"`
	Id      int      `xml:"id" json:"jd"`
	Type    string   `xml:"type" json:"type"`
	Name    string   `xml:"name" json:"name"`
	Domain  string   `xml:"domain" json:"domain"`
	Admin   string   `xml:"admin" json:"admin"`
	Created string   `xml:"created" json:"created"`
	Updated string   `xml:"updated" json:"updated"`
}

type Account struct {
	Id        int    `xml:"id" json:"id"`
	Email     string `xml:"email" json:"email"`
	FirstName string `xml:"first_name" json:"first_name"`
	LastName  string `xml:"last_name" json:"last_name"`
	Title     string `xml:"title" json:"title"`
	Nickname  string `xml:"nickname" json:"nickname"`
	Signature string `xml:"signature" json:"signature"`
	Image     string `xml:"image" json:"image"`
	Address   string `xml:"address" json:"address"`
	Phone     string `xml:"phone" json:"phone"`
	Skype     string `xml:"skype" json:"skype"`
	Website   string `xml:"website" json:"website"`
	Created   string `xml:"created" json:"created"`
	Updated   string `xml:"updated" json:"updated"`
}

type Group struct {
	Id          int    `xml:"id" json:"id"`
	Type        string `xml:"type" json:"type"`
	Name        string `xml:"name" json:"name"`
	Url         string `xml:"url" json:"url"`
	Description string `xml:"description" json:"description"`
	Password    string `xml:"password" json:"password"`
	Created     string `xml:"created" json:"created"`
	Updated     string `xml:"updated" json:"updated"`
}

type Subject struct {
	Id   int    `xml:"id" json:"id"`
	Name string `xml:"name" json:"name"`
	Url  string `xml:"url" json:"url"`
}

type Tag struct {
	Id   int    `xml:"id" json:"id"`
	Name string `xml:"name" json:"name"`
}

type Vendor struct {
	Id   int    `xml:"id" json:"id"`
	Name string `xml:"name" json:"name"`
}

type Owner struct {
	XMLName   xml.Name `xml:"owner" json:"owner"`
	Id        int      `xml:"id" json:"id"`
	Email     string   `xml:"email" json:"email"`
	FirstName string   `xml:"first_name" json:"first_name"`
	LastName  string   `xml:"last_name" json:"last_name"`
	Image     string   `xml:"image" json:"image"`
}

type Asset struct {
	Id   int    `xml:"id" json:"id"`
	Name string `xml:"name" json:"name"`
	Type string `xml:"type" json:"type"`
	// Description contains HTML encoded text, double encoding existing encoded text
	Description string `xml:"description" json:"description"`
	Url         string `xml:"url" json:"url"`
	Owner       Owner  `xml:"owner" json:"owner"`
	MapId       string `xml:"map_id" json:"map_id"`
	Position    int    `xml:"position" json:"position"`
	Created     string `xml:"created" json:"created"`
	Updated     string `xml:"updated" json:"updated"`
}

type Pane struct {
	Assets []*Asset `xml:"assets>asset" json:"assets"`
}

type Box struct {
	XMLName  xml.Name `xml:"box" json:"box"`
	Id       int      `xml:"id" json:"id"`
	Name     string   `xml:"name" json:"name"`
	Type     string   `xml:"type" json:"type"`
	MapId    string   `xml:"map_id" json:"map_id"`
	Column   int      `xml:"column" json:"column"`
	Position int      `xml:"position" json:"position"`
	Hidden   int      `xml:"hidden" json:"hidden"`
	Created  string   `xml:"created" json:"created"`
	Updated  string   `xml:"updated" json:"updated"`
	Assets   []*Asset `xml:"assets>asset" json:"assets"`
	Panes    []*Pane  `xml:"panes>pane,omitempty" json:"panes,omitempty"`
}

type Page struct {
	Id           int    `xml:"id" json:"id"`
	Name         string `xml:"name" json:"name"`
	Description  string `xml:"description" json:"description"`
	Url          string `xml:"url" json:"url"`
	Redirect     string `xml:"redirect" json:"redirect"`
	SourcePageId int    `xml:"source_page_id" json:"source_page_id"`
	ParentPageId int    `xml:"parent_page_id" json:"parent_page_id"`
	Position     int    `xml:"position" json:"position"`
	Hidden       int    `xml:"hidden" json:"hidden"`
	Created      string `xml:"created" json:"created"`
	Updated      string `xml:"updated" json:"updated"`
	Modified     string `xml:"modified" json:"modified"`
	Boxes        []*Box `xml:"boxes>box" json:"boxes"`
}

type Guide struct {
	Id          int        `xml:"id" json:"id"`
	Type        string     `xml:"type" json:"type"`
	Name        string     `xml:"name" json:"name"`
	Description string     `xml:"description" json:"description"`
	Url         string     `xml:"url" json:"url"`
	Owner       Owner      `xml:"owner" json:"owner"`
	Group       Group      `xml:"group" json:"group"`
	Redirect    string     `xml:"redirect" json:"redirect"`
	Status      string     `xml:"status" json:"status"`
	Created     string     `xml:"created" json:"created"`
	Updated     string     `xml:"updated" json:"updated"`
	Modified    string     `xml:"modified" json:"modified"`
	Published   string     `xml:"published" json:"published"`
	Subjects    []*Subject `xml:"subjects>subject" json:"subjects"`
	Tags        []*Tag     `xml:"tags>tag" json:"tags"`
	Pages       []*Page    `xml:"pages>page" json:"pages"`
}

type LibGuides struct {
	XMLName  xml.Name   `json:"-"`
	Customer *Customer  `xml:"customer" json:"customer"`
	Site     *Site      `xml:"site" json:"site"`
	Accounts []*Account `xml:"accounts>account" json:"accounts"`
	Groups   []*Group   `xml:"groups>group" json:"groups"`
	Subjects []*Subject `xml:"subjects>subject" json:"subjects"`
	Tags     []*Tag     `xml:"tags>tag" json:"tags"`
	Vendors  []*Vendor  `xml:"vendors>vendor" json:"vendors"`
	Guides   []*Guide   `xml:"guides>guide" json:"guides"`
}

// FromXML takes a LibGuides Object, []bytes of XML source
// populates the LibGuides object and returns any error.
func (lg *LibGuides) FromXML(src []byte) error {
	if lg == nil {
		lg = new(LibGuides)
	}
	return xml.Unmarshal(src, &lg)
}

// ToJSON takes a LibGuides object and renders JSON output and error
func (lg *LibGuides) ToJSON() ([]byte, error) {
	var (
		src []byte
		err error
	)
	src, err = json.MarshalIndent(lg, "", "    ")
	if err != nil {
		return nil, err
	}
	return src, nil
}
