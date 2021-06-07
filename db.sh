#!/bin/bash
docker run --name some-postgres -e POSTGRES_PASSWORD=pass123 -p 5432:5432 -d postgres:13.2
