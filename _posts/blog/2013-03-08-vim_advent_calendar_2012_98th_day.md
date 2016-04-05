---
layout: blog
category: blog
title: カラースキームを360倍使うためのスクリプト
---

[Vim Advent Calendar 2012](http://atnd.org/events/33746) 98日目の記事です。

みなさんお気に入りのカラースキームはありますか。もう長いこと同じものを使っている方も多いのではないでしょうか。しかしずっと同じ色だと飽きますよね。たまには色を変えて気分をリフレッシュしたいことでしょう。とはいえ、星の数ほどあるカラースキームの中から気に入ったものを探し出すのは大変です。自分でカスタマイズしたり一から作るのは面倒だしなによりセンスがない。そんなあなたにぴったりのスクリプトを紹介します。

[huerotation.vim](http://www.vim.org/scripts/script.php?script_id=2283) (GUI only)

Hue (色相) を Rotation (回転) させるスクリプトです。

お絵描きソフトの色選択ダイアログでこういうのを見たことがあると思いますが、

![](/assets/images/post-vac2012-98-gimp.png)

この虹色の円の部分が色相です。これをくるくる回すようなイメージです。

使い方は簡単。

スクリプトを読み込んで:

    :source huerotation.vim

色相を回転させる (例: 30°):

    :call RotateHue(30)

![](/assets/images/post-vac2012-98-gvim1.png)
![](/assets/images/post-vac2012-98-gvim2.png)

360°回転させると元の色に戻ります。

センスのいいカラースキームはそれぞれの色の相関関係が優れています。そのいいセンスを保ったまま全体的に色を変えることができるわけです。

というわけでたまには気分を変えてプログラミングをしてみてはいかがでしょうか。

ちなみに私が書いたものではありませんが huerotation.vim よりもっと細かく調整できる [toner.vim](http://www.vim.org/scripts/script.php?script_id=3417) というものもあります。

-- Yukihiro Nakadaira

