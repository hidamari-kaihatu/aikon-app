FROM golang:1.18-alpine
# アップデートとgitとvimのインストール
RUN apk update && apk add git && apk add vim
# appディレクトリの作成
RUN mkdir /go/src/api-go
# ワーキングディレクトリの設定
WORKDIR /go/src/api-go
# ホストのファイルをコンテナの作業ディレクトリに移行
# やっぱり COPY ./api-go /go/src/api-goだとデプロイできない
COPY . .

VOLUME /go/src/api-go
EXPOSE 8080
RUN go mod tidy
# GoApp起動
CMD ["go","run","main.go"]