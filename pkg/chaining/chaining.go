package chaining

type Modifier func(int) int

func Double(i int) int {
	return i * 2
}

func AddTen(i int) int {
	return i + 10
}

func Chain(i int, funcs ...Modifier) int {
	result := i
	for _, fn := range funcs {
		result = fn(result)
	}
	return result
}
