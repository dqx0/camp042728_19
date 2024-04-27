# ベースイメージを指定します。Node.jsを使用します。
FROM node:latest

# 作業ディレクトリを設定します。
WORKDIR /frontend

# ローカルのpackage.jsonとpackage-lock.jsonをコピーし、依存関係をインストールします。
COPY ../../frontend/package*.json ./
RUN npm install

# ローカルのソースコードをコンテナにコピーします。
COPY ../../frontend/ .

# Reactアプリケーションのビルドを実行します。
RUN npm run build

# ポート番号を指定します。
EXPOSE 3000

# アプリケーションを起動します。
CMD ["npm", "start"]