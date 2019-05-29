# gitdesc
これはgit branch --edit-descriptionなどでブランチに説明を追加するコマンドの代わりとなるCLIツールです。

## Install
```
go get github.com/miyazi777/git-desc
```

## Usage
### 現在のブランチに説明を追加
```
git-desc set "説明"
```

### 現在のブランチの説明を表示
```
git-desc info
```

### 全ブランチの説明を表示
```
git-desc list
```
