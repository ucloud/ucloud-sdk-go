package request

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSchema(t *testing.T) {
	s := "foo"
	assert.Equal(t, s, StringValue(String(s)))
	assert.Equal(t, "", StringValue(nil))

	i := 42
	assert.Equal(t, i, IntValue(Int(i)))
	assert.Equal(t, 0, IntValue(nil))

	f := 42.0
	assert.Equal(t, f, Float64Value(Float64(f)))
	assert.Equal(t, 0.0, Float64Value(nil))

	assert.Equal(t, true, BoolValue(Bool(true)))
	assert.Equal(t, false, BoolValue(nil))

	d := 1 * time.Second
	assert.Equal(t, d, TimeDurationValue(TimeDuration(d)))
	assert.Equal(t, time.Duration(0), TimeDurationValue(nil))
}
