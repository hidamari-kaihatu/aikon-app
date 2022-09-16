#　aikon-app

## 学童連絡帳アプリ　あいこん

## 私たちひだまり開発における最初の開発アプリです

## 開発者用Dockre関係の注意事項
### git clone してきたら最初に打つコマンド
```docker-compose
docker-compose run -w /usr/src/app/next --rm app-next npm install
```
その後に
```docker-compose
docker-compose up -d
```
### 今はnpm installなどのRUNコマンドなしの状態だけど、これを入れて立ち上げられるように後々やる。


## 開発者用GitHub関係の注意事項
