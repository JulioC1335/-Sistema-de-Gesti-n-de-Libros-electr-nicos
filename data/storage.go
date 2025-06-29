package data

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"biblioteca/models"
)

// Thread-safe in-memory storage with file persistence
var (
	mu      sync.RWMutex
	Books   []models.Book
	Users   []models.User
	Loans   []models.Loan
	BookID  int
	UserID  int
	LoanID  int
	dataDir = "data"
)

// Initialize loads data from JSON files, seeds default user, and sets counters
func Initialize() error {
	mu.Lock()
	defer mu.Unlock()

	// Load existing data
	if err := loadFromFile("books.json", &Books); err != nil {
		return fmt.Errorf("load books: %w", err)
	}
	if err := loadFromFile("users.json", &Users); err != nil {
		return fmt.Errorf("load users: %w", err)
	}
	// Seed default admin user if none exist
	if len(Users) == 0 {
		admin := models.User{
			ID:       NextUserID(),
			Name:     "admin",
			Password: "admin",
		}
		Users = append(Users, admin)
		if err := saveToFile("users.json", Users); err != nil {
			return fmt.Errorf("seed default user: %w", err)
		}
	}
	if err := loadFromFile("loans.json", &Loans); err != nil {
		return fmt.Errorf("load loans: %w", err)
	}

	// Set ID counters based on existing data
	for _, b := range Books {
		if b.ID > BookID {
			BookID = b.ID
		}
	}
	for _, u := range Users {
		if u.ID > UserID {
			UserID = u.ID
		}
	}
	for _, l := range Loans {
		if l.ID > LoanID {
			LoanID = l.ID
		}
	}
	return nil
}

// SaveAll persists all slices to their JSON files
func SaveAll() error {
	mu.RLock()
	defer mu.RUnlock()

	if err := saveToFile("books.json", Books); err != nil {
		return err
	}
	if err := saveToFile("users.json", Users); err != nil {
		return err
	}
	if err := saveToFile("loans.json", Loans); err != nil {
		return err
	}
	return nil
}

// saveToFile writes data to filename under dataDir
func saveToFile(filename string, data interface{}) error {
	path := filepath.Join(dataDir, filename)
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create %s: %w", path, err)
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(data); err != nil {
		return fmt.Errorf("encode %s: %w", path, err)
	}
	return nil
}

// loadFromFile reads JSON into out, ignores missing file
func loadFromFile(filename string, out interface{}) error {
	path := filepath.Join(dataDir, filename)
	f, err := os.Open(path)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("open %s: %w", path, err)
	}
	defer f.Close()
	if err := json.NewDecoder(f).Decode(out); err != nil {
		return fmt.Errorf("decode %s: %w", path, err)
	}
	return nil
}

// Helper functions to allocate IDs and append
func NextBookID() int {
	mu.Lock()
	defer mu.Unlock()
	BookID++
	return BookID
}

func AddBook(b models.Book) error {
	b.ID = NextBookID()
	mu.Lock()
	Books = append(Books, b)
	mu.Unlock()
	return saveToFile("books.json", Books)
}

func DeleteBookByID(id, ownerID int) error {
	mu.Lock()
	defer mu.Unlock()
	newList := make([]models.Book, 0, len(Books))
	for _, b := range Books {
		if !(b.ID == id && b.UserID == ownerID) {
			newList = append(newList, b)
		}
	}
	Books = newList
	return saveToFile("books.json", Books)
}

func NextUserID() int {
	mu.Lock()
	defer mu.Unlock()
	UserID++
	return UserID
}

func AddUser(u models.User) error {
	u.ID = NextUserID()
	mu.Lock()
	Users = append(Users, u)
	mu.Unlock()
	return saveToFile("users.json", Users)
}

func DeleteUserByID(id int) error {
	mu.Lock()
	defer mu.Unlock()
	newList := make([]models.User, 0, len(Users))
	for _, u := range Users {
		if u.ID != id {
			newList = append(newList, u)
		}
	}
	Users = newList
	return saveToFile("users.json", Users)
}

func NextLoanID() int {
	mu.Lock()
	defer mu.Unlock()
	LoanID++
	return LoanID
}

func AddLoan(l models.Loan) error {
	l.ID = NextLoanID()
	mu.Lock()
	Loans = append(Loans, l)
	mu.Unlock()
	return saveToFile("loans.json", Loans)
}

func DeleteLoanByID(id int) error {
	mu.Lock()
	defer mu.Unlock()
	newList := make([]models.Loan, 0, len(Loans))
	for _, l := range Loans {
		if l.ID != id {
			newList = append(newList, l)
		}
	}
	Loans = newList
	return saveToFile("loans.json", Loans)
}
