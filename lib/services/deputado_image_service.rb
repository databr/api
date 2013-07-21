class DeputadoImageService
  def self.save_images_from_deputies_json_parser
    parser = DeputiesJsonParser.new('db/data/deputies.json')
    parser.deputados.each do |deputado_data|
      deputado = Deputado.find_by_parlamentar_id(deputado_data["id_dep"])
      deputado.update_attribute :image_url, deputado_data["image_urls"].first if deputado
    end
  end
end
