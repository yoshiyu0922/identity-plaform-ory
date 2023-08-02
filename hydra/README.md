# Quick Start


- OAuth2 Client を作成する
```shell
curl -i -X POST \
   -H "Content-Type:application/json" \
   -d \
'{
  "client_name": "yoshikawa_test2",
  "redirect_uris": ["http://127.0.0.1:5555/callback"],
  "grant_types": ["authorization_code","refresh_token"],
  "response_types": ["code", "id_token"],
  "scope": "openid offline",
  "token_endpoint_auth_method": "client_secret_post",
  "skip_consent": true
}' \
 'http://localhost:4445/admin/clients'
```

- 以下のコマンドを実行するとトークンを取得できる
  - `client id` と `client secret` は上記で作成した Client の ID と Secret を設定する
```shell
docker exec -it db-hydra-1 \
    hydra perform authorization-code \
    --client-id (client id) \
    --client-secret (client secret)\
    --endpoint http://127.0.0.1:4444/ \
    --port 5555
```