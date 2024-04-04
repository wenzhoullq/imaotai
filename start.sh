export ENV="dev"
kill -s 9 `pgrep zuoxingtao`
chmod 777 zuoxingtao
nohup ./zuoxingtao &