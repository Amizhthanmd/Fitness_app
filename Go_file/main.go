package main

import (
    "net/http"
	"fmt"
	"encoding/json"
)

func main() {
	fmt.Println("Running...")
	http.HandleFunc("/GetBmi", GetBmi)
	http.HandleFunc("/GetPro", GetPro)
	http.HandleFunc("/GetWat", GetWat)
	http.ListenAndServe(":8080", nil)
}

// BMI Calculation --------------------------------------------------------------------------------------------------------------------

func GetBmi(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	type value struct {
		Weight float32	`json:"weight"`
		Height float32  `json:"height"`
	}

	type Response struct{
		Bmivalue float32 `json:"bmivalue"`
	}

	var Value value

	decoder.Decode(&Value)

	r.ParseForm()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	// cross-origin resource sharing (CORS)
	if r.Method == "OPTIONS" {
		return
	}

    if err := json.NewEncoder(w).Encode(Value); err != nil {
        panic(err)                                                 
    }

	meterheight := Value.Height/100 

	var Bmivalue float32 = Value.Weight / (meterheight * meterheight)

	response := Response {Bmivalue: Bmivalue}

	if err := json.NewEncoder(w).Encode(response); 
	   err != nil {
       panic(err)                                                 
    }

	fmt.Println("Weight:", Value.Weight)
	fmt.Println("Height:", Value.Height)
	fmt.Println("BMI:", Bmivalue)

}

// Protein Intake Calculation --------------------------------------------------------------------------------------------------------------------

func GetPro(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	type value struct {
		Weight float32	`json:"weight"`
		Workoutdaily bool `json:"workout`
	}

	type Response struct{
		Proteinvalue float32 `json:"proteinvalue"`
		WorkProteinvalue float32 `json:"workproteinvalue"`

	}
	var Value value

	decoder.Decode(&Value)

	r.ParseForm()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	
	if r.Method == "OPTIONS" {
		return
	}

    if err := json.NewEncoder(w).Encode(Value); err != nil {
        panic(err)                                                 
    }

	var Protein float32
	var Gymprotein float32

	if Value.Workoutdaily == false{
		Protein  = float32(Value.Weight) * 0.8
		fmt.Printf("Your protein is: %.1f \n", Protein)

	}
	if Value.Workoutdaily == true {
		Gymprotein = float32(Value.Weight) * 1.6
		fmt.Printf("Your gymprotein is: %.1f \n", Gymprotein)

	}
	response := Response {Proteinvalue: Protein ,WorkProteinvalue: Gymprotein}

	if err := json.NewEncoder(w).Encode(response); err != nil {
        panic(err)                                                 
    }

	fmt.Println("Weight:", Value.Weight)
	fmt.Println("Workoutdaily:", Value.Workoutdaily)
}

// Water Intake Calculation --------------------------------------------------------------------------------------------------------------------

func GetWat(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	type value struct {
		Weight float32	`json:"weight"`
	}

	type Response struct{
		Watervalue float32 `json:"watervalue"`

	}
	var Value value

	decoder.Decode(&Value)

	r.ParseForm()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	
	if r.Method == "OPTIONS" {
		return
	}

    if err := json.NewEncoder(w).Encode(Value); err != nil {
        panic(err)                                                 
    }

	Waterintakes := Value.Weight * 30/1000

	response := Response {Watervalue: Waterintakes}

	if err := json.NewEncoder(w).Encode(response); err != nil {
        panic(err)                                                 
    }

	fmt.Println("Weight:", Value.Weight)
	fmt.Println("Waterintakes:", Waterintakes)



}

//-----------------------------------------------------------------------------------------------------------------------------------------------