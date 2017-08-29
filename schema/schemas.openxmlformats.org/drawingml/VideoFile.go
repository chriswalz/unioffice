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

type VideoFile struct {
	CT_VideoFile
}

func NewVideoFile() *VideoFile {
	ret := &VideoFile{}
	ret.CT_VideoFile = *NewCT_VideoFile()
	return ret
}
func (m *VideoFile) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "a:videoFile"
	return m.CT_VideoFile.MarshalXML(e, start)
}
func (m *VideoFile) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_VideoFile = *NewCT_VideoFile()
	for _, attr := range start.Attr {
		if attr.Name.Local == "link" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LinkAttr = parsed
		}
		if attr.Name.Local == "contentType" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ContentTypeAttr = &parsed
		}
	}
lVideoFile:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
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
			break lVideoFile
		case xml.CharData:
		}
	}
	return nil
}
func (m *VideoFile) Validate() error {
	return m.ValidateWithPath("VideoFile")
}
func (m *VideoFile) ValidateWithPath(path string) error {
	if err := m.CT_VideoFile.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}