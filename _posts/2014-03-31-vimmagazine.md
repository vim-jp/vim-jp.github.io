---
layout: post
category: vimmagazine
title: Vim Magazine 3 月号
---

## 話題

- [Yokohama.vim #5](http://atnd.org/events/48511)

  約1年ぶりの開催です。 4月5日(土) 神奈川

- [TokyoVim#18](http://connpass.com/event/5658/)

  Vimに関する何かをやる会です。 4月12日(土) 東京

- [Vimプラグイン読書会 第5回](http://haya14busa.github.io/reading-vimplugin/)

  Vim scriptでリスト、文字列、辞書の効率的かつ汎用的な扱い方を学ぶために、[vim-jp/vital.vim](https://github.com/vim-jp/vital.vim)のData.List, Data.String, Data.Dictionaryあたりを読みます。
  4月12日(土) 21:00~ [LingrのVim部屋](http://lingr.com/room/vim)にて

- [第2回 Sapporo.vim](http://atnd.org/events/49000)

  Sapporo.vimは札幌でVimに興味がある！Vimが好きだ！ という方が対象の勉強会です。 皆さんでVimVimしましょう。
  5月25日(日) 札幌

## 今月の新機能

- 7.4.200 : ビジュアルモードの機能が最小構成ビルドでも有効化されました。
- 7.4.201 : オプション `'lispwords'` が global-local になりました。
- 7.4.202 : Windowsで `enc=utf-8` のとき、フォント名にマルチバイト文字が使えるようになりました。
- 7.4.204 : マルチバイト文字に map 出来るようになりました。
- 7.4.213 : `:noswapfile` が追加されました。
- 7.4.218 : `uniq()` 関数が追加されました。

## リリース情報

- [7.4.193 : typos in messages](https://github.com/vim/vim/commit/4de6a212f922aa132e9a76d0392bee3adfee6e29)
- [7.4.194 : can't build for Android](https://github.com/vim/vim/commit/4064e24a0f7581b47dcc9de9e24891d295d4f6e2)
- [7.4.195 : (after 7.4.193) Python tests fail](https://github.com/vim/vim/commit/038e5d47660b1b6ffe9005a4bcc51c88a3b38416)
- [7.4.196 : tests fail on Solaris 9 and 10](https://github.com/vim/vim/commit/f7dc2b551708315a833fcb19c288759c277dcfe1)
- [7.4.197 : various problems on VMS](https://github.com/vim/vim/commit/206f011829920e41c4d3782af7591da6e1b86655)
- [7.4.198 : can't build with non-threding Perl and dynamic loading](https://github.com/vim/vim/commit/d8619997c47557a5f8acd52fcab04cd3eb01a0d6)
- [7.4.199 : (issue 197) \]P doesn't paste over Visual selection](https://github.com/vim/vim/commit/27bed20452ad2f8a9c9d749d2db20660c3d0e45a)
- [7.4.200 : too many #ifdefs in the code](https://github.com/vim/vim/commit/24ff9e33a9b87f63547e7bce95d09ade1a04fcfb)
- [7.4.201 : 'lispwords' is a global option](https://github.com/vim/vim/commit/af6c131bf7f86dc85fbc2e4a79f2547786228126)
- [7.4.202 : MS-Windows: non-ASCII font names don't work](https://github.com/vim/vim/commit/b1692e2b8f074926f75b11e0d0a519d502b95c82)
- [7.4.203 : parsing 'errorformat' is not correct](https://github.com/vim/vim/commit/8eded09bf7a1c4f7994db7ef66780a32334c2538)
- [7.4.204 : a mapping where the second byte is 0x80 doesn't work](https://github.com/vim/vim/commit/1d9ff43f585da9dc51c65a8c64bc19b281cc73f6)
- [7.4.205 : ":mksession" does not handle the argument list correctly](https://github.com/vim/vim/commit/f0bdd2f28db28707385535ac9c270903e16c4637)
- [7.4.206 : compiler warnings on 64 bit Windows](https://github.com/vim/vim/commit/a0844a117068ff4d541eb17ef3c0566995e5d3c8)
- [7.4.207 : cursor report sometimes not recognized causing replace mode](https://github.com/vim/vim/commit/9c8c8c5d302ad40a3cd89fcd76f718b7e9fe86d5)
- [7.4.208 : Mercurial picks up some files that are not distributed](https://github.com/vim/vim/commit/2fb9fa19e2263f077163d0da844b637d0a356863)
- [7.4.209 : when repeating a filter command "%" and "#" are expanded](https://github.com/vim/vim/commit/529d2d63699bd43fde8c04fd0c84fd7b3e6bf20a)
- [7.4.210 : Visual block mode + virtual edit doesn't work well with tabs](https://github.com/vim/vim/commit/4c9a949d00468dde748653b269338ddfb30910cc)
- [7.4.211 : ":lu" is abbreviation for ":lua", but it should be ":lunmap"](https://github.com/vim/vim/commit/ed287f9a4e3f4ed5528ad2af65b7b23bce14a688)
- [7.4.212 : #ifdefs for the +visual feature are not useful](https://github.com/vim/vim/commit/f7ff6e85e8a4e84cff023f5db4b66ef85986605a)
- [7.4.213 : it's not possible to open a new buffer without a swapfile](https://github.com/vim/vim/commit/5803ae6c076b1d61381afe27fcdedac61dd2cda9)
- [7.4.214 : compilation problems on HP&#x5f;nonStop (Tandem)](https://github.com/vim/vim/commit/03305f396ff7be377cb608f1a281c1eb5aa80d0b)
- [7.4.215 : ":sp foo" reloads "foo" if it is the current buffer](https://github.com/vim/vim/commit/7b44934037cf2a603df8c9dcd6f82cdde4ec4940)
- [7.4.216 : compiler warnings](https://github.com/vim/vim/commit/c410530637aad9e7ab187e1fb0ffc6c3de2b5b20)
- [7.4.217 : "make clean" may run configure pointlessly](https://github.com/vim/vim/commit/1a3eb8e1b51015b99d2e68242dc5326153edc72f)
- [7.4.218 : it's not easy to remove duplicates from a list](https://github.com/vim/vim/commit/327aa02ddaafa7c4585120ed152e078b9d74acea)
- [7.4.219 : redraw too often when 'relativenumber' or 'cursorline' set](https://github.com/vim/vim/commit/b679875b29b519b5a612ffa6f4379ddfa5da5b12)
- [7.4.220 : test 105 does not work in a shadow dir](https://github.com/vim/vim/commit/c666b5b467df3a3c5c7755bf1cb45724879709c9)
- [7.4.221 : quickfix doesn't resize on ":copen 20"](https://github.com/vim/vim/commit/158864120d3ee8baa2b5c9c2cb15bd5c5ea5731b)
- [7.4.222 : the Ruby directory is constructed from parts](https://github.com/vim/vim/commit/a6fd37be4fa01cef4bc4368f8fa814d9c6d084fd)
- [7.4.223 : still using an older autoconf version](https://github.com/vim/vim/commit/7db7784103306cfa7db4e94acf9b1bc2b0bc3888)
- [7.4.224 : /usr/bin/grep on Solaris does not support -F](https://github.com/vim/vim/commit/2bcaec320aff11e777618e5b8c112aa15ec6ca21)
- [7.4.225 : dynamic Ruby doesn't work on Solaris](https://github.com/vim/vim/commit/73b044dca94c8d427144b920c12fe5c47007a6f6)
- [7.4.226 : (after 7.4.219) cursorline highlight not always redrawn](https://github.com/vim/vim/commit/3d6db1467b9c3737fcf61e03209f2ffcea5a4bca)
- [7.4.227 : (after 7.4.225) can't build with Ruby 1.8](https://github.com/vim/vim/commit/498af70e066c66b66f1737b553287c1339846842)
- [7.4.228 : compiler warnings when building with Python 3.2](https://github.com/vim/vim/commit/922a4664fe51662a24097b8e74e5f716beef12f4)
- [7.4.229 : ":let" fails when listing vars with a curly braces expression](https://github.com/vim/vim/commit/a392038db5871af6f78fe4d822d9f7126f25eab6)

## 新着スクリプト

- [TWCommand : Tab / Window move & management commands for vim.](http://www.vim.org/scripts/script.php?script_id=4876)
- [magnum.vim : Pure Vim script big integer library](http://www.vim.org/scripts/script.php?script_id=4877)
- [Abridge : Vim code snippets creation made easy](http://www.vim.org/scripts/script.php?script_id=4878)
- [vim-autoflake : vim-autoflake is a Vim plugin that applies autoflake to your current file. ](http://www.vim.org/scripts/script.php?script_id=4879)
- [columnmove : Bring cursor vertically in similar ways as line-wise commands.](http://www.vim.org/scripts/script.php?script_id=4880)
- [patternjump : Move cursor as you like.](http://www.vim.org/scripts/script.php?script_id=4881)
- [wipeout : Destroy all buffers that are not open in any tabs or windows.](http://www.vim.org/scripts/script.php?script_id=4882)
- [pipe2eval : simple REPL inside vim](http://www.vim.org/scripts/script.php?script_id=4883)
- [ctrlp-py-matcher : Fast matcher for CtrlP plugin written in python.](http://www.vim.org/scripts/script.php?script_id=4884)
- [FilePathConvert : Convert filespec between absolute, relative, and URL formats.](http://www.vim.org/scripts/script.php?script_id=4885)
- [ConsultADict.vim : ConsultADict is plugin that provides easy way to consult a dictionary](http://www.vim.org/scripts/script.php?script_id=4886)
- [REPL plugin : A plugin for handling a REPL interpreter in a scratch window](http://www.vim.org/scripts/script.php?script_id=4887)
- [svnj.vim : vim svn plugin (subversion svn )](http://www.vim.org/scripts/script.php?script_id=4888)
- [likelycomplete : Collect & rank words for completion](http://www.vim.org/scripts/script.php?script_id=4889)
- [quickfix-reflector.vim : Change code right in the quickfix window](http://www.vim.org/scripts/script.php?script_id=4890)
- [nomad : Keep your environment variables in sync with tmux](http://www.vim.org/scripts/script.php?script_id=4891)
- [vim-less : Less syntax highlighting](http://www.vim.org/scripts/script.php?script_id=4892)
- [vim-bookmarks : Toggle line-based bookmarks](http://www.vim.org/scripts/script.php?script_id=4893)
- [Markology : Mark Visualization, Navigation and Management](http://www.vim.org/scripts/script.php?script_id=4894)
- [Pterosaur : Makes firefox text-inputs use vim.](http://www.vim.org/scripts/script.php?script_id=4895)
- [clone : Create a duplicate clone of the current buffer.](http://www.vim.org/scripts/script.php?script_id=4896)
- [cloneSimilar : Create a duplicate clone of the current buffer with a similar name.](http://www.vim.org/scripts/script.php?script_id=4897)
- [VIDE : Simple IDE interface for vim](http://www.vim.org/scripts/script.php?script_id=4898)
- [cra.vim : yearly timesheet](http://www.vim.org/scripts/script.php?script_id=4899)
- [Arrow Key Repurpose : A plugin for repurposing the arrow keys mostly for shifting lines of text around](http://www.vim.org/scripts/script.php?script_id=4900)
- [vim-lsdyna : Initial upload](http://www.vim.org/scripts/script.php?script_id=4901)
- [VimCompletesMe : Simple tab completion for Vim](http://www.vim.org/scripts/script.php?script_id=4902)
- [github-issues.vim : Github Issues integration with Vim](http://www.vim.org/scripts/script.php?script_id=4903)
- [avahi-syntax : VIM syntax highlighting for Avahi](http://www.vim.org/scripts/script.php?script_id=4904)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (2795)
2. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (2536)
3. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (1483)
4. [pathogen.vim : Easy manipulation of 'runtimepath', 'path', 'tags', etc](http://www.vim.org/scripts/script.php?script_id=2332) (1388)
5. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (1164)
6. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (1012)
7. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (985)
8. [AutomaticLaTeXPlugin : Background compilation, completion, bib serch, toc and other nice features.](http://www.vim.org/scripts/script.php?script_id=2945) (938)
9. [bufexplorer.zip : Buffer Explorer / Browser](http://www.vim.org/scripts/script.php?script_id=42) (872)
10. [snipMate : TextMate-style snippets for Vim](http://www.vim.org/scripts/script.php?script_id=2540) (868)

## vim-jp issues

Open : 183 (+1) | Closed : 359 (+10)

- [Issue #532 : vimのlatexの書き方　、や。の扱い](https://github.com/vim-jp/issues/issues/532)
- [Issue #533 : vim-surroundについて](https://github.com/vim-jp/issues/issues/533)
- [Issue #534 : Windows の gf/gF で grep 結果が開けない](https://github.com/vim-jp/issues/issues/534)
- [Issue #535 : BufAdd で :e すると SEGV](https://github.com/vim-jp/issues/issues/535)
- [Issue #536 : Test 86 and 87 failed with 7.4.193](https://github.com/vim-jp/issues/issues/536)
- [Issue #537 : vim-over](https://github.com/vim-jp/issues/issues/537)
- [Issue #538 : .cファイルで#ifdef AA ～ #endifを#if 1 ～ #endifで囲むとAAのSyntax highlightが変化する](https://github.com/vim-jp/issues/issues/538)
- [Issue #539 : if&#x5f;python3 コンパイル時の警告について](https://github.com/vim-jp/issues/issues/539)
- [Issue #540 : jsx のファイルを開くと ft=javascript になる](https://github.com/vim-jp/issues/issues/540)
- [Issue #541 : help に &#x2a;String&#x2a; タグがない](https://github.com/vim-jp/issues/issues/541)
- [Issue #542 : cwについて](https://github.com/vim-jp/issues/issues/542)

