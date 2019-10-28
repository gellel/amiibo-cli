package main

type mixAmiibo struct {
	Compatability *compatabilityAmiibo
	Lineup        *lineupAmiibo
}

func newMixAmiibo(c *compatabilityAmiibo, l *lineupAmiibo) (*mixAmiibo, error) {
	var (
		err error

		m = &mixAmiibo{c, l}
	)
	return m, err
}
