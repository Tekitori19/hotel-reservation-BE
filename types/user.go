package types

type User struct {
	ID        string `bson:"_id,omitemty" json:"id,omitempty"`
	FirstName string `bson:"firstName" json:"first_Name"`
	LastName  string `bson:"lastName" json:"last_Name"`
}
