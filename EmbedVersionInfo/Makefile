
VersionPath=github.com/zaoshi00/examples-go/EmbedVersionInfo/version

buildTimeFlag=-X '$(VersionPath).BuildTime=`date -u '+%Y-%m-%d %H:%M:%S'`'
gitFlag=-X '$(VersionPath).GitHash=`git rev-parse HEAD`'
goVersion=-X '$(VersionPath).GoVersion=`go version`'

build:
	go build -ldflags="$(buildTimeFlag) $(gitFlag) $(goVersion)"

