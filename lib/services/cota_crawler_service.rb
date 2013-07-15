class CotaCrawlerService
  def self.save_from_cota_xml_parser
    file = 'db/data/cota_2013.xml'
    parser = CotaXMLParser.new(file)
    parser.cotas.each do |cota|
      deputado = Deputado.where(nome_parlamentar: cota[:nome_parlamentar]).first
      if deputado.nil?
        File.open('deputado_not_found', 'a') {|f| f.puts(cota) }
        next
      end
      cota.delete(:nome_parlamentar)
      deputado.cotas.create! cota
    end
  end
end
