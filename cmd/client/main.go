package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	// .env is in the bin directory, because only that directory is seen by the docker container
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	envPath := filepath.Join(wd, "/bin/.env")
	log.Printf(".env filepath: %s", envPath)

	godotenv.Load(envPath)

	// Modbus server address
	port := os.Getenv("MB_PORT")
	addr := os.Getenv("SERVER_ADDR")

	log.Printf("Target address: %s", fmt.Sprintf("%s:%s", addr, port))
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", addr, port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Printf("Connection to modbus server established")

	// first commandline argument for the client is the message

	msg := []byte(os.Args[1])
	buf := make([]byte, 1024)

	conn.Write(msg)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Println("Server closed the connection")
				break
			}
			log.Fatal("Error reading from the server: ", err)
		}

		if n > 0 {
			res := string(buf[:n])
			fmt.Printf("Server response: %s", res)
			break
		}

	}
}
