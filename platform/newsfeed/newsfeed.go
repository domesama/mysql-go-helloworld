package newsfeed

import "database/sql"

type Feed struct{
	DB *sql.DB
}

func NewFeed(db *sql.DB) *Feed{
		statement, _ := db.Prepare(`
		create table if not exists news_feed
		(
			id int auto_increment,
			content varchar(500) null,
			constraint news_feed_pk
		primary key (id)
		);
	`)
		statement.Exec()
	return &Feed{DB: db}
}

func (f *Feed) AddFeed(item Item){
	statement, _ := f.DB.Prepare(`INSERT INTO news_feed (content) values (?)`)
	statement.Exec(item.Content)
}

func (f *Feed) GetFeed() []Item{
	items := []Item{}
	rows, _ := f.DB.Query(`SElECT * FROM news_feed`)

	var id int
	var content string
	
	for rows.Next(){
		rows.Scan(&id, &content)
		items = append(items, Item{Id: id, Content: content})
	}

	return items
}