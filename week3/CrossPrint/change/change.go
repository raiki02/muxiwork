package change

func Change(s string) (ss []string) {
	for _, v := range s {
		ss = append(ss, string(v))
	}
	return ss
}
