# ベースになる docker イメージを指定
FROM golang:1.17.1-alpine

# 作業ディレクトリの作成
RUN mkdir /go/src/app

# 作業ディレクトリの設定
WORKDIR /go/src/app

# ホストのファイルをコンテナの作業ディレクトリにコピー
ADD . /go/src/app