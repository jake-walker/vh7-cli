// ISO 8601 UTC date parser

package vh7

import (
	"strings"
	"time"
)

type UtcTime time.Time

func (u *UtcTime) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`)
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse("2006-01-02T15:04:05", value)
	if err != nil {
		return err
	}
	*u = UtcTime(t)
	return nil
}

func (u UtcTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(u).Format("2006-01-02T15:04:05") + `"`), nil
}
