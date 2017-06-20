# Section-1

`export PROJECT_ID=(GCP-Project-ID)`

## Building the first Docker image
`docker build -t us.gcr.io/$PROJECT_ID/test:v1 .`
`docker run --rm us.gcr.io/$PROJECT_ID/test:v1`

## Pushing the Docker image to Google Container Registry (GCR)
`gcloud docker -- push us.gcr.io/$PROJECT_ID/test:v1`
