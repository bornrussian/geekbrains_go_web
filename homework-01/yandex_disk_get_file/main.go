//
// Задача:
// * Напишите функцию, которая получает на вход публичную ссылку на файл с «Яндекс.Диска»
// и сохраняет полученный файл на диск пользователя.
//

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
)

// Скачиваем ссылку в файл на диске
func urlToFile (url string, destFileName string) (bool, string) {

	var filename string
	if destFileName == "parse_from_url" {
		re := regexp.MustCompile(`filename=.*`)
		filename = re.FindString(url)
		re = regexp.MustCompile(`filename=`)
		filename = re.ReplaceAllString(filename, "");
		re = regexp.MustCompile(`&.*`)
		filename = re.ReplaceAllString(filename, "");
		fmt.Println("Имя файла из URL =", filename)
	} else {
		filename = destFileName
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http.Get error:", err)
		return false, ""
	} else {
		defer resp.Body.Close()

		out, err := os.Create(filename)
		if err != nil {
			fmt.Println("os.Create error:", err)
			return false, ""
		} else {
			defer out.Close()

			_, err = io.Copy(out, resp.Body)
			if err != nil {
				fmt.Println("io.Copy error:", err)
				return false, ""
			}

			return true, filename
		}
	}
}

// документация по Yandex API :
// https://yandex.ru/dev/disk/api/reference/public-docpage/
var yandex_api_url = "https://cloud-api.yandex.net/v1/disk/public/resources/download?public_key="

// Скачиваем Яндекс.Диск-ссылку в файл на диске
func yandexDiskToFile (yandexDiskUrl string, destFileName string) (bool, string) {
	// отправляем яндексу запрос на получение ссылки на скачивание файла 'yandexDiskUrl'
	linkResp, err := http.Get(yandex_api_url + url.QueryEscape(yandexDiskUrl))
	if err != nil {
		fmt.Println("http.Get error:", err)
		return false, ""
	} else {
		defer linkResp.Body.Close()
		linkBody, _ := ioutil.ReadAll(linkResp.Body)
		if linkResp.Status == "200 OK" {
			fmt.Println("Ответ от Яндекс API:", string(linkBody))

			var jsonInterface interface{}
			err := json.Unmarshal(linkBody, &jsonInterface)
			if err != nil {
				fmt.Println("json.Unmarshal error:", err)
				return false, ""
			}

			jsonMap := jsonInterface.(map[string]interface{})
			link := jsonMap["href"].(string)
			// скачаем его в файл destFileName
			// теперь у нас есть ссылка на файл: link
			return urlToFile(link, destFileName)
		} else {
			fmt.Println("Yandex API error:", string(linkBody))
			return false, ""
		}
	}
}

func main () {
	flagUrl := flag.String("url", "https://yadi.sk/i/eK0nO8P0SPfWyg", "Link to Yandex.Disk public shared file")
	flagFilename := flag.String("filename", "parse_from_url", "Filename")
	flag.Parse()

	fmt.Println("Утилита для скачивания файла с Яндекс.Диска")
	fmt.Println("Принимаемые аргументы командной строки:")
	fmt.Println("  --url \"https://yadi.sk/i/eK0nO8P0SPfWyg\" # Ссылка на публичнорасшаренный файл")
	fmt.Println("  --filename \"filename.png\" # Имя файла, которое создадим на диске. Не заполняйте, если нужно прочитать имя файла из Яндекс.Диск")
	fmt.Println("")

	ok, filename := yandexDiskToFile(*flagUrl, *flagFilename)
	if ok {
		fmt.Println("Получилось скачать файл", filename)
	} else {
		fmt.Println("Что-то пошло не так :(")
	}
}
