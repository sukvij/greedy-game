# greedy-game



gorm --> postman --> field --> map with model --> map with database
gorm --> serial is not working still.


query := dbcon
query = query.Where("name = ?", "John").Limit(10).Offset(5).Order("name ASC")

queryString := dbcon.ToSQL(func(tx *gorm.DB) *gorm.DB) {
  var users []User{}
  return tx.Find(&user)
})
fmt.Println("The query is ", queryString)