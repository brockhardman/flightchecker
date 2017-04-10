// // main.go
// package main

// import (
//     "github.com/astaxie/beego"
//     "bytes"
//     "net/http"
//     "io/ioutil"
//     "strings"
//     "encoding/json" 
//     "fmt"
//     "time"
//     "strconv"
// )

// //GLOBAL VARIABLES AND CONSTANTS
// var searchParams = &Params{}
// const apiUrl string = "https://www.googleapis.com/qpxExpress/v1/trips/search?key=AIzaSyDNqowpRs4ctjCp-EYzfzlcrQuY9lT-CGI"

// /*********************
// BEGIN STRUCTS SECTION
// *********************/

// // This is the controller that this application uses
// type mainController struct {
//     beego.Controller
// }

// type TripResponse struct {
//     Trips Trips `json:"trips"`
// }

// type Data struct {  
//     Airpoirt []Airport `json:"airport"`
//     City []City `json:"city"`
//     Aircraft []Aircraft `json:"aircraft"`
//     Tax []Tax `json:"tax"`
//     Carrier []Carrier `json:"carrier"`
// }

// type Airport struct {
//     Code string `json:"code"`
//     City string `json:"city"`
//     Name string `json:"name"`
// }

// type City struct {
//     Code string `json:"code"`
//     Name string `json:"name"`
// }

// type Aircraft struct {
//     Code string `json:"code"`
//     Name string `json:"name"`
// }

// type Tax struct {
//     Id string `json:"id"`
//     Name string `json:"name"`
// }

// type Carrier struct {
//     Code string `json:"code"`
//     Name string `json:"name"`
// }

// type Trips struct {
//     RequestId string `json:"requestId"`
//     Data Data `json:"data"`
//     TripOption []TripOptionEntry `json:"tripOption"`
// }

// type TripOptionEntry struct {
//     SaleTotal string `json:"saleTotal"`
//     Slice []Slice `json:"slice"`
//     Pricing []Pricing `json:"pricing"`
// }

// type Slice struct {
//     Duration int `json:"duration"`
//     Segment []Segment `json:"segment"`
// }

// type Segment struct {
//     Flight Flight `json:"flight"`
//     Leg []Leg `json:"leg"`
// }

// type Leg struct {
//     ArrivalTime string `json:"arrivalTime"`
//     DepartureTime string `json:"departureTime"`
//     Origin string `json:"origin"`
//     Destination string `json:"destination"`
//     Duration int `json:"duration"`
//     DurationString string
//     MealInfo string `json:"meal"`
// }

// type Pricing struct {
//     BaseFareTotal string `json:"baseFareTotal"`
//     SaleFareTotal string `json:"saleFareTotal"`
//     SaleTaxTotal string `json:"saleTaxTotal"`
//     SaleTotal string `json:"saleTotal"`
// }

// type Flight struct {
//     Number string `json:"number"`
//     Carrier string `json:"carrier"`
// }

// type Params struct {
//     From string `form:"departureAirport"`
//     To string `form:"destinationAirport"`
//     FromDate string `form:"departureDate"`
//     ToDate string `form:"returnDate"`
// }

// /*******************
// END STRUCTS SECTION
// *******************/


// func main() {
//     beego.Router("/:from/:to/:fromDate/:toDate", &mainController{}, "get:URLSearch")
//     beego.Router("/search", &mainController{}, "post:FormSearch")
//     beego.Router("/", &mainController{})
//     beego.Run()
// }

// func (this *mainController) Get() {
//     this.updateFlightData()
// }

// //This will be called when the user passes the params in the URL
// func (this *mainController) URLSearch() {
//     searchParams.From = this.Ctx.Input.Param(":from")
//     searchParams.To = this.Ctx.Input.Param(":to")
//     searchParams.FromDate = this.Ctx.Input.Param(":fromDate")
//     searchParams.ToDate = this.Ctx.Input.Param(":toDate")
//     this.updateFlightData()
// }

