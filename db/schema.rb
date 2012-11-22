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

ActiveRecord::Schema.define(:version => 20121122040505) do

  create_table "deputados", :force => true do |t|
    t.string  "nome"
    t.integer "cadastro_id",      :limit => 255
    t.string  "nome_parlamentar"
    t.string  "sexo"
    t.string  "uf"
    t.integer "partido_id"
    t.integer "gabinete",         :limit => 255
    t.integer "anexo",            :limit => 255
    t.string  "fone"
    t.string  "email"
    t.integer "parlamentar_id"
  end

end
