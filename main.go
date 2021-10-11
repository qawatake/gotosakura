package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("コマンドライン引数の数が不正です")
	}

	rawurl := os.Args[1]
	browser := os.Args[2]
	asin, err := getASIN(rawurl)
	if err != nil {
		log.Fatalf("getASIN failed: %v", err)
	}
	if (asin == "") {
		fmt.Print("不正なページを開いているようです")
		return
	}

	sakuraurl := fmt.Sprintf("https://sakura-checker.jp/search/%v/", asin)
	alfredjson := NewAlfredJSON(sakuraurl, browser)

	ec := json.NewEncoder(os.Stdout)
	if err := ec.Encode(alfredjson); err != nil {
		log.Fatal(err)
	}
}

func getASIN(rawurl string) (string, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return "", nil
	}

	parts := strings.Split(u.Path, "/")

	// http://www.amazon.co.jp/dp/[ASIN]/
	if len(parts) >= 3 && strings.Compare(parts[1], "dp") == 0 {
		matched, err := regexp.MatchString(`^[0-9A-Z]{10}$`, parts[2])
		if err != nil {
			return "", err
		}
		if matched {
			return parts[2], nil
		}
	}

	// http://www.amazon.co.jp/[商品名]/dp/[ASIN]
	if len(parts) >= 4 && strings.Compare(parts[2], "dp") == 0 {
		matched, err := regexp.MatchString(`^[0-9A-Z]{10}$`, parts[3])
		if err != nil {
			return "", nil
		}
		if matched {
			return parts[3], nil
		}
	}

	// http://amazon.co.jp/gp/product/[ASIN]/
	if len(parts) >= 4 && strings.Compare(parts[1], "gp") == 0 && strings.Compare(parts[2], "product") == 0 {
		matched, err := regexp.MatchString(`^[0-9A-Z]{10}$`, parts[3])
		if err != nil {
			return "", err
		}
		if matched {
			return parts[3], nil
		}
	}

	return "", nil
}
