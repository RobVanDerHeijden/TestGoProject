package main

// Importing Packaging (fmt is a basic package for I/O text)
import (
	"fmt"
)

// UserRepository
// type UserRepository interface {
// 	Create(ctx context.Context, description string, priority internal.Priority, dates internal.Dates) (internal.Task, error)
// 	Find(ctx context.Context, id string) (internal.Task, error)
// 	Update(ctx context.Context, id string, description string, priority internal.Priority, dates internal.Dates, isDone bool) error
// }

func main() {
	// Basic Print function
	fmt.Println("Hello, World!")

	// p := user{name: "Rob", age: 28}
	// fmt.Println(p.age, p.name)

}

// func (productModel ProductModel) Create(product *entities.Product) (int64, error) {
// 	if err != nil {
// 		return 0, err
// 	} else {
// 		product.Id, _ = result.LastInsertId()
// 		rowsAffected, _ := result.RowsAffected()
// 		return rowsAffected, nil
// 	}
// }
