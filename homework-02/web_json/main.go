package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

// Структура, которую мы будем ожидать от пользователя
type jsonQueryStruct struct {
	Search string `json:"search"`
	Sites []string `json:"sites"`
}

// Структура, которой мы ответим пользователю
type jsonReplyStruct struct {
	Sites []string `json:"found-at"`
}

// Стартуем веб-сервер
func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", rootHandle)		// тут мы подготовим post-запрос с json-ом
	router.HandleFunc("/query", queryHandle)	// тут мы обработаем json-запрос и ответим
	fmt.Println("Starting web-server at *:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Функция, котороя поможет подготовить POST-запрос с JSON-ом
func rootHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(
		"<html>" +
		"<header>" +
		"</header>" +
		"<body>" +
			"make 'search' on 'sites':<br><br>" +
			"<form action='/query' method='post'>" +
				"<textarea rows='10' cols='45' name='json'>" +
					"{\n" +
					"  \"search\":\"google-analytics\",\n" +
					"  \"sites\": [\n" +
					"    \"https://mail.ru\",\n" +
					"    \"https://github.com\",\n" +
					"    \"https://lenta.ru\"\n"+
					"  ]\n" +
					"}\n"	+
				"</textarea><br><br>" +
				"<input type='submit'>" +
			"</form>" +
		"</body>" +
		"</html>"))
}

// Функция, которая принимает POST-запрос с JSON-ом
func queryHandle(w http.ResponseWriter, r *http.Request) {
	if isParsed := r.ParseForm(); isParsed != nil {

		w.Write([]byte("{ \"error\": \"something is wrong while post html form parsing\" }"))
	} else {
		jsonQueryText := r.PostFormValue("json")
		jsonQueryCode := jsonQueryStruct{}
		if isUnmarsh := json.Unmarshal([]byte(jsonQueryText), &jsonQueryCode); isUnmarsh != nil {
			w.Write([]byte("{ \"error\": \"something is wrong while json unmarshall\" }"))
		} else {
			jsonReplyCode := jsonReplyStruct{}
			jsonReplyCode.Sites = whichSitesHasContent(jsonQueryCode.Sites, jsonQueryCode.Search)
			jsonReplyText, isMarsh := json.MarshalIndent(jsonReplyCode,"","\t")
			if isMarsh != nil {
				w.Write([]byte("{ \"error\": \"something is wrong while json marshall\" }"))
			} else {
				// Сообщаем ответ в виде JSON
				w.Write([]byte(jsonReplyText))
			}
		}
	}
}

// Проверяем на одном конкретном сайте, есть ли там искомый контент
func checkOneSiteHasContent (site string, needle string) bool {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Error!:", err)
		return false
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if strings.Contains(string(body), needle) {
		return true
	}
	return false
}

// Берём массив со ссылками на сайт, и проверяем, не найдется ли там content
// Результат - массив с ссылками, где нашлось.
func whichSitesHasContent (sites []string, needle string) []string {
	result := []string{}
	mutex := &sync.Mutex{}

	company := &sync.WaitGroup{}
	for id := 0; id < len(sites); id++ {
		company.Add(1)
		go worker(company, mutex, sites[id], needle, &result)
	}
	company.Wait()
	return result
}
func worker(wg *sync.WaitGroup, mu *sync.Mutex, site string, needle string, found *[]string) {
	defer func () {
		//fmt.Println("Worker закончил проверять сайт", site)
		wg.Done()
	} ()
	if checkOneSiteHasContent(site, needle) {
		//fmt.Println("Worker нашёл совпадение на сайте", site)
		mu.Lock()
		*found = append(*found, site)
		mu.Unlock()
	} else {
		//fmt.Println("Worker не нашёл совпадение на сайте", site)
	}
}
