# Turion Space Take Home

This coding assignment consist of 3 parts. Each part is its own application and it can run independent from each other. The first 2 parts are backend services one for ingesting Telemetry data provided by a server (code provided on the problem statement) and a Telemetry API service which is a REST API that can interact with the data the Ingestion service has collected and stored in a PostgreSQL DB instance.

## Included functionality
This repository contains functionality to support the execution of 4 applications and the database needed to run them. The apps included in this repository are:

## PostgreSQL DB instance 
PostgreSQL instance which will be run in a Docker container. For this database, a migration step is included so that the tables are created and ready to be used by the applications. Both `make migrate-up` and `make migrate-down` actions are available as Makefile commands.

## Telemetry Server
This is provided in problem statement, it will transmit UDP packages at a cadence of 1 second and it is to be consumed by the Telemetry Ingest Service

## Telemetry Ingestion Service
Application that will consume the UDP packets sent by the Telemetry Server. This application will:
* Listen for UDP packets containing spacecraft telemetry.
* Decodes CCSDS-formatted packets according to provided structure
* Validates telemetry values against defined ranges.
* Persist data to a PostgreSQL database.
* Implements an alerting mechanism for out-of-range values in 2 ways:
  * Using Log-Based alerting, logging an `[ALERT]` level log message.
  * Saving the entries in the DB and flagging the entries with anomalies.
* Both alerting mechanisms can be used to produce alerts on dashboard tools like the ELK stack, Grafana or Prometheus

## Telemetry API
Application that retrieves data gathered by the Telemetry Ingestion Service. This API was built using `gorilla/mux` even though this router library is deprecated. The explanation can be found on the notes below.

The endpoints included are documented and can be seen at `http://localhost:8080/docs/index.html`. In an ideal situation, we would be able to see examples for requests and responses and even try the endpoints from that page.

Instead I will provide a set of curl commands that can be used to test the endpoints:

```shell
# For the current data set
curl localhost:8080/api/v1/telemetry/current | jq '.'
```

```shell
# For data between 2 points in time
# Start and End times are optional. Ommitting either of them will remove that boundary
curl "http://localhost:8080/api/v1/telemetry?start_time=2024-12-04T16:04:00&end_time=2024-12-04T16:05:00" | jq '.'
```

```shell
# For anomalies data between 2 points in time
# Start and End times are optional. Ommitting either of them will remove that boundary
curl "http://localhost:8080/api/v1/telemetry/anomalies?start_time=2024-12-04T16:04:00&end_time=2024-12-04T16:05:00" | jq '.'
```

## Telemetry Frontend
This is a simple React app that has 2 panels:
* One showing the current Telemetry data that is updated in real time. This panel also turns orange when an anomaly is detected.
* The bottom panel contains historical data. This data is fetched when the page is loaded, however, I added a button to `Refetch Data` so that we can get the latest data. I added a limit of 200 entries per request but this can be adjusted if needed.

### Notes:
* A very simplified structure has been used to create the Telemetry Ingestion Service.
* During implementation, simple logging has been added. In a complete system, a more robust logging library shall be used or be written.
* For simplicity, these endpoints do not have pagination. This will be necessary because of the amount of data entries.
* The idea of using Gorilla/Mux was to make a quick implementation and integrate with Swagger UI and the OpenAPI spec to showcase it. However, my efforts were in vain because the libraries for SwaggerUI are not playing well. But you can partially visualize how the documentation UI would show. With more time, I would have learned how to integrate these libraries with Fiber/Echo.

## How to Run the Apps

### Running the Apps Locally

In order to run the apps locally, we have a Makefile available with the necessary commands to get up and going. Before running the applications, it is necessary to spin up a PostgreSQL container and migrate the initial tables. 

### Note: It is recommended to run each app separate in a distinct terminal window

* Spin up a PostgreSQL container. This will create the PG container, create the schema and migrate the tables. These steps are available to be run separately as `make pg`, `make pg-schema` and `make migrate-up`
  ```shell
  # This command will run all 3 DB steps
  make local-db  
  ```

* Run the `Telemetry Server`: This is the server that will be sending UDP packages with Telemetry data.
  ```shell
  make run-server
  ```
* Run the `Telemetry Ingestion Service`: This is the service that consumes UDP packages with Telemetry data comming from the Telemetry Server.
  ```shell
  make run-ingestion-service
  ```
  
* Run the `Telemetry API`: This REST API retrieves the telemetry data collected by the Telemetry Ingestion Service.
  ```shell
  make run-telemetry-api
  ```
  
* Run the `Telemetry Frontend`: This is the React app that displays the data retrieved from the Telemetry API
  ```shell
  make start-frontend
  ```
