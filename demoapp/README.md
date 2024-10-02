# Demo APP

This is a Golang + React demo cloud-native app. No bullshit, only production-tested code.

Features:

* Backend
  * ✅ Golang
  * ✅ GORM
  * ✅ Dependency injection
  * ✅ Error handling
  * ✅ Logging
  * ✅ Unit/integration testing (examples only)
  * ✅ Data fixture generator
  * ✅ Docker
  * ⌛ Authentication with Auth0
  * ❌ CICD
  * ❌ S3 / GCS
  * ❌ O11y (metrics, Prometheus, Grafana)
* Frontend
  * ✅ TypeScript
  * ✅ React
  * ✅ create-react-app
  * ✅ react router
  * ✅ mui/joy
  * ✅ react-hooks
  * ✅ unstated-next
  * ❌ Authentication with Auth0
  * ❌ Unit testing

## Running locally

### Pre-requisites

Install [Golang](https://github.com/moovweb/gvm) and [Node](https://github.com/nvm-sh/nvm).

### Install packages

~~~shell
make install
~~~

### Prepare env vars

~~~
cd apps/api
cp .env.example .env.local
# populate the values
cd ../dashboard
cp .env.example .env
~~~

### Run infra

~~~shell
make run_local_infra
~~~

### Prepare database

~~~shell
cd apps/api
make migrate_db
make seed
~~~

### Run API app

~~~shell
make run app=api
~~~

### Run Dashboard app

~~~shell
make run app=dashboard
~~~

### Further improvements

* Deploy somewhere :)
* Alerting and O11y with Prometheus, Grafana and OpsGenie
* Logging with Grafana Loki
* Better support for Swagger rebuilds, possible migration to Protobuf
