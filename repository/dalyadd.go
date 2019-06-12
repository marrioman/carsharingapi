package repository

import (
	"delycarapi/entity"
	"time"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func PostData() {
	response, err := http.Get(entity.IssuesURL)
	if err != nil {
		log.Fatalln("The HTTP request failed with error %s\n", err)
	}

	if response.StatusCode > 499 {
		log.Println("StatusCode:", response.StatusCode)
		time.Sleep(time.Second)
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	defer response.Body.Close()

	data := []byte(body)

	d := &entity.Dely{}
	err = json.Unmarshal(data, d)
	if err != nil {
		log.Println("Incorrect unmarshal", err)
		return
	}

	// var delIDs []int ---
	for _, car := range d.Cars {
		err = db.Table("dely").Save(&car).Error
		if err != nil {
			log.Fatalln(err)
			return
		}
		//	delIDs = append(delIDs, car.ID)
	}

	err = SetGeo()
	if err != nil {
		log.Println("Geo insert error", err)
		return
	}
}
