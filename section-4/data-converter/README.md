# Convert Repo Supervisor data into BigQuery ingest-able format using a Data Converter

## Running Locally
* Create a BigQuery dataset `reposupervisords` and an empty table `reposupervisor` in Google BigQuery from the GCP UI to store the processed nmap results with the following schema (all nullable):
```
File:string
Secret:string
```
* Run Repo Supervisor by typing `docker run -it abhartiya/tools_gitallsecrets:v3 -token <> -orgname <> -toolName repo-supervisor`. Copy the `/data/results.txt` from the container to `results.json`.
* Complete the `.env` file in the `data-converter` folder with the appropriate `PROJECT_ID`, `DATASET_NAME` and `TABLE_NAME`
* In that folder, type `go run dataconvert.go results.json` - Run the data converter locally