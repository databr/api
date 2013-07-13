class CotaCrawlerService
  def self.save_from_cota_xml_parser
    file = 'db/data/cota_2013.xml'
    parser = CotaXMLParser.new(file)
    parser.cotas.each do |cota|
      deputado = Deputado.find_or_create_by_nome_parlamentar(cota.delete(:nome_parlamentar))
      deputado.save!
      deputado.cotas.create! cota
    end
  end
end
