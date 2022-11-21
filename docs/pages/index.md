# POCs, demos, katas

## Project list

| Project 	       | Public URL 	 | Description      |
|-----------------|---------------|------------------|
| service_name  	 |  	            | App description. |

## Local launch

1. Clone the repo, `cd` to the folder.
2. Run `make install` to install all local dependencies for every application.

### Localstack profile (if needed)

~~~
mkdir ~/.aws
printf "[pocs-demos-katas]\naws_access_key_id=doesnt-matter\naws_secret_access_key=doesnt-matter\nregion=eu-central-1\n" >> ~/.aws/credentials
~~~

### Env file

1. Copy the env file template: `cp ./.env.example ./.env.local`
2. Fill the variables up

### Running all services

1. Run `make run_infra` to launch local infrastructure.
2. Wait until the infrastructure is ready.
3. If not done before, in another terminal run `make create_resources` to create the resources in the Localstack.
4. If not done before, run `make seed_database` to fill the database with some data.
5. In another terminal run: `make run app=service_name` to launch the API microservice.

## CI/CD

Before running CI/CD make sure that infrastructure was pre-created.

### Secrets

So far we do not use any software like Vault for managing secrets.
The following secrets should be obtained/generated and then added as [environment variables on GitHub](https://github.com/gannochenko/pocs-demos-katas/settings/environments):

* `SOME_ENV_VAR`

The other env vars come from the infrastructure.
