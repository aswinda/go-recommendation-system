package main

import (
	"bufio"
    "encoding/csv"
    "fmt"
    "io"
	"os"
	"strconv"
	"log"
)

func DotProduct(a, b []float64) (float64, error) {
	product := float64(0)

	for i := 0; i < len(a); i++ {
		product += a[i] * b[i]
	}

	return product, nil
}

type Like struct {
    UserID      float64  `json:"user_id"`
    MusicID   	float64  `json:"music_id"`
}

// main function to boot up everything
func main() {
    file, _ := os.Open("data/music_like.csv")
    reader := csv.NewReader(bufio.NewReader(file))

    var likes []Like

    for {
        line, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
        }

		userId, _ := strconv.ParseFloat(line[0], 64)
		musicId, _ := strconv.ParseFloat(line[1], 64)
        likes = append(likes, Like{
            UserID:		userId,
            MusicID:	musicId,
        })
    }

    fmt.Printf("%v", likes)
}