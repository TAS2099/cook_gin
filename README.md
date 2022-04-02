# cook_gin
## 概要
料理のレシピを投稿するアプリを作ろうと思っています。(cookpad的なもの)

現在は、todoアプリの雛形を作成し終えた状態です。

これから、これをもとにユーザーログインやデータベースのテーブルを増築したりしたいと考えています。

## ディレクトリ構成
<pre>
.
├── Dockerfile
├── cmd
│   └── main.go
├── db
│   └── my.cnf
├── docker-compose.yml
├── go.mod
├── go.sum
├── model
│   └── model.go
├── my.cnf
└── templates
    ├── detail.html
    └── index.html
</pre>
