// generated by jsonenums -type=AppStatus; DO NOT EDIT

package models

import (
	"encoding/json"
	"fmt"
)

var (
	_AppStatusNameToValue = map[string]AppStatus{
		"Running":    Running,
		"Stopped":    Stopped,
		"Undeployed": Undeployed,
	}

	_AppStatusValueToName = map[AppStatus]string{
		Running:    "Running",
		Stopped:    "Stopped",
		Undeployed: "Undeployed",
	}
)

func init() {
	var v AppStatus
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_AppStatusNameToValue = map[string]AppStatus{
			interface{}(Running).(fmt.Stringer).String():    Running,
			interface{}(Stopped).(fmt.Stringer).String():    Stopped,
			interface{}(Undeployed).(fmt.Stringer).String(): Undeployed,
		}
	}
}

// MarshalJSON is generated so AppStatus satisfies json.Marshaler.
func (r AppStatus) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _AppStatusValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid AppStatus: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so AppStatus satisfies json.Unmarshaler.
func (r *AppStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("AppStatus should be a string, got %s", data)
	}
	v, ok := _AppStatusNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid AppStatus %q", s)
	}
	*r = v
	return nil
}