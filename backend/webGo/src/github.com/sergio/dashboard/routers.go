package dashboard

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

type NewAddress struct {
	Status string   `json:"status"`
	Data   *Address `json:"data"`
}
type Address struct {
	Address string `json:"address"`
}

func GetNewAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dat, _ := ioutil.ReadAll(r.Body)
	var params map[string]string
	json.Unmarshal(dat, &params)
	fmt.Println(params["token"])
	t := params["token"]
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("do or do not there is no try"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["sub"])
	} else {
		fmt.Println(err)
	}
	// res, err := http.Get(os.Getenv("BASE_API") + "/get_new_address/?api_key=" + os.Getenv(params["currency"]))
	// if err != nil {
	// 	panic(err.Error())
	// }

	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// data := &NewAddress{
	// 	Data: &Address{},
	// }
	// err = json.Unmarshal(body, data)
	// fmt.Println(data.Data.Address)
	// s2, _ := json.Marshal(data)
	// fmt.Println(string(s2))

	// db := common.DbConn()

	// insForm, err := db.Prepare("INSERT INTO addrs (id_user,address,currency,create_at) VALUES(?,?,?,?)")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// insForm.Exec("barvjsivcst3i5u87eog", data.Data.Address, params["currency"], time.Now())
	// defer db.Close()
	// w.WriteHeader(http.StatusOK)
	// w.Write(s2)
}
