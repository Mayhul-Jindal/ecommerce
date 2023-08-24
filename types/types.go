// TODO
/*
instead of this package , models.go should be used
*/

package types

type contextKey string

const (
	CtxKey contextKey = "reqID"
)

type Book struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

type Review struct {
	Username string `json:"username"`
	Rating   int    `json:"rating"`
	Comment  string `json:"comment"`
}
