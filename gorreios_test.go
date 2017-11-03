package gorreios

import "testing"

import "github.com/stretchr/testify/assert"
import "time"

var expectedTracking = Tracking{
	ID: "PO875228513BR",
	Events: Events{
		{
			Description: "Objeto encaminhado de Unidade de Tratamento em PORTO ALEGRE/RS para Unidade de Distribuição em PORTO ALEGRE/RS",
			DateRaw:     "01/11/2017 20:32 PORTO ALEGRE/RS",
			DateTime:    time.Date(2017, time.November, 1, 20, 32, 0, 0, timeLocation),
			City:        "PORTO ALEGRE/RS",
		},
		{
			Description: "Objeto encaminhado de Unidade de Tratamento em RIO DE JANEIRO/RJ para Unidade de Tratamento em PORTO ALEGRE/RS",
			DateRaw:     "30/10/2017 14:32 RIO DE JANEIRO/RJ",
			DateTime:    time.Date(2017, time.October, 30, 14, 32, 0, 0, timeLocation),
			City:        "RIO DE JANEIRO/RJ",
		},
		{
			Description: "Objeto encaminhado de Agência dos Correios em Rio De Janeiro/RJ para Unidade de Tratamento em RIO DE JANEIRO/RJ",
			DateRaw:     "24/10/2017 18:31 Rio De Janeiro/RJ",
			DateTime:    time.Date(2017, time.October, 24, 18, 31, 0, 0, timeLocation),
			City:        "RIO DE JANEIRO/RJ",
		},
		{
			Description: "Objeto postado",
			DateRaw:     "24/10/2017 16:25 Rio De Janeiro/RJ",
			DateTime:    time.Date(2017, time.October, 24, 16, 25, 0, 0, timeLocation),
			City:        "RIO DE JANEIRO/RJ",
		},
	},
}

func TestGetTrackingInfo(t *testing.T) {
	tracking, err := GetTrackingInfo("PO875228513BR")
	assert.Nil(t, err)
	assert.Equal(t, expectedTracking.ID, tracking.ID)
	for i := 0; i < len(tracking.Events); i++ {
		assert.Equal(t, expectedTracking.Events[i], tracking.Events[i])
	}
}
