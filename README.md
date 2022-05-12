# state-estimation-api
API to configure and launch new state-estimation jobs

## Development

### Compiling and running

    $ go build
    $ go run main.go

## Configuration

Configuration parameters can be specified in the confg.yaml file:

```yml  
"address": "localhost:6379",
"password": "", 
"database": 0
 ``` 