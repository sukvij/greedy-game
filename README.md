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





# instruction to run this app
  start redis
  go run main.go


  http://localhost:8080/campaign
  # create campaign
              // {
            //     "cid":"spotify",
            //     "campaign_name":"Spotify - Music for everyone",
            //     "img":"https://somelink",
            //     "cta":"Download",
            //     "status":"ACTIVE"
            // }


            // {
            //     "cid":"duolingo",
            //     "campaign_name":"Duolingo: Best way to learn",
            //     "img":"https://somelink2",
            //     "cta":"Install",
            //     "status":"ACTIVE"
            //}
            {
                "cid":"subwaysurfer",
                "campaign_name":"Subway Surfer",
                "img":"https://somelink3",
                "cta":"Play",
                "status":"ACTIVE"
            }


  http://localhost:8080/targeting-rule
  # create targeting rules
            {
          "cid":"duolingo",
          "rules": {
                  "include_country": ["US", "Canada"],
                  // "exclude_country": ["US"],
                  "include_os": ["Android"],
                  // "exclude_os": ["Windows"],
                  "include_app": ["com.duolingo.ludokinggame"]
                  // "exclude_app": ["App3"]
          }
      }



  http://localhost:8080/delivery?app=com.duolingo.ludokinggame&country=US&os=Android
  # get all deliveries
  response will be like this

  {
    "success": true,
    "statusCode": 200,
    "data": [
        {
            "cid": "duolingo",
            "campaign_name": "Duolingo: Best way to learn",
            "img": "https://somelink2",
            "cta": "Install"
        }
    ],
    "meta": {
        "version": "",
        "latencyMs": 12
    }
}



