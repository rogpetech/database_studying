require 'mysql2'


def database_connection
  Mysql2::Client.new(host: 'localhost', username: 'root', password: '', database: 'rodolfopeixoto')
end

def menu
  puts "\n1. Create table"
  puts "2. Insert data"
  puts "3. Read data"
  puts "4. Update data"
  puts "5. Delete data"
  puts "6. Exit"
  print "Choose an option: "
end

def create_table(client)
  puts "Enter table name: "
  table_name = gets.strip
  puts "Enter columns and types (e.g. id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255) )"
  columns_details = gets.strip
  client.query("CREATE TABLE #{table_name} (#{columns_details})")
  puts "Table created."
end

def insert_data(client)
  puts "Enter table name: "
  table_name = gets.strip
  puts "Enter columns (e.g. id, name)"
  columns = gets.strip
  puts "Enter values (e.g. 'Jhon Doe')"
  values = gets.strip

  client.query("INSERT INTO #{table_name} (#{columns}) VALUES (#{values})")
  puts "Data inserted."
end

def read_data(client)
  puts "Enter table name: "
  table_name = gets.strip
  puts "Enter wildcard name (e.g. * = All columns or Specific = id, name)"
  wildcards = gets.strip

  rows = client.query("SELECT #{wildcards} FROM #{table_name}")
  rows.each do |row|
    row.each do |key, value|
      puts "#{key}: #{value}"
    end
  end
end

def update_data(client)
  puts "Enter table name: "
  table_name = gets.strip
  puts "Enter Set clause (e.g. name='Rodolfo Peixoto' where id = 1)"
  clause = gets.strip

  client.query("UPDATE #{table_name} SET (#{clause})")
  puts "Data updated."
end

def delete_data(client)
  puts "Enter table name: "
  table_name = gets.strip
  puts "Enter with condiction: (e.g. id = 1)"
  condiction = gets.strip
  
  client.query("DELETE FROM #{table_name} where #{condiction}")
end

def main
  begin
    client = database_connection
    loop do
      menu
      option = gets.to_i
      
      case option
      when 1
        create_table(client)
      when 2
        insert_data(client)
      when 3
        read_data(client)
      when 4
        update_data(client)
      when 5
        delete_data(client)
      when 6
        puts "Exiting program."
        exit(0)
      else
        puts "Invalid option."
      end
    end
  rescue Mysql2::Error => e
    puts "Error while connection is MySQL: #{e.message}"
  ensure
    client.close if client
    puts "MySQL connection is closed."
  end
end

main
