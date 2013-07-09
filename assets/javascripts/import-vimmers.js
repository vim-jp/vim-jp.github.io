// vim:set ts=8 sts=2 sw=2 tw=0 et:

$(function() {

  var elementCache = [];
  var limits = 50;

  function escapeHTML(str) {
    return str.replace('&', '&amp;')
      .replace('<', '&lt;')
      .replace('>', '&gt;')
      .replace('"', '&quot;')
      .replace("'", '&apos;');
  }

  function to_jsonml(s) {
    var ml = [];
    var re = /\n/;
    while (s.length > 0) {
      if (s.match(re)) {
        ml.push(RegExp.leftContext, [ 'br' ]);
        s = RegExp.rightContext;
      } else {
        ml.push(s);
        s = '';
      }
    }
    re = /(?:http|https):\/\/[\x21-\x7e]+/i;
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

  function shuffleNumbers(max) {
    var rows = [];
    for (var i = 0; i < max; ++i) {
      rows.push(i);
    }
    for (var i = 1; i < max - 1; ++i) {
      var n = Math.floor((max - i) * Math.random()) + i;
      if (n != i) {
        var tmp = rows[i];
        rows[i] = rows[n];
        rows[n] = tmp;
      }
    }
    return rows;
  }

  function toElement(getter) {
    var name         = getter('name');
    var shortname    = getter('short_name');
    var image        = getter('logo');
    var twitter      = getter('twitter');
    var twitter_icon = getter('twitter_icon');
    var github       = getter('github');
    var facebook     = getter('facebook');
    var googleplus   = getter('googleplus');
    var vimorg       = getter('vimorg');
    var website      = getter('website');
    var description  = getter('description');

    var items = [];
    var ready_hooks = [];

    items.push([ 'li', { 'class': 'name' }, name + (shortname ? ' \u0028' + shortname + '\u0029' : '') ]);

    if (twitter) {
      var twitter_url = 'https://twitter.com/' + twitter;
      if (!image && twitter_icon) {
        image = twitter_icon
      }
      items.push([
        'li', { 'class': 'link' },
        [ 'a', { 'href': twitter_url }, [ 'img', { 'src': '/assets/images/icon-twitter.png', 'title': 'Twitter' } ] ]
      ]);
    }

    if (github) {
      var github_url = 'https://github.com/' + github;
      var gh_id = 'github-id-' + github;
      if (!image) {
        items.unshift([
          'li', { 'class': 'faceicon ' },
          [ 'img', { 'src': '/assets/images/icon-loading.gif', 'id': gh_id } ]
        ]);
        // Register gh-icon loader to "ready_hooks".
        ready_hooks.push(function(element) {
          $.getJSON('https://api.github.com/users/' + github + '?callback=?', function(res) {
            element.find('img#' + gh_id).attr('src', res.data.avatar_url);
          });
        });
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

    var element = $.jqml([
      'div', { 'class': 'person', },
      [ 'ul' ].concat(items)
    ]);
    // Fire all "ready_hooks" events.
    element.ready(function() {
      for (var i = 0, max = ready_hooks.length; i < max; ++i) {
        ready_hooks[i](element);
      }
    })

    return element;
  }

  function refreshVimmers(first) {
    var container = $('#vimmers-container');
    container.empty();
    var order = shuffleNumbers(elementCache.length);
    var max = Math.min(order.length, limits);
    for (var i = 0; i < max; ++i) {
      container.append(elementCache[order[i]]);
    }
    if (first) {
      container.masonry({
        isAnimated: true,
        isResizable: true
      });
    } else {
      console.log('reload');
      container.masonry('reload');
    }
  }

  function loadVimmers(end_cb) {
    elementCache = [];
    $('#vimmers-container').empty();
    $.getJSON("http://jsonp.ap01.aws.af.cm/?url=https://raw.github.com/vim-jp/vim-jp.github.com/master/vimmers/vimmers.json&callback=?", function(data) {
      $.each(data, function(n, e) {
        var element = toElement(function(name) {
          return e[name];
        });
        elementCache.push(element);
      });
      refreshVimmers(true);
      if (end_cb) {
        end_cb();
      }
    });
  }

  loadVimmers(function() {
    $('#vimmers-shuffle').click(function() {
      refreshVimmers(false);
    });
    $('#vimmers-showall').click(function() {
      limits = elementCache.length;
      refreshVimmers(false);
      $('#vimmers-showall').hide();
    });
  });
});
