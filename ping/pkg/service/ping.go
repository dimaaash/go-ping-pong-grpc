package service

// type PingService interface {
// 	Ping() string
// }

type PingServer struct {
}

func (s *PingServer) Ping() string {

	return "ping!"
}
