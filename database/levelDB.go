package database

import (
	"fmt"
	"gonode/result"
	"sort"

	"github.com/syndtr/goleveldb/leveldb"
	"google.golang.org/protobuf/proto"
)

func DBAddResult(source string, form *result.MineResult) {
	db, err := leveldb.OpenFile("../.history/index", nil)
	if err != nil {
		panic(err)
	}
	spv, notIn := DBIndexLookup(source, *db)
	mineResult := &result.Results{
		Res: []*result.MineResult{},
	}
	if !notIn {
		mineResult = &result.Results{
			Res: []*result.MineResult{form},
		}
		res, _ := proto.Marshal(mineResult)
		fmt.Printf("Adding %s result to database.\n", source)
		db.Put([]byte(source), res, nil)
	} else {
		proto.Unmarshal(spv, mineResult)
		spvSlice := mineResult.GetRes()
		for _, entries := range spvSlice {
			if entries.Rotation == form.Rotation {
				fmt.Println("Entry with this rotation already in DB")
				break
			} else {
				spvSlice = append(mineResult.GetRes(), form)
				sort.Slice(spvSlice, func(p, q int) bool {
					return spvSlice[p].Rotation > spvSlice[q].Rotation
				})
				mineResult = &result.Results{
					Res: spvSlice,
				}
				h, _ := proto.Marshal(mineResult)
				fmt.Printf("Adding %s result to database.\n", source)
				db.Put([]byte(source), h, nil)
			}
		}
	}

	// db.Put([]byte(source), form, nil

	db.Close()
}

func DBAddData(datahash string, data string) {
	db, err := leveldb.OpenFile("../.history/data", nil)
	if err != nil {
		panic(err)
	}
	db.Put([]byte(datahash), []byte(data), nil)

	db.Close()
}

func DBDataLookup(datahash string) []byte {
	db, err := leveldb.OpenFile("../.history/data", nil)
	if err != nil {
		panic(err)
	}
	data, err := db.Get([]byte(datahash), nil)
	if err != nil {
		fmt.Printf("Not in database")
	}
	fmt.Printf("Data being sent to client\n")
	db.Close()
	return data
}

func DBIndexLookup(root string, db leveldb.DB) ([]byte, bool) {
	//db, err := leveldb.OpenFile("../.history/index", nil)
	inDb := true
	data, err := db.Get([]byte(root), nil)

	if err != nil {
		inDb = false
		fmt.Println("Not a key in the index")
	}
	values := &result.Results{
		Res: []*result.MineResult{},
	}
	proto.Unmarshal(data, values)
	//db.Close()
	return data, inDb
}

func DBListData() {
	db, err := leveldb.OpenFile("../.history/index", nil)
	if err != nil {
		panic(err)
	}
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Println(string(key), string(value))
	}
	iter.Release()
	err = iter.Error()
	db.Close()
}
