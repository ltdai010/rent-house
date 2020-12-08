package convert

func ArrayToString(a []string) string {
	s := ""
	for _, i := range a {
		s += i;
	}
	return s
}
