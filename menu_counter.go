package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type MenuCounter struct {
	counts map[string]int
}

func NewMenuCounter() *MenuCounter {
	return &MenuCounter{
		counts: make(map[string]int),
	}
}

func (mc *MenuCounter) ProcessLog(logPath string) error {
	file, err := os.Open(logPath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")

		if len(fields) != 2 {
			return fmt.Errorf("invalid log format: %s", line)
		}

		eaterID := fields[0]
		menuID := fields[1]
		if _, ok := mc.counts[eaterID]; ok {
			return fmt.Errorf("duplicate entry found: %s", line)
		}

		mc.counts[menuID]++
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (mc *MenuCounter) GetTopMenuItems(numItems int) []string {
	type menuItem struct {
		ID    string
		Count int
	}

	topItems := make([]menuItem, 0, len(mc.counts))
	for id, count := range mc.counts {
		topItems = append(topItems, menuItem{ID: id, Count: count})
	}

	// Sort the items based on count in descending order
	for i := 0; i < len(topItems)-1; i++ {
		for j := 0; j < len(topItems)-i-1; j++ {
			if topItems[j].Count < topItems[j+1].Count {
				topItems[j], topItems[j+1] = topItems[j+1], topItems[j]
			}
		}
	}

	// Retrieve the top numItems menu IDs
	topMenuItems := make([]string, 0, numItems)
	for i := 0; i < numItems && i < len(topItems); i++ {
		topMenuItems = append(topMenuItems, topItems[i].ID)
	}

	return topMenuItems
}

func main() {
	menuCounter := NewMenuCounter()

	err := menuCounter.ProcessLog("log.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	topItems := menuCounter.GetTopMenuItems(3)
	fmt.Println("Top 3 Menu Items Consumed:")
	for i, item := range topItems {
		fmt.Printf("%d. %s\n", i+1, item)
	}
}
