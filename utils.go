package sienna

import (
	"encoding/json"

	"github.com/sate-infra/sienna/errs"
)

func jsonToStr(v any) (string, error) {
	out, err := json.Marshal(v)
	if err != nil {
		return "", errs.NewCantConvertJsonToStrErr()
	}
	str := string(out)
	return str, nil
}

func strToJson(v any, str string) error {
	if err := json.Unmarshal([]byte(str), &v); err != nil {
		return errs.NewCantConvertStrToJsonErr()
	}
	return nil
}
