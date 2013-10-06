class DeputadoCrawlerService
  def self.save_from_pesquisa_parser
    parser = PesquisaDeputadosParser.new
    parser.deputados.each do |deputado|
      Deputado.first_or_create! nome_parlamentar: deputado[:nome_parlamentar], cadastro_id: deputado[:id]
    end
  end

  def self.save_from_deputado_xml_parser
    parser = DeputadoXMLParser.new
    parser.deputados.each do |deputado_result|
      deputado_result.delete(:partido) # without partido for now
      deputado_result.delete(:comissoes)
      deputado = Deputado.find_or_create_by_cadastro_id deputado_result.delete(:cadastro_id)
      deputado.update_attributes deputado_result
    end
  end
end
