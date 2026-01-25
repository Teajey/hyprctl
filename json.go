package hyprctl

import "encoding/json"

type inputJson struct {
	Label     string   `json:"label"`
	Type      string   `json:"type,omitempty"`
	Name      string   `json:"name"`
	Required  bool     `json:"required,omitempty"`
	Value     *string  `json:"value,omitempty"`
	MinLength uint     `json:"minlength,omitempty"`
	MaxLength uint     `json:"maxlength,omitempty"`
	Step      float32  `json:"step,omitempty"`
	Min       string   `json:"min,omitempty"`
	Max       string   `json:"max,omitempty"`
	Error     string   `json:"error,omitempty"`
	Multiple  bool     `json:"multiple,omitempty"`
	Options   []Option `json:"options,omitempty"`
}

func (i Input) MarshalJSON() ([]byte, error) {
	vals := i.Values()
	j := inputJson{
		Label:     i.Label,
		Type:      i.Type,
		Name:      i.Name,
		Required:  i.Required,
		MinLength: i.MinLength,
		MaxLength: i.MaxLength,
		Step:      i.Step,
		Min:       i.Min,
		Max:       i.Max,
		Error:     i.Error,
		Multiple:  i.Multiple,
		Options:   i.Options,
	}
	if len(vals) < 2 {
		j.Value = &vals[0]
	}
	return json.Marshal(j)
}
