package main

import (
    "errors"
	"bufio"
    "encoding/csv"
    "fmt"
    "io"
    "os"
)

func DotProduct(a, b []float64) (float64, error) {
	product := float64(0)

	for i := 0; i < len(a); i++ {
		product += a[i] * b[i]
	}

	return product, nil
}

// main function to boot up everything
func main() {
    file, _ := os.Open("data/music_like.csv")
    reader := csv.NewReader(bufio.NewReader(file))

    var music []Music

    for {
        line, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
        }

        music = append(music, Music{
            ID:         line[0],
            Title:      line[1],
            Artist:      line[2],
        })
    }

    fmt.Println(string(result))
}