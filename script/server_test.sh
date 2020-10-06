#!/bin/sh

set -eu

cd `dirname $0`/..
wd=`pwd`

cd ${wd}/docker/mysql

sudo docker-compose up -d

cd ${wd}/server/test

go test .