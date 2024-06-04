package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

/*
   Adicionar o gorm.Model traz as vantagens de termos mais opracoes possiveis no objeto e
   fará com que o dado no banco possua o create_at, updated_at, deleted_at.
   No SQL podemos ver todos os campos de uma tabela pelo sgdb com o comando:
   desc <nome_tabela>
*/

func main() {
	dsn := "root:password@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}) // Migração para criar um tabela a partir dessa struct

	// db.Create(&Product{ // Como o campo ID é primary key ele vai auto incrementar se vc não passar para ele.
	// 	Name:  "Playstation",
	// 	Price: 49.99,
	// })

	var p1 Product
	db.Where("name like ?", "%station%").Find(&p1, 1)
	fmt.Println("Produto encontrado:", p1)
	p1.Name = "PS2"
	db.Save(&p1)

	var p2 Product
	db.Where("id = ?", p1.ID).Find(&p2)
	fmt.Println("Produto encontrado:", p2)

	fmt.Println()
	db.Delete(&p2)

	/*
				Aqui aconteceu uns erros bizarros.
				Aparentemente pq não passamos o parse time como um argumento no conexão, as comunicações com o banco
				apresentam erro mas não são bloqueantes.
				Para corrigir trocamos
				dsn := "root:password@tcp(localhost:3306)/goexpert"
				para:
				dsn := "root:password@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

				Outro erro é que a lógica está incorreta. Por exemplo, quando fazemos um db.Where("name like ?", "%station%").Find(&p1, 1) e não
				encontramos nenhum objeto corresponente no banco atribuimos falores paradrão nulos para p1.
				Logo p1 é: { 0 "" 0 {0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC {0001-01-01 00:00:00 +0000 UTC false}}}
				Após isso alteramos o nome do produto para PS2 e salvamos devolta ao banco. Isso por sí só não pode ser um update pois
				recupera uma valor inexistente para ser atualizado.

		        O conceito de soft delete, é relacionado a nunca perder a inofrmação mesmo que haja o processo de delete. Para o Gorm
		        manteém um campo de time stamp para quando o objeto for deletado, deleted_at.
	*/
}
