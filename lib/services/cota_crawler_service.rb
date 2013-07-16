# encoding: utf-8
#
class CotaCrawlerService
  def self.save_from_cota_xml_parser
    unzip('db/data/cota_2013.xml.zip') do |file|
      parser = CotaXMLParser.new(file)
      parser.cotas.each do |cota|
        deputado = Deputado.where(nome_parlamentar: cota[:nome_parlamentar]).first
        print "Creating cota..."
        cota.delete(:nome_parlamentar)
        if deputado.nil?
          Cota.create! cota
        else
          deputado.cotas.create! cota
        end
        puts " Done!"
      end
    end
  end

  def self.unzip(file, &block)
    require 'zip/zipfilesystem'
    xml_file = file.gsub('.zip', '')
    entry = xml_file.split("/").last
    dest = "tmp/#{entry}"

    system "rm #{dest}"
    Zip::ZipFile.open(file) do |zipfile|
      zipfile.extract(entry, dest)
    end

    yield(dest)
  end
end
