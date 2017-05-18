package tests

import (
	"github.com/sddysz/leanote/app/db"
	"testing"
	//	. "github.com/sddysz/leanote/app/lea"
	//	"github.com/sddysz/leanote/app/service"
	//	"gopkg.in/mgo.v2"
	//	"fmt"
)

func TestDBConnect(t *testing.T) {
	db.Init("mongodb://localhost:27017/leanote", "leanote")
}
