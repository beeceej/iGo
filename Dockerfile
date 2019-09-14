## Generate the protobuf artifacts
FROM namely/protoc-all:1.23_0 AS pbgen
LABEL MAINTAINER=beeceej.code@gmail.com
ADD proto ./proto
ADD igo/igopb ./igo/igopb
RUN entrypoint.sh \
  -d ./proto \
  -l go \
  -o igo/igopb

## Build The iGo binaries (igod, igoclient)
FROM golang:1.13.0 AS binbuild
WORKDIR /iGo
ADD cmd ./cmd
ADD igo ./igo
ADD go.mod go.sum ./
COPY --from=pbgen /defs/igo/igopb igo/igopb

RUN go get golang.org/x/tools/cmd/goimports
RUN go get ./...
RUN CGO_ENABLED=0 go build -a -ldflags '-w -extldflags "-static"' -o igod cmd/igod/main.go
RUN CGO_ENABLED=0 go build -a -ldflags '-w -extldflags "-static"' -o igoclient cmd/igoclient/main.go

## Copy the binaries from binbuild for usage
FROM alpine
COPY --from=binbuild /iGo/igod bin/igod
COPY --from=binbuild /iGo/igoclient bin/igoclient
COPY --from=binbuild /go/bin/goimports /bin/goimports

ENTRYPOINT [ "sh", "-c" ]
CMD [ "igod" ]
