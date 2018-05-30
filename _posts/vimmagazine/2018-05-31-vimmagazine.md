---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2018 年 05 月号

---

## 話題

*   [vim-jp.org](https://vim-jp.org) が HTTPS になりました
*   Vim 8.1 が 05/18(金) に[リリース](https://vim-jp.org/blog/2018/05/18/Vim-8.1-released.html)されました
*   [asakusa.vim #0](https://asakusavim.connpass.com/event/88381/) が 05/24(木) に開催されました

## 今月の新機能

今月の新機能及びユーザーに影響のある変更は以下のとおりです。

*   8.0.1782: `getqflist()` と `setqflist()` で取り扱う辞書型のキーに `"module"` が追加されました
*   8.0.1783: Windows のコンソールで256色が使えるようになりました
*   8.0.1786: コマンドラインにおいて `<C-R><C-L>` でバッファのカーソルのある行全体の内容を貼り付けられるようになりました
*   8.0.1792: Windows ではオプション `-?` でもヘルプを表示できるようになりました
*   8.0.1832: 環境変数に対して `:unlet` ができるようになりました
*   8.1.0000: Vim 8.1 としてリリースされました!
*   8.1.0020: マクロ記録中と実行中のレジスタを知るための関数 `reg_recording()` と `reg_executing()` が追加されました

*   ランタイムの更新
    *   `debcopyright` ファイルタイプ及びシンタックスが追加されました
    *   [chrisbra/vim-xml-ftplugin](https://github.com/chrisbra/vim-xml-ftplugin) がバンドルされました
        *   フォーマッタについての詳細は `:help format-formatexpr` も参照してください

## Vimに関する脆弱性

特筆すべき脆弱性の報告はありませんでした。
