package handler

import (
	"congo"
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

// func (h *Handler) filter(c *gin.Context) {
// 	fmt.Print("Yeeees\n")
// }

type getAllListsResponse struct {
	Data []congo.Account `json:"accounts"`
}

func (h *Handler) getAll(c *gin.Context) {

	lists, err := h.services.AccountsList.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) filter(c *gin.Context) {

	var types = []string{"limit", "sex", "email", "status", "fname", "sname", "phone", "country", "city", "birth", "interests", "likes", "premium"}

	type Filter struct {
		Filter   string // Имя параметра фильрации
		Method   string // Метод фильтрации
		Parametr string // Параметр
	}

	// type Filters struct {
	// 	Limit int // limit - аналогично остальным

	// 	Sex_eq string //eq - соответствие конкретному полу - "m" или "f";

	// 	Email_domain string // domain - выбрать всех, чьи email-ы имеют указанный домен
	// 	Email_lt     string // lt - выбрать всех, чьи email-ы лексикографически раньше
	// 	Email_gt     string // gt - то же, но лексикографически позже

	// 	Status_eq  string // eq - соответствие конкретному статусу
	// 	Status_neq string // neq - выбрать всех, чей статус не равен указанному

	// 	Fname_eq   string // eq - соответствие конкретному имени
	// 	Fname_any  string // any - соответствие любому имени из перечисленных через запятую
	// 	Fname_null string // null - выбрать всех, у кого указано имя (если 0) или не указано (если 1)

	// 	Sname_eq     string // eq - соответствие конкретной фамилии
	// 	Sname_starts string // starts - выбрать всех, чьи фамилии начинаются с переданного префикса
	// 	Sname_null   string // null - выбрать всех, у кого указана фамилия (если 0) или не указана (если 1)

	// 	Phone_code string // code - выбрать всех, у кого в телефоне конкретный код (три цифры в скобках)
	// 	Phone_null string // null - аналогично остальным полям

	// 	Country_eq   string // eq - всех, кто живёт в конкретной стране
	// 	Country_null string // null - аналогично

	// 	City_eq   string // eq - всех, кто живёт в конкретном городе
	// 	City_any  string // any - в любом из перечисленных через запятую городов
	// 	City_null string // null - аналогично

	// 	Birth_lt   string // lt - выбрать всех, кто родился до указанной даты
	// 	Birth_gt   string // gt - после указанной даты
	// 	Birth_year string // year - кто родился в указанном году

	// 	Interests_contains string // contains - выбрать всех, у кого есть все перечисленные интересы
	// 	Interests_any      string // any - выбрать всех, у кого есть любой из перечисленных интересов

	// 	Likes_contains string // contains - выбрать всех, кто лайкал всех перечисленных пользователей (в значении - перечисленные через запятые id)

	// 	Premium_now  string // now - все у кого есть премиум на текущую дату
	// 	Premium_null string // null - аналогично остальным
	// }

	filter := []Filter{}

	paramPairs := c.Request.URL.Query()

	keys := make([]string, 0, len(paramPairs))
	for k := range paramPairs {
		keys = append(keys, k)
	}

	params := make(map[string]bool)
	for _, val := range keys {
		param := strings.Split(val, "_")
		_, exists := params[param[0]]
		if !exists {
			params[param[0]] = true

			found := slices.Contains(types, param[0])
			if found {
				elem := Filter{param[0], "", paramPairs.Get(val)}
				if len(param) == 2 {
					elem.Method = param[1]
				}
				filter = append(filter, elem)
			} else {
				fmt.Println("test ", param)
			}
		} else {
			newErrorResponse(c, http.StatusBadRequest, "duplicate filters "+param[0])
			return
		}
	}
	fmt.Println(filter)

	lists, err := h.services.AccountsList.FilterSex("d")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}
