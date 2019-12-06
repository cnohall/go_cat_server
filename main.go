package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
    "io/ioutil"
    "os"
    // "strconv"
)

type Cat struct {
	Name string `json:"name"`
	Image string `json:"image"`
	CutenessLevel int `json:"cutenessLevel"`
	AllergyInducingFur bool `json:"allergyInducingFur"`
	LivesLeft int `json:"livesLeft"`
}

type Cats struct {
    Cats []Cat `json:"cats"`
}

func allCats(w http.ResponseWriter, r *http.Request) {
	// cats := Cats{
	// 	Cat  {
	// 		Name: "Parker",
	// 		Image: "1.jpg",
	// 		CutenessLevel: 82,
	// 		AllergyInducingFur: false,
	// 		LivesLeft: 8,
	// 	  },
	// 	}

	fmt.Println("Endpoint Hit: All Cats Endopoint")

}

func homePage(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Homepage Endpoint Hit!")
	// Open our jsonFile
	jsonFile, err := os.Open("catdata.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	
	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var cats Cats

	json.Unmarshal(byteValue, &cats)

	json.NewEncoder(w).Encode(cats)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/allCats", allCats)
	log.Fatal(http.ListenAndServe(":8081",nil))
}

func main() {
	handleRequests()



}

// for i := 0; i < len(cats.Cats); i++ {
// 	fmt.Println("Cat Name: " + cats.Cats[i].Name)
// 	fmt.Println("Cat Image: " + cats.Cats[i].Image)
// 	fmt.Println("Cat CutenessLevel: " + strconv.Itoa(cats.Cats[i].CutenessLevel))
// 	fmt.Println("Cat AllergyInducingFur: " + strconv.FormatBool(cats.Cats[i].AllergyInducingFur))
// 	fmt.Println("Cat LivesLeft: " + strconv.Itoa(cats.Cats[i].LivesLeft))
// }