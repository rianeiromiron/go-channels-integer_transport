package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Creamos un canal para enviar números enteros
	canal := make(chan int)
	// Productor (Goroutine emisora)
	go func() {
		fmt.Println("Prueba de función anónima")
	}()
	go func() {
		for i := 1; i <= 5; i++ {
			n := rand.Intn(100)
			fmt.Printf("📤 Productor: Enviando %d...\n", n)
			canal <- n
			time.Sleep(500 * time.Millisecond) // Simulamos una pequeña demora
		}

		// IMPORTANTE: Cerramos el canal cuando terminamos de enviar
		fmt.Println("🔒 Productor: Ya no hay más datos. Cerrando el canal...")
		close(canal)
	}()

	// Consumidor (Goroutine principal 'main')
	fmt.Println("📥 Consumidor (Main): Esperando y recibiendo datos...")

	// El bucle 'range' lee del canal automáticamente hasta que el canal sea CERRADO
	for numero := range canal {
		fmt.Printf("✅ Consumidor (Main): Recibido el número %d\n", numero)
	}

	fmt.Println("🏁 Main: El canal fue cerrado y se procesaron todos los datos. ¡Fin del programa!")
}
