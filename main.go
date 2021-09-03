package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func filter(rows []row) (out []row) {
	for _, row := range rows {
		if row.pi != "Sergei Grebenev" {
			if row.target == "Gal. Bulge region" || row.target == "Galactic Center" {
				out = append(out, row)
			}
		}
	}
	return out
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("type the rev number")
	}
	if rev, err := strconv.Atoi(os.Args[1]); err != nil {
		log.Fatal(err)
	} else {
		c := http.Client{}
		req, err := newRequest(rev)
		if err != nil {
			log.Fatal(err)
		}
		resp, err := c.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		table, err := GetTable(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		data := rowsFromTable(table)
		for _, e := range filter(data) {
			if req, err := http.Get(e.pattern); err != nil {
				log.Fatal(err)
			} else {
				if ids, err := GetIds(req.Body); err != nil {
					log.Fatal(err)
				} else {
					for _, id := range ids {
						fmt.Println(id)
					}
				}
			}
		}
	}
}
