// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"

	"baliance.com/gooxml"
)

type CT_ShapeDefaults struct {
	Any []gooxml.Any
}

func NewCT_ShapeDefaults() *CT_ShapeDefaults {
	ret := &CT_ShapeDefaults{}
	return ret
}
func (m *CT_ShapeDefaults) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Any != nil {
		for _, c := range m.Any {
			c.MarshalXML(e, start)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_ShapeDefaults) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ShapeDefaults:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			default:
				if anyEl, err := gooxml.CreateElement(el); err != nil {
					return err
				} else {
					if err := d.DecodeElement(anyEl, &el); err != nil {
						return err
					}
					m.Any = append(m.Any, anyEl)
				}
			}
		case xml.EndElement:
			break lCT_ShapeDefaults
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_ShapeDefaults) Validate() error {
	return m.ValidateWithPath("CT_ShapeDefaults")
}
func (m *CT_ShapeDefaults) ValidateWithPath(path string) error {
	return nil
}
