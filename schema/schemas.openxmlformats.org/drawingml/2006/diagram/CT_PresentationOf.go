// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"encoding/xml"
	"fmt"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_PresentationOf struct {
	ExtLst            *drawingml.CT_OfficeArtExtensionList
	AxisAttr          *ST_AxisTypes
	PtTypeAttr        *ST_ElementTypes
	HideLastTransAttr *ST_Booleans
	StAttr            *ST_Ints
	CntAttr           *ST_UnsignedInts
	StepAttr          *ST_Ints
}

func NewCT_PresentationOf() *CT_PresentationOf {
	ret := &CT_PresentationOf{}
	return ret
}
func (m *CT_PresentationOf) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.AxisAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "axis"},
			Value: fmt.Sprintf("%v", *m.AxisAttr)})
	}
	if m.PtTypeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ptType"},
			Value: fmt.Sprintf("%v", *m.PtTypeAttr)})
	}
	if m.HideLastTransAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hideLastTrans"},
			Value: fmt.Sprintf("%v", *m.HideLastTransAttr)})
	}
	if m.StAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "st"},
			Value: fmt.Sprintf("%v", *m.StAttr)})
	}
	if m.CntAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cnt"},
			Value: fmt.Sprintf("%v", *m.CntAttr)})
	}
	if m.StepAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "step"},
			Value: fmt.Sprintf("%v", *m.StepAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_PresentationOf) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "axis" {
			parsed, err := ParseSliceST_AxisTypes(attr.Value)
			if err != nil {
				return err
			}
			m.AxisAttr = &parsed
		}
		if attr.Name.Local == "ptType" {
			parsed, err := ParseSliceST_ElementTypes(attr.Value)
			if err != nil {
				return err
			}
			m.PtTypeAttr = &parsed
		}
		if attr.Name.Local == "hideLastTrans" {
			parsed, err := ParseSliceST_Booleans(attr.Value)
			if err != nil {
				return err
			}
			m.HideLastTransAttr = &parsed
		}
		if attr.Name.Local == "st" {
			parsed, err := ParseSliceST_Ints(attr.Value)
			if err != nil {
				return err
			}
			m.StAttr = &parsed
		}
		if attr.Name.Local == "cnt" {
			parsed, err := ParseSliceST_UnsignedInts(attr.Value)
			if err != nil {
				return err
			}
			m.CntAttr = &parsed
		}
		if attr.Name.Local == "step" {
			parsed, err := ParseSliceST_Ints(attr.Value)
			if err != nil {
				return err
			}
			m.StepAttr = &parsed
		}
	}
lCT_PresentationOf:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "extLst":
				m.ExtLst = drawingml.NewCT_OfficeArtExtensionList()
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
			break lCT_PresentationOf
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_PresentationOf) Validate() error {
	return m.ValidateWithPath("CT_PresentationOf")
}
func (m *CT_PresentationOf) ValidateWithPath(path string) error {
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}