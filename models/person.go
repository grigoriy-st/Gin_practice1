package models

import "encoding/xml"

type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:"firstName"`
	LastName  string   `xml:"lastName"`
}
