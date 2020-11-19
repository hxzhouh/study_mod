package test

import (
	"fmt"
	"testing"
)

func Test_lastValueBisection(t *testing.T) {

	a := []int{1, 3, 4, 5, 6, 8, 8, 8, 8, 11, 18}
	fmt.Println(lastValueBisection(a, 8))
}
