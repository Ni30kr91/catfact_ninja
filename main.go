package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

type CatFact struct {
	Breed   string `json:"breed"`
	Country string `json:"country,omitempty"`
	Origin  string `json:"origin"`
	Coat    string `json:"coat"`
	Pattern string `json:"pattern"`
}

type WordCountRequest struct {
	String string `json:"str"`
}

type CatResponse struct {
	Data     []CatFact `json:"data"`
	LastPage int       `json:"last_page"`
}

type CatBreeds map[string][]CatFact

func main() {
	router := gin.Default()
	router.POST("/", CheckWordCount)
	router.GET("/cat-breeds", GetCatBreeds)
	router.Run(":8080")
}

func CheckWordCount(c *gin.Context) {
	var request WordCountRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Using regex to count words
	regex := regexp.MustCompile(`\w+`)
	words := regex.FindAllString(request.String, -1)

	if len(words) >= 8 {
		c.JSON(http.StatusOK, gin.H{"message": "200 OK"})
	} else {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Not Acceptable"})
	}
}

func GetCatBreeds(c *gin.Context) {
	data, err := GetCatBreedsByPage()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetCatBreedsByPage() (data map[string][]CatFact, errr error) {
	pageNo := 1
	printedOK := false
	data = make(map[string][]CatFact)
	for i := 1; i <= pageNo; i++ {
		resp, err := http.Get("https://catfact.ninja/breeds?page=" + fmt.Sprint(i))
		if err != nil {
			log.Println(err)
			errr = errors.New("Failed to fetch cat breeds")
			break
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			errr = errors.New("Failed to read response body")
			break
		}

		// Log the response as-is to a text file
		err = ioutil.WriteFile("cat_breeds_response.txt", body, 0644)
		if err != nil {
			log.Println(err)
			errr = errors.New("Failed to write response to file")
			break
		}

		catBreeds := CatResponse{}
		err = json.Unmarshal(body, &catBreeds)
		if err != nil {
			log.Println(err)
			errr = errors.New("Failed to parse cat breeds")
			break
		}

		for _, catBreed := range catBreeds.Data {
			country := catBreed.Country
			catBreed.Country = ""
			data[country] = append(data[country], catBreed)
		}

		// Console log the number of pages of data available
		if !printedOK {
			fmt.Printf("Number of pages of data available: %d\n", catBreeds.LastPage)
			printedOK = true
		}
		pageNo = catBreeds.LastPage

	}
	return
}
