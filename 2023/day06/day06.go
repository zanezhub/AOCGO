package day02

func SolvePart1(time int, dis int) int {
	var (
		res int
	)

	for i := 0; i <= time; i++ {
		cur := i * (time - i)
		if cur > dis {
			res++
		}
	}

	return res
}
