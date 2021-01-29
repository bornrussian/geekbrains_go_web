//
//  Задача:
//
//  1. Подключите логирование к своему проекту
//	2. Выберите параметры, которые стоит вынести в конфигурационный файл (по своему усмотрению). Напишите код для загрузки конфигурации и используйте полученные данные в проекте.
//	*3. Задеплоить свой проект на хостинге
//

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	flagConfigPath := flag.String("c", "./config.yaml", "config file with yaml format")
	flag.Parse()
	conf, err := ReadConfig(*flagConfigPath)
	if err != nil {
		panic(fmt.Sprintf("config file read failed: %s", err))
	}
	lg, err := ConfigureLogger(&conf.Logger)
	if err != nil {
		panic(fmt.Sprintf("can't configure logger: %s", err))
	}

	// HTTP START
	serv := NewServer(context.TODO(), lg)
	serv.SetConfig(conf.Server)
	serv.Start()
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt, os.Kill)
	<-stopChan
	serv.Stop()
}
