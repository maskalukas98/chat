## Table of Contents

- [ABOUT](#about)
- [LAUNCH](#launch)
- [TEST](#tests)
- [DIAGRAM](#diagram)


## ABOUT
**migration** <br>
make migrate-up
- It should be executed after the PostgreSQL container is created.

**swagger** <br>
http://localhost:8080/swagger/index.html

## TESTS
**unit tests** <br>
make unit-test

**repository tests** <br>
make repository-test  <br>


**api tests** <br>
make api-test  <br>
- Docker-compose and http server must be running.


## DIAGRAM
![Alt text](./doc/diagram.png?raw=true "High level diagram")


