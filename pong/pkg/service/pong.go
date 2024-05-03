package service

// type PongService interface {
// 	Pong() string
// }

type PongServer struct {
}

func (s *PongServer) Pong() string {

	return "pong!"
}
