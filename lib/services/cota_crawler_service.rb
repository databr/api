# encoding: utf-8
#
class CotaCrawlerService
  def self.save_from_cota_xml_parser
    puts "\e[32m* Starting cota parser \e[0m"
    start = Time.now
    unzip('db/data/cota_2013.xml.zip') do |file|
      # file = 'db/data/cota_2013.xml'
      puts "\e[32m  * Unziped and get #{file} \e[0m"
      parser = CotaXMLParser.new(file)
      parser.cotas.each_slice(50) do |slice_cotas|
        slice_cotas.each do |cota|
          puts "\e[32m    * Try save #{cota[:numero]}\e[0m"
          if Cota.where(numero: cota[:numero]).first
            puts "\e[31m    - Cota #{cota[:numero]} already on database\e[0m"
            puts "    "
            next
          end
          cota[:data_emissao] = cota[:data_emissao].nil? ? Time.new(cota[:ano], cota[:mes], 1, 0, 0, 0, '-03:00') : Time.iso8601("#{cota[:data_emissao]}-03:00")

          puts "\e[32m    * Get deputado: #{cota[:nome_parlamentar]}\e[0m"
          deputado = get_deputado(cota[:nome_parlamentar])

          if deputado.nil?
            puts "\e[31m    - Not found deputado: #{cota[:nome_parlamentar]}\e[0m"
            puts "    "
            open('not_found', 'w') do |f|
              f << "#{cota}\n"
            end
            next
          end
          puts "\e[32m    * Found deputado: #{cota[:nome_parlamentar]} is ##{deputado.id}\e[0m"
          puts "\e[32m    * Creating cota: #{cota[:numero]}\e[0m"

          cota.delete(:nome_parlamentar)
          cota[:deputado_id] = deputado.id
          Cota.create! cota
          puts "    "
        end
      end
    end
    puts "Finished!: #{(start-Time.now) / 60}"
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

  def self.get_deputado(nome_parlamentar)
    deputado = Deputado.where(nome_parlamentar: nome_parlamentar).first
    deputado = Deputado.where(other_name: nome_parlamentar).first unless deputado
    #unless deputado
      #deputado = search(nome_parlamentar, 6)
      #deputado.update_attribute :other_name, nome_parlamentar if deputado
    #end
    deputado
  end

  def self.search(name, words, where_old = nil)
    puts "\e[32m    * Try search #{name} with #{words} words\e[0m"
    return matchs(name, where_old) if words == 0
    regexp = Regexp.new("\\w{#{words}}")
    match = name.match(regexp)
    return matchs(name, Deputado.all) if match.nil? && match.to_s.length == 0
    deputados = (where_old || Deputado).where('nome_parlamentar LIKE ?', "%#{match}%")
    return search(name, 0, Deputado.all) if deputados.count == 0
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
    possibles[match_index] if match_index
  end
end
