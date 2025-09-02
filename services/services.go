package services

type Service interface {
	Start()
	Stop()
}

var (
	services []Service
)

func RegisterService(s Service) {
	services = append(services, s)
}

// StartAll 启动所有服务
func StartAll() {
	for _, s := range services {
		go s.Start()
	}
}

// StopAll 停止所有服务
func StopAll() {
	for _, s := range services {
		s.Stop()
	}
}
