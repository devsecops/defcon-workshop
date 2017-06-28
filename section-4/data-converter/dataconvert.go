// gets the filepath of the results file as an argument to the program
// checks if the results file exists and the size of it, if it exists
// converts the results file into BQ ingest-able format
// continues to upload the results to a BQ dataset

package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"cloud.google.com/go/bigquery"

	"github.com/astaxie/flatmap"
	"github.com/subosito/gotenv"
)

var ks []string

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

//CheckIfError function to check for errors
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

//Info function to print pretty output
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

//flatteneachRes function to flatten ecah result JSON array
func flatteneachRes(result string, wg *sync.WaitGroup) {
	var mp map[string]interface{}

	//reading the result string into a map of string interface
	err := json.Unmarshal([]byte(result), &mp)
	CheckIfError(err)

	//flattening that map
	fm, err := flatmap.Flatten(mp)
	CheckIfError(err)

	//Ranging through the flattened map and creating a string of File|Secret and adding to the ks string array
	for k := range fm {
		ks = append(ks, k+"|"+fm[k])
	}

	wg.Done()
}

func main() {

	//Loading the environment variables
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

	//Opening the results file to read each line
	resultsFile, err := os.Open(filepath)
	CheckIfError(err)
	defer resultsFile.Close()

	//Instantiating the bufio scanner to read the results file
	resScanner := bufio.NewScanner(resultsFile)

	//For each result object being read, pass it to the flatteneachRes function
	var wg sync.WaitGroup
	for resScanner.Scan() {
		wg.Add(1)
		res := resScanner.Text()
		go flatteneachRes(res, &wg)
	}
	wg.Wait()

	//Creating a temp csv file to store the final results
	outputFile, err := os.Create("/tmp/final.csv")
	CheckIfError(err)
	defer outputFile.Close()

	//Iterating through the string array and writing each one of them in the above CSV file
	for _, el := range ks {
		outputFile.WriteString(el + "\n")
	}

	Info("Now, uploading the results to BigQuery...")
	f, err := os.Open("/tmp/final.csv") //Need to open the file before uploading
	CheckIfError(err)
	defer f.Close()

	//defining the schema
	schema := bigquery.Schema{
		&bigquery.FieldSchema{Name: "File", Required: false, Type: bigquery.StringFieldType},
		&bigquery.FieldSchema{Name: "Secret", Repeated: false, Type: bigquery.StringFieldType},
	}

	//reading the processed file into BQ reader
	rs := bigquery.NewReaderSource(f)
	rs.AllowJaggedRows = true
	rs.AllowQuotedNewlines = true
	rs.FieldDelimiter = "|"
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
