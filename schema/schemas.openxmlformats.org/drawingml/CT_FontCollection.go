// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_FontCollection struct {
	Latin  *CT_TextFont
	Ea     *CT_TextFont
	Cs     *CT_TextFont
	Font   []*CT_SupplementalFont
	ExtLst *CT_OfficeArtExtensionList
}

func NewCT_FontCollection() *CT_FontCollection {
	ret := &CT_FontCollection{}
	ret.Latin = NewCT_TextFont()
	ret.Ea = NewCT_TextFont()
	ret.Cs = NewCT_TextFont()
	return ret
}
func (m *CT_FontCollection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	selatin := xml.StartElement{Name: xml.Name{Local: "a:latin"}}
	e.EncodeElement(m.Latin, selatin)
	seea := xml.StartElement{Name: xml.Name{Local: "a:ea"}}
	e.EncodeElement(m.Ea, seea)
	secs := xml.StartElement{Name: xml.Name{Local: "a:cs"}}
	e.EncodeElement(m.Cs, secs)
	if m.Font != nil {
		sefont := xml.StartElement{Name: xml.Name{Local: "a:font"}}
		e.EncodeElement(m.Font, sefont)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_FontCollection) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Latin = NewCT_TextFont()
	m.Ea = NewCT_TextFont()
	m.Cs = NewCT_TextFont()
lCT_FontCollection:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "latin":
				if err := d.DecodeElement(m.Latin, &el); err != nil {
					return err
				}
			case "ea":
				if err := d.DecodeElement(m.Ea, &el); err != nil {
					return err
				}
			case "cs":
				if err := d.DecodeElement(m.Cs, &el); err != nil {
					return err
				}
			case "font":
				tmp := NewCT_SupplementalFont()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Font = append(m.Font, tmp)
			case "extLst":
				m.ExtLst = NewCT_OfficeArtExtensionList()
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
			break lCT_FontCollection
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_FontCollection) Validate() error {
	return m.ValidateWithPath("CT_FontCollection")
}
func (m *CT_FontCollection) ValidateWithPath(path string) error {
	if err := m.Latin.ValidateWithPath(path + "/Latin"); err != nil {
		return err
	}
	if err := m.Ea.ValidateWithPath(path + "/Ea"); err != nil {
		return err
	}
	if err := m.Cs.ValidateWithPath(path + "/Cs"); err != nil {
		return err
	}
	for i, v := range m.Font {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Font[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}