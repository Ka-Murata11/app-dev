# 概要

本リポジトリは、WEBアプリ開発に必要な知識をつけるための勉強用WEBアプリです。  
フロントエンドは React/NEXT.js を用いて、APP Router で実装しています。サーバーコンポーネントやクライアントコンポーネントの理解、操作性や分かりやすさなどのユーザエクスぺリンスを高めるための知識を深めます。  
バックエンドは golang/echo フレームワークを用いた REST API です。ディレクトリ構成は、クリーンアーキテクチャを参考にしたものになっています。

# 使用言語

|カテゴリー|技術スタック|
| ---- | ---- |
| 開発環境 | docker, Dev Container |
| バージョン管理 | Git, GitHub|
| フロントエンド | TypeScript, React, Node.js, NEXT.js |
| バックエンド | Go, Echo |
| データベース | mysql |

# ディレクトリ構成

````plaintext
.
├── README.md               
├── backend
├── docker-compose.yml
├── frontend
└── json-server

```` 

# 起動方法

1. コマンドパレットから「開発コンテナ―: コンテナ―でフォルダーを開く」を選択

   ![alt text](/.image/image.png)

2. 編集したいコンテナを選択  
   ※ 選択していないコンテナも起動します

3. 以下の図の通りに表示できれば成功

   ![alt text](/.image/image2.png)