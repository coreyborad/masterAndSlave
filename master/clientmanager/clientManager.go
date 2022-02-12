package clientmanager

import (
	"log"
	"sync"
)

var manager *ClientManager

type ClientManager struct {
	sync.Mutex
	Clients map[*Client]bool
}

func Init() error {
	manager = &ClientManager{
		Clients: make(map[*Client]bool),
	}

	return nil
}

func GetManager() *ClientManager {
	return manager
}

func (m *ClientManager) Register(client *Client) {
	m.Lock()
	defer m.Unlock()
	m.Clients[client] = true
}

// Unregister Unregister
func (m *ClientManager) Unregister(client *Client) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.Clients[client]; ok {
		log.Printf("Unregister: %v", client.ID)
		// client.Lock()
		// client.Cancel()
		close(client.ReadMessage)
		close(client.WriteMessage)
		delete(m.Clients, client)
		client = nil
	}
}

func (m *ClientManager) BroadCast(msg []byte) {
	m.Lock()
	defer m.Unlock()

	for client, _ := range m.Clients {
		client.Send(msg)
	}
}
