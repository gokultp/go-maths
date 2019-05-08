package math

func Sum (args ...int) int {
	s := 0
	for _, i :=  range args {
		s += i
	}
	return s
}