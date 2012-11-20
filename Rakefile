require "rubygems"
require "bundler"
Bundler.require

StandaloneMigrations::Tasks.load_tasks

namespace :data do
  desc "get deputados"
  task :deputados do
    require './data/camara_parser'
    require './data/pesquisa_deputados_parser'
    require './data/deputados_xml_parser'
    require './lib/activerecord'

    Deputado.save_from_pesquisa_parser
  end
end
