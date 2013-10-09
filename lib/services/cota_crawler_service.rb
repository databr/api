# encoding: utf-8
#
class CotaCrawlerService
  def self.save_from_cota_xml_parser
    #unzip('db/data/cota_2013.xml.zip') do |file|
    start = Time.now.to_i
    file = 'db/data/cota_2013.xml'
    parser = CotaXMLParser.new(file)
    parser.cotas.each_slice(50) do |slice_cotas|
      slice_cotas.each do |cota|
        next if Cota.where(numero: cota[:numero]).first
        deputado = Deputado.where(nome_parlamentar: cota[:nome_parlamentar]).first
        cota[:data_emissao] = cota[:data_emissao].nil? ? Time.new(cota[:ano], cota[:mes], 1, 0, 0, 0, '-03:00') : Time.iso8601("#{cota[:data_emissao]}-03:00")
        if deputado.nil?
          unless deputado = Deputado.where(other_name: cota[:nome_parlamentar]).first
            begin
              deputado = search(cota[:nome_parlamentar], 6);
              deputado.update_attribute :other_name, cota[:nome_parlamentar]
            rescue
              open('not_found', 'a') do |f|
                f << cota
              end
              next
            end
          end
        end
        cota.delete(:nome_parlamentar)
        deputado.cotas.create cota
        puts "."
      end
    end
    #end
    puts "Finished!: #{Time.now.to_i - start}"
  end

  def self.unzip(file, &block)
    require 'zip/zipfilesystem'
    xml_file = file.gsub('.zip', '')
    entry = xml_file.split("/").last
    dest = "tmp/#{entry}"

    system "rm #{dest}"
    Zip::ZipFile.open(file) do |zipfile|
      zipfile.extract('AnoAtual.xml', dest)
    end

    yield(dest)
  end

  def self.search(name, words, where_old = nil, try_matchs = false)
    return matchs(name, where_old) if words == 0
    regexp = Regexp.new("\\w{#{words}}")
    match = name.match(regexp)
    return matchs(name, Deputado.all) if match.nil? && match.to_s.length == 0
    deputados = (where_old || Deputado).where('nome_parlamentar LIKE ?', "%#{match}%")
    return matchs(name, deputados) if try_matchs
    return search(name, words+1, nil, true) if deputados.count == 0
    if deputados.count > 1
      search(name, words-1, deputados)
    else
      deputados.first
    end
  end

  def self.matchs(search, possibles)
    include Amatch
    m = Sellers.new(search)
    result = m.match(possibles.map(&:nome_parlamentar))
    match_index = result.index(result.min)
    possibles[match_index]
  end
end
