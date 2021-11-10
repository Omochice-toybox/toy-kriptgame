package set

import "fmt"

type Set struct {
	data map[string]struct{}
}

func NewSet() *Set {
	return &Set{data: map[string]struct{}{}}
}

func (s *Set) Add(n string) {
	s.data[n] = struct{}{}
}

func (s *Set) Exists(n string) bool {
	_, ok := s.data[n]
	return ok
}

func (s *Set) Delete(n string) error {
	if !s.Exists(n) {
		return fmt.Errorf("%s is not exists in the Set", n)
	}
	delete(s.data, n)
	return nil
}

func (s *Set) Pop() (string, error) {
	for k := range s.data {
		delete(s.data, k)
		return k, nil
	}
	return "", fmt.Errorf("The set is empty")
}
