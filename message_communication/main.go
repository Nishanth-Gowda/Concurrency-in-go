package main

import (
	"fmt"
	"time"

)

type Message struct {
	From    string
	Payload string
}

type Server struct {
	msgCh chan Message
	quitCh chan struct{}
}

func (s *Server) StartAndListen() {

	//THis corresponds to naming for loop
	free:
		for {
			//select is like a switch more of a concurrent switch
			select {
			//Blocks here until someone is sending a message to channel.
			case msg := <-s.msgCh:
				fmt.Printf("Recieved Message from %s with Payload %s\n", msg.From, msg.Payload)
			case <- s.quitCh:
				fmt.Println("Entering Graceful Shutdown")
				//Use your logic 
				break free
			default:
			}
	}

	fmt.Println("The Server is shut down")
}

func sendMessageToServer(msgCh chan Message, payload string) {
	msg := Message{
		From:    "Naredra Modi",
		Payload: payload,
	}

	msgCh <- msg
	fmt.Println("Sending Message")
}

func gracefullShutDown(quitCh chan struct{}) {
	close(quitCh)
}

func main() {
	s := &Server{
		msgCh: make(chan Message),
		quitCh: make(chan struct{}),
	}

	go s.StartAndListen()

	go func() {
		time.Sleep(2 * time.Second)
		sendMessageToServer(s.msgCh, "Hello Modi")
	} ()

	go func() {
		time.Sleep(4 * time.Second)
		gracefullShutDown(s.quitCh)
	} ()

	//!!use this only for demonstration, this forces the deadlock
	select {}

}
