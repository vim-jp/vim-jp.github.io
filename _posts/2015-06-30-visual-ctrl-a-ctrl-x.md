---
layout: post
category: blog
title: Visual モード時の CTRL-A/CTRL-X について
---

Vim をお使いの方であれば、数字の上で CTRL-A や CTRL-X を押すとインクリメント/デクリメントされる事はご存じだと思います。

[Patch 7.4.754](https://groups.google.com/forum/#!topic/vim_dev/5u4nnnUbA70) により Visual モードで選択している数字を一括でインクリメント/デクリメント出来る様になりました。

もちろん `20<CTRL-A>` で一気に20インクリメントする事も出来ます。

![](http://vim-jp.org/assets/images/visual-ctrl-a-ctrl-x.gif)

単純に `<CTRL-A>`, `<CTRL-X>` をタイプした場合、アクション後にビジュアル選択が解けてしまうので、続けて操作したいという場合は

```vim
vnoremap <c-a> <c-a>gv
vnoremap <c-x> <c-x>gv
```
*Windows だと `vunmap <C-X>` が必要*

この様に vimrc に書いておけば連続で `<CTRL-A>`, `<CTRL-X>` をタイプする事も出来ます。詳しくは `:help v_CTRL-A` を参照下さい。

また `g<C-A>` と `g<C-X>` を使う事で簡単に連番を付与する事が出来るようになりました。

```
1.
1.
1.
1.
```

この様なテキストの2行目以降を Visual 選択し、`g<C-A>` をタイプすると

```
1.
2.
3.
4.
```

一気に連番が出来上がりました。便利ですね。詳しくは `:help v_g_CTRL-A` を参照下さい。
