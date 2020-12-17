package yahb

import (
	"io/ioutil"
	"log"
	"testing"
)

var placeBuf, _ = ioutil.ReadFile("./testdata/place.json")
var placeObj = new(Place)

var settingsBuf, _ = ioutil.ReadFile("./testdata/setting.json")
var settingsObj = new(Setting)

var reqBuf, _ = ioutil.ReadFile("./testdata/req.json")
var reqObj = new(BidRequest)

func _testPlaceUnmarshalJSON(t *testing.T) {
	err := placeObj.UnmarshalJSON(placeBuf)
	if err != nil {
		t.Fatal(err)
	}
}

func _testPlaceValidate(t *testing.T) {
	err := placeObj.Validate()
	if err != nil {
		t.Fatal(err)
	}
}

func TestPlace(t *testing.T) {

	t.Parallel()

	_testPlaceUnmarshalJSON(t)
	_testPlaceValidate(t)
}

func _testSettingsUnmarshalJSON(t *testing.T) {
	err := settingsObj.UnmarshalJSON(settingsBuf)
	if err != nil {
		t.Fatal(err)
	}
}

func _testSettingsValidate(t *testing.T) {
	err := settingsObj.Validate()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSettings(t *testing.T) {

	t.Parallel()

	_testSettingsUnmarshalJSON(t)
	_testSettingsValidate(t)
}

func _testBidRequestUnmarshalJSON(t *testing.T) {
	err := reqObj.UnmarshalJSON(reqBuf)
	if err != nil {
		t.Fatal(err)
	}
}

func _testBidRequestValidate(t *testing.T) {
	err := reqObj.Validate()
	if err != nil {
		t.Fatal(err)
	}
}

func TestBidRequest(t *testing.T) {

	t.Parallel()

	_testBidRequestUnmarshalJSON(t)
	_testBidRequestValidate(t)
}

func BenchmarkBidRequest_MarshalJSON(b *testing.B) {
	bidReq := new(BidRequest)
	bidReq.Setting = &Setting{
		Currency: "RUB",
		WinSize: &Size{
			Width:  100,
			Height: 100,
		},
	}
	bidReq.Places = append(bidReq.Places, Place{
		ID:          "1",
		PlacementId: "1",
		Sizes:       [][]int{[]int{100, 100}, []int{200, 200}},
	})

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, e := bidReq.MarshalJSON()
			if e != nil {
				b.Fatal(e)
			}
		}
	})
}

func BenchmarkBidRequest_UnMarshalJSON(b *testing.B) {
	buf := []byte("{\"places\":[{\"id\":\"1\",\"placementId\":\"1\",\"sizes\":[[100,100],[200,200]]}],\"settings\":{\"currency\":\"RUB\",\"windowSize\":{\"width\":100,\"height\":100}}}")
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bidReq := new(BidRequest)
			log.Fatal(string(buf))
			e := bidReq.UnmarshalJSON(buf)
			if e != nil {
				b.Fatal(e)
			}
		}
	})
}
