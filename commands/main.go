package main

import (
	"database/sql"
	"github.com/domesama/mysql-go-helloworld/platform/newsfeed"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/qkgo/yin"
	"log"
	"net/http"
)

func main() {

	db, err := sql.Open("mysql", "OwO:OwO@tcp(localhost)/news_feed")
	if err != nil{
		log.Fatal(err)
	}

	feed := newsfeed.NewFeed(db)
	//
	//feed.AddFeed(newsfeed.Item{Content: "UwU"})
	//items := feed.GetFeed()
	//
	//fmt.Println(items)

	r := chi.NewRouter()
	r.Use(yin.SimpleLogger)

	r.Get("/posts",func(w http.ResponseWriter, r *http.Request){
		res, _ := yin.Event(w,r)
		items := feed.GetFeed()

		res.SendJSON(items)
	})

	r.Post("/posts", func(w http.ResponseWriter, r *http.Request){
		res, req := yin.Event(w,r)
		body := map[string]string{}
		req.BindBody(&body)

		item := newsfeed.Item{
			Content: body["content"],
		}

		feed.AddFeed(item)

		res.SendStatus(200)
	})
	
	http.ListenAndServe(":8080", r)


}

