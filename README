# CLI tool to convert between 4 Byte ASN formats
Convert easily between ASPLAIN and ASDOT at CLI.

## Install
```
$ go get github.com/darylturner/as-convert
$ go build 
---or---
$ go install	# place binary in $GOPATH/bin
```

## Usage
```
$ as-convert 65535.7
4294901767
$ as-convert 4294901767
65535.7
```

or read from stdin
```
echo '4294901767' | as-convert
65535.7
echo '4294901767' | as-convert | as-convert
4294901767
```
