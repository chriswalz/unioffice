// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_QueryTable struct {
	// QueryTable Name
	NameAttr string
	// First Row Column Titles
	HeadersAttr *bool
	// Row Numbers
	RowNumbersAttr *bool
	// Disable Refresh
	DisableRefreshAttr *bool
	// Background Refresh
	BackgroundRefreshAttr *bool
	// First Background Refresh
	FirstBackgroundRefreshAttr *bool
	// Refresh On Load
	RefreshOnLoadAttr *bool
	// Grow Shrink Type
	GrowShrinkTypeAttr ST_GrowShrinkType
	// Fill Adjacent Formulas
	FillFormulasAttr *bool
	// Remove Data On Save
	RemoveDataOnSaveAttr *bool
	// Disable Edit
	DisableEditAttr *bool
	// Preserve Formatting On Refresh
	PreserveFormattingAttr *bool
	// Adjust Column Width On Refresh
	AdjustColumnWidthAttr *bool
	// Intermediate
	IntermediateAttr *bool
	// Connection Id
	ConnectionIdAttr uint32
	// QueryTable Refresh Information
	QueryTableRefresh *CT_QueryTableRefresh
	// Future Feature Data Storage Area
	ExtLst                      *CT_ExtensionList
	AutoFormatIdAttr            *uint32
	ApplyNumberFormatsAttr      *bool
	ApplyBorderFormatsAttr      *bool
	ApplyFontFormatsAttr        *bool
	ApplyPatternFormatsAttr     *bool
	ApplyAlignmentFormatsAttr   *bool
	ApplyWidthHeightFormatsAttr *bool
}

