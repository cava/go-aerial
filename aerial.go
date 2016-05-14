package main

import "net/http"
import "fmt"
import "io/ioutil"
import "encoding/json"
import "os/exec"
import "math/rand"

const catalogUrl = "http://a1.phobos.apple.com/us/r1000/000/Features/atv/AutumnResources/videos/entries.json"

type AerialEntry struct {
	Id     string `json:"id"`
	Assets []AerialAsset
}
type AerialAsset struct {
	Url       string `json:"url"`
	Label     string `json:"accessibilityLabel"`
	Type      string `json:"type"`
	Id        string `json:"id"`
	TimeOfday string `json:"timeOfDay"`
}

func GetEntries() []AerialEntry {
	res, err := http.Get(catalogUrl)
	if err != nil {
		fmt.Println(err.Error())
	}
	buf, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	entries := make([]AerialEntry, 0)
	err = json.Unmarshal(buf, &entries)
	if err != nil {
		fmt.Println(err.Error())
	}
	return entries
}

func Random(entries []AerialEntry) AerialAsset {
	fmt.Println(rand.Intn(len(entries)))
	assets := entries[rand.Intn(len(entries))].Assets
	return assets[rand.Intn(len(assets))]
}

func main() {
	entries := GetEntries()
	randomAsset := Random(entries)
	fmt.Println(randomAsset)
	cmd := exec.Command("mpv", randomAsset.Url)
	cmd.Start()
}
