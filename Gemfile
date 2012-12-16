source 'http://rubygems.org'

gem 'mechanize'
gem 'activerecord'
gem 'rake'
gem 'standalone_migrations'
gem 'grape'
gem 'rack-cors', :require => 'rack/cors'

group :production do
  gem 'pg'
end

group :development do
  gem 'sqlite3'
  gem 'shotgun'
end

group :test do
  gem 'shoulda-matchers'
  gem 'rspec'
  gem 'autotest'
  gem 'pry'
  gem 'webmock'
end
