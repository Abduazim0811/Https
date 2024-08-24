package items


type Items struct{
	Id 		string 			`json:"_id" bson:"_id"`
	Name 	string			`json:"name" bson:"name"`
	Amount	int				`json:"amount" bson:"amount"`
}

type CreateItems struct{
	Name 	string			`json:"name" bson:"name"`
	Amount	int				`json:"amount" bson:"amount"`
}

type GetItems struct{
	Id 		string 			`json:"_id" bson:"_id"`
}