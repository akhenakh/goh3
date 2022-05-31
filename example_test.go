package h3

import (
	"fmt"
)

func ExampleFromGeo() {
	geo := GeoCoord{
		Latitude:  37.775938728915946,
		Longitude: -122.41795063018799,
	}
	resolution := 9

	h := FromGeo(geo, resolution)
	fmt.Printf("%#x\n", h)
	// Output:
	// 0x8928308280fffff

}

func ExampleToGeo() {
	h := H3Index(0x8928308280fffff)

	g := ToGeo(h)
	fmt.Printf("%f %f\n", g.Latitude, g.Longitude)
	// Output:
	// 37.776702 -122.418459

}
