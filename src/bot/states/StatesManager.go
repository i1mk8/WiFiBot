/*
Код для управления состояниями юзера.
Cостояния необходимы, чтобы бот понимал в каком меню находится юзер и исходя из этого корректно реагировал на команды
*/

package states

var (
	users []StatesStruct
)

// Получение текущего состояния юзера
func GetUser(userId int64) (*StatesStruct, *int) {
	for index, user := range users {
		if user.UserId == userId {
			return &user, &index
		}
	}

	return nil, nil
}

// Установка текущего состояния юзера
func SetUser(userId int64, state string) {
	user, index := GetUser(userId)

	if index != nil {
		user.State = state
		users[*index] = *user

	} else {
		stateStruct := &StatesStruct{
			UserId: userId,
			State:  state,
		}
		users = append(users, *stateStruct)
	}
}
