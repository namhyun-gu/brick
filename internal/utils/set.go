package utils

type Set struct {
	data map[string]interface{}
}

func NewSet() Set {
	return Set{
		data: make(map[string]interface{}),
	}
}

func (s Set) Add(key string) {
	if !s.Contains(key) {
		s.data[key] = struct{}{}
	}
}

func (s Set) Contains(key string) bool {
	_, contain := s.data[key]
	return contain
}

func (s Set) Remove(key string) {
	delete(s.data, key)
}

func (s Set) Items() []string {
	items := make([]string, 0)
	for key, _ := range s.data {
		items = append(items, key)
	}
	return items
}
