package utils

import (
	"fmt"
)

func Throw(err any) {
	panic(err)
}

func ThrowIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func Try(f func()) error {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
			return
		}
	}()

	f()
	return err
}
