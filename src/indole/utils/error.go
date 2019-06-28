package utils

import "log"

// FirstError ...
func FirstError(errs ...error) error {
	for _, v := range errs {
		if v != nil {
			return v
		}
	}
	return nil
}

// Recover ...
func Recover(v ...interface{}) {
	if r := recover(); r != nil {
		log.Println(append(v, r)...)
	}
}
