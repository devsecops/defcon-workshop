// gets the filepath of the results file as an argument to the program
// checks if the results file exists and the size of it, if it exists
// converts the results file into BQ ingest-able format
// continues to upload the results to a BQ dataset

package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/subosito/gotenv"

	"cloud.google.com/go/bigquery"
)

type error interface {
	Error() string
}

var (
	projectID   string
	datasetName string
	tableName   string
)

func exists(path string) (bool, int64, error) {
	fi, err := os.Stat(path)
	if err == nil {
		return true, fi.Size(), nil
	}
	if os.IsNotExist(err) {
		return false, int64(0), nil
	}
	return false, int64(0), err
}

func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

func main() {

	err := gotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	projectID = os.Getenv("PROJECT_ID")
	datasetName = os.Getenv("DATASET_NAME")
	tableName = os.Getenv("TABLE_NAME")

	// instantiating the BQ client
	ctx := context.Background()
	bqclient, err := bigquery.NewClient(ctx, projectID)
	CheckIfError(err)

	//getting the filepath as a command line argument
	filepath := os.Args[1]
	Info("Filepath: " + filepath + "\n")

	value := false
	fsize := int64(0)

	//checking for file existence and file size
	for (value == false) || (fsize == int64(0)) {
		i, s, err := exists(filepath)
		CheckIfError(err)
		value = i
		fsize = s
	}

	fmt.Println("File exists:", value)
	fmt.Println("File size:", fsize)
	fmt.Println("")
	//now, we know that the file exists and that its size>0

	sa := strings.Split(filepath, "/")                              //getting the filename from the filepath
	t := time.Now()                                                 //getting the current time
	rawFilename := sa[len(sa)-1] + "_" + t.Format("20060102150405") //adding the timestamp to the filename

	//convert data into BQ ingest-able
	processedFilename := rawFilename + ".csv"
	cmd := exec.Command("python", "nmaptocsv.py", "-i", filepath, "-o", "/tmp/"+processedFilename, "-n", "-s") //converting nmap output to csv
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	CheckIfError(err)

	Info("Now, uploading the results to BigQuery...")

	//opening the processed file
	f, err := os.Open("/tmp/" + processedFilename)
	CheckIfError(err)
	defer f.Close()

	//defining the schema
	schema := bigquery.Schema{
		&bigquery.FieldSchema{Name: "IP", Required: false, Type: bigquery.StringFieldType},
		&bigquery.FieldSchema{Name: "FQDN", Repeated: false, Type: bigquery.StringFieldType},
		&bigquery.FieldSchema{Name: "Port", Required: false, Type: bigquery.StringFieldType},
		&bigquery.FieldSchema{Name: "Protocol", Required: false, Type: bigquery.StringFieldType},
		&bigquery.FieldSchema{Name: "Service", Required: false, Type: bigquery.StringFieldType},
		&bigquery.FieldSchema{Name: "Version", Required: false, Type: bigquery.StringFieldType},
	}

	//reading the processed file into BQ reader
	rs := bigquery.NewReaderSource(f)
	rs.AllowJaggedRows = true
	rs.AllowQuotedNewlines = true
	rs.FieldDelimiter = ";"
	rs.IgnoreUnknownValues = true
	rs.Schema = schema

	//instantiating the dataset
	ds := bqclient.Dataset(datasetName)
	loader := ds.Table(tableName).LoaderFrom(rs) //loading the results
	loader.CreateDisposition = bigquery.CreateNever

	//checking the job status
	job, err := loader.Run(ctx)
	CheckIfError(err)
	status, err := job.Wait(ctx)
	CheckIfError(err)
	err = status.Err()
	CheckIfError(err)
	fmt.Println("Done")

}
