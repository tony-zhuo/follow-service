package libs

import (
	"strconv"
	"time"
)

func StringConvertToTime(data string) (*time.Time, error) {
	v, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return nil, err
	}
	s := v / 1000
	ns := (v % 1000) * 1e6
	res := time.Unix(s, ns)
	return &res, nil
}
