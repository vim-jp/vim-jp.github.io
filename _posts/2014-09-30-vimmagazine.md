---
layout: post
category: vimmagazine
title: Vim Magazine 9 月号
---

## 話題

- [rogue.vim : Porting of Rogue-clone II for Vim](http://www.vim.org/scripts/script.php?script_id=5017)

  名作RPG『Rogue Clone II』が Vim に移植されました。

## 今月の新機能

  - 7.4.434: gettabvar() でスコープ変数が取れるようになりました。
  - 7.4.449: `:helpclose` コマンドが追加されました。
  - 7.4.456: 'backupcopy' オプションが global-local になりました。バッファごとに設定を変更できるようになります。

## リリース情報

- [7.4.431 : compiler warning](http://code.google.com/p/vim/source/detail?r=5672a06e4bd84cae4194ff0584d487448bad8c38)
- [7.4.432 : MS-Windows: setting 'encoding' does not convert arguments](http://code.google.com/p/vim/source/detail?r=59cd2b16f71846dec2e842491e509385beeda5e9)
- [7.4.433 : test 75 fails on MS-Windows](http://code.google.com/p/vim/source/detail?r=dd3dac42cb9b5f95ad16d5e08220b0c7adf4b024)
- [7.4.434 : gettabvar() is not consistent with getwinvar() and getbufvar()](http://code.google.com/p/vim/source/detail?r=4176c48c7dd51d0fcdb7403cbfb489a75b0d819d)
- [7.4.435 : line formatting behaves differently when 'linebreak' is set](http://code.google.com/p/vim/source/detail?r=b98af9af378eebe674052aa9a1b1956ad3f96487)
- [7.4.436 : ml&#x5f;get error for autocommand that moves the cursor](http://code.google.com/p/vim/source/detail?r=52fa8300ce205a8f3c45b3b1874564b7768630ea)
- [7.4.437 : new and old regexp engine are not consistent](http://code.google.com/p/vim/source/detail?r=8515b42f939cdf3a59fa6120e989d9f8fe6f4571)
- [7.4.438 : cached values for 'cino' not reset for ":set all&"](http://code.google.com/p/vim/source/detail?r=db3b8fe8330ea2afabd6c4856be5c76ef86d4728)
- [7.4.439 : duplicate message in message history](http://code.google.com/p/vim/source/detail?r=07b28e96af8bace1af65bac661f22716781103fd)
- [7.4.440 : omni complete popup drawn incorrectly](http://code.google.com/p/vim/source/detail?r=1f578cd9a65779d2597e0135a5916db621d65734)
- [7.4.441 : endless loop and other problems when 'cedit' is set to CTRL-C](http://code.google.com/p/vim/source/detail?r=e754b23b7d1c3825dc2c6028867d631520a8cdca)
- [7.4.442 : (after 7.4.434) using unitinialized variable](http://code.google.com/p/vim/source/detail?r=8c3c067b4ae3cc02ce4411042df05f97df5bc316)
- [7.4.443 : error reported by ubsan when running test 72](http://code.google.com/p/vim/source/detail?r=f1ba154c3a12d06cd0f40edbfc675103eb2dd7ad)
- [7.4.444 : reversed question mark not recognized as punctuation](http://code.google.com/p/vim/source/detail?r=d4cc9f48d5a40382e3e3d52dfe6e948745ee13b4)
- [7.4.445 : clipboard may be cleared on startup](http://code.google.com/p/vim/source/detail?r=1138726736fbdb92aa5259ca4b76573b8f4b3f79)
- [7.4.446 : environment for executing autocommand may not be restored.](http://code.google.com/p/vim/source/detail?r=63121fdd093ff71081725d5495337ea45d38daab)
- [7.4.447 : spell files from Hunspell may generate a lot of errors](http://code.google.com/p/vim/source/detail?r=0d2c821cdc2505879823a551baeeb8d3c61ce758)
- [7.4.448 : MS-Windows: Using ETO&#x5f;IGNORELANGUAGE causes problems](http://code.google.com/p/vim/source/detail?r=2fd96725b0632cd9183581ce13d7b2cbd734d8d4)
- [7.4.449 : can't easily close the help window](http://code.google.com/p/vim/source/detail?r=cb5480096f1b9e8f44e9d742fa190a90d77e1c54)
- [7.4.450 : some commands that edit another buffer don't support +cmd](http://code.google.com/p/vim/source/detail?r=7c9abc70ffc187f9bb60d968cb7e3cc5ed6d4231)
- [7.4.451 : calling system() with empty input gives an error](http://code.google.com/p/vim/source/detail?r=ac6cfdc02695b31b07e9beb60fd704f74f4d4c81)
- [7.4.452 : can't build with tiny features](http://code.google.com/p/vim/source/detail?r=78aad99db7627a19574d10415780ded69936d7d7)
- [7.4.453 : still can't build with tiny features](http://code.google.com/p/vim/source/detail?r=c72eb8499a9d3f618de3528287d1de62025fdda4)
- [7.4.454 : CTRL-W-\] does not work when using a Visual selection](http://code.google.com/p/vim/source/detail?r=0cdff7c268559f8f34eae073a013ece71b62b9e3)
- [7.4.455 : completion for :buf does not use 'wildignorecase'](http://code.google.com/p/vim/source/detail?r=ebc72764fa1e74758ae1ef9d3f49301f80fc3aa4)
- [7.4.456 : 'backupcopy' is global, can't set a local value](http://code.google.com/p/vim/source/detail?r=54194bd6ed607aa0cec07d8a27d4794573eabba9)
- [7.4.457 : cannot recognize getchar() returning K&#x5f;CURSORHOLD](http://code.google.com/p/vim/source/detail?r=96761b6789f61a8dee1551b7cea98e669fb90fb3)
- [7.4.458 : issue 252: Cursor moves in a zero-height window](http://code.google.com/p/vim/source/detail?r=fe1827921d50cb59ac680ea13e3f9103cd9db660)
- [7.4.460 : (after 7.4.454) can't build without the quickfix feature](http://code.google.com/p/vim/source/detail?r=d5eba03293a6193a1b462822800641b820252ce3)
- [7.4.461 : MS-Windows: When collate is on the number of copies is too high](http://code.google.com/p/vim/source/detail?r=4eba2bdb26f2be61bfdb1f635c0ed65f4bc35b0b)
- [7.4.462 : Setting the local value of 'backupcopy' empty gives an error](http://code.google.com/p/vim/source/detail?r=c0d93d67572b2d4161d0c5a5ab09477d6447ff49)
- [7.4.463 : tests 86 and 87 may hang on MS-Windows](http://code.google.com/p/vim/source/detail?r=6f80b9a773db0f3c69156f8768c8b6cfa4ad6785)

## 新着スクリプト

- [vim-nerdtree&#x5f;plugin&#x5f;open : Open files independent from vim (NERD tree extension)](http://www.vim.org/scripts/script.php?script_id=5014)
- [C64.vim : A scheme using the 16 Commodore 64 colors](http://www.vim.org/scripts/script.php?script_id=5015)
- [autocdls : Automatically shell-like ls after ':cd' in Vim](http://www.vim.org/scripts/script.php?script_id=5016)
- [rogue.vim : Porting of Rogue-clone II for Vim](http://www.vim.org/scripts/script.php?script_id=5017)
- [VullScreen : Full-screen mode for Vim on GNU/Linux](http://www.vim.org/scripts/script.php?script_id=5018)
- [oengsimple.vim : simplified keymap for Old English (Anglo Saxon)](http://www.vim.org/scripts/script.php?script_id=5019)
- [VIM PHP Refactoring Toolbox : Easily refactor your PHP Code](http://www.vim.org/scripts/script.php?script_id=5020)
- [vim-lamdify : Conceals lambdas as lambda characters in various languages](http://www.vim.org/scripts/script.php?script_id=5021)
- [cobalt : A dark blue color theme that works in 256 color terminals as well as GUI.](http://www.vim.org/scripts/script.php?script_id=5022)
- [vim-line-jump : easier way to jump lines in Tagbar and NERDTree](http://www.vim.org/scripts/script.php?script_id=5023)

## 月間ダウンロードランキング

1. [AutomaticLaTeXPlugin : Background compilation, completion, bib serch, toc and other nice features.](http://www.vim.org/scripts/script.php?script_id=2945) (3517)
2. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (2067)
3. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (1978)
4. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (998)
5. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (834)
6. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (742)
7. [pathogen.vim : Poor man's package manager. Easy manipulation of 'runtimepath' et al](http://www.vim.org/scripts/script.php?script_id=2332) (736)
8. [winmanager : A windows style IDE for Vim 6.0](http://www.vim.org/scripts/script.php?script_id=95) (711)
9. [minibufexpl.vim : Elegant buffer explorer - takes very little screen space](http://www.vim.org/scripts/script.php?script_id=159) (680)
10. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (664)

## vim-jp issues

Open : 224 (+9) | Closed : 420 (+12)

- [Issue #624 : E888 が重複している](https://github.com/vim-jp/issues/issues/624)
- [Issue #625 : Windows で encoding=utf-8 時にワイルドカード展開が iconv 変換されていない](https://github.com/vim-jp/issues/issues/625)
- [Issue #626 : signを操作するVim script関数がほしい](https://github.com/vim-jp/issues/issues/626)
- [Issue #627 : Windows で encoding=utf-8 を使用すると swap が残る](https://github.com/vim-jp/issues/issues/627)
- [Issue #628 : スコープ消滅時にスコープ変数がクリアされる](https://github.com/vim-jp/issues/issues/628)
- [Issue #629 : \[提案\]Vim に好きなアイコンを設定出来るパッチ](https://github.com/vim-jp/issues/issues/629)
- [Issue #630 : ビジュアルモードでの複数行選択後に入ったインサートモードでのカーソルキーを用いた編集が意図しない結果となる](https://github.com/vim-jp/issues/issues/630)
- [Issue #631 : 公式パッチのクローンが欲しい](https://github.com/vim-jp/issues/issues/631)
- [Issue #632 : lambda 関数作ろうぜ！](https://github.com/vim-jp/issues/issues/632)
- [Issue #633 : Cygwinのclangでif&#x5f;perlがリンクエラー](https://github.com/vim-jp/issues/issues/633)
- [Issue #634 : ヒアドキュメント作ろうぜ！](https://github.com/vim-jp/issues/issues/634)
- [Issue #635 : Windows の netrw で scp な URL を開けない](https://github.com/vim-jp/issues/issues/635)
- [Issue #636 : vi ftp://サーバー名/  コマンドを実行すると"netrw.vimの処理中にエラーが検出されました”のメッセージ](https://github.com/vim-jp/issues/issues/636)
- [Issue #637 : winrestview()に渡すtoplineに0以下の値を設定すると描画が崩れる](https://github.com/vim-jp/issues/issues/637)
- [Issue #638 : Windows で enc=utf-8 時にコマンドプロンプトのタイトルが化ける](https://github.com/vim-jp/issues/issues/638)
- [Issue #639 : Windows で複数部数を印刷すると、増える](https://github.com/vim-jp/issues/issues/639)
- [Issue #640 : プリンタ名にマルチバイト文字列があると印刷中ダイアログで?で表示される](https://github.com/vim-jp/issues/issues/640)
- [Issue #641 : Windowsで一部のテストが停止する](https://github.com/vim-jp/issues/issues/641)
- [Issue #642 : count関数の引数の順番が間違っている。](https://github.com/vim-jp/issues/issues/642)
- [Issue #643 : vim7.4.456以降でmkviewの後loadviewするとエラーになる。](https://github.com/vim-jp/issues/issues/643)
- [Issue #644 : n/N で normal!, noremap した時とデフォルトのfoldの扱いが違う](https://github.com/vim-jp/issues/issues/644)

