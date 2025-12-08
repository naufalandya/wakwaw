package handler

import (
	"belajar/app/db"
	"database/sql"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserDB struct {
	ID             int
	Username       sql.NullString
	Email          sql.NullString
	CreatedAt      sql.NullTime
	FullName       sql.NullString
	Bio            sql.NullString
	ProfilePicture sql.NullString
	UpdatedAt      sql.NullTime
	DeletedAt      sql.NullTime
	IsPirate       sql.NullInt64
	NewestBounty   sql.NullInt64
	PreviousBounty sql.NullInt64
}

// Struct untuk RESPONSE agar JSON rapi
type UserResponse struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	CreatedAt      string `json:"created_at"`
	FullName       string `json:"full_name"`
	Bio            string `json:"bio"`
	ProfilePicture string `json:"profile_picture"`
	UpdatedAt      string `json:"updated_at"`
	DeletedAt      string `json:"deleted_at"`
	IsPirate       int64  `json:"is_pirate"`
	NewestBounty   int64  `json:"newest_bounty"`
	PreviousBounty int64  `json:"previous_bounty"`
}

// ===============================
// Helper konversi Null Types
// ===============================

func ns(v sql.NullString) string {
	if v.Valid {
		return v.String
	}
	return ""
}

func ni(v sql.NullInt64) int64 {
	if v.Valid {
		return v.Int64
	}
	return 0
}

func nt(v sql.NullTime) string {
	if v.Valid {
		return v.Time.Format(time.RFC3339)
	}
	return ""
}

// ===============================
// Handler utama
// ===============================

func GetUsersWow(c *fiber.Ctx) error {

	// Ambil query params
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 20)

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 20
	}

	offset := (page - 1) * limit

	// Query total count
	var total int
	err := db.Oracle.QueryRow(`SELECT COUNT(*) FROM WOWAA.USERS`).Scan(&total)
	if err != nil {
		log.Println("Count query failed:", err)
		return c.Status(500).JSON(fiber.Map{"error": "count query failed"})
	}

	// Query dengan pagination
	query := `
		SELECT 
			ID,
			USERNAME,
			EMAIL,
			CREATED_AT,
			FULL_NAME,
			BIO,
			PROFILE_PICTURE_URL,
			UPDATED_AT,
			DELETED_AT,
			IS_PIRATE,
			NEWEST_BOUNTY,
			PREVIOUS_BOUNTY
		FROM WOWAA.USERS
		OFFSET :offset ROWS FETCH NEXT :limit ROWS ONLY
	`

	rows, err := db.Oracle.Query(query, offset, limit)
	if err != nil {
		log.Println("Query failed:", err)
		return c.Status(500).JSON(fiber.Map{"error": "query failed"})
	}
	defer rows.Close()

	var dbResult []UserDB

	for rows.Next() {
		var u UserDB

		err := rows.Scan(
			&u.ID,
			&u.Username,
			&u.Email,
			&u.CreatedAt,
			&u.FullName,
			&u.Bio,
			&u.ProfilePicture,
			&u.UpdatedAt,
			&u.DeletedAt,
			&u.IsPirate,
			&u.NewestBounty,
			&u.PreviousBounty,
		)
		if err != nil {
			log.Println("Scan failed:", err)
			continue
		}

		dbResult = append(dbResult, u)
	}

	// ===============================
	// CONVERT KE RESPONSE NORMAL
	// ===============================
	var result []UserResponse

	for _, u := range dbResult {
		result = append(result, UserResponse{
			ID:             u.ID,
			Username:       ns(u.Username),
			Email:          ns(u.Email),
			CreatedAt:      nt(u.CreatedAt),
			FullName:       ns(u.FullName),
			Bio:            ns(u.Bio),
			ProfilePicture: ns(u.ProfilePicture),
			UpdatedAt:      nt(u.UpdatedAt),
			DeletedAt:      nt(u.DeletedAt),
			IsPirate:       ni(u.IsPirate),
			NewestBounty:   ni(u.NewestBounty),
			PreviousBounty: ni(u.PreviousBounty),
		})
	}

	return c.JSON(fiber.Map{
		"status": "ok",
		"page":   page,
		"limit":  limit,
		"total":  total,
		"data":   result,
	})
}
