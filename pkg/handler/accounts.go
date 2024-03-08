package handler

import (
	"congo"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// func (h *Handler) filter(c *gin.Context) {
// 	fmt.Print("Yeeees\n")
// }

type getAllListsResponse struct {
	Data []congo.Account `json:"accounts"`
}

func (h *Handler) initFilters() {
	h.types = make(map[string][]string)

	// 	eq - соответствие конкретному полу - "m" или "f"
	h.types["sex"] = []string{"eq"}
	// 	domain - выбрать всех, чьи email-ы имеют указанный домен
	// 	lt - выбрать всех, чьи email-ы лексикографически раньше
	// 	gt - то же, но лексикографически позже
	h.types["email"] = []string{"domain", "lt", "gt"}
	// 	eq - соответствие конкретному статусу
	//  neq - выбрать всех, чей статус не равен указанному
	h.types["status"] = []string{"eq", "neq"}
	// 	eq - соответствие конкретному имени
	// 	any - соответствие любому имени из перечисленных через запятую
	// 	null - выбрать всех, у кого указано имя (если 0) или не указано (если 1)
	h.types["fname"] = []string{"eq", "any", "null"}
	// 	eq - соответствие конкретной фамилии
	//  starts - выбрать всех, чьи фамилии начинаются с переданного префикса
	// 	null - выбрать всех, у кого указана фамилия (если 0) или не указана (если 1)
	h.types["sname"] = []string{"eq", "starts", "null"}
	// 	code - выбрать всех, у кого в телефоне конкретный код (три цифры в скобках)
	// 	null - аналогично остальным полям
	h.types["phone"] = []string{"code", "null"}
	//  eq - всех, кто живёт в конкретной стране
	// 	null - аналогично
	h.types["country"] = []string{"eq", "null"}
	// 	eq - всех, кто живёт в конкретном городе
	// 	any - в любом из перечисленных через запятую городов
	// 	null - аналогично
	h.types["city"] = []string{"eq", "any", "null"}
	// 	lt - выбрать всех, кто родился до указанной даты
	// 	gt - после указанной даты
	// 	year - кто родился в указанном году
	h.types["birth"] = []string{"lt", "gt", "year"}
	// 	contains - выбрать всех, у кого есть все перечисленные интересы
	// 	any - выбрать всех, у кого есть любой из перечисленных интересов
	h.types["interests"] = []string{"contains", "any"}
	// 	contains - выбрать всех, кто лайкал всех перечисленных пользователей (в значении - перечисленные через запятые id)
	h.types["likes"] = []string{"contains"}
	// 	now - все у кого есть премиум на текущую дату
	// 	null - аналогично остальным
	h.types["premium"] = []string{"now", "null"}
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
	filter := []congo.Filter{}

	paramPairs := c.Request.URL.Query()

	// Параметры запроса
	keys := make([]string, 0, len(paramPairs))
	for k := range paramPairs {
		keys = append(keys, k)
	}

	// Проходим по всем параметрам запроса
	filters := make(map[string]bool)
	for _, val := range keys {
		// Получаем параметр запроса и его метод
		typeFilter := strings.Split(val, "_")
		_, exists := filters[typeFilter[0]]
		if exists {
			newErrorResponse(c, http.StatusBadRequest, "duplicate filters "+typeFilter[0])
			return
		}
		filters[typeFilter[0]] = true
		// Если данный запрос есть в API
		method, found := h.types[typeFilter[0]]
		if found {
			// Должено быть 2 значения, параметр поиска и его метод
			if len(typeFilter) != 2 {
				newErrorResponse(c, http.StatusBadRequest, "error parametrs filters "+typeFilter[0])
				return
			}
			// Неизвестный тип метода
			if !slices.Contains(method, typeFilter[1]) {
				newErrorResponse(c, http.StatusBadRequest, "error method filters "+typeFilter[1])
				return
			}

			filter = append(filter, congo.Filter{Filter: typeFilter[0], Method: typeFilter[1], Parametr: paramPairs.Get(val)})
		}

	}

	if _, ok := filters["limit"]; !ok {
		newErrorResponse(c, http.StatusBadRequest, "no limit parameters")
		return
	}

	limit, err := strconv.Atoi(paramPairs.Get("limit"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "limit not a number")
		return
	}

	lists, err := h.services.AccountsList.Filter(filter, limit)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}
