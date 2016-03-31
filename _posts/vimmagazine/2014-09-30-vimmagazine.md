---
layout: post
category: vimmagazine
title: Vim Magazine 9 月号
---

## 話題

- [rogue.vim : Porting of Rogue-clone II for Vim](http://www.vim.org/scripts/script.php?script_id=5017)

  名作RPG『Rogue Clone II』が Vim に移植されました。

- [Osaka.vim #3](http://osaka-vim.connpass.com/event/9119/)

  待ちに待った第三回です。 今回はもくもく会のみになります。 10月25日(土) 大阪

## 今月の新機能

  - 7.4.434: gettabvar() でスコープ変数が取れるようになりました。
  - 7.4.447: IGNOREEXTRA フラグが追加されました。(spell関連)
  - 7.4.449: `:helpclose` コマンドが追加されました。
  - 7.4.450: +cmd 引数が関連のあるコマンドに追加されました。
  - 7.4.456: 'backupcopy' オプションが global-local になりました。バッファごとに設定を変更できるようになります。
  - 7.4.459: MS-Windows: Vim のアイコンを変更できるようになりました。

## リリース情報

- [7.4.431 : compiler warning](https://github.com/vim/vim/commit/120f4a878f61d78b9be34e3516aa4cdf0fc682d7)
- [7.4.432 : MS-Windows: setting 'encoding' does not convert arguments](https://github.com/vim/vim/commit/14993322b1bdcd507f439fbbf6dd3d051f777373)
- [7.4.433 : test 75 fails on MS-Windows](https://github.com/vim/vim/commit/13e2a0af665cffa3c4485be843feca70d90a7918)
- [7.4.434 : gettabvar() is not consistent with getwinvar() and getbufvar()](https://github.com/vim/vim/commit/0e2ea1beb471a24dd86a45c439a98e5d758b4270)
- [7.4.435 : line formatting behaves differently when 'linebreak' is set](https://github.com/vim/vim/commit/0026d47d8cb6e6b135527ca489bbf81ac8e0045a)
- [7.4.436 : ml&#x5f;get error for autocommand that moves the cursor](https://github.com/vim/vim/commit/371932a7754453b5a3adbd41959056fc9a45a9fd)
- [7.4.437 : new and old regexp engine are not consistent](https://github.com/vim/vim/commit/fb031407304b82b49f6f367bfba31f602837fc49)
- [7.4.438 : cached values for 'cino' not reset for ":set all&"](https://github.com/vim/vim/commit/5a4eceb02a12f1746b84eea680ddb511924248ac)
- [7.4.439 : duplicate message in message history](https://github.com/vim/vim/commit/f638cbce7623e2233f7912a9493cd51e6adff159)
- [7.4.440 : omni complete popup drawn incorrectly](https://github.com/vim/vim/commit/834def3872425b04aa0f52f868d056e60e614516)
- [7.4.441 : endless loop and other problems when 'cedit' is set to CTRL-C](https://github.com/vim/vim/commit/58da707d2412f60eb5c0e158ade48dd8a13c96ed)
- [7.4.442 : (after 7.4.434) using unitinialized variable](https://github.com/vim/vim/commit/3089a1068dd5a0645cea0df09a1c8a1f00146649)
- [7.4.443 : error reported by ubsan when running test 72](https://github.com/vim/vim/commit/3516928c75f38e4304b76f0697175dbb055af6a2)
- [7.4.444 : reversed question mark not recognized as punctuation](https://github.com/vim/vim/commit/103650de9083ba5451869a3f8ed729ba47746b3a)
- [7.4.445 : clipboard may be cleared on startup](https://github.com/vim/vim/commit/714db3bb819bb7f9f52270f2352b274bba3a9ff4)
- [7.4.446 : environment for executing autocommand may not be restored.](https://github.com/vim/vim/commit/5d2bae8b1c2bb8059010ea604e092de9cf3c1baf)
- [7.4.447 : spell files from Hunspell may generate a lot of errors](https://github.com/vim/vim/commit/b4b43bb572ca4c6fab846515274913bdb5ee2aed)
- [7.4.448 : MS-Windows: Using ETO&#x5f;IGNORELANGUAGE causes problems](https://github.com/vim/vim/commit/4ee40b051545c3749d9d215f3500977792d6db72)
- [7.4.449 : can't easily close the help window](https://github.com/vim/vim/commit/5bfa2ed6e47a2bc4e15e43cf6504e97401f2e5e6)
- [7.4.450 : some commands that edit another buffer don't support +cmd](https://github.com/vim/vim/commit/9c8d9e19523a312e888f413ca65d0863fce08f5f)
- [7.4.451 : calling system() with empty input gives an error](https://github.com/vim/vim/commit/1ecfd9cb09d7cd2231940a519c1af36d9636966a)
- [7.4.452 : can't build with tiny features](https://github.com/vim/vim/commit/eb21e4ccf5453ecd289e9f7e09fa27a978fc0360)
- [7.4.453 : still can't build with tiny features](https://github.com/vim/vim/commit/3fa57e0d50d90248de9ad1867203b8ce6eedc92a)
- [7.4.454 : CTRL-W-\] does not work when using a Visual selection](https://github.com/vim/vim/commit/d355c50a2a07eaec87729d405553e15e5e680ca0)
- [7.4.455 : completion for :buf does not use 'wildignorecase'](https://github.com/vim/vim/commit/4b9d637e9cb4864c28a331ecf4f9c67696240897)
- [7.4.456 : 'backupcopy' is global, can't set a local value](https://github.com/vim/vim/commit/b8ee25acabe887aec5008fd254c4e5c61bebbb04)
- [7.4.457 : cannot recognize getchar() returning K&#x5f;CURSORHOLD](https://github.com/vim/vim/commit/1db60c47d9a0f14f4586702bfd8fccd8ea96f83f)
- [7.4.458 : issue 252: Cursor moves in a zero-height window](https://github.com/vim/vim/commit/cfc216e7b00690805aef8683fab000654937191d)
- [7.4.459 : Can't change the icon after building Vim.](https://github.com/vim/vim/commit/cddc91ccb4cc0989fa0169a343289873a6d2d651)
- [7.4.460 : (after 7.4.454) can't build without the quickfix feature](https://github.com/vim/vim/commit/da014b94417a6e270787704d24c2cd735bc83e38)
- [7.4.461 : MS-Windows: When collate is on the number of copies is too high](https://github.com/vim/vim/commit/7ddc642a0c786f834ab0b6851b80dbed6e04204b)
- [7.4.462 : Setting the local value of 'backupcopy' empty gives an error](https://github.com/vim/vim/commit/84d17a6f9e3aee1adaeae3b073b5c25401bf70d3)
- [7.4.463 : tests 86 and 87 may hang on MS-Windows](https://github.com/vim/vim/commit/56f6227224c6c0434fc758ab82245069d11c4fd4)

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

