# バックエンド

## 技術スタック

言語  **GO**

フレームワーク  **Echo**

## バックエンドのディレクトリ構成
````
 backend
   ├── .devcontainer            // dev containerの設定
   ├── .vscode                  // デバッグ用
   ├── Dockerfile               // dockerコンテナの設定
   └── golang
       ├── .env                 // 環境変数
       ├── db                   // DBとの接続やマイグレーション
       ├── go.mod
       ├── go.sum
       ├── infrastructure       // サーバー起動用
       │   ├── cmd              // main.go
       │   ├── di               // 依存性注入用
       │   ├── entity           // DBのテーブル定義
       │   └── router           // エンドポイント設定用
       ├── internal             // API作成用
       │   ├── auth             // 認証の共通関数
       │   ├── authMiddleware   // 認証ミドルウェア
       │   ├── handler
       │   ├── model            // レスポンスやリクエストのスキーマ定義
       │   ├── repository
       │   ├── usecase
       │   └── util             // 共通関数
       ├── tests                // テストコード用
       │   ├── internal
       │   │   ├── handler
       │   │   ├── repository
       │   │   └── usecase
       │   └── testutil         // テストで使用する共通関数
       └── validate             // バリデーション設定用

````

## 環境構築

### デバッグ
１．左メニューバーの**実行とデバッグ(Ctrl + Shift + G)** を選択

２．**デバックの起動** もしくは **F5** でデバックが起動

３．以下の画像の通りになれば起動成功

![aaa](https://res.cloudinary.com/zenn/image/fetch/s--BYdfS68U--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_1200/https://storage.googleapis.com/zenn-user-upload/deployed-images/25899a9ddcfd4e6a27bdacc4.png%3Fsha%3Df33afdac342ea3e25717b557c9f751a5437bf235)


