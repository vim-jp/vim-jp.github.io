#!/bin/sh

set -eu

OLDDIR=$PWD
BUILDDIR=`mktemp -d`
cd "$BUILDDIR"
git clone git@github.com:vim-jp/vimdoc-ja.git
cd vimdoc-ja
MASTER_TIME=`git log -1 --pretty=format:%ct origin/master`
GHPAGES_TIME=`git log -1 --pretty=format:%ct origin/gh-pages`
echo "$MASTER_TIME $GHPAGES_TIME"
if [ $MASTER_TIME -gt $GHPAGES_TIME ] ; then
  git checkout -b devel origin/devel
  make htmlbatch
  git push origin gh-pages
fi
cd $OLDDIR
rm -rf "$BUILDDIR"
