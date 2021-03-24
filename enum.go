//go:generate enumer -type=Pill -json

package main

import (
	"encoding/json"
	"fmt"
)

type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)

type Help struct {
	Drug Pill
}

func main() {
	fmt.Printf("value: %v\n", Paracetamol)

	help := Help{Drug: Aspirin}
	fmt.Printf("struct: %+v\n", help)

	js, _ := json.Marshal(help)
	fmt.Println("json: " + string(js))

	um := Help{}
	_ = json.Unmarshal(js, &um)
	fmt.Printf("unmarshalled: %+v\n", um)

	err := json.Unmarshal([]byte(`{"Drug":"invalid"}`), &um)
	fmt.Println(err)

}
