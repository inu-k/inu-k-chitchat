# inu-k-chitchat
GoとReactを用いたフォーラムサイトの作成

## setup
.envファイルを作成し、postgresのユーザー名とパスワードを記述してください。
```.env
DB_USER=<postgresのユーザー名>
DB_PASSWORD=<postgresのパスワード>
```

```bash
$ sudo /etc/init.d/postgresql start # postgresの起動
$ createdb chitchat
$ cd backend
$ psql -f data/setup.sql -d chitchat  # テーブルの作成、初期データの挿入
$ go build
$ ./inu-k-chitchat  # localhost:8999でサーバーが起動
```

(別のターミナルで)
```bash
$ cd ../frontend
$ npm install
$ npm start  # localhost:3000でフロントエンドが起動
```

## 機能
- 作成されたスレッドの一覧表示
- スレッドに投稿された投稿の表示

現状ではスレッドの作成、投稿の作成はできず、データベースにあるデータを表示するだけの状態です。
