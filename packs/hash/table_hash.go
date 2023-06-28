package hash

type Tupla struct {
	key   int
	value string
}

type TableHash struct {
	bucker   []Tupla
	tamTable int
}

func (table *TableHash) Init() {

}

func (table *TableHash) Add(key int, value string) {

}

func (table *TableHash) Remove(key int) {

}

func (table *TableHash) Search(key int) bool {
	return true
}

func (table *TableHash) Hash_function(key int) int {
	return 1
}
