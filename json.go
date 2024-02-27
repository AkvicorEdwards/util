package util

import (
	"encoding/json"
)

type JSON struct {
	result  []byte
	err     error
	isArray bool
}

func NewJSONResult(data []byte) *JSON {
	m1 := make(map[string]any)
	err1 := json.Unmarshal(data, &m1)
	if err1 != nil {
		m2 := make([]map[string]any, 0)
		err2 := json.Unmarshal(data, &m2)
		if err2 != nil {
			return nil
		}
		return NewJSON(m2, true)
	}
	return NewJSON(m1, false)
}

func NewJSON(v any, isArray bool) *JSON {
	res := &JSON{}
	res.result, res.err = json.Marshal(v)
	res.isArray = isArray
	return res
}

func (j *JSON) Bytes() []byte {
	return j.result
}

func (j *JSON) String() string {
	return string(j.result)
}

func (j *JSON) Map(idx ...int) map[string]any {
	if j.result == nil {
		return nil
	}
	if !j.isArray {
		m := make(map[string]any)
		err := json.Unmarshal(j.result, &m)
		if err != nil {
			return nil
		}
		return m
	}
	i := 0
	if len(idx) > 0 {
		i = idx[0]
	}
	m := make([]map[string]any, 0)
	err := json.Unmarshal(j.result, &m)
	if err != nil {
		return nil
	}
	ml := len(m)
	if ml == 0 {
		return nil
	}
	if i >= ml {
		i = ml - 1
	}
	return m[i]
}

func (j *JSON) MapArray() []map[string]any {
	if j.isArray {
		m := make([]map[string]any, 0)
		err := json.Unmarshal(j.result, &m)
		if err != nil {
			return nil
		}
		return m
	} else {
		m := make(map[string]any)
		err := json.Unmarshal(j.result, &m)
		if err != nil {
			return nil
		}
		return []map[string]any{m}
	}
}

func (j *JSON) Error() error {
	return j.err
}
