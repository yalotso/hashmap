package hashmap

func Index(key, size int) int {
	return abs(hash(key)) % size
}

func hash(x int) int {
	return (x >> 15) ^ x
}

func abs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}
