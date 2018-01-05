---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2017 年 12 月号

---


## 話題

*   [Fablic.vim #4](https://fablicvim.connpass.com/event/71673/) が 12/04 (月) に東京で開催されました
*   [Meguro.vim #7](https://megurovim.connpass.com/event/72311/) が 12/16 (土) に東京で開催されました
*   [Yokohama.vim #10](https://yokohamavim.connpass.com/event/73077/) が 12/10 (日) に神奈川で開催されました
* [vital.vim開発者会議2018-02](https://fablicvim.connpass.com/event/74842/)が来年である2018-02-15 (木) に開催されます

## 今月の新機能

今月の新機能及びユーザーに影響のある変更は以下のとおりです。

*   8.0.1361: `hiddenoff` が `diffopt` オプションに追加されました
*   8.0.1364: `win_screenpos()` 関数が追加されました
*   8.0.1381: `ch_read()` 関数で最終行がないテキストを読めるようになりました

## 今月の新機能

今月の新機能及びユーザーに影響のある変更は以下のとおりです。

## Vimに関する脆弱性

*   8.0.1263 で修正された .swp ファイルのパーミッションに関する問題に、 [CVE-2017-17087](https://nvd.nist.gov/vuln/detail/CVE-2017-17087) が割り当てられました。

## リリース情報

- [8.0.1360 : the Terminal highlighting doesn't work in a terminal](https://github.com/vim/vim/commit/76bb7196f5102c9929959d710e8ed97a19affa4d)
- [8.0.1361 : some users don't want to diff with hidden buffers](https://github.com/vim/vim/commit/97ce419201421f65f4764549ed80307a7ef9c7a6)
- [8.0.1362 : terminal window colors wrong when using Terminal highlighting](https://github.com/vim/vim/commit/a7c54cfcf825e8e99db03f4ccdb1a32cd0714c52)
- [8.0.1363 : recovering does not work when swap file ends in .stz](https://github.com/vim/vim/commit/af903e5d490ec9c6c49079f67de7e92e3c35a725)
- [8.0.1364 : there is no easy way to get the window position](https://github.com/vim/vim/commit/22044dc31788d9f1c2da7725269884d9923b4795)
- [8.0.1365 : when one channel test fails others fail as well](https://github.com/vim/vim/commit/24820691e6ba9dae41ef16a3d3e55589843b34f4)
- [8.0.1366 : balloon shows when cursor is in WinBar](https://github.com/vim/vim/commit/1ad022a9b81d7829d5dc98cf5b8d0ee410558040)
- [8.0.1367 : terminal test hangs, executing abcde](https://github.com/vim/vim/commit/1ad022a9b81d7829d5dc98cf5b8d0ee410558040)
- [8.0.1368 : cannot drag status or separator of new terminal window](https://github.com/vim/vim/commit/ce6179c799468e471c3b7fc71c9924f57a2253c5)
- [8.0.1369 : MS-Windows: drawing underline slow, mFallbackDC not updated](https://github.com/vim/vim/commit/92467d3351853de769329f62121bf34d28647546)
- [8.0.1370 : channel test for callback is flaky](https://github.com/vim/vim/commit/1eca6f13d62d2201f3bf0e83400c2e3de7df3951)
- [8.0.1371 : Shift-Insert doesn't always work in MS-Windows console](https://github.com/vim/vim/commit/feeb4d0901c7b16958e8f02ffcdac83b605b0b43)
- [8.0.1372 : profile log may be truncated halfway a character](https://github.com/vim/vim/commit/ac112f01a6930c9d15cf0360b657373699916bfd)
- [8.0.1373 : no error when settting 'renderoptions' before starting GUI](https://github.com/vim/vim/commit/3767c6e9ee1bd585a2afba7e932854b24d194a2d)
- [8.0.1374 : CTRL-A does not work with an empty line](https://github.com/vim/vim/commit/5fe6bdf858a7f2f288d599ffb5efb3c08449c817)
- [8.0.1375 : window size wrong after maximizing with WinBar](https://github.com/vim/vim/commit/415a6939a4e8d4e26b4af26c24eb75243d3a2756)
- [8.0.1376 : cursor in terminal not always updated](https://github.com/vim/vim/commit/23c1b2b018c8121ca5fcc247e37966428bf8ca66)
- [8.0.1377 : cannot call a dict function in autoloaded dict](https://github.com/vim/vim/commit/6e65d594aa33be11f6074f26e9ff81b52504c62b)
- [8.0.1378 : autoload script sources itself when defining function](https://github.com/vim/vim/commit/3388d334572f9d65a603d09d75e363805d96c5d9)
- [8.0.1379 : configure check for selinux does not check for header file](https://github.com/vim/vim/commit/e4b78e2a42e9030348770764c478b73caa925539)
- [8.0.1380 : using "vim -r swapfile" the hit-enter prompt is misplaced.](https://github.com/vim/vim/commit/05684310a53c8a4804441c1c6f8b7fc9e8194940)
- [8.0.1381 : ch&#x5f;readraw() waits for NL if channel mode is NL](https://github.com/vim/vim/commit/620ca2da372dc9c892022faff83d363c67cc5c45)
- [8.0.1382 : get "no write since last change" message if terminal is open](https://github.com/vim/vim/commit/f405c8fe85bba6dc96a68a12ab976f745fc51a38)
- [8.0.1383 : local additions in help skips some files](https://github.com/vim/vim/commit/35c5e8155da797f14124d98fdc6189067b965688)
- [8.0.1384 : not enough quickfix help; confusing winid](https://github.com/vim/vim/commit/74240d3febd1e3bc7cf086c647c9348b20716c33)
- [8.0.1385 : Python 3.5 is getting old](https://github.com/vim/vim/commit/59eb016dff3fd4f764cfe326c80a9b840f0e1a02)
- [8.0.1386 : cannot select modified buffers with getbufinfo()](https://github.com/vim/vim/commit/8e6a31df81113bbf0e4bb5324a74dc5f6c62a490)
- [8.0.1387 : wordcount test is old style](https://github.com/vim/vim/commit/a703aaee4d6230ed81469ef0cb37f331bf255136)
- [8.0.1388 : char not overwritten with ambiguous width char](https://github.com/vim/vim/commit/fae8ed1fc8c06b28528d726e8440dfc66852bca8)
- [8.0.1389 : getqflist() items are missing if not set](https://github.com/vim/vim/commit/a6d4849c711379b773529afaed640455287ac934)
- [8.0.1390 : DirectX scrolling can be slow, vertical positioning is off](https://github.com/vim/vim/commit/7f88b65f6cde3d69386e461f61271a271b79b7e8)
- [8.0.1391 : encoding empty string to JSON sometimes gives "null"](https://github.com/vim/vim/commit/b29d328eb403a243431f4d0d5b360b4ac46bbc3e)
- [8.0.1392 : build fails with --with-features=huge --disable-channel](https://github.com/vim/vim/commit/8b42328cef294a5d6059ee7c452da75b845f6f8c)
- [8.0.1393 : too much highlighting with 'hlsearch' and 'incsearch' set](https://github.com/vim/vim/commit/6621605eb97cf5fbc481282fd4d349a76e168f16)
- [8.0.1394 : cannot intercept a yank command](https://github.com/vim/vim/commit/7e1652c63c96585b9e2235c195a3c322b1f11595)
- [8.0.1395 : it is not easy to see if a colorscheme is well written](https://github.com/vim/vim/commit/200d0e36bc5384beb9dc76ac75806ac0aecf84ac)
- [8.0.1396 : memory leak when CTRL-G in search command line fails](https://github.com/vim/vim/commit/a1d5c154dbd5fbe317726bbf2ba99632b91878f4)
- [8.0.1397 : pattern with \\& following nothing gives an error](https://github.com/vim/vim/commit/890dd05492d88d48eee1dda7f7a1811d027ce7ca)
- [8.0.1398 : :packadd does not load packages from the "start" directory](https://github.com/vim/vim/commit/9e1d399e63903c6f84d7888ad8d84ebf4e29d8a1)
- [8.0.1399 : warnings and errors when building tiny version](https://github.com/vim/vim/commit/ee219b0e9faab4b7159ed1725c5b82cea4f4d4f8)
- [8.0.1400 : color scheme check script shows up as color scheme](https://github.com/vim/vim/commit/8ee2d36e216756e712a3a9122ce1e1203378a9c8)
- [8.0.1401 : cannot build with GTK but without XIM](https://github.com/vim/vim/commit/cb89c98c264b0fe9fd26e3501eab5a062e57c064)
- [8.0.1402 : crash with nasty autocommand](https://github.com/vim/vim/commit/9bca805ec49eb0d2d0d0b2093f418ff425500169)
- [8.0.1403 : using freed buffer in grep command](https://github.com/vim/vim/commit/4fb921e388b9a042573ab06cce36e92874954197)
- [8.0.1404 : invalid memory access on exit](https://github.com/vim/vim/commit/606d45ccd8a2ad2956e2729f6135fd79fd2f6d72)
- [8.0.1405 : duplicated code for getting a typed character](https://github.com/vim/vim/commit/c9e649ae816cdff0d1da8a97d40e695c6d3991bd)
- [8.0.1406 : difficult to track changes to a quickfix list](https://github.com/vim/vim/commit/b254af312d1696b12367085acfbe41a41b7f1ea5)
- [8.0.1407 : GUI: CursorHold may trigger before 'updatetime'](https://github.com/vim/vim/commit/4af031dbc8d62f89c94072a406f6d2ec0e5200be)
- [8.0.1408 : crash in setqflist()](https://github.com/vim/vim/commit/a0ca7d002d4efcf4bce0af6943146a339677ed3d)
- [8.0.1409 : buffer overflow in :tags command](https://github.com/vim/vim/commit/132f75255ecea17ff621f71236568c5d8d8e0163)
- [8.0.1410 : hang when using count() with an empty string](https://github.com/vim/vim/commit/338e47fdfdf0d918dae50a5cbf0cf4f7be45b4f0)
- [8.0.1411 : reading invalid memory with CTRL-W :](https://github.com/vim/vim/commit/2efb323e875d2852f63e41c40641760d1d6b069f)
- [8.0.1412 : using free memory using setloclist()](https://github.com/vim/vim/commit/12237448499aaeb8c4f2be7a1deda60c0f160627)
- [8.0.1413 : accessing freed memory in :cbuffer](https://github.com/vim/vim/commit/aaf6e43b7a99cedb89d73ba749a46f7a0f16bbb6)
- [8.0.1414 : accessing freed memory in :lfile.](https://github.com/vim/vim/commit/14a4deb064610c30a50f00d524dde9b3292aad59)
- [8.0.1415 : warning for unused function without timers feature](https://github.com/vim/vim/commit/4ce46c2a6b59586e329fb41ad25799872bffc2c8)
- [8.0.1416 : crash when searching for a sentence](https://github.com/vim/vim/commit/8ada6aa9298b4764d9ca0024dd21b17e815595ce)
- [8.0.1417 : test doesn't search for a sentence](https://github.com/vim/vim/commit/1bd999f982e783219a06e6c8f219df1d53ac7e77)
- [8.0.1418 : no test for expanding backticks](https://github.com/vim/vim/commit/ae6f8651251013bafef9de1aed09069deaae8122)
- [8.0.1419 : cursor column is not updated after \]s](https://github.com/vim/vim/commit/b73fa629d6d3d705c1f8e8d5f8109fc9abd7bb6f)
- [8.0.1420 : accessing freed memory in vimgrep](https://github.com/vim/vim/commit/3c09722600e3218905b5d4a7b635a9e6560f87b3)
- [8.0.1421 : accessing invalid memory with overlong byte sequence](https://github.com/vim/vim/commit/e6640ad44e2186bd3642b972115496d347cd1fdd)
- [8.0.1422 : no fallback to underline when undercurl is not set ](https://github.com/vim/vim/commit/45a0000d5c4a5ea934d5a16a2ab5bbcf307623ac)
- [8.0.1423 : error in return not caught by try/catch](https://github.com/vim/vim/commit/fabaf753e26df5a89a854673d14c15a1fa6b321a)
- [8.0.1424 : the timer&#x5f;pause test is flaky on Travis](https://github.com/vim/vim/commit/a47ebdbd222ae488a65be4e8bc3fd87c6301c000)
- [8.0.1425 : execute() does not work in completion of user command](https://github.com/vim/vim/commit/2095148277cf1c4e7b3bbaf4e34812b7cfe3011b)
- [8.0.1426 : "gf" and \<cfile> don't accept ? and & in URL](https://github.com/vim/vim/commit/9e3dfc650190e96739abc004eb9948afa68136b4)
- [8.0.1427 : the :leftabove modifier doesn't work for :copen](https://github.com/vim/vim/commit/de04654ddc865af94ef04b1738b335a924be7923)
- [8.0.1428 : compiler warning on 64 bit MS-Windows system](https://github.com/vim/vim/commit/200ea8ffaa90e1ccc156b24ee097be87acdd5214)

## 新着スクリプト

- [SQHell.vim : A lightweight SQL wrapper for vim](https://vim.sourceforge.io/scripts/script.php?script_id=5623)
- [match-up : even better %, navigate & highlight matching word, modern matchit replacement](https://vim.sourceforge.io/scripts/script.php?script_id=5624)
- [Github.vim : Another github v3 api implemented in vim script](https://vim.sourceforge.io/scripts/script.php?script_id=5625)
- [Libertine : A complex mix of seven strong bodied colors with citrus flavors and pine notes](https://vim.sourceforge.io/scripts/script.php?script_id=5626)
- [i3config.vim : Vim syntax highlighting for i3 config](https://vim.sourceforge.io/scripts/script.php?script_id=5627)
- [bioSyntax : Syntax Highlighting for Computational Biology](https://vim.sourceforge.io/scripts/script.php?script_id=5628)
- [rainbow&#x5f;levels.vim : A different approach to code highlighting.](https://vim.sourceforge.io/scripts/script.php?script_id=5629)
- [vim-terminal : A Vim plugin that opens an interactive terminal in a buffer for running programs](https://vim.sourceforge.io/scripts/script.php?script_id=5630)
- [place.vim : Enter characters without moving](https://vim.sourceforge.io/scripts/script.php?script_id=5631)
- [DumpX : Disassemble C/C++ code](https://vim.sourceforge.io/scripts/script.php?script_id=5632)
- [readline.vim : Readline-style bindings for command mode](https://vim.sourceforge.io/scripts/script.php?script_id=5633)
- [debugstring : Debug printf()-style at the speed of light ](https://vim.sourceforge.io/scripts/script.php?script_id=5634)
- [Hearts of Iron Syntax : Add syntax highlight for Hearts of Iron files](https://vim.sourceforge.io/scripts/script.php?script_id=5635)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](https://vim.sourceforge.io/scripts/script.php?script_id=273) (1266)
2. [The NERD tree : A tree explorer plugin for navigating the filesystem](https://vim.sourceforge.io/scripts/script.php?script_id=1658) (947)
3. [wombat256.vim : Wombat for 256 color xterms](https://vim.sourceforge.io/scripts/script.php?script_id=2465) (926)
4. [OmniCppComplete : C/C++ omni-completion with ctags database](https://vim.sourceforge.io/scripts/script.php?script_id=1520) (427)
5. [molokai : A port of the monokai scheme for TextMate](https://vim.sourceforge.io/scripts/script.php?script_id=2340) (393)
6. [nginx.vim : initial version](https://vim.sourceforge.io/scripts/script.php?script_id=1886) (390)
7. [python.vim : Enhanced version of the python syntax highlighting script](https://vim.sourceforge.io/scripts/script.php?script_id=790) (373)
8. [AutomaticLaTeXPlugin : Background compilation, completion, bib serch, toc and other nice features.](https://vim.sourceforge.io/scripts/script.php?script_id=2945) (373)
9. [AutoComplPop : Automatically opens popup menu for completions](https://vim.sourceforge.io/scripts/script.php?script_id=1879) (347)
10. [minibufexpl.vim : Elegant buffer explorer - takes very little screen space](https://vim.sourceforge.io/scripts/script.php?script_id=159) (338)

## vim-jp issues

Open : 255 (+1) | Closed : 882 (+11)

- [Issue #1126 : Ambiguous width character の直後に文字が残ることがある](https://github.com/vim-jp/issues/issues/1126)
- [Issue #1127 : ファイルをドロップすると E341: 内部エラー: lalloc(0,) が発生する](https://github.com/vim-jp/issues/issues/1127)
- [Issue #1128 : readdir() が欲しい](https://github.com/vim-jp/issues/issues/1128)
- [Issue #1129 : ユーザー定義コマンドの補完の関数内で execute() を実行すると結果が得られない](https://github.com/vim-jp/issues/issues/1129)
- [Issue #1130 : Linux版 gvimで imactivatefunc / imstatusfunc が 8.0.1344以降、呼ばれなくなっている](https://github.com/vim-jp/issues/issues/1130)
- [Issue #1131 : hl-Terminal のヘルプがない。:hi \<C-D>で Ternimal が候補にない。](https://github.com/vim-jp/issues/issues/1131)
- [Issue #1132 : Windows で印刷する際のドキュメント名が文字化けする](https://github.com/vim-jp/issues/issues/1132)
- [Issue #1133 : set linebreak+矩形ビジュアルモードで加算、減算、置き換えでおかしな位置が編集される](https://github.com/vim-jp/issues/issues/1133)
- [Issue #1134 : elseif 行のプロファイル結果がおかしい](https://github.com/vim-jp/issues/issues/1134)
- [Issue #1135 : return の右項で例外が発生すると try/catch で補足されない。](https://github.com/vim-jp/issues/issues/1135)
- [Issue #1136 : Windows の job でダブルコーテーションを付けるとコマンドを正常実行できない場合がある](https://github.com/vim-jp/issues/issues/1136)
- [Issue #1137 : 入力補完中にコマンドラインウインドウに表示される"match N of M"といったメッセージを抑制したい](https://github.com/vim-jp/issues/issues/1137)
