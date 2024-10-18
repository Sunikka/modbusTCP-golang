package main

import (
	"io"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {

	// .env is in the bin directory, because it's the only directory that the docker environment sees
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	envPath := filepath.Join(wd, ".env")

	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env")
	}

	serverAddr := os.Getenv("SERVER_ADDR") + ":" + os.Getenv("MB_PORT")

	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	log.Println("TCP server listening on:", serverAddr)
	// Currenly a basic echo server
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept connection: ", err)
			continue
		}
		log.Println("New client connection!")

		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()

	buf := make([]byte, 1024)
	n, err := c.Read(buf)
	if err != nil {
		if err == io.EOF {
			log.Println("Client disconnected")
		} else {
			log.Println("Error reading from client: ", err)
		}
		return
	}

	_, err = c.Write(buf[:n])
	if err != nil {
		log.Println("Error writing to client: ", err)
		return
	}

	log.Printf("Echoed %d bytes back to the client", n)

}
