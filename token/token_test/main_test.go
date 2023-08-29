package token_test

import (
	"os"
	"testing"

	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/token"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/util"
)

var testJwtMaker token.Maker
var testPasetoMaker token.Maker
func TestMain(m *testing.M) {
	var err error
	testJwtMaker, err = token.NewJWTMaker(util.RandomString(32))
	if err != nil {
		os.Exit(1)
	}

	testPasetoMaker = token.NewPasetoMaker(util.RandomString(32))
	os.Exit(m.Run())
}
