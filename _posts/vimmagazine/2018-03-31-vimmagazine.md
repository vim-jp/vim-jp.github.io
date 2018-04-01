---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2018 年 03 月号

---
## 話題

*   [vital.vim開発者会議2018-03](https://fablicvim.connpass.com/event/74842/) が 03/01 (木) に開催されました
*   [Yokohama.vim #11](https://yokohamavim.connpass.com/event/81539/) が 03/31 (土) に開催されました
*   [Meguro.vim #9](https://megurovim.connpass.com/event/81678/) が 04/07 (土) に開催されます

## 今月の新機能

今月の新機能及びユーザーに影響のある変更は以下のとおりです。

*   8.0.1554: `--clean` オプションで起動された場合は `'runtimepath'` に `~/.vim` (`~/vimfiles`) が含まれないようになりました
*   8.0.1557: `printf()` 関数が引数が1つのみの場合にエラーが出ないようになりました
*   8.0.1558: 端末ウィンドウ内で右クリックした際にポップアップメニューが出るようになりました
*   8.0.1562: 端末デバッガーが右クリックでブレークポイントを設定できるようになりました
*   8.0.1563: `getwinpos()` 関数が追加されました (関連: 8.0.1574)
*   8.0.1570: `:popup` コマンドが端末でも動くようになりました
*   8.0.1587: クリップボードレジスタ (`*`, `+`) からインサートモードでペーストする場合は文字通り挿入されるようになりました
*   8.0.1592: [terminal セッションの復帰](http://vim-jp.org/blog/2018/03/12/recent-changes.html)が実装されました
*   8.0.1593: [端末ウィンドウを閉じようとした時の挙動を設定できるようになりました (デフォルトでは閉じようとするとE947エラーが発生します)](http://vim-jp.org/blog/2018/03/12/recent-changes.html)
    * `term_setkill()` 関数が追加されました
    * `term_setrestore()` 関数のオプションに `term_kill` が追加されました
    * `:terminal` コマンドのオプションに `++kill` が追加されました
    * GVim でウィンドウの × ボタンをクリックすると実行中の端末ウィンドウがあっても自動終了するようになりました
*   8.0.1595: [`ExitPre` イベントが追加されました](http://vim-jp.org/blog/2018/03/12/recent-changes.html)
*   8.0.1596: [`TerminalOpen` イベントが追加されました](http://vim-jp.org/blog/2018/03/12/recent-changes.html)
*   8.0.1602: `{[]: 42}` のようなキーが配列の JSON を `json_decode()` するとクラッシュする問題が修正されました
*   8.0.1607: `--clean` オプションで起動した際に .gvimrc も読み込まないようになりました
*   8.0.1608: `+directx` 機能がデフォルトで有効になりました
*   8.0.1609: `'guioptions'` に `!` フラグが追加されました (関連: 8.0.1616, 8.0.1617)
*   8.0.1612: 引数なしの `:terminal` コマンドがデフォルトでシェル終了後に閉じるようになりました
*   8.0.1624: `term_dumpdiff()`, `term_dumpload()` 関数の各種未実装オプションが実装されました
*   8.0.1630: `trim()` 関数が追加されました
*   8.0.1641: 端末ウィンドウ内のコマンドが Vim へメッセージを送信するためのエスケープシーケンスを用いた端末 API が追加されました (関連: 8.0.1643-1645, 8.0.1647)
*   8.0.1649: 引数リストコマンドの補完が追加されました
*   8.0.1651: `:ls` コマンドの出力に端末ウィンドウのバッファを示すフラグが付くようになりました

## Vimに関する脆弱性

特筆すべき脆弱性の報告はありませんでした。

## リリース情報

- [8.0.1554 : custom plugins loaded with --clean](https://github.com/vim/vim/commit/072687032683b1994d25a114893d9a6f8bc36612)
- [8.0.1555 : build error for some combination of features](https://github.com/vim/vim/commit/77780b66f4c5573df999a7319d77a852ea4ca837)
- [8.0.1556 : may not parse the t&#x5f;RS response correctly](https://github.com/vim/vim/commit/590ec878a52b3b3d4453475f1eb4899f2b37969f)
- [8.0.1557 : printf() does not work with only one argument](https://github.com/vim/vim/commit/c71807db9c1821baf86796cd76952df36ff1a29a)
- [8.0.1558 : no right-click menu in a terminal](https://github.com/vim/vim/commit/aef8c3da2ba59285b7cfde559ae21cdce6ba6919)
- [8.0.1559 : build failure without GUI](https://github.com/vim/vim/commit/658a154cbf2362c49a9797d9b2b9b69f481cd940)
- [8.0.1560 : build failure without GUI on MS-Windows](https://github.com/vim/vim/commit/3767b61ad99fe6ff2cace2dccb5d35331786735f)
- [8.0.1561 : crash with rust syntax highligting](https://github.com/vim/vim/commit/069dafc1ded60d9ee0fee4bcecce78ac8a235d87)
- [8.0.1562 : the terminal debugger can't set a breakpoint with the mouse](https://github.com/vim/vim/commit/71137fed4d77e985d49ca32c79f030512767b8ce)
- [8.0.1563 : timeout of getwinposx() can be too short](https://github.com/vim/vim/commit/3f54fd319f6641b4bed478bcc90cdb39ede68e31)
- [8.0.1564 : too many #ifdefs](https://github.com/vim/vim/commit/f2bd8ef2b4507d02c6043affff8f7e85e3414d5f)
- [8.0.1565 : can't build Mac version without GUI](https://github.com/vim/vim/commit/107279c17b3a21899e0a9d0293ada040216539ba)
- [8.0.1566 : too many #ifdefs](https://github.com/vim/vim/commit/8a3bb56230d220b8e1b4dcca517ed95f5341b8c7)
- [8.0.1567 : cannot build Win32 GUI without IME](https://github.com/vim/vim/commit/6e35a11490f1f9061677671220ce4cb73176b44a)
- [8.0.1568 : can't build on older Mac, header file is missing](https://github.com/vim/vim/commit/24fe4758941a5a59669f135e45ea7ec69d16fe9f)
- [8.0.1569 : warning for uninitialized variable from gcc](https://github.com/vim/vim/commit/28ada699c13833219acaeb7e768acc2acae50e02)
- [8.0.1570 : can't use :popup for a menu in the terminal](https://github.com/vim/vim/commit/29a2c08d792e4458a0af8371f5341394829fce29)
- [8.0.1571 : can't build without GUI](https://github.com/vim/vim/commit/40d77b0ec14a4999d8f0d52643a0ad289bf88dfa)
- [8.0.1572 : Mac: getting memory size doesn't work everywhere](https://github.com/vim/vim/commit/362dc33835e3ec5d41b7f29d8877da0a9e74c15c)
- [8.0.1573 : getwinpos(1) may cause response to be handled as command](https://github.com/vim/vim/commit/89894aa671ed1db03d95d38cab300702c242239d)
- [8.0.1574 : show cursor in wrong place when using popup menu](https://github.com/vim/vim/commit/987723e084660290270b3c3d943eb13bd828d5da)
- [8.0.1575 : crash when using virtual replace](https://github.com/vim/vim/commit/63e82db6fc910b2d8f1cd018894e50e8b4448155)
- [8.0.1576 : Perl VIM::Buffers() does not find every buffer](https://github.com/vim/vim/commit/16896a10199572ff05149786868ff402aa8e31ea)
- [8.0.1577 : virtual replace test fails on MS-Windows](https://github.com/vim/vim/commit/e7808481507b9e11ae73c8f865e95eb2d20f6cc8)
- [8.0.1578 : no test for :popup in terminal](https://github.com/vim/vim/commit/69f5a3011d5a93e0d54837e97d625bd6c253d897)
- [8.0.1579 : virtual replace test fails in GUI](https://github.com/vim/vim/commit/df0d24b62742edd3ea73795b96a771501e642970)
- [8.0.1580 : FEAT&#x5f;CURSORBIND and FEAT&#x5f;SCROLLBIND are unused](https://github.com/vim/vim/commit/946acdac5b4e334f816e78f7876bde9d97b3d96a)
- [8.0.1581 : cannot build Win32 GUI without +eval](https://github.com/vim/vim/commit/3b3a9a5609df2cbaef3512032ac47c1779fbc775)
- [8.0.1582 : in the MS-Windows console mouse movement is not used](https://github.com/vim/vim/commit/157d813be41c122e210b0eb1cd1e862fbddbf665)
- [8.0.1583 : using C99 comment](https://github.com/vim/vim/commit/fc6f16b57ce988cea6d07b5f2d5054711aad6bbf)
- [8.0.1584 : using C99 in Mac file gives compiler warning messages](https://github.com/vim/vim/commit/f536bf6d4518e2c66e635e39820d53f672021d87)
- [8.0.1585 : enabling beval&#x5f;term feature in Win32 GUI](https://github.com/vim/vim/commit/ebf142a1ed1824ed17a3eb4f64e75616e3b3d0e8)
- [8.0.1586 : imactivatefunc does not work on non-GUI Mac ](https://github.com/vim/vim/commit/201dc67db587e77a845fd24681352d61e57156e2)
- [8.0.1587 : inserting from the clipboard doesn't work literally](https://github.com/vim/vim/commit/3324d0a86421a634572758dcfde917547f4d4c67)
- [8.0.1588 : popup menu hangs after typing CTRL-C](https://github.com/vim/vim/commit/52f18a112a073c39187cd93f26115d1e2bfd29ce)
- [8.0.1589 : error for setting 'modifiable' when resetting it](https://github.com/vim/vim/commit/d7db27bafd1045281c965d9483539748a744de70)
- [8.0.1590 : padding in list type wastes memory](https://github.com/vim/vim/commit/1a840240376f2858d489736f9eed6d2975225fdf)
- [8.0.1591 : MS-Windows: when reparsing the arguments 'wildignore' matters](https://github.com/vim/vim/commit/20586cb4f4d516a60b96cc02a94b810fea8b8cdb)
- [8.0.1592 : terminal windows in a session are not properly restored](https://github.com/vim/vim/commit/4d8bac8bf593ff087517ff79090c2d224325aae6)
- [8.0.1593 : :qall never exits with an active terminal window](https://github.com/vim/vim/commit/25cdd9c33b21ddbd31321c075873bb225450d2d2)
- [8.0.1594 : :confirm qall not tested with active terminal window](https://github.com/vim/vim/commit/435acdb88c849c6bac44ef30f24f29618b36af50)
- [8.0.1595 : no autocommand triggered before exiting](https://github.com/vim/vim/commit/12a96de430779b88795fac87a2be666d9f661d1e)
- [8.0.1596 : no autocommand specifically for opening a terminal window](https://github.com/vim/vim/commit/b852c3e64d319d6ec47dd780c8654ae095e1d8c2)
- [8.0.1597 : autocommand events are not sorted](https://github.com/vim/vim/commit/e87303af3236b8fb5e1e3be4d0e2209344fbf8b2)
- [8.0.1598 : cannot select text in a terminal with the mouse](https://github.com/vim/vim/commit/c48369c3fc507f398abbc933a60f653c6abe6701)
- [8.0.1599 : no error message when gdb does not support debugger](https://github.com/vim/vim/commit/3e4b84d0b55936ac017d20df6651ddd05e38f58f)
- [8.0.1600 : crash when setting t&#x5f;Co to zero when 'termguicolors' is set](https://github.com/vim/vim/commit/f708ac592f47100a36f2bc12ec98ea6357fdfa27)
- [8.0.1601 : highlight test fails on Win32](https://github.com/vim/vim/commit/ff1e8795772a0175017c4c4f74ce33614ea8e73a)
- [8.0.1602 : crash in parsing JSON](https://github.com/vim/vim/commit/625f0c1eb75da08229843fa393b1ee4e6547d285)
- [8.0.1603 : cannot build with +terminal but without +menu](https://github.com/vim/vim/commit/f118d4847eb0dc2473260aebce301643568c5bc3)
- [8.0.1604 : paste test may fail if $DISPLAY is not set](https://github.com/vim/vim/commit/a903472cfa048e7a29d9c6ed0945ef03e48c6a08)
- [8.0.1605 : terminal test is a bit flaky](https://github.com/vim/vim/commit/012eb6629337fdf8afca78a24faa132e9b42e7b4)
- [8.0.1606 : singular/plural variants not translated](https://github.com/vim/vim/commit/228de1dfd2687af16dc7bc88942c65b5af9d7e04)
- [8.0.1607 : --clean loads user settings from .gvimrc](https://github.com/vim/vim/commit/62dd452d02f4a17f7e9fdfb3656b71bea62e12fe)
- [8.0.1608 : Win32: directx not enabled by default](https://github.com/vim/vim/commit/43cb6262141c0a0853680bd8d0433e04db3bf182)
- [8.0.1609 : shell commands in the GUI use a dumb terminal](https://github.com/vim/vim/commit/135682517bc378cfdb63fe3a6e3553935f69f6ce)
- [8.0.1610 : cannot build without GUI](https://github.com/vim/vim/commit/4ac31eeff0409de0b1fe2d9e9f678859656ac4ae)
- [8.0.1611 : CTRL-W in system terminal does not go to job](https://github.com/vim/vim/commit/af23bad0fd05b1f6495287089387a1f97efa5b49)
- [8.0.1612 : need to close terminal after shell stopped](https://github.com/vim/vim/commit/1dd98334d6daee8abefcd640291d4b777d9f0f96)
- [8.0.1613 : warning for unused variable in tiny build](https://github.com/vim/vim/commit/a2150ac01643ef99b0e7b2273c6ad84ce79c8f90)
- [8.0.1614 : "make tags" doesn't include libvterm](https://github.com/vim/vim/commit/9ef2a30e6ff5830984b57dd3d682656789a7ad78)
- [8.0.1615 : term&#x5f;dumpload() does not use the right colors](https://github.com/vim/vim/commit/52acb110ac025b57dad5f5ec6004abbe4726a718)
- [8.0.1616 : Win32: shell commands in the GUI open a new console](https://github.com/vim/vim/commit/f05fa378229004fb64493b411b864709ce7da453)
- [8.0.1617 : Win32: :shell command in the GUI crashes](https://github.com/vim/vim/commit/42f652f733550a0d8bc9b030c9a5a62e7e2d8668)
- [8.0.1618 : color Grey50 is missing in the compiled-in table](https://github.com/vim/vim/commit/ecadf4377f120200252e023ef6a9aa369e98850a)
- [8.0.1619 : Win32 GUI: crash when winpty is not installed](https://github.com/vim/vim/commit/81c3c89a28859de3aa167c53355f374394ada093)
- [8.0.1620 : reading spell file has no good EOF detection](https://github.com/vim/vim/commit/e26e0d2b83c2875b9829b884c2ababf8ca771f7e)
- [8.0.1621 : using invalid default value for highlight attribute](https://github.com/vim/vim/commit/6185903e3d07eb53326fc1403fc2de97ca31b775)
- [8.0.1622 : possible NULL pointer dereference](https://github.com/vim/vim/commit/6ed86ad170b60517eeddb54c2b22fdc888a22c0b)
- [8.0.1623 : terminal kill tests are flaky](https://github.com/vim/vim/commit/3e8d385347b23b2925d1b8ca64b78764d37f21fe)
- [8.0.1624 : options for term&#x5f;dumpdiff()/term&#x5f;dumpload() not implemented](https://github.com/vim/vim/commit/5a3a49ed592e8300553590392721a5e55fede6e7)
- [8.0.1625 : test&#x5f;quotestar is flaky when run in GTK GUI](https://github.com/vim/vim/commit/29dfa5af3c14406573d08609d2e95add4d6b45fb)
- [8.0.1626 : compiler warning for possible loss of data](https://github.com/vim/vim/commit/b571c63d481b065dc22465300e41a4b8cff66120)
- [8.0.1627 : compiler warning for visibility attribute not supported](https://github.com/vim/vim/commit/7735dafb584c77cf1191eceaa0b4db46266c9cb8)
- [8.0.1628 : channel log doesn't mention exiting](https://github.com/vim/vim/commit/0bd052ba12f672b46415ed1d5d1d2597a326f85a)
- [8.0.1629 : Mac: getpagesize() is deprecated](https://github.com/vim/vim/commit/62b7f6a139a19eb6c79eb428c66a7118e9741b5d)
- [8.0.1630 : trimming white space is not that easy](https://github.com/vim/vim/commit/295ac5ab5e840af6051bed5ec9d9acc3c73445de)
- [8.0.1631 : testing with Vim running in terminal is a bit flaky](https://github.com/vim/vim/commit/948a796bfea55d1e5a3a079d4976fdd4d093cd05)
- [8.0.1632 : in a terminal dump NUL and space are different](https://github.com/vim/vim/commit/47015b80a0b0ff74ba7cb597d5959604b6e9a511)
- [8.0.1633 : a TextChanged autocmd triggers when it is defined](https://github.com/vim/vim/commit/8c64a36e40b8746404f7151abe6849393396af10)
- [8.0.1634 : the ex&#x5f;vimgrep() function is too long](https://github.com/vim/vim/commit/75b0a888e41bcda4163072f41bb7b5471fef7651)
- [8.0.1635 : undefining &#x5f;POSIX&#x5f;THREADS causes problems with Python 3](https://github.com/vim/vim/commit/16d7eced1a08565a9837db8067c7b9db5ed68854)
- [8.0.1636 : no test for term&#x5f;dumpload() and term&#x5f;dumpdiff()](https://github.com/vim/vim/commit/45d2a64f964238753403017b6d0ef30be88ccf37)
- [8.0.1637 : no test for term&#x5f;dumpdiff() options argument](https://github.com/vim/vim/commit/897e63c4c0623078b5b39e2715b10f456a2453e7)
- [8.0.1638 : popup test fails depending on environment variable](https://github.com/vim/vim/commit/e7499ddc33508d3d341e96f84a0e7b95b2d6927c)
- [8.0.1639 : libvterm code lags behind master](https://github.com/vim/vim/commit/b5b49a3b430ea0aaf8cce6b7bc0e444f0211ddfb)
- [8.0.1640 : Test&#x5f;cwd() is flaky](https://github.com/vim/vim/commit/65873846e088bb94028e9d591ea03e377cb206b5)
- [8.0.1641 : job in terminal can't communicate with Vim](https://github.com/vim/vim/commit/8fbaeb195d9298c3a2a80300b5f96f1adddd2f59)
- [8.0.1642 : running Vim in terminal fails with two windows](https://github.com/vim/vim/commit/15a1c3fdc19499c267e4dc664aa0de9d65a08cad)
- [8.0.1643 : terminal API tests fail](https://github.com/vim/vim/commit/2de50f87622cd4e631fd17845e16a94ed5992b80)
- [8.0.1644 : terminal API tests still fail](https://github.com/vim/vim/commit/cf67a509e93167f14c892301e13de14636cedc61)
- [8.0.1645 : test for terminal response to escape sequence may fail](https://github.com/vim/vim/commit/086eb876951aa0b0bee3de4ed0569bf4b70c6f31)
- [8.0.1646 : MS-Windows: executable contains unreferenced functions](https://github.com/vim/vim/commit/4368d5ce8a2fa16826d76eeb32347bbbb4fad2bc)
- [8.0.1647 : terminal API may call any user function](https://github.com/vim/vim/commit/2a77d21f7893ba14e682a3c5891d606f117a3f36)
- [8.0.1648 : resource fork tool doesn't work on Python 3](https://github.com/vim/vim/commit/1b2f61e732a961c1345bf3bb6826c1caa870c10d)
- [8.0.1649 : no completion for argument list commands](https://github.com/vim/vim/commit/cd43effecab02c6c28b1c4a3a14f91b8c3f26c0d)
- [8.0.1650 : too many #ifdefs](https://github.com/vim/vim/commit/0c72fe4ed8430db41f43c5878e6ee60265dc49e9)
- [8.0.1651 : cannot filter :ls output for terminal buffers](https://github.com/vim/vim/commit/0751f51a5b428805a8c1e9fe529693d032bec991)
- [8.0.1652 : term&#x5f;dumpwrite() does not output composing characters](https://github.com/vim/vim/commit/f06b0b6c8f85ea9c320f2be30b25ed084969c1e2)
- [8.0.1653 : screen dump is made too soon](https://github.com/vim/vim/commit/1834d37396e046ccbc4aa2678ba16a38197da6b4)
- [8.0.1654 : warnings for conversion of void to function pointer](https://github.com/vim/vim/commit/7b24ce08fe99345cac035215fca29c7e174a6456)
- [8.0.1655 : outdated gdb message in terminal debugger unclear](https://github.com/vim/vim/commit/f3ba14ffd32faa1856a99cf657c426bf9d1204ae)

## 新着スクリプト

- [copy paste finder : utility for viewing copied-pasted code side-by-side](https://www.vim.org/scripts/script.php?script_id=5664)
- [dadbod.vim : Modern database interface for Vim](https://www.vim.org/scripts/script.php?script_id=5665)
- [open-browser-unicode.vim : Open fileformat.info page about character on current cursor / given character](https://www.vim.org/scripts/script.php?script_id=5667)
- [open-browser-github.vim : Open GitHub URL of file, issue, pull request (GitHub Enterprise supported)](https://www.vim.org/scripts/script.php?script_id=5668)
- [nimble-anyjump.vim : Initial upload.](https://www.vim.org/scripts/script.php?script_id=5669)
- [simplebuffer : a very simple buffer manager.](https://www.vim.org/scripts/script.php?script_id=5670)
- [vinegar.vim : Combine with netrw to create a delicious salad dressing](https://www.vim.org/scripts/script.php?script_id=5671)
- [rested : A desert based scheme with more colors.](https://www.vim.org/scripts/script.php?script_id=5672)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](https://www.vim.org/scripts/script.php?script_id=273) (1000)
2. [The NERD tree : A tree explorer plugin for navigating the filesystem](https://www.vim.org/scripts/script.php?script_id=1658) (732)
3. [wombat256.vim : Wombat for 256 color xterms](https://www.vim.org/scripts/script.php?script_id=2465) (717)
4. [python.vim : Enhanced version of the python syntax highlighting script](https://www.vim.org/scripts/script.php?script_id=790) (445)
5. [OmniCppComplete : C/C++ omni-completion with ctags database](https://www.vim.org/scripts/script.php?script_id=1520) (348)
6. [c.vim : C/C++ IDE -- Write and run programs. Insert statements, idioms, comments etc.](https://www.vim.org/scripts/script.php?script_id=213) (347)
7. [nginx.vim : initial version](https://www.vim.org/scripts/script.php?script_id=1886) (338)
8. [AutoComplPop : Automatically opens popup menu for completions](https://www.vim.org/scripts/script.php?script_id=1879) (323)
9. [jellybeans.vim : Colorful, dark color scheme](https://www.vim.org/scripts/script.php?script_id=2555) (317)
10. [yaml.vim : Syntax coloring and functions for YAML](https://www.vim.org/scripts/script.php?script_id=739) (275)

## vim-jp issues

Open : 262 (+1) | Closed : 899 (+6)

- [Issue #1155 : マルチバイト文字があるパス上のドット付きディレクトリ内のファイルが読み込めない](https://github.com/vim-jp/issues/issues/1155)
- [Issue #1156 : :tjump で一覧から選択してバッファに戻ったときにカーソルアイコンが戻らない](https://github.com/vim-jp/issues/issues/1156)
- [Issue #1157 : termguicolors有効時にt&#x5f;Coの値がおかしいと落ちる](https://github.com/vim-jp/issues/issues/1157)
- [Issue #1158 : pty付きで起動したジョブが即座に終了してしまうことがある](https://github.com/vim-jp/issues/issues/1158)
- [Issue #1159 : ":autocmd TextChanged &#x2a; FOO" 自体が TextChanged イベントを発生させる](https://github.com/vim-jp/issues/issues/1159)
- [Issue #1160 : インサートモードでの変更時に TextChanged (not TextChangedI) が発生する](https://github.com/vim-jp/issues/issues/1160)
- [Issue #1161 : macOS で job&#x5f;start を使い一定サイズ以上の文字列を送信すると SEGV する](https://github.com/vim-jp/issues/issues/1161)

