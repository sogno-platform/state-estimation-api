# state-estimation-api
API to configure and launch new state-estimation jobs

## Development

### Compiling and running

    $ go build
    $ go run main.go

## Configuration

Configuration parameters can be specified in the confg.yaml file:

```yml  
database:
  address: "localhost:1234"
  password: ""
  database: 1234

amqpBroker:
  host: "localhost"
  port: 5672
  user: ""
  password: ""

``` 