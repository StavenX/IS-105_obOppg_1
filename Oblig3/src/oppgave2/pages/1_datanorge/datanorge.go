package _datanorge

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"time"
	"strings"
)

func main() {
	getJson()
	http.HandleFunc("/1", PrintToServer)
	http.ListenAndServe(":8080", nil)
}

// ---------------------------------------------------

type Datasets struct {
	Datasets []Dataset `json:"datasets"`
}

type Dataset struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Antall string `json:"antall"`
	Description []Description `json:"description"`
	Distribution []Distribution `json:"distribution"`
}

type Distribution struct {
	ID string `json:"id"`
	Title string `json:"title"`
}

type Description struct {
	Language string `json:"language"`
	Value string `json:"value"`
}

var datasets Datasets
var URL = "https://data.norge.no/api/dcat/data.json?page=1"

// ---------------------------------------------------

func getJson() {
	url := "https://data.norge.no/api/dcat/data.json?page=1"

	client := http.Client{
		Timeout: time.Second *2,
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Request error.")
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Response error.")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Body error.")
	}

	jsonerr := json.Unmarshal(body, &datasets)
	if jsonerr != nil {
		fmt.Println("JSON unmarshal error.")
	}

	fmt.Printf("Length: %d\n", len(datasets.Datasets))

	for i := 0; i < len(datasets.Datasets); i++ {
		PrintAll(i)

	}
}

func PrintAll(index int) {
	d:= datasets.Datasets[index]
	fmt.Printf("\nID: %s\n", d.ID)
	fmt.Printf("Tittel: %s\n", d.Title)
	fmt.Printf("Antall: %s\n", d.Antall)
	if len(d.Description) > 0 {
		s := d.Description[0].Value
		s = strings.Replace(s, "\u003Cp\u003E", "", -1)
		s = strings.Replace(s, "\u003C/p\u003E", "", -1)
		fmt.Printf("Beskrivelse: %s\n", s)
		fmt.Println()
	}
}


func PrintToServer(writer http.ResponseWriter, request *http.Request) {
	for i := 0; i < len(datasets.Datasets); i++ {
		d := datasets.Datasets[i]
		fmt.Fprintf(writer,"\nID: %s\n", d.ID)
		fmt.Fprintf(writer,"Tittel: %s\n", d.Title)
		fmt.Fprintf(writer,"Antall: %s\n", d.Antall)
		if len(d.Description) > 0 {
			fmt.Fprintf(writer, "Beskrivelse: %s\n", d.Description[0].Value)
		}
	}
}