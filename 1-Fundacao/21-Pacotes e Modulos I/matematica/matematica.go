package matematica

func Soma[T int | float32 | float64](a, b T) T {
	return a + b
}
