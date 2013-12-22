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
        Deputado.allcached
      end

      get ":uri" do
        Deputado.cached(params[:uri])
      end

      get ":uri/about" do
        Deputado.about(params[:uri])
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


