package makerQueue

import (
	"fmt"
	"log"
)

func DataInput(rowLine []string) {

	log.Printf(`В карте сайта - элементов= %v `, len(rowLine))
	fmt.Println(`Обработка...`)

	cQue := make(chan string)
	go func() {

		for _, elemList := range rowLine {
			//fmt.Println(`Пишу в Канал`)
			//fmt.Println(elemList)
			cQue <- elemList
		}

		close(cQue)

	}()

	queueMaker(cQue)

}
