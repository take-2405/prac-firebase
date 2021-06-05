package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {
	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile("./path/myproject-63870-firebase-adminsdk-iv088-d9893bb49c.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// ドキュメント削除
	_, errorDelete := client.Collection("users").Doc("uesr2").Delete(ctx)
	if errorDelete != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}

	//切断
	defer client.Close()
}