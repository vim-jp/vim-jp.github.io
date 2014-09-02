---
layout: post
category: blog
title: vim-go-extra を公開致します。
---
golang オフィシャル配布物として提供されてきた misc/vim という vim プラグインが、開発対象から外すという理由により先日リポジトリから削除されました。
その変更を受けて vim-jp ではそのコピーを `go-vim` というリポジトリ名で公開しておりましたが、本日それを `vim-go-extra` という名称に変更致しました。
以下これまでの流れ。

* golang オフィシャルリポジトリから misc/vim が削除される
* vim-jp が go-vim として misc/vim のコピーを配布
* Google が [vim-ft-go](https://github/google/vim-ft-go) というリポジトリで misc/vim の一部を公開する
* vim 本体リポジトリに vim-ft-go がマージされる

[vim-ft-go](https://github/google/vim-ft-go) は misc/vim の一部のみが含まれています。misc/vim からは以下のコマンドが削除されました。

* `:Godoc` コマンド
* `:Fmt` コマンド
* `:Import` コマンド

これらのコマンドは Fatih Arslan 氏が提供する  [vim-go](https://github.com/fatih/vim-go) に名称を変えて含まれていますが、vim-go に含まれている syntax 等は vim のリポジトリに含まれている物とバッティングしています。
そこで vim-jp では上記の消えたコマンド群のみを集めた vim-go-extra というリポジトリを提供する事にしました。

[https://github.com/vim-jp/vim-go-extra](https://github.com/vim-jp/vim-go-extra)

vim-go-extra の位置づけは以下の通りです。

* vim 本体に含まれる golang 関連ファイルとバッティングしない
* golang 関連コマンドをインストールしていれば自動でコマンドが使える様になる
* ファイル保存時等に自動で `:Fmt` したりする様な動作はユーザが混乱するのでデフォルトとしない

というポリシーで公開してまいります。なお、gocode、godef、lint コマンドがインストール済みの状態であれば、各コマンドに付属されている vim プラグインが自動で使える様になっています。
何か問題が発生した再は、上記 vim-go-extra の issues までご連絡下さい。
