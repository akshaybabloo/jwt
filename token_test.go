package jwt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUnixTimeToTime(t *testing.T) {
	unixTimeStamp := "1614581683"
	output, _ := time.Parse(time.RFC3339, "2021-03-01T19:54:43+13:00")

	toTime, err := UnixTimeToTime(unixTimeStamp)
	t.Log(toTime)
	 if assert.Nil(t, err) {
	 	assert.Equal(t, toTime, output)
	 }
}

func TestInTimeSpan(t *testing.T) {
	t1 := time.Date(1984, time.November, 3, 10, 23, 34, 0, time.UTC)
	t2 := time.Date(1984, time.November, 3, 11, 23, 34, 0, time.UTC)
	t3 := time.Date(1984, time.November, 3, 13, 0, 0, 0, time.UTC)

	span := InTimeSpan(t1, t3, t2)
	assert.True(t, span)
}
