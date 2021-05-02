# Test restaurant app on golang

## Installation:
run: docker-compose up -d

After this command, two containers will be raised:
- app (container with rest api application)
- mongo (container with mongodb storage)

## Available endpoints:

- POST http://127.0.0.1:9093/bills/{tableNumber} 
- POST http://127.0.0.1:9093/assign-table 
- POST http://127.0.0.1:9093/order/{tableNumber} 

#### postman collection for this api attached to project `postman_collection.json`.

## Additional information:
- there are no unit tests & integration tests (sorry about it)
- there are no lint issues
