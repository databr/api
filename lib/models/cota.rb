class Cota < ActiveRecord::Base
  self.table_name = "cotas"
  #validates :numero, uniqueness: true
end
