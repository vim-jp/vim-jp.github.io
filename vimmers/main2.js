// vim:set ts=8 sts=2 sw=2 tw=0 et:
(function(global){

var VIMMERS_DATA_URL = "http://vim-jp.herokuapp.com/vimmers?callback=?";
var DEFAULT_LIMITS = 50;

var allVimmers = [];

function main() {
  var ractive;

  ractive = new Ractive({
    debug: true,
    el: '#vimmers-body',
    template: '#vimmers-template',
    data: {
      persons: [],
      all: false
    }
  });
  ractive.on({
    shuffle: onShuffle,
    showAll: onShowAll
  });
  loadJSON(VIMMERS_DATA_URL).then(onFirstLoad.bind(ractive));
}

function loadJSON(url) {
  return new Promise(function(resolve, reject) {
    $.getJSON(url, function(data) {
      resolve(data);
    });
  });
}

function loadGithubIcon(id) {
  return loadJSON('https://api.github.com/users/' + id + '?callback=?');
}

function toVimmerData(raw, delay) {
  var v;

  v = {
    name: raw.name,
    shortName: raw.short_name,
    desc: raw.description,
    face_icon_url: raw.logo,
    links: []
  };

  if (raw.twitter) {
    if (!v.face_icon_url && raw.twitter_icon) {
      v.face_icon_url = raw.twitter_icon;
    }
    v.links.push({
      url: 'https://twitter.com/' + raw.twitter,
      type: 'Twitter',
      icon_url: '/assets/images/icon-twitter.png'
    });
  }

  if (raw.github) {
    if (!v.face_icon_url) {
      v.face_icon_url = '/assets/images/icon-loading.gif';
      delay(loadGithubIcon(raw.github).then(function(res) {
        try {
          if (res.data.avatar_url) {
            v.face_icon_url = res.data.avatar_url;
          } else {
            // FIXME: set rate limited error image.
            v.face_icon_url = null;
            console.warn('may reached rate limit of Github API for '
                + raw.name);
          }
        } catch (e) {
          console.error(e.stack);
        }
      }));
    }
    v.links.push({
      url: 'https://github.com/' + raw.github,
      type: 'Github',
      icon_url: '/assets/images/icon-github.png'
    });
  }

  if (raw.facebook) {
    v.links.push({
      url: 'https://www.facebook.com/' + raw.facebook,
      type: 'Facebook',
      icon_url: '/assets/images/icon-facebook.png'
    });
  }

  if (raw.googleplus) {
    v.links.push({
      url: 'https://plus.google.com/' + raw.googleplus,
      type: 'Google+',
      icon_url: '/assets/images/icon-googleplus.png'
    });
  }

  if (raw.website) {
    v.links.push({
      url: raw.website,
      type: 'WebSite',
      icon_url: '/assets/images/icon-website.png'
    });
  }

  if (raw.vimorg) {
    v.links.push({
      url: 'http://www.vim.org/account/profile.php?user_id=' + raw.vimorg,
      type: 'Vim',
      icon_url: '/assets/images/icon-vim.png'
    });
  }

  return v;
}

function showVimmers(ractive, all) {
  var order, persons;

  order = getShuffledNumbers(allVimmers.length);
  if (!all) {
    order = order.splice(0, DEFAULT_LIMITS);
  }
  persons = _.map(order, function(index) {
    return allVimmers[index];
  });

  ractive.set({
    persons: persons,
    all: all
  });
}

function getShuffledNumbers(max) {
  var rows = _.range(max);
  var i, n, tmp;

  // Start from 1, to keep Mr.Bram at first.
  for (i = 1; i < max - 1; ++i) {
    n = Math.floor((max - i) * Math.random()) + i;
    if (n != i) {
      tmp = rows[i];
      rows[i] = rows[n];
      rows[n] = tmp;
    }
  }
  return rows;
}

function onFirstLoad(data) {
  var ractive, delays;

  ractive = this;
  try {
    delays = [];
    allVimmers = _.map(data, function(raw) {
      return toVimmerData(raw, function(promise) {
        delays.push(promise);
      });
    });
    showVimmers(ractive, false);
    Promise.all(delays).then(function() {
      ractive.update();
    });

    $('#vimmers-container').masonry({
      isAnimated: true,
      isResizable: true
    });
  } catch (e) {
    console.error(e.stack);
  }
}

function onShuffle(e) {
  showVimmers(this, this.get('all'));
  $('#vimmers-container').masonry('reload');
}

function onShowAll(e) {
  showVimmers(this, true);
  $('#vimmers-container').masonry('reload');
}

global.addEventListener('load', main);

})(this);
