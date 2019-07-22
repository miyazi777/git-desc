# git-desc
これはgit branch --edit-descriptionなどでブランチに説明を追加するコマンドの代わりとなるCLIツールです。

## Environment
現状、macのみしか動作しません。

## Install
```
brew tap miyazi777/git-desc
brew install miyazi777/gitdesc/git-desc
```

or

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

また、--only-listオプションを付けることでリストのみを表示します。

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

### 現在のブランチの説明と関連するwebページを削除
以下のコマンドで現在のブランチの説明と関連webページに関する情報を削除します。
```
git-desc delete
```

## config
$HOME/.config/git-desc/config.yamlに設定ファイルを置くと、コマンド実行時にそこから設定を読み込み、設定に沿った動きをします。
設定内容は以下のとおりです。

|キー|内容|
|---|---|
|editor|set, page setコマンドを実行した時に起動するエディタ|

### 設定例
```
editor: nvim
```

# fzfとの連携例
## ブランチの移動
fzfを使用しているのであれば、以下のようなスクリプトを.bashrcや.zshrcに記述することでブランチを確認しながら、移動することができます。
以下のスクリプトではターミナルでblと打ち込むとこのツールのlistからブランチを選択し、移動することが出来ます。

```zsh
# git desc setting
bl() {
  branches=$(git-desc list --only-list) &&
  select_line=$(echo "$branches" | fzf +m)
  git checkout $(echo ${select_line} | cut -d" " -f 1)
}
```


