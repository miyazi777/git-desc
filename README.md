# git-desc
これはgit branch --edit-descriptionなどでブランチに説明を追加するコマンドの代わりとなるCLIツールです。

## Install
```
go get github.com/miyazi777/gitdesc
```

## Usage
### 現在のブランチに説明を追加
```
gitdesc set "説明"
```

### 現在のブランチの説明を表示
```
gitdesc info
```

### 全ブランチの説明を表示
```
gitdesc list
```
