package main



func main() {

	logger := NewLogger()

	service := NewService(logger)

	if err := service.Server.Serve(); err != nil {
		service.Logger.Error(err)
	}
}
