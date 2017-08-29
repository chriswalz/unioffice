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

type CT_pivotTableDefinition struct {
	// Name
	NameAttr string
	// PivotCache Definition Id
	CacheIdAttr uint32
	// Data On Rows
	DataOnRowsAttr *bool
	// Default Data Field Position
	DataPositionAttr *uint32
	// Data Field Header Name
	DataCaptionAttr string
	// Grand Totals Caption
	GrandTotalCaptionAttr *string
	// Error Caption
	ErrorCaptionAttr *string
	// Show Error
	ShowErrorAttr *bool
	// Caption for Missing Values
	MissingCaptionAttr *string
	// Show Missing
	ShowMissingAttr *bool
	// Page Header Style Name
	PageStyleAttr *string
	// Table Style Name
	PivotTableStyleAttr *string
	// Vacated Style
	VacatedStyleAttr *string
	// PivotTable Custom String
	TagAttr *string
	// PivotTable Last Updated Version
	UpdatedVersionAttr *uint8
	// Minimum Refreshable Version
	MinRefreshableVersionAttr *uint8
	// Asterisk Totals
	AsteriskTotalsAttr *bool
	// Show Item Names
	ShowItemsAttr *bool
	// Allow Edit Data
	EditDataAttr *bool
	// Disable Field List
	DisableFieldListAttr *bool
	// Show Calculated Members
	ShowCalcMbrsAttr *bool
	// Total Visual Data
	VisualTotalsAttr *bool
	// Show Multiple Labels
	ShowMultipleLabelAttr *bool
	// Show Drop Down
	ShowDataDropDownAttr *bool
	// Show Expand Collapse
	ShowDrillAttr *bool
	// Print Drill Indicators
	PrintDrillAttr *bool
	// Show Member Property ToolTips
	ShowMemberPropertyTipsAttr *bool
	// Show ToolTips on Data
	ShowDataTipsAttr *bool
	// Enable PivotTable Wizard
	EnableWizardAttr *bool
	// Enable Drill Down
	EnableDrillAttr *bool
	// Enable Field Properties
	EnableFieldPropertiesAttr *bool
	// Preserve Formatting
	PreserveFormattingAttr *bool
	// Auto Formatting
	UseAutoFormattingAttr *bool
	// Page Wrap
	PageWrapAttr *uint32
	// Page Over Then Down
	PageOverThenDownAttr *bool
	// Subtotal Hidden Items
	SubtotalHiddenItemsAttr *bool
	// Row Grand Totals
	RowGrandTotalsAttr *bool
	// Grand Totals On Columns
	ColGrandTotalsAttr *bool
	// Field Print Titles
	FieldPrintTitlesAttr *bool
	// Item Print Titles
	ItemPrintTitlesAttr *bool
	// Merge Titles
	MergeItemAttr *bool
	// Show Drop Zones
	ShowDropZonesAttr *bool
	// PivotCache Created Version
	CreatedVersionAttr *uint8
	// Indentation for Compact Axis
	IndentAttr *uint32
	// Show Empty Row
	ShowEmptyRowAttr *bool
	// Show Empty Column
	ShowEmptyColAttr *bool
	// Show Field Headers
	ShowHeadersAttr *bool
	// Compact New Fields
	CompactAttr *bool
	// Outline New Fields
	OutlineAttr *bool
	// Outline Data Fields
	OutlineDataAttr *bool
	// Compact Data
	CompactDataAttr *bool
	// Data Fields Published
	PublishedAttr *bool
	// Enable Drop Zones
	GridDropZonesAttr *bool
	// Stop Immersive UI
	ImmersiveAttr *bool
	// Multiple Field Filters
	MultipleFieldFiltersAttr *bool
	// Chart Format Id
	ChartFormatAttr *uint32
	// Row Header Caption
	RowHeaderCaptionAttr *string
	// Column Header Caption
	ColHeaderCaptionAttr *string
	// Default Sort Order
	FieldListSortAscendingAttr *bool
	// MDX Subqueries Supported
	MdxSubqueriesAttr *bool
	// Custom List AutoSort
	CustomListSortAttr *bool
	// PivotTable Location
	Location *CT_Location
	// PivotTable Fields
	PivotFields *CT_PivotFields
	// Row Fields
	RowFields *CT_RowFields
	// Row Items
	RowItems *CT_rowItems
	// Column Fields
	ColFields *CT_ColFields
	// Column Items
	ColItems *CT_colItems
	// Page Field Items
	PageFields *CT_PageFields
	// Data Fields
	DataFields *CT_DataFields
	// PivotTable Formats
	Formats *CT_Formats
	// Conditional Formats
	ConditionalFormats *CT_ConditionalFormats
	// PivotChart Formats
	ChartFormats *CT_ChartFormats
	// PivotTable OLAP Hierarchies
	PivotHierarchies *CT_PivotHierarchies
	// PivotTable Style
	PivotTableStyleInfo *CT_PivotTableStyle
	// Filters
	Filters *CT_PivotFilters
	// Row OLAP Hierarchy References
	RowHierarchiesUsage *CT_RowHierarchiesUsage
	// Column OLAP Hierarchy References
	ColHierarchiesUsage *CT_ColHierarchiesUsage
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

func NewCT_pivotTableDefinition() *CT_pivotTableDefinition {
	ret := &CT_pivotTableDefinition{}
	ret.Location = NewCT_Location()
	return ret
}
func (m *CT_pivotTableDefinition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cacheId"},
		Value: fmt.Sprintf("%v", m.CacheIdAttr)})
	if m.DataOnRowsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dataOnRows"},
			Value: fmt.Sprintf("%v", *m.DataOnRowsAttr)})
	}
	if m.DataPositionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dataPosition"},
			Value: fmt.Sprintf("%v", *m.DataPositionAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dataCaption"},
		Value: fmt.Sprintf("%v", m.DataCaptionAttr)})
	if m.GrandTotalCaptionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "grandTotalCaption"},
			Value: fmt.Sprintf("%v", *m.GrandTotalCaptionAttr)})
	}
	if m.ErrorCaptionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "errorCaption"},
			Value: fmt.Sprintf("%v", *m.ErrorCaptionAttr)})
	}
	if m.ShowErrorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showError"},
			Value: fmt.Sprintf("%v", *m.ShowErrorAttr)})
	}
	if m.MissingCaptionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "missingCaption"},
			Value: fmt.Sprintf("%v", *m.MissingCaptionAttr)})
	}
	if m.ShowMissingAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showMissing"},
			Value: fmt.Sprintf("%v", *m.ShowMissingAttr)})
	}
	if m.PageStyleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "pageStyle"},
			Value: fmt.Sprintf("%v", *m.PageStyleAttr)})
	}
	if m.PivotTableStyleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "pivotTableStyle"},
			Value: fmt.Sprintf("%v", *m.PivotTableStyleAttr)})
	}
	if m.VacatedStyleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "vacatedStyle"},
			Value: fmt.Sprintf("%v", *m.VacatedStyleAttr)})
	}
	if m.TagAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "tag"},
			Value: fmt.Sprintf("%v", *m.TagAttr)})
	}
	if m.UpdatedVersionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "updatedVersion"},
			Value: fmt.Sprintf("%v", *m.UpdatedVersionAttr)})
	}
	if m.MinRefreshableVersionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "minRefreshableVersion"},
			Value: fmt.Sprintf("%v", *m.MinRefreshableVersionAttr)})
	}
	if m.AsteriskTotalsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "asteriskTotals"},
			Value: fmt.Sprintf("%v", *m.AsteriskTotalsAttr)})
	}
	if m.ShowItemsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showItems"},
			Value: fmt.Sprintf("%v", *m.ShowItemsAttr)})
	}
	if m.EditDataAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "editData"},
			Value: fmt.Sprintf("%v", *m.EditDataAttr)})
	}
	if m.DisableFieldListAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "disableFieldList"},
			Value: fmt.Sprintf("%v", *m.DisableFieldListAttr)})
	}
	if m.ShowCalcMbrsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showCalcMbrs"},
			Value: fmt.Sprintf("%v", *m.ShowCalcMbrsAttr)})
	}
	if m.VisualTotalsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "visualTotals"},
			Value: fmt.Sprintf("%v", *m.VisualTotalsAttr)})
	}
	if m.ShowMultipleLabelAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showMultipleLabel"},
			Value: fmt.Sprintf("%v", *m.ShowMultipleLabelAttr)})
	}
	if m.ShowDataDropDownAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showDataDropDown"},
			Value: fmt.Sprintf("%v", *m.ShowDataDropDownAttr)})
	}
	if m.ShowDrillAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showDrill"},
			Value: fmt.Sprintf("%v", *m.ShowDrillAttr)})
	}
	if m.PrintDrillAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "printDrill"},
			Value: fmt.Sprintf("%v", *m.PrintDrillAttr)})
	}
	if m.ShowMemberPropertyTipsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showMemberPropertyTips"},
			Value: fmt.Sprintf("%v", *m.ShowMemberPropertyTipsAttr)})
	}
	if m.ShowDataTipsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showDataTips"},
			Value: fmt.Sprintf("%v", *m.ShowDataTipsAttr)})
	}
	if m.EnableWizardAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "enableWizard"},
			Value: fmt.Sprintf("%v", *m.EnableWizardAttr)})
	}
	if m.EnableDrillAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "enableDrill"},
			Value: fmt.Sprintf("%v", *m.EnableDrillAttr)})
	}
	if m.EnableFieldPropertiesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "enableFieldProperties"},
			Value: fmt.Sprintf("%v", *m.EnableFieldPropertiesAttr)})
	}
	if m.PreserveFormattingAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "preserveFormatting"},
			Value: fmt.Sprintf("%v", *m.PreserveFormattingAttr)})
	}
	if m.UseAutoFormattingAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "useAutoFormatting"},
			Value: fmt.Sprintf("%v", *m.UseAutoFormattingAttr)})
	}
	if m.PageWrapAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "pageWrap"},
			Value: fmt.Sprintf("%v", *m.PageWrapAttr)})
	}
	if m.PageOverThenDownAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "pageOverThenDown"},
			Value: fmt.Sprintf("%v", *m.PageOverThenDownAttr)})
	}
	if m.SubtotalHiddenItemsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "subtotalHiddenItems"},
			Value: fmt.Sprintf("%v", *m.SubtotalHiddenItemsAttr)})
	}
	if m.RowGrandTotalsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rowGrandTotals"},
			Value: fmt.Sprintf("%v", *m.RowGrandTotalsAttr)})
	}
	if m.ColGrandTotalsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "colGrandTotals"},
			Value: fmt.Sprintf("%v", *m.ColGrandTotalsAttr)})
	}
	if m.FieldPrintTitlesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fieldPrintTitles"},
			Value: fmt.Sprintf("%v", *m.FieldPrintTitlesAttr)})
	}
	if m.ItemPrintTitlesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "itemPrintTitles"},
			Value: fmt.Sprintf("%v", *m.ItemPrintTitlesAttr)})
	}
	if m.MergeItemAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "mergeItem"},
			Value: fmt.Sprintf("%v", *m.MergeItemAttr)})
	}
	if m.ShowDropZonesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showDropZones"},
			Value: fmt.Sprintf("%v", *m.ShowDropZonesAttr)})
	}
	if m.CreatedVersionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "createdVersion"},
			Value: fmt.Sprintf("%v", *m.CreatedVersionAttr)})
	}
	if m.IndentAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "indent"},
			Value: fmt.Sprintf("%v", *m.IndentAttr)})
	}
	if m.ShowEmptyRowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showEmptyRow"},
			Value: fmt.Sprintf("%v", *m.ShowEmptyRowAttr)})
	}
	if m.ShowEmptyColAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showEmptyCol"},
			Value: fmt.Sprintf("%v", *m.ShowEmptyColAttr)})
	}
	if m.ShowHeadersAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showHeaders"},
			Value: fmt.Sprintf("%v", *m.ShowHeadersAttr)})
	}
	if m.CompactAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "compact"},
			Value: fmt.Sprintf("%v", *m.CompactAttr)})
	}
	if m.OutlineAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "outline"},
			Value: fmt.Sprintf("%v", *m.OutlineAttr)})
	}
	if m.OutlineDataAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "outlineData"},
			Value: fmt.Sprintf("%v", *m.OutlineDataAttr)})
	}
	if m.CompactDataAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "compactData"},
			Value: fmt.Sprintf("%v", *m.CompactDataAttr)})
	}
	if m.PublishedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "published"},
			Value: fmt.Sprintf("%v", *m.PublishedAttr)})
	}
	if m.GridDropZonesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "gridDropZones"},
			Value: fmt.Sprintf("%v", *m.GridDropZonesAttr)})
	}
	if m.ImmersiveAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "immersive"},
			Value: fmt.Sprintf("%v", *m.ImmersiveAttr)})
	}
	if m.MultipleFieldFiltersAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "multipleFieldFilters"},
			Value: fmt.Sprintf("%v", *m.MultipleFieldFiltersAttr)})
	}
	if m.ChartFormatAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "chartFormat"},
			Value: fmt.Sprintf("%v", *m.ChartFormatAttr)})
	}
	if m.RowHeaderCaptionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rowHeaderCaption"},
			Value: fmt.Sprintf("%v", *m.RowHeaderCaptionAttr)})
	}
	if m.ColHeaderCaptionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "colHeaderCaption"},
			Value: fmt.Sprintf("%v", *m.ColHeaderCaptionAttr)})
	}
	if m.FieldListSortAscendingAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fieldListSortAscending"},
			Value: fmt.Sprintf("%v", *m.FieldListSortAscendingAttr)})
	}
	if m.MdxSubqueriesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "mdxSubqueries"},
			Value: fmt.Sprintf("%v", *m.MdxSubqueriesAttr)})
	}
	if m.CustomListSortAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "customListSort"},
			Value: fmt.Sprintf("%v", *m.CustomListSortAttr)})
	}
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
	selocation := xml.StartElement{Name: xml.Name{Local: "x:location"}}
	e.EncodeElement(m.Location, selocation)
	if m.PivotFields != nil {
		sepivotFields := xml.StartElement{Name: xml.Name{Local: "x:pivotFields"}}
		e.EncodeElement(m.PivotFields, sepivotFields)
	}
	if m.RowFields != nil {
		serowFields := xml.StartElement{Name: xml.Name{Local: "x:rowFields"}}
		e.EncodeElement(m.RowFields, serowFields)
	}
	if m.RowItems != nil {
		serowItems := xml.StartElement{Name: xml.Name{Local: "x:rowItems"}}
		e.EncodeElement(m.RowItems, serowItems)
	}
	if m.ColFields != nil {
		secolFields := xml.StartElement{Name: xml.Name{Local: "x:colFields"}}
		e.EncodeElement(m.ColFields, secolFields)
	}
	if m.ColItems != nil {
		secolItems := xml.StartElement{Name: xml.Name{Local: "x:colItems"}}
		e.EncodeElement(m.ColItems, secolItems)
	}
	if m.PageFields != nil {
		sepageFields := xml.StartElement{Name: xml.Name{Local: "x:pageFields"}}
		e.EncodeElement(m.PageFields, sepageFields)
	}
	if m.DataFields != nil {
		sedataFields := xml.StartElement{Name: xml.Name{Local: "x:dataFields"}}
		e.EncodeElement(m.DataFields, sedataFields)
	}
	if m.Formats != nil {
		seformats := xml.StartElement{Name: xml.Name{Local: "x:formats"}}
		e.EncodeElement(m.Formats, seformats)
	}
	if m.ConditionalFormats != nil {
		seconditionalFormats := xml.StartElement{Name: xml.Name{Local: "x:conditionalFormats"}}
		e.EncodeElement(m.ConditionalFormats, seconditionalFormats)
	}
	if m.ChartFormats != nil {
		sechartFormats := xml.StartElement{Name: xml.Name{Local: "x:chartFormats"}}
		e.EncodeElement(m.ChartFormats, sechartFormats)
	}
	if m.PivotHierarchies != nil {
		sepivotHierarchies := xml.StartElement{Name: xml.Name{Local: "x:pivotHierarchies"}}
		e.EncodeElement(m.PivotHierarchies, sepivotHierarchies)
	}
	if m.PivotTableStyleInfo != nil {
		sepivotTableStyleInfo := xml.StartElement{Name: xml.Name{Local: "x:pivotTableStyleInfo"}}
		e.EncodeElement(m.PivotTableStyleInfo, sepivotTableStyleInfo)
	}
	if m.Filters != nil {
		sefilters := xml.StartElement{Name: xml.Name{Local: "x:filters"}}
		e.EncodeElement(m.Filters, sefilters)
	}
	if m.RowHierarchiesUsage != nil {
		serowHierarchiesUsage := xml.StartElement{Name: xml.Name{Local: "x:rowHierarchiesUsage"}}
		e.EncodeElement(m.RowHierarchiesUsage, serowHierarchiesUsage)
	}
	if m.ColHierarchiesUsage != nil {
		secolHierarchiesUsage := xml.StartElement{Name: xml.Name{Local: "x:colHierarchiesUsage"}}
		e.EncodeElement(m.ColHierarchiesUsage, secolHierarchiesUsage)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_pivotTableDefinition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Location = NewCT_Location()
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
		}
		if attr.Name.Local == "cacheId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.CacheIdAttr = uint32(parsed)
		}
		if attr.Name.Local == "dataOnRows" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DataOnRowsAttr = &parsed
		}
		if attr.Name.Local == "dataPosition" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.DataPositionAttr = &pt
		}
		if attr.Name.Local == "dataCaption" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DataCaptionAttr = parsed
		}
		if attr.Name.Local == "grandTotalCaption" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.GrandTotalCaptionAttr = &parsed
		}
		if attr.Name.Local == "errorCaption" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ErrorCaptionAttr = &parsed
		}
		if attr.Name.Local == "showError" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowErrorAttr = &parsed
		}
		if attr.Name.Local == "missingCaption" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.MissingCaptionAttr = &parsed
		}
		if attr.Name.Local == "showMissing" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowMissingAttr = &parsed
		}
		if attr.Name.Local == "pageStyle" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PageStyleAttr = &parsed
		}
		if attr.Name.Local == "pivotTableStyle" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PivotTableStyleAttr = &parsed
		}
		if attr.Name.Local == "vacatedStyle" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.VacatedStyleAttr = &parsed
		}
		if attr.Name.Local == "tag" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TagAttr = &parsed
		}
		if attr.Name.Local == "updatedVersion" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint8(parsed)
			m.UpdatedVersionAttr = &pt
		}
		if attr.Name.Local == "minRefreshableVersion" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint8(parsed)
			m.MinRefreshableVersionAttr = &pt
		}
		if attr.Name.Local == "asteriskTotals" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AsteriskTotalsAttr = &parsed
		}
		if attr.Name.Local == "showItems" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowItemsAttr = &parsed
		}
		if attr.Name.Local == "editData" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.EditDataAttr = &parsed
		}
		if attr.Name.Local == "disableFieldList" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DisableFieldListAttr = &parsed
		}
		if attr.Name.Local == "showCalcMbrs" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowCalcMbrsAttr = &parsed
		}
		if attr.Name.Local == "visualTotals" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.VisualTotalsAttr = &parsed
		}
		if attr.Name.Local == "showMultipleLabel" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowMultipleLabelAttr = &parsed
		}
		if attr.Name.Local == "showDataDropDown" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowDataDropDownAttr = &parsed
		}
		if attr.Name.Local == "showDrill" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowDrillAttr = &parsed
		}
		if attr.Name.Local == "printDrill" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PrintDrillAttr = &parsed
		}
		if attr.Name.Local == "showMemberPropertyTips" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowMemberPropertyTipsAttr = &parsed
		}
		if attr.Name.Local == "showDataTips" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowDataTipsAttr = &parsed
		}
		if attr.Name.Local == "enableWizard" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.EnableWizardAttr = &parsed
		}
		if attr.Name.Local == "enableDrill" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.EnableDrillAttr = &parsed
		}
		if attr.Name.Local == "enableFieldProperties" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.EnableFieldPropertiesAttr = &parsed
		}
		if attr.Name.Local == "preserveFormatting" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PreserveFormattingAttr = &parsed
		}
		if attr.Name.Local == "useAutoFormatting" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UseAutoFormattingAttr = &parsed
		}
		if attr.Name.Local == "pageWrap" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.PageWrapAttr = &pt
		}
		if attr.Name.Local == "pageOverThenDown" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PageOverThenDownAttr = &parsed
		}
		if attr.Name.Local == "subtotalHiddenItems" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SubtotalHiddenItemsAttr = &parsed
		}
		if attr.Name.Local == "rowGrandTotals" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RowGrandTotalsAttr = &parsed
		}
		if attr.Name.Local == "colGrandTotals" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ColGrandTotalsAttr = &parsed
		}
		if attr.Name.Local == "fieldPrintTitles" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FieldPrintTitlesAttr = &parsed
		}
		if attr.Name.Local == "itemPrintTitles" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ItemPrintTitlesAttr = &parsed
		}
		if attr.Name.Local == "mergeItem" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.MergeItemAttr = &parsed
		}
		if attr.Name.Local == "showDropZones" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowDropZonesAttr = &parsed
		}
		if attr.Name.Local == "createdVersion" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint8(parsed)
			m.CreatedVersionAttr = &pt
		}
		if attr.Name.Local == "indent" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.IndentAttr = &pt
		}
		if attr.Name.Local == "showEmptyRow" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowEmptyRowAttr = &parsed
		}
		if attr.Name.Local == "showEmptyCol" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowEmptyColAttr = &parsed
		}
		if attr.Name.Local == "showHeaders" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowHeadersAttr = &parsed
		}
		if attr.Name.Local == "compact" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CompactAttr = &parsed
		}
		if attr.Name.Local == "outline" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.OutlineAttr = &parsed
		}
		if attr.Name.Local == "outlineData" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.OutlineDataAttr = &parsed
		}
		if attr.Name.Local == "compactData" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CompactDataAttr = &parsed
		}
		if attr.Name.Local == "published" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PublishedAttr = &parsed
		}
		if attr.Name.Local == "gridDropZones" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.GridDropZonesAttr = &parsed
		}
		if attr.Name.Local == "immersive" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ImmersiveAttr = &parsed
		}
		if attr.Name.Local == "multipleFieldFilters" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.MultipleFieldFiltersAttr = &parsed
		}
		if attr.Name.Local == "chartFormat" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.ChartFormatAttr = &pt
		}
		if attr.Name.Local == "rowHeaderCaption" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RowHeaderCaptionAttr = &parsed
		}
		if attr.Name.Local == "colHeaderCaption" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ColHeaderCaptionAttr = &parsed
		}
		if attr.Name.Local == "fieldListSortAscending" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FieldListSortAscendingAttr = &parsed
		}
		if attr.Name.Local == "mdxSubqueries" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.MdxSubqueriesAttr = &parsed
		}
		if attr.Name.Local == "customListSort" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CustomListSortAttr = &parsed
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
lCT_pivotTableDefinition:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "location":
				if err := d.DecodeElement(m.Location, &el); err != nil {
					return err
				}
			case "pivotFields":
				m.PivotFields = NewCT_PivotFields()
				if err := d.DecodeElement(m.PivotFields, &el); err != nil {
					return err
				}
			case "rowFields":
				m.RowFields = NewCT_RowFields()
				if err := d.DecodeElement(m.RowFields, &el); err != nil {
					return err
				}
			case "rowItems":
				m.RowItems = NewCT_rowItems()
				if err := d.DecodeElement(m.RowItems, &el); err != nil {
					return err
				}
			case "colFields":
				m.ColFields = NewCT_ColFields()
				if err := d.DecodeElement(m.ColFields, &el); err != nil {
					return err
				}
			case "colItems":
				m.ColItems = NewCT_colItems()
				if err := d.DecodeElement(m.ColItems, &el); err != nil {
					return err
				}
			case "pageFields":
				m.PageFields = NewCT_PageFields()
				if err := d.DecodeElement(m.PageFields, &el); err != nil {
					return err
				}
			case "dataFields":
				m.DataFields = NewCT_DataFields()
				if err := d.DecodeElement(m.DataFields, &el); err != nil {
					return err
				}
			case "formats":
				m.Formats = NewCT_Formats()
				if err := d.DecodeElement(m.Formats, &el); err != nil {
					return err
				}
			case "conditionalFormats":
				m.ConditionalFormats = NewCT_ConditionalFormats()
				if err := d.DecodeElement(m.ConditionalFormats, &el); err != nil {
					return err
				}
			case "chartFormats":
				m.ChartFormats = NewCT_ChartFormats()
				if err := d.DecodeElement(m.ChartFormats, &el); err != nil {
					return err
				}
			case "pivotHierarchies":
				m.PivotHierarchies = NewCT_PivotHierarchies()
				if err := d.DecodeElement(m.PivotHierarchies, &el); err != nil {
					return err
				}
			case "pivotTableStyleInfo":
				m.PivotTableStyleInfo = NewCT_PivotTableStyle()
				if err := d.DecodeElement(m.PivotTableStyleInfo, &el); err != nil {
					return err
				}
			case "filters":
				m.Filters = NewCT_PivotFilters()
				if err := d.DecodeElement(m.Filters, &el); err != nil {
					return err
				}
			case "rowHierarchiesUsage":
				m.RowHierarchiesUsage = NewCT_RowHierarchiesUsage()
				if err := d.DecodeElement(m.RowHierarchiesUsage, &el); err != nil {
					return err
				}
			case "colHierarchiesUsage":
				m.ColHierarchiesUsage = NewCT_ColHierarchiesUsage()
				if err := d.DecodeElement(m.ColHierarchiesUsage, &el); err != nil {
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
			break lCT_pivotTableDefinition
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_pivotTableDefinition) Validate() error {
	return m.ValidateWithPath("CT_pivotTableDefinition")
}
func (m *CT_pivotTableDefinition) ValidateWithPath(path string) error {
	if err := m.Location.ValidateWithPath(path + "/Location"); err != nil {
		return err
	}
	if m.PivotFields != nil {
		if err := m.PivotFields.ValidateWithPath(path + "/PivotFields"); err != nil {
			return err
		}
	}
	if m.RowFields != nil {
		if err := m.RowFields.ValidateWithPath(path + "/RowFields"); err != nil {
			return err
		}
	}
	if m.RowItems != nil {
		if err := m.RowItems.ValidateWithPath(path + "/RowItems"); err != nil {
			return err
		}
	}
	if m.ColFields != nil {
		if err := m.ColFields.ValidateWithPath(path + "/ColFields"); err != nil {
			return err
		}
	}
	if m.ColItems != nil {
		if err := m.ColItems.ValidateWithPath(path + "/ColItems"); err != nil {
			return err
		}
	}
	if m.PageFields != nil {
		if err := m.PageFields.ValidateWithPath(path + "/PageFields"); err != nil {
			return err
		}
	}
	if m.DataFields != nil {
		if err := m.DataFields.ValidateWithPath(path + "/DataFields"); err != nil {
			return err
		}
	}
	if m.Formats != nil {
		if err := m.Formats.ValidateWithPath(path + "/Formats"); err != nil {
			return err
		}
	}
	if m.ConditionalFormats != nil {
		if err := m.ConditionalFormats.ValidateWithPath(path + "/ConditionalFormats"); err != nil {
			return err
		}
	}
	if m.ChartFormats != nil {
		if err := m.ChartFormats.ValidateWithPath(path + "/ChartFormats"); err != nil {
			return err
		}
	}
	if m.PivotHierarchies != nil {
		if err := m.PivotHierarchies.ValidateWithPath(path + "/PivotHierarchies"); err != nil {
			return err
		}
	}
	if m.PivotTableStyleInfo != nil {
		if err := m.PivotTableStyleInfo.ValidateWithPath(path + "/PivotTableStyleInfo"); err != nil {
			return err
		}
	}
	if m.Filters != nil {
		if err := m.Filters.ValidateWithPath(path + "/Filters"); err != nil {
			return err
		}
	}
	if m.RowHierarchiesUsage != nil {
		if err := m.RowHierarchiesUsage.ValidateWithPath(path + "/RowHierarchiesUsage"); err != nil {
			return err
		}
	}
	if m.ColHierarchiesUsage != nil {
		if err := m.ColHierarchiesUsage.ValidateWithPath(path + "/ColHierarchiesUsage"); err != nil {
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