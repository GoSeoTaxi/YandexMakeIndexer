package readSiteMap

func ReadSitemapXML(pathSiteMap string) (out []string) {

	urls, _ := fetchUrls(pathSiteMap)
	/*	if err != nil {
		fmt.Println(err)
	}*/

	for _, url := range urls {

		if (url[len(url)-4 : len(url)-0]) == ".xml" {
			urlsOut1 := ReadSitemapXML(url)
			for _, url1 := range urlsOut1 {
				out = append(out, url1)
			}

		} else {
			out = append(out, url)
		}

		//fmt.Println(`+++111`)
		//time.Sleep(time.Second)

	}

	return out
}
