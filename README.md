# Guest List

## Summary
This API was created to handle tables and guests at an event, specifically the an end of year party. It features a layered design and a Swagger API specification for an extended API documentation.

## Architecture
The API was created with a layered architecture. The layers from highest to lowest are Handler --> Service --> Repository.

The handler layer is in charge of handling the HTTP request and returning the correct error codes, responses, etc. The service layer contains the buisness logic of the application. Lastly, the repository layer is in charge of interacting with the MySQL database. Each lower layer provides its services to the higher layer to manage communications and run the applications.

This architecture allows data from the upper layer to be passed to the lower layer through an interface. A layered architecture provides a clean-cut interface so that minimum information is shared among different layers. It also ensures that the implementation of one layer can be easily replaced by another implementation.

In addition, a global error handler wraps the handlers to provide a centralized place to handle errors.

## Application Handling
Standing on the root directory, start by moving to the golang directory:
```
cd golang
```
In this directory, you can run the Makefile commands to start the app, prune it, generate mock files, and run the unit tests. Read the sections below on how to run each case.

### Run the app
To start up the docker container with the mysql image, run:
```
make docker-up
```
This command uses the `docker-compose.yaml` file to start the application.

### Shut down and prune
To shut down the container and prune it, run:
```
make docker-down
```
This must be done after any changes to the MySQL schema at the file  `docker/mysql/dump.sql`.

### Generate Mocks
Upon changes to the repository or service interfaces, mock files must be regenerated for proper testing. 
To generate the mock files, run:
```
make generate-mocks
```
The ```mockgen``` package was used to generate the mock files.

### Run Tests
To run the unit tests for the handler and service packages, run:
```
make run-tests
```

## Documentation 
A Swagger API specification (`api-spec.yaml`) is included to detail the API endpoints, their parameters, and their responses. It can be visualized by opening it with the [Swagger Editor](https://editor.swagger.io/).