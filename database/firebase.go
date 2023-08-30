package database

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func Firebase() *firestore.Client {
	context := context.Background()
	// TODO: move this to .env
	sa := option.WithCredentialsFile("firebaseServiceAccount.json")
	app, err := firebase.NewApp(context, nil, sa)

	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(context)
	if err != nil {
		log.Fatalln(err)
	}

	return client
}
