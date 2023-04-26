package main

import "testing"

func TestGetTitle(t *testing.T) {
	content1 := "<html><head><title>Example Title</title></head><body><p>Example Paragraph</p></body></html>"
	expected1 := "Example Title"
	result1 := getTitle(content1)
	if result1 != expected1 {
		t.Errorf("Test case 1 failed: Expected '%s', but got '%s' ", expected1, result1)
	}

	content2 := "<html><head>Example Title</head><body><p>Example Paragraph</p></body></html>"
	expected2 := "no title"
	result2 := getTitle(content2)
	if result2 != expected2 {
		t.Errorf("Test case 2 failed: Expected '%s', but got '%s' ", expected2, result2)
	}

	content3 := ""
	expected3 := "no title"
	result3 := getTitle(content3)
	if result3 != expected3 {
		t.Errorf("Test case 3 failed: Expected '%s', but got '%s' ", expected3, result3)
	}
}
