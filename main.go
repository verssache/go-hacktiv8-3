package main

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
)

type Data struct {
	Status struct {
		Water int `json:"water"`
		Wind  int `json:"wind"`
	} `json:"status"`
	UpdatedAt string `json:"updated_at"`
}

type ResultData struct {
	Data struct {
		Water string `json:"water"`
		Wind  string `json:"wind"`
	} `json:"data"`
	Status    string `json:"status"`
	UpdatedAt string `json:"updated_at"`
}

func main() {
	tz, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatal(err.Error())
	}

	s := gocron.NewScheduler(tz)
	s.Every(15).Seconds().Do(func() {
		data, err := os.Open("data.json")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer data.Close()

		byteValue, _ := io.ReadAll(data)
		var result Data
		err = json.Unmarshal(byteValue, &result)
		if err != nil {
			log.Fatal(err.Error())
		}

		rand.Seed(time.Now().UnixNano())
		result.Status.Water = rand.Intn(100)
		result.Status.Wind = rand.Intn(100)
		result.UpdatedAt = time.Now().In(tz).Format("2006-01-02 15:04:05")

		file, err := os.Create("data.json")
		if err != nil {
			log.Fatal(err.Error())
		}

		jsonData, err := json.Marshal(result)
		if err != nil {
			log.Fatal(err.Error())
		}

		_, err = file.Write(jsonData)
		if err != nil {
			log.Fatal(err.Error())
		}
	})

	s.StartAsync()

	router := gin.Default()
	router.GET("/data", func(c *gin.Context) {
		data, err := os.Open("data.json")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer data.Close()

		byteValue, _ := io.ReadAll(data)
		var result Data
		err = json.Unmarshal(byteValue, &result)
		if err != nil {
			log.Fatal(err.Error())
		}

		var resultData ResultData

		if result.Status.Water < 5 {
			resultData.Status = "Aman"
		} else if result.Status.Water >= 6 && result.Status.Water <= 8 {
			resultData.Status = "Siaga"
		} else if result.Status.Water > 8 {
			resultData.Status = "Bahaya"
		} else if result.Status.Wind < 6 {
			resultData.Status = "Aman"
		} else if result.Status.Wind >= 7 && result.Status.Wind <= 15 {
			resultData.Status = "Siaga"
		} else if result.Status.Wind > 15 {
			resultData.Status = "Bahaya"
		}

		resultData.Data.Water = strconv.Itoa(result.Status.Water) + " m"
		resultData.Data.Wind = strconv.Itoa(result.Status.Wind) + " m/s"
		resultData.UpdatedAt = result.UpdatedAt

		c.JSON(200, resultData)
	})

	router.Run(":8081")

}
