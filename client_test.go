package ublox

import (
	"fmt"
	"os"
	"testing"

	"github.com/geotrace/geo"
)

var (
	token     = os.Getenv("UBLOX_TOKEN")
	pointWork = geo.NewPoint(37.57451, 55.715084)
	pointHome = geo.NewPoint(37.588248, 55.765944)
)

func TestClient(t *testing.T) {
	ubox := NewClient(token)
	data, err := ubox.GetOnline(pointWork, DefaultProfile)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(data)
}
