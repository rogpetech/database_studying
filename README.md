PROBLEMA

Você foi designado para desenvolver um sistema de gerenciamento de blog onde os usuários podem escrever artigos, comentar em artigos de outros usuários, interagir com o conteúdo e favoritar os artigos que mais gostam. Você precisa criar e gerenciar o banco de dados para este sistema, utilizando as seguintes informações sobre as tabelas:

Tabelas:

Users (Usuários):

id (Chave primária, inteiro)

username (Texto, até 100 caracteres)

first_name (Texto, até 100 caracteres)

email (Texto, até 50 caracteres)

password (Texto, até 100 caracteres)

avatar_url (Texto, até 100 caracteres)

birthdate (Data)

last_login (Data e hora)

is_admin (Booleano)

Articles (Artigos):

id (Chave primária, inteiro)

user_id (Chave estrangeira referenciando id da tabela Users)

title (Texto, até 100 caracteres)

content (Texto, até 2000 caracteres)

category (Texto)

tags (Texto)

created_at (Data e hora)

updated_at (Data e hora)

views (Inteiro)

likes (Inteiro)

is_published (Booleano)

published_at (Data e hora)

Comments (Comentários):

id (Chave primária, inteiro)

article_id (Chave estrangeira referenciando id da tabela Articles)

user_id (Chave estrangeira referenciando id da tabela Users)

content (Texto, até 2000 caracteres)

parent_comment_id (Chave estrangeira referenciando id da tabela Comments)

is_deleted (Booleano)

deleted_at (Data e hora)

created_at (Data e hora)

updated_at (Data e hora)

Favorites (Favoritos):

id (Chave primária, inteiro)

user_id (Chave estrangeira referenciando id da tabela Users)

article_id (Chave estrangeira referenciando id da tabela Articles)

favorited_at (Data e hora)

Relacionamentos:

Um usuário pode ter vários artigos, mas um artigo pertence a apenas um usuário. Relação de "um para muitos" entre Users e Articles.

Um artigo pode ter vários comentários, e um comentário pertence a apenas um artigo. Relação de "um para muitos" entre Articles e Comments.

Um usuário pode fazer vários comentários, mas cada comentário é feito por apenas um usuário. Relação de "um para muitos" entre Users e Comments.

Um usuário pode favoritar vários artigos e um artigo pode ser favoritado por vários usuários. Relação de "muitos para muitos" entre Users e Articles através da tabela Favorites.





### Install Debugger Go

https://github.com/golang/vscode-go/blob/master/docs/debugging.md