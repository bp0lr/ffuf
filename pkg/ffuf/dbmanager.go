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
	OriginalRes		Response `json:"response"`
}

func GetDbClient() *scribble.Driver{
	db, err := scribble.New("./resultsDB", nil)
	if err != nil {
		fmt.Printf("couldn't create db: %v\n", err.Error())
	}

	return db
}

func UpdateDB(db *scribble.Driver, folder string, name string, results ResultDB ){
	db.Write(folder, name, results)
}


func ReturnAll(db *scribble.Driver, folder string) []string{
	records, _ := db.ReadAll(folder)
	//if err != nil {
  		//fmt.Printf("\nError: %v\n", err)
	//}
	return records
}