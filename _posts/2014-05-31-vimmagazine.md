---
layout: post
category: vimmagazine
title: Vim Magazine 5 月号
---

## 話題

- [momonga.vim #4](http://connpass.com/event/6108/)

  Vimをします（かなり）。
  Vim pluginを作ったり、.vimrcを弄ったり、Vimで何かつくったり。自由にもくもくします。初心者歓迎ですよっ
  6月7日(土) 東京

- [TokyoVim#20](http://tokyovim.connpass.com/event/6478/)

  Vimに関する何かをやる会です。
  6月14日(土) 東京

- [Vimプラグイン読書会 第7回](http://haya14busa.github.io/reading-vimplugin/)

  [terryma/vim-expand-region](https://github.com/terryma/vim-expand-region)を読みます。
  6月21日(土) 21:00~ [LingrのVim部屋](http://lingr.com/room/vim)にて

- [Aizu.vim #2](http://atnd.org/events/51752)

  Aizuの夏、Vimの夏
  6月22日(日) 福島

- [Osaka.vim #1](http://osaka-vim.connpass.com/event/6491/)

  Vimと仲良くなりたいでしょ？じゃあ参加してなんかしよっ。
  6月28日(土) 大阪

## 今月の新機能

- 7.4.276 : fish shellがサポートされました。
- 7.4.279 : `globpath()` に `glob()` と同様な任意の引数が追加されました。
- 7.4.293 : `%C` 指定子が追加されました。合成文字を無視することができます。
- 7.4.312 : `arglistid()` 関数が追加されました。ウィンドウで使用されている引数リストのIDを得ることができます。
- 7.4.313 : `getcurpos()` 関数が追加されました。
- 7.4.314 : `'shortmess'` オプションに `'c'` フラグが追加されました。補完ポップアップメニュー関連のメッセージを表示しないようにできます。

## リリース情報

- [7.4.274 : check for timestamp does not always detect a changed file](http://code.google.com/p/vim/source/detail?r=1ee3fc5b40ae94c2a7fc5a62bca38d4f730f9bb2)
- [7.4.275 : no error message when changing type of unplaced sign](http://code.google.com/p/vim/source/detail?r=8a3117a4887c1e12a1165c9719491f96753787d6)
- [7.4.276 : the fish shell is not supported](http://code.google.com/p/vim/source/detail?r=a6b59ee633a355095e6473ec5e2a7d9088bfb853)
- [7.4.277 : using ":sign unplace &#x2a;" may leave the cursor in wrong position](http://code.google.com/p/vim/source/detail?r=373204662d82e894b27ee76bc3319bc62c91f6ae)
- [7.4.278 : list&#x5f;remove() conflicts with Sun header file](http://code.google.com/p/vim/source/detail?r=b4ce0e1fb5a67d7d6b0bca8eaa3edc2e94a085d8)
- [7.4.279 : globpath() returns a string, a list would be better](http://code.google.com/p/vim/source/detail?r=8e9db1f27a0063df023cc05a760fce73255dad24)
- [7.4.280 : relative position of cursor not restored if session has 2 tabs](http://code.google.com/p/vim/source/detail?r=daf7e98675cf395e1ef96f8040567affb2782a11)
- [7.4.281 : positions are off when loading session with two tabpages](http://code.google.com/p/vim/source/detail?r=24c90f1fec859b54cf2b854b98c4c9e614c46061)
- [7.4.282 : (after 7.4.279) test 97 fails on Mac](http://code.google.com/p/vim/source/detail?r=6d0a1132dd71c7f55f7ed53fe99e97c79bfd05a4)
- [7.4.283 : (after 7.4.276) compiler warning about unused variable](http://code.google.com/p/vim/source/detail?r=aa99d04fa7e288a8580e3a5d4a9d6433a1572b48)
- [7.4.284 : setting 'langmap' in the modeline can cause trouble](http://code.google.com/p/vim/source/detail?r=3c35ca9666e88a8024af6dab585b8e79ab295f83)
- [7.4.285 : line numbers are not always updated for 'relativenumber'](http://code.google.com/p/vim/source/detail?r=5cb1828fd0056de3c166e71fbafc67a74c57d7b1)
- [7.4.286 : error messages are inconsistant](http://code.google.com/p/vim/source/detail?r=be19015ef43cc17825929206790696c2e716035d)
- [7.4.287 : patches for .hgignore don't work](http://code.google.com/p/vim/source/detail?r=66fe4908b649ba18426af6f69e8ccb01b487dcbd)
- [7.4.288 : when 'spellfile' is set the screen is not redrawn](http://code.google.com/p/vim/source/detail?r=7965cb6a435ae1ea331c7c2f8740d3d4c3625f3b)
- [7.4.289 : NFA regexp with repeated backreference does not match](http://code.google.com/p/vim/source/detail?r=99374096a76b96d1128f5e6aa1fa92b4ba70fee9)
- [7.4.290 : a non-greedy match followed by a branch is too greedy](http://code.google.com/p/vim/source/detail?r=b871734bf54ea185dbd2cc759d86dbfbe21cde26)
- [7.4.291 : compiler warning for int to pointer of different size](http://code.google.com/p/vim/source/detail?r=b5972833add9de714f4651e26fd9ea63ec4a880c)
- [7.4.292 : pattern "a" matches accented "a" when pattern is one byte](http://code.google.com/p/vim/source/detail?r=60cdaa05a6ad31cef55eb6b3dc1f57ecac6fcf79)
- [7.4.293 : can't ignore composing characters in a pattern](http://code.google.com/p/vim/source/detail?r=10fc95f48546f438648b8357062e93c9c2c0a377)
- [7.4.294 : (after 7.4.293) test files missing from patch](http://code.google.com/p/vim/source/detail?r=fdea5ea9afd139ea59dee6bdb3f1675b8b882bdf)
- [7.4.295 : various typos, bad white space and unclear comments](http://code.google.com/p/vim/source/detail?r=662ae48e7e246a63d38c9f3165b15b62252edaee)
- [7.4.296 : can't run tests on Solaris](http://code.google.com/p/vim/source/detail?r=53b87d790574b6d19034fb3390987c22fb928c58)
- [7.4.297 : memory leak from result of get&#x5f;isolated&#x5f;shell&#x5f;name()](http://code.google.com/p/vim/source/detail?r=81f5a056b2a582c8109da10cc538dc16a326a34d)
- [7.4.298 : can't have a funcref start with "t:"](http://code.google.com/p/vim/source/detail?r=156f891d520e93eab5d3ce02784660fb13a3b0d3)
- [7.4.299 : when running configure twice DYNAMIC&#x5f;PYTHON&#x5f;DLL becomes empty](http://code.google.com/p/vim/source/detail?r=daebf8ce66089c0c179fb436ceba359ef8d593d5)
- [7.4.300 : the way config.cache is removed doesn't always work](http://code.google.com/p/vim/source/detail?r=1157079ca5f167bcf8746dfc52ea5a85e6c87a30)
- [7.4.301 : (after 7.4.280) scrolling problem when loading a session file](http://code.google.com/p/vim/source/detail?r=8cb42aa3c4957a543e5dffe307475dbab969612f)
- [7.4.302 : signs don't show up after filler lines](http://code.google.com/p/vim/source/detail?r=df141c80ea3a1ffcbf82d05c1314675231fcfa75)
- [7.4.303 : double-width chars on the command line sometimes truncated](http://code.google.com/p/vim/source/detail?r=463ef551e9f62b63ac3f85f1f297b668b14bcd09)
- [7.4.304 : cannot always use Python with Vim](http://code.google.com/p/vim/source/detail?r=fed2e0967f8133ba9a44b0701f151c8d88c4896a)
- [7.4.305 : making 'ttymouse' empty causes problems](http://code.google.com/p/vim/source/detail?r=63e7cc62402dffb180b40c04c63ceeb5f53957d7)
- [7.4.306 : getchar(0) does not return Esc](http://code.google.com/p/vim/source/detail?r=05e1d8afcc5e375bf708ccc9810e2fd1a5a8a3cf)
- [7.4.307 : can't build without the +termresponse feature](http://code.google.com/p/vim/source/detail?r=06c10522d321d98874546b2a4d3b0ae145386f2e)
- [7.4.308 : ":diffsplit" on empty file cursor is displayed on command line](http://code.google.com/p/vim/source/detail?r=e3d2b8d83bb30c428a051f50791e454fcbc080af)
- [7.4.309 : increasing size of lower window, upper window jumps to the top](http://code.google.com/p/vim/source/detail?r=88a6e9f33822d33b6c32db578750c6c178c63f50)
- [7.4.310 : getpos()/setpos() don't include curswant](http://code.google.com/p/vim/source/detail?r=ccac0aa34eeaf46dad4b831461a532fc3fe71096)
- [7.4.311 : can't use winrestview to only restore part of the view](http://code.google.com/p/vim/source/detail?r=f6f7543043246107075f0d3739c471d51b7226da)
- [7.4.312 : cannot know what argument list is being used for a window](http://code.google.com/p/vim/source/detail?r=66eead134d6800fd4cf2d5d4b135d300c933f09a)
- [7.4.313 : (after 7.4.310) changing return value of getpos() causes error](http://code.google.com/p/vim/source/detail?r=332a5c2b2956d9b18d85268a724d01deea27ec83)
- [7.4.314 : completion messages can get in the way of a plugin](http://code.google.com/p/vim/source/detail?r=4d7af1962d6ce61df65fdc5c86544a61951f9517)
- [7.4.315 : (after 7.4.309) fixes for computation of topline not tested](http://code.google.com/p/vim/source/detail?r=646616b6ff4defcc7bef0b198b540f6d965a8126)
- [7.4.316 : warning from 64-bit compiler](http://code.google.com/p/vim/source/detail?r=0fc665889e8f0af532cb4e3be2f0ff0421bf2fbc)

## 新着スクリプト

- [diffchar.vim : Highlight the differences, based on characters and words](http://www.vim.org/scripts/script.php?script_id=4932)
- [lite-dfm : A plugin to remove distractions](http://www.vim.org/scripts/script.php?script_id=4933)
- [Syntax for Clustal : Highlighting for Clustal Alignment](http://www.vim.org/scripts/script.php?script_id=4935)
- [mrtg.vim : Multi Router Traffic Grapher config file](http://www.vim.org/scripts/script.php?script_id=4936)
- [RubyJump : Vim plugin for ruby editing. Quick jump to method, class, module defenitions.](http://www.vim.org/scripts/script.php?script_id=4937)
- [cmake : CMake syntax update](http://www.vim.org/scripts/script.php?script_id=4938)
- [TWstack : Smart commands for moving between tab pages and windows with history stacks.](http://www.vim.org/scripts/script.php?script_id=4939)
- [BufLine : This script will show your buffer list in the tabline.](http://www.vim.org/scripts/script.php?script_id=4940)
- [Vebugger : frontend for interactive shell debuggers](http://www.vim.org/scripts/script.php?script_id=4941)
- [vim-do-syntax : Syntax highlighting for Mentor Graphics Questa Sim do files](http://www.vim.org/scripts/script.php?script_id=4942)
- [vim-ralf-syntax : Syntax highlighting for Synopsys Register Abstraction Layer ralf files](http://www.vim.org/scripts/script.php?script_id=4943)
- [Switch c++ header files : Switch current vim buffer between c++ header and cpp files](http://www.vim.org/scripts/script.php?script_id=4944)
- [SemanticHL : Semantic highlighting for C/C++ (v1.0 20140521)](http://www.vim.org/scripts/script.php?script_id=4945)
- [Gtags And Man Proto (GAMP) : Full function completion using Linux manual and GTags (not mandatory).](http://www.vim.org/scripts/script.php?script_id=4946)
- [Find in Files, feature like Notepad++ : "Find in Files and Directories" like Notepad++](http://www.vim.org/scripts/script.php?script_id=4947)
- [SearchSpecial : Generic functions for special search modes.](http://www.vim.org/scripts/script.php?script_id=4948)
- [SearchRepeat : Repeat the last type of search via n/N.](http://www.vim.org/scripts/script.php?script_id=4949)
- [Lua Support 2 : Lua IDE. Insert codesnippets, run, compile, and check the code and look up help.](http://www.vim.org/scripts/script.php?script_id=4950)

## 月間ダウンロードランキング

1. [AutomaticLaTeXPlugin : Background compilation, completion, bib serch, toc and other nice features.](http://www.vim.org/scripts/script.php?script_id=2945) (4402)
2. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (2397)
3. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (2189)
4. [pathogen.vim : Poor man's package manager. Easy manipulation of 'runtimepath' et al](http://www.vim.org/scripts/script.php?script_id=2332) (1345)
5. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (1337)
6. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (1007)
7. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (928)
8. [The NERD Commenter : A plugin that allows for easy commenting of code for many filetypes.](http://www.vim.org/scripts/script.php?script_id=1218) (862)
9. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (814)
10. [github theme : A gvim colorscheme based on github's syntax highlighting](http://www.vim.org/scripts/script.php?script_id=2855) (782)

## vim-jp issues

Open : 191 (+6) | Closed : 379 (+10)

- [Issue #555 : E416番の日本語エラーメッセージのtypo](https://github.com/vim-jp/issues/issues/555)
- [Issue #556 : NFAエンジンの \[X-Y\] のYに composing 文字だけを指定すると別の文字にマッチする](https://github.com/vim-jp/issues/issues/556)
- [Issue #557 : 正規表現: \\v,\\V と パターンの先頭,末尾の対応が一貫してない](https://github.com/vim-jp/issues/issues/557)
- [Issue #558 : nnoremap \<silent>で呼ぶ関数の中でechoを使うと文字が途切れ途切れになることがある](https://github.com/vim-jp/issues/issues/558)
- [Issue #559 : 画面最下行での utf8 文字列の表示結果が不正確](https://github.com/vim-jp/issues/issues/559)
- [Issue #560 : ターミナル上のvimでgetchar(0)またはgetchar(1)の\<ESC>がすぐに取得できない](https://github.com/vim-jp/issues/issues/560)
- [Issue #561 : 「プログラムから開く」に登録できない(らしい)](https://github.com/vim-jp/issues/issues/561)
- [Issue #562 : t:+ 小文字な関数参照型変数を作成できない](https://github.com/vim-jp/issues/issues/562)
- [Issue #563 : configure を 2 回実行すると if&#x5f;python が有効にならない](https://github.com/vim-jp/issues/issues/563)
- [Issue #564 : helpの motion.txt 中の '\[, &#x60;\[, '\], &#x60;\] に関する記述がわかりにくい](https://github.com/vim-jp/issues/issues/564)
- [Issue #565 : \<C-V>$ で作成された矩形に対するオペレータ](https://github.com/vim-jp/issues/issues/565)
- [Issue #566 : :e scp://user@host/するとuserhostに繋ぎにいく](https://github.com/vim-jp/issues/issues/566)
- [Issue #567 : Windows でディレクトリ/ファイル名の頭に ( や ) があると %,# が展開できない](https://github.com/vim-jp/issues/issues/567)
- [Issue #568 : ビジュアルモードで複数行yank後の\[count\]付きのビジュアル選択範囲がおかしい](https://github.com/vim-jp/issues/issues/568)
- [Issue #569 : quickfix を開いた状態ででコマンドラインウィンドウを使用するとメッセージが消える](https://github.com/vim-jp/issues/issues/569)
- [Issue #570 : BufLeave で :bd! すると SEGV](https://github.com/vim-jp/issues/issues/570)

