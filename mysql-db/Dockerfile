# FROM mysql:8-debian

# RUN apt-get update -y && \
#     apt-get install -y locales && \
#     echo "ja_JP.UTF-8 UTF-8" > /etc/locale.gen && \
#     locale-gen && \
#     update-locale LANG=ja_JP.UTF-8 
# # CentOS Linuxの日本語設定
# ENV LANG ja_JP.UTF-8
# # ubuntuの日本語設定(これも上と一緒な気がする・・・いらないのでは？)
# # ENV LANGUAGE ja_JP:ja

# #ポートを開ける
# EXPOSE 3306

# #MySQL設定ファイルをイメージ内にコピー
# # ADD ./my.cnf /etc/mysql/conf.d/my.cnf
# # mysql8.0を使用する場合の認証のための引数
# CMD ["mysqld", "--character-set-server=utf8", "--collation-server=utf8_unicode_ci","--default-authentication-plugin=mysql_native_password"]
FROM mysql:8-debian

RUN apt-get update -y && \
    apt-get install -y locales && \
    echo "ja_JP.UTF-8 UTF-8" > /etc/locale.gen && \
    locale-gen && \
    update-locale LANG=ja_JP.UTF-8 

ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL=ja_JP.UTF-8

CMD ["mysqld", "--character-set-server=utf8", "--collation-server=utf8_unicode_ci"]