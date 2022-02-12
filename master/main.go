package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"master/clientmanager"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Initial Server")
	go clientmanager.Init()
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer listener.Close()

	// we need wait 3 client server online
	types := []string{"Mean", "Mode", "Median"}
	quitChan := make(chan string)
	// for receive every client's output
	typeChans := map[string]chan []byte{
		"Mean":   make(chan []byte, 1),
		"Mode":   make(chan []byte, 1),
		"Median": make(chan []byte, 1),
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Printf("Some connection error: %s\n", err)
			}
			manager := clientmanager.GetManager()
			lensOfClients := len(manager.Clients)
			if lensOfClients < len(types) {
				client := clientmanager.NewClient(conn, types[lensOfClients])
				// Assign type to client(to json format)
				typeString := map[string]string{
					"type": types[lensOfClients],
				}
				jsonString, _ := json.Marshal(typeString)
				client.Send(jsonString)
				// Read client result
				go ReadSocket(client, typeChans[types[lensOfClients]])

				fmt.Printf("Client [%s] is ready\n", types[lensOfClients])
				manager = clientmanager.GetManager()
				lensOfClients = len(manager.Clients)
				if lensOfClients == len(types) {
					fmt.Println("All client is ready now")
					// Start to read user's input
					go ReadInput(quitChan)
				}
			} else {
				fmt.Println("We don't need add more client")
			}
		}
	}()

	// Make sure output ordering is always same
	for {
		select {
		case message, ok := <-quitChan:
			if !ok {
				return
			}
			if message == "quit" {
				os.Exit(0)
			}
		default:
			if CheckChanIsFull(typeChans) {
				fmt.Println(string(<-typeChans["Mean"]))
				fmt.Println(string(<-typeChans["Mode"]))
				fmt.Println(string(<-typeChans["Median"]))
			}
		}

	}
}

func CheckChanIsFull(chans map[string]chan []byte) bool {
	for _, clientChan := range chans {
		if len(clientChan) == 0 {
			return false
		}
	}
	return true
}

func ReadSocket(client *clientmanager.Client, outputChan chan []byte) {
	ok := true
	for ok {
		msg, ok := <-client.ReadMessage
		if !ok {
			break
		}
		outputChan <- msg
	}
}

func ReadInput(signalChan chan string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input numbers")
	fmt.Println("---------------------")

	for {
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		// quit
		if strings.Compare("quit", text) == 0 {
			signalChan <- "quit"
			return
		}
		text = strings.Trim(text, " ")
		// Format nums
		splitSlice := strings.Split(text, " ")
		nums := []float64{}
		for _, str := range splitSlice {
			num, err := strconv.ParseFloat(str, 64)
			if err != nil {
				nums = []float64{}
				fmt.Println("Your number is invaild")
				break
			}
			nums = append(nums, num)
		}
		if len(nums) <= 0 {
			continue
		}
		// Broadcast to every client
		manager := clientmanager.GetManager()
		// Send nums to clients
		numsString := map[string][]float64{}
		numsString["nums"] = nums
		jsonString, _ := json.Marshal(numsString)
		manager.BroadCast(jsonString)

	}
}
