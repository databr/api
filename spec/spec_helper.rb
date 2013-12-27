# encoding: utf-8
ENV['RACK_ENV'] = 'test'
ENV['USE_PARSER'] = 'true'
ENV['YEAR'] = '2013'

require './config/boot'

require 'webmock/rspec'

Dir[File.join('spec/support/*.rb')].each { |f| require File.absolute_path(f) }

RSpec.configure do |config|
  config.treat_symbols_as_metadata_keys_with_true_values = true
  config.run_all_when_everything_filtered = true
  config.filter_run :focus

  config.order = 'random'

  config.expect_with :rspec do |c|
    c.syntax = :expect
  end

  config.around do |example|
    ActiveRecord::Base.transaction do
      example.run
      fail ActiveRecord::Rollback
    end
  end
end
