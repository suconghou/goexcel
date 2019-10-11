#!/bin/bash
make build && \
docker build -t=suconghou/tools:goexcel .
