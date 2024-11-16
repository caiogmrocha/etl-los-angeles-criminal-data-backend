package value_objects

type Sex string

const (
	Male      Sex = "M"
	Female    Sex = "F"
	NonBinary Sex = "X"
	Unknown   Sex = "U"
)
