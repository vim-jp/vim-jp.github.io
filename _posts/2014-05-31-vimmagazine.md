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

- [7.4.274 : check for timestamp does not always detect a changed file](https://github.com/vim/vim/commit/a7611f60cde29639b655ddb4494b98886c318870)
- [7.4.275 : no error message when changing type of unplaced sign](https://github.com/vim/vim/commit/f4d7f167f3c1e26f26c4b5905db553eb94aa9e81)
- [7.4.276 : the fish shell is not supported](https://github.com/vim/vim/commit/75a8d74cc22d66482a1149f004b04bcc0a8326f2)
- [7.4.277 : using ":sign unplace &#x2a;" may leave the cursor in wrong position](https://github.com/vim/vim/commit/0d3d5e0483ddad68d333284d988b4269674f71dd)
- [7.4.278 : list&#x5f;remove() conflicts with Sun header file](https://github.com/vim/vim/commit/3ec7f4e4025c5a78ccd312f8516ac0740aa65dfe)
- [7.4.279 : globpath() returns a string, a list would be better](https://github.com/vim/vim/commit/1b1063af58b015b7827168f8fa2631efb60c252b)
- [7.4.280 : relative position of cursor not restored if session has 2 tabs](https://github.com/vim/vim/commit/56b3bf89aa2021d20afabd8ec7b513beb1b8df25)
- [7.4.281 : positions are off when loading session with two tabpages](https://github.com/vim/vim/commit/04ad7fe05c1c15c31949022d48e3008158f0657a)
- [7.4.282 : (after 7.4.279) test 97 fails on Mac](https://github.com/vim/vim/commit/b7547420a4fa576cfc6fb8b620bfc05ba8429f38)
- [7.4.283 : (after 7.4.276) compiler warning about unused variable](https://github.com/vim/vim/commit/d442ec7fe2f929d098ee2739e3c98c70e3bf45a4)
- [7.4.284 : setting 'langmap' in the modeline can cause trouble](https://github.com/vim/vim/commit/ed7547da438e44e7907bd9fd488f6b961f23e1e0)
- [7.4.285 : line numbers are not always updated for 'relativenumber'](https://github.com/vim/vim/commit/fd859c9878c9457f281941d55d3ca1124e9f90d5)
- [7.4.286 : error messages are inconsistant](https://github.com/vim/vim/commit/59838520c7a452adbe82b583dcdf967483b00a98)
- [7.4.287 : patches for .hgignore don't work](https://github.com/vim/vim/commit/46ae3cedeb8785299e825970f46763eff73399c5)
- [7.4.288 : when 'spellfile' is set the screen is not redrawn](https://github.com/vim/vim/commit/be578edae30d1292ddbf1f63518175fa45567ab8)
- [7.4.289 : NFA regexp with repeated backreference does not match](https://github.com/vim/vim/commit/ee4825331aeb3f517f8a51d0c21c743655352fb3)
- [7.4.290 : a non-greedy match followed by a branch is too greedy](https://github.com/vim/vim/commit/9c235069b237caed426eff14ce91df46bc69e234)
- [7.4.291 : compiler warning for int to pointer of different size](https://github.com/vim/vim/commit/6e13207d82aa925e39eabc0f190c8696481fdf04)
- [7.4.292 : pattern "a" matches accented "a" when pattern is one byte](https://github.com/vim/vim/commit/6082bea6acae266c392cd25317414cf3a167a596)
- [7.4.293 : can't ignore composing characters in a pattern](https://github.com/vim/vim/commit/8df5acfda9e5a41e643066e044f4adb47bfb83ba)
- [7.4.294 : (after 7.4.293) test files missing from patch](https://github.com/vim/vim/commit/caca646bf0f5e55fe32604e90b87b02034e2812b)
- [7.4.295 : various typos, bad white space and unclear comments](https://github.com/vim/vim/commit/24dc230871258018d2a6a616039283cb7f17684d)
- [7.4.296 : can't run tests on Solaris](https://github.com/vim/vim/commit/1c5a2134dbf880859281c676b834c5a636f78f52)
- [7.4.297 : memory leak from result of get&#x5f;isolated&#x5f;shell&#x5f;name()](https://github.com/vim/vim/commit/050fe7ebad159f48704bd3a991632d0330ef0cf2)
- [7.4.298 : can't have a funcref start with "t:"](https://github.com/vim/vim/commit/cbc67723bde3c05e4bc0788457961b41304576ee)
- [7.4.299 : when running configure twice DYNAMIC&#x5f;PYTHON&#x5f;DLL becomes empty](https://github.com/vim/vim/commit/cf1b05793948d10401ba07c77b4b3821994755d1)
- [7.4.300 : the way config.cache is removed doesn't always work](https://github.com/vim/vim/commit/7b129a5ac475a22a363e348be045932233c6d386)
- [7.4.301 : (after 7.4.280) scrolling problem when loading a session file](https://github.com/vim/vim/commit/b42970535c205baede141c8be2fab0f8cc4e16d9)
- [7.4.302 : signs don't show up after filler lines](https://github.com/vim/vim/commit/bc6cf6c5117d9ef19904c6fb56a2a45faa33afdb)
- [7.4.303 : double-width chars on the command line sometimes truncated](https://github.com/vim/vim/commit/e4c21e6899b6e4d9d4e0b29fc43c604edd3bc507)
- [7.4.304 : cannot always use Python with Vim](https://github.com/vim/vim/commit/b777618d111e51b55ee452d3881d5ef5bc836765)
- [7.4.305 : making 'ttymouse' empty causes problems](https://github.com/vim/vim/commit/90013c6a751e6574fbe49b529a3dbf1eacbef43d)
- [7.4.306 : getchar(0) does not return Esc](https://github.com/vim/vim/commit/9a665ba296bfe0e96d9f72d9f50b956e210e50c2)
- [7.4.307 : can't build without the +termresponse feature](https://github.com/vim/vim/commit/cbc17d66be624d911354ef7a6f18a9c903e80731)
- [7.4.308 : ":diffsplit" on empty file cursor is displayed on command line](https://github.com/vim/vim/commit/846a2ff5f3a006fb612d1ece1962cb75ca06d0fe)
- [7.4.309 : increasing size of lower window, upper window jumps to the top](https://github.com/vim/vim/commit/dd0402a759488fba55aef4350da521e7b65e8855)
- [7.4.310 : getpos()/setpos() don't include curswant](https://github.com/vim/vim/commit/493c178a2851587ca3f9deae3b41ae3acf5c900b)
- [7.4.311 : can't use winrestview to only restore part of the view](https://github.com/vim/vim/commit/82c2585eb840a84240ef927c09df28d5a04302b1)
- [7.4.312 : cannot know what argument list is being used for a window](https://github.com/vim/vim/commit/2d1fe05fc06a91886635680886c200ad90f0abd7)
- [7.4.313 : (after 7.4.310) changing return value of getpos() causes error](https://github.com/vim/vim/commit/6f6c0f8085a5b0855f9dce8378086fd5e06a219b)
- [7.4.314 : completion messages can get in the way of a plugin](https://github.com/vim/vim/commit/ea389e910326a0b44a9c534535c1c4c8b1841388)
- [7.4.315 : (after 7.4.309) fixes for computation of topline not tested](https://github.com/vim/vim/commit/c35b0fb181264658e8221b1ccb249e8551af8a59)
- [7.4.316 : warning from 64-bit compiler](https://github.com/vim/vim/commit/7116aa0f7d4c3f50f6898798c3047f5232aabe4a)

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

