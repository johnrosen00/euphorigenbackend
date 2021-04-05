export GOOS=linux
go clean
go build
docker build -t johnrosen00/eugateway .
go clean

docker push johnrosen00/eugateway


docker rm -f eugateway
docker pull johnrosen00/eugateway


export TLSKEY="/etc/letsencrypt/live/api.johnrosen.me/privkey.pem"
export TLSCERT="/etc/letsencrypt/live/api.johnrosen.me/fullchain.pem"


export REDISADDR="redisServer:6379"
export DSN="root:test@tcp(eudb:3306)/store"
export SESSIONKEY="bnottomtext"
export MANPASS="eleven"
#docker run -d -p 443:443 --net api -v /etc/letsencrypt:/etc/letsencrypt:ro -e REDISADDR=$REDISADDR -e DSN=$DSN -e SESSIONKEY=$SESSIONKEY -e TLSKEY=$TLSKEY -e TLSCERT=$TLSCERT --name eugateway johnrosen00/eugateway
docker run -d -p 443:443 --net api -e REDISADDR=$REDISADDR -e MANPASS=$MANPASS -e DSN=$DSN -e SESSIONKEY=$SESSIONKEY -e TLSKEY=$TLSKEY -e TLSCERT=$TLSCERT --name eugateway johnrosen00/eugateway