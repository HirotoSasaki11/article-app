#!/bin/sh

cd `dirname $0`/..
wd=`pwd`

cd ${wd}
docker-compose down -v
docker-compose up