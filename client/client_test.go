package client_test

import (
	"etl-app/client"
	"etl-app/client/mocks"
	"etl-app/entity"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type funcSortDirection func() error
type funcSortByDiscovered func()
type funcCsv func(csvFile *os.File) error

type IncidentClientMock struct {
	SortDirectionField    funcSortDirection
	SortByDiscoveredField funcSortByDiscovered
	CsvField              funcCsv
}

func (in *IncidentClientMock) SortDirection() error {
	return in.SortDirectionField()
}

func (in *IncidentClientMock) SortByDiscovered() {
	in.SortByDiscoveredField()
}

func (in *IncidentClientMock) Csv(csvFile *os.File) error {
	return in.CsvField(csvFile)
}

func TestIncidentClientCsv(t *testing.T) {
	ascending, desceding := "ascending", "desceding"
	tmpdir := t.TempDir()
	csvFile, _ := ioutil.TempFile(tmpdir, "incidents.csv")
	cases := []struct {
		name          string
		incidents     entity.Incidents
		csv           *os.File
		sortDirection *string
		assert        func(err error)
	}{
		{
			name:          "success csv ascending",
			sortDirection: &ascending,
			incidents:     mocks.OkIncident,
			csv:           csvFile,
			assert: func(err error) {
				assert.Nil(t, err)
			},
		},
		{
			name:          "success csv descending",
			sortDirection: &desceding,
			incidents:     mocks.OkIncident,
			csv:           csvFile,
			assert: func(err error) {
				assert.Nil(t, err)
			},
		},
		{
			name:          "Error making csv",
			sortDirection: &desceding,
			incidents:     mocks.OkIncident,
			csv:           nil,
			assert: func(err error) {
				assert.Error(t, err)
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			incidentsClient := client.NewIncidentsClient(c.incidents, c.sortDirection)
			incidentsClient.SortByDiscovered()
			err := incidentsClient.Csv(c.csv)
			c.assert(err)

		})

	}
}

func TestIncidentSortByDirection(t *testing.T) {
	ascending, desceding := "ascending", "desceding"
	cases := []struct {
		name          string
		incidents     entity.Incidents
		sortDirection *string
		assert        func(err error)
	}{
		{
			name:          "success csv ascending",
			sortDirection: &ascending,
			incidents:     mocks.OkIncident,
			assert: func(err error) {
				assert.Nil(t, err)
			},
		},
		{
			name:          "success csv descending",
			sortDirection: &desceding,
			incidents:     mocks.OkIncident,
			assert: func(err error) {
				assert.Nil(t, err)
			},
		},
		{
			name:      "Error sorting",
			incidents: mocks.OkIncident,
			assert: func(err error) {
				assert.Error(t, err)
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			incidentsClient := client.NewIncidentsClient(c.incidents, c.sortDirection)
			err := incidentsClient.SortDirection()
			c.assert(err)

		})

	}
}
