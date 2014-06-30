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

- [7.4.193 : typos in messages](http://code.google.com/p/vim/source/detail?r=a8650e2a0b5a5936f7d503429180df47df2aa775)
- [7.4.194 : can't build for Android](http://code.google.com/p/vim/source/detail?r=767103b5bcb05171fc1be0d6f480e0aa39aff731)
- [7.4.195 : (after 7.4.193) Python tests fail](http://code.google.com/p/vim/source/detail?r=c2e88a40d530b7999cd29218555515bcd282d661)
- [7.4.196 : tests fail on Solaris 9 and 10](http://code.google.com/p/vim/source/detail?r=1ed6acdbfef55e6e80aea55c8bc52d05f8458bbd)
- [7.4.197 : various problems on VMS](http://code.google.com/p/vim/source/detail?r=47a673b20e493db91844999f74aa3724808ac624)
- [7.4.198 : can't build with non-threding Perl and dynamic loading](http://code.google.com/p/vim/source/detail?r=705f398ce6f58016fae80c945a570721fc989117)
- [7.4.199 : (issue 197) \]P doesn't paste over Visual selection](http://code.google.com/p/vim/source/detail?r=54b1a90c937380195fad6a52408aa3b4eed6d8d1)
- [7.4.200 : too many #ifdefs in the code](http://code.google.com/p/vim/source/detail?r=0559091e6f09b5b96619513639ac14578a019612)
- [7.4.201 : 'lispwords' is a global option](http://code.google.com/p/vim/source/detail?r=06e5f65c34d8136c3a9d2219429b7eca35cb3a21)
- [7.4.202 : MS-Windows: non-ASCII font names don't work](http://code.google.com/p/vim/source/detail?r=22d7af9ff3e5e2b93fdbe8603df2f15155a5976b)
- [7.4.203 : parsing 'errorformat' is not correct](http://code.google.com/p/vim/source/detail?r=fb24b025c7cf07db79a559a3091db42e02c1af86)
- [7.4.204 : a mapping where the second byte is 0x80 doesn't work](http://code.google.com/p/vim/source/detail?r=f5120cbf16b9a9c6e0fbb599a6524e05ecf11393)
- [7.4.205 : ":mksession" does not handle the argument list correctly](http://code.google.com/p/vim/source/detail?r=0ace3a24c2a0153f0aaf9b619d3958e7f486705f)
- [7.4.206 : compiler warnings on 64 bit Windows](http://code.google.com/p/vim/source/detail?r=7e826028d3999b6125b4741d45b22a2d0a520679)
- [7.4.207 : cursor report sometimes not recognized causing replace mode](http://code.google.com/p/vim/source/detail?r=2aa909427e44cd3aac7def024b66e41d0c9d0e0d)
- [7.4.208 : Mercurial picks up some files that are not distributed](http://code.google.com/p/vim/source/detail?r=9b5541e276aa56fd72c3bbba85a17ef023c54ba8)
- [7.4.209 : when repeating a filter command "%" and "#" are expanded](http://code.google.com/p/vim/source/detail?r=bb402c49379de97fcd475fbbbbdc5ed41e5dff07)
- [7.4.210 : Visual block mode + virtual edit doesn't work well with tabs](http://code.google.com/p/vim/source/detail?r=420fd9cb86d51a92c4307a746557e81914c6d6c4)
- [7.4.211 : ":lu" is abbreviation for ":lua", but it should be ":lunmap"](http://code.google.com/p/vim/source/detail?r=e90bef2240c8d187da6e8d8fa5007ec5afc12284)
- [7.4.212 : #ifdefs for the +visual feature are not useful](http://code.google.com/p/vim/source/detail?r=50dbef5e774af6a998d1798569566429c38378b1)
- [7.4.213 : it's not possible to open a new buffer without a swapfile](http://code.google.com/p/vim/source/detail?r=e25a04c1c515e6eb32197291472f89bcadfabf89)
- [7.4.214 : compilation problems on HP&#x5f;nonStop (Tandem)](http://code.google.com/p/vim/source/detail?r=fe02fdfbdec077d0dc501bca0c65c51e56e1bd11)
- [7.4.215 : ":sp foo" reloads "foo" if it is the current buffer](http://code.google.com/p/vim/source/detail?r=f069a3a0f84451aa498c6c22d8f922d1e695e96d)
- [7.4.216 : compiler warnings](http://code.google.com/p/vim/source/detail?r=bf1775553d3bf561e289bbf367d014915d09eaad)
- [7.4.217 : "make clean" may run configure pointlessly](http://code.google.com/p/vim/source/detail?r=2f225a17c26b26eab87cfb89142fc0f95726ee92)
- [7.4.218 : it's not easy to remove duplicates from a list](http://code.google.com/p/vim/source/detail?r=ddc3f32a4b2191f829206322d46f0e9c7e365e22)
- [7.4.219 : redraw too often when 'relativenumber' or 'cursorline' set](http://code.google.com/p/vim/source/detail?r=37af1e6e91bb1e8ceb89d3ba1c49a04ffd889880)
- [7.4.220 : test 105 does not work in a shadow dir](http://code.google.com/p/vim/source/detail?r=1e272e318daaba51ec932be83a2481987a133429)
- [7.4.221 : quickfix doesn't resize on ":copen 20"](http://code.google.com/p/vim/source/detail?r=a548aae15b3a27a56d814900049785c29c01a37a)
- [7.4.222 : the Ruby directory is constructed from parts](http://code.google.com/p/vim/source/detail?r=9123b4d3ef05c7434b6f6cd011b63f63a9a4a539)
- [7.4.223 : still using an older autoconf version](http://code.google.com/p/vim/source/detail?r=4b78922d557c324de6b6a75515fb658235fa782c)
- [7.4.224 : /usr/bin/grep on Solaris does not support -F](http://code.google.com/p/vim/source/detail?r=092b8f61021b91024b127de3983d1afb74b64996)
- [7.4.225 : dynamic Ruby doesn't work on Solaris](http://code.google.com/p/vim/source/detail?r=eeb150c07647b3691d2696a3a4cdb96f6db5a63e)
- [7.4.226 : (after 7.4.219) cursorline highlight not always redrawn](http://code.google.com/p/vim/source/detail?r=b650f2db8f9604124c0ddfb14af0c04bd4ae0580)
- [7.4.227 : (after 7.4.225) can't build with Ruby 1.8](http://code.google.com/p/vim/source/detail?r=b0a9df477096e4abe938f998476e84e77db42e0b)
- [7.4.228 : compiler warnings when building with Python 3.2](http://code.google.com/p/vim/source/detail?r=3ee5808a293c7436e34a696a09ed29ce38750d9a)
- [7.4.229 : ":let" fails when listing vars with a curly braces expression](http://code.google.com/p/vim/source/detail?r=839cca5ec18d560e3714065e54ed38b6e812aaf7)

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

