package main

import (
	"encoding/json"
	"fmt"
	"net"
	"sort"
	"strings"
	"time"
)

var mode string = ""

func main() {
	var conn net.Conn
	var connErr error
	for {
		conn, connErr = net.Dial("tcp", "master:8888")
		if connErr == nil {
			break
		} else {
			fmt.Println(connErr)
		}
		time.Sleep(1 * time.Second)
	}

	// conn.Write([]byte("first"))

	buf := make([]byte, 1024)
	for {
		// Read the incoming connection into the buffer.
		reqLen, err := conn.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("Disconned from server")
				break
			} else {
				fmt.Println("Error reading:", err.Error())
				break
			}
		}
		msg := buf[:reqLen]
		// Try to decode to typing mode json
		typeJsonObj := map[string]string{}
		json.Unmarshal(msg, &typeJsonObj)
		if _, ok := typeJsonObj["type"]; ok {
			mode = typeJsonObj["type"]
		}
		nums := []float64{}
		numsJsonObj := map[string][]float64{}
		json.Unmarshal(msg, &numsJsonObj)
		if _, ok := numsJsonObj["nums"]; ok {
			nums = numsJsonObj["nums"]
			result := ""
			switch mode {
			case "Mean":
				result = fmt.Sprintf("%v", ClacMean(nums))
			case "Mode":
				modes := ClacMode(nums)
				result = strings.Join(modes, ",")
			case "Median":
				result = fmt.Sprintf("%v", ClacMedian(nums))
			}
			msg := fmt.Sprintf("%s is %v\n", mode, result)
			conn.Write([]byte(msg))
		}
	}
	defer conn.Close()
}

func ClacMean(nums []float64) float64 {
	total := 0.0
	for _, num := range nums {
		total += num
	}
	return total / float64(len(nums))
}

func ClacMode(nums []float64) []string {
	counts := map[float64]int{}

	for _, num := range nums {
		counts[num]++
	}
	max := 0
	resultMap := map[float64]bool{}
	// result := []string{}
	for num, count := range counts {
		if count > max {
			max = count
			resultMap = map[float64]bool{}
			resultMap[num] = true
		}
		if count == max {
			resultMap[num] = true
		}
	}
	// trun to slice
	result := []string{}
	for num, _ := range resultMap {
		result = append(result, fmt.Sprintf("%v", num))
	}
	return result
}

func ClacMedian(nums []float64) float64 {
	sort.Float64s(nums)
	midNumber := len(nums) / 2

	if len(nums)%2 == 0 {
		return (nums[midNumber-1] + nums[midNumber]) / 2
	} else {
		return nums[midNumber]
	}
}
