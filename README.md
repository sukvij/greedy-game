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
  run .sql file comamnd
  docker run -d -p 6831:6831/udp -p 16686:16686 jaegertracing/all-in-one:latest
  run      http://localhost:16686/   --> jaeger

# create campaign
  http://localhost:8080/v1/campaign

  
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

# create targeting rules
  http://localhost:8080/v1/targeting-rule
 
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



  
  # get all deliveries
  http://localhost:8080/v1/delivery?app=com.duolingo.ludokinggame&country=US&os=Android
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



