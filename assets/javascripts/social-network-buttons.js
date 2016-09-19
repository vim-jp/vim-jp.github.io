'use strict';

window.addEventListener('load', function() {
  (function(doc, script) {
    window.___gcfg = {lang: 'ja'};
    //Async Social Buttons
    var fjs = doc.getElementsByTagName(script)[0];
    var add = function(url, id) {
      if (doc.getElementById(id)) {return;}
      var js = doc.createElement(script);
      js.src = url;
      id && (js.id = id);
      fjs.parentNode.insertBefore(js, fjs);
    };

    // Twitter SDK
    add('//platform.twitter.com/widgets.js', 'twitter-wjs');
    // Google+ button
    add('//apis.google.com/js/plusone.js');
    // Facebook SDK
    add('//connect.facebook.net/ja_JP/all.js#xfbml=1', 'facebook-jssdk');
    // Hatena
    add('//b.st-hatena.com/js/bookmark_button_wo_al.js');
  })(document, 'script');
});
