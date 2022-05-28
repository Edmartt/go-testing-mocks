package service

import (
	"go-testing-mocks/internal/users"
	"go-testing-mocks/internal/users/repository"
	"testing"
)


func TestMakeUserAvailable(t *testing.T){	

	mockedDBUser := users.User{
		Id: "1",
		Name: "John",
	}


	table := []struct{
		expectedState string
		mockFuncOne func()
		userForPass users.User
		userForFail users.User
	}{
		{
			expectedState: "on",
			mockFuncOne: func(){
				//The behaviour of our function is changed for using the data we want
				//without calling an external service like DB
				repository.UserExist = func(username string) bool {
					if username == mockedDBUser.Name{
						return true
					}
					return false
				}
			},

			userForPass: users.User{
				Id: "1",
				Name: "Edinson",
			},
			userForFail: users.User{
				Id: "1",
				Name: "John",
			},
		},

	}

	originalExist := repository.UserExist

	for _, test := range table{
		test.mockFuncOne()

		value := makeUserAvailable(test.userForPass)

		if test.expectedState != value{
			t.Errorf("expected %s and got %s", test.expectedState, value)
		}

		newValue := makeUserAvailable(test.userForFail)

		if test.expectedState == newValue{
			t.Errorf("expected to be different")
		}
	}

	repository.UserExist = originalExist
}


/*This section is for mocking the CreateUser function and changing it's behaviour too, because it calls a database,
and we do it implementing the same interface as we do with the original repository*/


//The mock of our createUser function
var createUserMock func(user users.User)(*string, error)

//This struct is for adding a method that implements the Repository interface
type RepositoryMock struct{}

//This method implements the repository interface and instead of calling a database function, we return a mock of it with
//the data we want
func(r RepositoryMock) CreateUser(user users.User)(*string, error){
	return createUserMock(user)
}


func TestCreateUser(t *testing.T){
	dbUserForChecking := users.User{
		Id: "1",
		Name: "Michael",
	}

	//In the testing context our var is turning into an instance of our repository mock
	//and this is very helpful because in this context the method that our RegisterUser function will use
	//for creating our users is our createUserMock instead of the function using third party service or dependency
	repository.AccessRepo = RepositoryMock{}

	//Definition of the behaviour we want for our createUser function
	createUserMock = func(user users.User) (*string, error) {
		return &user.Id, nil
	}

	table := []struct{
		userForCreate users.User
		expectedResponse string
		badResponse string
		mockFuncCheck func()
	}{
		{
			userForCreate: users.User{
				Id: "1",
				Name: "John Doe",
			}, 

			expectedResponse: "1", 

			badResponse: "User already exists",

			mockFuncCheck: func() {
				repository.UserExist = func(username string) bool {
					if username == dbUserForChecking.Name{
						return true
					}
					return false
				}
			},
		},
		{
			userForCreate: users.User{
				Id: "23",
				Name: "Jane Doe",
			}, 

			expectedResponse: "23", 
			badResponse: "User already exists",

			mockFuncCheck:func(){
				repository.UserExist = func(username string) bool {
					if username == dbUserForChecking.Name{
						return true
					}
					return false
				}
			},
		},
{
			userForCreate: users.User{
				Id: "10",
				Name: "Michael",
			}, 

			expectedResponse: "User already exists", 
			badResponse: "10",

			mockFuncCheck:func(){
				repository.UserExist = func(username string) bool {
					if username == dbUserForChecking.Name{
						return true
					}
					return false
				}
			},
		},
	}

	originalExist := repository.UserExist

	for _, property := range table{
		property.mockFuncCheck()

		response := RegisterUser(property.userForCreate)

		if response != property.expectedResponse{
			t.Fatalf("Expected response is %s and got %s", property.expectedResponse, response)
		}
	}

	repository.UserExist = originalExist
}
