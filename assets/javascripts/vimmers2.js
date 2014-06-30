function loadVimmers(data) {
  function shuffleArray(array) {
    var i = array.length;

    while (--i) {
      var j = Math.floor(Math.random() * (i + 1));
      if (i == j) continue;
      var k = array[i];
      array[i] = array[j];
      array[j] = k;
    }

    return array;
  }

  var bram = data.shift();
  data = shuffleArray(data);
  data.unshift(bram);

  var vimmers = document.createElement('div');

  data.forEach(function (d) {
    var icon = d.logo;
    var twitter = d.twitter;
    var iconTwitter = d.twitter_icon;
    var github = d.github;
    var facebook = d.facebook;
    var googleplus = d.googleplus;
    var vimorg = d.vimorg;
    var website = d.website;
    var description = d.description;
    var links = [];

    var grid = document.createElement('div');
    grid.className = 'grid';
    vimmers.appendChild(grid);

    var vimmer = document.createElement('section');
    vimmer.className = 'vimmer';
    grid.appendChild(vimmer);

    var name = document.createElement('h1');
    name.textContent = d.name;
    vimmer.appendChild(name);

    var avatar = document.createElement('div');
    avatar.className = 'avatar';
    vimmer.appendChild(avatar);

    if (twitter) {
      if (!icon && iconTwitter) {
        icon = iconTwitter;
      }

      links.push({
        'url': 'https://twitter.com/' + twitter,
        'label': 'Twitter',
        'class': 'twitter'
      });
    }

    if (github) {
      // TODO: Load GitHub icon here

      links.push({
        'url': 'https://github.com/' + github,
        'label': 'GitHub',
        'class': 'github'
      });
    }

    if (facebook) {
      links.push({
        'url': 'https://www.facebook.com/' + facebook,
        'label': 'Facebook',
        'class': 'facebook'
      });
    }

    if (googleplus) {
      links.push({
        'url': 'https://plus.google.com/' + googleplus,
        'label': 'Google+',
        'class': 'googleplus'
      });
    }

    if (vimorg) {
      links.push({
        'url': 'http://www.vim.org/account/profile.php?user_id=' + vimorg,
        'label': 'Vim',
        'class': 'vimorg'
      });
    }

    if (website) {
      links.push({
        'url': website,
        'label': 'Website',
        'class': 'website'
      });
    }

    if (!icon) {
      icon = '/assets/images/icon2-default';

      if (Modernizr && Modernizr.svg) {
        icon += '.svg';
      } else {
        icon += '.png';
      }
    }

    var imgAvatar = document.createElement('img');
    imgAvatar.src = icon;
    avatar.appendChild(imgAvatar);

    var elsewhere = document.createElement('ul');
    elsewhere.className = 'elsewhere';
    vimmer.appendChild(elsewhere);

    links.forEach(function (link) {
      var li = document.createElement('li');
      elsewhere.appendChild(li);

      var a = document.createElement('a');
      a.className = link.class;
      a.href = link.url;
      a.textContent = link.label;
      li.appendChild(a);
    });
  });

  document.getElementById('vimmers').innerHTML = vimmers.innerHTML;
}

(function () {
  var js = document.createElement('script');
  js.src = 'http://vim-jp.herokuapp.com/vimmers?callback=loadVimmers';
  js.async = true;
  var s = document.getElementsByTagName('script')[0];
  s.parentNode.insertBefore(js, s);
})();
