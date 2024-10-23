# Demo APP

This is a Golang + React demo cloud-native app. No bullshit, only production-tested code.

It's based on (but does not strictly follow) [the Swagger-petstore demo API](https://github.com/swagger-api/swagger-petstore).

Under the hood:

* Backend
  * ✅ Golang
  * ✅ GORM
  * ✅ Dependency injection
  * ✅ Error handling
  * ✅ Logging
  * ✅ Unit/integration testing (examples only)
  * ✅ Data fixture generator
  * ✅ Docker
  * ✅ Authentication with Auth0
  * ❌ RBAC
  * ❌ Metrics instrumentation
  * ❌ CICD
  * ❌ S3 / GCS
  * ❌ O11y (Prometheus, Grafana)
  * ❌ Loadtest
* Frontend
  * ✅ TypeScript
  * ✅ React
  * ✅ create-react-app
  * ✅ react router
  * ✅ mui/joy
  * ✅ react-hooks
  * ✅ unstated-next
  * ✅ Authentication with Auth0
  * ❌ CICD
  * ❌ RBAC
  * ❌ Notifications https://mui.com/joy-ui/react-snackbar/
  * ❌ Unit testing
  * ❌ Accessibility (aria-*)
  * ❌ E2E testing (Cypress or Playwright)
  * ❌ I18n

## Running locally

### Pre-requisites

* [Docker](https://www.docker.com/products/docker-desktop/)
* [Golang](https://github.com/go-nv/goenv)
  * [godotenv](https://github.com/joho/godotenv)
  * golang migrate: `brew install golang-migrate`
* [Node + Yarn](https://github.com/nvm-sh/nvm)
* psql: `brew install postgresql`

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
