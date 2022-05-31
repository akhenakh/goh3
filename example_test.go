package h3

import (
	"fmt"
)

func ExampleGeoToH3() {
	geo := GeoCoord{
		Latitude:  37.775938728915946,
		Longitude: -122.41795063018799,
	}
	resolution := 9

	fmt.Printf("%#x\n", GeoToH3(geo, resolution))
	// Output:
	// 0x8928308280fffff
}
