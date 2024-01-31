## Table of Contents

- [ABOUT](#about)
- [LAUNCH](#launch)
- [TEST](#tests)
- [DIAGRAM](#diagram)


## ABOUT
#### migration
make migrate-up
- It should be executed after the PostgreSQL container is created.

#### swagger
http://localhost:8080/swagger/index.html

## TESTS
#### unit tests
make unit-test

#### repository tests
make repository-test  <br>


#### api tests
make api-test  <br>
- Docker-compose and http server must be running.


## DIAGRAM
![Alt text](./doc/diagram.png?raw=true "High level diagram")


