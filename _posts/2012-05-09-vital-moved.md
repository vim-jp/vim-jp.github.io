---
layout: post
category: blog
title: vital.vim を vim-jp に移転しました。
---

vital.vim はモダンでエレガントな汎用ライブラリを目指すプロジェクトで、有名な物としては

* neocomplcache
* neobundle
* unite
* quickrun

といったプロジェクトで利用されています。これまでは ujihisa さんの個人リポジトリで開発されてきましたが今後は vim-jp からの提供物として管理していく事になります。
vital.vim は plugin's plugin、つまりプラグインが使うライブラリ的なプラグインであり、上記の様なプラグインを使っておられる方は意識せず使っている事になりますし、vital.vim は各プラグインに同梱される形式で配布されていますので、今回の移転において何か作業をして頂く事もありません。

ただし vital.vim のコードを追っかけておられる方は、これまの ujihisa さんのリポジトリ

    git://github.com/ujihisa/vital.vim

から

    git://github.com/vim-jp/vital.vim

に移動となりますのでご注意下さい。

### git スキーマを使っておられる方

    cd /path/to/your/.vim/bundle/vital.vim
    git remote set-url origin git://github.com/vim-jp/vital.vim 

### http スキーマを使っておられる方

    cd /path/to/your/.vim/bundle/vital.vim
    git remote set-url origin https://github.com/vim-jp/vital.vim 

vim-jp では vital.vim の開発に参加頂ける方を募集しています。
