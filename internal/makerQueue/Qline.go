package makerQueue

import (
	"YandexMakeIndex/internal/browserWorker"
)

func queueMaker(c chan string) error {

	var err error
	var closeC bool

	for {
		closeC, err = browserWorker.T1Open(c)
		if closeC {
			break
		}
	}

	return err
}
