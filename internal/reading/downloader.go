package reading

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/guergeiro/fator-conversao-gas-portugal/internal/interval"
)

func DownloadCsv(interval interval.Interval) (io.ReadCloser, error) {
	page, cookies, err := getInitialPage()
	if err != nil {
		return nil, err
	}
	formUri, err := extractFormActionUri(page)
	if err != nil {
		return nil, err
	}
	csvUri, err := searchForIntervalCsv(formUri, cookies, interval)
	if err != nil {
		return nil, err
	}
	return getCsv(csvUri, cookies)
}

func getInitialPage() (*goquery.Document, []*http.Cookie, error) {
	res, err := http.Get("https://www.ign.ren.pt/web/guest/monitorizacao-horaria-qualidade")
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	return doc, res.Cookies(), err
}

func searchForIntervalCsv(
	formUri string,
	cookies []*http.Cookie,
	interval interval.Interval,
) (string, error) {

	layout := "02/01/2006"
	form := url.Values{}
	form.Add("servicePoint", "")
	form.Add("startTime", interval.StartTime().Format(layout))
	form.Add("stopTime", interval.StopTime().Format(layout))

	req, err := http.NewRequest(
		"POST",
		formUri,
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	for _, c := range cookies {
		req.AddCookie(c)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}
	return extractCsvHref(doc)
}

func getCsv(csvUri string, cookies []*http.Cookie) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", csvUri, nil)
	if err != nil {
		return nil, err
	}
	for _, c := range cookies {
		req.AddCookie(c)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func extractFormActionUri(doc *goquery.Document) (string, error) {
	// Load the HTML document
	uri, exists := doc.Find("form#qualityReadingsSearchCriteria").First().Attr("action")
	if exists == false {
		return "", fmt.Errorf("Cant find form action")
	}
	return uri, nil
}

func extractCsvHref(doc *goquery.Document) (string, error) {
	uri, exists := doc.Find("span.csvIcon").Parent().Attr("href")
	if exists == false {
		return "", fmt.Errorf("Cant find csv uri")
	}

	return uri, nil
}
