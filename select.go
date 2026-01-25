package hyprctl

import (
	"cmp"
	"encoding/xml"
	"iter"
	"net/url"
)

type Option struct {
	Value    string
	Label    string `json:",omitempty"`
	Selected bool   `json:",omitempty"`
	Disabled bool   `json:",omitempty"`
}

func (o Option) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "c:Option"
	if o.Selected && o.Value != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "selected"}})
	}
	if o.Disabled {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "disabled"}})
	}
	if o.Label != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "value"}, Value: o.Value})
		if err := e.EncodeToken(start); err != nil {
			return err
		}
		if err := e.EncodeToken(xml.CharData(o.Label)); err != nil {
			return err
		}
		if err := e.EncodeToken(start.End()); err != nil {
			return err
		}
	} else {
		if err := e.EncodeToken(start); err != nil {
			return err
		}
		if err := e.EncodeToken(xml.CharData(o.Value)); err != nil {
			return err
		}
		if err := e.EncodeToken(start.End()); err != nil {
			return err
		}
	}
	return nil
}

type Select struct {
	Multiple bool     `json:"multiple,omitempty"`
	Label    string   `json:"label"`
	Name     string   `json:"name"`
	Required bool     `json:"required,omitempty"`
	Options  []Option `json:"options"`
}

func (s *Select) SetValues(values ...string) {
	for _, v := range values {
		found := false
		for i, o := range s.Options {
			if o.Value == v {
				s.Options[i].Selected = true
				found = true
			} else {
				s.Options[i].Selected = false
			}
		}
		if !found {
			s.Options = append([]Option{{
				Value:    v,
				Selected: true,
			}}, s.Options...)
		}
	}
}

func (s *Select) Values() iter.Seq[string] {
	return iter.Seq[string](func(yield func(string) bool) {
		for _, o := range s.Options {
			if o.Selected {
				if !yield(o.Value) {
					return
				}
			}
		}
	})
}

func (s *Select) ExtractFormValues(form url.Values) {
	formValue, ok := form[s.Name]
	if !ok {
		return
	}
	if s.Multiple {
		s.SetValues(formValue...)
		delete(form, s.Name)
	} else {
		s.SetValues(formValue[0])
		form[s.Name] = formValue[1:]
	}
}

func (i Select) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "c:Select"

	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "label"}, Value: i.Label})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"}, Value: i.Name})
	if i.Required {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "required"}, Value: "true"})
	}

	if err := e.EncodeToken(start); err != nil {
		return nil
	}

	for _, o := range i.Options {
		optionElement := xml.StartElement{Name: xml.Name{Local: "c:Option"}}
		label := cmp.Or(o.Label, o.Value)
		if o.Label != "" {
			optionElement.Attr = append(optionElement.Attr, xml.Attr{Name: xml.Name{Local: "value"}, Value: o.Value})
		}
		if err := e.EncodeElement(label, optionElement); err != nil {
			return err
		}
	}

	return e.EncodeToken(start.End())
}
