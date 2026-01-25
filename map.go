package hyprctl

import (
	"encoding/xml"
	"net/url"
	"sort"
	"strings"
)

type Map struct {
	Label   string              `json:"label"`
	Name    string              `json:"name"`
	Entries map[string][]string `json:"entries"`
}

func (m *Map) ExtractFormValues(form url.Values) {
	if m.Entries == nil {
		m.Entries = make(map[string][]string, len(form))
	}
	for k, v := range form {
		if m.Name != "" && !strings.HasPrefix(k, m.Name+"[") && !strings.HasSuffix(k, "]") {
			continue
		}
		delete(form, k)
		m.Entries[k] = v
	}
}

func (m Map) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "c:Map"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "label"}, Value: m.Label})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"}, Value: m.Name})

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	keys := make([]string, 0, len(m.Entries))
	for k := range m.Entries {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		for _, v := range m.Entries[k] {
			keyElem := xml.StartElement{Name: xml.Name{Local: "c:Input"}}
			keyElem.Attr = append(keyElem.Attr, xml.Attr{Name: xml.Name{Local: "name"}, Value: k})
			keyElem.Attr = append(keyElem.Attr, xml.Attr{Name: xml.Name{Local: "value"}, Value: v})
			if err := e.EncodeElement("", keyElem); err != nil {
				return err
			}
		}
	}

	return e.EncodeToken(start.End())
}
