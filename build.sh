cd  gincmd
GOOS=darwin GOARH=amd64 go build -o ../dist/ginsupercmd_mac
GOOS=linux GOARH=amd64 go build -o ../dist/ginsupercmd_linux
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ../dist/ginsupercmd_linux
docker build -t supercmd .

