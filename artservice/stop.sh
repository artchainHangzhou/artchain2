ARTSERVICEPID=`ps -ef | grep 'artservice' | grep -v grep | awk '{print $2}'`
echo $ARTSERVICEPID
kill -15 $ARTSERVICEPID
