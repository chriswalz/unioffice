// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_SlideIdListEntry struct {
	// Slide Identifier
	IdAttr  uint32
	RIdAttr string
	ExtLst  *CT_ExtensionList
}

func NewCT_SlideIdListEntry() *CT_SlideIdListEntry {
	ret := &CT_SlideIdListEntry{}
	return ret
}
func (m *CT_SlideIdListEntry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
		Value: fmt.Sprintf("%v", m.RIdAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_SlideIdListEntry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.IdAttr = uint32(parsed)
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RIdAttr = parsed
		}
	}
lCT_SlideIdListEntry:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SlideIdListEntry
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_SlideIdListEntry) Validate() error {
	return m.ValidateWithPath("CT_SlideIdListEntry")
}
func (m *CT_SlideIdListEntry) ValidateWithPath(path string) error {
	if m.IdAttr < 256 {
		return fmt.Errorf("%s/m.IdAttr must be >= 256 (have %v)", path, m.IdAttr)
	}
	if m.IdAttr >= 2147483648 {
		return fmt.Errorf("%s/m.IdAttr must be < 2147483648 (have %v)", path, m.IdAttr)
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}