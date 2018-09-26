# Image-resizer: sample app

This project is a demo app used for Kubernetes and Docker Swarm deployment examples.

There isn't any Swarm or Kubernetes manifest in this repository, you'll need to create them. This project require Postgres, nats.io and Minio. You'll need to create deployments for those dependencies too. In the `Helpful informations` section, you will find docker commands to start Postgres, nats and Minio.

## Project architecture:

![Architecture](docs/architecture.png?raw=true 'Architecture')

## Services configuration:

### Frontend:

To build the frontend you first need to edit the `.env.prod` file. Then you can build the project by running `yarn build`.

Environment variables (at build time):

VUE_APP_JOB_API: The Jobs-API url
VUE_APP_IMAGES_API: The Images-API url

### Jobs-API:

The jobs-API is a REST API providing image upload endpoint to create a resize job.
When you send an image on `/upload` its uploaded to minio and a job is sent to nats.io

Environment variables:

API_PORT: The port to listen
STORAGE_BUCKETNAME: The bucket's name to use and to create in minio
STORAGE_SERVER: The minio's access url
STORAGE_PORT: The minio's access port
STORAGE_ACCESSKEY: The minio's access key
STORAGE_SECRETKEY: The minio's secret key
STORAGE_SSL: Define if minio is using SSL
NATS_CONNECTIONSTRINGS: Nats.io connection url
NATS_QUEUE: The name of the queue

### Worker:

The worker gets resize jobs from nats.io, resize the image, saves it to minio and store metadata on postgres.
Environment variables:

NATS_SERVERS: The connection stringS to connect to nats.io
NATS_QUEUE: The name of the queue
POSTGRES_CONNECTION_STRING: Postgres connection url
STORAGE_SERVER: The minio's access url (with port)
STORAGE_BUCKETNAME: The bucket's name to use and to create in minio
STORAGE_ACCESSKEY: The minio's access key
STORAGE_SECRETKEY: The minio's secret key
STORAGE_SSL: Define if minio is using SSL

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
  minio/minio server /data
```
