# backend/Dockerfile
FROM node:18

# 作業ディレクトリを作成
WORKDIR /app

# db.json ファイルをコピー
COPY db.json /app/db.json

# json-server をインストール
RUN npm install -g json-server

# json-server をポート 3001 で起動
CMD ["json-server", "--watch", "db.json", "--host", "0.0.0.0", "--port", "3001"]