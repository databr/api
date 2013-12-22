class DeputadoImageService
  def self.save_images_from_deputies_json_parser
    parser = DeputiesJsonParser.new('db/data/deputies.json')
    puts "\e[32m  * Starting save_images_from_deputies_json_parser\e[0m"
    parser.deputados.each do |deputado_data|
      deputado = Deputado.find_by_parlamentar_id(deputado_data["id_dep"])
      if deputado
        puts "\e[32m    * Saving image to #{deputado.nome_parlamentar}(#{deputado.id}) from json parser\e[0m"
        deputado.update_attribute :image_url, deputado_data["image_urls"].first
      end
    end
    Deputado.where(image_url: nil).each do |deputado|
      puts "\e[32m    * Adding default image to #{deputado.nome_parlamentar}(#{deputado.id}) from json parser\e[0m"
      deputado.update_attribute :image_url,  "http://www.camara.gov.br/internet/deputado/bandep/#{deputado.cadastro_id}.jpg"
    end
  end
end
