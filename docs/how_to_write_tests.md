---
layout: docs
title: テストの書き方
---

それでは、具体的なテストを参照しながらテストの作り方を学んでいくことにしましょう。

次に紹介するのは私のパッチに対して、 `h_east` さんに実際に書いてもらったテストコードになります。
このテストコードでは、getbufvar(), getwinvar(), gettabvar() and gettabwinvar() に追加された機能をテストします。
getbufvar(), getwinvar(), gettabwinvar() にはそもそもテストが無かったため、これらの関数のテストにもなっています。

    Tests for getbufvar(), getwinvar(), gettabvar() and gettabwinvar().
    vim: set ft=vim :

    STARTTEST
    :so small.vim
    :"
    :" test for getbufvar()
    :let b:var_num = 1234
    :let def_num = 5678
    :$put =string(getbufvar(1, 'var_num'))
    :$put =string(getbufvar(1, 'var_num', def_num))
    :$put =string(getbufvar(1, ''))
    :$put =string(getbufvar(1, '', def_num))
    :unlet b:var_num
    :$put =string(getbufvar(1, 'var_num', def_num))
    :$put =string(getbufvar(1, ''))
    :$put =string(getbufvar(1, '', def_num))
    :$put =string(getbufvar(9, ''))
    :$put =string(getbufvar(9, '', def_num))
    :unlet def_num
    :$put =string(getbufvar(1, '&autoindent'))
    :$put =string(getbufvar(1, '&autoindent', 1))
    :"
    :" test for getwinvar()
    :let w:var_str = "Dance"
    :let def_str = "Chance"
    :$put =string(getwinvar(1, 'var_str'))
    :$put =string(getwinvar(1, 'var_str', def_str))
    :$put =string(getwinvar(1, ''))
    :$put =string(getwinvar(1, '', def_str))
    :unlet w:var_str
    :$put =string(getwinvar(1, 'var_str', def_str))
    :$put =string(getwinvar(1, ''))
    :$put =string(getwinvar(1, '', def_str))
    :$put =string(getwinvar(9, ''))
    :$put =string(getwinvar(9, '', def_str))
    :$put =string(getwinvar(1, '&nu'))
    :$put =string(getwinvar(1, '&nu',  1))
    :unlet def_str
    :"
    :" test for gettabvar()
    :tabnew
    :tabnew
    :let t:var_list = [1, 2, 3]
    :let def_list = [4, 5, 6, 7]
    :tabrewind
    :$put =string(gettabvar(3, 'var_list'))
    :$put =string(gettabvar(3, 'var_list', def_list))
    :$put =string(gettabvar(3, ''))
    :$put =string(gettabvar(3, '', def_list))
    :tablast
    :unlet t:var_list
    :tabrewind
    :$put =string(gettabvar(3, 'var_list', def_list))
    :$put =string(gettabvar(9, ''))
    :$put =string(gettabvar(9, '', def_list))
    :$put =string(gettabvar(3, '&nu'))
    :$put =string(gettabvar(3, '&nu', def_list))
    :unlet def_list
    :tabonly
    :"
    :" test for gettabwinvar()
    :tabnew
    :tabnew
    :tabprev
    :split
    :split
    :wincmd w
    :vert split
    :wincmd w
    :let w:var_dict = {'dict': 'tabwin'}
    :let def_dict = {'dict2': 'newval'}
    :wincmd b
    :tabrewind
    :$put =string(gettabwinvar(2, 3, 'var_dict'))
    :$put =string(gettabwinvar(2, 3, 'var_dict', def_dict))
    :$put =string(gettabwinvar(2, 3, ''))
    :$put =string(gettabwinvar(2, 3, '', def_dict))
    :tabnext
    :3wincmd w
    :unlet w:var_dict
    :tabrewind
    :$put =string(gettabwinvar(2, 3, 'var_dict', def_dict))
    :$put =string(gettabwinvar(2, 3, ''))
    :$put =string(gettabwinvar(2, 3, '', def_dict))
    :$put =string(gettabwinvar(2, 9, ''))
    :$put =string(gettabwinvar(2, 9, '', def_dict))
    :$put =string(gettabwinvar(9, 3, ''))
    :$put =string(gettabwinvar(9, 3, '', def_dict))
    :unlet def_dict
    :$put =string(gettabwinvar(2, 3, '&nu'))
    :$put =string(gettabwinvar(2, 3, '&nu', 1))
    :tabonly
    :"
    :/^start/,$wq! test.out
    ENDTEST
    
    start:
    iff -r b89e2bdcc6e5 src/testdir/test89.ok
    -- /dev/null	Thu Jan 01 00:00:00 1970 +0000
    ++ b/src/testdir/test89.ok	Wed Jan 02 09:01:40 2013 +0900
    @ -0,0 +1,45 @@
    start:
    1234
    1234
    {'var_num': 1234}
    {'var_num': 1234}
    5678
    {}
    {}
    ''
    ''
    0
    0
    'Dance'
    'Dance'
    {'var_str': 'Dance'}
    {'var_str': 'Dance'}
    'Chance'
    {}
    {}
    ''
    ''
    0
    0
    [1, 2, 3]
    [1, 2, 3]
    ''
    ''
    [4, 5, 6, 7]
    ''
    ''
    ''
    ''
    {'dict': 'tabwin'}
    {'dict': 'tabwin'}
    {'var_dict': {'dict': 'tabwin'}}
    {'var_dict': {'dict': 'tabwin'}}
    {'dict2': 'newval'}
    {}
    {}
    ''
    ''
    ''
    ''
    0
    0
