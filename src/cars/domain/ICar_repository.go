package domain

//Separa el modelo de la capa de negocio
//que necesitan hacer mis actores con los casos de uso
type ICar interface {
	Save(cars Car) (uint, error) //podemos retornar algo acá, como un string, etc.
	GetAll() ([]Car) //tenemos nuestros métodos los cuales los usuarios hacen la acción con nuestra base de datos
	Delete(id int) (uint, error)
	Update(id int, car Car) (uint, error)
}