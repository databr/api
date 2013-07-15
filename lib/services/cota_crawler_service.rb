class CotaCrawlerService
  def self.save_from_cota_xml_parser
    unzip('db/data/cota_2013.xml.zip') do |file|
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

  def self.unzip(file, &block)
    require 'zip/zipfilesystem'
    xml_file = file.gsub('.zip', '')
    entry = xml_file.split("/").last
    dest = "tmp/#{entry}"

    Zip::ZipFile.open(file) do |zipfile|
      zipfile.extract(entry, dest)
    end

    yield(dest)
  end
end
