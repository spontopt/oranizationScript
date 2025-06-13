package main

import (
	"log"
	"organizationScript/data"
	"organizationScript/process"
)

func main() {
	log.Println("running script")

	/* 	s := []data.Subscription{{
	   		Id:    "0196e7cf-8706-7186-af0c-f1c385131e2d",
	   		Name:  "MC Testing",
	   		Email: "mc+2025051901@kamae.pt",
	   	},
	   	} */

	log.Println("connecting to db")
	connect, err := data.DBConnect()
	if err != nil {
		log.Fatal(err)
	}

	db := data.MongoDb{
		Db: connect,
	}

	log.Println("starting processing")
	process.ImportStatus(db)

	/* process.ProcessSubscriptions(s, db) */

	//process.GetAllSubscriptions(db)
	log.Println("finished")
}
