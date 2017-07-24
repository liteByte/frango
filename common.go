package frango

import (
	"fmt"
)

func PrintErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
