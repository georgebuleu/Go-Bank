package main

func main() {

	server := NewApiServer(":7890")
	server.run()

}
