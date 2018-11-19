package main



func main() {

	logger := NewLogger()

	service := NewService(logger)

	if err := service.Server.Server.ListenAndServe(); err != nil {
		service.Logger.Error(err)
	}
}
