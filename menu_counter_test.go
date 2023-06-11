package main

import (
	"os"
	"testing"
)

func TestMenuCounter_ProcessLog(t *testing.T) {
	menuCounter := NewMenuCounter()

	err := menuCounter.ProcessLog("testlog.txt")
	if err != nil {
		t.Errorf("error processing log: %v", err)
		return
	}

	// Check if the counts are accurate
	expectedCounts := map[string]int{
		"menu1": 2,
		"menu2": 3,
		"menu3": 1,
	}
	for id, expectedCount := range expectedCounts {
		if count, ok := menuCounter.counts[id]; !ok || count != expectedCount {
			t.Errorf("incorrect count for menu ID %s, expected: %d, got: %d", id, expectedCount, count)
		}
	}

	// Check for a duplicate entry
	err = menuCounter.ProcessLog("testlog_duplicate.txt")
	if err == nil {
		t.Error("expected error for duplicate entry, got nil")
	}
}

func TestMenuCounter_GetTopMenuItems(t *testing.T) {
	menuCounter := NewMenuCounter()

	// Prepare some counts for testing
	menuCounter.counts = map[string]int{
		"menu1": 4,
		"menu2": 2,
		"menu3": 6,
		"menu4": 1,
		"menu5": 3,
	}

	// Test getting top 3 items
	expectedTopItems := []string{"menu3", "menu1", "menu5"}
	topItems := menuCounter.GetTopMenuItems(3)
	if len(topItems) != len(expectedTopItems) {
		t.Errorf("incorrect number of top menu items, expected: %d, got: %d", len(expectedTopItems), len(topItems))
	}
	for i, item := range topItems {
		if item != expectedTopItems[i] {
			t.Errorf("incorrect top menu item at position %d, expected: %s, got: %s", i+1, expectedTopItems[i], item)
		}
	}

	// Test getting top 5 items (more than available)
	allItems := []string{"menu3", "menu1", "menu5", "menu2", "menu4"}
	topItems = menuCounter.GetTopMenuItems(5)
	if len(topItems) != len(allItems) {
		t.Errorf("incorrect number of top menu items, expected: %d, got: %d", len(allItems), len(topItems))
	}
	for i, item := range topItems {
		if item != allItems[i] {
			t.Errorf("incorrect top menu item at position %d, expected: %s, got: %s", i+1, allItems[i], item)
		}
	}

	// Test getting top 0 items
	topItems = menuCounter.GetTopMenuItems(0)
	if len(topItems) != 0 {
		t.Errorf("incorrect number of top menu items, expected: 0, got: %d", len(topItems))
	}
}

func TestMain(m *testing.M) {
	// Create test log files
	testLogFile := "testlog.txt"
	testLogDuplicateFile := "testlog_duplicate.txt"

	defer func() {
		// Remove test log files after running tests
		os.Remove(testLogFile)
		os.Remove(testLogDuplicateFile)
	}()

	// Create a test log file
	testLogContent := "eater1,menu1\n" +
		"eater2,menu2\n" +
		"eater3,menu2\n" +
		"eater4,menu3\n" +
		"eater5,menu2\n" +
		"eater6,menu1\n"
	writeTestLog(testLogFile, testLogContent)

	// Create a test log file with a duplicate entry
	testLogDuplicateContent := "eater1,menu1\n" +
		"eater2,menu2\n" +
		"eater3,menu2\n" +
		"eater4,menu3\n" +
		"eater5,menu2\n" +
		"eater6,menu1\n" +
		"eater7,menu1\n" // Duplicate entry
	writeTestLog(testLogDuplicateFile, testLogDuplicateContent)

	// Run the tests
	os.Exit(m.Run())
}

func writeTestLog(filename, content string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		panic(err)
	}
}
