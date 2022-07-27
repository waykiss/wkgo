package utilsdate

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetDaysBetweenDates(t *testing.T) {
	baseDate := time.Date(2022, 7, 27, 0, 0, 0, 0, time.Local)
	testCases := []struct {
		startDate, endDate time.Time
		days               int
	}{
		{baseDate, baseDate, 0},
		{baseDate, baseDate, 0},
		{baseDate, baseDate.AddDate(0, 0, -5), -5},
		{baseDate, baseDate.AddDate(0, 0, 3), 3},
		{baseDate, baseDate.AddDate(1, 0, 0), 365},
		{baseDate.AddDate(0, 0, 3), baseDate.AddDate(0, 0, 0), -3},
		{baseDate.AddDate(1, 0, 0), baseDate.AddDate(0, 0, 0), -365},
	}
	for _, v := range testCases {
		got := GetDaysBetweenDates(v.startDate, v.endDate)
		assert.Equal(t, v.days, got)
	}
}
