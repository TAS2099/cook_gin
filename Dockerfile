# ベースとなるDockerイメージ指定
FROM golang:latest

# 環境変数
ENV GOPATH=
ENV GO111MODULE=on

# ホストマシン上の作業ディレクトリをコンテナ環境にコピー
COPY cmd/main.go /go
COPY go.mod go.sum ./

# ginパッケージ、gorm、mysqlをinstall
RUN go get github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm
RUN go get github.com/go-sql-driver/mysql
RUN go mod download

# コンテナログイン時のディレクトリ指定
WORKDIR /go

# コンテナ起動時に実行するコマンド
CMD ["go","run", "/go/cmd/main.go"]



