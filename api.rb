require './config/boot'

module SocialCamara
  class API < Grape::API
    format :json

    resource :deputados do
      desc "Returns all deputados"
      get do
        Deputado.all
      end

      get ":id" do
        Deputado.find(params[:id])
      end

      get ":id/feed" do
        deputado = Deputado.find(params[:id])
        deputado.cotas
      end
    end
  end
end
