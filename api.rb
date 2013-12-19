require './config/boot'

module SocialCamara
  class API < Grape::API
    format :json
    logger LOGGER
    helpers do
      def logger
        LOGGER
      end
    end

    resource :deputados do
      desc "Returns all deputados"
      get do
        if ENV['USE_CACHE'] == 'true' && deputados = REDIS.hgetall('dep')
          deputados.map{|i, d| Oj.load(d) }
        else
          Deputado.all.map{ |deputado|
             REDIS.hset('dep', deputado.id, Oj.dump(deputado.attributes) )
             deputado
          }
        end
      end

      get ":uri" do
        Deputado.find_by_uri(params[:uri])
      end

      get ":uri/about" do
        deputado = Deputado.find_by_uri(params[:uri])
        About.where(cadastro_id: deputado.cadastro_id)
      end

      get ":uri/feed" do
        # deputado = Deputado.find_by_uri(params[:uri])
        # cotas = CotaEntity.new(deputado.cotas.limit(30)).results
        # videos = VideoEntity.new(deputado.videos.limit(30)).results
        Aggregator.build([[], []])
      end
    end
  end
end


