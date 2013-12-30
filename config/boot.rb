# encoding: utf-8
require 'rubygems'
env = ENV['RACK_ENV'] ? ENV['RACK_ENV'] : 'development'
require 'bundler'
Bundler.require(:default, env)

require 'active_record'
require './lib/api_logger.rb'
LOGGER = ApiLogger.new

if ENV['USE_PARSER'] == 'true'
  require './data/camara_parser'
  Dir[File.join('data/*.rb')].each { |f| require File.absolute_path(f) }
  Dir[File.join('lib/services/*.rb')].each { |f| require File.absolute_path(f) }
end

require './lib/entity/base_on_feed_entity.rb'
Dir[File.join('lib/models/*.rb')].each { |f| require File.absolute_path(f) }
Dir[File.join('lib/entity/*.rb')].each { |f| require File.absolute_path(f) }

require 'aggregator'

database_config_file = File.read('config/database.yml')

if env == 'production'
  require 'erubis'
  e = Erubis::Eruby.new(database_config_file)
  database_config_file = e.result(binding)
end

database_config = YAML.load(database_config_file)[env]
ActiveRecord::Base.include_root_in_json = false
ActiveRecord::Base.logger = LOGGER
ActiveRecord::Base.establish_connection(database_config)

CACHE = Dalli::Client.new(
                    (ENV['MEMCACHED_SERVERS'] || 'localhost:11211').split(',') ,
                    {
                      username: ENV['MEMCACHED_USERNAME'],
                      password: ENV['MEMCACHED_PASSWORD'],
                      failover: true,
                      socket_timeout: 1.5,
                      socket_failure_delay: 0.2
                    })

uri = URI.parse(ENV['REDIS_URL'] || 'redis://127.0.0.1:6379')
REDIS = Redis.new(host: uri.host, port: uri.port, password: uri.password)
