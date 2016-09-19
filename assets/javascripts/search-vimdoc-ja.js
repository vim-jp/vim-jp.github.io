(function(google) {
  'use strict';

  google.load("search", "1", {language : "ja"});
  google.setOnLoadCallback(function() {
    // search
    var customSearchControl = new google.search.CustomSearchControl("001325159752250591701:65aunpq8rlg");
    customSearchControl.setResultSetSize(google.search.Search.FILTERED_CSE_RESULTSET);
    var options = new google.search.DrawOptions();
    options.enableSearchboxOnly("//www.google.com/cse?cx=001325159752250591701:65aunpq8rlg");
    customSearchControl.draw('VimdocJaSearch', options);
  }, true);
})(google);
