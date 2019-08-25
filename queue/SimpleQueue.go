package queue

type SimpleQueue struct {
	DB []interface{}
}

func (s *SimpleQueue) Put(e interface{}) {
	s.DB = append(s.DB, e)
}

func (s *SimpleQueue) Poll() interface{} {
	e := s.DB[0]
	s.DB = s.DB[1:len(s.DB)]
	return e
}

func (s *SimpleQueue) IsEmpty() bool {
	return len(s.DB) == 0
}
