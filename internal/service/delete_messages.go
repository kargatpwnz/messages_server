package service

import "fmt"

func (s *MessagesService) DeleteMessage(id uint64) bool {
	fmt.Println("delete id ", id)
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	_, hasId := s.Storage[id]
	if hasId {
		delete(s.Storage, id)
	}
	return hasId
}
