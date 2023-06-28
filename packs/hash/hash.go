package hash

type Ihash interface {
	Hash_function(key int) int
	Init()
	Add(key int, value string)
	Remove(key int)
	Search(key int) bool
}
