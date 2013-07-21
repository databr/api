require './config/boot'

module SocialCamara
  class API < Grape::API
    format :json

    resource :deputados do
      desc "Returns all deputados"
      get do
        Deputado.all
      end

      get ":uri" do
        Deputado.find_by_uri(params[:uri])
      end

      get ":uri/feed" do
        deputado = Deputado.find_by_uri(params[:uri])
        deputado.cotas
      end
    end
  end
end
