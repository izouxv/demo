package api

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Id		bson.ObjectId	`bson:"_id,omitempty" json:"id"`
	Name  	string			`bson:"name,omitempty" json:"name"`
	Phone 	string			`bson:"phone,omitempty" json:"phone"`
	Age 	int				`bson:"age,omitempty" json:"age"`
}

// 文章
type Post struct {
	Content string `yaml:"context"`
	*Meta
}

type Meta struct {
	Title   string   `yaml:"title"`
	Path    string   `yaml:"path"`
	PubTime string   `yaml:"pub_time"`
	Tags    []string `yaml:"tags"`
	PostTime time.Time `yaml:"post_time"`
}
