package set_map

type void struct{}

type any interface{}

var isMember void

type setMap map[any]void

func New() setMap {
	return setMap{}
}

func (s setMap) SADD(members ...interface{}) { s.set(members...) }

func (s setMap) SREM(members ...interface{}) { s.remove(members...) }

func (s setMap) SISMEMBER(member interface{}) bool { return s.exists(member) }

func (s setMap) SCARD() int { return s.size() }

func (s setMap) set(members ...interface{}) {
	for _, v := range members {
		s[v] = isMember
	}
}

func (s setMap) remove(members ...interface{}) {
	for _, v := range members {
		delete(s, v)
	}
}

func (s setMap) exists(value any) bool {
	_, ok := s[value]
	return ok
}

func (s setMap) size() int { return len(s) }
