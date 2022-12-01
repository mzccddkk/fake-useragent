package useragent

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

var ua = loadData()
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

type useragent struct {
	data     map[string][]string
	isLoaded bool
	current  []string
}

func loadData() *useragent {
	u := &useragent{}
	l, err := u.loadData()
	if err != nil {
		log.Println("load data error:", err)
		return u
	}
	return l
}

func (u *useragent) get(key string) *useragent {
	u.current = u.data[key]
	return u
}

func (u *useragent) getAll() *useragent {
	var all []string
	for _, v := range u.data {
		all = append(all, v...)
	}
	u.current = all
	return u
}

func (u *useragent) random() string {
	current := u.current
	n := len(current)
	if n <= 0 {
		return ""
	}
	return current[r.Intn(n)]
}

// loadData from json file
func (u *useragent) loadData() (*useragent, error) {
	if u.isLoaded {
		return u, nil
	}
	var data map[string][]string
	dir, err := currentAbsolutePath()
	if err != nil {
		return nil, err
	}
	if !strings.HasSuffix(dir, "/") {
		dir = dir + "/"
	}
	file, err := ioutil.ReadFile(dir + "data/useragent.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}
	u.data = data
	u.isLoaded = true
	return u, nil
}
