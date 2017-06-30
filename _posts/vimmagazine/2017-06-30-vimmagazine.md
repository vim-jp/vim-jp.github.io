---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2017 年 6 月号

---

## 話題

*   [Meguro.vim #4](https://megurovim.connpass.com/event/57955/) が 07/15 に開催されます。既に定員は超えているようですが、キャンセルや増員がないとも限らないので注目です。
*   [Fabric.vim #2](https://fablicvim.connpass.com/event/60339/) が 07/20 に開催されます。こちらはまだ定員に余裕があります。

## 今月の新機能

今月の新機能及びユーザーに影響のある変更は以下のとおりです。

*   8.0.0609: 終了方法の説明が改善されました
*   8.0.0612: パッケージディレクトリが `'runtimepath'` に追加されるタイミングが速まりました
*   8.0.0613: `:setfiletype` コマンドに `FALLBACK` オプションが追加されました。これはファイルタイプの推測のために使用され、具体的には `did_filetype()` が true ではなく false を返します。
*   8.0.0621: `:stag` コマンドが `'switchbuf'` オプションを尊重するようになりました
*   8.0.0625: `shellescale()` が第2引数に応じて、改行をエスケープするようになりました

*   ランタイムの更新

## Vimに関する脆弱性

特筆すべき脆弱性の報告はありませんでした。
