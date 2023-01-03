package main

func main() {
	for _, line := range Search("Hello", "-i", []string{"input.txt", "greeting.txt"}) {
		println(line)
	}
}

func Search(pattern string, flag string, files []string) []string {
	lst := []File{}
	
	for _, filename := range files{
		lst = append(lst, createFile(filename))		
	}

	switch flag {
	case "-n":
		return first(pattern, lst);
	case "-l":
		return second(pattern, lst);
	case "-i":
		return third(pattern, lst);
	case "-v":
		return fourth(pattern, lst);
	case "-x":
		return fifth(pattern, lst);
	default:
		println("Please enter a correct flag");
	}
	return []string{};	
}