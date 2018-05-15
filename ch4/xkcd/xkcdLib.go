package xkcd

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Comic struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

func getUrlForComic(id int) string {
	base := "https://xkcd.com/###/info.0.json"
	url := strings.Replace(base, "###", strconv.Itoa(id), 1)
	return url
}

func getFileNameForComic(id int) string {
	workDir := "/tmp"
	return workDir + "/" + strconv.Itoa(id) + ".json"
}

/*
func GetAll() {
	var i int
	for {
		comic, err := GetComic(i)
		if err != nil {
			log.Printf("Error getting url %s: %err\n", url, err)
			break
		}

		// save the contents

		ioutil.Copy(getFileNameForComic(id), comic)
		if i > 5 {
			break
		}
		i++

	}
}
*/
func getComicFromUrl(url string) (Comic, error) {
	var result Comic
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error getting url %s: %err\n", url, err)
		resp.Body.Close() // do not leak descriptors. There has got to be a better pattern
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("Error decoding comic for url %s: %s (%s)", url, err, resp.Body)
		resp.Body.Close() // do not leak descriptors. There has got to be a better pattern
	}

	return result, err
}

// TODO: implement for real so that we don't retrieve the same comic multiple times
func GetComic(id int) (Comic, error) {
	var c Comic
	f, err := os.Open(getFileNameForComic(id))
	if err != nil {
		url := getUrlForComic(id)
		c, err = getComicFromUrl(url)
	}
	json.NewDecoder(bufio.NewReader(f)).Decode(&c)
	return c, err
}

/*
func SimpleIndex(c &Comic) map([string]Comic){
	// The index is comprised of single words. Pretty lame, but the point is to exercise code not make a search engine
	t =

}
*/
