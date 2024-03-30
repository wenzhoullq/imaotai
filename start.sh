export ENV="dev"
chmod 777 zuoxingtao
ps -ef|grep zuoxingtao | awk -F' ' '{print $2}'| xargs kill -9
nohup ./zuoxingtao &