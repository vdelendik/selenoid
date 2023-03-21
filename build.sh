export PATH=$PATH:/usr/local/go/bin

export GO111MODULE="on"
go install github.com/mitchellh/gox@latest # cross compile
CGO_ENABLED=0 gox -os "linux darwin windows" -arch "amd64" -osarch="darwin/arm64" -osarch="darwin/arm64" -osarch="linux/arm64" -osarch="windows/386" -output "dist/{{.Dir}}_{{.OS}}_{{.Arch}}" -ldflags "-X main.buildStamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.gitRevision=`git describe --tags || git rev-parse HEAD` -s -w"

#build again in current folder with s3 support:
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags s3,metadata

# ./selenoid is linux_amd64 binary with s3 support

docker build .

