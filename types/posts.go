package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Posts struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID    primitive.ObjectID `bson:"userID" json:"userID"`
	Content   string             `bson:"content" json:"content"`
	ImagePath []string           `bson:"imagePath" json:"imagePath"`
	VideoPath []string           `bson:"videoPath" json:"videoPath"`
}
