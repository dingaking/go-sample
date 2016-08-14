# go-sample
go lang basic samples

## Set PATH
export PATH=$PATH:/usr/local/go/bin;<br />

## mkdir & set PATH
mkdir src; mkdir pkg; mkdir bin;<br />
export GOPATH=\`pwd\`;<br />
cd src;


## GOPATH error chack
Go install always fails no install directory outside GOPATH<br /><br />

check go env<br />

mkdir bin <br />
export GOBIN=$GOPATH/bin<br /><br />

```html
    GOPATH/
        bin/
        src/
        pkg/
```


source code in GOPATH/src/source_code.go<br />
go get command in GOPATH/src

