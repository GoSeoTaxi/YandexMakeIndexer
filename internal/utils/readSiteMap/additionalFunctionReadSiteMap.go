package readSiteMap

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SitemapIndexMaps struct {
	Locations []string `xml:"sitemap>loc"`
}

type SitemapIndexUrls struct {
	LocationsU []string `xml:"url>loc"`
}

func (s *SitemapIndexMaps) getUrlsM() []string {
	return s.Locations
}

func (s *SitemapIndexUrls) getUrlsU() []string {
	return s.LocationsU
}

func fetchUrls(url string) (out []string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(`Не сработал GET`)
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var sM SitemapIndexMaps
	xml.Unmarshal(bytes, &sM)

	out = sM.getUrlsM()
	if len(out) > 1 {
		return sM.getUrlsM(), nil
	}

	var sU SitemapIndexUrls
	xml.Unmarshal(bytes, &sU)
	return sU.getUrlsU(), nil
}
