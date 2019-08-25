package engine

func contains(list []string, find string) bool {
	for _, d := range list {
		if d == find {
			return true
		}
	}
	return false
}
