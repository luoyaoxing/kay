package dec

import "encoding/json"

func JsonEncode(v interface{}) string {
	if v == nil {
		return ""
	}
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}

func JsonDecode(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}
