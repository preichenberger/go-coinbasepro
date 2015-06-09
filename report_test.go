package coinbase

import(
  "testing"
  "time"
)

func TestCreateReportAndStatus(t *testing.T) {
  client := NewTestClient()
  newReport := Report{
    Type: "fill",
    StartDate: time.Now().Add(-24 * 4  * time.Hour),
    EndDate: time.Now().Add(-24 * 2 * time.Hour),
  }

  report, err := client.CreateReport(&newReport)
  if err != nil {
    t.Error(err)
  }

  println(report.Status)
}
