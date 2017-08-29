// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_TxChoice struct {
	StrRef *CT_StrRef
	Rich   *drawingml.CT_TextBody
}

func NewCT_TxChoice() *CT_TxChoice {
	ret := &CT_TxChoice{}
	return ret
}
func (m *CT_TxChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.StrRef != nil {
		sestrRef := xml.StartElement{Name: xml.Name{Local: "strRef"}}
		e.EncodeElement(m.StrRef, sestrRef)
	}
	if m.Rich != nil {
		serich := xml.StartElement{Name: xml.Name{Local: "rich"}}
		e.EncodeElement(m.Rich, serich)
	}
	return nil
}
func (m *CT_TxChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TxChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "strRef":
				m.StrRef = NewCT_StrRef()
				if err := d.DecodeElement(m.StrRef, &el); err != nil {
					return err
				}
			case "rich":
				m.Rich = drawingml.NewCT_TextBody()
				if err := d.DecodeElement(m.Rich, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TxChoice
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_TxChoice) Validate() error {
	return m.ValidateWithPath("CT_TxChoice")
}
func (m *CT_TxChoice) ValidateWithPath(path string) error {
	if m.StrRef != nil {
		if err := m.StrRef.ValidateWithPath(path + "/StrRef"); err != nil {
			return err
		}
	}
	if m.Rich != nil {
		if err := m.Rich.ValidateWithPath(path + "/Rich"); err != nil {
			return err
		}
	}
	return nil
}