// //Generic method to update the flight information on the screen
// func (this *mainController) updateFlightData() {
//     this.TplName = "result.html"
//     //Build request/response objects 
//     var jsonRaw = `{"request": {"slice": [{"origin": "ORIGIN_CODE","destination": "DESTINATION_CODE","date": "FROM_DATE_STRING"},{"origin": "DESTINATION_CODE","destination": "ORIGIN_CODE","date": "TO_DATE_STRING"}],"passengers": {"adultCount": 1,"infantInLapCount": 0,"infantInSeatCount": 0,"childCount": 0,"seniorCount": 0},"solutions": 100,"refundable": true}}`
//     jsonRaw = strings.Replace(jsonRaw, "ORIGIN_CODE", searchParams.From, 2)
//     jsonRaw = strings.Replace(jsonRaw, "DESTINATION_CODE", searchParams.To, 2)
//     jsonRaw = strings.Replace(jsonRaw, "FROM_DATE_STRING", searchParams.FromDate, 1)
//     jsonRaw = strings.Replace(jsonRaw, "TO_DATE_STRING", searchParams.ToDate, 1)
//     var jsonString = []byte(jsonRaw)
    
//     req,err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonString))
//     if err != nil {
//         panic(err)
//     }
//     req.Header.Add("Content-Type", "application/json")
    
//     client := http.Client{}
//     resp,err := client.Do(req)
//     defer resp.Body.Close()

//     body, _ := ioutil.ReadAll(resp.Body)

//     //Parse response into meaningful representation
//     tripsData, error := generateTripObjectWithFlightsData([]byte(body))
//     if(error != nil){
//         fmt.Println("error:", error)
//     }

//     //This will populate the data that will be used in the HTML results table
//     this.Data["flightData"] = this.sanitizeFlightData(tripsData.Trips)
// }

// //Various data elements need to be formatted properly
// //TODO: MAKE THIS BETTER!
// func (this *mainController) sanitizeFlightData(TripsIncoming Trips) (TripsOutgoing Trips) {

//     for i, tripOption := range TripsIncoming.TripOption {
//         tripOptionPointer := &TripsIncoming.TripOption[i]

//         for _, entry := range tripOption.Slice {
//             for _, segment := range entry.Segment {
//                 for i2, leg := range segment.Leg {
//                     legPointer := &segment.Leg[i2]
//                     time1, err1 := time.Parse("2006-01-02T15:04-07:00", leg.DepartureTime)
//                     time2, err2 := time.Parse("2006-01-02T15:04-07:00", leg.ArrivalTime)
//                     layout := "3:04PM Mon Jan _2"
                    
//                     if(err1 != nil){
//                         fmt.Println("error:", err1)
//                     }
//                     if(err2 != nil){
//                         fmt.Println("error:", err2)
//                     }
                    
//                     legPointer.DepartureTime = time1.Format(layout)
//                     legPointer.ArrivalTime = time2.Format(layout)
//                     hours := int(leg.Duration / 60)
//                     minutes := int(leg.Duration % 60)
//                     legPointer.DurationString = strconv.Itoa(hours) + "h " + strconv.Itoa(minutes) + "m"
//                 }    
//             }
//         } 

//         tripOptionPointer.SaleTotal = strings.Replace(tripOptionPointer.SaleTotal, "USD", "$", 1)    
//     }

//     return TripsIncoming
// }

// //Parse JSON body into model objects (Structs) defined above
// func generateTripObjectWithFlightsData(body []byte) (*TripResponse, error) {
//     var responseObj = new(TripResponse)
//     error := json.Unmarshal(body, &responseObj)
//     if(error != nil){
//         fmt.Println("error:", error)
//     }
//     return responseObj, error
// }

// //This will be called when the user submits the form
// func (this *mainController) FormSearch() {
//     params := Params{}
//     error := this.ParseForm(&params) 
//     if error != nil {
//         fmt.Println("error:", error)
//     }
//     searchParams.From = params.From
//     searchParams.To = params.To
//     searchParams.FromDate = params.FromDate
//     searchParams.ToDate = params.ToDate
//     this.updateFlightData()
// }
