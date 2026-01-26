package hyprctl

import (
	"encoding/xml"
	"net/url"
)

// Form is analogous to HTML's <form> which represents a state transition that requires input from the client.
// It describes what data is needed, how it should be submitted, and where
// it should be sent.
//
// Elements should contain a struct representing,
// and elements semantically related to the form,
// which would usually be [Input], [Select], [Map], [Link], etc.
// but might also be something like `Error string` or `Warning string` fields.
type Form struct {
	Method       string `json:"method,omitempty"`
	FormElements `json:"elements"`
}

func (i Form) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Local: "c:Form"}

	if i.Method != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "method"}, Value: i.Method})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	err = e.Encode(i.FormElements)
	if err != nil {
		return err
	}

	return e.EncodeToken(start.End())
}

type FormElements interface {
	ExtractValues(form url.Values)
	Validate()
}
