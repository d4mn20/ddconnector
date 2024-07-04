# DDConnector
The “ddconnect” service is responsible for receiving and storing security scan results (SAST, SCA, and DAST) sent via HTTP requests from Azure Pipelines. It is part of a larger system that integrates security scans into a DevSecOps environment and connects to DefectDojo and OpenAI (ddmitigator) for processing and managing the results.

# Features
Receives security scan data via HTTP requests from Azure Pipelines.
Stores scan results in a database for processing within DefectDojo.

# Configuration
Define the port on which the service will run in the <router.go> file and ensure it is accessible over the network.

# Usage
To use the “ddconnector” service, follow these steps:

## Clone this repository:
```bash
git clone https://dev.azure.com/bbts-lab/DevSecOps/_git/ms-ddconnector
```

Configure environment variables and configuration files according to your implementation.
Run the microservice:
```go
go run main.go
```

AI-generated code. Review and use carefully. More info on FAQ.
Send HTTP requests containing scan results to the appropriate endpoint.
The service will process the requests and store the results in the configured database.

## Docker
To build and run the Docker container for “ms-ddconnector,” use the following commands:

```bash
sudo docker build -t ms-ddconnectori .
sudo docker run --name ms-ddconnector -d -e DD_API_KEY=<API_KEY> -p 21777:21777 ms-ddconnectori
```

# License
This microservice was developed by BBTS and is licensed under the MIT License.

# To Do
[ ] Add “Environment” field to import and reimport.
