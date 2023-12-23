pt-query-digest /var/log/mysql/mysql-slow.log > /tmp/slow_query_$(date +%Y%m%d%H%M%S).digest
target=$(ls -1t /tmp/slow_query_*.digest |head -n 1)
echo file is ${target}
cat ${target}