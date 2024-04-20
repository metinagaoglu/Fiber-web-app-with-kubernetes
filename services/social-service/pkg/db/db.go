package db

import (
	"github.com/couchbase/gocb/v2"
	"log"
)

type Handler struct {
	DB *gocb.Cluster
}

func Init(url string, username string, password string) Handler {
	cluster, err := gocb.Connect("couchbase://"+url, gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: username,
			Password: password,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Get a reference to the default collection, required for older Couchbase server versions
	// col := bucket.DefaultCollection()
	return Handler{cluster}
}
