package service

import "strconv"

func IsStatusCodeOK(code int) bool {
	str := strconv.Itoa(code)
	firstDigit := int(str[0] - '0')

	return firstDigit == 2 || firstDigit == 3
}

func IsInternalServerError(code int) bool {
	str := strconv.Itoa(code)
	firstDigit := int(str[0] - '0')

	return firstDigit == 5
}
