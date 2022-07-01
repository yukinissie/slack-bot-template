# Slack Bot Template

Go 製の SlackBot テンプレート

## 開発方法

### 1. インストール

```
git clone git@github.com:yukinissie/slack-bot-template.git
cd slack-bot-template/bot
make init
cd ..
```

### 2. SlackApp のボットトークンを設定

```
cp conf/dev/dev.yml.example conf/dev/dev.yml
```

dev.yml に SlackApp のボットトークンを設定する。

※SlackApp は[ここ](https://api.slack.com/apps/)で作る。

OAuth & Permissions > Scopes > Bot Token Scopes には channels:history と chat:write を追加する。

OAuth & Permissions の Bot User OAuth Token を slackToken に追記する。

```
slackToken: 'xoxb-...'
```

Event Subscriptions > Subscrive to be bot events に message.channels を追加する。

### 3. ローカルサーバーを起動

3000 ポートに ApiGateway をエミュレートしたサーバーが立ちます。

```
cd bot
make local_run
```

### 4. SlackApp とローカルサーバーを連携

ngrok を使って一時的にローカルサーバーを公開します。これは SlackApp と連携するためです。

```
ngrok http 3000
```

表示された TLS 化済みの URL を SlackApp の Event Subscriptions > Request URL に登録します。

例：`https://022d-131-129-4-95.ngrok.io/dev/example`

Go のコードを書き換えた場合は 3 を繰り返し実行します。

ngrok を一度でも落とした場合は 4 を繰り返します。

## デプロイ

AWS 上にデプロイする方法です。2 種類あります。うまくいく保証はありません。

- 手動デプロイ（dev 用）
- GitHub Actions でデプロイ（prod 用）

### 手動デプロイ（dev 用）

開発環境用の AWS リソースを自動的に作成し、デプロイします。

```
cd bot
make dev_deploy
```

### GitHub Actions でデプロイ（prod 用）

GitHub リポジトリの`Settings>Secrets`に`SLACK_BOT_TOKEN`、`AWS_ACCESS_KEY_ID`、`AWS_SECRET_ACCESS_KEY`、`SLACK_WEBHOOK`キーを登録します。

あとは`main`ブランチにプッシュすることでデプロイされます。
