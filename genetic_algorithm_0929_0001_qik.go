// 代码生成时间: 2025-09-29 00:01:40
package main

import (
    "fmt"
    "math/rand"
    "time"
)

// GeneticAlgorithm 封装遗传算法的参数和函数
type GeneticAlgorithm struct {
    // 基因编码的长度
    geneLength int
    // 种群大小
    populationSize int
    // 交叉率
    crossoverRate float64
    // 变异率
    mutationRate float64
    // 适应度函数
    fitnessFunction func([]int) float64
    // 种群
    population [][]int
}

// NewGeneticAlgorithm 初始化遗传算法
func NewGeneticAlgorithm(geneLength, populationSize int, crossoverRate, mutationRate float64, fitnessFunction func([]int) float64) *GeneticAlgorithm {
    ga := &GeneticAlgorithm{
        geneLength:     geneLength,
        populationSize: populationSize,
        crossoverRate:  crossoverRate,
        mutationRate:   mutationRate,
        fitnessFunction: fitnessFunction,
    }
    ga.createInitialPopulation()
    return ga
}

// createInitialPopulation 随机生成初始种群
func (ga *GeneticAlgorithm) createInitialPopulation() {
    ga.population = make([][]int, ga.populationSize)
    for i := range ga.population {
        ga.population[i] = make([]int, ga.geneLength)
        for j := range ga.population[i] {
            ga.population[i][j] = rand.Intn(2) // 假设是二进制编码
        }
    }
}

// EvaluatePopulation 计算整个种群的适应度
func (ga *GeneticAlgorithm) EvaluatePopulation() []float64 {
    fitnessScores := make([]float64, ga.populationSize)
    for i, individual := range ga.population {
        fitnessScores[i] = ga.fitnessFunction(individual)
    }
    return fitnessScores
}

// SelectParents 选择父代个体
func (ga *GeneticAlgorithm) SelectParents(fitnessScores []float64) (parent1, parent2 []int) {
    // 这里可以添加更复杂的选择算法，如轮盘赌选择、锦标赛选择等
    // 简单随机选择两个不同的父代
    r1, r2 := rand.Intn(ga.populationSize), rand.Intn(ga.populationSize)
    for r1 == r2 {
        r2 = rand.Intn(ga.populationSize)
    }
    return ga.population[r1], ga.population[r2]
}

// Crossover 交叉产生新的个体
func (ga *GeneticAlgorithm) Crossover(parent1, parent2 []int) []int {
    child := make([]int, ga.geneLength)
    crossoverPoint := rand.Intn(ga.geneLength)
    for i := 0; i < crossoverPoint; i++ {
        child[i] = parent1[i]
    }
    for i := crossoverPoint; i < ga.geneLength; i++ {
        child[i] = parent2[i]
    }
    return child
}

// Mutate 变异
func (ga *GeneticAlgorithm) Mutate(child []int) {
    for i := range child {
        if rand.Float64() < ga.mutationRate {
            child[i] = 1 - child[i]
        }
    }
}

// EvolvePopulation 进化种群
func (ga *GeneticAlgorithm) EvolvePopulation() {
    newPopulation := make([][]int, ga.populationSize)
    fitnessScores := ga.EvaluatePopulation()
    for i := 0; i < ga.populationSize; i++ {
        parent1, parent2 := ga.SelectParents(fitnessScores)
        child := ga.Crossover(parent1, parent2)
        ga.Mutate(child)
        newPopulation[i] = child
    }
    ga.population = newPopulation
}

// Run 运行遗传算法
func (ga *GeneticAlgorithm) Run(generations int) []int {
    for i := 0; i < generations; i++ {
        ga.EvolvePopulation()
    }
    // 选择最佳个体
    bestFitness := -1.0
    bestIndividual := nil
    for _, individual := range ga.population {
        fitness := ga.fitnessFunction(individual)
        if fitness > bestFitness {
            bestFitness = fitness
            bestIndividual = individual
        }
    }
    return bestIndividual
}

func main() {
    rand.Seed(time.Now().UnixNano())

    // 示例适应度函数：简单的最大值适应度
    fitnessFunction := func(gene []int) float64 {
        sum := 0
        for _, value := range gene {
            sum += value
        }
        return float64(sum)
    }

    ga := NewGeneticAlgorithm(10, 100, 0.8, 0.01, fitnessFunction)
    bestGene := ga.Run(1000)
    fmt.Printf("Best gene: %v
", bestGene)
}
