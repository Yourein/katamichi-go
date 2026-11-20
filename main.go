package main

import (
    "fmt"
    "log"
    "strings"
    "net/http"

    "github.com/PuerkitoBio/goquery"
)

type Reservation struct {
    Car       string
    Departure string
    Arrival   string
}

func normalize_string(s string) string {
    whitespace_stripped_s := strings.Join(strings.Fields(s), " ")
    return strings.ReplaceAll(whitespace_stripped_s, " （下記参照）", "")
}

func fetch_data() []Reservation {
    available := []Reservation{}

    res, err := http.Get("https://toyota-rentacar.com/?fp=t%2FWGf6fpT40ToDKVIpylTEeVTHfyJepR%2Bq8BsDQTmEw89Cf%2B1lcCCP6bkMbERjSICfGC3SxajVUjzGygz1UEtA%3D%3D&prvtof=ZqqiGB99cN%2FowlGpVw6sSkXMAzP10XMIncNfX%2FFwjGs%3D&poru=2NYiaB8r4Z6C90%2F0467BTq5ljD7HhVJPndWIiJrgtRdftvfd9m52OgaLrVtvEyRJ")

    if err != nil {
        log.Fatal(err)
    }

    defer res.Body.Close()
    if res.StatusCode != 200 {
        log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
    }

    doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
        log.Fatal(err)
    }

    doc.Find("#service-items-shop-type-start .service-item").Each(func(i int, s *goquery.Selection){
            car := s.Find(".service-item__info__car-type p").Eq(1).Text()
            raw_depart_shop := s.Find(".service-item__shop-start p").Eq(1).Text()
            raw_return_shop := s.Find(".service-item__shop-return p").Eq(1).Text()
            depart_shop := normalize_string(raw_depart_shop) 
            return_shop := normalize_string(raw_return_shop) 
            reservation := Reservation {
                Car: car,
                Departure: depart_shop,
                Arrival: return_shop,
            }
            available = append(available, reservation)
        })

    return available
}

func main() {
    reservations := fetch_data()
    fmt.Printf("%v\n", reservations)
}
