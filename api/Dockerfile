# ベースイメージを指定
FROM golang:1.21

# 作業ディレクトリを設定
WORKDIR /app

# ホストのカレントディレクトリの内容をコンテナの作業ディレクトリにコピー
COPY . .

# Goアプリをビルド
RUN go build -o main main.go

# ポートをエクスポート
EXPOSE 8080

# アプリを実行
CMD ["./main"]
