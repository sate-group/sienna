package sienna

import "encoding/json"

type State struct {
	payload string
}

func NewState(p string) *State {
	return &State{
		payload: p,
	}
}

func (p *State) Decode(v any) error {
	err := json.Unmarshal([]byte(p.payload), &v)
	return err
}