func NewCT_QueryTable() *CT_QueryTable {
	ret := &CT_QueryTable{}
	return ret
}
func (m *CT_QueryTable) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	if m.HeadersAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "headers"},
			Value: fmt.Sprintf("%v", *m.HeadersAttr)})
	}
	if m.RowNumbersAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rowNumbers"},
			Value: fmt.Sprintf("%v", *m.RowNumbersAttr)})
	}
	if m.DisableRefreshAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "disableRefresh"},
			Value: fmt.Sprintf("%v", *m.DisableRefreshAttr)})
	}
	if m.BackgroundRefreshAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "backgroundRefresh"},
			Value: fmt.Sprintf("%v", *m.BackgroundRefreshAttr)})
	}
	if m.FirstBackgroundRefreshAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "firstBackgroundRefresh"},
			Value: fmt.Sprintf("%v", *m.FirstBackgroundRefreshAttr)})
	}
	if m.RefreshOnLoadAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "refreshOnLoad"},
			Value: fmt.Sprintf("%v", *m.RefreshOnLoadAttr)})
	}
	if m.GrowShrinkTypeAttr != ST_GrowShrinkTypeUnset {
		attr, err := m.GrowShrinkTypeAttr.MarshalXMLAttr(xml.Name{Local: "growShrinkType"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.FillFormulasAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fillFormulas"},
			Value: fmt.Sprintf("%v", *m.FillFormulasAttr)})
	}
	if m.RemoveDataOnSaveAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "removeDataOnSave"},
			Value: fmt.Sprintf("%v", *m.RemoveDataOnSaveAttr)})
	}
	if m.DisableEditAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "disableEdit"},
			Value: fmt.Sprintf("%v", *m.DisableEditAttr)})
	}
	if m.PreserveFormattingAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "preserveFormatting"},
			Value: fmt.Sprintf("%v", *m.PreserveFormattingAttr)})
	}
	if m.AdjustColumnWidthAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "adjustColumnWidth"},
			Value: fmt.Sprintf("%v", *m.AdjustColumnWidthAttr)})
	}
	if m.IntermediateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "intermediate"},
			Value: fmt.Sprintf("%v", *m.IntermediateAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "connectionId"},
		Value: fmt.Sprintf("%v", m.ConnectionIdAttr)})
	if m.AutoFormatIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "autoFormatId"},
			Value: fmt.Sprintf("%v", *m.AutoFormatIdAttr)})
	}
	if m.ApplyNumberFormatsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyNumberFormats"},
			Value: fmt.Sprintf("%v", *m.ApplyNumberFormatsAttr)})
	}
	if m.ApplyBorderFormatsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyBorderFormats"},
			Value: fmt.Sprintf("%v", *m.ApplyBorderFormatsAttr)})
	}
	if m.ApplyFontFormatsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyFontFormats"},
			Value: fmt.Sprintf("%v", *m.ApplyFontFormatsAttr)})
	}
	if m.ApplyPatternFormatsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyPatternFormats"},
			Value: fmt.Sprintf("%v", *m.ApplyPatternFormatsAttr)})
	}
	if m.ApplyAlignmentFormatsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyAlignmentFormats"},
			Value: fmt.Sprintf("%v", *m.ApplyAlignmentFormatsAttr)})
	}
	if m.ApplyWidthHeightFormatsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyWidthHeightFormats"},
			Value: fmt.Sprintf("%v", *m.ApplyWidthHeightFormatsAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.QueryTableRefresh != nil {
		sequeryTableRefresh := xml.StartElement{Name: xml.Name{Local: "x:queryTableRefresh"}}
		e.EncodeElement(m.QueryTableRefresh, sequeryTableRefresh)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_QueryTable) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
		}
		if attr.Name.Local == "headers" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HeadersAttr = &parsed
		}
		if attr.Name.Local == "rowNumbers" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RowNumbersAttr = &parsed
		}
		if attr.Name.Local == "disableRefresh" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DisableRefreshAttr = &parsed
		}
		if attr.Name.Local == "backgroundRefresh" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.BackgroundRefreshAttr = &parsed
		}
		if attr.Name.Local == "firstBackgroundRefresh" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FirstBackgroundRefreshAttr = &parsed
		}
		if attr.Name.Local == "refreshOnLoad" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RefreshOnLoadAttr = &parsed
		}
		if attr.Name.Local == "growShrinkType" {
			m.GrowShrinkTypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "fillFormulas" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FillFormulasAttr = &parsed
		}
		if attr.Name.Local == "removeDataOnSave" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RemoveDataOnSaveAttr = &parsed
		}
		if attr.Name.Local == "disableEdit" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DisableEditAttr = &parsed
		}
		if attr.Name.Local == "preserveFormatting" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PreserveFormattingAttr = &parsed
		}
		if attr.Name.Local == "adjustColumnWidth" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AdjustColumnWidthAttr = &parsed
		}
		if attr.Name.Local == "intermediate" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.IntermediateAttr = &parsed
		}
		if attr.Name.Local == "connectionId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.ConnectionIdAttr = uint32(parsed)
		}
		if attr.Name.Local == "autoFormatId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.AutoFormatIdAttr = &pt
		}
		if attr.Name.Local == "applyNumberFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyNumberFormatsAttr = &parsed
		}
		if attr.Name.Local == "applyBorderFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyBorderFormatsAttr = &parsed
		}
		if attr.Name.Local == "applyFontFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyFontFormatsAttr = &parsed
		}
		if attr.Name.Local == "applyPatternFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyPatternFormatsAttr = &parsed
		}
		if attr.Name.Local == "applyAlignmentFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyAlignmentFormatsAttr = &parsed
		}
		if attr.Name.Local == "applyWidthHeightFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyWidthHeightFormatsAttr = &parsed
		}
	}
lCT_QueryTable:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "queryTableRefresh":
				m.QueryTableRefresh = NewCT_QueryTableRefresh()
				if err := d.DecodeElement(m.QueryTableRefresh, &el); err != nil {
					return err
				}
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
			break lCT_QueryTable
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_QueryTable) Validate() error {
	return m.ValidateWithPath("CT_QueryTable")
}
func (m *CT_QueryTable) ValidateWithPath(path string) error {
	if err := m.GrowShrinkTypeAttr.ValidateWithPath(path + "/GrowShrinkTypeAttr"); err != nil {
		return err
	}
	if m.QueryTableRefresh != nil {
		if err := m.QueryTableRefresh.ValidateWithPath(path + "/QueryTableRefresh"); err != nil {
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