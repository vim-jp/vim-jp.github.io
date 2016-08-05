serve 'bundle exec jekyll serve'
port 4000

before_build {
  run 'bundle install --path vendor --binstubs'
}
