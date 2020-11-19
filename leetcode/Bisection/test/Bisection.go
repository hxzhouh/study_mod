package test

func Bisection(a []int, value int) bool {
	if len(a) == 1 {
		return a[0] == value
	}
	mid := len(a) / 2
	if a[mid] == value {
		return true
	} else if a[mid] < value {
		return Bisection(a[mid+1:], value)
	} else {
		return Bisection(a[:mid], value)
	}
}
