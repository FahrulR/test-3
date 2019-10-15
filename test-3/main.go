package main

import (
	"encoding/json"
	f "fmt"
	"net/http"
	"strings"
)

type CountStruct struct {
	Consonants int
	Vowels     int
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome! try POST /countchar"))
	})
	http.HandleFunc("/countchar", CountCharControllers)
	f.Println("server running in port 3333")
	http.ListenAndServe(":3333", nil)
}

func CountCharControllers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		input := r.FormValue("input")
		consonants, vowels := CountCharModel(input)
		data := []CountStruct{
			CountStruct{consonants, vowels},
		}
		result, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "Bad request", http.StatusBadRequest)
}

func CountCharModel(input string) (int, int) {
	var countC, countV int
	check := []bool{false, false, false, false, false}
	input = strings.ToLower(input)
	input = strings.Replace(input, " ", "", -1)
	for _, str := range input {
		switch str {
		case 'a':
			if check[0] == false {
				countV++
				check[0] = true
			}
		case 'e':
			if check[1] == false {
				countV++
				check[1] = true
			}
		case 'i':
			if check[2] == false {
				countV++
				check[2] = true
			}
		case 'o':
			if check[3] == false {
				countV++
				check[3] = true
			}
		case 'u':
			if check[4] == false {
				countV++
				check[4] = true
			}
		default:
			countC++
		}

	}
	return countC, countV
}
