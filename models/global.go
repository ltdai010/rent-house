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
	"sync"
)

var (
	initOnce     sync.Once
	client       *firestore.Client
	clientSearch *search.Client
	searchIndex  *search.Index
	bucket       *storage.BucketHandle
	ctx          context.Context
)

func InitDataBase() {
	initOnce.Do(func() {
		initCloudStore()
		initStorage()
		initSearch()
	})
}

func initCloudStore() {
	ctx = context.Background()
	sa := option.WithCredentialsFile(beego.AppConfig.String("firebase::key_path"))
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}
	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
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