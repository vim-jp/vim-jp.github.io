---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2018 年 04 月号

---
## 話題

*   [Meguro.vim #9](https://megurovim.connpass.com/event/81678/) が 04/07 (土) に開催されました

## 今月の新機能

今月の新機能及びユーザーに影響のある変更は以下のとおりです。

*   8.0.1656: `xxd` コマンドに `-i` オプション使用時に C 変数をキャピタライズする `-C` オプションが追加されました (関連: 8.0.1658)
*   8.0.1660: 端末 API の `drop` コマンドが追加されました
*   8.0.1665: 端末を起動する時 GUI 版では `'term'` オプションの代わりに元の `TERM` 環境変数の値が端末の `TERM` 環境変数として渡されるようになりました
*   8.0.1668: 端末デバッガに `:Source` コマンドが追加されました
*   8.0.1685: `term_setansicolors()`, `term_getansicolors()`, `g:terminal_ansi_colors` が追加され、端末の ANSI カラーを設定できるようになりました
*   8.0.1704: Mac の `'backupskip'` オプションのデフォルト値が `/private/tmp` に変更されました
*   8.0.1708: `mkdir()` 関数に `p` フラグを渡した時、ディレクトリが存在した場合にエラーを出さなくなりました
*   8.0.1711: 端末のサイズを変更する `term_setsize()` 関数が追加されました (関連: 8.0.1714)
*   8.0.1712: 端末のスクロールバック可能な行数を制限する `'terminalscroll'` オプションが追加されました (デフォルトは 10000 行) (関連: 8.0.1715)
*   8.0.1719: Python の実行ファイルをビルド時に指定可能な `--with-python-command` と `--with-python3-command` が追加されました
*   8.0.1725: 端末デバッガに `:TermdebugCommand` コマンドが追加されました
*   8.0.1734: パッケージディレクトリ名の先頭が他のパッケージディレクトリ名にマッチした場合読み込まれない不具合が解消されました
*   8.0.1738: `:args` コマンドの出力が複数行の場合、余白付きで表示されるようになりました
*   8.0.1742: `job_info()` 関数の引数に何も渡されなかった場合、全てのジョブの情報が取得できるようになりました
*   8.0.1743: 端末ウィンドウに関するオプション名は `termwin` が先頭に付くようになりました (`termkey` → `termwinkey`, `termsize` → `termwinsize`, `terminalscroll` → `termwinscroll`)
*   8.0.1747: `term_start()` 関数が `job_info()` 関数で取得できる情報を設定するようになりました

## Vimに関する脆弱性

特筆すべき脆弱性の報告はありませんでした。

## リリース情報

- [8.0.1656 : no option to have xxd produce upper case variable names](https://github.com/vim/vim/commit/8b31a6ff7a0a8184ddc9997a0cc4aa5f2350b279)
- [8.0.1657 : crash when reading a channel](https://github.com/vim/vim/commit/9af9778209cd0cbb284f6ff7e88f707418089975)
- [8.0.1658 : xxd capitalize argument not available in long form](https://github.com/vim/vim/commit/79cf7c0d4574f15cfad5ad706f175e226c528f2a)
- [8.0.1659 : scroll events not recognized for some xterm emulators](https://github.com/vim/vim/commit/1f8495cf48888ff3a8ab3d0eb22db92389191a60)
- [8.0.1660 : the terminal API "drop" command doesn't support options](https://github.com/vim/vim/commit/333b80acf3a44e462456e6d5730e47ffa449c83d)
- [8.0.1661 : warnings from 64 bit compiler](https://github.com/vim/vim/commit/878c96d5b9416170dfd28a02cba0db683f91c220)
- [8.0.1662 : showing dump diff doesn't mention both file names](https://github.com/vim/vim/commit/4a69634b1b55e06c4bf7f05b54125b1669b1c363)
- [8.0.1663 : cannot build without multi-byte feature](https://github.com/vim/vim/commit/86676c949fcec939441620bc747851e315e2702a)
- [8.0.1664 : test failure because of not allocating enough space](https://github.com/vim/vim/commit/3aa67fb453373109132b71eeb488595d651dfb7d)
- [8.0.1665 : when running a terminal from the GUI 'term' is not useful](https://github.com/vim/vim/commit/9a993e3c09371bb80d71be62fca53cf954a98f72)
- [8.0.1666 : % argument in ch&#x5f;log() causes trouble](https://github.com/vim/vim/commit/d5359b24ceefefed4b3a37efd9bf54176935f3a5)
- [8.0.1667 : terminal window tests are flaky](https://github.com/vim/vim/commit/3680bab7e797444a7f41646888d9dd04b01c3816)
- [8.0.1668 : terminal debugger: can't re-open source code window](https://github.com/vim/vim/commit/c4b533e1e93151658cb170c6796d327a8c0f8612)
- [8.0.1669 : :vimgrep may add entries to the wrong quickfix list](https://github.com/vim/vim/commit/e1bb879f49665bb828197135b80aaf72cc190073)
- [8.0.1670 : terminal window tests are still a bit flaky](https://github.com/vim/vim/commit/b6ed74fef8e088128c45928ee7d02458b2882c1e)
- [8.0.1671 : crash when passing non-dict argument as env to job&#x5f;start()](https://github.com/vim/vim/commit/22efba47deb7e504142487fdc1617ec638cf0683)
- [8.0.1672 : error during completion causes command to be cancelled](https://github.com/vim/vim/commit/72532d354e699f1cceec34c0b08e1de4d3ea9641)
- [8.0.1673 : terminal window tests are still a bit flaky](https://github.com/vim/vim/commit/4791015e6f0adf7f3a0a6a59884c4092ca3c19ef)
- [8.0.1674 : libvterm can't handle an OSC string split](https://github.com/vim/vim/commit/77557a7ae66830cb21c79d3a2b48a93b086599b3)
- [8.0.1675 : unused macro argument in libvterm](https://github.com/vim/vim/commit/820082914122d5bf7b6e89ff30cf364ad40a560f)
- [8.0.1676 : no compiler warning for wrong printf format](https://github.com/vim/vim/commit/4ac2e8d8e60dcc7dbff662e177b86ccfbda7cd9e)
- [8.0.1677 : no compiler warning for wrong format in vim&#x5f;snprintf()](https://github.com/vim/vim/commit/ea39176baab52b646d1e2676e662def718ddd365)
- [8.0.1678 : errorformat "%r" implies "%>"](https://github.com/vim/vim/commit/e333e79f9bdff82432e0fd7fcf7ae30ef8e3d092)
- [8.0.1679 : compiler warning for printf format](https://github.com/vim/vim/commit/672afb9f66c64e031a2b609bdee0cb873883c9ec)
- [8.0.1680 : memory allocated by libvterm is not profiled](https://github.com/vim/vim/commit/756ef113d14428e598274f87672d7f0e34ff9781)
- [8.0.1681 : the format attribute fails with MinGW](https://github.com/vim/vim/commit/e80757c1545286240d687e9a303cf8eeb3f9a6de)
- [8.0.1682 : auto indenting breaks inserting a block](https://github.com/vim/vim/commit/8c87a2b1fec85e4aac33f71586ac1514536fc66b)
- [8.0.1683 : Python upgrade breaks Vim when defining PYTHON&#x5f;HOME](https://github.com/vim/vim/commit/0424958bde3d3e74c721ba39ab4d5a4744897393)
- [8.0.1684 : ml&#x5f;get errors when using terminal window for shell command](https://github.com/vim/vim/commit/07b46af645fa7a2b95afec3f3d87c91942d9cbd0)
- [8.0.1685 : can't set ANSI colors of a terminal window](https://github.com/vim/vim/commit/f59c6e8cee092433d325ba21a107654a8d84f776)
- [8.0.1686 : Python does not work when configuring with specific dir](https://github.com/vim/vim/commit/d088240b96bc8b2bdba9242dcb9e3d36f1e16c48)
- [8.0.1687 : 64 bit compiler warnings](https://github.com/vim/vim/commit/d6b4f2dd760b425c63c359b9ff299f73448f8854)
- [8.0.1688 : some macros are used without a semicolon](https://github.com/vim/vim/commit/6f4700233fd925fe122b851f937929fb5e5da707)
- [8.0.1689 : no tests for xxd](https://github.com/vim/vim/commit/29f9ed20b42e7badea6cfb82fb1fb8d71f33e282)
- [8.0.1690 : not easy to run one test with gvim instead of vim](https://github.com/vim/vim/commit/8dd1761e69d70bb4b5848dc4c213792d9121787a)
- [8.0.1691 : xxd test sometimes fails](https://github.com/vim/vim/commit/164268d4cf84fa0efd984fddbef4a41366aabd46)
- [8.0.1692 : Python may not work when using statically linked library ](https://github.com/vim/vim/commit/3a21d9c7bb4d58cccd2152e319daa8f8ffa2e664)
- [8.0.1693 : xxd is excluded from coverage statistics](https://github.com/vim/vim/commit/1acdb3f30104576c36f83d422fb67a805b0ce58a)
- [8.0.1694 : terminal API test is a bit flaky](https://github.com/vim/vim/commit/b377457e693210d1b0c0bbac7e04671fc98f038b)
- [8.0.1695 : xxd test not run on MS-Windows](https://github.com/vim/vim/commit/6995c0a7a91986caa0b495dd926d440c1e99f651)
- [8.0.1696 : coverage statistics don't work](https://github.com/vim/vim/commit/d60e0a1d202d49503ba95a40c165172c8f61db7a)
- [8.0.1697 : various tests are still a bit flaky](https://github.com/vim/vim/commit/769e9d21ac3e8dff43b9ef5e46cdc4523833b51e)
- [8.0.1698 : coverage statistics don't work on coveralls](https://github.com/vim/vim/commit/1ed00330bcd169d5fcd72608e2745612b5b1d5b5)
- [8.0.1699 : leftover stuff for Python 1.4](https://github.com/vim/vim/commit/9ce4213b0322477ace1df5e0aa356b1a5551b445)
- [8.0.1700 : coverage statistics still don't work on coveralls](https://github.com/vim/vim/commit/56f99550444949e12eca4a31ee60be91b2ab3a6e)
- [8.0.1701 : can disable COLOR&#x5f;EMOJI with MSVC but not MinGW](https://github.com/vim/vim/commit/cccd62d435097814ea93e384920443fb2419e663)
- [8.0.1702 : leaking memory when autocommands make quickfix list invalid](https://github.com/vim/vim/commit/ee5b94a8ecbaeafa11c7c7c8cf7bef765fbea912)
- [8.0.1703 : in the tutor 'showcmd' is not set](https://github.com/vim/vim/commit/527dec3f94c72096bf6baf4bb71f1a21dc7df89c)
- [8.0.1704 : 'backupskip' default doesn't work for Mac](https://github.com/vim/vim/commit/b8e22a053b68774dfd86fd829d9dba2333f09c05)
- [8.0.1705 : when making a vertical split the mode message isn't updated](https://github.com/vim/vim/commit/5bab555c2f1b3b86d57e4adeb86d908eff477fc9)
- [8.0.1706 : cannot sent CTRL-\\ to a terminal window](https://github.com/vim/vim/commit/b59118d5012a6a07e2d54504fbb04837b5a1208c)
- [8.0.1707 : when 'wfh' is set ":bel 10new" scrolls window](https://github.com/vim/vim/commit/98da6ecab905df48a67da36ce60233f45726c979)
- [8.0.1708 : mkdir with 'p' flag fails on existing directory](https://github.com/vim/vim/commit/78a16b0f2a142aae1fdc96c50ab0f25194d0e755)
- [8.0.1709 : some non-C89 code may slip through](https://github.com/vim/vim/commit/2e324950b83fcdf60843b54a6a339183370f338a)
- [8.0.1710 : building with Ruby fails](https://github.com/vim/vim/commit/2a43230ce39eea340aab15fb50a083bc527fb8d0)
- [8.0.1711 : term&#x5f;setsize() is not implemented yet](https://github.com/vim/vim/commit/a42d363bac8a581afe769c370db70cf833767c41)
- [8.0.1712 : terminal scrollback is not limited](https://github.com/vim/vim/commit/8c041b6b95f49f7383cf00e2036cf009b326fa8d)
- [8.0.1713 : terminal debugger doesn't handle arguments](https://github.com/vim/vim/commit/b3623a382abc8f8e2bcfba4c1a2aa06b7578eb8d)
- [8.0.1714 : term&#x5f;setsize() does not give an error in a normal buffer](https://github.com/vim/vim/commit/6e72cd0d7267b9545ef966a53e62706914a96042)
- [8.0.1715 : terminal buffer can be 1 more than 'terminalscroll' lines](https://github.com/vim/vim/commit/8c94a549051cc4d4cbb8cabd321724a85fe40c23)
- [8.0.1716 : test for term&#x5f;setsize() does not give a good error message](https://github.com/vim/vim/commit/ac3e830065f1e54c422cdd2f3157fb35fac27e04)
- [8.0.1717 : C89 check causes too much trouble](https://github.com/vim/vim/commit/3cba73423e5304ef7ccdb2280d65562d2f06f483)
- [8.0.1718 : terminal scrollback test fails on MS-Windows](https://github.com/vim/vim/commit/6c672194cd06fb1ab1289475aad31a9769bdb2ad)
- [8.0.1719 : cannot specify which Python executable configure should use](https://github.com/vim/vim/commit/e1a32310d5e9684608b893ab8295191d63178b4e)
- [8.0.1720 : when a timer is running a terminal window may not close](https://github.com/vim/vim/commit/802bfb14636b24d86d0ca8e0947d808b9b7c941e)
- [8.0.1721 : no test for using the 'termsize' option](https://github.com/vim/vim/commit/a7eef3d87fa36d527d1cfc749b400df1e69dcb3d)
- [8.0.1722 : cannot specify a minimal size for a terminal window](https://github.com/vim/vim/commit/498c2562e1bcc72492fe8da8a76504f893e9b5fe)
- [8.0.1723 : using one item array size declaration is misleading](https://github.com/vim/vim/commit/f3a411783c9736645d6ba480c5ff9d861164c040)
- [8.0.1724 : declarations cannot be halfway a block](https://github.com/vim/vim/commit/06965b838f3677251668dca726b07c5ad2499341)
- [8.0.1725 : terminal debugger doesn't handle command arguments](https://github.com/vim/vim/commit/32c67ba733abcd4f2d729595fdccb2295445636a)
- [8.0.1726 : older MSVC doesn't support declarations halfway a block](https://github.com/vim/vim/commit/df2c774ded4c2ed169390a352099c23b6db657c4)
- [8.0.1727 : qf&#x5f;get&#x5f;properties() function is too long](https://github.com/vim/vim/commit/353eeeaca269ed5e83900bd4a24dc6dc80bb4880)
- [8.0.1728 : condition always false, useless code](https://github.com/vim/vim/commit/1feab2b9cc355edc2f24814c35b72ae5e74eb2ae)
- [8.0.1729 : no comma after last enum item](https://github.com/vim/vim/commit/ea3ece405ab55f44018257bd2f5021231af8e87f)
- [8.0.1730 : no configure check for the used C99 features](https://github.com/vim/vim/commit/561f8a5a4612751c2a4ebd6bf918cbc3be867350)
- [8.0.1731 : characters deleted on completion](https://github.com/vim/vim/commit/e87edf3b85f607632e5431640071fdbc36b685b2)
- [8.0.1732 : crash when terminal API call deletes the buffer](https://github.com/vim/vim/commit/a997b45c7e350ea5b378ca0c52ed3d4cc610975c)
- [8.0.1733 : incomplete testing for completion fix](https://github.com/vim/vim/commit/bad0ce7b26be5eed8524347018f4c835b212f8d1)
- [8.0.1734 : package directory not added to 'rtp' if prefix matches](https://github.com/vim/vim/commit/f98a39ca57d001ba3e24831bae1e375790fb41f0)
- [8.0.1735 : flexible array member feature not supported by HP-UX](https://github.com/vim/vim/commit/285e3358696b1bc6296e5d4c53425680ce8fbd54)
- [8.0.1736 : check for C99 features is incomplete](https://github.com/vim/vim/commit/226400830b470774bf5a552e1af10706b609720c)
- [8.0.1737 : fchown() used when it is not supported](https://github.com/vim/vim/commit/02e802b2da9e25b5824185976b163691b5bbd558)
- [8.0.1738 : ":args" output is hard to read](https://github.com/vim/vim/commit/5d69da462f584a3aefb3427b127334bf9af3a4b0)
- [8.0.1739 : MS-Windows with msys2 cannot build Ruby statically](https://github.com/vim/vim/commit/5d69da462f584a3aefb3427b127334bf9af3a4b0)
- [8.0.1740 : warning for signed-unsigned incompatibility](https://github.com/vim/vim/commit/405dadb63ea2b7aa4c8c659807506a35a8a9504c)
- [8.0.1741 : MS-Windows with msys2 cannot build Ruby statically](https://github.com/vim/vim/commit/259a90f7ce4d6d983a7fe5461da88cc52c46d9fe)
- [8.0.1742 : cannot get a list of all the jobs](https://github.com/vim/vim/commit/e1fc51558dc14906a8d9f6a6e7bb1ac2a6ba8f3d)
- [8.0.1743 : terminal window options are named inconsistently](https://github.com/vim/vim/commit/6d150f783d5d3820fe69734dda1e79b8276a84d2)
- [8.0.1744 : on some systems /dev/stdout isn't writable](https://github.com/vim/vim/commit/9980b37a80dc72eef05bf8862aaf475ab17790a5)
- [8.0.1745 : build failure on MS-Windows](https://github.com/vim/vim/commit/2060892028e05b1325dc0759259254180669eb5e)
- [8.0.1746 : MS-Windows: channel tests fail](https://github.com/vim/vim/commit/a69b39511d769bc88f0854bd026556fdb9bd2817)
- [8.0.1747 : MS-Windows: term&#x5f;start() does not set job&#x5f;info() cmd](https://github.com/vim/vim/commit/ebe74b73677b06db7d483987a863b41cee051cc0)
- [8.0.1748 : CmdlineEnter command uses backslash instead of slash](https://github.com/vim/vim/commit/a4baf5b32519855bb176a7aa0e9397c137ca890a)
- [8.0.1749 : VMS: 100% CPU use, redefining mch&#x5f;open() fails](https://github.com/vim/vim/commit/6537c421bb626ee92a2487a543e6287d33461bd0)
- [8.0.1750 : crash when clearing location list in autocommand](https://github.com/vim/vim/commit/3b9474b4ad4d85b5396f7f641b436f193dc9d486)
- [8.0.1751 : #ifdef causes bad highlighting](https://github.com/vim/vim/commit/4e601e3c82899b3a8334856397946efeaa007f29)
- [8.0.1752 : qf&#x5f;set&#x5f;properties() is to long](https://github.com/vim/vim/commit/a2aa8a2b22de909619d7faa3ff5383a6224defc5)
- [8.0.1753 : various warnings from a static analyser](https://github.com/vim/vim/commit/1c17ffa4611f4efe68c61f7cdd9ed692a866ba75)
- [8.0.1754 : ex&#x5f;helpgrep() is too long](https://github.com/vim/vim/commit/2225ebb48644f3924311b8df02a1319ab7675d42)
- [8.0.1755 : MS-Windows: high unicode char received as two utf-16 words](https://github.com/vim/vim/commit/f1f2f837bc85c93ba81176383cccd6734f0129fb)
- [8.0.1756 : GUI: after prompting for a number the mouse shape is wrong](https://github.com/vim/vim/commit/73658317bacd9a0264dfaa32288de6ea1f236fe5)
- [8.0.1757 : unnecessary changes in libvterm](https://github.com/vim/vim/commit/b691de05f69905fe417f583083d7e3cc16eb865e)
- [8.0.1758 : open&#x5f;line() returns TRUE/FALSE for success/failure](https://github.com/vim/vim/commit/24a2d7264fc2394f69a9a890b34d1f09cf63b14e)
- [8.0.1759 : memory leak from duplicate options](https://github.com/vim/vim/commit/09d1d51df5d9b215e583b5bbe36df46afb3db35f)
- [8.0.1760 : wrong number of arguments to vms&#x5f;read()](https://github.com/vim/vim/commit/4994373c5db80f8fd57e802997ae430afa5ec9ec)
- [8.0.1761 : job in terminal window with no output channel is killed](https://github.com/vim/vim/commit/4e9d443a25b451e3f2de62e9eeea439aa4c3f039)
- [8.0.1762 : terminal debug logging is a bit complicated](https://github.com/vim/vim/commit/b255b90503a986931904c23dafb5b6d4e512a17e)
- [8.0.1763 : :argedit does not reuse an empty unnamed buffer](https://github.com/vim/vim/commit/46a53dfc29689c6a0d80e3820e8b0a48dba6b6ec)
- [8.0.1764 : lgtm considers tutor.es to be EcmaScript](https://github.com/vim/vim/commit/856905bd838fdf71be2c3dabc98453b85b80a83b)
- [8.0.1765 : CTRL-G j in Insert mode is incorrect when 'virtualedit' set](https://github.com/vim/vim/commit/db0eedec16621854c772760d02427804bc0a298d)
- [8.0.1766 : expanding abbreviation doesn't work](https://github.com/vim/vim/commit/c3c3e698966fac86dee94799b70947defb85440d)
- [8.0.1767 : with 'incsearch' text may jump up and down](https://github.com/vim/vim/commit/9d34d90210ba52ebaf45973282e5921f5af364c7)
- [8.0.1768 : SET&#x5f;NO&#x5f;HLSEARCH() used in a wrong way](https://github.com/vim/vim/commit/451fc7b954906069f1830a8092ad85616049a828)
- [8.0.1769 : repeated saving and restoring viewstate for 'incsearch'](https://github.com/vim/vim/commit/9b25af36204c0511eab08d621688f0f2008fc68e)
- [8.0.1770 : assert functions don't return anything](https://github.com/vim/vim/commit/65a5464985f980d2bbbf4e14d39d416dce065ec7)
- [8.0.1771 : in tests, when WaitFor() fails it doesn't say why](https://github.com/vim/vim/commit/50182fa84e20a0547f3e2bd6683ef799fcd27855)
- [8.0.1772 : quickfix: mixup of FALSE and FAIL, returning -1](https://github.com/vim/vim/commit/29ce409bfca52bb8a07e2975d06fd788458e9861)
- [8.0.1773 : dialog messages are not translated](https://github.com/vim/vim/commit/c36651b4b946333dce0a916326d821d2562cf39d)
- [8.0.1774 : reading very long lines can be slow](https://github.com/vim/vim/commit/13d3b05ed2cf9a54b18b4e8236f0af2c5386200c)
- [8.0.1775 : MS-Windows: warning for unused variable](https://github.com/vim/vim/commit/850d427f070a61267ef384a449992e72d2bea12c)
- [8.0.1776 : in tests, when WaitFor() fails it doesn't say why](https://github.com/vim/vim/commit/0e9d1ae3216a5940b36bb56d155fb300b2e55b00)
- [8.0.1777 : cannot cleanup before loading another colorscheme](https://github.com/vim/vim/commit/60a68362aa73f4a6cb534688978f9dc2b16e60fe)
- [8.0.1778 : script to check translations does not always work](https://github.com/vim/vim/commit/b07bbb0d29493fcf4ed080fe018535e64441d663)
- [8.0.1779 : deleting in a block selection causes problems](https://github.com/vim/vim/commit/35e802e713382d7e76232ad344af7dcd577e43de)
- [8.0.1780 : test fails because Vim in a terminal uses wrong 'encoding'](https://github.com/vim/vim/commit/77bfd756a02c3b3150d6c08ddbf89b66362dad8d)

## 新着スクリプト

- [vim-hardaway : Vim colorscheme with Airline integration. ](https://www.vim.org/scripts/script.php?script_id=5673)
- [move-less : plugin for faster navigation without changing the cursor position](https://www.vim.org/scripts/script.php?script_id=5674)
- [git-remote-open : Copy/Open url of files directly from vim on github/bitbucket](https://www.vim.org/scripts/script.php?script_id=5675)
- [Flexagon : Fold switcher with variety of prepackaged folds](https://www.vim.org/scripts/script.php?script_id=5676)
- [traces.vim : Range, pattern and substitute preview for Vim ](https://www.vim.org/scripts/script.php?script_id=5677)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](https://www.vim.org/scripts/script.php?script_id=273) (973)
2. [wombat256.vim : Wombat for 256 color xterms](https://www.vim.org/scripts/script.php?script_id=2465) (724)
3. [The NERD tree : A tree explorer plugin for navigating the filesystem](https://www.vim.org/scripts/script.php?script_id=1658) (707)
4. [OmniCppComplete : C/C++ omni-completion with ctags database](https://www.vim.org/scripts/script.php?script_id=1520) (412)
5. [AutoComplPop : Automatically opens popup menu for completions](https://www.vim.org/scripts/script.php?script_id=1879) (294)
6. [minibufexpl.vim : Elegant buffer explorer - takes very little screen space](https://www.vim.org/scripts/script.php?script_id=159) (287)
7. [python.vim : Enhanced version of the python syntax highlighting script](https://www.vim.org/scripts/script.php?script_id=790) (284)
8. [nginx.vim : initial version](https://www.vim.org/scripts/script.php?script_id=1886) (277)
9. [c.vim : C/C++ IDE -- Write and run programs. Insert statements, idioms, comments etc.](https://www.vim.org/scripts/script.php?script_id=213) (262)
10. [jellybeans.vim : Colorful, dark color scheme](https://www.vim.org/scripts/script.php?script_id=2555) (245)

## vim-jp issues

Open : 261 (-1) | Closed : 904 (+5)

- [Issue #1162 : Vim script に複合代入演算子を追加したい](https://github.com/vim-jp/issues/issues/1162)
- [Issue #1163 : float 型の等値比較について](https://github.com/vim-jp/issues/issues/1163)
- [Issue #1164 : timer callback 内で起動した job の exit-cb 実行が遅れることがある](https://github.com/vim-jp/issues/issues/1164)
- [Issue #1165 : 'runtimepath' に登録済みのパスに前方一致するパスが :packadd されない](https://github.com/vim-jp/issues/issues/1165)
