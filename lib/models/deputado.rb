class Deputado < ActiveRecord::Base
  validates_presence_of :nome_parlamentar, :cadastro_id

  def self.save_from_pesquisa_parser
    parser = PesquisaDeputadosParser.new
    parser.deputados.each do |deputado|
      create! nome_parlamentar: deputado[:nome_parlamentar], cadastro_id: deputado[:id]
    end
  end
end
