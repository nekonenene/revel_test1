# Revel であれこれ

## Start the web server:

```
revel run revel_test1
```

そして http://localhost:9000 へ


## 気付き

* GORM、CRUDがあんまり直感的な書き方じゃないように思える。[SQLBoiler](https://github.com/volatiletech/sqlboiler)のほうがその点はうまく出来てる。
* `Gorm, err := gorm.Open` だとスコープが関数内になるのはハマりどころだった。  
`Gorm, err = gorm.Open` とするのが正しい。
* Railsの迷わず作れる良さを改めて感じるのだった……


## メモ

* `go fmt ./**/*.go` をたまにして矯正


## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites


## リンク集

### Revel

* [GitHub](https://github.com/revel/revel)
* [チュートリアル](http://revel.github.io/tutorial/gettingstarted.html)
* [ガイド](http://revel.github.io/manual/index.html)
* [サンプル集](http://revel.github.io/examples/index.html)
* [ドキュメント](https://godoc.org/github.com/revel/revel)

### GORM

* [GitHub](https://github.com/jinzhu/gorm)
* [ガイド](http://jinzhu.me/gorm/)
