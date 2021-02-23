package main

import (
	"database/sql"
	"fmt"
	"github.com/domesama/mysql-go-helloworld/platform/newsfeed"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	db, err := sql.Open("mysql", "OwO:OwO@tcp(localhost)/news_feed")
	if err != nil{
		log.Fatal(err)
	}

	feed := newsfeed.NewFeed(db)

	//feed.AddFeed(newsfeed.Item{Content: "UwU"})
	items := feed.GetFeed()

	fmt.Println(items)


}

