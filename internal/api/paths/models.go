package paths

type UserId struct {
	Id string `uri:"userId" json:"userId" bson:"userId" db:"userId"`
}
