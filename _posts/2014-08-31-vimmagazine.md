---
layout: post
category: vimmagazine
title: Vim Magazine 8 月号
---

## 話題

- [Nagoya.vim #2](http://nagoyavim.connpass.com/event/8299/)

  しばらく休止していた Nagoya.vim ですが、この度もくもく会として復活します。  9月20日(土) 名古屋

## 今月の新機能

- 7.4.392 : `getcmdwintype()` 関数が追加されました。
- 7.4.393 : MS-Windows: Direct-X (DirectWrite) によるテキスト描画を有効にするためのオプション `'renderoptions'` が追加されました。
- 7.4.399 : 暗号化方法に "blowfish2" が追加されました。
- 7.4.414 : 自動コマンドイベント `CmdUndefined` が追加されました。

## リリース情報

- [7.4.390 : advancing pointer over end of a string](http://code.google.com/p/vim/source/detail?r=d3a674f6c7370f70d6b66dcefb9e36d22d73c7b7)
- [7.4.391 : no 'cul' highlight with cursor on a line with diff highlight](http://code.google.com/p/vim/source/detail?r=f051e50a6a5feb0f0c6d209440d9d305b82f5836)
- [7.4.392 : not easy to detect type of command line window](http://code.google.com/p/vim/source/detail?r=589fd07888abc8a1208616c56b21898d3352fc5b)
- [7.4.393 : text drawing on newer MS-Windows systems is suboptimal](http://code.google.com/p/vim/source/detail?r=1bff71d202621d97acc4ef21ce9f47b78f68bda3)
- [7.4.394 : (after 7.4.393) with DirectX last italic char is incomplete](http://code.google.com/p/vim/source/detail?r=d594c15521377493d5e23b03e4dbf5108df69acd)
- [7.4.395 : C indent wrong below an if with wrapped () followed by {}](http://code.google.com/p/vim/source/detail?r=65b0974c5bebf5d6b56e9abc90d86759d1c4f6d7)
- [7.4.396 : when 'clipboard' is "unnamed", :g/pat/d is very slow](http://code.google.com/p/vim/source/detail?r=7766142fc7d3e90c2e15a9c606efcd97331edef8)
- [7.4.397 : matchparen only uses the topmost syntax item](http://code.google.com/p/vim/source/detail?r=8d361608fe86c534e51e4c7e583375642d34b7a4)
- [7.4.398 : (after 7.4.393) gcc error for argument of InterlockedIncrement](http://code.google.com/p/vim/source/detail?r=f62b2e76dd809dc57cad9d0fd8f04a4f3a7193e9)
- [7.4.399 : crypt implementation is messy; blowfish crypt has a weakness](http://code.google.com/p/vim/source/detail?r=18ac55444b37bcae8ea090f7b3abf4826c679bc3)
- [7.4.400 : list of distributed files is incomplete](http://code.google.com/p/vim/source/detail?r=50a658bf4cff030cbe6a97a53bb838b5976c2e1a)
- [7.4.401 : (after 7.4.399) new files missing in Makefiles](http://code.google.com/p/vim/source/detail?r=ab71bb81b84e9a1cdf030b172f61d779ccda975c)
- [7.4.402 : test 72 crashes under certain conditions](http://code.google.com/p/vim/source/detail?r=3aa37ad2c4ed4d4f979e7c7e7d5d88586c6b74ba)
- [7.4.403 : valgrind reports errors when running test 72](http://code.google.com/p/vim/source/detail?r=86fb698a38d5c10f59928e57da94ce5b3e124a40)
- [7.4.404 : Windows 64 bit compiler warnings](http://code.google.com/p/vim/source/detail?r=0242c27e40e1ce913ec2b6540371d5ea5fceeac9)
- [7.4.405 : screen updating is slow when using matches](http://code.google.com/p/vim/source/detail?r=4c25af735304ff17a27bb3ce21e7f7d3b304ff88)
- [7.4.406 : test 72 and 100 fail on MS-Windows](http://code.google.com/p/vim/source/detail?r=1134bba900c17abcecdc414c6e30da22c17d32a5)
- [7.4.407 : Visual block insert with cursor move repeats the wrong text](http://code.google.com/p/vim/source/detail?r=3f92ef156c664fb8fea1f036a4a4c87af487ba6b)
- [7.4.408 : Visual block insert breaks a multi-byte character](http://code.google.com/p/vim/source/detail?r=cb3218a69c2f0949bbff5bcf9beeccc83ef2ef59)
- [7.4.409 : can't build with Perl on Fedora 20](http://code.google.com/p/vim/source/detail?r=de35e123e63c326be4a576f8e678ed9e63ed1f48)
- [7.4.410 : fold does not open after search with CmdwinLeave autocmd](http://code.google.com/p/vim/source/detail?r=2fd550c75256cc1ecf06a4afc6a856b305cf7a8a)
- [7.4.411 : "foo bar" sorts before "foo" with sort()](http://code.google.com/p/vim/source/detail?r=d4e0c197e97968050596f4c441552383e31db63b)
- [7.4.412 : can't build on Windows XP with MSVC](http://code.google.com/p/vim/source/detail?r=4c845a1600d28db8edabd211e6d7a9cd03883ea6)
- [7.4.413 : MS-Windows: inserting dead key does not always work](http://code.google.com/p/vim/source/detail?r=121613e72e397db495f5e9b640b249aead5a61c1)
- [7.4.414 : cannot define a command only when it's used](http://code.google.com/p/vim/source/detail?r=67edf731a2403a8023440184872159ddc073093c)
- [7.4.415 : (after 7.4.414) cannot build, warning for shadowed variable](http://code.google.com/p/vim/source/detail?r=20dbceb6f4713ccd01be45dc531abc269fbb7579)
- [7.4.416 : problem with breakindent/showbreak and tabs](http://code.google.com/p/vim/source/detail?r=735bd597a8ffcb00edff68b34a26573b8f3cad9a)
- [7.4.417 : minimum width used for breakindent is not set after split](http://code.google.com/p/vim/source/detail?r=e60327caf9095d5ec067a5ab47bf4a69c0f58580)
- [7.4.418 : when leaving ":append" the cursor shape is like in Insert mode](http://code.google.com/p/vim/source/detail?r=00f58478e38b68faf783d8fe70c1439fe9dd732d)
- [7.4.419 : when part of a list is locked it's possible to make changes](http://code.google.com/p/vim/source/detail?r=0a42938f449cf306a429db373106d5e29d8d4312)
- [7.4.420 : it's not obvious how to add a new test](http://code.google.com/p/vim/source/detail?r=93aa0340130912c69d4093fc626a3cbb0cc7ad68)
- [7.4.421 : crash when searching for "\\ze&#x2a;"](http://code.google.com/p/vim/source/detail?r=3ee39fe2df7d60d6376b3429383c5c5e3f9bb3fe)
- [7.4.422 : with conceal + linebreak some text is not displayed correctly](http://code.google.com/p/vim/source/detail?r=69da1498ce89dd093935c0e5013e1df8060a4e80)
- [7.4.423 : expand("$shell") does not work as documented](http://code.google.com/p/vim/source/detail?r=96a4fa8e530c12137af235a9aaad57d578f9eebd)
- [7.4.424 : ml&#x5f;get error when using Python to delete lines in a buffer](http://code.google.com/p/vim/source/detail?r=85a1802ae810f0e840f8b5d8f2827b4199b39750)
- [7.4.425 : when 'showbreak' is used "gj" may move to the wrong position](http://code.google.com/p/vim/source/detail?r=2e804ebb3ee8cd59451f8f1c1be2f09ebce0d77a)
- [7.4.426 : README File missing from list of files](http://code.google.com/p/vim/source/detail?r=7ab9a9cf2859ae4b7c00c4fff016fd5a0869450e)
- [7.4.427 : system() typeahead may be echoed and mess up the display](http://code.google.com/p/vim/source/detail?r=fca35aa9380a9e3acecc78f85e05f26436e87ae1)
- [7.4.428 : executable() may return a wrong result on MS-Windows](http://code.google.com/p/vim/source/detail?r=1fe61f6d52079f44f61b3e03cf9b0c300f2890d9)
- [7.4.429 : build fails with fewer features](http://code.google.com/p/vim/source/detail?r=e98e415ea3b8c743a5ab04ad9f498ab1b11ae8cf)
- [7.4.430 : test&#x5f;listlbr fails when compiled with normal features](http://code.google.com/p/vim/source/detail?r=f4e9c65a18d4a5753747f327ed55c10f05d8d91c)

## 新着スクリプト

- [vCoolor : A Simple color selector/picker plugin](http://www.vim.org/scripts/script.php?script_id=4995)
- [AdvancedDiffOptions : Additional diff options and commands to manage them.](http://www.vim.org/scripts/script.php?script_id=4996)
- [SearchInRange : Limit search to range when jumping to the next search result.](http://www.vim.org/scripts/script.php?script_id=4997)
- [doctest.vim :   DocTest in vim](http://www.vim.org/scripts/script.php?script_id=4998)
- [InstantRst : Preview rst document instantly.](http://www.vim.org/scripts/script.php?script_id=4999)
- [clickable.vim :  make things clickable](http://www.vim.org/scripts/script.php?script_id=5000)
- [logstash.vim : logstash.vim highlights configuration files for logstash (see http://logstash.ne](http://www.vim.org/scripts/script.php?script_id=5001)
- [vim-autoprefixer : Adds CSS autoprefixer support to Vim.](http://www.vim.org/scripts/script.php?script_id=5002)
- [Dutyl : Coordinate D tools to work together for you](http://www.vim.org/scripts/script.php?script_id=5003)
- [vim-rubygems : The plugin provides ability to work with rubygems.org inside Vim](http://www.vim.org/scripts/script.php?script_id=5004)
- [Session-Viminfo-Management : The plug-in improves upon Vim's built-in :mksession and :wviminfo commands.](http://www.vim.org/scripts/script.php?script_id=5005)
- [Vim-EPUB : Plugin for EPUB ebooks' edition](http://www.vim.org/scripts/script.php?script_id=5006)
- [phpmanual : PHP manual for vim, update to PHP5.5](http://www.vim.org/scripts/script.php?script_id=5007)
- [a indent : a&#x5f;indent helps you to select/move in different code block](http://www.vim.org/scripts/script.php?script_id=5008)
- [256&#x5f;darkdot : DarkDot color scheme for 256 color terminals](http://www.vim.org/scripts/script.php?script_id=5009)
- [json-formatter.vim : A VIM plugin for formatting saved JSON file.](http://www.vim.org/scripts/script.php?script_id=5010)
- [autorepeat.vim : Repeat the normal mode command automatically.](http://www.vim.org/scripts/script.php?script_id=5011)
- [Vim Markdown Preview : A Vim plugin for previewing markdown files in a browser.](http://www.vim.org/scripts/script.php?script_id=5012)
- [autotype.vim : Yes, Vim will work for you.](http://www.vim.org/scripts/script.php?script_id=5013)

## 月間ダウンロードランキング

1. [AutomaticLaTeXPlugin : Background compilation, completion, bib serch, toc and other nice features.](http://www.vim.org/scripts/script.php?script_id=2945) (10672)
2. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (2233)
3. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (2186)
4. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (1143)
5. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (892)
6. [pathogen.vim : Poor man's package manager. Easy manipulation of 'runtimepath' et al](http://www.vim.org/scripts/script.php?script_id=2332) (869)
7. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (782)
8. [winmanager : A windows style IDE for Vim 6.0](http://www.vim.org/scripts/script.php?script_id=95) (769)
9. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (745)
10. [minibufexpl.vim : Elegant buffer explorer - takes very little screen space](http://www.vim.org/scripts/script.php?script_id=159) (719)

## vim-jp issues

Open : 215 (+16) | Closed : 408 (+6)

- [Issue #602 : Windowsでvim -gしたらgvimを起動する](https://github.com/vim-jp/issues/issues/602)
- [Issue #603 : vim 7.4.364 で hlsearch と cursorline をオンにしていると描画が遅くなった](https://github.com/vim-jp/issues/issues/603)
- [Issue #604 : $でカーソルを行末尾に移動直後にcurswantの値がおかしい](https://github.com/vim-jp/issues/issues/604)
- [Issue #605 : removing bevel (inset) borders. (Linux)](https://github.com/vim-jp/issues/issues/605)
- [Issue #606 : Vim script の実行に CTRL-C で割り込めなくできる](https://github.com/vim-jp/issues/issues/606)
- [Issue #607 : updatetime の間隔で謎のキーシーケンス \<80>\<fd>&#x60; が送られてくる](https://github.com/vim-jp/issues/issues/607)
- [Issue #608 : アイドル時に実行されるGC処理のパフォーマンスが悪く、フリーズすることがある](https://github.com/vim-jp/issues/issues/608)
- [Issue #609 : 補完候補の表示位置が本来の位置からズレる](https://github.com/vim-jp/issues/issues/609)
- [Issue #610 : VimからPython 3.xを使うと落ちることがある](https://github.com/vim-jp/issues/issues/610)
- [Issue #611 : vim を終了しても端末のハイライトが消えない](https://github.com/vim-jp/issues/issues/611)
- [Issue #612 : 複数行を削除した時のメッセージが重複している？](https://github.com/vim-jp/issues/issues/612)
- [Issue #613 : cursorline と cursorcolumn の動きが違う](https://github.com/vim-jp/issues/issues/613)
- [Issue #614 : filetype indent on を設定し行挿入を行うと意図してないコマンドが実行される](https://github.com/vim-jp/issues/issues/614)
- [Issue #615 : プロセス置換ファイルオープン時のfileencodingsの振る舞いについて](https://github.com/vim-jp/issues/issues/615)
- [Issue #616 : ユーザー定義コマンドで -complete=file を使うとスペースが受け取れない](https://github.com/vim-jp/issues/issues/616)
- [Issue #617 : 'set t&#x5f;ut='時、'hi Normal ctermbg'の指定が右下半角1文字分効かない](https://github.com/vim-jp/issues/issues/617)
- [Issue #618 : grepformat の実際の既定値とドキュメントでの既定値が異なっている](https://github.com/vim-jp/issues/issues/618)
- [Issue #619 : text-objectsのi)の挙動が理解できない](https://github.com/vim-jp/issues/issues/619)
- [Issue #620 : cygwin/msys/Linux で enc=cp932 ft=diff とするとエラー](https://github.com/vim-jp/issues/issues/620)
- [Issue #621 : 'cedit' を \<C-c> に設定すると Vim が固まる](https://github.com/vim-jp/issues/issues/621)
- [Issue #622 : gettabvar(1, '') はスコープの辞書を返して欲しい](https://github.com/vim-jp/issues/issues/622)
- [Issue #623 : NFAエンジンでは '\[  のマーク位置を示す正規表現パターンが使えない？](https://github.com/vim-jp/issues/issues/623)

