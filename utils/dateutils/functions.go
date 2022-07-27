package dateutils

import "time"

//FromString function to return a time.Time given a string, it automatically tries to convert from several and
// common formats, if it has error, return time zeroValue
func FromString(date string) (r time.Time) {
	formats := []string{
		"2006-01-02",
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05Z",
		"2006-01-02 15:04:05",
		"02/01/2006 15:04:05", //BR format
		"02/01/2006",
	}
	var err error
	for _, f := range formats {
		r, err = time.Parse(f, date)
		if err == nil {
			return
		}
	}
	return
}
