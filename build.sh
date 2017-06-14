export GOPATH="$(pwd)"
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

go get github.com/sadbox/mediawiki
go get github.com/PuerkitoBio/goquery
go get github.com/shiyou0130011/p2mfmt

go build -o p2m
