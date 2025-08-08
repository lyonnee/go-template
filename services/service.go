package services

// StartAllServices 启动所有服务
func StartAllServices() {
    go StartHTTPServer()
    go StartCronScheduler()
    go StartGRPCServer()
}

// StopAllServices 停止所有服务
func StopAllServices() {
    StopHTTPServer()
    StopCronScheduler()
    StopGRPCServer()
}
