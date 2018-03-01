---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2018 年 02 月号

---
## 話題

*   [Meguro.vim #8](https://megurovim.connpass.com/event/76881/)が 2018-02-17 (土) に開催されました
*   [Osaka.vim #12](https://osaka-vim.connpass.com/event/77504/)が 2018-02-17 (土) に開催されました
*   [vital.vim開発者会議2018-03](https://fablicvim.connpass.com/event/74842/)が 参加者募集中です
*   <https://www.vim.org/> が HTTPS になりました。

## 今月の新機能

今月の新機能及びユーザーに影響のある変更は以下のとおりです。

*   8.0.1459: ディレクトリが変更されたときに実行される自動コマンドイベント `DirChanged` が追加されました (関連パッチ: 8.0.1460, 8.0.1461, 8.0.1462, 8.0.1463)
*   8.0.1489: `getcwd(-1)` でグローバルなディレクトリを取得できるようになりました
*   8.0.1491: ポップアップメニューの幅を制御するためのオプション `'pumwidth'` が追加されました (関連パッチ: 8.0.1495, 8.0.1522, 8.0.1540)
*   8.0.1497: ジャンプリストを取得する関数 `getjumplist()` が追加されました (関連パッチ: 8.0.1498, 8.0.1513)
*   8.0.1508: `:drop` コマンドが非 GUI 環境でも動作するようになりました
*   8.0.1510: ビープのテスト用関数 `assert_beeps()` が追加されました
*   8.0.1514: 変更リストを取得する関数 `getchangelist()` が追加されました (関連パッチ: 8.0.1519)
*   8.0.1521: 端末ウィンドウ内で Shift-Tab が使えるようになりました
*   8.0.1523: 端末ウィンドウのダンプを取って比較するための関数 `term_dumpwrite()`, `term_dumpload()`, `term_dumpdiff()`, `assert_equalfile()` が追加されました (関連パッチ: 8.0.1529, 8.0.1542)
*   8.0.1531: Windows 10 Version 1703 (Creators Update) 以降のコマンドプロンプトで 24bit 色が使えるようになりました
*   8.0.1553: `ga` でダイグラフが表示されるようになりました

## Vimに関する脆弱性

特筆すべき脆弱性の報告はありませんでした。

## リリース情報

- [8.0.1452 : terminal test fails on some systems](https://github.com/vim/vim/commit/d4a282f7a5efb2bf28fbd947251e512f08563738)
- [8.0.1453 : terminal test fails on some slow terminals](https://github.com/vim/vim/commit/218101442e8775bd1f05d1dda18d3946ee032dd1)
- [8.0.1454 : when in silent mode too much output is buffered](https://github.com/vim/vim/commit/42b23fad1d9cdd6266f52d1ed7e0f3f17ce2d04b)
- [8.0.1455 : if $SHELL contains a space then 'shell' is incorrect](https://github.com/vim/vim/commit/4bfa8af14142e54f509048239f4e8596911f56aa)
- [8.0.1456 : timer test on travis Mac is still flaky](https://github.com/vim/vim/commit/8dce6c54c81f13582617db3127071b3436b8163b)
- [8.0.1457 : clojure now supports a shebang line](https://github.com/vim/vim/commit/8fd2ffc530f8f01afac53903e6ff95867afd94d7)
- [8.0.1458 : filetype detection test does not check all scripts](https://github.com/vim/vim/commit/ddb349369d107c14fad9c38baf2f0e2b8514fbf0)
- [8.0.1459 : cannot handle change of directory](https://github.com/vim/vim/commit/b7407d3fc9496f9048fb65ab17b5ba3444965c0e)
- [8.0.1460 : (after 8.0.1459) missing file in patch](https://github.com/vim/vim/commit/b5cb65ba2bcc6bbc6d2798a2dea18b95f0b38f5e)
- [8.0.1461 : (after 8.0.1459) missing another file in patch](https://github.com/vim/vim/commit/15833239a4131279935a4bd574b74fe3a2b0f49f)
- [8.0.1462 : (after 8.0.1459) missing yet another file in patch](https://github.com/vim/vim/commit/f4aba797cb57d6d5b3b491bd5ae2c0b40e07baaf)
- [8.0.1463 : (after 8.0.1459) test fails without 'autochdir' option](https://github.com/vim/vim/commit/ec48a9c58989babcad23d73483955f35b6e41492)
- [8.0.1464 : completing directory after :find does not add slash](https://github.com/vim/vim/commit/8a37b032895b40dd6953280c33585bcba0c7ef8b)
- [8.0.1465 : python2 and python3 detection not tested](https://github.com/vim/vim/commit/4bc0bed53695ac67db8d601f2a15e48e7a196688)
- [8.0.1466 : older GTK versions don't have gtk&#x5f;entry&#x5f;get&#x5f;text&#x5f;length()](https://github.com/vim/vim/commit/06b77ef69f252e1ba8a2136dcbed6622bc2371bb)
- [8.0.1467 : libvterm doesn't handle illegal byte sequence correctly](https://github.com/vim/vim/commit/fef4ddd5eb8816a6607a624aa401bcfa71a63def)
- [8.0.1468 : illegal memory access in del&#x5f;bytes()](https://github.com/vim/vim/commit/191f18bad0b5c48afa05c3e8a00f3ced993f6a38)
- [8.0.1469 : when package path is a symlink 'runtimepath' is wrong](https://github.com/vim/vim/commit/2374faae111057ee28e8d487f9a52a95855e2206)
- [8.0.1470 : integer overflow when using regexp pattern](https://github.com/vim/vim/commit/2c7b906afb86b986476cfc959732e433b1b4a3b1)
- [8.0.1471 : on MS-Windows CursorIM highlighting no longer works](https://github.com/vim/vim/commit/28944fecff3c40b44325921d45aaf67451b0937f)
- [8.0.1472 : MS-Windows: nsis installer is a bit slow](https://github.com/vim/vim/commit/5d4247402b7195c6872485ddf1600a1cea723027)
- [8.0.1473 : MS-Windows: D&D fails between 32 and 64 bit apps](https://github.com/vim/vim/commit/511ffdd65d48b0597ed10614d161b5e811342058)
- [8.0.1474 : Visual C 2017 has multiple MSVCVER numbers](https://github.com/vim/vim/commit/dd08b6a32b639b8c7a11275e04ae0a7ffc43aed0)
- [8.0.1475 : invalid memory access in read&#x5f;redo()](https://github.com/vim/vim/commit/f12519dec88251305793f1651f558d16506b4be2)
- [8.0.1476 : screen isn't always updated right away](https://github.com/vim/vim/commit/acda04f5c641330cd589ca52eb61d1ab0d62385f)
- [8.0.1477 : redraw flicker when moving mouse outside of terminal window](https://github.com/vim/vim/commit/d317b38a4dbcdbd4cf587cf78cc1f4289374214e)
- [8.0.1478 : unnecessary condition](https://github.com/vim/vim/commit/dff72ba4459f54cac2ce40eea3d92097660c7b9f)
- [8.0.1479 : insert mode completion state is confusing](https://github.com/vim/vim/commit/bc0e9adae9c253f36803665180e4b576d1e725ab)
- [8.0.1480 : (after 8.0.1479) patch missing change](https://github.com/vim/vim/commit/0562532c2eee6205d225aa1dc7e3e89af0dfd990)
- [8.0.1481 : clearing a pointer takes two lines](https://github.com/vim/vim/commit/1567558b20575e1b17c3808c6bd622b0b4810e36)
- [8.0.1482 : using feedkeys() does not work to test completion](https://github.com/vim/vim/commit/02ae9b4a93deea4993d7abe20485f91f1cce5e36)
- [8.0.1483 : searchpair() might return an invalid value on timeout](https://github.com/vim/vim/commit/9d32276b52a63fccfae681f0d1d6ccb66efec1c0)
- [8.0.1484 : reduntant conditions](https://github.com/vim/vim/commit/a15ef4588c057da87f64be5b17aef35aaead8ac8)
- [8.0.1485 : weird autocmd may cause arglist to be changed recursively](https://github.com/vim/vim/commit/9e33efd1523b85a70533930dd43a26925a2b648c)
- [8.0.1486 : accessing invalid memory with "it"](https://github.com/vim/vim/commit/82846a00ac0c135946c93c48c1657018a5c96b11)
- [8.0.1487 : (after 8.0.1486) test 14 fails](https://github.com/vim/vim/commit/8846ac5aedb568b9aae969f0ad2b9b1606522ca9)
- [8.0.1488 : emacs tags no longer work](https://github.com/vim/vim/commit/0d2073773218736e368786f0db7024bd9b9e7912)
- [8.0.1489 : there is no easy way to get the global directory](https://github.com/vim/vim/commit/5459129af2a832a027a1e7ca2d6177c26647d64f)
- [8.0.1490 : number of spell regions is spread out through the code](https://github.com/vim/vim/commit/2993ac5fce5450428322ce43aaa5e643e6994443)
- [8.0.1491 : the minimum width of the popup menu is hard coded](https://github.com/vim/vim/commit/a8f04aa275984183bab5bb583b128f38c64abb69)
- [8.0.1492 : memory leak in balloon&#x5f;split()](https://github.com/vim/vim/commit/b301f6b950975b9d7ae87a4f551b32bba63ccdcf)
- [8.0.1493 : completion items cannot be annotated](https://github.com/vim/vim/commit/9b56a57cdae31f7a2c85d440392bf63d3253a158)
- [8.0.1494 : no autocmd triggered in Insert mode with visible popup menu](https://github.com/vim/vim/commit/5a093437199001a0d60d8e18e2b9539b99a7757c)
- [8.0.1495 : having 'pumwidth' default to zero has no merit](https://github.com/vim/vim/commit/42443c7d7fecc3a2a72154bb6139b028438617c2)
- [8.0.1496 : clearing a pointer takes two lines](https://github.com/vim/vim/commit/d23a823669d93fb2a570a039173eefe4856ac806)
- [8.0.1497 : getting the jump list requires parsing the output of :jumps](https://github.com/vim/vim/commit/4f50588ba336e7f086a72c53f5688c2494fc34b3)
- [8.0.1498 : getjumplist() returns duplicate entries](https://github.com/vim/vim/commit/a7e18d237f817637815f0de44b08df1e0ca0f4f9)
- [8.0.1499 : out-of-memory situation not correctly handled](https://github.com/vim/vim/commit/4b7e7bed6cb16c8256e6973418701cfb15b30b8e)
- [8.0.1500 : possible NULL pointer dereference](https://github.com/vim/vim/commit/0549a1e184d33674f4c2b8fb44ccdf6b9b9808a3)
- [8.0.1501 : out-of-memory situation not correctly handled](https://github.com/vim/vim/commit/a0221df149aa3773450b3f930299a409dd75bd5b)
- [8.0.1502 : in out-of-memory situation character is not restored](https://github.com/vim/vim/commit/71a43c01377cb0c5cdc5f2d9a357b5ef1aa69ee3)
- [8.0.1503 : access memory beyond end of string](https://github.com/vim/vim/commit/cdd09aa51a8d34bb384460af4f91026dbff5bf48)
- [8.0.1504 : Win32: the screen may be cleared on startup](https://github.com/vim/vim/commit/c33ecb291518044f661c5fc10b704fc718321794)
- [8.0.1505 : debugger can't break on a condition](https://github.com/vim/vim/commit/c6f9f739d32084923c3031cbf6f581f8c8bf7fd2)
- [8.0.1506 : new version of HP NonStop (Tandem) doesn't like a header](https://github.com/vim/vim/commit/c19e1d1ba52e4f19f7fd03b45ea236808a754b92)
- [8.0.1507 : timer test is a bit flaky](https://github.com/vim/vim/commit/bfbea567d89fdaa08ed987fd80daa53a6ce399d1)
- [8.0.1508 : the :drop command is not always available](https://github.com/vim/vim/commit/5a656864a0610547da28e0c8c1649ecd1d782948)
- [8.0.1509 : (after 8.0.1508) failing drag-n-drop command no longer fails](https://github.com/vim/vim/commit/294959528e02403cd7ef6541208835f0c621c63b)
- [8.0.1510 : cannot test if a command causes a beep](https://github.com/vim/vim/commit/b48e96f61c87a64e38e3ac50732c92a84a4833b8)
- [8.0.1511 : some code for the debugger watch expression is clumsy](https://github.com/vim/vim/commit/3198870137df64214317151726648af8e56f1729)
- [8.0.1512 : warning for possibly using NULL pointer](https://github.com/vim/vim/commit/e4db7aedab65abadcc84c78e7a10ec7bb62f11cf)
- [8.0.1513 : the jumplist is not always properly cleaned up](https://github.com/vim/vim/commit/486797413791f6be12dcec6e5faf4f952e4647ae)
- [8.0.1514 : getting the list of changes is not easy](https://github.com/vim/vim/commit/07ad816525da67cab3c0db21d1286d221dbc7477)
- [8.0.1515 : BufWinEnter event fired when opening hidden terminal](https://github.com/vim/vim/commit/ab5e7c3deb40328e1c7d40534ed8cd9da47a92f8)
- [8.0.1516 : errors for job options are not very specific](https://github.com/vim/vim/commit/b3292fa2d362b064ff8a115fc0ad794c1f1265d7)
- [8.0.1517 : invalid memory acces with pattern using look-behind match](https://github.com/vim/vim/commit/bc197195b097707d08fd44a476dbc374366504cb)
- [8.0.1518 : error messages suppressed after ":silent! try"](https://github.com/vim/vim/commit/2be57331524e93da52a0663f4a334d21c05123bb)
- [8.0.1519 : getchangelist() does not use argument as bufname()](https://github.com/vim/vim/commit/341a64c9cabff08e4a7dc8cd932a598e12134457)
- [8.0.1520 : cursor in wrong line when using a WinBar in Terminal window](https://github.com/vim/vim/commit/181ca99e163b145cd1a4ec47a50cc25cf3994109)
- [8.0.1521 : Shift-Tab does not work in a terminal window](https://github.com/vim/vim/commit/73cddfd559152ea9b7e978ea7cf9c0d3a41e7316)
- [8.0.1522 : popup menu is positioned in the wrong place](https://github.com/vim/vim/commit/4287ed33ddc324d26dd05d3e19596dd74cf479d6)
- [8.0.1523 : cannot write and read terminal screendumps](https://github.com/vim/vim/commit/d96ff165113ce5fe62107add590997660e3d4802)
- [8.0.1524 : (after 8.0.1523) compiler warnings for uninitialized vars](https://github.com/vim/vim/commit/9c8816bd306a003c2ac3dce161be3fef481c9902)
- [8.0.1525 : using :wqa exits even if a job runs in a terminal window](https://github.com/vim/vim/commit/7a76092a51fc5446426a4bfd9eb6503ec61bf9e9)
- [8.0.1526 : no test using a screen dump yet](https://github.com/vim/vim/commit/da65058a9c4774dc534c7ae98d24c58b5db669fa)
- [8.0.1527 : screen dump test fails on MS-Windows](https://github.com/vim/vim/commit/3cc9f7440d857ff8360c15bb11e4e6229463920e)
- [8.0.1528 : dead code found](https://github.com/vim/vim/commit/81226e03102dd00b7cdce0e00775e1e30462f9a6)
- [8.0.1529 : assert&#x5f;equalfile() does not close file descriptors](https://github.com/vim/vim/commit/3049418f3dbc571463a04d068069f6c5b7a8ccf1)
- [8.0.1530 : dump test fails when using a shadow directory](https://github.com/vim/vim/commit/19eb6658eced1b1ce3d1097187ee28c28e352f2e)
- [8.0.1531 : cannot use 24 bit colors in MS-Windows console](https://github.com/vim/vim/commit/cafafb381a04e33f3ce9cd15dd9f94b73226831f)
- [8.0.1532 : compiler warnings without termguicolors feature](https://github.com/vim/vim/commit/cc0f2be88046bd1c07efa444bba6c05efe15ddd5)
- [8.0.1533 : libterm doesn't support requesting fg and bg color](https://github.com/vim/vim/commit/674e482d1346aa1afddab62675f3a7d7a00a4894)
- [8.0.1534 : C syntax test fails when using gvim](https://github.com/vim/vim/commit/b7ea7cb8e430ea096b4c452cdc9c3299819e6d6b)
- [8.0.1535 : C syntax test still fails when using gvim](https://github.com/vim/vim/commit/6acadda8d60892ddf06449f1cc4286912b0c0c2b)
- [8.0.1536 : quotestar test is flaky when using the GUI](https://github.com/vim/vim/commit/791010e648a68490440e202222938c8a93b9d0d3)
- [8.0.1537 : xxd does not skip NUL lines when using ebcdic](https://github.com/vim/vim/commit/085346f5a1ab5828b1fd80990d93172440c54724)
- [8.0.1538 : popupmenu is too far left when completion is long](https://github.com/vim/vim/commit/bb008dd3239c5fe3ac04501e38e4c950fa9426c8)
- [8.0.1539 : no test for the popup menu positioning](https://github.com/vim/vim/commit/6bb2cdfe604e51eec216cbe23bb6e8fb47810347)
- [8.0.1540 : popup menu positioning fails with longer string](https://github.com/vim/vim/commit/2b10bcbfc1c025bf7e6358326ee70105e7d30e96)
- [8.0.1541 : synpat&#x5f;T is taking too much memory](https://github.com/vim/vim/commit/36f923014a7eb7e24c4b0b88719cad14351e3a60)
- [8.0.1542 : terminal screen dump does not include cursor position](https://github.com/vim/vim/commit/9271d058c92c94b696eed5da24a69c077f42bc91)
- [8.0.1543 : with 'termguicolors' Normal color doesn't work correctly](https://github.com/vim/vim/commit/33ef5bb0e4ea9a049c4311cfe59d1926dcb963a4)
- [8.0.1544 : when using 'termguicolors' SpellBad doesn't show](https://github.com/vim/vim/commit/d4fc577e60d325777d38c00bd78fb9a32c7b1dfa)
- [8.0.1545 : screen dumps not included in distribution](https://github.com/vim/vim/commit/8226ac6b5916c736dde751df9cb8d03c19cb5611)
- [8.0.1546 : using feedkeys() in a terminal may trigger mappings](https://github.com/vim/vim/commit/c8bcfe7efd4cfdfd7f503700aba45c42765cd4dc)
- [8.0.1547 : undo in the options window makes it empty](https://github.com/vim/vim/commit/9c474b277336235012f0c058bbaef2961d1ec83e)
- [8.0.1548 : screen dump test script not included in distribution](https://github.com/vim/vim/commit/75542ec9f67146d075117003b92a5f86736754d2)
- [8.0.1549 : various small problems in test files](https://github.com/vim/vim/commit/5d7ead3bc85eefd0929bfcbb579510c8164ea1be)
- [8.0.1550 : various small problems in source files](https://github.com/vim/vim/commit/792f0e36593d1ec13ccb8a622ca5542c500577b4)
- [8.0.1551 : on Mac 'maxmemtot' is set to a weird value](https://github.com/vim/vim/commit/988615f26f262d9ef6472c53b48868968a6b6d16)
- [8.0.1552 : may leak file descriptors when executing job](https://github.com/vim/vim/commit/8195247054a659fe5cbc238197634d5e13e8e8e9)
- [8.0.1553 : cannot see what digraph is used to insert a character](https://github.com/vim/vim/commit/5f73ef8d20070cd45c9aea4dc33c2e0657f5515c)

## 新着スクリプト

- [vim-multiselect : A library for multiple selection](https://www.vim.org/scripts/script.php?script_id=5654)
- [vim-masquerade : Tools to edit multiple selections](https://www.vim.org/scripts/script.php?script_id=5655)
- [Toki Pona : syntax coloring and filetype settings for the Toki Pona minimalist conlang](https://www.vim.org/scripts/script.php?script_id=5656)
- [JumpToVerticalBlock : Like W / E, but vertically in the same column.](https://www.vim.org/scripts/script.php?script_id=5657)
- [vim-ruby-block-helpers : Helpers for maneuvering between ruby blocks, and context output in rspec files](https://www.vim.org/scripts/script.php?script_id=5658)
- [copytoggle.vim : Vim's pastetoggle for copying](https://www.vim.org/scripts/script.php?script_id=5659)
- [streamline.vim : Essential wares, to get there](https://www.vim.org/scripts/script.php?script_id=5660)
- [textwins.vim : Windows that text](https://www.vim.org/scripts/script.php?script_id=5661)
- [txtmux.vim : ( tabwins x textwins ) multiplexer](https://www.vim.org/scripts/script.php?script_id=5662)
- [aa : A plugin for managing time and note taking](https://www.vim.org/scripts/script.php?script_id=5663)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](https://www.vim.org/scripts/script.php?script_id=273) (650)
2. [wombat256.vim : Wombat for 256 color xterms](https://www.vim.org/scripts/script.php?script_id=2465) (534)
3. [The NERD tree : A tree explorer plugin for navigating the filesystem](https://www.vim.org/scripts/script.php?script_id=1658) (422)
4. [nginx.vim : initial version](https://www.vim.org/scripts/script.php?script_id=1886) (251)
5. [OmniCppComplete : C/C++ omni-completion with ctags database](https://www.vim.org/scripts/script.php?script_id=1520) (219)
6. [python.vim : Enhanced version of the python syntax highlighting script](https://www.vim.org/scripts/script.php?script_id=790) (217)
7. [AutomaticLaTeXPlugin : Background compilation, completion, bib serch, toc and other nice features.](https://www.vim.org/scripts/script.php?script_id=2945) (212)
8. [rails.vim : Ruby on Rails: easy file navigation, enhanced syntax highlighting, and more](https://www.vim.org/scripts/script.php?script_id=1567) (207)
9. [AutoComplPop : Automatically opens popup menu for completions](https://www.vim.org/scripts/script.php?script_id=1879) (201)
10. [c.vim : C/C++ IDE -- Write and run programs. Insert statements, idioms, comments etc.](https://www.vim.org/scripts/script.php?script_id=213) (184)

## vim-jp issues

Open : 261 (+2) | Closed : 893 (+6)

- [Issue #1147 : ユーザー定義コマンドの中から i&#x5f;CTRL-O 内で実行されたかを知りたい](https://github.com/vim-jp/issues/issues/1147)
- [Issue #1148 : ch&#x5f;open() で開いた channel からの出力をバッファに出力させたい](https://github.com/vim-jp/issues/issues/1148)
- [Issue #1149 : Vim script で CPU のエンディアンを取得する方法](https://github.com/vim-jp/issues/issues/1149)
- [Issue #1150 : :options コマンドで生成されるバッファが空になるまで undo 出来る](https://github.com/vim-jp/issues/issues/1150)
- [Issue #1151 : 文字入力開始時に確実にIMEをオフにしておきたい](https://github.com/vim-jp/issues/issues/1151)
- [Issue #1152 : マクロ記録中かどうかを取得・判断する手段が欲しい](https://github.com/vim-jp/issues/issues/1152)
- [Issue #1153 : インサートモード補完ポップアップメニューの表示位置と幅がおかしい](https://github.com/vim-jp/issues/issues/1153)
- [Issue #1154 : 外部プログラムの補完に重複が生じうる](https://github.com/vim-jp/issues/issues/1154)
