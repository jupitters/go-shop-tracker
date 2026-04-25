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
