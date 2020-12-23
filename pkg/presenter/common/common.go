package common

import "github.com/go-email-validator/go-email-validator/pkg/ev/utils"

func MX2String(MXs utils.MXs) []string {
	var result = make([]string, len(MXs))
	for i, mx := range MXs {
		result[i] = mx.Host
	}

	return result
}
