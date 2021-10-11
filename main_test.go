package main

import "testing"

func TestGetASIN(t *testing.T) {
	cases := []struct {
		rawurl string
		want   string
	}{
		{
			rawurl: "https://www.amazon.co.jp/dp/B08DR7YL5J/ref=dp-kindle-redirect?_encoding=UTF8&btkr=1",
			want:   "B08DR7YL5J",
		},
		{
			rawurl: "https://www.amazon.co.jp/gp/product/B08G1PWN9X?",
			want:   "B08G1PWN9X",
		},
		{
			rawurl: "https://www.amazon.co.jp/%E3%82%BA%E3%83%AB%E3%81%84-%E5%90%88%E6%A0%BC%E6%B3%95-%E5%8C%BB%E8%96%AC%E5%93%81%E7%99%BB%E9%8C%B2%E8%B2%A9%E5%A3%B2%E8%80%85%E8%A9%A6%E9%A8%93%E5%AF%BE%E7%AD%96-%E9%B7%B9%E3%81%AE%E7%88%AA%E5%9B%A3%E7%9B%B4%E4%BC%9D-%E5%8F%82%E8%80%83%E6%9B%B8/dp/4910243062/ref=zg_bs_2541731051_1?_encoding=UTF8&psc=1&refRID=AP5RV4T9T6MT4FRRYJ2W",
			want:   "4910243062",
		},
	}

	for _, tt := range cases {
		got, err := getASIN(tt.rawurl)
		if err != nil {
			t.Errorf("[ERROR] %v\n", err)
		}
		if got != tt.want {
			t.Errorf("[ERROR] got: %v, want: %v\n", got, tt.want)
		}
	}
}
