# Pismo

Pismo helps handle accounts and transactions APIs for a bank/agency 

## Tech

- Language - [Golang]
- DB - [PostgreSQL]
- ORM - [GORM]
- Image - [Docker]

And of course Pismo itself is open source with a [public repository][pismo] on GitHub
## Development

[pismo.tf](https://github.com/preetamkv/pismo/blob/master/pismo.tf) is used to generate the DB cluster in Azure CosmosDB
[cmd/service](https://github.com/preetamkv/pismo/tree/master/cmd/service) has the logic for server initialization
[internal/app/pismo](https://github.com/preetamkv/pismo/tree/master/internal/app/pismo) has the business logic for the APIs
[intenal/pkg](https://github.com/preetamkv/pismo/tree/master/internal/pkg) has the helpers required for the business logic

## Running

Pismo requires [Golang] latest and [Docker] to run

DB is already created and [settings.json](https://github.com/preetamkv/pismo/blob/master/settings.json) needs to be updated with the host and password details
Update the settings.json file before running [build.sh](https://github.com/preetamkv/pismo/blob/master/build.sh)

[build.sh](https://github.com/preetamkv/pismo/blob/master/build.sh) file takes care of
- Building code
- Generating docker image
- Running the container using the latest image

### Command to Run
```sh
./build.sh
```

   [pismo]: <https://github.com/preetamkv/pismo>
   [Golang]: <https://go.dev/>
   [PostgreSQL]: <https://learn.microsoft.com/en-us/azure/cosmos-db/postgresql/introduction>
   [GORM]: <https://pkg.go.dev/gorm.io/gorm>
   [Docker]: <https://www.docker.com/>
