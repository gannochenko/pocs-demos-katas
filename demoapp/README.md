# Demo APP

## Running locally

### Pre-requisites

Install [Golang](https://github.com/moovweb/gvm) and [Node](https://github.com/nvm-sh/nvm).

### Install packages

~~~shell
make install
~~~

### Copy and modify .env.local

~~~
cp .env.example .env.local
~~~

### Run infra

~~~shell
make run_local_infra
~~~

### Run API app

### Run Dashboard app

~~~shell
make run app=dashboard
~~~

### Further improvements

* Deploy somewhere :)
* CI/CD
* Upload files directly to S3 or GCS
* Alerting and O11y with Prometheus, Grafana and OpsGenie
* Logging with Grafana Loki
* Proper migrations
* Better support for Swagger rebuilds, possible migration to Protobuf
