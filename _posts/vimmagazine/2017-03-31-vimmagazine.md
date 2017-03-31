---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2017 年 3 月号

---

## 話題

## 今月の新機能

今月の新機能及びユーザーに影響のある変更は以下のとおりです。

*   8.0.0392: テスト実行時に特定のエラーを無視するための `test_ignore_error()` 関数が追加されました
*   8.0.0393: タグサーチにおいて同じタグが複数見つかった場合に、タグファイル中での順番で候補を表示するようになりました
*   8.0.0394: 水平スクロール時のタブ文字が正しく整列するようになりました
*   8.0.0396: バルーンの内容を非同的に更新するために `balloon_show()` 関数が追加されました
*   8.0.0402: `:map` 入力中の補完候補に `<special>` が追加されました
*   8.0.0405: `v:progpath` がフルパスに変更されました
*   8.0.0411: メニュー翻訳である `menutranslate` が `{english}` の大文字小文字を無視するようになりました
*   8.0.0420: `:make` 等のコマンドの出力の文字コードを指定する `'makeencoding'` オプションが追加されました
*   8.0.0431: `'cinoptions'` オプションに C 言語の extern ブロックのインデントを調整するための `E` フラグが追加されました
*   8.0.0439: 引数リストが空の時の `:%argdel` がエラーにならなくなりました
*   8.0.0475: リモートサーバを開始するための関数 `remote_startserver()` が追加されました
*   8.0.0492: リモートサーバを操作する関数 `remote_expr()` 及び `remote_read()` にタイムアウトを指定できるようになりました
*   8.0.0499: `taglist()` 関数に特定のバッファを指定することで、そのバッファのタグを優先できるようになりました
*   8.0.0501: Windows の `:!start` コマンドがコマンドプロンプトの `start` コマンドに近い動作をするようになりました
*   8.0.0504: Ex コマンドの探索効率が改善されました。これにより Vim script の実行速度が若干速くなります
*   8.0.0517: `setqflist()` と `setloclist()` 関数にリストを削除するための `'f'` フラグが追加されました
*   8.0.0522: Windows において `'clipboard'` が `"unnanmed"` の際に、 `:global` コマンドと組み合わせると機能していなかった `yyp` が修正されました
*   8.0.0523: `dv}` がマルチバイト文字を破壊する問題が修正されました
*   ランタイムの更新
    *   ターミナルマルチプレクサ tmux の設定ファイル用のファイルタイプ tmux が追加されました
    *   プログラミング言語 Rust 用のファイルタイプ rust が追加されました

## Vimに関する脆弱性

特筆すべき脆弱性の報告はありませんでした。

## リリース情報

