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
    end
  end
end
