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
make run_infra
~~~

### Run API app

### Run Dashboard app

~~~shell
make run app=dashboard
~~~