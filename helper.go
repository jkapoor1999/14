package main

import (	
	"fmt"	
	"strings"
)

func first(pattern string, lst []File) []string {
	res := []string{}
	
	for _, file := range lst{
		for ind, line := range file.lines{
			if strings.Contains(line, pattern) {
				if len(lst) > 1{
					res = append(res, file.name + ":" + fmt.Sprint(ind + 1) + " " + line)
				}else{
					res = append(res, fmt.Sprint(ind + 1) + " " + line)
				}
				
			}
		}
	}
	return res
}

func second(pattern string, lst []File) []string {
	res := []string{}
	for _, file := range lst{
		for _, line := range file.lines{
			if strings.Contains(line, pattern){
				res = append(res, file.name)
				break
			}
		}
	}
	return res
}

func third(pattern string, lst []File) []string{
	res := []string{}
	for _, file := range lst{
		for _, line := range file.lines{
			if strings.Contains(strings.ToLower(line), strings.ToLower(pattern)) {
				if len(lst) > 1{
					res = append(res, file.name + ":" + " " + line)
				}else{
					res = append(res, line)
				}
				
			}
		}
	}
	return res
}

func fourth(pattern string, lst []File) []string {
	res := []string{}
	for _, file := range lst{
		for _, line := range file.lines{
			if !strings.Contains(line, pattern) {
				if len(lst) > 1{
					res = append(res, file.name + ":" + " " + line)
				}else{
					res = append(res, " " + line)
				}				
			}
		}
	}
	return res
}

func fifth(pattern string, lst []File) []string {
	res := []string{}
	for _, file := range lst{
		for _, line := range file.lines{
			if line == pattern {
				if len(lst) > 1{
					res = append(res, file.name + ":" + " " + line)
				}else{
					res = append(res, " " + line)
				}				
			}
		}
	}
	return res
}

func stringSlicesEqual(a, b []string) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}