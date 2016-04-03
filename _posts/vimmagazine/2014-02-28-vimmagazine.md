---
layout: post
category: vimmagazine
title: Vim Magazine 2 月号
---

## 話題

- [Neovim 登場](http://neovim.org/)

  Vim をフォークした新しいプロジェクト、 Neovim が登場しました。20 年以上にわたって開発されてきた Vim の古いコードベース、古い開発体制を改め、開発を加速させようという試み。 リファクタリングによってメンテナンス性を高め、新機能追加やプラグイン作成を容易にし、コミュニティによる開発を目標としています。発起人の Thiago de Arruda さんが <https://www.bountysource.com/fundraisers/539-neovim-first-iteration> でプロジェクトの開発資金を募集しています。

- [momonga.vim #3](http://partake.in/events/52a7b249-5183-4f1a-9e4f-22c1f0e436a5)

  若い女性に大人気と噂のVimについて便利します 3月21日(金) 東京

- [Vimプラグイン読書会 第4回](http://haya14busa.github.io/reading-vimplugin/)

  プラグインマネージャの仕組みを理解するために、[tpope/vim-pathogen](https://github.com/tpope/vim-pathogen)を読みます。
  3月15日(土) 21:00~ [LingrのVim部屋](http://lingr.com/room/vim)にて

## 今月の新機能

- 7.4.191: ファイル名修飾子`:S`が追加されました。

## リリース情報

- [7.4.161 : crash in Python exception handling](https://github.com/vim/vim/commit/4315f26500934bff48ed9bda85d2c9db0c8c73f7)
- [7.4.162 : running tests in shadow dir doesn't work](https://github.com/vim/vim/commit/09edc46d01ee0804e29594738c5b86a8a2c06ca0)
- [7.4.163 : (after 7.4.142) MS-Windows 7 input doesn't work properly](https://github.com/vim/vim/commit/f50eb7882e89a6ec1670a0377224f174c7899338)
- [7.4.164 : problem with event handling on Windows 8](https://github.com/vim/vim/commit/dd415a6c5acf5c7c054a16331b34f6b84f9c004d)
- [7.4.165 : by default, after closing a buffer changes can't be undone](https://github.com/vim/vim/commit/767568846b322a52cfbf21f048b1e9e968e393c5)
- [7.4.166 : auto-loading a function for code that won't be executed](https://github.com/vim/vim/commit/f31ecce97eb2aa77a9f43725710af0619e409035)
- [7.4.167 : (after 7.4.149) fixes are not tested](https://github.com/vim/vim/commit/4f943c09a5f09b330d92177ab50b5dd2a7986836)
- [7.4.168 : can't compile with Ruby 2.1.0](https://github.com/vim/vim/commit/a1a118b159ae9555a7fdce3d707ddc3064717067)
- [7.4.169 : ":sleep" puts cursor in the wrong column](https://github.com/vim/vim/commit/10395d8d6c9f85d5af42d30fb5bd2225bd8ef626)
- [7.4.170 : some help tags don't work with ":help"](https://github.com/vim/vim/commit/c467d9b2745642df836e89c026b2e84d9b227b59)
- [7.4.171 : redo does not set v:count and v:count1](https://github.com/vim/vim/commit/0a36fece9234d423aa5e35a8df1e33b3f79e9ede)
- [7.4.172 : blowfish code says OFB, but it's doing CFB](https://github.com/vim/vim/commit/4d504a3e11c42ca1f7ad29848b943eb1537283cb)
- [7.4.173 : when using scrollbind the cursor can end up below last line](https://github.com/vim/vim/commit/12a0f22340ef1b0fa9a4e96d918e282c3d9c2c4e)
- [7.4.174 : compiler warnings for Python interface](https://github.com/vim/vim/commit/3b5226121ddefaa51c1dbb844bfb7124386fe4a8)
- [7.4.175 : wrong fall-back to non-wide function if wide function fails](https://github.com/vim/vim/commit/cd981f2e0f00613a63b46e1e6b5227d5993ba994)
- [7.4.176 : Python: Dictionary.update() thows an unexpected error](https://github.com/vim/vim/commit/2d5f38ff10a955058416b93aae774aeef1c34486)
- [7.4.177 : compiler warning for unused variable](https://github.com/vim/vim/commit/b8e2305656adeb9357e726742bfe0e91e88cda28)
- [7.4.178 : the J command does not update '\[ and '\] marks](https://github.com/vim/vim/commit/f92d8a254ac092e525049ab930924d34eca0f114)
- [7.4.179 : Compiler warning for type-punned pointer](https://github.com/vim/vim/commit/0e4eebd4254017970b7389f66315ebed1486e1ed)
- [7.4.180 : (after 7.4.174) older Python versions don't support %ld](https://github.com/vim/vim/commit/403b3cf0a1ee4c7c1862ef52d42713b90c75a5ee)
- [7.4.181 : when using 'pastetoggle' the status lines are not updated](https://github.com/vim/vim/commit/06811f33ce34575e32f890f8b70c675d44858821)
- [7.4.182 : building with mzscheme and racket does not work](https://github.com/vim/vim/commit/45e2bcc2750dc09121bab718da454ed0b0b77c26)
- [7.4.183 : MSVC Visual Studio update not supported](https://github.com/vim/vim/commit/c86438205ba1d7b5ecadaa700c7f9943fdaf9838)
- [7.4.184 : match() does not work properly with a {count} argument](https://github.com/vim/vim/commit/9feaf6206f73b367f3167ffff927c1a5d2d90954)
- [7.4.185 : Clang gives warnings](https://github.com/vim/vim/commit/631abc35dfc7ad87255c940fd22eee416c731658)
- [7.4.186 : (after 7.4.085) Visual mode insert mode is sometimes wrong](https://github.com/vim/vim/commit/b1d90a3d770c5e513e6cad8284a5168ac5535883)
- [7.4.187 : delete that crosses line break splits multi-byte character](https://github.com/vim/vim/commit/7d311c52ce2859f67532e0a4b604eb8147473c01)
- [7.4.188 : SIZEOF&#x5f;LONG clashes with similar defines in header files](https://github.com/vim/vim/commit/a2aa31a9269641839178b0c0663dc67405c33d3e)
- [7.4.189 : compiler warning for unused argument](https://github.com/vim/vim/commit/f30caaf08d2ccbf429cf1eeb910b0ab08b88a40d)
- [7.4.190 : compiler warning for using %lld for off&#x5f;t](https://github.com/vim/vim/commit/581966e8323c2bab6f9e54729708dc46de8f9fc5)
- [7.4.191 : can't escape a file name for shell commands without a function](https://github.com/vim/vim/commit/26df092843de91ea0c5c5c130d0d0695d2d81c07)
- [7.4.192 : memory leak when giving E853](https://github.com/vim/vim/commit/47b8342abd3dfee39f5961dbb15883c3a95d8488)

## 新着スクリプト

- [Enchanted Vim : persistent very magic option](http://www.vim.org/scripts/script.php?script_id=4849)
- [vim-colors-pencil : A light (and dark) color scheme inspired by iA Writer](http://www.vim.org/scripts/script.php?script_id=4850)
- [grin.vim :   Vim plugin for 'grin', the Python port of 'ack'](http://www.vim.org/scripts/script.php?script_id=4851)
- [vawa.vim : A plugin that will highlight all occurances of the php variable under the cursor](http://www.vim.org/scripts/script.php?script_id=4852)
- [scd.vim : smart change of directory in Vim](http://www.vim.org/scripts/script.php?script_id=4853)
- [CRDispatcher : dispatch CR in the command line](http://www.vim.org/scripts/script.php?script_id=4856)
- [APL character picker : A pop-up menu for inserting APL characters](http://www.vim.org/scripts/script.php?script_id=4857)
- [stdtabs : Indentation for some languages according to their most used style guidelines.](http://www.vim.org/scripts/script.php?script_id=4858)
- [better-whitespace : A better whitespace highlighter for Vim](http://www.vim.org/scripts/script.php?script_id=4859)
- [VimMahewinHub : Interface for hub command.](http://www.vim.org/scripts/script.php?script_id=4860)
- [sql.snippets : T-SQL snippets](http://www.vim.org/scripts/script.php?script_id=4861)
- [hybris ImpEx : Providing syntax highlight for hybris ImpEx file](http://www.vim.org/scripts/script.php?script_id=4862)
- [jdaddy.vim : JSON manipulation and pretty printing](http://www.vim.org/scripts/script.php?script_id=4863)
- [c.vim : Highlight more syntax items of C program](http://www.vim.org/scripts/script.php?script_id=4864)
- [writable&#x5f;search.vim : Perform a grep and edit the results directly](http://www.vim.org/scripts/script.php?script_id=4865)
- [prop.vim : Android system properties highlight](http://www.vim.org/scripts/script.php?script_id=4866)
- [fuckpep8 : Make life feasible in the presense of crazy space-hippies.](http://www.vim.org/scripts/script.php?script_id=4867)
- [wildfire.vim : Smart selection of the closest text object](http://www.vim.org/scripts/script.php?script_id=4868)
- [searchmatch : Easy search highlight pinning in Vim with :match](http://www.vim.org/scripts/script.php?script_id=4869)
- [ExtractLinks : Replace inline links with unique references and a link table.](http://www.vim.org/scripts/script.php?script_id=4870)
- [Licenses : Provides commands to add licenses at the top of the buffer](http://www.vim.org/scripts/script.php?script_id=4871)
- [vim-decfile : A plugin for EvtGen decay files](http://www.vim.org/scripts/script.php?script_id=4872)
- [gnuplot syntax highlighting : syntax highlighting for gnuplot](http://www.vim.org/scripts/script.php?script_id=4873)
- [LeaderF : A replacement of the famous ctrlp.vim yet written in Python.](http://www.vim.org/scripts/script.php?script_id=4874)
- [industry.vim : Standard industrial color scheme for Vim](http://www.vim.org/scripts/script.php?script_id=4875)

## 月間ダウンロードランキング

1. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (1939)
2. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (1929)
3. [pathogen.vim : Easy manipulation of 'runtimepath', 'path', 'tags', etc](http://www.vim.org/scripts/script.php?script_id=2332) (1207)
4. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (1173)
5. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (905)
6. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (832)
7. [github theme : A gvim colorscheme based on github's syntax highlighting](http://www.vim.org/scripts/script.php?script_id=2855) (735)
8. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (642)
9. [snipMate : TextMate-style snippets for Vim](http://www.vim.org/scripts/script.php?script_id=2540) (637)
10. [minibufexpl.vim : Elegant buffer explorer - takes very little screen space](http://www.vim.org/scripts/script.php?script_id=159) (601)

## vim-jp issues

Open : 182 (-3) | Closed : 349 (+12)

- [Issue #523 : vim-latexについて](https://github.com/vim-jp/issues/issues/523)
- [Issue #524 : 、や。のマッピング](https://github.com/vim-jp/issues/issues/524)
- [Issue #525 : guicursorの内部の情報を正しく得ることができない](https://github.com/vim-jp/issues/issues/525)
- [Issue #526 : match('foo', '$', 0, count)の count が 2 以上の時の挙動がおかしい](https://github.com/vim-jp/issues/issues/526)
- [Issue #527 : vim 突然落ちる現象](https://github.com/vim-jp/issues/issues/527)
- [Issue #528 : 改行コードCR/LFが条件によって変換されてしまう(?)](https://github.com/vim-jp/issues/issues/528)
- [Issue #529 : vim-latexの表示について](https://github.com/vim-jp/issues/issues/529)
- [Issue #530 : --enable-perlinterp=dynamic かつ非 USE&#x5f;ITHREAD で if&#x5f;perl がビルド出来ない(was: Perl 5.19 で if&#x5f;perl がビルドできない)](https://github.com/vim-jp/issues/issues/530)
- [Issue #531 : perl5.19.9 の signatures でプロトタイプのsyntaxがエラー色になる](https://github.com/vim-jp/issues/issues/531)

