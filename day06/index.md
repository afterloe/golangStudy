# 第六条

```sbtshell
curl --request POST \
  --url 'http://127.0.0.1:8082/?name=afterloe' \
  --header 'Cache-Control: no-cache' \
  --header 'Content-Type: application/x-www-form-urlencoded' \
  --header 'Postman-Token: 57315e28-b94f-df3f-721c-127450eac187' \
  --data 'name=joe&age=12'
```

```sbtshell
curl --request POST \
  --url 'http://127.0.0.1:8081/v1/user?name=afterloe' \
  --header 'Cache-Control: no-cache' \
  --header 'Content-Type: application/x-www-form-urlencoded' \
  --header 'Postman-Token: 9dc41a2e-6f9a-9f92-d3e8-a316055f9956' \
  --data 'name=joe&age=12'
```

```sbtshell
curl --request POST \
  --url http://127.0.0.1:8081/ \
  --header 'Cache-Control: no-cache' \
  --header 'Content-Type: application/x-www-form-urlencoded' \
  --header 'Postman-Token: 3c45ca83-d791-c3ba-9144-33dc87070627' \
  --data age=12
```

```sbtshell
curl --request POST \
  --url http://127.0.0.1:8081/ \
  --header 'Cache-Control: no-cache' \
  --header 'Content-Type: application/x-www-form-urlencoded' \
  --header 'Postman-Token: 30a9abe9-a19c-db3a-484d-3b7efe6d55bf' \
  --data 'name=joe&age=12'
```
这里推荐scrypt方案，scrypt是由著名的FreeBSD黑客Colin Percival为他的备份服务Tarsnap开发的。 目前Go语言里面支持的库http://code.google.com/p/go/source/browse?repo=crypto#hg%2Fscrypt
