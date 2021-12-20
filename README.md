# go RESTful API テンプレート

## 構成概要

Docker outside of docker を利用

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

`--parameter-overrides` で `template.yaml` の parameter を上書きできる

```sh
sam local start-api --host 0.0.0.0 --container-host host.docker.internal --debug --docker-volume-basedir $PWD --docker-network sample-network --parameter-overrides ParameterKey=AllowOrigin,ParameterValue="*"
```

### Cognito を使う場合

```sh

# ローカルから cognito を操作するためのシークレットキーを config に設定
vi ~/.aws/config
[profile sample]
aws_access_key_id = xxxxxxxxxxxxx
aws_secret_access_key = xxxxxxxxxxxxxxxxxxxxxxxx

# sam local の実行 Cognito を使う場合
sam local start-api --host 0.0.0.0 --container-host host.docker.internal --debug --docker-volume-basedir $PWD --docker-network sample-network  --profile sample --parameter-overrides ParameterKey=AllowOrigin,ParameterValue="*"
```
