---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2017 年 4 月号

---

## 話題

* [k-takata/minpac](https://github.com/k-takata/minpac) の正式リリース v1.0 がリリースされました

## 今月の新機能

今月の新機能及びユーザーに影響のある変更は以下のとおりです。

*   8.0.0580: `setqflist()` 関数に `"valid"` フラグを指定できるようになりました

## Vimに関する脆弱性

(特筆すべき脆弱性の報告はありません)

## リリース情報

- [8.0.0530 : buffer overflow when 'columns' is very big](https://github.com/vim/vim/commit/658a3a2caf5852d071b6b1be92d9d6614a6208dc)
- [8.0.0531 : test with long directory name fails on non-unix systems](https://github.com/vim/vim/commit/9b81079ddd839a666682f6bdbc24890bf4d1a42c)
- [8.0.0532 : test with long directory name fails on Mac](https://github.com/vim/vim/commit/c77d6757471fa207520586bbdbc1b30af84cf5c8)
- [8.0.0533 : abbreviation doesn't work after backspacing newline](https://github.com/vim/vim/commit/878c263a489b7e211eda31fa13a3d5ad9e120554)
- [8.0.0534 : defaults.vim does not work well with tiny features](https://github.com/vim/vim/commit/33ccb24cf7891911561adccf9f2af8859e581198)
- [8.0.0535 : memory leak when exiting from within a user function](https://github.com/vim/vim/commit/6914c64ee58ce68f31fb8a8793293a9b3f2f6240)
- [8.0.0536 : quickfix window not updated when freeing quickfix stack](https://github.com/vim/vim/commit/69f40be64555d50f603c6f22722cf762aaa6bbc1)
- [8.0.0537 : illegal memory access with :z and large count](https://github.com/vim/vim/commit/fa0ad0bb0b4255e64ebcf9269d60a942e0ae7ff9)
- [8.0.0538 : no test for falling back to default term value](https://github.com/vim/vim/commit/85045a73db258a054a17fd52a67eb5cd02a788dc)
- [8.0.0539 : (after 8.0.0538) startup test fails on Mac](https://github.com/vim/vim/commit/08f88b139d2f93661ed2b17214ee6b308b4edb5b)
- [8.0.0540 : (after 8.0.0540) building unit tests fails](https://github.com/vim/vim/commit/a604429529ce3d3c97309c7cf6bba8b699f6bfb5)
- [8.0.0541 : compiler warning on MS-Windows](https://github.com/vim/vim/commit/04000560ca81cc2608b291d0990e661b41ca8c68)
- [8.0.0542 : getpos() can return a negative line number](https://github.com/vim/vim/commit/a1d5fa65bc7e8a548858e9c295a192b63dcd011b)
- [8.0.0543 : test&#x5f;edit causes older xfce4-terminal to close](https://github.com/vim/vim/commit/ba6ec182973af726ce9b7b7eb3753fc3a7ae7d1b)
- [8.0.0544 : cppcheck warnings](https://github.com/vim/vim/commit/866c68861071f8cd1ef5a82445bebaafc8626e7e)
- [8.0.0545 : edit test may fail on some systems](https://github.com/vim/vim/commit/15ecbd6f3d39ff04862999a577962ef9369a9e53)
- [8.0.0546 : swap file exists briefly when opening the command window](https://github.com/vim/vim/commit/3bab93998d01a01b7f2a071fa3b8054bb0094625)
- [8.0.0547 : extra line break in verbosefile](https://github.com/vim/vim/commit/52604f2454e5369f861d3ce34764f74a0999c773)
- [8.0.0548 : saving the redo buffer only works one time](https://github.com/vim/vim/commit/d4863aa99e0527e9505c79cbeafc68a6832200bf)
- [8.0.0549 : no test for the 8g8 command](https://github.com/vim/vim/commit/395b6bab33e5ed9a0377d7f140e98fd3ab682672)
- [8.0.0550 : cannot parse some etags format tags file](https://github.com/vim/vim/commit/9585a1655ba0d34ea88574617112093a9bd4f2e9)
- [8.0.0551 : the typeahead buffer is reallocated too often](https://github.com/vim/vim/commit/d34f9b1155a6b470e1dc766ff98192f440e7eba7)
- [8.0.0552 : toupper and tolower don't work properly for Turkish](https://github.com/vim/vim/commit/3317d5ebbe8304da82b8088446060afcae0012af)
- [8.0.0553 : toupper/tolower test fails on Mac](https://github.com/vim/vim/commit/9f4de1f5435b900e43e19766da1a5bed4686cf44)
- [8.0.0554 : toupper and tolower don't work properly for Turkish](https://github.com/vim/vim/commit/1cc482069a3407132aeb43a55d6dc284153e79c7)
- [8.0.0555 : toupper/tolower test fails on OSX without Darwin](https://github.com/vim/vim/commit/d2381a2cadb9ef359ad5efb916734c635b29bd13)
- [8.0.0556 : getting window position fails if GUI and term possible](https://github.com/vim/vim/commit/7860bac57baeda408d5ef84f3334b0e52cc690bd)
- [8.0.0557 : GTK: using static gravities is not useful](https://github.com/vim/vim/commit/59053e1d9f41f9afa16ac7e266ee825a2eb2dd59)
- [8.0.0558 : :ownsyntax is not tested](https://github.com/vim/vim/commit/f8ec998613d8037e345f4e7e08460dfc15c528a9)
- [8.0.0559 : setting ttytype to xxx does not always fail](https://github.com/vim/vim/commit/f803a7697808f63949ea26fc256d4c9a8647089e)
- [8.0.0560 : :windo allows for ! but it's not supported](https://github.com/vim/vim/commit/451a4a1cb7797e5d9b9fd625671cb5c652e7da00)
- [8.0.0561 : undefined behavior when using backslash after empty line](https://github.com/vim/vim/commit/478af67dd6a9adc456464c3736bda328ae3a28cb)
- [8.0.0562 : not enough test coverage for syntax commands](https://github.com/vim/vim/commit/ea588154d00aec288e48b344b2685e46bea99cd6)
- [8.0.0563 : crash when getting the window position in tmux](https://github.com/vim/vim/commit/9f9288627795cb6cc6bff332d28eac86b1dbcf13)
- [8.0.0564 : cannot detect Bazel BUILD files on some systems](https://github.com/vim/vim/commit/39170e2d9761345df4be67d4d3928ac1094b9adf)
- [8.0.0565 : using freed memory in :caddbuf](https://github.com/vim/vim/commit/31bdd13c335533c749993b57dcd980a87373139e)
- [8.0.0566 : setting nocompatible for the tiny version moves the cursor](https://github.com/vim/vim/commit/43d1ac6e81a6796f2e6b3514bbe029cf9bd0a0ae)
- [8.0.0567 : call for requesting color and ambiwidth is too early](https://github.com/vim/vim/commit/fc8f1118e7ccd65b0537e47e6323c93a1df6a61a)
- [8.0.0568 : 1gd may hang](https://github.com/vim/vim/commit/60402d68da09997cacdeec71fd22c9344f8f40d5)
- [8.0.0569 : bracketed paste is still enabled in a shell command](https://github.com/vim/vim/commit/62cf09b5dcb9a04b8e4f5e6cf3b14a787758e1a2)
- [8.0.0570 : can't run make with several jobs](https://github.com/vim/vim/commit/c03944151fec2d5469746642a949c8fcb5a744a8)
- [8.0.0571 : negative line number when using :z&#x5e; in an empty buffer](https://github.com/vim/vim/commit/a364cdb648ae009fa7aa05382f5659335683d349)
- [8.0.0572 : building the command table requires Perl](https://github.com/vim/vim/commit/6de5e126018b6f92526795cc06b1d73fac965db1)
- [8.0.0573 : running parallel make after distclean fails](https://github.com/vim/vim/commit/91b6e4591a839139f720798a2efbe83b2ecf8446)
- [8.0.0574 : get only one quickfix list after :caddbuf](https://github.com/vim/vim/commit/99895eac1cf71be43ece7e14b50e206e041fbe9f)
- [8.0.0575 : using freed memory when resetting 'indentexpr'](https://github.com/vim/vim/commit/a701b3b6f0f06ac0c9fcc75c6c34a1258fc3b1a2)
- [8.0.0576 : can't build when configure choses "install-sh"](https://github.com/vim/vim/commit/1d4be82c7d794e62bcf892310fe2a9b1eb5e2fec)
- [8.0.0577 : warning for uninitialized variable](https://github.com/vim/vim/commit/97db5541a65e4614f8753af27996c204946061a9)
- [8.0.0578 : :simalt on MS-Windows does not work properly](https://github.com/vim/vim/commit/7a85b0f028e0960ba70650ab462a9b6cca32139d)
- [8.0.0579 : duplicate test case for quickfix](https://github.com/vim/vim/commit/9b77016545d5ef1a1f4a90c9bb4b7a6693af8918)
- [8.0.0580 : cannot set the valid flag with setqflist()](https://github.com/vim/vim/commit/f1d21c8cc83f40c815b6bf13cd2043152db533ee)
- [8.0.0581 : moving folded text is sometimes not correct](https://github.com/vim/vim/commit/94be619e30e82d28cadeea5e0766c6f5c321ff8b)
- [8.0.0582 : illegal memory access with z= command](https://github.com/vim/vim/commit/5b276aa80e112ae1993bd43e28f599f257827c54)
- [8.0.0583 : fold test hangs on MS-Windows](https://github.com/vim/vim/commit/b11c826ddc459813f9f991cdb8e8736b686a6328)
- [8.0.0584 : memory leak when executing quickfix tests](https://github.com/vim/vim/commit/d788f6fe89c77262c474de323f5dab6d1c814e27)
- [8.0.0585 : test&#x5f;options fails when run in the GUI](https://github.com/vim/vim/commit/86e5792906621be5a6c4a1de9b8852e9caf65ea1)
- [8.0.0586 : no test for mapping timing out](https://github.com/vim/vim/commit/b7637c44c26b057d1f3721d932bbab06d9f74393)
- [8.0.0587 : configure check for return value of tgetent skipped](https://github.com/vim/vim/commit/696cbd224b4f78b99b8fc34538867303a8a922fc)
- [8.0.0588 : job&#x5f;stop() often assumes the channel will be closed](https://github.com/vim/vim/commit/1a9020d4cba6d90688506102ab587c371ff72f8b)

## 新着スクリプト

- [vim-custom-surround :  Define your own surround patterns ](http://www.vim.org/scripts/script.php?script_id=5547)
- [js&#x5f;alternate : A JavaScript plugin to easily navigate between alternate files](http://www.vim.org/scripts/script.php?script_id=5548)
- [tstatus : Monitor buffer-local options and variables in the statusline](http://www.vim.org/scripts/script.php?script_id=5549)
- [autoproject : Set buffer-local options, the working directory etc. specific to a project](http://www.vim.org/scripts/script.php?script_id=5550)
- [vikibase : A slimmed-down version of the viki plugin: a personal wiki](http://www.vim.org/scripts/script.php?script_id=5551)
- [foldtext : Folding support for some filetypes that don't do so by default](http://www.vim.org/scripts/script.php?script_id=5552)
- [vim-daps : Adds daps (https://opensuse.github.io/daps) functionality to the Vim text editor](http://www.vim.org/scripts/script.php?script_id=5553)
- [SumOfSelection : Calculate sum of visual (block) selected numbers](http://www.vim.org/scripts/script.php?script_id=5554)
- [Multiplayer : Collaborative Editing](http://www.vim.org/scripts/script.php?script_id=5555)
- [Multiplayer : Collaborative Editing](http://www.vim.org/scripts/script.php?script_id=5556)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (1317)
2. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (1061)
3. [wombat256.vim : Wombat for 256 color xterms](http://www.vim.org/scripts/script.php?script_id=2465) (554)
4. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (518)
5. [c.vim : C/C++ IDE -- Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (479)
6. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (436)
7. [nginx.vim : initial version](http://www.vim.org/scripts/script.php?script_id=1886) (412)
8. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (412)
9. [bufexplorer.zip : Buffer Explorer / Browser](http://www.vim.org/scripts/script.php?script_id=42) (395)
10. [winmanager : A windows style IDE for Vim 6.0](http://www.vim.org/scripts/script.php?script_id=95) (380)

## vim-jp issues

Open : 251 (-4) | Closed : 796 (+10)

- [Issue #1042 : screen が invalid 状態のとき，'w0', 'w$' の Position の返す値がおかしい](https://github.com/vim-jp/issues/issues/1042)
- [Issue #1043 : Gnome デスクトップで winpos が効かない?](https://github.com/vim-jp/issues/issues/1043)
- [Issue #1044 : Gnome デスクトップで winpos が効かない?](https://github.com/vim-jp/issues/issues/1044)
- [Issue #1045 : 特定の文字列を編集して k で移動するとターミナルごと落ちる](https://github.com/vim-jp/issues/issues/1045)
- [Issue #1046 : Windows のコンソール版が色鮮やかになる](https://github.com/vim-jp/issues/issues/1046)
- [Issue #1047 : あいまい幅文字の表示がおかしい](https://github.com/vim-jp/issues/issues/1047)
