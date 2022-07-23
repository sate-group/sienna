package sienna

import "encoding/json"

func jsonToStr(v any) (string, error) {
	out, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	str := string(out)
	return str, nil
}

func strToJson(v any, str string) error {
	if err := json.Unmarshal([]byte(str), &v); err != nil {
		return err
	}
	return nil
}
