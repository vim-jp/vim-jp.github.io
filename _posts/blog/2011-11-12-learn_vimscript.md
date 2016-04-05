---
layout: blog
category: blog
title: Learn Vim script
redirect_from:
  - /2011/11/12/learn_vimscript.html

---

Vim scriptは他の言語と比べ癖があり、Vimの仕様と直結している為、他の言語と同等の機能が無かったり、もしくはやり方が異なったりするものも中にはあります。
しかしながらVim scriptはVim本体を強力にカスタマイズ出来る本来の言語でもあります。
以前、Vimハッカーでもお馴染みのthincaさんが書かれた「[Vimスクリプト基礎文法最速マスター](http://d.hatena.ne.jp/thinca/20100201/1265009821)」をvimの`:help`から参照出来る様にhelp形式にしてみました。
pathogen-vim/bundle系プラグインでVimプラグインを管理されておられる方であれば

    cd ~/.vim/bundle
    git clone https://github.com/mattn/learn-vimscript.git

でインストール出来ます。Vim scriptの文法等で分からない事があれば`:help learn-vimscript`を実行してみて下さい。
