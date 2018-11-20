package main

func main() {

	service := NewService(NewLogger(), NewSlack(), NewRouter())

	server := NewServer().New(service.Router)

	if err := server.Serve(); err != nil {
		service.Logger.Error(err)
	}

}
