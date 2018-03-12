---
layout: post
category: blog
title: 最近追加された機能
---
最近追加された機能について説明します。

# terminal バッファ強制終了の動作

これまで GVim で `:terminal` コマンドを実行した状態で画面の × ボタンをクリックすると、エラーが発生して GVim を終了する事が出来ませんでした。これは Vim の設計方針としてバッファを閉じる際にジョブを強制終了するかどうかが明確に決まっていなかった事が原因です。

[Patch 8.0.1593](https://github.com/vim/vim/commit/25cdd9c33b21ddbd31321c075873bb225450d2d2) にて × ボタンを押した時の動作 `:qall` に変更が加えられました。

バッファを閉じる際にはジョブを自動で終了するべきではないというスタンスには変わりはありませんが、× ボタンを押した場合のみ、今回追加された `term_setkill` 関数により内部的にジョブを自動終了する為のフラグ付けが行われます。

更新されたバッファもしくは実行中の :terminalが開かれている状態で、× ボタンをクリックすると、全てを破棄するかを確認されます。全てを破棄を選択すると :terminal も閉じられます。

# TerminalOpen イベント

また [Patch 8.0.1596](https://github.com/vim/vim/commit/b852c3e64d319d6ec47dd780c8654ae095e1d8c2) にて TerminalOpen イベントも導入されました。 :terminal が開かれた際に固有の設定を施したい方は以下の様に設定する事が出来る様になりました。

```vim

function! s:my_terminal()
  " :terminal バッファ固有の設定
endfunction

augroup MyTerminal
  au!
  au TerminalOpen * call s:my_terminal()
augroup END
```

# ExitPre イベント

Vim プラグイン開発者はこれまで Vim の終了を検知する為に QuitPre イベントを使ってきました。しかしこの QuitPre イベントは、バッファの終了を確認する手前で発火される為、バッファに割り当てた channel を閉じたいといったニーズに使用するには少し無理がありました。

[Patch 8.0.1595](https://github.com/vim/vim/commit/12a96de430779b88795fac87a2be666d9f661d1e) にて ExitPre イベントが導入され、バッファが全て閉じられる事が確定した後に発火するイベントが追加されました。Vim プラグインから Vim が ExitPre イベントをサポートしているかどうかを確認するには以下の様に実行します。

```vim
:echo exists("##ExitPre")
```

良く似たイベントに QuitPre、VimLeavePre、VimLeave があります。発火するタイミングは以下の様になります。

|イベント   |発火するタイミング      |
|-----------|------------------------|
|QuitPre    |終了するかも(しないかも)|
|ExitPre    |今からバッファを閉じる  |
|VimLeavePre|バッファを閉じた後      |
|VimLeave   |終了する直前            |

# terminal セッションの復帰

Vim では `:mksession` コマンドを使う事で、編集中のバッファの状態を保存し、次回 Vim を起動した際に復帰できる機能があります。[Patch 8.0.1592](https://github.com/vim/vim/commit/4d8bac8bf593ff087517ff79090c2d224325aae6) にて sessionoptions オプションに terminal という値を含める事が出来る様になりました。これによりセッションを読み込んだ際に terminal ウィンドウも復帰されます。

※起動していたプロセスの状態などは復帰しません
