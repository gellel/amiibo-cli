package main

var (
	_ valuer = (&item{})
)

type item struct{}

func (i *item) Value() interface{} {
	return *i
}

func newItem(c *compatabilityItem, l *lineupItem) {}
