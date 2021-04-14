export GOOS=linux
go clean
go build
docker build -t johnrosen00/eugateway .
go clean

docker push johnrosen00/eugateway


docker rm -f eugateway
docker pull johnrosen00/eugateway



export TLSKEY="/usr/local/ssl/certs/api.lab.cip.uw.edu.key"
export TLSCERT="/usr/local/ssl/certs/api.lab.cip.uw.edu.crt"


export REDISADDR="redisServer:6379"
export DSN="root:test@tcp(eudb:3306)/store"
export SESSIONKEY="bnottomtext"
export MANPASS="CalvinGarfieldWilson"

docker run -d -p 443:443 --net api \
-e REDISADDR=$REDISADDR \
-e MANPASS=$MANPASS \
-e DSN=$DSN \
-e SESSIONKEY=$SESSIONKEY \
-e TLSKEY=$TLSKEY \
-e TLSCERT=$TLSCERT \
-v /usr/local/ssl/:/usr/local/ssl/:ro \
--name "eugateway" \
johnrosen00/eugateway
