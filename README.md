# Turion Space Take Home

This coding assignment consist of 3 parts. Each part is its own application and it can run independent from each other. The first 2 parts are backend services one for ingesting Telemetry data provided by a server (code provided on the problem statement) and a Telemetry API service which is a REST API that can interact with the data the Ingestion service has collected and stored in a PostgreSQL DB instance.

## Included functionality
This repository contains functionality to support the execution of 4 applications and the database needed to run them. The apps included in this repository are:

* `PostgreSQL DB instance` which will be run in a Docker container. For this database, a migration step is included so that the tables are created and ready to be used by the applications. Both `make migrate-up` and `make migrate-down` actions are available as Makefile commands.

* `Telemetry Server`. This is provided in problem statement, it will transmit UDP packages at a cadence of 1 second and it is to be consumed by the Telemetry Ingest Service

* `Telemetry Ingestion Service` is an application that will consume the UDP packages sent by the Telemetry Server. This application will not only receive and parse the messages, but also store the package data into the PostgreSQL database.

### Notes:
* A very simplified structure has been used to create the Telemetry Ingestion Service.
* During implementation, simple logging has been added. In a complete system, a more robust logging library shall be used or be written.

## How to Run the Apps

There are a couple of ways to run the applications. They can be executed locally as well as with the aid of docker-compose. We'll describe both ways below.

### Running the Apps Locally

In order to run the apps locally, we have a Makefile available with the necessary commands to get up and going. Before running the applications, it is necessary to spin up a PostgreSQL container and migrate the initial tables. 

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
