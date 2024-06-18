package service

import "github.com/bigxxby/get-counter-test/internal/repository"

type Service interface {
	Counter() (int, error)
	ClearCounter() error
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}

// Counter increments the counter and returns its value
func (s *service) Counter() (int, error) {
	err := s.repo.CounterExists()
	if err != nil {
		return -1, err
	}
	count, err := s.repo.GetCount()
	if err != nil {
		return -1, err
	}
	count++

	err = s.repo.SetCount(count)
	if err != nil {
		return -1, err
	}

	return count, nil
}
func (s *service) ClearCounter() error {
	err := s.repo.SetCount(0)
	if err != nil {
		return err
	}
	return nil
}
