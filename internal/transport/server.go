package transport

import (
	"github.com/bulbosaur/dish-hub/internal/dao"
	"github.com/spf13/viper"
)

func RunHttp(repo *dao.DishHubModel) {
	host := viper.GetString("http.HOST")
	port := viper.GetString("http.PORT")

	addr := host + ":" + port

	print(addr)
}
