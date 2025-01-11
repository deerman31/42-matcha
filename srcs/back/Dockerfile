# ビルドステージ
FROM golang:1.23.4-bullseye AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y git && rm -rf /var/lib/apt/lists/*

COPY ./golang-echo/go.mod ./golang-echo/go.sum ./

RUN go mod download

COPY ./golang-echo/ .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 実行ステージ
# 実行ステージであるため、golangのイメージをベースにする必要はない
FROM debian:bullseye-slim

# 必要なパッケージのインストールとクリーンアップ（オプション）
# セキュリティ: ca-certificatesは安全なHTTPS接続に必要です。
# 最小化: --no-install-recommendsで推奨パッケージのインストールを避けます。
# クリーンアップ: /var/lib/apt/lists/*の削除でイメージサイズを削減します。
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    #netcat-openbsd \
    && rm -rf /var/lib/apt/lists/*

# セキュリティのために非root ユーザーを作成
# 目的: セキュリティ強化のために、特権のない専用ユーザーを作成します。
# セキュリティ: rootでアプリケーションを実行するリスクを軽減します。
# 最小権限の原則: アプリケーションに必要最小限の権限のみを与えます。
# 詳細: -mオプションでホームディレクトリを作成します。
RUN useradd -m appuser

# 作業ディレクトリを設定
WORKDIR /app

# ビルドステージからバイナリをコピー
# 目的: ビルドステージで作成されたバイナリを実行環境にコピーします。
# 最小化: ビルド環境の依存関係を含めずに、実行可能ファイルのみをコピーします。
# マルチステージビルド: --from=builderで前段階の成果物を参照します。
COPY --from=builder /app/main .

# 非root ユーザーに切り替え
USER appuser

# /app/uploads/images
RUN mkdir -p /home/appuser/uploads/images1
RUN mkdir -p /home/appuser/uploads/images2
RUN mkdir -p /home/appuser/uploads/images3
RUN mkdir -p /home/appuser/uploads/images4
RUN mkdir -p /home/appuser/uploads/images5

RUN mkdir -p /home/appuser/uploads/default_image
COPY ./default_image/napoleon.png /home/appuser/uploads/default_image

COPY ./default_image/male1.png /home/appuser/uploads/images1
COPY ./default_image/male2.png /home/appuser/uploads/images1
COPY ./default_image/male3.png /home/appuser/uploads/images1
COPY ./default_image/male4.png /home/appuser/uploads/images1
COPY ./default_image/male5.png /home/appuser/uploads/images1
COPY ./default_image/female1.png /home/appuser/uploads/images1
COPY ./default_image/female2.png /home/appuser/uploads/images1
COPY ./default_image/female3.png /home/appuser/uploads/images1
COPY ./default_image/female4.png /home/appuser/uploads/images1
COPY ./default_image/female5.png /home/appuser/uploads/images1


EXPOSE 3000

# コンテナ起動時に実行されるコマンド
#CMD ["./golang-echo"]
CMD ["./main"]