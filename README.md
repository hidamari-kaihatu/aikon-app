# aikon-app

## 学童連絡帳アプリ　あいこん

#### 私たちひだまり開発における最初の開発アプリです

##### 開発者用Dockre関係の注意事項

###### docker-compose.yml
platformはosがぞれぞれ違うから必要に応じてコメントアウト(M1に必要)
    platform: linux/x86_64 
###### git clone してきたら最初に打つコマンド
```docker-compose
docker-compose run -w /usr/src/app/next --rm app-next npm install
```
その後に
```docker-compose
docker-compose up -d
```
##### 開発者用GitHub関係の注意事項
- ファイルが未完成でコミットしておくときは、コード内に//TODO:として自分のやり残していることを書き記しておく。
- リモートリポジトリのブランチについて
    - main（完成品の置き場）
    - dev（みんなの編集したものを合わせてテストをするためのブランチ）
- ローカルリポジトリのブランチについて
    - 作業するときは必ずmain(master)からブランチを切る！
    
    - ブランチ名は作業内容を表す
    ex) change-tite, update-api-students,
    - margeしたら消すかは各自の判断
- 最新の状態をリモートからpullするときはmain (master)に戻ってから
  リモートのdevブランチからpullするコマンド
  ```git
  git pull origin dev
  ```
- devへのPRについて
ながれとしては以下のように！

⚠️かならずブランチをmainからきっておく

```git
git add ファイル名

git commit -m “コミット名”

git push origin HEAD
```
・originは今いるところ、HEADはpush先で、HEADは今のブランチの別名
・つまり、git push origin HEADは、ローカルの今いるブランチの内容をgithubの同名のリモートブランチにpushするということ
・git pushするときに、devをかいちゃうと、devにpushされてしまい、PR（確認してmergeできなくなってしまうので、devはかかないでね）

↓

git push origin HEADをターミナルでうったら、githubのPRのページに行く。

**そこで、PR先を、mainからdevに変えてください！**

- コンフリクトが発生したとき・・・他の作業は一旦止めて、関係者はコンフリクトの解消＆マージしきる

  1. git status コマンドで確認

  2.GitHub上で確認

  3.コンフリクトの原因となっている箇所を特定

  4.ファイルを正しく修正

  5.修正したファイルを　git add

  6.パラメーターなしでコミット　git commit

  7.リモートリポジトリへpush.    git push origin ブランチ名

  8.PRを確認する