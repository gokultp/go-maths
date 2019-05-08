package math

func Sum (args ...int) int {
	s := 0
	for _, i :=  range args {
		s += i
	}
	return s
}

func Mul (args ...int) int {
	p := 0
	for _, i :=  range args {
		p *= i
	}
	return p
}