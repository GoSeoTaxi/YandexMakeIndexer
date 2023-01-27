package main

import (
	"YandexMakeIndex/internal/config"
	"YandexMakeIndex/internal/makerQueue"
	"YandexMakeIndex/internal/utils/readSiteMap"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	fmt.Println(`Привет это программа для ускорения индексации в яндексе`)
	fmt.Println(`Все прикольные программы я складываю тут = https://t.me/dev_etc`)
	fmt.Println(`Заказная разработка автоматизации и ботов = https://t.me/backender_dev`)

	fmt.Println(`version CLI XML`)
	cfg, err := config.InitConfig()
	if err != nil {
		fmt.Printf("No load config %v", err)
	}

	checkLic()

	//siteMap := "https://etecom.ru/sitemap.xml"

	fmt.Println(`Начинаем считывать карту сайта ` + cfg.URLSiteMap)
	makerQueue.DataInput(readSiteMap.ReadSitemapXML(cfg.URLSiteMap))
	log.Println(`Завершил свою работу`)
}

func checkLic() {

	deadLine := time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC)

	if time.Now().After(deadLine) {
		fmt.Println(`Ошибка системы, напишите мне в Телеграмм https://t.me/backender_dev`)
		os.Exit(1)
	}

}
