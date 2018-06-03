---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2018 年 06 月号

---

## 話題

とくに大きな話題はありません。

## 今月の新機能

今月の新機能及びユーザーに影響のある変更は以下のとおりです。

*   8.1.0011: `maparg()` 関数と `mapcheck()` 関数が rhs が `<Nop>` であるマッピングに対して `"<Nop>"` という文字列を返すようになりました (今までは rhs が `<Nop>` あるいは存在しないマッピング両方に対して空文字列を返していました)
*   8.1.0027: プロンプトバッファ (`set buftype=promopt`) が追加されました (関連: XXX) (TODO: 8.1.0027 では初期実装が入ったのみ)
    *   `prompt_addtext({buf}, {expr})`
    *   `prompt_setprompt({buf}, {text})`
    *   `prompt_setcallback({buf}, {expr})`

*   ランタイムの更新
    *   PHP のインデントスクリプトに `g:PHP_noArrowMatching`, `b:PHP_noArrowMatching` 変数が追加されました (`:help PHP_noArrowMatching`)。1 にするとメソッド呼び出しの `->` の位置を揃えないようになります

## Vimに関する脆弱性

特筆すべき脆弱性の報告はありませんでした。
