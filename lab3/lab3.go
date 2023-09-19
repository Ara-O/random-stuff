package main

import (
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

type Data struct {
	BMRHasBeenCalculated bool
	BMR                  float64
}

type UserInput struct {
	Weight float64 `json:"weight"`
	Height float64 `json:"height"`
}

func main() {

	data := Data{
		BMRHasBeenCalculated: false,
		BMR:                  0,
	}

	str, err := os.ReadFile("index.html")

	if err != nil {
		log.Fatal(err)
	}

	htmlTemplate := string(str)

	t, err := template.New("").Parse(htmlTemplate)

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := t.Execute(w, data)

		if err != nil {
			log.Fatal("error", err)
		}
	})

	http.HandleFunc("/findBMR", func(w http.ResponseWriter, r *http.Request) {
		var input UserInput

		if r.Method != "POST" {
			return
		}

		height, err := strconv.Atoi(r.FormValue("height"))

		if err != nil {
			log.Fatal(err)
		}

		weight, err := strconv.Atoi(r.FormValue("weight"))

		if err != nil {
			log.Fatal(err)
		}

		input.Height = float64(height)
		input.Weight = float64(weight)

		data.BMRHasBeenCalculated = true
		data.BMR = (input.Weight * 703) / math.Pow(input.Height, 2)
		err = t.Execute(w, data)

		if err != nil {
			log.Fatal("error", err)
		}
	})

	http.ListenAndServe(":8080", nil)

	// var userWeight, userHeight, userBMI float64
	// fmt.Print("Enter your weight in pounds: ")
	// fmt.Scanln(&userWeight)
	// fmt.Print("Enter your height in inches: ")
	// fmt.Scanln(&userHeight)

	// userBMI = (userWeight * 703) / (math.Pow(userHeight, 2))

	// fmt.Printf("Your BMI is %.2f\n", userBMI)

	// switch true {
	// case userBMI < 18.5:
	// 	fmt.Println("You are underweight")
	// case userBMI <= 25:
	// 	fmt.Println("Your weight is optimal :D")
	// case userBMI > 25:
	// 	fmt.Println("You are overweight!")
	// }
}
