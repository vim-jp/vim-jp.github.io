---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2017 年 5 月号

---

## 話題

* [meguro.vim #3](https://megurovim.connpass.com/event/53990/) が 5/20 に開催されました
* [Fablic.vim #1](https://connpass.com/event/56803/) が 2017-05-22 に[開催されました](http://in.fablic.co.jp/entry/2017-05-fablic.vim)。次回は7月です。
*   [Udon.vim #0](https://connpass.com/event/55348/) が開催されました。

## 今月の新機能

今月の新機能及びユーザーに影響のある変更は以下のとおりです。

*   8.0.0590: `setqflist()` に `"context"` を指定できるようになりました。 `"context"` に指定した値は `getqflist()` で取得できます。
*   ランタイムの更新
    *   Scala Build Tool 用のファイルタイプ sbt が追加されました。

## Vimに関する脆弱性

特筆すべき脆弱性の報告はありませんでした。

## リリース情報

- [8.0.0589 : (after 8.0.0578) :simalt still does not work](https://github.com/vim/vim/commit/a21ccb7a974a96550c5cd99c4633e166d2083172)
- [8.0.0590 : cannot add a context to locations](https://github.com/vim/vim/commit/8f77c5a4ec756f3f866bd6b18feb6fca6f2a2e91)
- [8.0.0591 : changes to eval functionality not documented](https://github.com/vim/vim/commit/45d2cca1ea3f90fc70ad99d0c6812a9d8536303c)
- [8.0.0592 : if a job writes to a buffer screen is not updated](https://github.com/vim/vim/commit/29ae377ea7039874337bc79ace9ab2b37b9056e5)
- [8.0.0593 : duplication of code for adding a list or dict return value](https://github.com/vim/vim/commit/45cf6e910c6d162775ca9d470fac4b6db844001f)
- [8.0.0594 : (after 8.0.0592) build failure when windows feature is missing](https://github.com/vim/vim/commit/66c0e70b806cd8afdc9da9ddae11945ddf29841a)
- [8.0.0595 : (after 8.0.0590) Coverity warning for not checking return value](https://github.com/vim/vim/commit/beb9cb19c660484488a71a25eda46ab0fa579278)
- [8.0.0596 : crash when complete() called after complete&#x5f;add()](https://github.com/vim/vim/commit/4475b623960671898dac6a72b13a8d140402afa6)
- [8.0.0597 : off-by-one error in size computation](https://github.com/vim/vim/commit/253f9128779f315ea670f9b4a17446b7b4c74927)
- [8.0.0598 : building with gcc 7.1 yields new warnings](https://github.com/vim/vim/commit/9e0f6ec0762575d229b86798b284ca4876bc3d73)
- [8.0.0599 : diff mode is insufficiently tested](https://github.com/vim/vim/commit/79a213d6a4e909703524a7f1ad985c7bd40650a6)
- [8.0.0600 : test&#x5f;recover fails on some systems](https://github.com/vim/vim/commit/2a0b06def49198aef45ef173f524a638a700f9c9)
- [8.0.0601 : no test coverage for :spellrepall](https://github.com/vim/vim/commit/545cb79da586be3333a0a55616046f94b01f6b1a)
- [8.0.0602 : when gF fails to edit the file the cursor still moves](https://github.com/vim/vim/commit/2a79ed293c077ba791db962f0e121cf97ba4f7a7)
- [8.0.0603 : gF test fails on MS-Windows](https://github.com/vim/vim/commit/d7aca7a71c7254501c08f04db1798df479df0e89)
- [8.0.0604 : gF test still fails on MS-Windows](https://github.com/vim/vim/commit/712598f210570627534246cb5dcbb4f213367997)
- [8.0.0605 : the quickfix cached buffer may become invalid](https://github.com/vim/vim/commit/6dd4a53502fb4ec1b66104eab1805e7254ad9e41)
- [8.0.0606 : cannot set the context for a specified quickfix list](https://github.com/vim/vim/commit/6e62da3e14d32f76f60d5cc8b267059923842f17)

## 新着スクリプト

- [auto&#x5f;update&#x5f;cscope&#x5f;ctags&#x5f;database : 1:auto update cscope ctags database2:easy create database 3: auto plugn database](http://www.vim.org/scripts/script.php?script_id=5557)
- [vim-sass-colors : sass/scss color variable highlighting (works with imports)](http://www.vim.org/scripts/script.php?script_id=5558)
- [vim-yaml-folds : Folding rules for YAML, RAML and SaltStack SLS files](http://www.vim.org/scripts/script.php?script_id=5559)
- [VIM MySQL Suggestions : A plugin of VIM to show suggestions from MySQL database structure.](http://www.vim.org/scripts/script.php?script_id=5560)
- [bal.vim : Provides ballerina support to vim, when coupled with the bal syntax plugin. ](http://www.vim.org/scripts/script.php?script_id=5562)
- [vim-game-snake : snake game](http://www.vim.org/scripts/script.php?script_id=5563)
- [opensips.vim : Syntax highlighting for OpenSIPS script](http://www.vim.org/scripts/script.php?script_id=5564)
- [rst-robotframework-syntax-vim : Extends RST syntax to highlight Robot Framework code snippets](http://www.vim.org/scripts/script.php?script_id=5565)
- [tiny-todo : A simple todo list.](http://www.vim.org/scripts/script.php?script_id=5566)
- [bal.vim : Provides ballerina support to vim, when coupled with the bal ft script.](http://www.vim.org/scripts/script.php?script_id=5567)
- [vim-transit : Translates a given text from source language to target language via Google Trans](http://www.vim.org/scripts/script.php?script_id=5568)
- [botlang-syntax : Botlang support for vim](http://www.vim.org/scripts/script.php?script_id=5569)
- [quickEditTab : A new command to open files in the same or another tab.](http://www.vim.org/scripts/script.php?script_id=5570)
- [remaining-todos : There are remaining TODOs in this file!](http://www.vim.org/scripts/script.php?script_id=5571)
- [vim-prettier : A vim plugin wrapper for prettier](http://www.vim.org/scripts/script.php?script_id=5572)
- [yankee : minimal plugin to make yanking more fun](http://www.vim.org/scripts/script.php?script_id=5573)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (1193)
2. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (1012)
3. [wombat256.vim : Wombat for 256 color xterms](http://www.vim.org/scripts/script.php?script_id=2465) (520)
4. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (497)
5. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (423)
6. [AutomaticLaTeXPlugin : Background compilation, completion, bib serch, toc and other nice features.](http://www.vim.org/scripts/script.php?script_id=2945) (422)
7. [c.vim : C/C++ IDE -- Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (408)
8. [nginx.vim : initial version](http://www.vim.org/scripts/script.php?script_id=1886) (398)
9. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (389)
10. [pathogen.vim : Poor man's package manager. Easy manipulation of 'runtimepath' et al](http://www.vim.org/scripts/script.php?script_id=2332) (310)

## vim-jp issues

Open : 255 (+4) | Closed : 800 (+4)

- [Issue #1048 : 複数ファイル編集後、終了時にフリーズ](https://github.com/vim-jp/issues/issues/1048)
- [Issue #1049 : 整数の最大値に近い値をスライスの添字に指定すると間違った結果が返ってくる](https://github.com/vim-jp/issues/issues/1049)
- [Issue #1050 : execute('throw 1') が2度目から表示されない](https://github.com/vim-jp/issues/issues/1050)
- [Issue #1051 : &#x60;vim -c startinsert&#x60; で起動すると表示がおかしい](https://github.com/vim-jp/issues/issues/1051)
- [Issue #1052 : timer&#x5f;info()を連続でcallし続けるとtimerが止まらない](https://github.com/vim-jp/issues/issues/1052)
- [Issue #1053 : 262040byte以上の文字列をクリップボードにコピーすると正常に貼り付けできずSEGV](https://github.com/vim-jp/issues/issues/1053)
- [Issue #1054 : set hidden 時に diff 状態が window を閉じても解除されない](https://github.com/vim-jp/issues/issues/1054)
- [Issue #1055 : /&#x5f;CTRL-L で "~(チルダ)" があるとおかしくなる](https://github.com/vim-jp/issues/issues/1055)
