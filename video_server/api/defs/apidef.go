package defs

type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}
type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	Displayctime string
}
type SimpleSession struct {
	UserName string
	TTL      int64
}
