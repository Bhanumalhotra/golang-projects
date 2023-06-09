package GetUser


import(
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/dynamodb"
	"github.com/aws/aws-sdk-go/dynamodb/dynamodbattribute"

	"github.com/aws/aws-sdk-go/dynamodb/dynamodbiface"

)

var(
    ErrorFailedToFetchRecord = "Failed to Fetch Record"
)

type User struct{
	Email         string  `json:"email"`
	FirstName     string  `json:"firstName"`
	LastName      string  `json:"lastName"`
}
func FetchUser(email, tableName string, dynamodbiface.DynamoDBAPI)(*User, error){
	input := &dynamodb.GetItemInput{
		key: map[string]*dynamodb.AttributeValue{
			"email":{
				S: aws.String(email)

			}
			
		},
		TableName: aws.String(tableName)
	}
	result, err := dynaClient.GetItem(input)
	if err!= nil {
		return nil, error.New(ErrorFailedToFetchRecord)
	}

	item := new(User)
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}
    return item, nil
}

func FetchUsers(tableName string, dynaClient dynamodbiface.DynamoDBAPI)(*[]Users){

}

func CreateUser()(){
	
}

func UpdateUser()(){

}

func DeletUser()error{
	
}