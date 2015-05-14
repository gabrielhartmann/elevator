package elevator

func sliceContains(reqs requests, floor int) (index []int, ok bool) {
	indices := []int{}
	for i, req := range reqs {
		if req.floor == floor {
			indices = append(indices, i)
		}
	}

	if len(indices) > 0 {
		return indices, true
	} else {
		return []int{}, false
	}
}
