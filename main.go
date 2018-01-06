package main

import (
	"io/ioutil"
	"os"
	"time"

	log "github.com/inconshreveable/log15"
	yaml "gopkg.in/yaml.v2"
)

//go:generate go-assets-builder -s="/data" -o holiday.go data

var (
	// DateFormat is "yyyy-MM-dd"
	DateFormat = "2006-01-02"

	// JST is Japan's time zone
	JST = time.FixedZone("Asia/Tokyo", 9*60*60)
)

func main() {
	os.Exit(run())
}

func run() int {
	f, err := Assets.Open("/holidays.yml")
	if err != nil {
		log.Error("Open", "Error", err)
		return 2
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Error("Read", "Error", err)
		return 2
	}

	holiday := make(map[interface{}]interface{})
	if err = yaml.Unmarshal(b, &holiday); err != nil {
		log.Error("YAML Unmarshal", "Error", err)
		return 2
	}

	now := time.Now().In(JST)
	d := now.Format(DateFormat)

	if _, ok := holiday[d]; ok {
		return 0
	}

	if now.Weekday() == time.Saturday || now.Weekday() == time.Sunday {
		return 0
	}

	return 1
}
