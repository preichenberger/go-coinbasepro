package coinbase

import(
  "fmt"
  "time"
)

type ReportParams struct {
  StartDate time.Time
  EndDate time.Time
}

type CreateReportParams struct {
  Start time.Time
  End time.Time
}


type Report struct {
  Type string `json:"type"`
  Status string `json:"status"`
  CreatedAt time.Time `json:"created_at"`
  CompletedAt time.Time `json:"created_at"`
  ExpiresAt time.Time `json:"expires_at"`
  FileURL string `json:"file_url"`
  Params ReportParams `json:"params"`
  StartDate time.Time
  EndDate time.Time
}

func (c *Client) CreateReport(newReport *Report) (Report, error) {
  var savedReport Report

  url := fmt.Sprintf("/reports")
  _, err := c.Request("POST", url, newReport, &savedReport)

  return savedReport, err
}

func (c *Client) GetReportStatus(id int) (Report, error) {
  report := Report{}

  url := fmt.Sprintf("/reports/%s", id)
  _, err := c.Request("GET", url, nil, &report)

  return report, err
}
