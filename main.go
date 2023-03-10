package main

import (
	"fmt"
	"math/big"
	"math/bits"
	"time"

	"github.com/d9c818ae24/multiexp"
	"github.com/d9c818ae24/rsa_accumulator/accumulator"
	"github.com/d9c818ae24/rsa_accumulator/experiments"
)

func testPreCompute() {
	setSize := 65536 // 2 ^ 16 65536
	fmt.Println("Test set size = ", setSize)
	fmt.Println("GenRepresentatives with MultiDIHashFromPoseidon")
	set := accumulator.GenBenchSet(setSize)
	setup := *accumulator.TrustedSetup()
	rep := accumulator.GenRepresentatives(set, accumulator.MultiDIHashFromPoseidon)

	startingTime := time.Now().UTC()
	//accumulator.ProveMembershipParallel(setup.G, setup.N, rep, 2)
	duration := time.Now().UTC().Sub(startingTime)
	fmt.Printf("Running ProveMembershipParallel2 Takes [%.3f] Seconds \n", duration.Seconds())

	startingTime = time.Now().UTC()
	maxLen := setSize * 256 / bits.UintSize
	table := multiexp.NewPrecomputeTable(setup.G, setup.N, maxLen)
	duration = time.Now().UTC().Sub(startingTime)
	fmt.Printf("Running PreCompute Takes [%.3f] Seconds \n", duration.Seconds())
	startingTime = time.Now().UTC()
	accumulator.ProveMembershipParallelWithTable(setup.G, setup.N, rep, 2, table)
	duration = time.Now().UTC().Sub(startingTime)
	fmt.Printf("Running ProveMembershipParallelWithTable Takes [%.3f] Seconds \n", duration.Seconds())
}

func testBigInt() {
	var temp, temp2 big.Int
	temp.SetInt64(1024)
	bytes := temp.Bytes()
	fmt.Println("byte[0] = ", bytes[0])
	fmt.Println("byte[1] = ", bytes[1])
	//fmt.Println("byte[2] = ", bytes[2])
	//fmt.Println("byte[4] = ", bytes[3])
	temp2.SetBytes(bytes)
	fmt.Println("temp2 = ", temp2.String())

	tempBits := temp.Bits()
	fmt.Println("bit[0] = ", tempBits[0])
	tempBits[0]++
	//fmt.Println("bit[1] = ", bits[1])
	temp2.SetBits(tempBits)
	fmt.Println("temp = ", temp.String())
	fmt.Println("temp2 = ", temp2.String())
}

func testExp() {
	setup := *accumulator.TrustedSetup()
	var ret1, ret2 big.Int
	ret1.Exp(setup.G, setup.G, setup.N)
	ret2.Exp(setup.G, setup.N, setup.N)
	temp := multiexp.DoubleExp(setup.G, [2]*big.Int{setup.G, setup.N}, setup.N)
	temp2 := multiexp.FourfoldExp(setup.G, setup.N, [4]*big.Int{setup.G, setup.N, setup.G, setup.N})
	fmt.Println("ret1 in main = ", ret1.String())
	fmt.Println("ret1.2 in main = ", ret2.String())
	fmt.Println("ret2 in main = ", temp[0].String())
	fmt.Println("ret3 in main = ", temp[1].String())
	fmt.Println("---ret4 in main = ", temp2[0].String())
	fmt.Println("ret5 in main = ", temp2[1].String())
	fmt.Println("ret6 in main = ", temp2[2].String())
	fmt.Println("ret7 in main = ", temp2[3].String())
}

func main() {
	experiments.TestBasicZKrsa()
	experiments.TestDifferentGroupSize()
	experiments.TestDifferentMembership()
	experiments.TestDifferentMembershipForDI()
	experiments.TestDifferentPrecomputationTableSize()
	experiments.TestFirstLayerPercentage()
	experiments.TestMembership()
	experiments.TestProduct()
	experiments.TestProduct2()
	experiments.TestProduct3()
	experiments.TestRange()
}
