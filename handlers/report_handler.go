package handlers

import (
	"fmt"
	"net/http"
	"os/exec"
	"path/filepath"
	"strings"
)

func DownloadReport(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("C:\\Users\\USER\\AppData\\Local\\Programs\\Python\\Python314\\python.exe", "print_report.py")
	out, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, "Report generation failed", http.StatusInternalServerError)
		return
	}

	outputStr := string(out)
	var filePath string
	lines := strings.Split(outputStr, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "READY:") {
			filePath = strings.TrimSpace(strings.TrimPrefix(line, "READY:"))
			break
		}
	}

	if filePath == "" {
		http.Error(w, "Could not find file path", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filepath.Base(filePath)))
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	http.ServeFile(w, r, filePath)
}
