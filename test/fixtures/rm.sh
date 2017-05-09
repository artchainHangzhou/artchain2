docker kill $(docker ps -qa)
docker rm $(docker ps -qa)
docker network rm $(docker network ls -q)
docker rmi $(docker images | grep "\dev-peer[0-9]*\-" | tr -s ' ' | cut -d ' ' -f 1 | less)
rm -rf /var/hyperledger/couchdb*
