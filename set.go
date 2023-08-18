package set_map

type void struct{}

type any interface{}

var isMember void

type SetMap map[any]void

func New() SetMap {
	return SetMap{}
}

func (s SetMap) SADD(members ...interface{}) { s.set(members...) }

func (s SetMap) SREM(members ...interface{}) { s.remove(members...) }

func (s SetMap) SISMEMBER(member interface{}) bool { return s.exists(member) }

func (s SetMap) SCARD() int { return s.size() }

func (s SetMap) set(members ...interface{}) {
	for _, v := range members {
		s[v] = isMember
	}
}

func (s SetMap) remove(members ...interface{}) {
	for _, v := range members {
		delete(s, v)
	}
}

func (s SetMap) exists(value any) bool {
	_, ok := s[value]
	return ok
}

func (s SetMap) size() int { return len(s) }
