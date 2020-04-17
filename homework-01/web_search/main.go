//
// Задача:
//
//  Напишите функцию, которая будет получать на вход строку с поисковым запросом (string) и
//  массив ссылок на страницы, по которым стоит произвести поиск ([]string). Результатом работы
//  функции должен быть массив строк со ссылками на страницы, на которых обнаружен
//  поисковый запрос. Функция должна искать точное соответствие фразе в тексте ответа от
//  сервера по каждой из ссылок.
//

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

// Проверяем на одном конкретном сайте, есть ли там искомый контент
func checkOneSiteHasContent (site string, content string) bool {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Error!:", err)
		return false
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if strings.Contains(string(body), content) {
		return true
	}
	return false
}

// Берём массив со ссылками на сайт, и проверяем, не найдется ли там content
// Результат - массив с ссылками, где нашлось.
type channelMessage struct {
	site string
	found bool
}
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

func main () {
	flagUrls := flag.String("urls", "https://google.ru,https://mail.ru/,https://lenta.ru,https://github.com", "URLs separated by comma (,)")
	flagSearch := flag.String("search", "google-analytics", "String to find to")
	flag.Parse()

	sites := strings.Split(*flagUrls,",")
	fmt.Printf("Ищем фразу '%v' на сайтах: %v\n", *flagSearch, sites)
	fmt.Printf("Найдено на сайтах: %v\n",whichSitesHasContent(sites, *flagSearch))
}
