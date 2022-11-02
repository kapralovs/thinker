package server

type server struct {
}

func New() *server {
	return &server{}
}

func (s *server) Run() error {
	return nil
}
