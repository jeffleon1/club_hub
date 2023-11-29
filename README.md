## Franchise API with GRPC

This is a small application for the management of franchises created with __grpc__ that has two microservices, one for the management of the metadata of the franchises and the other in charge of the most revealing attributes of the franchise such as where it is located, website, name, etc.

This application has the particularity that thanks to __grpc-gateway__ it is also exposed in __http__ which allows to make requests through both protocols.

## Structure of this Repo

- Proto files: https://github.com/jeffleon1/club_hub/blob/main/api/franchise.proto
- Franchise microservice: https://github.com/jeffleon1/club_hub/tree/main/franchise
- Metadata microservice: https://github.com/jeffleon1/club_hub/tree/main/metadata
- Go api build: https://github.com/jeffleon1/club_hub/blob/main/Makefile


## To build locally 

You need a machine (linux) that has Go installed > 1.19 and make, you also need to change the `.env-sample` file to `.env` once you have done this you can run `make up_build`.

## How to make request

these are the two postman collections to test for [gRPC](https://lunar-shuttle-128744.postman.co/workspace/Duo-Din%C3%A1mico~2274368e-001a-4d64-8bba-5f040623782d/collection/656678b1e12291959a483034?action=share&creator=15046674) and [Http](https://lunar-shuttle-128744.postman.co/workspace/Duo-Din%C3%A1mico~2274368e-001a-4d64-8bba-5f040623782d/collection/15046674-00ef5288-1089-4b72-a78e-a8a8ada6fd98?action=share&creator=15046674)

also if you have grpcurl or curl you can use it here are some examples of how to test the application

### gRPC

```bash
grpcurl -plaintext -d '{"host":"marriot.com"}' localhost:8080 FranchiseService/CreateFranchise
```

```bash
grpcurl -plaintext -d '{"company_id":0}' localhost:8080 FranchiseService/GetFranchise
```

```bash
grpcurl -plaintext -d '{"franchise":{"name":"hilton"}, "id":1}' localhost:8080 FranchiseService/UpdateFranchise
```

```bash
grpcurl -plaintext -d '{"key":"url", "value":"marriott.com"}' localhost:8080 FranchiseService/GetFranchiseByFilter
```

### Http

```bash
curl --location 'localhost:8090/v1/franchises' \
--header 'Content-Type: application/json' \
--data '{
    "host": "marriot.com"
}'
```

```bash
curl --location 'localhost:8090/v1/franchise?company_id=0'
```

```bash
curl --location 'localhost:8090/v1/franchise?key=name&value=marriot'
```

```bash
curl --location --request PUT 'localhost:8090/v1/franchise/1' \
--header 'Content-Type: application/json' \
--data '{
    "franchise": {
        "name":"hilton"
    }
}'
```

### Note: 
how to install grpcurl
```bash
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```

