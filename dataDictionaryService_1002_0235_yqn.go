// 代码生成时间: 2025-10-02 02:35:32
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "log"
)

// DataDictionaryEntry represents a single entry in the data dictionary.
type DataDictionaryEntry struct {
    ID    string `json:"id"`
    Key   string `json:"key"`
    Value string `json:"value"`
}

// DataDictionaryService manages the operations on the data dictionary entries.
type DataDictionaryService struct {
    entries map[string]DataDictionaryEntry
}

// NewDataDictionaryService creates a new instance of DataDictionaryService.
func NewDataDictionaryService() *DataDictionaryService {
    return &DataDictionaryService{
        entries: make(map[string]DataDictionaryEntry),
    }
}

// AddEntry adds a new entry to the data dictionary.
func (s *DataDictionaryService) AddEntry(entry DataDictionaryEntry) error {
    if _, exists := s.entries[entry.ID]; exists {
        return fmt.Errorf("entry with ID '%s' already exists", entry.ID)
    }
    s.entries[entry.ID] = entry
    return nil
}

// UpdateEntry updates an existing entry in the data dictionary.
func (s *DataDictionaryService) UpdateEntry(entry DataDictionaryEntry) error {
    if _, exists := s.entries[entry.ID]; !exists {
        return fmt.Errorf("entry with ID '%s' does not exist", entry.ID)
    }
    s.entries[entry.ID] = entry
    return nil
}

// DeleteEntry removes an entry from the data dictionary.
func (s *DataDictionaryService) DeleteEntry(id string) error {
    if _, exists := s.entries[id]; !exists {
        return fmt.Errorf("entry with ID '%s' does not exist", id)
    }
    delete(s.entries, id)
    return nil
}

// GetEntry retrieves an entry from the data dictionary by ID.
func (s *DataDictionaryService) GetEntry(id string) (DataDictionaryEntry, error) {
    entry, exists := s.entries[id]
    if !exists {
        return DataDictionaryEntry{}, fmt.Errorf("entry with ID '%s' does not exist", id)
    }
    return entry, nil
}

// GetEntries retrieves all entries from the data dictionary.
func (s *DataDictionaryService) GetEntries() []DataDictionaryEntry {
    entries := make([]DataDictionaryEntry, 0, len(s.entries))
    for _, entry := range s.entries {
        entries = append(entries, entry)
    }
    return entries
}

// SetupRoutes sets up the HTTP routes for the data dictionary service.
func SetupRoutes(app *fiber.App, service *DataDictionaryService) {
    app.Post("/entries", func(c *fiber.Ctx) error {
        var entry DataDictionaryEntry
        if err := c.BodyParser(&entry); err != nil {
            return err
        }
        if err := service.AddEntry(entry); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.Status(fiber.StatusCreated).JSON(entry)
    })

    app.Put("/entries/:id", func(c *fiber.Ctx) error {
        id := c.Params("id")
        var entry DataDictionaryEntry
        if err := c.BodyParser(&entry); err != nil {
            return err
        }
        entry.ID = id
        if err := service.UpdateEntry(entry); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.Status(fiber.StatusOK).JSON(entry)
    })

    app.Delete("/entries/:id", func(c *fiber.Ctx) error {
        id := c.Params("id")
        if err := service.DeleteEntry(id); err != nil {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.SendStatus(fiber.StatusOK)
    })

    app.Get("/entries/:id", func(c *fiber.Ctx) error {
        id := c.Params("id\)
        entry, err := service.GetEntry(id)
        if err != nil {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.Status(fiber.StatusOK).JSON(entry)
    })

    app.Get("/entries", func(c *fiber.Ctx) error {
        entries := service.GetEntries()
        return c.Status(fiber.StatusOK).JSON(entries)
    })
}

func main() {
    app := fiber.New()
    service := NewDataDictionaryService()
    SetupRoutes(app, service)
    log.Fatal(app.Listen(":3000"))
}