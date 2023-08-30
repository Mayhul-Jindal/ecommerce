package email_test

import (
	"os"
	"testing"

	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/util"
)

var config util.Config

func TestMain(m *testing.M) {
	var err error
	config, err = util.LoadConfig("../..")
	if err != nil {
		os.Exit(1)
	}

	os.Exit(m.Run())
}
