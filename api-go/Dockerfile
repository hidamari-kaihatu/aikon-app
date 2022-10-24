FROM golang:1.18-alpine
# アップデートとgitとvimのインストール
RUN apk update && apk add git && apk add vim && apk --no-cache add make gcc g++ musl-dev binutils autoconf automake libtool pkgconfig check-dev file patch
# appディレクトリの作成
RUN mkdir /go/src/api-go
# ワーキングディレクトリの設定
WORKDIR /go/src/api-go
# ホストのファイルをコンテナの作業ディレクトリに移行
# やっぱり COPY ./api-go /go/src/api-goだとデプロイできない
COPY . .

VOLUME /go/src/api-go
EXPOSE 8080
ENV GOOGLE_APPLICATION_CREDENTIALS aikon-570be-firebase-adminsdk-bh8yt-73dc8a9673.json

# RUN go mod tidy
# GoApp起動
CMD ["go","run","."]