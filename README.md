# Image-resizer: sample app

This project is a demo app used for Kubernetes and Docker Swarm deployment examples.

There isn't any Swarm or Kubernetes manifest in this repository, you'll need to create them. This project require Postgres, nats.io and Minio. You'll need to create deployments for those dependencies too. In the `Helpful informations` section, you will find docker commands to start Postgres, nats and Minio.

## Project architecture:

![Architecture](docs/architecture.png?raw=true 'Architecture')

## Services configuration:

### Frontend:

To build the frontend you first need to edit the `.env.prod` file. Then you can build the project by running `yarn build`.

_Environment variables (at build time):_

- VUE_APP_JOB_API: The Jobs-API url
- VUE_APP_IMAGES_API: The Images-API url

Docker image: alexandrevilain/image-resizer-frontend

This image needs to be rebuild with working env variables

### Jobs-API:

The jobs-API is a REST API providing image upload endpoint to create a resize job.
When you send an image on `/upload` its uploaded to minio and a job is sent to nats.io

_Environment variables:_

- API_PORT: The port to listen
- STORAGE_BUCKETNAME: The bucket's name to use and to create in minio
- STORAGE_SERVER: The minio's access url
- STORAGE_PORT: The minio's access port
- STORAGE_ACCESSKEY: The minio's access key
- STORAGE_SECRETKEY: The minio's secret key
- STORAGE_SSL: Define if minio is using SSL
- NATS_CONNECTIONSTRINGS: Nats.io connection url
- NATS_QUEUE: The name of the queue

Docker image: alexandrevilain/image-resizer-jobs-api

### Worker:

The worker gets resize jobs from nats.io, resize the image, saves it to minio and store metadata on postgres.

_Environment variables:_

- NATS_SERVERS: The connection stringS to connect to nats.io
- NATS_QUEUE: The name of the queue
- POSTGRES_CONNECTION_STRING: Postgres connection url
- STORAGE_SERVER: The minio's access url (with port)
- STORAGE_PUBLIC_ENDPOINT: The minio's public url (For k8s, this is the ingress). Don't put the scheme at the beginning. It uses the STORAGE_SSL env var to know if it's https or http.
- STORAGE_BUCKETNAME: The bucket's name to use and to create in minio
- STORAGE_ACCESSKEY: The minio's access key
- STORAGE_SECRETKEY: The minio's secret key
- STORAGE_SSL: Define if minio is using SSL

Docker image: alexandrevilain/image-resizer-worker

### Images-API:

The image-API is using PostgREST. It returns all resized images from PostgreSQL.
The config file is in its dedicated folder.

## Helpful informations:

To start Postgres using docker:

```bash
docker run --name images \
  -e POSTGRES_PASSWORD=supinfo \
  -e POSTGRES_USER=supinfo \
  -e POSTGRES_DB=images \
  -v $(pwd)/tables.sql:/docker-entrypoint-initdb.d/tables.sql \
  -p 5432:5432 \
  -d postgres
```

To start Nats.io using docker:

```bash
docker run --name gnatsd \
  -p 4222:4222 \
  -p 8222:8222 \
  -p 6222:6222 \
  -d nats:latest
```

To start minio using docker:

```bash
docker run --name minio \
  -e MINIO_ACCESS_KEY=supinfo \
  -e MINIO_SECRET_KEY=supinfo1234 \
  -p 9000:9000 \
  -v /tmp/minio:/data \
  -d minio/minio server /data
```

To start image-api using docker:

```bash
docker run --rm --link images:db -p 3000:3000 \
  -e PGRST_DB_URI="postgres://supinfo:supinfo@db:5432/images" \
  -e PGRST_DB_ANON_ROLE="web_anon" \
  -e PGRST_DB_SCHEMA="api" \
  -d postgrest/postgrest
```

Or you can use the file `images-api/postgrest.conf`.
