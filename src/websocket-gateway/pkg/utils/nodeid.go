package utils

import (
	"math/rand"
)

var NodeId string

func GenerateNodeId() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 16)
	for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GetNodeId() string {
	if NodeId == "" {
		NodeId = GenerateNodeId()
	}
	return NodeId
}