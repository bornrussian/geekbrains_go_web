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
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Берём массив со ссылками на сайт, и проверяем, не найдется ли там content
// Результат - массив с ссылками, где нашлось.

func whichSitesHaveContent (sites []string, content string) []string {
	result := []string{}
	for _, site := range sites {
		resp, err := http.Get(site)
		if err != nil {
			fmt.Println("Error!:", err)
		} else {
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			//fmt.Println("processing site: ", site, "...")
			if strings.Contains(string(body), content) {
				result = append(result, site)
			}
		}
	}
	return result
}

func main () {
	sites := []string{"https://google.ru", "https://yandex.ru", "https://mail.ru/", "https://lenta.ru", "https://github.com"}
	fmt.Println("search sites are:", sites)
	fmt.Println("'google-analytics' found at: ",whichSitesHaveContent(sites, "google-analytics"))
}
