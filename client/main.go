package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	chat "main/chat"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := chat.NewChatServiceClient(conn)
	ctx := context.Background()

	stream, err := client.Chat(ctx)
	if err != nil {
		log.Fatalf("Error creating chat stream: %v", err)
	}

	isFirstBatch := true
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter message: ")
		scanned := scanner.Scan()
		if !scanned {
			if err := scanner.Err(); err != nil {
				log.Printf("Error reading from input: %v", err)
			}
			return
		}
		message := scanner.Text()
		err := stream.Send(&chat.Message{Content: message})
		if err != nil {
			log.Printf("Error sending message: %v", err)
		}

		for {
			response, err := stream.Recv()
			if err != nil {
				log.Printf("Error receiving response: %v", err)
				break
			}

			if response.Content == "" {
				if isFirstBatch {
					fmt.Print("Response: ")
					isFirstBatch = false
					continue
				}
				isFirstBatch = true
				fmt.Println()
				break
			}

			fmt.Printf("%s", response.Content)
		}
	}
}
