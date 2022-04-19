## [ローカル下におけるKongでのルーティング検証に関する手順書、およびルーティングがうまく行かない際の原因と対策について]



## 手順1：

検証用のgolangのディレクトリとファイルを用意する。

#### [具体例]

go_kong_practiceディレクトリを用意して、その下にUser_apiディレクトリとClient_apiディレクトリを用意。

双方のディレクトリにはそれぞれmain.goファイルを作成する。


## 手順2： 

下記のKongの公式サイトに入り、KongをmacOSにインストールする。

https://docs.konghq.com/gateway/2.8.x/install-and-run/macos/#download-and-install

また、下記のコマンドの通りに実行しても良い。

#### --コマンド--

 brew tap kong/kong
 
 brew install kong
 
 
## 手順３：

下記のコマンドを実行してKongがインストールされているか確認する。

#### --コマンド--

kong


## 手順4：　

下記のKongの公式サイトに入る。

なお今回は「without a database」で検証する。

https://docs.konghq.com/gateway/2.8.x/install-and-run/macos/#without-a-database


## 手順5： 

検証したいgolangディレクトリ下にて下記のコマンドを実行してkong.ymlファイルを生成する。

#### --コマンド--
 
 kong config init


## 手順6: 

下記のURLに進み、内容をコピー。

その後、検証したいgolangディレクトリ下にてkong.confファイルを作成してコピー分を貼り付ける。


https://raw.githubusercontent.com/Kong/kong/master/kong.conf.default


## 手順7：　

作成したkong.confファイルを下記のように編集する。

この時、declarative_config部分に記載するkong.ymlファイルへのパスに関しては絶対パスとして、編集したい下記の部分に関してはコメントアウトを解除する。
 
 database = off　-- 966行目
 
 declarative_config = /path/to/kong.yml　-- 1150行目
 
 
## 手順8: 

kong.ymlファイルにルーティング情報を記載する

#### [注意点]

services以下にルーティング情報を記載していくが、ルーティングに関しては一つのservices下にまとめて記載すること。

二つのservicesを用意すると先述のkong.confファイルのdeclarative_config部分に記載したkong.ymlファイルへうまく繋ぐことが出来ず、エラーの原因となる。

#### 成功例：

services:

 name: user-api
  
  url: http://localhost:8002
  
  tags:
  
  example
 
  routes:
  
  name: user-api-routes
    
    paths:
    
    /user-api
  
  name: client-api
  
  url: http://localhost:8003
  
  routes:
  
  name: client-api-routes
    
   　paths:
    
    /client-api


#### 失敗例:

services:

 name: user-api
  
  url: http://localhost:8002
  
  
  tags:
  
  example
  
  routes:
  
  name: user-api-routes
    
    paths:
    
    /user-api
    
 **services:**  **--失敗の原因--**
  
  name: client-api
  
  url: http://localhost:8003
  
  routes:
  
  name: client-api-routes
    
    paths:
    
    /client-api


## 手順9: 
下記のコマンドを実行して、Kongを起動する。

#### --コマンド--
 
 kong start -c kong.conf


## 手順10:
検証したいgolangディレクトリ（User_apiディレクトリ、またはClient_apiディレクトリ）へ移動して、下記のコマンドでwebサーバーを起動させる。

#### --コマンド--

go run ./main.go


## 手順11:
ブラウザ上で起動させたwebサーバーからのレスポンスを確認する。

http://localhost:8000/user-api/users

または

http://localhost:8000/client-api/clients


## 手順12:

検証が済んだら、Kongをストップする。

その時、下記のコマンドを実行するとKongがストップする。

#### --コマンド--

kong stop


## 手順13:最終ステップ

起動させたwebサーバーを停止させる。

その時、「control + c」といったショートカットキーを実行する。
