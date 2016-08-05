serve 'jekyll serve'
port 4000

before_build {
  run 'bundle install'
}
