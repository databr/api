class Deputado < ActiveRecord::Base
  validates_presence_of :name
  validates_presence_of :deputado_id

  def self.save_from_pesquisa_parser
    parser = PesquisaDeputadosParser.new
    parser.deputados.each do |deputado|
      create! name: deputado[:name], deputado_id: deputado[:id]
    end
  end
end
