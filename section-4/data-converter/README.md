# Convert data into BigQuery ingest-able format using a Generic Data Converter

## Running Locally
* Replace the directory path `/path/where/gac/is/stored/` in the command below to the **directory** where GAC credentials file is stored locally on your workstation. Replace the `PROJECTID` as well. Replace the `gacfilename` with the name of the GAC credentials file.

`docker run -it -v /path/where/gac/is/stored/:/tmp/data/ abhartiya/utils_bqps:v1 -project <PROJECTID> -gac /tmp/data/<gacfilename> -wfdataset wfuzzds -wftable wfuzz_tomcat_test -rsdataset reposupervisords -rstable reposupervisor_test` 

the output should look simliar to this: 

```wfuzzds already exists
wfuzz_tomcat_test already exists. Deleting it now..
wfuzz_tomcat_test deleted. Creating it again now..
wfuzz_tomcat_test recreated
reposupervisords already exists
reposupervisor_test already exists. Deleting it now..
reposupervisor_test deleted. Creating it again now..
reposupervisor_test recreated```


* The above command will mount the local directory where you stored your GAC credentials file to `/tmp/data` inside the container. Once, it does that, it will run the `abhartiya/utils_bqps:v1` container with the arguments - `-project defcon-workshop -gac /tmp/data/<gacfilename> -wfdataset wfuzzds -wftable wfuzz_tomcat_test -rsdataset reposupervisords -rstable reposupervisor_test`. Once the container runs, the following will happen:

    * Create a BiqQuery dataset `wfuzzds` and an empty table `wfuzz_tomcat_test` in Google BigQuery to store the processed wfuzz results with the following schema (all nullable):

    ```
    ID:string
    Response:string
    Lines:string
    Word:string
    Chars:string
    Request:string
    Success:string
    ```

    * Create a BigQuery dataset `reposupervisords` and an empty table `reposupervisor_test` in Google BigQuery to store the processed repo-supervisor results with the following schema (all nullable):

    ```
    File:string
    Secret:string
    ```

* Run Repo Supervisor by typing `docker run -it abhartiya/tools_gitallsecrets:v3 -token <> -org <> -toolName repo-supervisor`. Copy the `/data/results.txt` from the container to `results.json` to do that first find the container ID by typing `docker ps -a`.  Then to copy the file type `docker cp contID:/data/results.txt ./data-converter/results.json`

* Run WFUZZ by typing `docker run -it abhartiya/tools_wfuzz -w /data/SecLists/Discovery/Web_Content/tomcat.txt --hc 404,429,400 -o csv <URL>/FUZZ /data/out.csv`. Copy the `/data/out.csv` from the container to `out.csv` in the data-converter folder

* Complete the `.env.sample` file in the `data-converter` folder with the appropriate values and rename it to `.env`

* `go get github.com/astaxie/flatmap`

* Now, in order to convert Repo Supervisor's output and upload it to BQ, type `go run dataconvert.go -toolName repo-supervisor -filePath results.json`

* Next, in order to convert WFUZZ's output and upload it to BQ, type `go run dataconvert.go -toolName wfuzz -filePath out.csv`
