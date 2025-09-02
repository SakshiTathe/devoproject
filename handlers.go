package main

import (
    "net/http"
    "strconv"
    "strings"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	err := pageTmpl.Execute(w, struct{ Items []Item }{Items: items})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	text := strings.TrimSpace(r.FormValue("text"))
	if text != "" {
		mu.Lock()
		items = append(items, Item{ID: nextID, Text: text})
		nextID++
		mu.Unlock()
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	idStr := r.FormValue("id")
	id, _ := strconv.Atoi(idStr)

	mu.Lock()
	defer mu.Unlock()
	for i := range items {
		if items[i].ID == id {
			items = append(items[:i], items[i+1:]...)
			break
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
