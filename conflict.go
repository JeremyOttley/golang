package main

import (
	"encoding/xml"
	"fmt"
)

type Conflict struct {
	XMLName  xml.Name `xml:"conflict"`
	Text     string   `xml:",chardata"`
	ID       string   `xml:"id,attr"`
	Created  string   `xml:"created,attr"`
	CauseID  string   `xml:"causeID,attr"`
	OtherIDs string   `xml:"otherIDs,attr"`
	DoiData  []struct {
		Text     string `xml:",chardata"`
		Doi      string `xml:"doi"`
		Metadata struct {
			Text         string `xml:",chardata"`
			JournalTitle string `xml:"journal_title"`
			Volume       string `xml:"volume"`
			Issue        string `xml:"issue"`
			FirstPage    string `xml:"first_page"`
			Year         string `xml:"year"`
			ArticleTitle string `xml:"article_title"`
		} `xml:"metadata"`
		OtherConflicts struct {
			Text     string `xml:",chardata"`
			Conflict struct {
				Text   string `xml:",chardata"`
				ID     string `xml:"id,attr"`
				Status string `xml:"status,attr"`
			} `xml:"conflict"`
		} `xml:"other_conflicts"`
	} `xml:"doi_data"`
}


func (c Conflict) String() string {
	return fmt.Sprintf("Cause ID=%v",
		c.CauseID)
}

func main() {

	data := `<conflict id="6457489" created="2022-03-28 16:23:58.0" causeID="1525000047" otherIDs="1245940516,">
  <doi_data>
    <doi>10.3098/ah.2010.84.4.423</doi>
    <metadata>
      <journal_title>Agricultural History</journal_title>
      <volume>84</volume>
      <issue>4</issue>
      <first_page>423</first_page>
      <year>2010</year>
      <article_title>Origins of Pioneer Apple Orchards in the American West: Random Seeding versus Artisan Horticulture</article_title>
    </metadata>
    <other_conflicts>
      <conflict id="6457489" status="N"/>
    </other_conflicts>
  </doi_data>
  <doi_data>
    <doi>10.1215/00021482-84.4.423</doi>
    <metadata/>
    <other_conflicts>
      <conflict id="6457489" status="N"/>
    </other_conflicts>
  </doi_data>
</conflict>`

	var c Conflict
	if err := xml.Unmarshal([]byte(data), &c); err != nil {
		panic(err)
	}
	fmt.Println(c)
}

