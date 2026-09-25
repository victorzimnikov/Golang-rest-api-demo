package v1

import (
	"bytes"
	"encoding/json"
	"time"
)

type OptionalDate struct {
	Set   bool
	Value *time.Time
}

func (d *OptionalDate) UnmarshalJSON(data []byte) error {
	d.Set = true

	if bytes.Equal(bytes.TrimSpace(data), []byte("null")) {
		d.Value = nil

		return nil
	}

	var value time.Time
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	d.Value = &value
	return nil
}
