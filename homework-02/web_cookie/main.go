//
//	Задача:
//
//	Напишите два роута: один будет записывать информацию в Cookie (например, имя),
//	а второй — получать ее и выводить в ответе на запрос.
//

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const COOKIE_NAME = "geekbrains_go_web_homework02_cookie_name"

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", rootHandle)
	router.HandleFunc("/setcookie_post_method", setCookiePostHandle)
	router.HandleFunc("/setcookie_get_method", setCookieGetHandle)
	router.HandleFunc("/getcookie", getCookieHandle)
	fmt.Println("Starting web-server at *:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Функция, котороя продемонстрирует работу функций в браузере
func rootHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(
		"<html>" +
			"<header>" +
			"</header>" +
			"<body>" +

			"<form action='/setcookie_post_method' method='post'>" +
			"<input type='text' value='12345678' name='cookie'><br>" +
			"<input type='submit' value='set cookie with post method'>" +
			"</form>" +

			"<hr>" +

			"<form action='/setcookie_get_method' method='get'>" +
			"<input type='text' value='87654321' name='cookie'><br>" +
			"<input type='submit' value='set cookie with get method'>" +
			"</form>" +

			"<hr>" +

			"<form action='/getcookie' method='get'>" +
			"<input type='submit' value='get cookie'>" +
			"</form>" +

			"</body>" +
			"</html>"))
}

func setCookiePostHandle(w http.ResponseWriter, r *http.Request) {
	setCookieHandle(w, r,true)
}
func setCookieGetHandle(w http.ResponseWriter, r *http.Request) {
	setCookieHandle(w, r,false)
}
func setCookieHandle(w http.ResponseWriter, r *http.Request, isPostForm bool) {
	if isParsed := r.ParseForm(); isParsed != nil {
		w.Write([]byte("something is wrong while post html form parsing"))
	} else {
		var value string
		if isPostForm {
			value = r.PostFormValue("cookie")
		} else {
			value = r.FormValue("cookie")
		}
		expire := time.Now().AddDate(0, 1, 0)
		cookie := http.Cookie{
			Name:    COOKIE_NAME,
			Value:   value,
			Expires: expire,
		}
		http.SetCookie(w, &cookie)
		w.Write([]byte("ok"))
	}
}

func getCookieHandle(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(COOKIE_NAME)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Write([]byte(cookie.Value))
	}
}

