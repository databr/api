# encoding: UTF-8
# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# Note that this schema.rb definition is the authoritative source for your
# database schema. If you need to create the application database on another
# system, you should be using db:schema:load, not running all the migrations
# from scratch. The latter is a flawed and unsustainable approach (the more migrations
# you'll amass, the slower it'll run and the greater likelihood for issues).
#
# It's strongly recommended to check this file into your version control system.

ActiveRecord::Schema.define(:version => 20130721033800) do

  create_table "cotas", :force => true do |t|
    t.integer  "deputado_id"
    t.string   "carteira_parlamentar"
    t.string   "legislatura"
    t.string   "uf"
    t.string   "partido"
    t.integer  "partido_id"
    t.string   "codigo_legislatura"
    t.string   "sub_cota"
    t.string   "descricao"
    t.string   "beneficiario"
    t.string   "cnpjcpf"
    t.string   "numero"
    t.string   "tipo_documento"
    t.datetime "data_emissao"
    t.decimal  "valor_documento",      :precision => 10, :scale => 2
    t.decimal  "valor_glossa",         :precision => 10, :scale => 2
    t.decimal  "valor_liquido",        :precision => 10, :scale => 2
    t.integer  "mes"
    t.integer  "ano"
    t.string   "parcela"
    t.string   "lote"
    t.string   "ressarcimento"
  end

  create_table "deputados", :force => true do |t|
    t.integer "partido_id"
    t.integer "cadastro_id"
    t.integer "parlamentar_id"
    t.string  "nome"
    t.string  "nome_parlamentar"
    t.string  "sexo"
    t.string  "uf"
    t.integer "gabinete"
    t.integer "anexo"
    t.string  "fone"
    t.string  "email"
    t.string  "uri"
    t.string  "image_url"
  end

  create_table "events", :force => true do |t|
    t.integer  "session_id"
    t.datetime "starts_at"
    t.string   "local"
    t.text     "title"
    t.text     "description"
  end

  create_table "videos", :force => true do |t|
    t.integer "event_id"
    t.integer "deputado_id"
    t.string  "video_url"
  end

end
