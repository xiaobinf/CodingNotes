package go_util

type HashSet struct {
	set map[int]bool
}

func NewHashSet() *HashSet {
	return &HashSet{make(map[int]bool)}
}

func (set *HashSet) Add(i int) bool {
	_, found := set.set[i]
	set.set[i] = true
	return !found //False if it existed already
}

func (set *HashSet) Get(i int) bool {
	_, found := set.set[i]
	return found //true if it existed already
}

func (set *HashSet) Remove(i int) {
	delete(set.set, i)
}
