## このリポジトリは？

構築したMariadbへの接続を確認するためのアプリケーション

webサーバからdbサーバに接続できることを確認するために作成した簡易的なアプリケーションであり、単純にsqlオープンできるかだけを確認対象にしている。

## 使い方

適宜環境に合わせてコードを変更する。

### webサーバのListenポート

使用しているポート番号とバッティングしないよう、以下を変更。

```golang
log.Fatal(http.ListenAndServe(":8080", nil))
```

### dbサーバの情報

以下のユーザ、パスワード、IPアドレス（dnsも可）を書き換える。

```golang
db, err := sql.Open("mysql", "mysqluser:mysqlpassword@tcp(x.x.x.x:3306)/mysql")
```