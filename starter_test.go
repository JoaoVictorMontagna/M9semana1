package starter_test

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	starter "github.com/williaminfante/go_test_starter"
)


// Testa a função SayHello para ver se ela cria a mensagem de certinha.
// A ideia aqui é seguir o TDD, ou seja, primeiro a gente cria o teste e depois
// ajusta o código para passar nesse teste.
func TestSayHello(t *testing.T) {
	greeting := starter.SayHello("William")
	assert.Equal(t, "Hello William. Welcome!", greeting)
	another_greeting := starter.SayHello("asdf ghjkl")
	assert.Equal(t, "Hello asdf ghjkl. Welcome!", another_greeting)
}



// Testa a função OddOrEven para ver se ela identifica números pares e ímpares corretamente.
// O TDD aqui ajuda a garantir que a função funcione para todos os casos: positivos, negativos e zero.
func TestOddOrEven(t *testing.T) {
	t.Run("Check Non Negative Numbers", func(t *testing.T) {
		assert.Equal(t, "45 is an odd number", starter.OddOrEven(45))
		assert.Equal(t, "42 is an even number", starter.OddOrEven(42))
		assert.Equal(t, "0 is an even number", starter.OddOrEven(0))
	})
	t.Run("Check Negative Numbers", func(t *testing.T) {
		assert.Equal(t, "-45 is an odd number", starter.OddOrEven(-45))
		assert.Equal(t, "-42 is an even number", starter.OddOrEven(-42))
	})
}



// Testa a função Checkhealth para garantir que ela retorna o status de saúde certinho.
// Esse teste simula uma requisição HTTP para ver se o endpoint está respondendo como deveria.

func TestCheckhealth(t *testing.T) {
	t.Run("Check health status", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://mysite.com/example", nil)
		writer := httptest.NewRecorder()
		starter.Checkhealth(writer, req)
		response := writer.Result()
		body, err := io.ReadAll(response.Body)

		assert.Equal(t, "health check passed", string(body))
		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t,
			"text/plain; charset=utf-8",
			response.Header.Get("Content-Type"))

		assert.Equal(t, nil, err)
	})
}
