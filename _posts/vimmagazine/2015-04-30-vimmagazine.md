---
layout: post
category: vimmagazine
title: Vim Magazine 4 月号
---

## 話題

- [Support Neovim: The next generation of Vim! - Salt - Bountysource](https://salt.bountysource.com/teams/neovim)

  Neovim が新たな資金調達を開始しました。


- [Pure Python Vim clone.](https://github.com/jonathanslenders/pyvim)

  Python で実装された Vim クローン登場。


## 今月の新機能

- 7.4.687: `'t_SR'` オプションが追加されました。
- 7.4.710: `'listchars'` オプションに `space` が追加されました。

## リリース情報

- [7.4.684 : using non-unique temp file names when running Vim in diff mode](https://github.com/vim/vim/commit/e5c421cfd70a4d864faa0fac4e9f2dd6cdf5881d)
- [7.4.685 : with illegal utf-8 chars old regexp engine may crash](https://github.com/vim/vim/commit/0e462411cafdd908356792b2c229ab6369103bca)
- [7.4.686 : "zr" and "zm" do not take a count](https://github.com/vim/vim/commit/7d2757a47204d00cd47e3db94f1bd248c499d4e3)
- [7.4.687 : there is no way to use a different Replace cursor in terminal](https://github.com/vim/vim/commit/1e7813a1872df80eec1760bdd9d80352b306ac56)
- [7.4.688 : when "$" is in 'cpo' the popup menu isn't undrawn correctly](https://github.com/vim/vim/commit/478c46e50fd94f270369ec1c5f76aa65af7ee671)
- [7.4.689 : MS-Windows: problem with diff mode when 'autochdir' is set](https://github.com/vim/vim/commit/d87c36ea5eae50580f3c733734669250cc969019)
- [7.4.690 : memory access errors when changing indent in Ex mode](https://github.com/vim/vim/commit/da6365774d658b66f8d337f6342f4077594dd494)
- [7.4.691 : (after 7.4.689) can't build with MzScheme](https://github.com/vim/vim/commit/f6303878274b9aedfb9ac5b80d95002965093e5b)
- [7.4.692 : defining SOLARIS for no good reason](https://github.com/vim/vim/commit/de1502f2b48f77c98547a66b324e76be422228d3)
- [7.4.693 : session file is not correct when there are multiple tab pages](https://github.com/vim/vim/commit/695baeefe17adcd95f91d089efee87bd96fc98db)
- [7.4.694 : running tests changes the .viminfo file](https://github.com/vim/vim/commit/29d54df5cdd0ef4caa9c050854ddffdabece7a61)
- [7.4.695 : out-of-bounds read, dectected by Coverity](https://github.com/vim/vim/commit/7c94ce910f07e0d76d66b46bccfe39a07d8e1f6d)
- [7.4.696 : not freeing memory when encountering an error](https://github.com/vim/vim/commit/50ab994ab34fcbe0046a9a998ee420be196334ba)
- [7.4.697 : the filename used for ":profile" must be given literally](https://github.com/vim/vim/commit/d94682ffb4fb479e99f375431f82d67b38255629)
- [7.4.698 : various problems with locked and fixed lists and dictionaries](https://github.com/vim/vim/commit/9bc174b69d2775b0f27ed74bf372608d18c9f3d4)
- [7.4.699 : E315 when trying to delete a fold](https://github.com/vim/vim/commit/05b20fb14ec3219d73014595cfb6cc28ae1e1fc3)
- [7.4.700 : fold can't be opened after ":move"](https://github.com/vim/vim/commit/d5f6933d5c57ea6f79bbdeab6c426cf66a393f33)
- [7.4.701 : compiler warning for using uninitialized variable](https://github.com/vim/vim/commit/ec45c4a6dfd8acf52c8d1412188f54f79dde895e)
- [7.4.702 : joining an empty list does uneccessary work](https://github.com/vim/vim/commit/d39a7512804f860a20319c181613771a79a9ea4e)
- [7.4.703 : compiler warning for start&#x5f;dir unused when building unittests](https://github.com/vim/vim/commit/f9bde2b152d3344c694c5e1ca629156b34f1a627)
- [7.4.704 : invalid memory access if char search matches an illegal byte](https://github.com/vim/vim/commit/d82a2a990bc329754e1b61c5af669c76ac202497)
- [7.4.705 : can't build with Ruby 2.2](https://github.com/vim/vim/commit/bbc1a592a0e0efd5101a58032e7e7cc9f3e9f417)
- [7.4.706 : drawing error when 'laststatus' zero and a cmdline window](https://github.com/vim/vim/commit/5b14f89164fa2f40f52283226c98005887568317)
- [7.4.707 : undo files can have their executable bit set](https://github.com/vim/vim/commit/3cbac309f78f92434f75da5c4963d8caf4b6d3d3)
- [7.4.708 : gettext() is called too often](https://github.com/vim/vim/commit/77354e78a887e1b59ac519c5a1cb0e7fe9fc5899)
- [7.4.709 : ":tabmove" does not work as documented](https://github.com/vim/vim/commit/40ce3a4e1f50badb75ca812e26557a9bc5fde8c6)
- [7.4.710 : it is not possible to make spaces visibible in list mode](https://github.com/vim/vim/commit/79278367337be1e9a12222698642dcce0af1e059)
- [7.4.711 : (after 7.4.710) missing change in one file](https://github.com/vim/vim/commit/4c6b3b20dd86bb44247bb82ffc3b67eaee830954)
- [7.4.712 : missing change in another file](https://github.com/vim/vim/commit/1510f996c53243e5fe65a773087a1fa14879f503)

## 新着スクリプト

- [underscore.vim : Vim script utility library](http://www.vim.org/scripts/script.php?script_id=5149)
- [vim-mediawiki-editor : Edit MediaWiki right from Vim!](http://www.vim.org/scripts/script.php?script_id=5150)
- [tsuquyomi : Vim plugin for typescript](http://www.vim.org/scripts/script.php?script_id=5151)
- [createfilenext.vim : Creates a new file in the directory of the current buffer.](http://www.vim.org/scripts/script.php?script_id=5152)
- [AtelierPlateau : Plateau is a warm colorscheme with desaturated colors](http://www.vim.org/scripts/script.php?script_id=5153)
- [AtelierCave : Cave is a cool-warm colorscheme](http://www.vim.org/scripts/script.php?script_id=5154)
- [AtelierSavanna : Savanna is an understated greenish colorscheme with desaturated colors](http://www.vim.org/scripts/script.php?script_id=5155)
- [AtelierEstuary : Estuary is a yellow/greenish colorscheme](http://www.vim.org/scripts/script.php?script_id=5156)
- [Treemap : Create Treemaps with VIM](http://www.vim.org/scripts/script.php?script_id=5157)
- [XMLExpander : XML Pretty Printer](http://www.vim.org/scripts/script.php?script_id=5158)
- [gcov-marker : Show Gcov coverage within vim](http://www.vim.org/scripts/script.php?script_id=5159)
- [primary.vim : Inspired by Google's official colors](http://www.vim.org/scripts/script.php?script_id=5160)
- [yats : Yet Another TypeScript Syntax](http://www.vim.org/scripts/script.php?script_id=5161)
- [vison : Completion JSON files with JSON Schema](http://www.vim.org/scripts/script.php?script_id=5162)
- [mpi.vim : Syntax highlight, matching rules and mappings for the mpi](http://www.vim.org/scripts/script.php?script_id=5163)

## 月間ダウンロードランキング

1. [AutomaticLaTeXPlugin : Background compilation, completion, bib serch, toc and other nice features.](http://www.vim.org/scripts/script.php?script_id=2945) (3583)
2. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (2049)
3. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (1646)
4. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (1062)
5. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (791)
6. [winmanager : A windows style IDE for Vim 6.0](http://www.vim.org/scripts/script.php?script_id=95) (660)
7. [pathogen.vim : Poor man's package manager. Easy manipulation of 'runtimepath' et al](http://www.vim.org/scripts/script.php?script_id=2332) (648)
8. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (643)
9. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (597)
10. [minibufexpl.vim : Elegant buffer explorer - takes very little screen space](http://www.vim.org/scripts/script.php?script_id=159) (581)

## vim-jp issues

Open : 231 (+5) | Closed : 499 (+4)

- [Issue #722 : syntax/css.vim のハイライトがあまい](https://github.com/vim-jp/issues/issues/722)
- [Issue #723 : vimdiffを提供する良い方法](https://github.com/vim-jp/issues/issues/723)
- [Issue #724 : リストの map() が極端に遅くなる](https://github.com/vim-jp/issues/issues/724)
- [Issue #725 : groovy のコメントハイライトが間違っている](https://github.com/vim-jp/issues/issues/725)
- [Issue #726 : round() の第二引数に桁が欲しい](https://github.com/vim-jp/issues/issues/726)
- [Issue #727 : Vim 7.4.711でtest&#x5f;listcharsがFAILする。](https://github.com/vim-jp/issues/issues/727)
- [Issue #728 : windowタイトルの表示幅の計算がおかしい？](https://github.com/vim-jp/issues/issues/728)
- [Issue #729 : Vim基本セットのjavascript環境をちょっといい感じにする](https://github.com/vim-jp/issues/issues/729)
- [Issue #730 : fcitx の未変換状態が gvim で表示されない](https://github.com/vim-jp/issues/issues/730)

