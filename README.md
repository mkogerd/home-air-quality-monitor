Welcome to the Awair Element Exporter Docker setup!

This repository provides a Docker Compose file ([docker-compose.yml](docker-compose.yml)) to run Prometheus and Grafana containers alongside an Awair Element Exporter container. The exporter fetches data from Awair, which is then scraped by Prometheus for visualization in Grafana.

## Running the Setup

To get started:

1. Clone this repository using `git clone`.
2. Create a new directory named `.volumes` to store persistent volumes.
3. Run the following command to start the services:

```
docker-compose up -d
```

This will create and start all containers in detached mode.

## Monitoring

The setup exposes Prometheus on host port 9090, allowing you to access its web interface at `http://localhost:9090`. The Grafana server is exposed on host port 3000, accessible at `http://localhost:3000`.

To visualize the Awair data in Grafana:

1. Open a web browser and navigate to `http://localhost:3000`.
2. Log in with the default admin credentials (`admin` for both username and password).
3. Navigate to the "Connections" tab.
4. Add a new Prometheus connection using `http://prometheus:9090`
5. Navigate to the "Dashboards" tab.
6. Search for the "Awair Metrics" dashboard.

## Configuration

The [docker-compose.yml](docker-compose.yml) file uses a single environment variable:

- `AWAIR_HOST`: Set this variable to your Awair device hostname or IP address to enable data fetching.

You can adjust these values in the `.env` file (not shown, but assumed to exist) or by modifying the [docker-compose.yml](docker-compose.yml) file directly.

## Troubleshooting

If you encounter issues, refer to the following resources:

- Docker Compose documentation: https://docs.docker.com/compose/
- Prometheus documentation: https://prometheus.io/docs/
- Grafana documentation: https://grafana.com/docs/

Feel free to open an issue or ask questions on GitHub if you need further assistance!
