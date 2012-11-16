require 'bundler'
Bundler.require

require 'pry'
require 'shoulda-matchers'
require 'webmock/rspec'


Dir[File.join("data/*.rb")].each {|f| require File.absolute_path(f) }
Dir[File.join("spec/support/*.rb")].each {|f| require File.absolute_path(f) }


RSpec.configure do |config|
  config.treat_symbols_as_metadata_keys_with_true_values = true
  config.run_all_when_everything_filtered = true
  config.filter_run :focus

  config.order = 'random'

  config.expect_with :rspec do |c|
    c.syntax = :expect
  end

  config.around do |example|
    stub_request(:get, 'http://www2.camara.leg.br/deputados/pesquisa').
      to_return(:status => 200, :body => File.read('spec/fixtures/pesquisa.html'), :headers => {:'Content-Type' => 'text/html'})


    ActiveRecord::Base.transaction do
      example.run
      raise ActiveRecord::Rollback
    end
  end
end
