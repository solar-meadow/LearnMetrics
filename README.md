# Learn metrics


This repository contains code for a system that generates metrics using the node-exporter, collects the data using Prometheus, and displays it in a visually appealing format using Grafana.

Prerequisites
```
    Docker
    docker-compose
```
# Usage
To start the server, navigate to the root directory of the repository in your terminal and enter the following command:

```
make compose

```

This command will build the necessary Docker images and start the containers for Grafana, Prometheus, and the node-exporter. You can access Grafana at localhost:3000, Prometheus at localhost:9090, and the node-exporter metrics endpoint at localhost:9100/metrics. The web application is accessible at localhost:8080/metrics.

To stop the server, enter the following command:

```
make down
```
This command will stop and remove the containers, networks, and volumes created by docker-compose up.