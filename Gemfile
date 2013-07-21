source 'http://rubygems.org'

gem 'mechanize'
gem 'activerecord'
gem 'rake'
gem 'standalone_migrations'
gem 'grape'
gem 'rack-cors', :require => 'rack/cors'
gem 'rubyzip'
gem 'thin'

group :production do
  gem 'pg'
end

group :development, :test do
  gem 'sqlite3'
  gem 'shotgun'
  gem 'autotest'
  gem 'ZenTest', '4.9.2'
  gem 'pry'
end

group :test do
  gem 'shoulda-matchers'
  gem 'rspec'
  gem 'autotest'
  gem 'webmock'
end
