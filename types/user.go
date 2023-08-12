package types

type User struct {
	ID        string `bson:"_id" json:"id,omitempty"`
	FirstName string `bson:"firstName" json:"first_Name"`
	LastName  string `bson:"lastName" json:"last_Name"`
}
