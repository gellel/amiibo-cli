package main

type mixAmiibo struct {
	*compatabilityAmiibo
	*lineupAmiibo
}

func newMixAmiibo(c *compatabilityAmiibo, l *lineupAmiibo) (*mixAmiibo, error) {
	var (
		err error

		m = &mixAmiibo{c, l}
	)
	return m, err
}
