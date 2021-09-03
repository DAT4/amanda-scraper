package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func urlize(k string) string {
	return fmt.Sprintf("http://integral.esac.esa.int/isocweb/%s", strings.Split(k, "'")[1])
}

func newRequest(rev int) (*http.Request, error) {
	url := "http://integral.esac.esa.int/isocweb/schedule.html"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("selectMode", "rev")
	q.Add("action", "schedule")
	q.Add("startRevno", strconv.Itoa(rev))
	q.Add("endRevno", strconv.Itoa(rev))
	req.URL.RawQuery = q.Encode()
	return req, nil
}

func GetTable(data io.ReadCloser) ([][]string, error) {
	doc, err := goquery.NewDocumentFromReader(data)
	var out [][]string
	if err != nil {
		return nil, err
	}
	table := doc.Find("table").Last()
	table.Find("tr").NextAll().Each(func(i int, s *goquery.Selection) {
		out = append(out, make([]string, 12))
		s.Find("td").Each(func(j int, t *goquery.Selection) {
			if j == 7 {
				if k, ok := t.Find("a").Attr("href"); ok {
					out[i][j] = urlize(k)
					return
				}
			}
			out[i][j] = strings.TrimSpace(t.Text())
		})
	})
	return out, nil
}

func GetIds(data io.ReadCloser) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(data)
	var out []string
	if err != nil {
		return nil, err
	}
	table := doc.Find("table").Last()
	table.Find("tr").NextAll().Each(func(i int, s *goquery.Selection) {
		out = append(out, s.Find("td").First().Text())
	})
	return out, nil
}

