
build:
	go build -ldflags="-X 'github.com/zaoshi00/examples-go/EmbedVersionInfo/version.BuildTime=`date -u '+%Y-%m-%d'`'\
		-X 'github.com/zaoshi00/examples-go/EmbedVersionInfo/version.GoVersion=$(shell go version)' \
		-X 'github.com/zaoshi00/examples-go/EmbedVersionInfo/version.GitHash=$(git rev-parse --short HEAD)' \
	"