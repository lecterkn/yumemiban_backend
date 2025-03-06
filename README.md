# yumemiban_backend

## 開発環境

- 言語
  - Go
- データベース
  - MySQL
  - Redis

## プロジェクト構成

[アーキテクチャについて](./docs/arch.md)

```
root
|  README.md    # readme、困ったときは読むように
|  dbconfig.yml # マイグレーション時に参照するデータベースの接続情報
|
|- docs # readmeから派生したドキュメントやswagger
|  | arch.md      # アーキテクチャについて
|  | startup.md   # 環境構築手順書
|  | swagger.yaml # swaggerのyaml形式のドキュメント
|  | swagger.json # swaggerのjson形式のドキュメント
|
|- cmd
|  | main.go # 実行ファイル
|
|- internal
|  |- app
|     |- common     # JWTの暗号化・復号化などの様々な場所で利用する共通の関数やロジックを集約
|     |- di         # wireによる依存性注入を行う
|     |- entity     # Domain層 エンティティを格納
|     |- handler    # Presentation層 HTTPのエンドポイントとなる関数を役割ごとにhandlerに分けて定義
|     |  |- request    # HTTPのリクエストボディを格納
|     |  |- response   # HTTPのレスポンスボディを格納
|     |- port       # Domain層 リポジトリなどのインターフェイスを格納
|     |- repository # Infrastructure層 リポジトリの実装を格納
|     |  |- mysql      # mysqlのコネクタやリポジトリの実装
|     |  |- redis      # redisのコネクタやリポジトリの実装
|     |- usecase    # Application層    ユースケースを役割ごとに分けて定義
|     |  |- input      # ユースケースの入力データを格納
|     |  |- output     # ユースケースの出力データを格納
|
|- migrations # データーベースのマイグレーションファイルを格納
|  | ...
|
```
