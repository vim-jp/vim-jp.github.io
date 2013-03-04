// vim:set ts=8 sts=2 sw=2 tw=0 et:

google.load("visualization", "1");
$(function() {

  function escapeHTML(str) {
    return str.replace('&', '&amp;')
      .replace('<', '&lt;')
      .replace('>', '&gt;')
      .replace('"', '&quot;')
      .replace("'", '&apos;');
  }

  function to_jsonml(s) {
    var ml = [];
    var re = /(?:http|https):\/\/[\x21-\x7e]+/i;
    while (s.length > 0) {
      if (s.match(re)) {
        ml.push(
          RegExp.leftContext,
          [ 'a', { 'href': RegExp.lastMatch }, RegExp.lastMatch ]
        );
        s = RegExp.rightContext;
      } else {
        ml.push(s);
        s = '';
      }
    }
    return ml;
  }

  var query = new google.visualization.Query('https://spreadsheets.google.com/tq?key=0AvP2GcrHKTfudEdwWm85ajZtd0ltZjRUcVFPWHpzWHc&gid=0&pub=1');
  query.send(function(response) {
    if(response.isError()) {
      return;
    }
    var data = response.getDataTable();
    var colNames = (function() {
      var colNames = {};
      for (var i = 0, max = data.getNumberOfColumns(); i < max; i++) {
        colNames[data.getColumnLabel(i)] = i;
      }
      return colNames;
    })();

    $('#vimmers-container').empty();
    for (var row = 0; row < data.getNumberOfRows(); row++) {
      var name        = data.getFormattedValue(row, colNames['Name']);
      var shortname   = data.getFormattedValue(row, colNames['Short Name']);
      var image       = data.getFormattedValue(row, colNames['Image']);
      var twitter     = data.getFormattedValue(row, colNames['Twitter Account']);
      var github      = data.getFormattedValue(row, colNames['Github Account']);
      var facebook    = data.getFormattedValue(row, colNames['Facebook']);
      var googleplus  = data.getFormattedValue(row, colNames['Google+']);
      var vimorg      = data.getFormattedValue(row, colNames['www.vim.org']);
      var website     = data.getFormattedValue(row, colNames['Website URL']);
      var description = data.getFormattedValue(row, colNames['Description']);

      var items = [];

      items.push([ 'li', { 'class': 'name' }, name + (shortname ? ' (' + shortname + ')' : '') ]);

      if (twitter) {
        var twitter_url = 'https://twitter.com/' + twitter;
        if (!image) {
          image = 'http://api.twitter.com/1/users/profile_image/'
            + twitter + '.png&size=mini';
        }
        items.push([
          'li', { 'class': 'link' },
          [ 'a', { 'href': twitter_url }, [ 'img', { 'src': '/assets/images/icon-twitter.png', 'title': 'Twitter' } ] ]
        ]);
      }

      if (github) {
        var github_url = 'https://github.com/' + github;
        if (!image) {
          items.unshift([
            'li', { 'class': 'faceicon ' },
            [ 'img', { 'src': '/assets/images/icon-loading.gif', 'class': github } ]
          ]);
          (function(github) {
            $.getJSON('https://api.github.com/users/' + github, function(data) {
              $('img.' + github).attr('src', data.avatar_url);
            });
          })(github);
        }
        items.push([
          'li', { 'class': 'link' },
          [ 'a', { 'href': github_url }, [ 'img', { 'src': '/assets/images/icon-github.png', 'title': 'GitHub' } ] ]
        ]);
      }

      if (facebook) {
        var facebook_url = 'https://www.facebook.com/' + facebook;
        items.push([
          'li', { 'class': 'link' },
          [ 'a', { 'href': facebook_url }, [ 'img', { 'src': '/assets/images/icon-facebook.png', 'title': 'Facebook' } ] ]
        ]);
      }

      if (googleplus) {
        var googleplus_url = 'https://plus.google.com/' + googleplus
        items.push([
          'li', { 'class': 'link' },
          [ 'a', { 'href': googleplus_url }, [ 'img', { 'src': '/assets/images/icon-googleplus.png', 'title': 'Google+' } ] ]
        ]);
      }

      if (website) {
        items.push([
          'li', { 'class': 'link' },
          [ 'a', { 'href': website }, [ 'img', { 'src': '/assets/images/icon-website.png', 'title': 'WebSite' } ] ]
        ]);
      }

      if (vimorg) {
        var vimorg_url = 'http://www.vim.org/account/profile.php?user_id=' + vimorg;
        items.push([
          'li', { 'class': 'link' },
          [ 'a', { 'href': vimorg_url }, [ 'img', { 'src': '/assets/images/icon-vim.png', 'title': 'Vim' } ] ]
        ]);
      }

      items.push([ 'li', { 'class': 'clear' } ]);

      if (description) {
        items.push([ 'li', { 'class': 'desc' } ]
            .concat(to_jsonml(escapeHTML(description))));
      }

      if (image) {
        items.unshift([
          'li', { 'class': 'faceicon' },
          [ 'img', { 'src': image } ]
        ]);
      }

      $.jqml([
        'div', { 'class': 'person', },
        [ 'ul' ].concat(items)
      ]).appendTo('#vimmers-container');
    }
    $('#vimmers-container').masonry({
      isAnimated: true,
      isResizable: true
    });
  });
});
