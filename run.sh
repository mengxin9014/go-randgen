#! /bin/bash
# cases dsn1 dsn2 queries ignore_errors
help() {
  echo "Usage:"
  echo "run.sh cases dsn1 dsn2 queries ignore_errors"
  echo "Description:"
  echo "cases. the cases to run, for example 'all' 'case1 case2'."
  echo "dsn1. the tikv engine dsn, for example 'root:@tcp(127.0.0.1:4000)'"
  echo "dsn2. the tiflash engine dsn, for example 'root:@tcp(127.0.0.1:4000)'"
  echo "queries. random sql num generated by zz, if it is negative(like -1), exec subcommand will generate endless sql"
  echo "ignore_errors. ignore errors regex, for example 'llegal types of arguments|Invalid compare_bytes'"
  exit 1
}

if [ ! -n "$1" ]; then
  help
fi

set -ex
cases=$1
dsn1=$2
dsn2=$3
queries=$4
ignore_errors=$5

cases2=""
declare -A run_cases
if [ 'all' = "$cases" ]; then
  for f in cases/*.yy; do
    s=${f##*/}
    cases2="$cases2 ${s%.*}"
  done
else
  for c in $cases; do
    if [ ! -f "cases/${c}.yy" ]; then
      echo "cases/${c}.yy" not exists
      exit 2
    fi
    cases2="${cases2} $c"
  done
fi

for c in $cases2; do
  if [ -f "resource/resource/${c}.zz.lua" ]; then
    run_cases[$c]="resource/resource/${c}.zz.lua"
  else
    run_cases[$c]="resource/resource/default.zz.lua"
  fi
done

echo "begin run cases ${cases2}"
date_d=$(TZ=Asia/Shanghai date '+%F')

for c in "${!run_cases[@]}"; do
  db1=${c}1
  db2=${c}2
  hp1=${dsn1##*(}
  host1=$(echo ${hp1%)*} | awk -F ':' '{print $1}')
  port1=$(echo ${hp1%)*} | awk -F ':' '{print $2}')
  user1=${dsn1%%:*}

  hp2=${dsn2##*(}
  host2=$(echo ${hp2%)*} | awk -F ':' '{print $1}')
  port2=$(echo ${hp2%)*} | awk -F ':' '{print $2}')
  user2=${dsn2%%:*}
  mysql -u $user1 -h $host1 -P $port1 -e "drop database if exists $db1;create database $db1"
  mysql -u $user2 -h $host2 -P $port2 -e "drop database if exists $db2;create database $db2"

  ./go-randgen gendata --dsns "${dsn1}/${db1},${dsn2}/${db2}" -Z "${run_cases[$c]}"
  sleep 60
  dump=res_$c
  ci=false
  if [[ "$c" =~ "collation" ]]; then
    ci=true
  fi
  rm -rf $dump && ./go-randgen exec --skip-zz -Y cases/${c}.yy --dsn1 "${dsn1}/${db1}?tidb_isolation_read_engines=tikv" --dsn2 "${dsn2}/${db2}?tidb_isolation_read_engines=tiflash" -Q $queries --dump $dump --ci $ci
  date_t=$(TZ=Asia/Shanghai date '+%H-%M-%S')
  tar czf ${dump}.tar.gz $dump
  curl -F lilinghai/tiflash-test/go-randgen/${date_d}/${dump}-${date_t}.tar.gz=@${dump}.tar.gz http://fileserver.pingcap.net/upload
  download_file="http://fileserver.pingcap.net/download/lilinghai/tiflash-test/go-randgen/${date_d}/${dump}-${date_t}.tar.gz"
  if [ -z "$skip_erros" ]; then
    if [ ! -z "$(ls -A $dump)" ]; then
      echo "$c failed. detail faile log is $download_file"
    fi
  else
    if [ ! -z $(grep -EL "$skip_errors") ]; then
      echo "$c failed. detail faile log is $download_file"
    fi
  fi
done
