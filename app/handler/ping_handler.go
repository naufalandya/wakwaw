package handler

import (
	"belajar/app/db"
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
)

type savedNumberDB struct {
	ID             int
	Alias          sql.NullString
	CustomerNumber sql.NullString
	CustomerName   sql.NullString
	ProductName    sql.NullString
	CompanyID      int
	CreatedAt      sql.NullString
	CreatedBy      sql.NullString
	UpdatedAt      sql.NullString
	UpdatedBy      sql.NullString
}

func GetSavedNumbersWow(c *fiber.Ctx) error {
	query := `
        SELECT 
            ID,
            ALIAS,
            CUSTOMER_NUMBER,
            CUSTOMER_NAME,
            PRODUCT_NAME,
            COMPANY_ID,
            CREATED_AT,
            CREATED_BY,
            UPDATED_AT,
            UPDATED_BY
        FROM ADDONS_EWALLET_PUBLIC.SAVED_NUMBERS
    `

	rows, err := db.Oracle.Query(query)
	if err != nil {
		log.Println("Query failed:", err)
		return c.Status(500).JSON(fiber.Map{"error": "query failed"})
	}
	defer rows.Close()

	var result []savedNumberDB

	for rows.Next() {
		var sn savedNumberDB

		err := rows.Scan(
			&sn.ID,
			&sn.Alias,
			&sn.CustomerNumber,
			&sn.CustomerName,
			&sn.ProductName,
			&sn.CompanyID,
			&sn.CreatedAt,
			&sn.CreatedBy,
			&sn.UpdatedAt,
			&sn.UpdatedBy,
		)
		if err != nil {
			log.Println("Scan failed:", err)
			continue
		}

		result = append(result, sn)
	}

	return c.JSON(fiber.Map{
		"status": "ok",
		"total":  len(result),
		"data":   result,
	})
}
