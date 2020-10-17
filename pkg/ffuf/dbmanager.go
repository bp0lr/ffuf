package ffuf

import (
	"fmt"
	"github.com/nanobox-io/golang-scribble"
)

type ResultDB struct{ 
	Rname 			string `json:"rname"`
	RstatusCode 	int64 `json:"rstatus"`
	RContentLength	int64 `json:"rlength"`
	RContentWords	int64 `json:"rwords"`
	RContentLines	int64 `json:"rlines"`
	RcleanLen		int64 `json:"rlen"`
}

func GetDbClient() *scribble.Driver{
	db, err := scribble.New("./resultsDB", nil)
	if err != nil {
		fmt.Printf("couldn't create db: " + err.Error())
	}

	return db
}

func UpdateDB(db *scribble.Driver, name string, key string, results ResultDB ){
	db.Write("db", name, results)
}


func ReturnAll(db *scribble.Driver) []string{
	records, err := db.ReadAll("db")
	if err != nil {
  		fmt.Println("Error", err)
	}
	return records
}