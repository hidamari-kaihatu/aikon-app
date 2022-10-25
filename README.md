## スマートGAKUDO
スマートGAKUDOは、学童のICT化を進めて学童職員の負担を減らすための学童運営効率化アプリです。
#### 【ひだまり開発が初めて開発したアプリ】
<img width="678" alt="スクリーンショット 2022-10-24 10 19 31" src="https://user-images.githubusercontent.com/94016735/197431411-614ad0dc-acbf-4939-bbae-2dcb03d655ab.png">

### 機能
#### 管理者
- 新規登録
- ログイン
- 学童施設一覧の表示・編集
- 職員一覧の表示・編集

#### 保護者
- 新規登録
- ログイン
- 出欠連絡
- メッセージの受信
- 児童の入退室時刻の表示

#### 先生
- 新規登録
- ログイン
- 今日の児童一覧の表示
- 児童名簿一覧の表示
- メッセージの送信（音声入力）
- 送信履歴の確認
- 児童の入退室時刻の表示

### 使い方
##### 開発環境での使用方法
1.リポジトリのクローン
```bash
git clone  https://github.com/hidamari-kaihatu/aikon-app.git
```
2.ルートディレクトリへ移動
```bash
cd aikon-app
```
3.dockerコマンドを打つ
```docker-compose
docker-compose run -w /usr/src/app/next --rm app-next npm install
```
その後に
```docker-compose
docker-compose up -d --build
```
4.localhostで各リンクへ
```bash
http://localhost:3000/parents/parent-login
```
```bash
http://localhost:3000/teachers/teacher-login
```
```bash
http://localhost:3000/admin/admin-login
```
### 注意事項！！
アプリを立ち上げる先にFirebaseに関するエラーが出ることがあります。その際は、以下の手順で解消してください。
- ルートディレクトリのaikon-appで以下のコマンドを打つ。
```bash
docker-compose exec app-next bash
```
```bash
npm uninstall firebase
```
```bash
npm install firebase@9.10.0
```
### 使用技術
<img width="669" alt="スクリーンショット 2022-10-24 10 22 00" src="https://user-images.githubusercontent.com/94016735/197431685-f5e26402-56e5-4d1d-a1c8-09b6e385127c.png">

- [Next.js](https://nextjs.org/): 12.3.0
- [TypeScript](https://www.typescriptlang.org/): 4.8.3
- [Firebase](https://github.com/firebase/firebase-js-sdk#readme):9.10.0
- [React Speech Recognition](https://webspeechrecognition.com/):3.9.1
- [React Fook Form](https://www.react-hook-form.com):7.37.0
- [yup](https://github.com/jquense/yup):0.32.11
- [Golang](https://go.dev/):1.18
- [MySQL](https://www.mysql.com/jp/): 8.0
- [Docker](https://www.docker.com/): 20.10.17
- [Stripe](https://stripe.com/jp): 10.11.0
- [AWS EC2,RDS,ECR,ECS,fargate](https://aws.amazon.com/jp/)

#### ※ひだまり開発関係者用GitHub関係の注意事項※
- 今後のリファクタリングにおいても以下の流れでいきましょう！
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
⚠️かならずブランチをmainからきっておく

```git
git add ファイル名

git commit -m “コミット名”

git push origin HEAD
```
・originは今いるところ、HEADはpush先で、HEADは今のブランチの別名
・つまり、git push origin HEADは、ローカルの今いるブランチの内容をgithubの同名のリモートブランチにpushするということ
・git pushするときに、devをかいちゃうと、devにpushされてしまい、PR（確認してmergeできなくなってしまうので、devはかかない）

↓

git push origin HEADをターミナルでうったら、githubのPRのページに行く。

**そこで、PR先を、mainからdevに変えてください！**
