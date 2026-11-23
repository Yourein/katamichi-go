package main

import (
    "errors"
    "fmt"
    "log"
    "strings"
    "net/http"
    "time"

    "github.com/PuerkitoBio/goquery"
)

type Reservation struct {
    Car           string
    Departure     string
    Arrival       string
    DepartureDate string
}

func normalize_string(s string) string {
    whitespace_stripped_s := strings.Join(strings.Fields(s), " ")
    return strings.ReplaceAll(whitespace_stripped_s, " （下記参照）", "")
}

func fetch_data() ([]Reservation, error) {
    available := []Reservation{}

    res, err := http.Get("https://toyota-rentacar.com/?fp=t%2FWGf6fpT40ToDKVIpylTEeVTHfyJepR%2Bq8BsDQTmEw89Cf%2B1lcCCP6bkMbERjSICfGC3SxajVUjzGygz1UEtA%3D%3D&prvtof=ZqqiGB99cN%2FowlGpVw6sSkXMAzP10XMIncNfX%2FFwjGs%3D&poru=2NYiaB8r4Z6C90%2F0467BTq5ljD7HhVJPndWIiJrgtRdftvfd9m52OgaLrVtvEyRJ")

    if err != nil {
        log.Fatal(err)
    }

    defer res.Body.Close()
    if res.StatusCode != 200 {
        errStr := fmt.Sprintf("status code error: %d %s", res.StatusCode, res.Status)
        return nil, errors.New(errStr)
    }

    doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
        log.Fatal(err)
    }

    doc.Find("#service-items-shop-type-start .service-item").Each(func(i int, s *goquery.Selection){
           if s.Find(".service-item__body.show-entry-end").Length() > 0 {
                return
            }

            car := s.Find(".service-item__info__car-type p").Eq(1).Text()
            raw_depart_shop := s.Find(".service-item__shop-start p").Eq(1).Text()
            raw_return_shop := s.Find(".service-item__shop-return p").Eq(1).Text()
            depart_shop := normalize_string(raw_depart_shop) 
            return_shop := normalize_string(raw_return_shop)
            departure_date := s.Find(".service-item__info__date").Text()
            reservation := Reservation {
                Car: car,
                Departure: depart_shop,
                Arrival: return_shop,
                DepartureDate: departure_date,
            }
            available = append(available, reservation)
        })

    return available, nil
}

// (before, after) -> (addition, deletion)
func find_diff(before []Reservation, after []Reservation) ([]Reservation, []Reservation) {
    matched := make([]bool, len(before), len(before))
    addition := []Reservation{}
    deletion := []Reservation{}

    for i, _ := range matched {
        matched[i] = false
    }

    for _, ax := range after {
        found := false
        for i, bx := range before {
            // 既にafterの要素とマッチング出来ていたら飛ばす
            if matched[i] {
                continue
            }

            if bx == ax {
                matched[i] = true
                found = true
                break
            }
        }

        // 前方からbeforeを探してマッチング対象がなければ新規掲載
        if !found {
            addition = append(addition, ax)
        }
   }

    // afterの要素とマッチングしなかったbeforeの要素は掲載終了
    for i, mx := range matched {
        if !mx {
            deletion = append(deletion, before[i])
        }
    }

    return addition, deletion
}

func main() {
    old_reservations := []Reservation{}

    for {
        reservations, err := fetch_data()

        if err != nil {
            whErr := sendError(err)
            
            if whErr != nil {
                log.Print(err)
            }

            time.Sleep(5 * time.Minute)
            continue
        }

        additions, deletions := find_diff(old_reservations, reservations)

        for _, addition := range additions {
            err := sendAddition(addition)

            if err != nil {
                log.Print(err)
            }

            time.Sleep(100 * time.Millisecond)
        }

        for _, deletion := range deletions {
            err := sendDeletion(deletion)

            if err != nil {
                log.Print(err)
            }

            time.Sleep(100 * time.Millisecond)
        }

        old_reservations = reservations
        time.Sleep(5 * time.Minute)
    }
}
