package domain

//Separa el modelo de la capa de negocio
//que necesitan hacer mis actores con los casos de uso
type ICustomer interface {
	Save(customers Customer) (uint, error) //podemos retornar algo acá, como un string, etc.
	GetAll() ([]Customer) //tenemos nuestros métodos los cuales los usuarios hacen la acción con nuestra base de datos
	Delete(id int) (uint, error)
	Update(id int, customer Customer) (uint, error)
}