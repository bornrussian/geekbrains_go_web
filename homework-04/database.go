package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func MysqlRealEscapeString(value string) string {
	replace := map[string]string{"\\":"\\\\", "'":`\'`, "\\0":"\\\\0", "\n":"\\n", "\r":"\\r", `"`:`\"`, "\x1a":"\\Z"}
	for b, a := range replace {
		value = strings.Replace(value, b, a, -1)
	}
	return value
}

func (b *Blog) GetAllJokesFromDB() ([]BlogEntry, error) {
	var res []BlogEntry

	rows, err := b.DB.Query("SELECT * FROM jokes ORDER BY date DESC")
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		joke := BlogEntry{}
		err := rows.Scan(&joke.ID, &joke.Autor, &joke.Date, &joke.Header, &joke.Content)
		if err != nil {
			log.Println(err)
			continue
		}
		res = append(res, joke)
	}
	return res, nil
}

func (b *Blog) GetSingleJokeFromDB (id int) (BlogEntry, error) {
	var res BlogEntry

	err := b.DB.QueryRow(fmt.Sprintf("SELECT * FROM jokes WHERE id = %v", id)).
		Scan(&res.ID, &res.Autor, &res.Date, &res.Header, &res.Content)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (b *Blog) PushJokeToDB (joke BlogEntry) error {
	if joke.ID == "" {
		// Если ID пустой, то вставляем новую joke в базу данных
		insert, err := b.DB.Prepare("INSERT INTO jokes (id, autor, date, header, content) VALUES (NULL, ?, ?, ?, ?);")
		if err != nil {
			log.Println(err)
		} else {
			res, err := insert.Exec(
				MysqlRealEscapeString(joke.Autor), MysqlRealEscapeString(joke.Date),
				MysqlRealEscapeString(joke.Header),	MysqlRealEscapeString(joke.Content))
			if err != nil {
				log.Println(res)
			}
			defer insert.Close()
		}
	} else {
		// Если ID не пустой, то обновляем в базе данных строчку joke с таким ID
		// Убедимся, что ID - это число
		if id, errConv := strconv.Atoi(joke.ID); errConv == nil {
			update, errUpdate := b.DB.Prepare("UPDATE jokes SET autor = ?, date = ?, header = ?, content = ? WHERE jokes.id = ?;")
			if errUpdate != nil {
				log.Println(errUpdate)
			} else {
				res, errExec := update.Exec(
					MysqlRealEscapeString(joke.Autor), MysqlRealEscapeString(joke.Date),
					MysqlRealEscapeString(joke.Header),	MysqlRealEscapeString(joke.Content), id)
				if errExec != nil {
					log.Println(res)
				}
				defer update.Close()
			}
		} else {
			return errors.New("ID is not number")
		}
	}
	return nil
}

func (b *Blog) DeleteJokeAtDB(id int) {
	delete, errDelete := b.DB.Prepare("DELETE FROM jokes WHERE id = ? LIMIT 1;")
	if errDelete != nil {
		log.Println(errDelete)
	} else {
		res, errExec := delete.Exec(id)
		if errExec != nil {
			log.Println(res)
		}
		defer delete.Close()
	}
}
