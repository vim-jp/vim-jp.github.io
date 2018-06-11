---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2018 年 06 月号

---

## 話題

* [asakusa.vim #1](https://asakusavim.connpass.com/event/90854/) が 2018-06-21 (木) に開催され、Vim使いたちの貴重な知見が共有されました。
* [VimConf 2018のスポンサー公募](https://vim-jp.org/blog/2018/06/07/VimConf2018-sponsor-ja.html) がはじまりました。

## 今月の新機能

今月の新機能及びユーザーに影響のある変更は以下のとおりです。

*   8.1.0011: `maparg()` 関数と `mapcheck()` 関数が rhs が `<Nop>` であるマッピングに対して `"<Nop>"` という文字列を返すようになりました (今までは rhs が `<Nop>` あるいは存在しないマッピング両方に対して空文字列を返していました)
*   8.1.0027: プロンプトバッファ (`set buftype=promopt`) が追加されました (関連: 8.1.0028, 8.1.0032, 8.1.0035-0036)
    *   関数 `prompt_addtext()`, `prompt_setprompt()`, `prompt_setcallback()` が追加されました
*   8.1.0037: 任意のバッファに行を追記できる関数 `appendbufline()` が追加されました
*   8.1.0039: 任意のバッファの行を削除できる関数 `deletebufline()` が追加されました

*   ランタイムの更新
    *   PHP のインデントスクリプトに `g:PHP_noArrowMatching`, `b:PHP_noArrowMatching` 変数が追加されました (`:help PHP_noArrowMatching`)。1 にするとメソッド呼び出しの `->` の位置を揃えないようになります

## Vimに関する脆弱性

特筆すべき脆弱性の報告はありませんでした。
