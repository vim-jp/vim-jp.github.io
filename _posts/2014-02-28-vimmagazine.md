---
layout: post
category: vimmagazine
title: Vim Magazine 2 月号
---

# Vim Magazine 2 月号

## 話題

- [Neovim 登場](http://neovim.org/)

  Vim をフォークした新しいプロジェクト、 Neovim が登場しました。20 年以上に渡って開発されてきた Vim の古いコードベース、古い開発体制を改め、開発を加速させようという試み。 リファクタリングによってメンテナンス性を高め、新機能追加やプラグイン作成を容易にし、コミュニティによる開発を目標としています。発起人の Thiago de Arruda さんが <https://www.bountysource.com/fundraisers/539-neovim-first-iteration> でプロジェクトの開発資金を募集しています。

- [momonga.vim #3](http://partake.in/events/52a7b249-5183-4f1a-9e4f-22c1f0e436a5)

  若い女性に大人気と噂のVimについて便利します 3月21日(金) 東京

- [Vimプラグイン読書会 第4回](http://haya14busa.github.io/reading-vimplugin/)

  [vim-jp/vital.vim](https://github.com/vim-jp/vital.vim)のData.List, Data.String, Data.Dictionaryあたりを読みます。 3月15日(土) 21:00~ [LingrのVim部屋](lingr.com/room/vim)にて

## 今月の新機能

- 7.4.191: ファイル名修飾子`:S`が追加されました。

## リリース情報

- [7.4.161 : crash in Python exception handling](http://code.google.com/p/vim/source/detail?r=91f3908748c29b0e8e759ac4cdafc95e536eab9a)
- [7.4.162 : running tests in shadow dir doesn't work](http://code.google.com/p/vim/source/detail?r=477665b4105fa80e3655981790e55fc9babfeed5)
- [7.4.163 : (after 7.4.142) MS-Windows 7 input doesn't work properly](http://code.google.com/p/vim/source/detail?r=7db84aadd37aaf9d59f8f15dd6783a4e28dcb63c)
- [7.4.164 : problem with event handling on Windows 8](http://code.google.com/p/vim/source/detail?r=a01819fb6e2b5c270dac492ab2fe923ea9301651)
- [7.4.165 : by default, after closing a buffer changes can't be undone](http://code.google.com/p/vim/source/detail?r=c82339d8fdae0d43d51cde5216c979a5c9415bc4)
- [7.4.166 : auto-loading a function for code that won't be executed](http://code.google.com/p/vim/source/detail?r=5d03c374712128077ac4c342aad02120ed98df70)
- [7.4.167 : (after 7.4.149) fixes are not tested](http://code.google.com/p/vim/source/detail?r=22387c8eec43ea8b1b704cad49c8f7187e2fd579)
- [7.4.168 : can't compile with Ruby 2.1.0](http://code.google.com/p/vim/source/detail?r=e61a2b709f693eea9108d475b9bab5ed10ac675d)
- [7.4.169 : ":sleep" puts cursor in the wrong column](http://code.google.com/p/vim/source/detail?r=4e3a9dd25d428e7c08ed401afc244972e27e08e6)
- [7.4.170 : some help tags don't work with ":help"](http://code.google.com/p/vim/source/detail?r=8122eab8fcdbbdaac62dfbf7c6458cb3e6f46b04)
- [7.4.171 : redo does not set v:count and v:count1](http://code.google.com/p/vim/source/detail?r=beb037a6c2708f539d50840637f70eed0811d93c)
- [7.4.172 : blowfish code says OFB, but it's doing CFB](http://code.google.com/p/vim/source/detail?r=391e10afccf6879dcfab8b28cb1587a13eb835c0)
- [7.4.173 : when using scrollbind the cursor can end up below last line](http://code.google.com/p/vim/source/detail?r=233ad7b960d0fbeb224b383918113b25c74ebe35)
- [7.4.174 : compiler warnings for Python interface](http://code.google.com/p/vim/source/detail?r=f2c8d86c460d75fec17a0527a2fe93ac7220693e)
- [7.4.175 : wrong fall-back to non-wide function if wide function fails](http://code.google.com/p/vim/source/detail?r=6b69d8dde19e32909f4ee3a6337e6a2ecfbb6f72)
- [7.4.176 : Python: Dictionary.update() thows an unexpected error](http://code.google.com/p/vim/source/detail?r=ae228baaec2c5eda4cd948382c2bab498d1a34cc)
- [7.4.177 : compiler warning for unused variable](http://code.google.com/p/vim/source/detail?r=df3b0b70d8c1d29817597f45d19bb72755f86bd1)
- [7.4.178 : the J command does not update '\[ and '\] marks](http://code.google.com/p/vim/source/detail?r=647e6bb15aa3f864eaf447fe77e3e3ae7e37b134)
- [7.4.179 : Compiler warning for type-punned pointer](http://code.google.com/p/vim/source/detail?r=d1c8c1d6403485ff0135f0cf5c1ef07272a07173)
- [7.4.180 : (after 7.4.174) older Python versions don't support %ld](http://code.google.com/p/vim/source/detail?r=b1b8b097539a001043d1f0aac4663d6ee358dee8)
- [7.4.181 : when using 'pastetoggle' the status lines are not updated](http://code.google.com/p/vim/source/detail?r=cb5683bcde03796baa7e845fd9a2fcaec3383538)
- [7.4.182 : building with mzscheme and racket does not work](http://code.google.com/p/vim/source/detail?r=708a6a5905217736a2a231e36ced82f9535c4954)
- [7.4.183 : MSVC Visual Studio update not supported](http://code.google.com/p/vim/source/detail?r=1e2bfe4f3e903110f27cb6231f6642e721808837)
- [7.4.184 : match() does not work properly with a {count} argument](http://code.google.com/p/vim/source/detail?r=9ac2fc63501d3eff92446c03b2822b30b169db5a)
- [7.4.185 : Clang gives warnings](http://code.google.com/p/vim/source/detail?r=b06e58f8dd3037f7f17145968d08432423ac8665)
- [7.4.186 : (after 7.4.085) Visual mode insert mode is sometimes wrong](http://code.google.com/p/vim/source/detail?r=4d12112c5efae071aecbeed1a7196f18950457b3)
- [7.4.187 : delete that crosses line break splits multi-byte character](http://code.google.com/p/vim/source/detail?r=a1c07956171a133583df42627d3498f935e59988)
- [7.4.188 : SIZEOF&#x5f;LONG clashes with similar defines in header files](http://code.google.com/p/vim/source/detail?r=136f05449f29388e9023f7934d6344bf4ca61d0f)
- [7.4.189 : compiler warning for unused argument](http://code.google.com/p/vim/source/detail?r=0a9990bbd94a7dbdf20ca3f897cdf364c356d228)
- [7.4.190 : compiler warning for using %lld for off&#x5f;t](http://code.google.com/p/vim/source/detail?r=6050f52d7cd8d6f79ed5ee8ef3efff332aca4932)
- [7.4.191 : can't escape a file name for shell commands without a function](http://code.google.com/p/vim/source/detail?r=40f18a1c1592c8b4047f6f2a413557f48a99c55f)
- [7.4.192 : memory leak when giving E853](http://code.google.com/p/vim/source/detail?r=04c4ef8c0a1b757494500e46400552b135135e94)

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

