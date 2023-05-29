##  GO REST API
================

### - Server layer
Built with net/http web server + chi router + postgresql database

### - Service layer via custom handlers

### - Repository layer generated with sqlc
Find config in sqlc.yml

- Migrations handled with goland-migrate via docker

Lightweight containerized application ready for production

- Most commonly used commands, executed via Tasks
Find config in Taskfile.yml
More info about Task on https://taskfile.dev

- Performance profiles via pprof go package 
Find profiles on http://localhost:6060/debug/pprof/
