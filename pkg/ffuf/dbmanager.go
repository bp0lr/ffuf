package ffuf

import (
	//"fmt"
	//"log"
	"github.com/nanobox-io/golang-scribble"
)

type ResultDB struct{ 
	key 			string
	statusCode 		int64
	ContentLength	int64
	ContentWords	int64
	ContentLines	int64
	cleanLen		int64
}

var mydb *scribble.Driver