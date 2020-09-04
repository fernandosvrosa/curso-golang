#!/bin/bash

docker run --name curso-golang -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=curso-golang -p 3306:3306 -d mysql:latest
