# git-desc
これはgit branch --edit-descriptionなどでブランチに説明を追加するコマンドの代わりとなるCLIツールです。

## Environment
現状、macのみで、go getでのインストールのみなので、goもインストールされている必要があります。

## Install
```
go get github.com/miyazi777/git-desc
```

## Usage
### 現在のブランチに説明を追加
以下のコマンドで環境変数$EDITORに設定されているエディタが起動し、編集したテキストが現在のブランチの説明になります。
```
git-desc set
```

また、-mオプションを付けることでエディタを起動せず、説明を設定できます。
```
git-desc set -m "branch description"
```

### 現在のブランチの説明と関連するwebページのURLを表示
```
git-desc info
```

### 全ブランチの説明を表示
```
git-desc list
```

### 現在のブランチに関連するwebページを登録
以下のコマンドで環境変数$EDITORに設定されているエディタが起動し、編集したテキストが現在のブランチの関連ページとなります。
```
git-desc page set
```

また、-mオプションを付けることでエディタを起動せず、webページを設定できます。
```
git-desc page set -m "web page url"
```

### 現在のブランチに関連するwebページをブラウザで開く
```
git-desc page open
```

## config
$HOME/.config/git-desc/config.yamlに設定ファイルを置くと、コマンド実行時にそこから設定を読み込み、設定に沿った動きをします。
設定内容は以下のとおりです。

|キー|内容|
|---|---|
|editor|set, page setコマンドを実行した時に起動するエディタ|
|list.branch.width|listコマンドのbranchの表示幅で、ターミナルの表示幅に対して0.0〜1.0でパーセンテージで指定します|
|list.description.width|listコマンドのdescriptionの表示幅で、ターミナルの表示幅に対して0.0〜1.0でパーセンテージで指定します|

### 設定例
```
editor: nvim
list:
  branch:
    width: 0.35
  description:
    width: 0.5
```
