package load

import (
	"encoding/json"
	"fmt"
	"os"
)

func File(fileName string) ([]map[string]int, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("error when opening file: %s", err)
	}

	var payload []map[string]int
	if err = json.Unmarshal(content, &payload); err != nil {
		return nil, fmt.Errorf("error during Unmarshal(): %s", err)
	}

	return payload, nil
}
