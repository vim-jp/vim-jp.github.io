---
layout: post
category: blog
title: Visual モード時の CTRL-A/CTRL-X について
---

Vim をお使いの方であれば、数字の上で CTRL-A や CTRL-X を押すとインクリメント/デクリメントされる事はご存じだと思います。

[Patch 7.4.754](https://groups.google.com/forum/#!topic/vim_dev/5u4nnnUbA70) により Visual モードで選択している数字を一括でインクリメント/デクリメント出来る様になりました。

ビジュアル選択した状態で `<CTRL-A>` や `<CTRL-X>` をタイプすると選択範囲内の数値が一括でインクリメント/デクリメント出来ます。もちろん `20<CTRL-A>` で一気に20インクリメントする事も出来ます。

![](http://vim-jp.org/assets/images/visual-ctrl-a-ctrl-x.gif)

単純に `<CTRL-A>`, `<CTRL-X>` をタイプした場合、アクション後にビジュアル選択が解けてしまうので、続けて操作したいという場合は

```vim
vnoremap <c-a> <c-a>gv
vnoremap <c-x> <c-x>gv
```

この様に vimrc に書いておけば連続で `<CTRL-A>`, `<CTRL-X>` をタイプする事も出来ます。
