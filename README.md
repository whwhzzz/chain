To run this exmaple, from the root of this project:

docker build -t dropbox .
docker run --publish 6060:8081 --name dropbox --rm dropbox
