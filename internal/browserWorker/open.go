package browserWorker

import (
	"YandexMakeIndex/internal/etc"
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

func T1Open(c chan string) (closeC bool, err error) {
	closeC = false

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		//	chromedp.Flag("user-data-dir", pathChrome),
		//	chromedp.Flag("no-sandbox", true),
		chromedp.ExecPath(etc.PathWindowsYaBrowser()),
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.Flag("restore-on-startup", false),
		chromedp.Flag("no-first-run", true),
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-extensions", false),
		chromedp.Flag("incognito", true),

		//chromedp.Flag("DefaultChromePath", etc.PathWindowsYaBrowser()),
		//	chromedp.WindowSize(912, 1368),
		chromedp.WindowSize(1366, 768),
	)

	// create context.allocCtx
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// create context.ctx
	ctx, cancel2 := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel2()

	//Не применимо тут

	time.AfterFunc(120*time.Second, func() {
		//	fmt.Println("Завершаем зависший процесс")
		cancel2()
	})

	//var res *runtime.RemoteObject

	//,
	//	chromedp.Navigate(sURL),
	//	wait5S(),
	//	chromedp.WaitNotVisible("div.loader"),

	/*
		func() {

			err = chromedp.Run(ctxNW,
				//	chromedp.Evaluate(`window.open("https://ya.ru", "", "resizable,scrollbars,status")`, &res),
				chromedp.Evaluate(`window.open("https://google.ru", "", "resizable,scrollbars,status")`, &res),
				chromedp.Sleep(10*time.Second),
				//	chromedp.WaitNotVisible("div.loader"),
				//chromedp.E

			)
			if err != nil {
				fmt.Println(err)
			}

			chromedp.WithPollingTimeout(2 * time.Second)
			fmt.Println(`START 78`)
			time.Sleep(11 * time.Second)
			fmt.Println(`++76`)
			err = chromedp.Cancel(ctxNW)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(`==81`)

		}()
	*/

	err = chromedp.Run(ctx)
	if err != nil {
		//	fmt.Println(`Err Start Browser`)
		//	fmt.Println(err)
	}
	func() {

		time.Sleep(1 * time.Second)

		//	for urlList := range c {

		for {
			urlList, ok := <-c
			//fmt.Println(`Читаем канал`)
			//fmt.Println(urlList)
			//fmt.Println(ok)
			if !ok {
				closeC = true
				break
				//fmt.Println(`Закончил чтение из канала`)
			}

			err = chromedp.Run(ctx,
				chromedp.Navigate(urlList),
				chromedp.Sleep(1*time.Second),
				//chromedp.Evaluate(`window.open("https://expired.badssl.com/")`, &res),
			)
			if err != nil {
				//fmt.Println(err)
				//fmt.Println(`ошибка в ранже браузера`)
				break
			}
		}
		//fmt.Println(`Закрываем контекст`)

		chromedp.Cancel(ctx)
	}()
	/*
			func() {

				err = chromedp.Run(ctx,
					//	chromedp.Evaluate(`window.open("https://ya.ru", "", "resizable,scrollbars,status")`, &res),
					chromedp.Evaluate(`window.open("https://google.ru", "", "resizable,scrollbars,status")`, &res),
					chromedp.Sleep(10*time.Second),

					//	chromedp.WaitNotVisible("div.loader"),
					//chromedp.E

				)
				if err != nil {
					fmt.Println(err)
				}
		chromedp.Stop()
			}()

			func() {

				err = chromedp.Run(ctx,
					//	chromedp.Evaluate(`window.open("https://ya.ru", "", "resizable,scrollbars,status")`, &res),
					chromedp.Evaluate(`window.open("https://mail.ru", "", "resizable,scrollbars,status")`, &res),

					//	chromedp.WaitNotVisible("div.loader"),
					//chromedp.E
				)
				if err != nil {
					fmt.Println(err)
				}

			}()
	*/
	return closeC, err
}
