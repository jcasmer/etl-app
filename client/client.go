package client

import (
	"encoding/csv"
	"errors"
	"etl-app/entity"
	"log"
	"os"
	"sort"
	"strconv"
	"sync"
)

type IncidentsInterface interface {
	SortDirection() error
	SortByDiscovered()
	Csv(csvFile *os.File) error
}

type IncidentsClient struct {
	mu            sync.Mutex
	incident      entity.Incidents
	sortDirection *string
}

func NewIncidentsClient(incident entity.Incidents, sortDirection *string) IncidentsInterface {
	return &IncidentsClient{
		incident:      incident,
		sortDirection: sortDirection,
	}
}

func (in *IncidentsClient) SortDirection() error {
	if in.sortDirection == nil || *in.sortDirection == "" {
		log.Println("sortdirection not found")
		return errors.New("sortdirection not found")
	}

	log.Println("making sortdirection")
	var sortFilter func(i, j int) bool
	if *in.sortDirection == "ascending" {
		sortFilter = func(i, j int) bool {
			return in.incident.Incident[i].ID < in.incident.Incident[j].ID
		}
	} else {
		sortFilter = func(i, j int) bool {
			return in.incident.Incident[i].ID > in.incident.Incident[j].ID
		}
	}
	sort.Slice(in.incident.Incident, sortFilter)
	return nil
}

func (in *IncidentsClient) SortByDiscovered() {
	log.Println("making sorting by discovered")
	sort.Slice(in.incident.Incident, func(i, j int) bool {
		return in.incident.Incident[i].Discovered.ParseTime().Before(in.incident.Incident[j].Discovered.ParseTime())
	})

}

func (in *IncidentsClient) Csv(csvFile *os.File) error {
	log.Println("making csv")
	in.mu.Lock()
	defer in.mu.Unlock()

	w := csv.NewWriter(csvFile)

	header := []string{"id", "name", "discovered", "description", "status"}
	if err := w.Write(header); err != nil {
		log.Println("header", err)
		return err
	}

	for _, r := range in.incident.Incident {
		var csvRow []string

		csvRow = append(csvRow, strconv.Itoa(r.ID), r.Name, r.Discovered.Format(), r.Description, r.Status)

		if err := w.Write(csvRow); err != nil {
			log.Println("error writing record to csv::", err)
			return err
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		log.Println("error writing csv:", err)
		return err
	}

	return nil
}
