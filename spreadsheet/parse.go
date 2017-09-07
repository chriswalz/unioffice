// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseCellReference parses a cell reference of the form 'A10' and splits it
// into column/row segments.
func ParseCellReference(s string) (col string, row uint32, err error) {
	split := -1
lfor:
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] >= '0' && s[i] <= '9':
			split = i
			break lfor
		}
	}
	switch split {
	case 0:
		return col, row, fmt.Errorf("no letter prefix in %s", s)
	case -1:
		return col, row, fmt.Errorf("no digits in %s", s)
	}

	col = s[0:split]
	r64, err := strconv.ParseUint(s[split:], 10, 32)
	row = uint32(r64)
	return col, row, err
}

// ColumnToIndex maps a column to a zero based index (e.g. A = 0, B = 1, AA = 26)
func ColumnToIndex(col string) uint32 {
	col = strings.ToUpper(col)
	res := uint32(0)
	for _, c := range col {
		res *= 26
		res += uint32(c - 'A' + 1)
	}
	return res - 1
}

// IndexToColumn maps a column number to a coumn name (e.g. 0 = A, 1 = B, 26 = AA)
func IndexToColumn(col uint32) string {
	col++
	res := []byte{}
	for col > 26 {
		res = append(res, byte('A'+(col-1)%26))
		col /= 26
	}
	res = append(res, byte('A'+(col-1)%26))
	// reverse it
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return string(res)
}