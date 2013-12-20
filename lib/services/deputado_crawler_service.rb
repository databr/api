# encoding: UTF-8
#
class DeputadoCrawlerService
  def self.save_from_pesquisa_parser
    parser = PesquisaDeputadosParser.new
    parser.deputados.each do |deputado|
      d = Deputado.where(nome_parlamentar: deputado[:nome_parlamentar], cadastro_id: deputado[:id])
      d.create! unless d
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

  def self.save_from_deputado_about_parser
    parser = PesquisaDeputadosParser.new
    parser.deputados.each do |deputado|
      DeputadoAboutParser.new(deputado[:id]).sections.each do |section|
        title = section.title #.force_encoding('iso8859-1').encode('utf-8')
        body = section.text   #.force_encoding('iso8859-1').encode('utf-8')
        About.create!(cadastro_id: deputado[:id], title: title, body: body, section_key: Digest::MD5.hexdigest(title)) unless section.text.strip.empty?
      end
    end
  end
end
