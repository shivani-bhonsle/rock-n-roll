package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"rock-n-roll/models"
	"sync"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func GetPhotos(c *gin.Context) {
	photos := fetchPhotos()
	c.IndentedJSON(http.StatusOK, photos)
}

func fetchPhotos() []models.Photos {
	url := `https://jsonplaceholder.typicode.com/photos`

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while making an api call")
	}

	var photos []models.Photos

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error while reading body")
	}

	err = json.Unmarshal(body, &photos)
	if err != nil {
		fmt.Println("Error while Unmarshalling body")
	}
	// fmt.Println(string(body))

	// fmt.Println(photos)
	return photos

}

func FetchPhotosRoutine(wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	photos := fetchPhotos()
	// fmt.Println(photos)
	for _, photo := range photos {
		// fmt.Println(photo.Id)
		ch <- photo.Id
	}
	close(ch)

}

func PrintPhotoId(wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()
	// _, closed := <-ch
	// if !closed {
	// 	fmt.Println("CLOSED CHANNEL")
	// 	return
	// }
	for val := range ch {
		fmt.Println("From PrintPhotoId func", val)
	}
	fmt.Println("CLOSED CHANNEL")

}

// func fetchPostAPI(){
// 	url:= `abc`
// 	payload := Data{
// 		ID: id
// 	}
// 	jsonData, err:=json.Marshal(payload)
// 	req, err :=http.NewRequest("POST", url, bytes.Buffer(jsonData))
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("api-key","ajsdhjhdjabbxmz")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	defer resp.Body.Close()
// 	io.ReadAll(resp.Body)

// }
