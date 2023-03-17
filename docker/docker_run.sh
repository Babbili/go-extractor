#!/bin/bash

docker build -f docker/Dockerfile --tag goget .

docker run --name goget-c goget
