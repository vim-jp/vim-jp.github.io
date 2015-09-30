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

- [7.4.390 : advancing pointer over end of a string](https://github.com/vim/vim/commit/ece29e80131afdd4bb7aa74f62b98cf97f1090b5)
- [7.4.391 : no 'cul' highlight with cursor on a line with diff highlight](https://github.com/vim/vim/commit/e0f148270a03e0da2bf21706bee4d2fe99146c55)
- [7.4.392 : not easy to detect type of command line window](https://github.com/vim/vim/commit/8c1329cb59e964797b8abdcf8d42af9877ea8daa)
- [7.4.393 : text drawing on newer MS-Windows systems is suboptimal](https://github.com/vim/vim/commit/b5a7a8b5451c6fe8a2cc1d5d86c42d9b9a50ef98)
- [7.4.394 : (after 7.4.393) with DirectX last italic char is incomplete](https://github.com/vim/vim/commit/9b352c46f0604288087c629381158e3e80943fbe)
- [7.4.395 : C indent wrong below an if with wrapped () followed by {}](https://github.com/vim/vim/commit/04d17ae1678846c4857cd86cf3eaf47d60c04c85)
- [7.4.396 : when 'clipboard' is "unnamed", :g/pat/d is very slow](https://github.com/vim/vim/commit/6b1ee34aa0236b50f675f3bbcd9bf0b7a3384f7f)
- [7.4.397 : matchparen only uses the topmost syntax item](https://github.com/vim/vim/commit/e98cfe1cc9c1f63fc280452593b4b4cbfe2f6356)
- [7.4.398 : (after 7.4.393) gcc error for argument of InterlockedIncrement](https://github.com/vim/vim/commit/0106b4b89127b043eddf711c750364b487deb794)
- [7.4.399 : crypt implementation is messy; blowfish crypt has a weakness](https://github.com/vim/vim/commit/8f4ac01544b44bdd906d241e4f203de7496e5ac8)
- [7.4.400 : list of distributed files is incomplete](https://github.com/vim/vim/commit/d21e8b3a17f1dbec348ff94f03466c5c3c4647bb)
- [7.4.401 : (after 7.4.399) new files missing in Makefiles](https://github.com/vim/vim/commit/07cf382e223b8b88f4b763d104e3e487c96d32d2)
- [7.4.402 : test 72 crashes under certain conditions](https://github.com/vim/vim/commit/5a669b94814a5ca1ff69f76f18c75a90f4404ca6)
- [7.4.403 : valgrind reports errors when running test 72](https://github.com/vim/vim/commit/2be795084f053928879d758443e199ef147cc07e)
- [7.4.404 : Windows 64 bit compiler warnings](https://github.com/vim/vim/commit/9b8f021d93de3a2a2905d0e46b9f0b6dbb3e12f9)
- [7.4.405 : screen updating is slow when using matches](https://github.com/vim/vim/commit/5307de0746359098d63053b5fb58dbfb924ce0fd)
- [7.4.406 : test 72 and 100 fail on MS-Windows](https://github.com/vim/vim/commit/0ddd1b5e3568f3402d5e58ee054c9c412eab0b94)
- [7.4.407 : Visual block insert with cursor move repeats the wrong text](https://github.com/vim/vim/commit/1fc7e972b00ce271328fcb2734a321562b7194da)
- [7.4.408 : Visual block insert breaks a multi-byte character](https://github.com/vim/vim/commit/b5cf6c34e84bd52ba43c3db06ff337abbcb08326)
- [7.4.409 : can't build with Perl on Fedora 20](https://github.com/vim/vim/commit/d5f62b14e1d70f23866e613449ca71ac0b903e0f)
- [7.4.410 : fold does not open after search with CmdwinLeave autocmd](https://github.com/vim/vim/commit/42f06f9dbc1e167a8ddef2ec500d952870600d5e)
- [7.4.411 : "foo bar" sorts before "foo" with sort()](https://github.com/vim/vim/commit/1b338d2022cb28fe143b134a966f13f18d14596c)
- [7.4.412 : can't build on Windows XP with MSVC](https://github.com/vim/vim/commit/ce2f2e061db91cc62e13023695dfa4a2755581f5)
- [7.4.413 : MS-Windows: inserting dead key does not always work](https://github.com/vim/vim/commit/d2e8087b532e676cc513a2b385deaf1bf7afbf94)
- [7.4.414 : cannot define a command only when it's used](https://github.com/vim/vim/commit/d5005164e1f367136c956d2941478a3aea30ec47)
- [7.4.415 : (after 7.4.414) cannot build, warning for shadowed variable](https://github.com/vim/vim/commit/92f4baace8af8fbc6ed610ba91268e38a4d555a2)
- [7.4.416 : problem with breakindent/showbreak and tabs](https://github.com/vim/vim/commit/95765089755c57d0d73252d4673c9d2deee337ff)
- [7.4.417 : minimum width used for breakindent is not set after split](https://github.com/vim/vim/commit/285ed7e049dc50c61672fb835752643bd01ed1ca)
- [7.4.418 : when leaving ":append" the cursor shape is like in Insert mode](https://github.com/vim/vim/commit/26f08b06b681085e8f40b109871fe7ae2f3bd7c6)
- [7.4.419 : when part of a list is locked it's possible to make changes](https://github.com/vim/vim/commit/f2d912e8bc64927f9089e5da9185b9f108d28c0d)
- [7.4.420 : it's not obvious how to add a new test](https://github.com/vim/vim/commit/a9537d238e8c2fe9afb9bbf4e137734372b7d9ba)
- [7.4.421 : crash when searching for "\\ze&#x2a;"](https://github.com/vim/vim/commit/2d46e6075ba3aa369172e610782810b9ac3f1f4b)
- [7.4.422 : with conceal + linebreak some text is not displayed correctly](https://github.com/vim/vim/commit/49f9dd7b916fb32bfd0cc9a50f8c4f9bb1bb760b)
- [7.4.423 : expand("$shell") does not work as documented](https://github.com/vim/vim/commit/e4df164692fef1a3c94149e1e2af68d44d335bf9)
- [7.4.424 : ml&#x5f;get error when using Python to delete lines in a buffer](https://github.com/vim/vim/commit/d7408fa0ceb7bf414ca5c17f8ae3bc13131b7b15)
- [7.4.425 : when 'showbreak' is used "gj" may move to the wrong position](https://github.com/vim/vim/commit/773b158de8f813b0683fc56e37639937e5c9c1ee)
- [7.4.426 : README File missing from list of files](https://github.com/vim/vim/commit/96f3a49b3bab421edebff24f070634ae876edd7f)
- [7.4.427 : system() typeahead may be echoed and mess up the display](https://github.com/vim/vim/commit/52a7246f015be40e54007c0c3a1d9e37dc14045a)
- [7.4.428 : executable() may return a wrong result on MS-Windows](https://github.com/vim/vim/commit/c40bdee42a58c48e00c76886289904849e8c4c3c)
- [7.4.429 : build fails with fewer features](https://github.com/vim/vim/commit/84d8cdd76900a8834260bc69d0be202923614ae0)
- [7.4.430 : test&#x5f;listlbr fails when compiled with normal features](https://github.com/vim/vim/commit/5d5fbe7a2b4cc948c54afeae5375b247bf34ba9b)

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

