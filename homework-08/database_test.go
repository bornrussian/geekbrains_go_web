package main

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
	"context"
	"github.com/google/uuid"
)

const DATABASE_LOCATION = "mongodb://127.0.0.1:27017"
const DATABASE_DBNAME = "geekbrains"

func TestInsert(t *testing.T) {
	// Подключаемся в базе данных
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(DATABASE_LOCATION))
	if err != nil {
		t.Error("Could not connect to database",DATABASE_LOCATION)
		return
	}
	db := client.Database(DATABASE_DBNAME)

	// Генерируем случайный пост
	uuid, _ := uuid.NewRandom()
	testuuid := "test-" + uuid.String()
	post := &Post{
		Autor: testuuid,
		Date: testuuid,
		Header: testuuid,
		Content: testuuid,
	}

	// Пробуем вставить его в базу данных
	if err := post.Insert(ctx, db); err != nil {
		t.Error("Insert failed", err)
	} else {
		// Если получилось вставить запись, то удаляем её за собой, чтобы не мусорить
		post.Delete(ctx, db)
	}
}

func TestGetPost(t *testing.T) {
	// Подключаемся в базе данных
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(DATABASE_LOCATION))
	if err != nil {
		t.Error("Could not connect to database",DATABASE_LOCATION)
		return
	}
	db := client.Database(DATABASE_DBNAME)

	// Генерируем случайный пост
	uuid, _ := uuid.NewRandom()
	testuuid := "test-" + uuid.String()
	origin := &Post{
		Autor: testuuid,
		Date: testuuid,
		Header: testuuid,
		Content: testuuid,
	}

	// Пробуем вставить его в базу данных
	if err := origin.Insert(ctx, db); err != nil {
		t.Error("Insert failed", err)
	} else {
		// Если получилось вставить запись, то пробуем его найти через GetPost
		search, errSearch := GetPost(ctx,db, origin.ID)
		if errSearch != nil {
			t.Error("GetPost failed", errSearch)
		} else {
			if search.Content != origin.Content {
				t.Error("GetPost search.Content != origin.Content")
			} else {
				// Все окей, убираем за собой тестовую запись, чтобы не мусорить
				origin.Delete(ctx, db)
			}
		}
	}
}

func TestUpdate(t *testing.T) {
	// Подключаемся в базе данных
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(DATABASE_LOCATION))
	if err != nil {
		t.Error("Could not connect to database",DATABASE_LOCATION)
		return
	}
	db := client.Database(DATABASE_DBNAME)

	// Генерируем случайный пост
	uuid, _ := uuid.NewRandom()
	testuuid := "test-" + uuid.String()
	post := &Post{
		Autor: testuuid,
		Date: testuuid,
		Header: testuuid,
		Content: testuuid,
	}

	// Пробуем вставить его в базу данных
	if err := post.Insert(ctx, db); err != nil {
		t.Error("Insert failed", err)
	} else {
		// Если получилось вставить запись, то пробуем метод Update

		post.Header = testuuid + "-updated"
		errUpdate := post.Update(ctx,db)
		if errUpdate != nil {
			t.Error("Update failed", errUpdate)
			return
		}

		// Если получилось вставить запись, то удаляем её за собой, чтобы не мусорить
		post.Delete(ctx, db)
	}
}

func TestDelete(t *testing.T) {
	// Подключаемся в базе данных
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(DATABASE_LOCATION))
	if err != nil {
		t.Error("Could not connect to database",DATABASE_LOCATION)
		return
	}
	db := client.Database(DATABASE_DBNAME)

	// Генерируем случайный пост
	uuid, _ := uuid.NewRandom()
	testuuid := "test-" + uuid.String()
	post := &Post{
		Autor: testuuid,
		Date: testuuid,
		Header: testuuid,
		Content: testuuid,
	}

	// Пробуем вставить его в базу данных
	if err := post.Insert(ctx, db); err != nil {
		t.Error("Insert failed", err)
	} else {
		// Если получилось вставить запись, то удаляем её за собой
		errDelete := post.Delete(ctx, db)
		if errDelete != nil {
			t.Error("Delete failed", errDelete)
		}
	}
}

func TestGetPosts(t *testing.T) {
	// Подключаемся в базе данных
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(DATABASE_LOCATION))
	if err != nil {
		t.Error("Could not connect to database",DATABASE_LOCATION)
		return
	}
	db := client.Database(DATABASE_DBNAME)

	// Генерируем случайный пост
	uuid, _ := uuid.NewRandom()
	testuuid := "test-" + uuid.String()
	post := &Post{
		Autor: testuuid,
		Date: testuuid,
		Header: testuuid,
		Content: testuuid,
	}

	// Пробуем вставить его в базу данных
	if err := post.Insert(ctx, db); err != nil {
		t.Error("Insert failed", err)
	} else {
		// Берём список всех постов и ищем там наш
		posts, errGet := GetPosts(ctx, db)
		if errGet != nil {
			t.Error("GetPosts error",errGet)
			return
		}

		found := false
		for _,p := range posts {
			if p.Content == post.Content {
				found = true
			}
		}
		if !found {
			t.Error("GetPosts test post was not found")
		}

		// Если получилось вставить запись, то удаляем её за собой
		post.Delete(ctx, db)
	}
}

func TestFind(t *testing.T) {
	// Подключаемся в базе данных
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(DATABASE_LOCATION))
	if err != nil {
		t.Error("Could not connect to database",DATABASE_LOCATION)
		return
	}
	db := client.Database(DATABASE_DBNAME)

	// Генерируем случайный пост
	uuid, _ := uuid.NewRandom()
	testuuid := "test-" + uuid.String()
	post := &Post{
		Autor: testuuid,
		Date: testuuid,
		Header: testuuid,
		Content: testuuid,
	}

	// Пробуем вставить его в базу данных
	if err := post.Insert(ctx, db); err != nil {
		t.Error("Insert failed", err)
	} else {
		// Ищем наш пост через метод Find
		posts, errFind := Find(ctx, db, "content", testuuid)
		if errFind != nil {
			t.Error("Find error",errFind)
			return
		}

		found := false
		for _,p := range posts {
			if p.Content == post.Content {
				found = true
			}
		}
		if !found {
			t.Error("Find test post was not found")
		}

		// Если получилось вставить запись, то удаляем её за собой
		post.Delete(ctx, db)
	}
}