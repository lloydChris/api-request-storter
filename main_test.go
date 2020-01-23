package main

import "testing"

func TestParseLine(t *testing.T) {
	var tests = make(map[string]string)
	tests["POST /Brands_Download.ashx?output=json HTTP/1.1"] = "POST /Brands_Download.ashx"
	tests["GET /v2/Get_Photo.ashx?key=IHUYuuqwfGo%3D&size=profile HTTP/1.1"] = "GET /v2/Get_Photo.ashx"
	tests[""] = ""
	tests["POST /en_US/ads/tracking/?authid=dd%3D%3D&key=w3rsder HTTP/1.1"] = "POST /en_US/ads/tracking"
	tests["POST /en_US/ads/tracking?authid=dd%3D%3D&key=43srw3 HTTP/1.1"] = "POST /en_US/ads/tracking"
	tests["POST /v2/en_us/prices/report?authId=dd%3D%3D&key=wefs&memberid=Truegas HTTP/1.1"] = "POST /v2/en_us/prices/report"

	for text, expected := range tests {
		actual := ParseURL(text)
		if actual != expected {
			t.Error(expected, "<>", actual)
		}
	}
}