- [8.0.0390 : when the window scrolls the popup menu may be garbled](https://github.com/vim/vim/commit/cb03642399ae4f3c47a006f608a6b912f9d98b5f)
- [8.0.0391 : arabic support is verbose and not well tested](https://github.com/vim/vim/commit/5f53dd3f747711be90879fa2f22a207970b86750)
- [8.0.0392 : GUI test fails with Athena and Motif](https://github.com/vim/vim/commit/e0c31f6a304496b597a51ce98af419ec815bec74)
- [8.0.0393 : (after 8.0.0190) order of duplicate tags is not preserved](https://github.com/vim/vim/commit/98e83b295628bc29bc67bcc1adb8ae75d01b8e07)
- [8.0.0394 : tabs are not aligned when scrolling horizontally](https://github.com/vim/vim/commit/abc39ab642791ae3d22a524516eeedb673a95d9d)
- [8.0.0395 : testing the + register fails with Motif](https://github.com/vim/vim/commit/f8ab1b14fd972a093e0c12146dd3becd511eb519)
- [8.0.0396 : 'balloonexpr' only works synchronously](https://github.com/vim/vim/commit/59716a27bd5c6c64def6c3ca430ff1246deae749)
- [8.0.0397 : (after 8.0.0392) can't build with +viminfo but without +eval](https://github.com/vim/vim/commit/5a66dfb25eb478c26176d993393a3b1b124edb39)
- [8.0.0398 : illegal memory access with "t"](https://github.com/vim/vim/commit/66727e16079fbac6db3897b5c3736ec9fba995bb)
- [8.0.0399 : crash when using balloon&#x5f;show() when not supported](https://github.com/vim/vim/commit/caf6434ac937cf26050276d7b474be2d2d6a06b3)
- [8.0.0400 : some tests have a one second delay](https://github.com/vim/vim/commit/358f6b0a36d25e0e2f8f1fc625974816b71fa50a)
- [8.0.0401 : test fails with missing balloon feature](https://github.com/vim/vim/commit/a0107bdf8762d81cb49909ef02ffff9954092fb6)
- [8.0.0402 : :map completion does not have \<special>](https://github.com/vim/vim/commit/cf5fdf7d1689ecb145b634dcb9c6e9fc60f63869)
- [8.0.0403 : GUI tests may fail](https://github.com/vim/vim/commit/24d7636e98b33ec6873c51fc72f6bf4b10188d8b)
- [8.0.0404 : not enough testing for quickfix](https://github.com/vim/vim/commit/391b1dd040af204b150d43c5a1c97477ee450a28)
- [8.0.0405 : v:progpath may become invalid after :cd](https://github.com/vim/vim/commit/08cab9608781c975b4acbad875862b842b29258d)
- [8.0.0406 : arabic shaping code is verbose](https://github.com/vim/vim/commit/7f73b54631af3f0e6f0acd1a1b4c9e8436784705)
- [8.0.0407 : filtering folds with marker method not tested](https://github.com/vim/vim/commit/3f3897e41f77a7bcd44ffd37acd4d524c0169ed4)
- [8.0.0408 : updating folds does not always work properly](https://github.com/vim/vim/commit/eadbc2b46176e2aff2cde5f3874b734ae2ae082a)
- [8.0.0409 : set&#x5f;progpath is defined but not always used](https://github.com/vim/vim/commit/595297d6d5658266dd0574361f77f7832bb18e26)
- [8.0.0410 : newer gettext/iconv library has extra dll file](https://github.com/vim/vim/commit/a54d2fc0d3ec7c4c8ca99b91e2ff7f7fee45c8bd)
- [8.0.0411 : menu translations don't match when case is changed](https://github.com/vim/vim/commit/11dd8c1201033dd74e2ea665ba277425b4b965b0)
- [8.0.0412 : (after 8.0.0411) menu test fails on MS-Windows](https://github.com/vim/vim/commit/5558d19432120696409c007c64d5ba52eed42670)
- [8.0.0413 : (after 8.0.0412) menu test fails on MS-Windows using gvim](https://github.com/vim/vim/commit/a1c8ecfda90c0e0e519762ae0521d7f6e297c32e)
- [8.0.0414 : balloon eval is not tested](https://github.com/vim/vim/commit/d5841f28d4b041830af0f3314979f9b9093d1a77)
- [8.0.0415 : (after 8.0.0414) balloon test fails on MS-Windows](https://github.com/vim/vim/commit/0f9ea22c115084a982c86214ef41200e0dfc1ea6)
- [8.0.0416 : setting v:progpath is not quite right](https://github.com/vim/vim/commit/436631969716e056feca47480976100b70283370)
- [8.0.0417 : test for the clipboard fails sometimes](https://github.com/vim/vim/commit/0fbff646d5beb1f48c9be3d0ce0339be90fd49f5)
- [8.0.0418 : ASAN logs are disabled](https://github.com/vim/vim/commit/a382868115e8f8e44c6a85bb9587f8eb31fa0033)
- [8.0.0419 : test for v:progpath fails on MS-Windows](https://github.com/vim/vim/commit/066029ef92b45dc4bd21a353b1fac25841062c26)
- [8.0.0420 : text garbled when the system encoding differs from 'encoding'](https://github.com/vim/vim/commit/2c7292dc5bbf155fe2192d417363b8c085759cad)
- [8.0.0421 : diff mode wrong when adding line at end of buffer](https://github.com/vim/vim/commit/f58a8475e17bd566760fc7e2a17d35ddf4edacf2)
- [8.0.0422 : Python test fails with Python 3.6](https://github.com/vim/vim/commit/f411a3c5aff644ec7a999be83ef22adf3e1132dd)
- [8.0.0423 : changing 'cinoptions' does not always work](https://github.com/vim/vim/commit/6b64394f346594404cffb9591d71ac693040679f)
- [8.0.0424 : compiler warnings on MS-Windows](https://github.com/vim/vim/commit/f7e894dfb60b6fd8f426aa3da1452803f8de010c)
- [8.0.0425 : build errors when building without folding](https://github.com/vim/vim/commit/a6c07603ca2fa98a4184f5fd5ab021adcf5156b2)
- [8.0.0426 : insufficient testing for statusline](https://github.com/vim/vim/commit/300af82eca3dd5fec6be5516f16c9fd344ae104b)
- [8.0.0427 : 'makeencoding' missing from the options window](https://github.com/vim/vim/commit/ad4187e6fc9c8e1083a172852d958a70a689a75c)
- [8.0.0428 : git and hg see new files after running tests](https://github.com/vim/vim/commit/65408f7dfbd4bb50d740a8b8f2171080571223f2)
- [8.0.0429 : (after 8.0.0428) options test does not always test everything](https://github.com/vim/vim/commit/e8512d7fedbddd8d464b2c589e6ba5ab9a41bc78)
- [8.0.0430 : (after 8.0.0429) options test fails or hangs on MS-Windows](https://github.com/vim/vim/commit/bb962262eb6b18d93161d7506317d215ae80bfe5)
- [8.0.0431 : 'cinoptions' cannot set indent for extern block](https://github.com/vim/vim/commit/7720ba8599162fbbb8f7fc034f674a2ccd3ca7f1)
- [8.0.0432 : "make shadow" creates an invalid link](https://github.com/vim/vim/commit/12d1dc9155a507f2bbffae19696c33a640fbc8d1)
- [8.0.0433 : beeps when running tests](https://github.com/vim/vim/commit/c3c766ea8c35f5b2bd45fb3d74d0ae46b2d8c24f)
- [8.0.0434 : clang version not correctly detected](https://github.com/vim/vim/commit/5f69fee26e0688c1646c8724ab794cb0c8d7c44d)
- [8.0.0435 : some functions are not tested](https://github.com/vim/vim/commit/41042f3cfdb91f946e553456278a995e61dd8578)
- [8.0.0436 : running the options test sometimes resizes the terminal](https://github.com/vim/vim/commit/ee1dd1cdc1f9a07e53b792e8762f515213b10c8a)
- [8.0.0437 : packadd test does not fully work](https://github.com/vim/vim/commit/644df41c44cbdfacdedbba55ef77a6c6031eccd8)
- [8.0.0438 : the fnamemodify test may cause later tests to fail](https://github.com/vim/vim/commit/056f700031602a2734b1ddf45f6bc2817e49b996)
- [8.0.0439 : ":%argdel" gives an error for an empty arglist](https://github.com/vim/vim/commit/69a92fb5aecdf2f9d5f6947790b18991b22d0e4c)
- [8.0.0440 : not enough test coverage in Insert mode](https://github.com/vim/vim/commit/eb992cb90fd79c77ad2743459ac898e6ac3de939)
- [8.0.0441 : dead code in #ifdef](https://github.com/vim/vim/commit/ad2cfb54b78523435b31de7234081ff70a53e736)
- [8.0.0442 : patch shell command not well escaped](https://github.com/vim/vim/commit/1ef73e33c9414eb02c229d8234aafd9d481a8856)
- [8.0.0443 : terminal width is set to 80 in test3](https://github.com/vim/vim/commit/38a3d6c9601b637a28f399059263300e9f65eba4)
- [8.0.0444 : (after 8.0.0442) diffpatch fails when file name has a quote](https://github.com/vim/vim/commit/a95ab321200f0239991bf53756b17cd7b90745f9)
- [8.0.0445 : getpgid is not supported on all systems](https://github.com/vim/vim/commit/2fcf6688bc3a8df2dff9c352d415b89db8b33668)
- [8.0.0446 : the ";" command does not work after some characters](https://github.com/vim/vim/commit/454709baffd3205bf2b7d2519419675a122f2bd2)
- [8.0.0447 : getting font name does not work on X11](https://github.com/vim/vim/commit/8774845ce1a7def122ea07c057a79417f3be3d17)
- [8.0.0448 : some macros are in lower case](https://github.com/vim/vim/commit/b5aedf3e228d35821591da9ae8501b61cf2e264c)
- [8.0.0449 : (after 8.0.0448) part of fold patch accidentally included](https://github.com/vim/vim/commit/abb4d5f73fd97d8f3c3e32311959929820e74b54)
- [8.0.0450 : v:progpath is not reliably set](https://github.com/vim/vim/commit/47ffb905f363571072faefbd417938e7ff351a9f)
- [8.0.0451 : some macros are in lower case](https://github.com/vim/vim/commit/91acfffc1e6c0d8c2abfb186a0e79a5bf19c3f3f)
- [8.0.0452 : some macros are in lower case](https://github.com/vim/vim/commit/1c46544412382db8b3203d6c78e550df885540bd)
- [8.0.0453 : adding fold marker creates new comment](https://github.com/vim/vim/commit/025a6b708a9bff54c73fb9c641b980da19e943a9)
- [8.0.0454 : compiler warnings for "always true" comparison](https://github.com/vim/vim/commit/977d03733601948ab334fec24d0da82bc18c57c3)
- [8.0.0455 : the mode test may hang](https://github.com/vim/vim/commit/ffea8c99d9658b0b51a848a6f674851851e78fa7)
- [8.0.0456 : typo in MinGW test makefile](https://github.com/vim/vim/commit/84be8b66604ef28c0e249284da3c6f0cab1c25ae)
- [8.0.0457 : using :move messes up manual folds](https://github.com/vim/vim/commit/88d298aed8682eac872ebfe40df3112a6acd83e8)
- [8.0.0458 : potential crash if adding list or dict to dict fails](https://github.com/vim/vim/commit/42f45b850df93b806ef44f3025aecedfffb9ae01)
- [8.0.0459 : (after 8.0.0457) old folding fix for :move no longer needed](https://github.com/vim/vim/commit/fe70c5198336dba7e8b0fa915e6c4350e7f80f6f)
- [8.0.0460 : can't build on HPUX](https://github.com/vim/vim/commit/d8492792f26c5bbdefc8f452a393621f055edd8c)
- [8.0.0461 : (after 8.0.0457) test 45 hangs on MS-Windows](https://github.com/vim/vim/commit/4a5a8dd6e5fa5581bd24b584a71b5fce9ba34bb8)
- [8.0.0462 : failure of an MS-Windows test may go unnoticed](https://github.com/vim/vim/commit/b27523ff7ab3083c20c0f33415d890f978726c4f)
- [8.0.0463 : side effects from resetting 'compatible' in defaults.vim](https://github.com/vim/vim/commit/0f39a82b07b285e0c54162e1038348b2988a715c)
- [8.0.0464 : can't find executable name on Solaris and FreeBSD](https://github.com/vim/vim/commit/f3757f0c87bbd52c7989c85dcbd21511bffcbdd6)
- [8.0.0465 : off-by-one error in using :move with folding](https://github.com/vim/vim/commit/40ebc0afda8d8e478d2090133ed6a3cd3d8da3ec)
- [8.0.0466 : still macros that should be all-caps](https://github.com/vim/vim/commit/8820b48654b62472821d9b155fe03ab7ac13a05c)
- [8.0.0467 : using g\< after :for does not show the right output](https://github.com/vim/vim/commit/57002ad70c4c32f3afefec24994a974cf3eef3ad)
- [8.0.0468 : after aborting an Ex command g\< does not work](https://github.com/vim/vim/commit/f2405ed2321da4a879fe0b0703af780fc0432c63)
- [8.0.0469 : compiler warnings on MS-Windows](https://github.com/vim/vim/commit/0c0d4eca4dd6252f22ec39f2d561a5e8a68e9a4e)
- [8.0.0470 : not enough testing for help commands](https://github.com/vim/vim/commit/751ba616d1c47de2c273b269df06c36a7ed141a2)
- [8.0.0471 : exit callback test sometimes fails](https://github.com/vim/vim/commit/0529b3eb01fcfd18c0644f8ece9ea107dd460a0f)
- [8.0.0472 : when a test fails another test may also fail](https://github.com/vim/vim/commit/c537947100d4a73754ad4b4c9228c6c439655e5d)
- [8.0.0473 : no test covering arg&#x5f;all()](https://github.com/vim/vim/commit/8c34aa09a449a5c1c2d1141b6fafa90f29b3fc12)
- [8.0.0474 : the client-server feature is not tested](https://github.com/vim/vim/commit/15bf76d40be1f1622ff5cc16596c308e76e2ca94)
- [8.0.0475 : not enough testing for the client-server feature](https://github.com/vim/vim/commit/7416f3e73ab2c4e7ae3adc2ff6e70234f7d40d2e)
- [8.0.0476 : (after 8.0.0475) missing change to main.c](https://github.com/vim/vim/commit/7a43cb9cb5909634469b0a3e865747f10ea38e9b)
- [8.0.0477 : the client-server test may hang when failing](https://github.com/vim/vim/commit/42205551b140bee8b419b24abe210f56bb80b35e)
- [8.0.0478 : tests use assert&#x5f;true(0) and assert&#x5f;false(1) to report errors](https://github.com/vim/vim/commit/37175409d766ce67f2548dffa6d73451379b5737)
- [8.0.0479 : remote&#x5f;peek() is not tested](https://github.com/vim/vim/commit/6caf606b144bbe2bf3ac515a194f5ff18fdfd331)
- [8.0.0480 : the remote&#x5f;peek() test fails on MS-Windows](https://github.com/vim/vim/commit/15e737f768542fcc516296b5c158e14cc7ba7feb)
- [8.0.0481 : unnecessary if statement](https://github.com/vim/vim/commit/aab93b12cb54fbe5efe9e8f6fde1c46802a3031e)
- [8.0.0482 : the setbufvar() function may mess up the window layout](https://github.com/vim/vim/commit/2c90d51123fba44a90e09aa4a4f2b7d972dadb94)
- [8.0.0483 : illegal memory access when using :all](https://github.com/vim/vim/commit/f79225ed4f81bc579bb3360ad2eb06adc8058153)
- [8.0.0484 : :lhelpgrep does not fail after a successful one](https://github.com/vim/vim/commit/ee85df37634dfb0c40ae5de0b4f246aef460b392)
- [8.0.0485 : not all windows commands are tested](https://github.com/vim/vim/commit/4520d440c59034452d1450b27fcd56825c090687)
- [8.0.0486 : crash and endless loop when closing windows in autocmd](https://github.com/vim/vim/commit/8c752bd6c4af54c0b7bac35a39acc2bf16015f85)
- [8.0.0487 : the autocmd test hangs on MS-Windows](https://github.com/vim/vim/commit/651e4056aceaa580df9e2fff575a6402d5c6f0dc)
- [8.0.0488 : running tests leaves an "xxx" file behind](https://github.com/vim/vim/commit/b4c5572e749063daeef30cfc2932bea1fd468165)
- [8.0.0489 : clipboard and "&#x2a; register is not tested](https://github.com/vim/vim/commit/7dd4850698e37c3ed4a3c3b23f4cdbf1ae6542f8)
- [8.0.0490 : vertical split makes 'winfixwidth' window smaller](https://github.com/vim/vim/commit/38e3483637c16e018f88c07b1dcff97cdb821a29)
- [8.0.0491 : quotestar test fails when features are missing](https://github.com/vim/vim/commit/bfd830d3e2dbd1e9b14c65625f18773074e6ac67)
- [8.0.0492 : a failing client-server request can make Vim hang](https://github.com/vim/vim/commit/81b9d0bd5c705815e903e671e81b0b05828efd9c)
- [8.0.0493 : crash with cd command with very long argument](https://github.com/vim/vim/commit/15618fa643867cf0d9c31f327022a22dff78a0cf)
- [8.0.0494 : build failure with older compiler on MS-Windows](https://github.com/vim/vim/commit/1662ce104e47f6a1eb0c2004d86b995f74dc4652)
- [8.0.0495 : quotestar test uses timer instead of timeout](https://github.com/vim/vim/commit/f5610da7a8a5908e747d1ae39dbed1408bbd2644)
- [8.0.0496 : insufficient testing for folding](https://github.com/vim/vim/commit/518c9b133baed4b1a555d7be0ece1b48bb234b1d)
- [8.0.0497 : arabic support is not fully tested](https://github.com/vim/vim/commit/3ff2f09525e3461419886763efa1a01f08882d8c)
- [8.0.0498 : two autocmd tests are skipped on MS-Windows](https://github.com/vim/vim/commit/e94260f3582cf928cc59e792b3b6ede57794a4a6)
- [8.0.0499 : taglist() does not prioritize tags for a buffer](https://github.com/vim/vim/commit/c6aafbaf3ea755e3ab4ee2e3045911126a08b038)
- [8.0.0500 : quotestar test is still a bit flaky](https://github.com/vim/vim/commit/4889ad7c6cd46368f553cc1140ca17508e8beaaa)
- [8.0.0501 : on MS-Windows ":!start" does not work as expected](https://github.com/vim/vim/commit/b2964f2570574b4c66f3645d69956fec99f2af3e)
- [8.0.0502 : Coverity complains about possible NULL pointer](https://github.com/vim/vim/commit/a37ffaa5e0a47e2db27bc0cc23f49e7094f47f3b)
- [8.0.0503 : endless loop in updating folds with 32 bit ints](https://github.com/vim/vim/commit/9d20ce6970158de69202a82529d9e97827a9e71b)
- [8.0.0504 : looking up an Ex command is a bit slow](https://github.com/vim/vim/commit/e5e0fbcd4244d032a0635ad7defe2831f251c639)
- [8.0.0505 : failed window split for :stag not handled](https://github.com/vim/vim/commit/ba6ad17378ddb9b33412d85174224997b8ff7a4f)
- [8.0.0506 : can't build with ANSI C](https://github.com/vim/vim/commit/6c0c1e8052811a818739e2f3d543291b7a347ad0)
- [8.0.0507 : client-server tests fail when $DISPLAY is not set](https://github.com/vim/vim/commit/a2845b8f5a3058c8c89699771ffd4d69513b097d)
- [8.0.0508 : Coveralls no longer shows per-file coverage](https://github.com/vim/vim/commit/d722fd74d84c8e7119d9599b67388bd26a2dacf5)
- [8.0.0509 : no link to codecov.io results](https://github.com/vim/vim/commit/dd00851e07f251c7c2696d87303f3df70a458731)
- [8.0.0510 : (after 8.0.0509) typo in link to codecov.io results](https://github.com/vim/vim/commit/8a0141d4e7054587281d3bfd09480376b1fc719f)
- [8.0.0511 : message for skipping client-server tests is unclear](https://github.com/vim/vim/commit/a683ec44c34f0717dcc6a0c03493ba39b879ac38)
- [8.0.0512 : check for available characters takes too long](https://github.com/vim/vim/commit/1572e30607e9a3bee9750242bf37168f7d2b6e66)
- [8.0.0513 : getting name of cleared highlight group is wrong](https://github.com/vim/vim/commit/c96272e30e2b81e5e0c8418f09d9db4e2fcd5d73)
- [8.0.0514 : script for creating cmdidxs can be improved](https://github.com/vim/vim/commit/980128c369451450743bdb90a67588fa72ec4b07)
- [8.0.0515 : ml&#x5f;get errors in silent Ex mode](https://github.com/vim/vim/commit/d5d37537d1fa46fd468bd378af2006dd09840f38)
- [8.0.0516 : a large count on a normal command causes trouble](https://github.com/vim/vim/commit/b1e04fca3704e272a30afbe062498819eaacd4f9)
- [8.0.0517 : there is no way to remove quickfix lists](https://github.com/vim/vim/commit/b6fa30ccc39cdb7f1d07b99fe2f4c6b61671dac2)
- [8.0.0518 : bad fold text when a multi-byte char has a zero byte](https://github.com/vim/vim/commit/c6cd8409c2993b1476e123fba11cb4b8d743b896)
- [8.0.0519 : character classes are not well tested](https://github.com/vim/vim/commit/0c078fc7db2902d4ccba04506db082ddbef45a8c)
- [8.0.0520 : using a function pointer while the function is known](https://github.com/vim/vim/commit/ace95989ed81929a84e205b26d0972cb9d6b4b19)
- [8.0.0521 : GtkForm handling is outdated](https://github.com/vim/vim/commit/99a6e8dd824399332563caa6cacfcda33da1f366)
- [8.0.0522 : Win32: when 'clipboard' is "unnamed" yyp does not work](https://github.com/vim/vim/commit/3fcfa35f82b9d1ef5e95051b3a45578e10c14ec3)
- [8.0.0523 : dv} deletes part of a multi-byte character.](https://github.com/vim/vim/commit/bf3d58073f7b34b2d65d1d08a728d1164e03cceb)
- [8.0.0524 : (after 8.0.0518) folds messed up](https://github.com/vim/vim/commit/8da1e6cedf839902e15987a98733ebd31b5f1b81)
- [8.0.0525 : completion for user command argument not tested](https://github.com/vim/vim/commit/a33ddbbd04ca9b81cba6114708f42b8e26293b99)
- [8.0.0526 : Coverity complains about possible negative value](https://github.com/vim/vim/commit/85325f839a14212c7d517a4cb3464c347bfd0c1e)
- [8.0.0527 : leftover file from RISC OS](https://github.com/vim/vim/commit/ce5c2742012e4793cb4cf82791e849f817fecd95)
- [8.0.0528 : highlight wrong text when 'wim' includes "longest"](https://github.com/vim/vim/commit/ef8eb0897819099fb00d675afb9bffe1d008c45e)
- [8.0.0529 : line in test commented out](https://github.com/vim/vim/commit/13489b9c41922b9ecb16fa47564ec76641515c08)

## 新着スクリプト

- [visual-search : Search for or only within visual selection](http://www.vim.org/scripts/script.php?script_id=5533)
- [nonclobbering-shift : Visual shift without clobbering formatting](http://www.vim.org/scripts/script.php?script_id=5534)
- [bargreybars.vim : remove gray bars from maximised gvim windows, on Windows ](http://www.vim.org/scripts/script.php?script_id=5535)
- [Agrep : Asynchronous grep plugin for Vim](http://www.vim.org/scripts/script.php?script_id=5536)
- [vimloo : Write VimL Script in an Object Orient Style](http://www.vim.org/scripts/script.php?script_id=5537)
- [vnote : A dairy or note edit and manage tool in vim](http://www.vim.org/scripts/script.php?script_id=5538)
- [GVFilter.vim : VIm plugin to filter current buffer by given arguments, like :g or :v](http://www.vim.org/scripts/script.php?script_id=5539)
- [vim-fileHeader : add information about source file at head of file](http://www.vim.org/scripts/script.php?script_id=5540)
- [R4vim : Plugin to work with R scripts.](http://www.vim.org/scripts/script.php?script_id=5541)
- [note.vim : A syntax highlighting for general text notes](http://www.vim.org/scripts/script.php?script_id=5542)
- [eclipse-lite : a vim scheme appears to eclipse](http://www.vim.org/scripts/script.php?script_id=5543)
- [vim-sourcer : Several new commands for sourcing vimscript files](http://www.vim.org/scripts/script.php?script_id=5544)
- [CurtineIncSw.vim : Switch from &#x2a;.c&#x2a; to &#x2a;.h&#x2a; and vice versa](http://www.vim.org/scripts/script.php?script_id=5545)
- [goyo-doc : Vim](http://www.vim.org/scripts/script.php?script_id=5546)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (1435)
2. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (1150)
3. [wombat256.vim : Wombat for 256 color xterms](http://www.vim.org/scripts/script.php?script_id=2465) (628)
4. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (557)
5. [c.vim : C/C++ IDE -- Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (513)
6. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (464)
7. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (441)
8. [nginx.vim : initial version](http://www.vim.org/scripts/script.php?script_id=1886) (437)
9. [winmanager : A windows style IDE for Vim 6.0](http://www.vim.org/scripts/script.php?script_id=95) (390)
10. [pathogen.vim : Poor man's package manager. Easy manipulation of 'runtimepath' et al](http://www.vim.org/scripts/script.php?script_id=2332) (375)

## vim-jp issues

Open : 255 (-1) | Closed : 786 (+9)

- [Issue #1034 : 複数の'filetype'を扱いたい](https://github.com/vim-jp/issues/issues/1034)
- [Issue #1035 : Batch mode で ":sign place" すると SEGV](https://github.com/vim-jp/issues/issues/1035)
- [Issue #1036 : 新しい filetype である tmux の検出パターンが正しくない](https://github.com/vim-jp/issues/issues/1036)
- [Issue #1037 : :windo コマンドで ! が指定できてしまう](https://github.com/vim-jp/issues/issues/1037)
- [Issue #1038 : f最 の後に ; が効かない](https://github.com/vim-jp/issues/issues/1038)
- [Issue #1039 : did&#x5f;filetype()の動きがhelpと異なる](https://github.com/vim-jp/issues/issues/1039)
- [Issue #1040 : インサートモードの補完が遅い](https://github.com/vim-jp/issues/issues/1040)
- [Issue #1041 : インサートモードでCommand+Vを使ったペーストを行うと正常にペーストされない。](https://github.com/vim-jp/issues/issues/1041)
