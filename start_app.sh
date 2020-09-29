# !/bin/bash

# # MySQLサーバーが起動するまでmain.goを実行せずにループで待機する
# until mysqladmin ping -h mysql --silent; do
#   echo 'waiting for mysqld to be connectable...'
#   sleep 2
# done
sleep 10

echo "app is starting...!"
exec go run server/cmd/main.go