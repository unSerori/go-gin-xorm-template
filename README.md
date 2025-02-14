# hoge

※go-gin-xorm-templateリポジトリ。goのAPIサーバのテンプレート。

## 概要

hogeのGo APIサーバ。

### 環境

Visual Studio Code: 1.88.1  
golang.Go: v0.41.4  
image Golang: go version go1.22.2 linux/amd64  

## 環境構築

1. 以下のDocker環境を作成  
[リポジトリURL](https://github.com/unSerori/docker-hoge)  
SSH URL:  

    ```SSH:SSH URL
    git@github.com:unSerori/hoge-api.git
    ```

2. ここまでが1の内容（フォルダーをVScodeで開きgo_serverをVScodeアタッチ。）
3. リポジトリをクローン

    ```bash
    # カレントディレクトリにリポジトリの中身を展開
    git clone git@github.com:unSerori/hoge-api.git .
    
    # developブランチに移動
    git switch develop
    ```

4. shareディレクトリ内で以下のコマンド。

    ```bash:Build an environment
    # vscode 拡張機能を追加　vscode-ext.txtにはプロジェクトごとに必要なものを追記している。  
    cat vscode-ext.txt | while read line; do code --install-extension $line; done
    ```

5. .envファイルをもらうか作成。[.envファイルの説明](#env)

## API仕様書

## エラー処理

APIがエラーを返す場合、詳細なエラーメッセージが含まれます。エラーに関する情報は[サーバーエラーコード](#server-error-code)を参照してください。　　

## SERVER ERROR CODE

サーバーレスポンスコードとして"srvResCode"キーで数値を返す。  
以下に意味を羅列。  

- 成功関連
  - 1000: Successful authentication.
  - 1001: Successful user registration.

- エラー関連
  - 7000: Authentication unsuccessful.
  - 7001: Failure to bind request.
  - 7002: There is already a user with the same primary key. Uniqueness constraint violation.
  - 7003: New user registration was not possible due to other DB problems.
  - 7004: Failure to hash passwords.
  - 7005: Failed to generate token.
  - 7006: New user registration was not possible due to other problems.

## .ENV

.evnファイルの各項目と説明

```env:.env
MYSQL_USER=DBに接続する際のログインユーザ名
MYSQL_PASSWORD=パスワード
MYSQL_HOST=ログイン先のDBホスト名。dockerだとサービス名。
MYSQL_PORT=ポート番号。dockerだとコンテナのポート。
MYSQL_DATABASE=使用するdatabase名
JWT_SECRET_KEY="openssl rand -base64 32"で作ったJWTトークン作成用のキー。
JWT_TOKEN_LIFETIME=JWTトークンの有効期限
```

## 開発者

- Author:[]
- Mail:[]
