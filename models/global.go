package models

import (
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	"context"
	firebase "firebase.google.com/go"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/astaxie/beego"
	"google.golang.org/api/option"
	"log"
	"net/smtp"
	"sync"
)

var (
	initOnce     sync.Once
	Client       *firestore.Client
	clientSearch *search.Client
	searchIndex  *search.Index
	bucket       *storage.BucketHandle
	ctx          context.Context
	EmailAuth    smtp.Auth
	EmailFrom    string
)

func InitDataBase() {
	initOnce.Do(func() {
		initCloudStore()
		initStorage()
		initSearch()
		initGmail()
	})
}

func initCloudStore() {
	ctx = context.Background()
	sa := option.WithCredentialsFile(beego.AppConfig.String("firebase::key_path"))
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}
	Client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}

func initGmail() {
	emailHost := "smtp.gmail.com"
	EmailFrom = "ltdai2468@gmail.com"
	emailPassword := "dajape0#"
	EmailAuth = smtp.PlainAuth("", EmailFrom, emailPassword, emailHost)
}

func initSearch()  {
	clientSearch = search.NewClient(beego.AppConfig.String("algolia::app_id"), beego.AppConfig.String("algolia::key"))
	searchIndex = clientSearch.InitIndex("house")
}

func initStorage()  {
	config := &firebase.Config{
		StorageBucket: beego.AppConfig.String("storage::bucket"),
	}
	ctx := context.Background()
	sa := option.WithCredentialsFile(beego.AppConfig.String("firebase::key_path"))
	app, err := firebase.NewApp(ctx, config, sa)
	if err != nil {
		log.Fatalln(err)
	}
	clientStorage, err := app.Storage(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	bucket, err = clientStorage.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}
}