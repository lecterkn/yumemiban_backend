# 環境構築手順

## 前提

- `go 1.2.40` がインストールされていること
- `docker` がインストールされていること

## 初回のみ

### sql-migrate をインストール

```sh
go install github.com/rubenv/sql-migrate/...@latest
```

### 環境設定ファイルを配置

```sh
cp .env.example .env
```

## 実行

### データベースの立ち上げ

```sh
docker compose up -d
```

### マイグレーション

```sh
sql-migrate up
```

### サーバー起動

```sh
go run cmd/main.go
```
