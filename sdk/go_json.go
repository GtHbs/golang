package sdk

import (
	"encoding/json"
	"fmt"
	"golang/base"
)

func BaseJson() {
	var user base.UserInfo
	s := `{"name":"alone","age":20}`
	json.Unmarshal([]byte(s), &user)
	fmt.Println("sss", user)
}
