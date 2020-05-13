// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-05-13 18:24:29.3960766 +0700 +07 m=+0.087061201

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "wwwIndex - показываем корневую страницу сайта",
                "tags": [
                    "handlers"
                ]
            }
        },
        "/delete": {
            "get": {
                "description": "wwwDelete - удаляем запись блога из базы данных",
                "tags": [
                    "handlers"
                ]
            }
        },
        "/docs/swagger.json": {
            "get": {
                "description": "Returns swagger.json docs",
                "tags": [
                    "system"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/edit": {
            "get": {
                "description": "wwwEdit - форма редактирования записи блога",
                "tags": [
                    "handlers"
                ]
            }
        },
        "/new": {
            "get": {
                "description": "wwwNew - показываем форму, в которой можно внести данные для нового поста блога",
                "tags": [
                    "handlers"
                ]
            }
        },
        "/push": {
            "post": {
                "description": "wwwPush - вставляем запись блога в базу данных",
                "tags": [
                    "handlers"
                ]
            }
        },
        "/view": {
            "get": {
                "description": "wwwView - отображает одну конкретную запись блога",
                "tags": [
                    "handlers"
                ]
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
