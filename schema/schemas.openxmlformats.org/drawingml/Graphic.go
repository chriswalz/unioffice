// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"log"
)

type Graphic struct {
	CT_GraphicalObject
}

func NewGraphic() *Graphic {
	ret := &Graphic{}
	ret.CT_GraphicalObject = *NewCT_GraphicalObject()
	return ret
}
func (m *Graphic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	return m.CT_GraphicalObject.MarshalXML(e, start)
}
func (m *Graphic) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_GraphicalObject = *NewCT_GraphicalObject()
lGraphic:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "graphicData":
				if err := d.DecodeElement(m.GraphicData, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lGraphic
		case xml.CharData:
		}
	}
	return nil
}
func (m *Graphic) Validate() error {
	return m.ValidateWithPath("Graphic")
}
func (m *Graphic) ValidateWithPath(path string) error {
	if err := m.CT_GraphicalObject.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}