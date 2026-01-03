package phpArray

type PArray struct {
	keys  []*key
	array map[string]interface{}
}

type key struct {
	key   string
	index int
	prv   *key
	next  *key
}
