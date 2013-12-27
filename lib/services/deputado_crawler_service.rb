# encoding: UTF-8
#
class DeputadoCrawlerService
  def self.save_from_pesquisa_parser
    parser = PesquisaDeputadosParser.new
    puts "\e[32m  * Starting save_from_pesquisa_parser\e[0m"
    parser.deputados.each do |deputado|
      puts "\e[32m    * Saving #{deputado[:nome_parlamentar]}(#{deputado[:id]}) from pesquisa\e[0m"
      d = Deputado.where(nome_parlamentar: deputado[:nome_parlamentar], cadastro_id: deputado[:id])
      d.create! unless d.first
    end
  end

  def self.save_from_deputado_xml_parser
    parser = DeputadoXMLParser.new
    puts "\e[32m  * Starting save_from_deputado_xml_parser\e[0m"
    parser.deputados.each do |deputado_result|
      deputado_result.delete(:partido) # without partido for now
      deputado_result.delete(:comissoes)
      deputado = Deputado.find_or_create_by_cadastro_id deputado_result.delete(:cadastro_id)
      puts "\e[32m    * Saving #{deputado.nome_parlamentar}(#{deputado.id}) from xml parser\e[0m"
      deputado.update_attributes deputado_result
    end
  end

  def self.save_from_deputado_about_parser
    puts "\e[32m  * Starting save_from_deputado_about_parser\e[0m"
    parser = PesquisaDeputadosParser.new
    parser.deputados.each do |deputado|
      DeputadoAboutParser.new(deputado[:id]).sections.each do |section|
        title = section.title # .force_encoding('iso8859-1').encode('utf-8')
        if section.text.strip.empty?
          puts "\e[31m    - Section #{title.gsub(/\s+/, " ").strip}(#{deputado[:id]}) empty :( \e[0m"
          next
        end
        body = section.text   # .force_encoding('iso8859-1').encode('utf-8')
        section_key = Digest::MD5.hexdigest(title)
        about = About.where(cadastro_id: deputado[:id], section_key: section_key).first
        if about
          puts "\e[32m    * Updating #{title}(#{deputado[:id]}) from deputado about parser\e[0m"
          about.update_attributes body: body
        else
          puts "\e[32m    * Creating #{title}(#{deputado[:id]}) from deputado about parser\e[0m"
          About.create!(cadastro_id: deputado[:id],
                        title: title,
                        body: body,
                        section_key: section_key)
        end
      end
    end
  end
end
