package cmd

import (
	"encoding/json"
	"etl-app/client"
	"etl-app/entity"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

func Commands() error {
	var incident entity.Incidents

	sortDirection := flag.String("sortdirection", "", "sort direction data")

	body, _ := ioutil.ReadFile("./incidents-json/incidents.json")
	flag.Parse()

	err := json.Unmarshal(body, &incident.Incident)
	if err != nil {
		log.Println("error unmarshal:", err)
		return err
	}
	records := client.NewIncidentsClient(incident, sortDirection)
	records.SortByDiscovered()
	records.SortDirection()

	csvFile, err := os.Create("./csv/incidents.csv")

	if err != nil {
		log.Println("failed creating file:", err)
		return err
	}
	err = records.Csv(csvFile)
	if err != nil {
		log.Println("error writing csv:", err)
		return err
	}
	csvFile.Close()
	return nil

}
