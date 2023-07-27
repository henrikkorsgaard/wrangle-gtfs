package repository

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/joho/godotenv"
	"henrikkorsgaard.dk/gtfs-service/testutils"
	"henrikkorsgaard.dk/gtfs-service/ingest"
)

var (
	testDataString string = "../testutils/data/GTFSDK.zip"
)

func init(){
	fmt.Println("Running repository basic ingest and fetch tests")
	godotenv.Load("../config_dev.env")
	err := testutils.ResetDatabase("./sql/gtfs.sql")
	if err != nil {
		panic(err)
	}
}

func TestIngestFetchAgency(t *testing.T){
	zbytes, err := testutils.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}

	gtfsFiles, err := ingest.ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Error TestIngestStops!")
	}

	data := gtfsFiles[0]
	agency, err := ingest.UnmarshallAgency(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestIngestAgency: " + err.Error())
	}

	// Singleton, so we will get the same each time anyways!
	repo, err := NewRepository()

	if err != nil {
		t.Error("Error TestIngestAgency: " + err.Error())
	}

	err = repo.IngestAgency(agency)
	assert.NoError(t, err)

	dbAgency, err :=  repo.FetchAgency();
	assert.Equal(t,len(data.Records), len(dbAgency))
}

func TestIngestFetchStops(t *testing.T){
	
	zbytes, err := testutils.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}

	gtfsFiles, err := ingest.ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Error TestIngestStops!")
	}

	data := gtfsFiles[7]
	stops, err := ingest.UnmarshallStops(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestIngestStops: " + err.Error())
	}

	// Singleton, so we will get the same each time anyways!
	repo, err := NewRepository()

	if err != nil {
		t.Error("Error TestIngestStops: " + err.Error())
	}

	err = repo.IngestStops(stops)
	assert.NoError(t, err)

	dbStops, err :=  repo.FetchStops();
	assert.Equal(t,len(data.Records), len(dbStops))

}

func TestIngestRoutes(t *testing.T){

	zbytes, err := testutils.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}

	gtfsFiles, err := ingest.ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Error TestIngestRoutes!")
	}

	data := gtfsFiles[5]
	routes, err := ingest.UnmarshallRoutes(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestIngestRoutes: " + err.Error())
	}

	repo, err := NewRepository()
	if err != nil {
		t.Error("Error TestIngestRoutes: " + err.Error())
	}

	err = repo.IngestRoutes(routes)
	assert.NoError(t, err)

	dbRoutes, err :=  repo.FetchRoutes();
	assert.Equal(t,len(data.Records), len(dbRoutes))
}

func TestIngestTrips(t *testing.T){
	
	zbytes, err := testutils.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}

	gtfsFiles, err := ingest.ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Error TestIngestTrips!")
	}

	data := gtfsFiles[10]
	trips, err := ingest.UnmarshallTrips(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestIngestTrips: " + err.Error())
	}

	repo, err := NewRepository()

	if err != nil {
		t.Error("Error TestIngestTrips: " + err.Error())
	}

	err = repo.IngestTrips(trips)
	assert.NoError(t, err)

	dbTrips, err :=  repo.FetchTrips();
	assert.Equal(t,len(data.Records), len(dbTrips))
}

func TestIngestShapes(t *testing.T){
	zbytes, err := testutils.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}

	gtfsFiles, err := ingest.ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Error TestIngestShapes!")
	}

	data := gtfsFiles[6]
	shapes, err := ingest.UnmarshallShapes(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestIngestShapes: " + err.Error())
	}

	repo, err := NewRepository()

	if err != nil {
		t.Error("Error TestIngestShapes: " + err.Error())
	}

	err = repo.IngestShapes(shapes)
	assert.NoError(t, err)

	dbShapes, err :=  repo.FetchShapes();
	assert.Equal(t,4, len(dbShapes))
}

func TestIngestStopTimes(t *testing.T){
	
	zbytes, err := testutils.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}

	gtfsFiles, err := ingest.ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Error TestStopTimes!")
	}

	data := gtfsFiles[8]
	stoptimes, err := ingest.UnmarshallStopTimes(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestStopTimes: " + err.Error())
	}

	repo, err := NewRepository()
	
	if err != nil {
		t.Error("Error TestStopTimes: " + err.Error())
	}

	err = repo.IngestStopTimes(stoptimes)
	assert.NoError(t, err)

	dbStopTimes, err :=  repo.FetchStopTimes();
	assert.NoError(t, err)
	assert.Equal(t,len(data.Records), len(dbStopTimes))
}
