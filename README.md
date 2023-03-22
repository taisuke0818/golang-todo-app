# golang-todo-app

## 実行環境

- goenv
- docker
- go-task
- direnv

## 環境構築

以下のコマンドにて以下の開発環境を構築する

```bash
task init
direnv allow
```

### 構築物

- buf:v1.15.1
- ko:v0.13.0
- go:1.19.5

## 開発

```bash
# リント
task lint

# フォーマット
task format

# ビルド（ko）
task build

# gRPCコード生成（buf）
task generate
```

## ローカルサーバ

```bash
# 起動
task start

# 停止
task down

# 起動確認
curl -XPOST http://localhost/todo.v1.TodoService/ListTodoTasks -v
```

## CLI アプリ

```bash
# タスク一覧取得
go run ./todo-cli/ list

# タスク作成
go run ./todo-cli/ create
```
