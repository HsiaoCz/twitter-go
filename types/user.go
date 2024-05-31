package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName       string             `bson:"firstName" json:"firstName"`
	LastName        string             `bson:"lastName" json:"lastName"`
	NickName        string             `bson:"nickName" json:"nickName"`
	Avatar          string             `bson:"avatar" json:"avatar"`
	BackgroundImage string             `bson:"backgroundImage" json:"backgroundImage"`
	Synopsis        string             `bson:"synopsis" json:"synopsis"`
	JoinedTime      string             `bson:"joinedTime" json:"joinedTime"`
	Following       string             `bson:"following" json:"following"` // use redis 
	Followers       string             `bson:"follwoers" json:"followers"` // use redis
}
