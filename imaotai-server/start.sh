export ENV="dev"
kill -s 9 `pgrep imaotai_helper`
chmod 777 imaotai_helper
nohup ./imaotai_helper &
