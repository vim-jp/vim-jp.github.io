serve 'bundle exec jekyll serve'
port 4000

before_build {
  run 'rbenv versions'
  run 'rbenv install --list'
  run 'bundle install --path vendor --binstubs'
}
