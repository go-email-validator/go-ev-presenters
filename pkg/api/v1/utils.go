package v1

import "time"

// OnlyPositiveInt returns positive or zero value
func OnlyPositiveInt(val int) int {
	if val < 0 {
		return 0
	}

	return val
}

// OnlyPositiveDuration returns positive or zero value
func OnlyPositiveDuration(val time.Duration) time.Duration {
	if val < 0 {
		return 0
	}

	return val
}

func SecondDuration(second float64) time.Duration {
	return time.Duration(second * float64(time.Second))
}
