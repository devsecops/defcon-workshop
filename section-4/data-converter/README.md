# Convert data into BigQuery ingest-able format using a Generic Data Converter

1. Replace the directory path `/path/where/gac/is/stored/` in the command below to the **directory** where GAC credentials file is stored locally on your workstation. Replace the `PROJECTID` as well. Replace the `gacfilename` with the name of the GAC credentials file.

`docker run -it -v /path/where/gac/is/stored/:/tmp/data/ abhartiya/utils_bqps:v1 -project <PROJECTID> -gac /tmp/data/<gacfilename> -wfdataset wfuzzds -wftable wfuzz_tomcat_test -rsdataset reposupervisords -rstable reposupervisor_test`

The above command will mount the local directory where you stored your GAC credentials file to `/tmp/data` inside the container. Once, it does that, it will run the `abhartiya/utils_bqps:v1` container with the arguments - `-project defcon-workshop -gac /tmp/data/<gacfilename> -wfdataset wfuzzds -wftable wfuzz_tomcat_test -rsdataset reposupervisords -rstable reposupervisor_test`. Once the container runs, the following will happen:

* A BiqQuery dataset `wfuzzds` and an empty table `wfuzz_tomcat_test` in Google BigQuery will be created to store the processed wfuzz results with the following schema (all nullable):

```
ID:string
Response:string
Lines:string
Word:string
Chars:string
Request:string
Success:string
```

* A BigQuery dataset `reposupervisords` and an empty table `reposupervisor_test` in Google BigQuery will be created to store the processed git-all-secrets results with the following schema (all nullable):

```
File:string
Secret:string
```

2. `cd` into the `data-converter` directory.

3. Run `git-all-secrets` by typing `docker run -it abhartiya/tools_gitallsecrets:v3 -token <git-personal-access-token> -org kubebot -toolName repo-supervisor`. Make sure you replace the `git-personal-access-token` with the appropriate value.

4. Copy the `/data/results.txt` from the container to `results.json`. To do that, first find the container ID by typing `docker ps -a`.  Then to copy the file, type `docker cp contID:/data/results.txt results.json`.

4. Run `WFUZZ` by typing `docker run -it abhartiya/tools_wfuzz -w /data/SecLists/Discovery/Web_Content/tomcat.txt --hc 404,429,400 -o csv http://104.198.4.57/FUZZ /data/out.csv`.

5. Copy the `/data/out.csv` from the container to `out.csv`. To do that, first find the container ID by typing `docker ps -a`.  Then to copy the file, type `docker cp contID:/data/out.csv out.csv`.

6. Complete the `.env.sample` file in the `data-converter` directory with the appropriate values and copy it to `.env`.

7. Download the `flatmap` library for the data converter by typing `go get github.com/astaxie/flatmap`.

8. Now, in order to convert git-all-secret's output and upload it to BQ, type `go run dataconvert.go -toolName repo-supervisor -filePath results.json`.

9. Next, in order to convert WFUZZ's output and upload it to BQ, type `go run dataconvert.go -toolName wfuzz -filePath out.csv`.

10. Ensure that the `reposupervisor_test` and `wfuzz_tomcat_test` BQ tables were updated.