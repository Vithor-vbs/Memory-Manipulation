## Trabalho de Desempenho em cache:

Matrícula: **2214624 / 6**

- a) Calcular o **tempo de acesso à cache** para cada tamanho de memória cache usando
  mapeamento direto;

        - 16Kb -> 0.55 ns
        - 32Kb -> 0.65 ns
        - 64Kb -> 0.70 ns
        - 128Kb -> 0.75 ns
        - 256Kb -> 0.85 ns
        - 512Kb -> 1.10 ns
        - 1024Kb -> 1.30 ns

- b) Calcular o **tempo de acesso à cache** para cada tamanho de memória cache usando
  mapeamento associativo por conjuntos de 2 vias;

      - 16Kb -> 0.72 ns
      - 32Kb -> 0.75 ns
      - 64Kb -> 0.85 ns
      - 128Kb -> 1.1 ns
      - 256Kb -> 1.25 ns
      - 512Kb -> 1.60 ns
      - 1024Kb -> 2.10 ns

- c) Calcular a **taxa de cache miss** para cada tamanho de memória cache usando
  mapeamento direto;

            - 16Kb -> 0.051
            - 32Kb -> 0.045
            - 64Kb -> 0.041
            - 128Kb -> 0.023
            - 256Kb -> 0.015
            - 512Kb -> 0.01
            - 1024Kb -> 0.005

- d) Calcular a **taxa de cache miss** para cada tamanho de memória cache usando
  mapeamento associativo por conjuntos de 2 vias;

          - 16Kb -> 0.04
          - 32Kb -> 0.038
          - 64Kb -> 0.031
          - 128Kb -> 0.02
          - 256Kb -> 0.011
          - 512Kb -> 0.005
          - 1024Kb -> 0.002

- e) Calcular a **taxa de cache miss** para cada tamanho de memória cache no
  mapeamento associativo por conjuntos de 2 vias usando cada uma das políticas de substituição
  **RANDOM, FIFO e LRU**;

  - RANDOM:

        - 16Kb -> 0.04
        - 32Kb -> 0.038
        - 64Kb -> 0.031
        - 128Kb -> 0.02
        - 256Kb -> 0.011
        - 512Kb -> 0.005
        - 1024Kb -> 0.002

  - FIFO: Diminui em 10% a probabilidade de cache miss

        - 16Kb -> 0.036
        - 32Kb -> 0.034
        - 64Kb -> 0.028
        - 128Kb -> 0.012
        - 256Kb -> 0.009
        - 512Kb -> 0.006
        - 1024Kb -> 0.003

  - LRU: Diminui em 20% a probabilidade de cache miss

          - 16Kb -> 0.032
          - 32Kb -> 0.031
          - 64Kb -> 0.02
          - 128Kb -> 0.016
          - 256Kb -> 0.008
          - 512Kb -> 0.005
          - 1024Kb -> 0.0024

- f) Calcular o tempo médio de acesso ao sistema de memória para cada tamanho de
  memória cache usando mapeamento direto;
