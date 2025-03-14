# 環境構築手順

## 前提

- `go 1.24.0` がインストールされていること
- `docker` がインストールされていること

## 初回のみ

### sql-migrate をインストール

```sh
go install github.com/rubenv/sql-migrate/...@latest
```

### swag をインストール

```sh
go install github.com/swaggo/swag/cmd/swag@latest
```

### wire をインストール

```sh
go install github.com/google/wire/cmd/wire@latest
```

### 環境設定ファイルを配置

```sh
cp .env.example .env
```

### OpenAI APIの設定

`.env`ファイルにOpenAIのAPI Keyを設定

```
OPENAI_API_KEY=<ここにAPIKEYを入力>
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
