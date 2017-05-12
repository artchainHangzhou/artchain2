NPMPID=`ps -u $USER| grep 'npm' | grep -v grep | awk '{print $1}'`
NODEPID=`ps -u $USER| grep 'node' | grep -v grep | awk '{print $1}'`
echo $NPMPID, $NODEPID
kill -15 $NPMPID $NODEPID
