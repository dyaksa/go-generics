package main

type ServiceInterface interface {
	Save() string
	Update() string
}

func Save[T ServiceInterface](param T) string {
	return param.Save()
}

func Update[T ServiceInterface](param T) string {
	return param.Update()
}

type Service struct {
	Name string
}

func (s *Service) Save() string {
	return s.Name + " saved"
}

func (s *Service) Update() string {
	return s.Name + " updated"
}
