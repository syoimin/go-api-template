# go RESTful API テンプレート

Docker outside of docker を利用した Go + SAM の RESTful API のサンプルです

## 技術要素

ランタイム・インフラ環境

-   Docker outside of docker ・・・ ローカル Docker 開発環境
-   AWS Lambda ・・・ Lambda のランタイム環境
    -   golang
-   Python ・・・ SAM CLI の実行環境
    -   AWS SAM cli

フレームワーク・モジュール

-   gin ・・・ go RESTFul API フレームワーク
-   migrate ・・・ DB マイグレーション（RDS を利用する場合）
-   go-outline ・・・ コードでバッグ用 （VSCode 専用かも）
-   mockgen ・・・ モック
-   godoc ・・・ GO ドキュメント生成

## SAM Local の実行

ローカル環境の確認用サーバの起動

SAM のコンテナに入って実行する

```sh
# コンテナへアクセス
docker exec -it sample-aws-sam /bin/sh

# ディレクトリ移動
cd ${VOLUME}

```

### Cognito を使わない場合

`--parameter-overrides` で `template.yaml` の parameter を上書きできる。

```sh
sam local start-api --host 0.0.0.0 --container-host host.docker.internal --debug --docker-volume-basedir $PWD --docker-network sample-network --parameter-overrides ParameterKey=AllowOrigin,ParameterValue="*"
```

### Cognito を使う場合

`--parameter-overrides` で `template.yaml` の parameter を上書きできる。

```sh

# ローカルから cognito を操作するためのシークレットキーを config に設定
vi ~/.aws/config
[profile sample]
aws_access_key_id = xxxxxxxxxxxxx
aws_secret_access_key = xxxxxxxxxxxxxxxxxxxxxxxx

# sam local の実行 Cognito を使う場合
sam local start-api --host 0.0.0.0 --container-host host.docker.internal --debug --docker-volume-basedir $PWD --docker-network sample-network  --profile sample --parameter-overrides ParameterKey=AllowOrigin,ParameterValue="*"
```
