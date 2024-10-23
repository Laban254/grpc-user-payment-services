// database/models.go
package database

type User struct {
    ID    uint   `gorm:"primaryKey"`
    Name  string `gorm:"size:100"`
    Email string `gorm:"uniqueIndex;size:100"`
}

type Transaction struct {
    ID        int32   `gorm:"primaryKey" json:"id"`      
    UserID    int32   `gorm:"not null" json:"user_id"`      
    Amount    float64 `gorm:"not null" json:"amount"`      
    Status    string  `gorm:"not null" json:"status"`     
}
