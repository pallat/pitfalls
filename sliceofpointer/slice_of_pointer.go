package sliceofpointer

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
)

type Data struct {
	Year               string `json:"year"`
	IndustryCodeAnzsic string `json:"industry_code_ANZSIC"`
	IndustryNameAnzsic string `json:"industry_name_ANZSIC"`
	RmeSizeGrp         string `json:"rme_size_grp"`
	Variable           string `json:"variable"`
	Value              string `json:"value"`
	Unit               string `json:"unit"`
}

func DoWithPointer() {
	f, err := os.Open("./annual-enterprise-survey-2023-financial-year-provisional-size-bands.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)

	var data []*Data
	var m = map[string]*Data{}
	for {
		line, err := r.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return
			}
			break
		}
		data = append(data, &Data{
			Year:               line[0],
			IndustryCodeAnzsic: line[1],
			IndustryNameAnzsic: line[2],
			RmeSizeGrp:         line[3],
			Variable:           line[4],
			Value:              line[5],
			Unit:               line[6],
		})
	}

	for i := range len(data) {
		key := data[i].IndustryCodeAnzsic + data[i].IndustryNameAnzsic
		m[key] = data[i]
	}

	x := make([]*Data, len(data))
	for _, v := range m {
		x = append(x, v)
	}
}

func DoWithNoPointer() {
	f, err := os.Open("./annual-enterprise-survey-2023-financial-year-provisional-size-bands.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)

	var data []Data
	var m = map[string]Data{}
	for {
		line, err := r.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return
			}
			break
		}
		data = append(data, Data{
			Year:               line[0],
			IndustryCodeAnzsic: line[1],
			IndustryNameAnzsic: line[2],
			RmeSizeGrp:         line[3],
			Variable:           line[4],
			Value:              line[5],
			Unit:               line[6],
		})

	}

	for i := range len(data) {
		key := data[i].IndustryCodeAnzsic + data[i].IndustryNameAnzsic
		m[key] = data[i]
	}

	x := make([]Data, len(data))
	for _, v := range m {
		x = append(x, v)
	}
}
