package elevator

type request struct {
	floor     int
	direction int
}

type requests []request

func (slice requests) Len() int {
	return len(slice)
}

func (slice requests) Less(i, j int) bool {
	return slice[i].floor < slice[j].floor
}

func (slice requests) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
