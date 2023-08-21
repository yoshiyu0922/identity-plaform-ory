# クライアントの作成方法

## 前提条件

- `just run` でコンテナを起動している

## API

- Reference: https://www.ory.sh/docs/hydra/reference/api

### クライアントの作成

以下のエンドポイントを実行してクライアントを作成する
```
POST http://localhost:4445/admin/clients
```
リクエストボディは以下の通り
- 詳細は[こちら](https://www.ory.sh/docs/hydra/reference/api#tag/oAuth2/operation/createOAuth2Client)
- レスポンスのクライアントシークレットはこのエンドポイント以外で取得できないので、この時点でメモしておく
```json
{
  "client_name": "yoshiyu0922",
  "redirect_uris": ["http://127.0.0.1:5555/callback"],
  "grant_types": ["authorization_code", "refresh_token"],
  "response_types": ["code", "id_token"],
  "scope": "openid offline",
  "token_endpoint_auth_method": "client_secret_post",
  "skip_consent": true
}
```

以下のエンドポイントを実行してクライアントが作成されていることを確認
- 詳細は[こちら](https://www.ory.sh/docs/hydra/reference/api#tag/oAuth2/operation/getOAuth2Client)
```
GET http://localhost:4445/admin/clients/(client id)
```