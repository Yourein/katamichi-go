package main

import (
    "io"
    "fmt"
    "log"
    "net/http"

    // "github.com/PuerkitoBio/goquery"
)

func main() {
    res, err := http.Get("https://toyota-rentacar.com/?fp=t%2FWGf6fpT40ToDKVIpylTEeVTHfyJepR%2Bq8BsDQTmEw89Cf%2B1lcCCP6bkMbERjSICfGC3SxajVUjzGygz1UEtA%3D%3D&prvtof=ZqqiGB99cN%2FowlGpVw6sSkXMAzP10XMIncNfX%2FFwjGs%3D&poru=2NYiaB8r4Z6C90%2F0467BTq5ljD7HhVJPndWIiJrgtRdftvfd9m52OgaLrVtvEyRJ")

    if err != nil {
        log.Fatal(err)
    }

    defer res.Body.Close()
    if res.StatusCode != 200 {
        log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
    }

    fmt.Printf("%d\n", res.StatusCode)
    fmt.Printf("%s\n", res.Body)

    content, err := io.ReadAll(res.Body)
    if err != nil {
        log.Fatal(err)
    }

    responseString := string(content)
    fmt.Println(responseString)
}
