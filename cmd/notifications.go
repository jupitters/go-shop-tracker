package main

import "sync"

type NotificationManager struct {
	clients map[string]map[chan string]bool
	mu      sync.RWMutex
}

func NewNotificationManager() *NotificationManager {
	return &NotificationManager{
		clients: make(map[string]map[chan string]bool),
	}
}

func (n *NotificationManager) AddClient(key string, client chan string) {
	n.mu.Lock()
	defer n.mu.Unlock()

	if n.clients[key] == nil {
		n.clients[key] = make(map[chan string]bool)
	}

	n.clients[key][client] = true
}
