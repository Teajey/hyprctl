package hyprctl

import (
	"encoding/xml"
)

// Submit represents a control that initiates a Form submission.
//
// A single form may contain many submit options with different names and values.
//
// A submit name-value pair is mutually exclusive with those of any other Submit in the same form.
type Submit struct {
	Label string `json:"label,omitempty"`
	// FIXME: why have a submit button without a name? Empty names and labels should be included to disincentivise useless c:Submit
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

func (i Submit) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Local: "c:Submit"}

	if i.Name != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"}, Value: i.Name})
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "value"}, Value: i.Value})
	} else if i.Value != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "value"}, Value: i.Value})
	}

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	if err := e.EncodeToken(xml.CharData(i.Label)); err != nil {
		return err
	}

	if err := e.EncodeToken(start.End()); err != nil {
		return err
	}

	return nil
}
