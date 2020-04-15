//
// Задача:
// * Напишите функцию, которая получает на вход публичную ссылку на файл с «Яндекс.Диска»
// и сохраняет полученный файл на диск пользователя.
//

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// Скачиваем ссылку в файл на диске
func urlToFile (url string, destFileName string) bool {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http.Get error:", err)
		return false
	} else {
		defer resp.Body.Close()

		out, err := os.Create(destFileName)
		if err != nil {
			fmt.Println("os.Create error:", err)
			return false
		} else {
			defer out.Close()

			_, err = io.Copy(out, resp.Body)
			if err != nil {
				fmt.Println("io.Copy error:", err)
				return false
			} else {
				return true
			}
		}
	}
}

// документация по Yandex API :
// https://yandex.ru/dev/disk/api/reference/public-docpage/
var yandex_api_url = "https://cloud-api.yandex.net/v1/disk/public/resources/download?public_key="

// Скачиваем Яндекс.Диск-ссылку в файл на диске
func yandexDiskToFile (yandexDiskUrl string, destFileName string) bool {
	// отправляем яндексу запрос на получение ссылки на скачивание файла 'yandexDiskUrl'
	linkResp, err := http.Get(yandex_api_url + url.QueryEscape(yandexDiskUrl))
	if err != nil {
		fmt.Println("http.Get error:", err)
		return false
	} else {
		defer linkResp.Body.Close()
		linkBody, _ := ioutil.ReadAll(linkResp.Body)
		if linkResp.StatusCode == 200 {
			fmt.Println("Ответ от API:", string(linkBody))
			var jsonInterface interface{}
			err := json.Unmarshal(linkBody, &jsonInterface)
			if err != nil {
				fmt.Println("json.Unmarshal error:", err)
				return false
			} else {
				jsonMap := jsonInterface.(map[string]interface{})
				link := jsonMap["href"].(string)
				// скачаем его в файл destFileName
				// теперь у нас есть ссылка на файл: link
				return urlToFile(link, destFileName)
			}
		} else {
			fmt.Println("Yandex API error:", string(linkBody))
			return false
		}
	}
}

func main () {
	if yandexDiskToFile("https://yadi.sk/i/eK0nO8P0SPfWyg","C:\\temp\\golang.png") {
		fmt.Println("Получилось скачать!")
	} else {
		fmt.Println("Что-то пошло не так :(")
	}
}
