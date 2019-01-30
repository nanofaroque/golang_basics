package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	Species string
	Description string
	Dimensions Dimensions

}

type Dimensions struct {
	Height int
	Width int
}

func main() {
	structureExample()
	jsonArrayExample()
	embeddedExample()
	unStructuredExample()
}
func structureExample()  {
	birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Printf("Species: %s, Description: %s", bird.Species, bird.Description)
	//Species: pigeon, Description: likes to perch on rocks
}

func jsonArrayExample()  {
	birdJson := `[{"species":"pigeon","decription":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`
	var birds []Bird
	json.Unmarshal([]byte(birdJson), &birds)
	fmt.Printf("Birds : %+v", birds)
	//Birds : [{Species:pigeon Description:} {Species:eagle Description:bird of prey}]
}

func embeddedExample()  {
	fmt.Println("===========================")
	birdJson := `{"species":"pigeon","description":"likes to perch on rocks", "dimensions":{"height":24,"width":10}}`
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Println(bird)
	// {pigeon likes to perch on rocks {24 10}}
}

func unStructuredExample()  {
	birdJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"},"animals":"none"}`
	var result map[string]interface{}
	json.Unmarshal([]byte(birdJson), &result)

	// The object stored in the "birds" key is also stored as
	// a map[string]interface{} type, and its type is asserted from
	// the interface{} type
	birds := result["birds"].(map[string]interface{})

	for key, value := range birds {
		// Each value is an interface{} type, that is type asserted as a string
		fmt.Println(key, value.(string))
	}
}
