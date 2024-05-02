import mysql.connector

from mysql.connector import Error

def main():
  try:
    database_connection = mysql.connector.connect(
      host="localhost",
      port="3306",
      user="root",
      password="",
      )
    cursor = database_connection.cursor()

    database_name = input('Enter database name: ')

    cursor.execute(f"SHOW DATABASES LIKE '{database_name}'")
    result = cursor.fetchone()
    if not result:
      print(f"Database {database_name} does not exists. Creating...")
      cursor.execute(f"CREATE DATABASE {database_name}")
      print("Database created.")
    
    database_connection.database = database_name
    
    while True:
      print("\n1. Create table")
      print("2. Insert data")
      print("3. Read data")
      print("4. Update data")
      print("5. Delete data")
      print("6. Exit")

      option = int(input("Choose an option: "))

      match option:
        case 1:
          create_table(cursor)
        case 2:
          insert_data(database_connection)
        case 3:
          read_data(cursor)
        case 4:
          update_data(database_connection)
        case 5:
          delete_data(database_connection)
        case 6:
          print("Exiting program.")
          break
        case _:
          print("===================")
          print("Invalid option.")
          print("===================")
  except Error as e:
    print("Error while connecting to MySQL", e)
  finally:
    if database_connection.is_connected:
      cursor.close()
      database_connection.close()
      print("MySQL connection is closed.")

def create_table(cursor):
  table_name = input("Enter table name: ")
  columns_details = input("Enter columns and types (e.g., id INT AUTOINCREMENT PRIMARY KEY, name VARCHAR(255)): ")

  cursor.execute(f"CREATE TABLE {table_name} ({columns_details})")

def insert_data(database_connection):
  cursor = database_connection.cursor()
  table_name = input("Enter table name: ")

  columns = input("Enter columns: (e.g., id, name) ")
  values = input("Enter values: (e.g., 1, 'Jhon Doe') ")
  query = f"INSERT INTO {table_name} ({columns}) VALUES ({values})"
  cursor.execute(query)
  database_connection.commit()
  print("Data inserted.")


def read_data(cursor):
  table_name = input("Enter table name: ")
  wildcard = input("Enter wildcard: ")

  cursor.execute(f"SELECT {wildcard} FROM {table_name}")
  rows = cursor.fetchall()

  for row in rows:
    print(row)
  print('--------------')

def update_data(database_connection):
  cursor = database_connection.cursor()
  table_name = input("Enter table name: ")
  set_clause = input("Enter SET clause (e.g. name='Jane Doe' where id = 1) ")
  cursor.execute(f"UPDATE {table_name} SET {set_clause}")
  database_connection.commit()
  print("Data updated.")

def delete_data(database_connection):
  table_name = input("Enter table name: ")
  condiction = input("Enter condiction: (e.g. id=1) ")

  database_connection.cursor.execute(f"DELETE FROM {table_name} WHERE {condiction}")
  database_connection.commit()

main()
