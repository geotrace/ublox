package ublox

import (
	"testing"

	"gopkg.in/mgo.v2"
)

func TestCache(t *testing.T) {
	mongodb, err := mgo.Dial("mongodb://localhost/geotrace")
	if err != nil {
		t.Fatal(err)
	}
	defer mongodb.Close()

	cache, err := InitCache(mongodb, "", token)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 1000; i++ {
		data, err := cache.Get(pointHome, DefaultProfile)
		if err != nil {
			t.Fatal(err)
		}
		// fmt.Println(data)
		data, err = cache.Get(pointWork, DefaultProfile)
		if err != nil {
			t.Fatal(err)
		}
		// fmt.Println(data)
		_ = data
		// jsondata, err := json.Marshal(data)
		// if err != nil {
		// 	t.Fatal(err)
		// }
		// fmt.Println("json:", string(jsondata))
	}
}
