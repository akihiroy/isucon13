
target=$(ls -1t /tmp/slow_query_*.digest |head -n 1)
echo file is ${target}
cat ${target}