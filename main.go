package main

import "fmt"

var cache_size [7]int
var T_memoria_principal float32 = 50 //nanosegundos

var T_acesso_cache [7]float32
var T_acesso_cache_2way [7]float32

var taxa_cache_miss [7]float32

// Políticas de substituição
var taxa_cache_miss_RANDOM [7]float32
var taxa_cache_miss_FIFO [7]float32
var taxa_cache_miss_LRU [7]float32

func main() {

	// Inserindo valores variádicos
	cache_size = [7]int{16, 32, 64, 128, 256, 512, 1024}
	taxa_cache_miss = [7]float32{0.051, 0.045, 0.041, 0.023, 0.015, 0.01, 0.005}
	T_acesso_cache = [7]float32{0.55, 0.65, 0.70, 0.75, 0.85, 1.10, 1.30}
	T_acesso_cache_2way = [7]float32{0.72, 0.75, 0.85, 1.1, 1.25, 1.60, 2.10}

	taxa_cache_miss_RANDOM = [7]float32{0.04, 0.038, 0.031, 0.02, 0.011, 0.005, 0.002}
	taxa_cache_miss_FIFO = [7]float32{0.036, 0.034, 0.028, 0.012, 0.009, 0.006, 0.003}
	taxa_cache_miss_LRU = [7]float32{0.032, 0.031, 0.02, 0.016, 0.008, 0.005, 0.0024}

	for i := 0; i < len(cache_size); i++ {
		fmt.Printf("Cache Size: %v Kb -----------------------------------------------------\n", cache_size[i])

		tempo_medio := taxa_cache_miss[i]*T_memoria_principal + (1 - taxa_cache_miss[i]*T_acesso_cache[i])
		fmt.Printf("Tempo médio de acesso ao sistema de memória usando mapeamento direto: %.4v ns\n", tempo_medio)

		tempo_medio_random := taxa_cache_miss_RANDOM[i]*T_memoria_principal + (1 - taxa_cache_miss_RANDOM[i]*T_acesso_cache_2way[i])
		fmt.Printf("Tempo médio de acesso usando mapeamento associativo de 2 vias | RANDOM: %.4v ns\n", tempo_medio_random)

		tempo_medio_fifo := taxa_cache_miss_RANDOM[i]*T_memoria_principal + (1 - taxa_cache_miss_RANDOM[i]*T_acesso_cache_2way[i])
		fmt.Printf("Tempo médio de acesso usando mapeamento associativo de 2 vias | FIFO: %.4v ns\n", tempo_medio_fifo)

		tempo_medio_lru := taxa_cache_miss_RANDOM[i]*T_memoria_principal + (1 - taxa_cache_miss_RANDOM[i]*T_acesso_cache_2way[i])
		fmt.Printf("Tempo médio de acesso usando mapeamento associativo de 2 vias | LRU: %.4v ns\n", tempo_medio_lru)

	}
}
